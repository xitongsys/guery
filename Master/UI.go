package Master

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

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
	path := request.URL.Path

	if strings.Contains(path[1:], ".html") {
		response.Header().Set("content-type", "text/html")
		fmt.Fprintf(response, getHtmlFile(path[1:]))
	} else if strings.Contains(path[1:], ".css") {
		response.Header().Set("content-type", "text/css")
		fmt.Fprintf(response, getHtmlFile(path[1:]))
	} else if strings.Contains(path[1:], ".js") {
		response.Header().Set("content-type", "text/javascript")
		fmt.Fprintf(response, getHtmlFile(path[1:]))
	} else {
		fmt.Fprintf(response, getHtmlFile("UI/index.html"))
	}

}

func getHtmlFile(path string) string {
	if data, err := ioutil.ReadFile(path); err != nil {
		return fmt.Sprintf("Error: %v", err)
	} else {
		return string(data)
	}
}
