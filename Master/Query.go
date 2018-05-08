package Master

import (
	"fmt"
	"net/http"

	"github.com/xitongsys/guery/Logger"
)

func (self *Master) QueryHandler(response http.ResponseWriter, request *http.Request) {
	Logger.Infof("QueryHandler")
	var err error

	if err = request.ParseForm(); err != nil {
		response.Write([]byte(fmt.Sprintf("Request Error: %v", err)))
		return
	}
	sqlStr := request.FormValue("sql")
	catalog := request.FormValue("catalog")
	schema := request.FormValue("schema")

	task, err := self.Scheduler.AddTask(sqlStr, catalog, schema, 0, response)
	if err != nil {
		response.Write([]byte(fmt.Sprintf("%s", err)))
		return
	}

	<-task.DoneChan

}
