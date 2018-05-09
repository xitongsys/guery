package Master

import (
	"html/template"
	"net/http"

	"github.com/xitongsys/guery/Logger"
	"github.com/xitongsys/guery/Scheduler"
)

type UIData struct {
	Todos, Doings, Dones, Fails Scheduler.TaskList
}

func (self *Master) UIHandler(response http.ResponseWriter, request *http.Request) {
	Logger.Infof("UIHandler")

	tmpl := template.Must(template.ParseFiles("UI/index.html"))
	data := UIData{
		Todos:  self.Scheduler.Todos,
		Doings: self.Scheduler.Doings,
		Dones:  self.Scheduler.Dones,
		Fails:  self.Scheduler.Fails,
	}

	tmpl.Execute(response, data)
}
