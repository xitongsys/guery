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

	todoLen, doingLen, doneLen, failLen := len(self.Scheduler.Todos), len(self.Scheduler.Doings), len(self.Scheduler.Dones), len(self.Scheduler.Fails)
	if todoLen > 100 {
		todoLen = 100
	}
	if doingLen > 100 {
		doingLen = 100
	}
	if doneLen > 100 {
		doneLen = 100
	}
	if failLen > 100 {
		failLen = 100
	}

	data := UIData{
		Todos:  self.Scheduler.Todos[:todoLen],
		Doings: self.Scheduler.Doings[:doingLen],
		Dones:  self.Scheduler.Dones[:doneLen],
		Fails:  self.Scheduler.Fails[:failLen],
	}

	tmpl.Execute(response, data)
}
