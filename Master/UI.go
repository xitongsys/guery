package Master

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/xitongsys/guery/Logger"
)

type UIInfo struct {
	Running  int
	Queued   int
	Finished int

	Active int
	Busy   int
	Free   int
}

func (self *Master) GetInfoHandler(response http.ResponseWriter, resquest *http.Request) {
	info := &UIInfo{
		Running:  len(self.Scheduler.Doings),
		Queued:   len(self.Scheduler.Todos),
		Finished: len(self.Scheduler.Dones) + len(self.Scheduler.Fails),

		Active: int(self.Scheduler.Topology.TotalExecutorNum),
		Busy:   int(self.Scheduler.Topology.TotalExecutorNum - self.Scheduler.Topology.IdleExecutorNum),
		Free:   int(self.Scheduler.Topology.IdleExecutorNum),
	}
	res, _ := json.Marshal(info)
	response.Write(res)
}

func (self *Master) UIHandler(response http.ResponseWriter, request *http.Request) {
	Logger.Infof("UIHandler")
	tmpl := template.Must(template.ParseFiles("UI/index.html"))
	tmpl.Execute(response, nil)
}
