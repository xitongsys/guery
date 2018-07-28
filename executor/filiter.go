package executor

import (
	"fmt"
	"io"
	"os"
	"runtime/pprof"
	"sync"
	"time"

	"github.com/vmihailenco/msgpack"
	"github.com/xitongsys/guery/config"
	"github.com/xitongsys/guery/eplan"
	"github.com/xitongsys/guery/logger"
	"github.com/xitongsys/guery/metadata"
	"github.com/xitongsys/guery/pb"
	"github.com/xitongsys/guery/row"
	"github.com/xitongsys/guery/util"
)

func (self *Executor) SetInstructionFilter(instruction *pb.Instruction) (err error) {
	var enode eplan.EPlanFilterNode
	if err = msgpack.Unmarshal(instruction.EncodedEPlanNodeBytes, &enode); err != nil {
		return err
	}
	self.Instruction = instruction
	self.EPlanNode = &enode
	self.InputLocations = []*pb.Location{&enode.Input}
	self.OutputLocations = []*pb.Location{&enode.Output}
	return nil
}

func (self *Executor) RunFilter() (err error) {
	fname := fmt.Sprintf("executor_%v_filter_%v_cpu.pprof", self.Name, time.Now().Format("20060102150405"))
	f, _ := os.Create(fname)
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	defer func() {
		if err != nil {
			self.AddLogInfo(err, pb.LogLevel_ERR)
		}
		self.Clear()
	}()

	if self.Instruction == nil {
		return fmt.Errorf("No Instruction")
	}
	enode := self.EPlanNode.(*eplan.EPlanFilterNode)

	md := &metadata.Metadata{}
	reader := self.Readers[0]
	writer := self.Writers[0]
	if err = util.ReadObject(reader, md); err != nil {
		return err
	}

	//write metadata
	if err = util.WriteObject(writer, md); err != nil {
		return err
	}

	rbReader := row.NewRowsBuffer(md, reader, nil)
	rbWriter := row.NewRowsBuffer(md, nil, writer)

	//write rows
	jobs := make(chan *row.Row)
	var wg sync.WaitGroup

	//init
	for _, be := range enode.BooleanExpressions {
		if err := be.Init(md); err != nil {
			return err
		}
	}

	for i := 0; i < int(config.Conf.Runtime.ParallelNumber); i++ {
		wg.Add(1)
		go func() {
			defer func() {
				wg.Done()
			}()

			for {
				r, ok := <-jobs
				//log.Println("========Filiter", row, ok)
				if ok {
					rg := row.NewRowsGroup(md)
					rg.Write(r)
					flag := true
					for _, booleanExpression := range enode.BooleanExpressions {
						if ok, err := booleanExpression.Result(rg); err == nil && !ok.([]interface{})[0].(bool) {
							flag = false
							break
						} else if err != nil {
							flag = false
							break
						}
					}

					if flag {
						err = rbWriter.WriteRow(r)
					}

					if err != nil {
						self.AddLogInfo(err, pb.LogLevel_ERR)
						break
					}

				} else {
					break
				}
			}
		}()
	}

	var r *row.Row
	for err == nil {
		r, err = rbReader.ReadRow()
		if err == io.EOF {
			err = nil
			break
		}
		if err != nil {
			break
		}
		jobs <- r
	}
	close(jobs)
	wg.Wait()

	if err = rbWriter.Flush(); err != nil {
		return err
	}

	logger.Infof("RunFilter finished")
	return err
}
