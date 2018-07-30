package executor

import (
	"fmt"
	"io"
	"os"
	"runtime/pprof"
	"sync"
	"time"

	"github.com/vmihailenco/msgpack"
	"github.com/xitongsys/guery/eplan"
	"github.com/xitongsys/guery/gtype"
	"github.com/xitongsys/guery/metadata"
	"github.com/xitongsys/guery/pb"
	"github.com/xitongsys/guery/row"
	"github.com/xitongsys/guery/util"
)

func (self *Executor) SetInstructionShuffle(instruction *pb.Instruction) (err error) {
	var enode eplan.EPlanShuffleNode
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

func (self *Executor) RunShuffle() (err error) {
	fname := fmt.Sprintf("executor_%v_hashjoinshuffle_%v_cpu.pprof", self.Name, time.Now().Format("20060102150405"))
	f, _ := os.Create(fname)
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	defer func() {
		if err != nil {
			self.AddLogInfo(err, pb.LogLevel_ERR)
		}
		self.Clear()
	}()
	enode := self.EPlanNode.(*eplan.EPlanShuffleNode)
	//read md
	md := &metadata.Metadata{}
	for _, reader := range self.Readers {
		if err = util.ReadObject(reader, md); err != nil {
			return err
		}
	}

	mdOutput := md.Copy()

	//write md
	if enode.Keys != nil && len(enode.Keys) > 0 {
		mdOutput.ClearKeys()
		mdOutput.AppendKeyByType(gtype.STRING)
	}
	for _, writer := range self.Writers {
		if err = util.WriteObject(writer, mdOutput); err != nil {
			return err
		}
	}

	rbWriters := make([]*row.RowsBuffer, len(self.Writers))
	for i, writer := range self.Writers {
		rbWriters[i] = row.NewRowsBuffer(mdOutput, nil, writer)
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
	var wg sync.WaitGroup
	for i, _ := range self.Readers {
		wg.Add(1)
		go func(index int) {
			defer func() {
				wg.Done()
			}()
			reader := self.Readers[index]
			rbReader := row.NewRowsBuffer(md, reader, nil)
			for {
				rg0, err := rbReader.Read()
				if err == io.EOF {
					break
				}
				if err != nil {
					self.AddLogInfo(err, pb.LogLevel_ERR)
					return
				}

				for i := 0; i < rg0.GetRowsNumber(); i++ {
					r := rg0.GetRow(i)
					index := 0
					if enode.Keys != nil && len(enode.Keys) > 0 {
						rg := row.NewRowsGroup(mdOutput)
						rg.Write(r)
						key, err := CalHashKey(enode.Keys, rg)
						if err != nil {
							self.AddLogInfo(err, pb.LogLevel_ERR)
							return
						}
						r.AppendKeys(key)
						index = ShuffleHash(key) % len(rbWriters)

					}

					if err = rbWriters[index].WriteRow(r); err != nil {
						self.AddLogInfo(err, pb.LogLevel_ERR)
						return
					}

					row.RowPool.Put(r)
				}
			}
		}(i)
	}

	wg.Wait()

	return nil
}
