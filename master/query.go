package master

import (
	"fmt"
	"net/http"

	"github.com/xitongsys/guery/config"
	"github.com/xitongsys/guery/logger"
)

func (self *Master) QueryHandler(response http.ResponseWriter, request *http.Request) {
	Logger.Infof("QueryHandler")
	var err error

	if err = request.ParseForm(); err != nil {
		response.Write([]byte(fmt.Sprintf("Request Error: %v", err)))
		return
	}

	maxConcurrentNumberStr := request.FormValue("maxconcurrentnumber")
	var maxConcurrentNumber int32
	fmt.Sscanf(maxConcurrentNumberStr, "%d", &maxConcurrentNumber)
	if maxConcurrentNumber <= 0 || maxConcurrentNumber > int32(Config.Conf.Runtime.MaxConcurrentNumber) {
		maxConcurrentNumber = int32(Config.Conf.Runtime.MaxConcurrentNumber)
	}

	sqlStr := request.FormValue("sql")
	catalog := request.FormValue("catalog")
	if catalog == "" {
		catalog = Config.Conf.Runtime.Catalog
	}
	schema := request.FormValue("schema")
	if schema == "" {
		schema = Config.Conf.Runtime.Schema
	}
	s3Region := request.FormValue("s3region")
	if s3Region == "" {
		s3Region = Config.Conf.Runtime.S3Region
	}
	parallelNumber := Config.Conf.Runtime.ParallelNumber
	fmt.Sprintf(request.FormValue("ParallelNumber"), "%d", &parallelNumber)
	if parallelNumber <= 0 {
		parallelNumber = Config.Conf.Runtime.ParallelNumber
	}

	if s3Region == "" {
		s3Region = Config.Conf.Runtime.S3Region
	}

	priorityStr := request.FormValue("priority")
	var priority int32
	fmt.Sscanf(priorityStr, "%d", &priority)
	if priority < 0 {
		priority = Config.Conf.Runtime.Priority
	}

	runtime := &Config.ConfigRuntime{
		MaxConcurrentNumber: maxConcurrentNumber,
		Catalog:             catalog,
		Schema:              schema,
		Priority:            priority,
		S3Region:            s3Region,
		ParallelNumber:      parallelNumber,
	}

	task, err := self.Scheduler.AddTask(runtime, sqlStr, response)
	if err != nil {
		response.Write([]byte(fmt.Sprintf("%s", err)))
		return
	}

	<-task.DoneChan

}
