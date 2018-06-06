package Master

import (
	"fmt"
	"net/http"

	"github.com/xitongsys/guery/Logger"
)

func (self *Master) ControlHandler(response http.ResponseWriter, request *http.Request) {
	Logger.Infof("ControlHandler")
	var err error

	if err = request.ParseForm(); err != nil {
		response.Write([]byte(fmt.Sprintf("Request Error: %v", err)))
		return
	}
	cmd := request.FormValue("cmd")
	switch cmd {
	case "killexecutor":
		name := request.FormValue("name")
		self.Topology.KillExecutor(name)
	case "canceltask":
		para := request.FormValue("taskid")
		var taskid int64
		fmt.Sscanf(para, "%d", &taskid)
		self.Scheduler.CancelTask(taskid)
	}
}
