package master

import (
	"fmt"
	"net/http"

	"github.com/xitongsys/guery/config"
	"github.com/xitongsys/guery/logger"
)

func (self *Master) QueryHandler(response http.ResponseWriter, request *http.Request) {
	logger.Infof("QueryHandler")
	var err error

	if err = request.ParseForm(); err != nil {
		response.Write([]byte(fmt.Sprintf("Request Error: %v", err)))
		return
	}

	maxConcurrentNumberStr := request.FormValue("maxconcurrentnumber")
	var maxConcurrentNumber int32
	fmt.Sscanf(maxConcurrentNumberStr, "%d", &maxConcurrentNumber)
	if maxConcurrentNumber <= 0 || maxConcurrentNumber > int32(config.Conf.Runtime.MaxConcurrentNumber) {
		maxConcurrentNumber = int32(config.Conf.Runtime.MaxConcurrentNumber)
	}

	sqlStr := request.FormValue("sql")
	catalog := request.FormValue("catalog")
	if catalog == "" {
		catalog = config.Conf.Runtime.Catalog
	}
	schema := request.FormValue("schema")
	if schema == "" {
		schema = config.Conf.Runtime.Schema
	}
	s3Region := request.FormValue("s3region")
	if s3Region == "" {
		s3Region = config.Conf.Runtime.S3Region
	}
	parallelNumber := config.Conf.Runtime.ParallelNumber
	fmt.Sprintf(request.FormValue("ParallelNumber"), "%d", &parallelNumber)
	if parallelNumber <= 0 {
		parallelNumber = config.Conf.Runtime.ParallelNumber
	}

	if s3Region == "" {
		s3Region = config.Conf.Runtime.S3Region
	}

	priorityStr := request.FormValue("priority")
	var priority int32
	fmt.Sscanf(priorityStr, "%d", &priority)
	if priority < 0 {
		priority = config.Conf.Runtime.Priority
	}

	runtime := &config.ConfigRuntime{
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
