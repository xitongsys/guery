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

func (self *Executor) SetInstructionDistinct(instruction *pb.Instruction) (err error) {
	var enode eplan.EPlanDistinctNode
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

func (self *Executor) RunDistinct() (err error) {
	fname := fmt.Sprintf("executor_%v_hashjoindistinct_%v_cpu.pprof", self.Name, time.Now().Format("20060102150405"))
	f, _ := os.Create(fname)
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	defer func() {
		if err != nil {
			self.AddLogInfo(err, pb.LogLevel_ERR)
		}
		self.Clear()
	}()
	enode := self.EPlanNode.(*eplan.EPlanDistinctNode)
	//read md
	md := &metadata.Metadata{}
	for _, reader := range self.Readers {
		if err = util.ReadObject(reader, md); err != nil {
			return err
		}
	}

	mdOutput := enode.Metadata

	//write md
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
	for _, e := range enode.Expressions {
		if err := e.Init(md); err != nil {
			return err
		}
	}

	var mutex sync.Mutex
	distinctMap := make([]map[string]bool, len(enode.Expressions))
	for i := 0; i < len(enode.Expressions); i++ {
		distinctMap[i] = make(map[string]bool)
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

				distCols := make([][]interface{}, len(enode.Expressions))
				for i, e := range enode.Expressions {
					res, err := e.Result(rg0)
					if err != nil {
						self.AddLogInfo(err, pb.LogLevel_ERR)
						return
					}
					distCols[i] = res.([]interface{})
					mutex.Lock()
					for j, c := range distCols[i] {
						ckey := gtype.ToKeyString(c)
						if _, ok := distinctMap[i][ckey]; ok {
							distCols[i][j] = nil
						} else {
							distinctMap[i][ckey] = true
						}
					}
					mutex.Unlock()
				}

				for i := 0; i < rg0.GetRowsNumber(); i++ {
					r := rg0.GetRow(i)
					flag := false
					for _, c := range distCols {
						r.AppendVals(c[i])
						if c[i] != nil {
							flag = true
						}
					}
					if !flag {
						continue
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
