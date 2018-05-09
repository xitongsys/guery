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

	doneLen, failLen := len(self.Scheduler.Dones), len(self.Scheduler.Fails)

	if doneLen > 100 {
		doneLen = 100
	}
	if failLen > 100 {
		failLen = 100
	}

	data := UIData{
		Todos:  self.Scheduler.Todos,
		Doings: self.Scheduler.Doings,
		Dones:  self.Scheduler.Dones[len(self.Scheduler.Dones)-doneLen : len(self.Scheduler.Dones)],
		Fails:  self.Scheduler.Fails[len(self.Scheduler.Fails)-failLen : len(self.Scheduler.Fails)],
	}

	tmpl.Execute(response, data)
}
