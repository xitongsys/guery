package Executor

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime/pprof"
	"sync"
	"time"

	"github.com/vmihailenco/msgpack"
	"github.com/xitongsys/guery/EPlan"
	"github.com/xitongsys/guery/Metadata"
	"github.com/xitongsys/guery/Row"
	"github.com/xitongsys/guery/Type"
	"github.com/xitongsys/guery/Util"
	"github.com/xitongsys/guery/pb"
)

func (self *Executor) SetInstructionHashJoinShuffle(instruction *pb.Instruction) (err error) {
	var enode EPlan.EPlanHashJoinShuffleNode
	if err = msgpack.Unmarshal(instruction.EncodedEPlanNodeBytes, &enode); err != nil {
		return err
	}
	self.Instruction = instruction
	self.EPlanNode = &enode
	self.InputLocations = []*pb.Location{}
	for i := 0; i < len(enode.Inputs); i++ {
		loc := enode.Inputs[i]
		self.InputLocations = append(self.InputLocations, &loc)
	}
	self.OutputLocations = []*pb.Location{}
	for i := 0; i < len(enode.Outputs); i++ {
		loc := enode.Outputs[i]
		self.OutputLocations = append(self.OutputLocations, &loc)
	}
	return nil
}

func ShuffleHash(s string) int {
	res := 0
	for _, c := range []byte(s) {
		res += int(c)
	}
	return res
}

func (self *Executor) RunHashJoinShuffle() (err error) {
	fname := fmt.Sprintf("executor_%v_hashjoinshuffle_%v_cpu.pprof", self.Name, time.Now().Format("20060102150405"))
	f, _ := os.Create(fname)
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	defer self.Clear(err)
	enode := self.EPlanNode.(*EPlan.EPlanHashJoinShuffleNode)
	//read md
	md := &Metadata.Metadata{}
	for _, reader := range self.Readers {
		if err = Util.ReadObject(reader, md); err != nil {
			return err
		}
	}

	mdOutput := md.Copy()

	//write md
	if enode.Keys != nil && len(enode.Keys) > 0 {
		mdOutput.ClearKeys()
		mdOutput.AppendKeyByType(Type.STRING)
	}
	for _, writer := range self.Writers {
		if err = Util.WriteObject(writer, mdOutput); err != nil {
			return err
		}
	}

	rbWriters := make([]*Row.RowsBuffer, len(self.Writers))
	for i, writer := range self.Writers {
		rbWriters[i] = Row.NewRowsBuffer(mdOutput, nil, writer)
	}

	defer func() {
		for _, rbWriter := range rbWriters {
			rbWriter.Flush()
		}
	}()

	//init
	for _, k := range enode.Keys {
		if err := k.Init(md); err != nil {
			return err
		}
	}

	//write rows
	var rg0 *Row.RowsGroup
	var wg sync.WaitGroup
	for i, _ := range self.Readers {
		wg.Add(1)
		go func(index int) {
			defer func() {
				wg.Done()
			}()
			reader := self.Readers[index]
			rbReader := Row.NewRowsBuffer(md, reader, nil)
			for {
				rg0, err = rbReader.Read()
				if err == io.EOF {
					break
				}
				if err != nil {
					return
				}
				log.Println("=======", self.Name, rg0.GetRowsNumber())

				for i := 0; i < rg0.GetRowsNumber(); i++ {
					row := rg0.GetRow(i)
					index := 0
					if enode.Keys != nil && len(enode.Keys) > 0 {
						rg := Row.NewRowsGroup(mdOutput)
						rg.Write(row)
						key, err := CalHashKey(enode.Keys, rg)
						if err != nil {
							return
						}
						row.AppendKeys(key)
						index = ShuffleHash(key) % len(rbWriters)
					}

					if err = rbWriters[index].WriteRow(row); err != nil {
						return
					}
				}
			}
		}(i)
	}

	wg.Wait()

	return nil
}
