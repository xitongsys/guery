package Master

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/xitongsys/guery/Logger"
)

func (self *Master) GetInfoHandler(response http.ResponseWriter, resquest *http.Request) {
	info := self.GetInfo()
	res, _ := json.Marshal(info)
	response.Write(res)
}

func (self *Master) UIHandler(response http.ResponseWriter, request *http.Request) {
	Logger.Infof("UIHandler")
	path := request.URL.Path

	if strings.Contains(path[1:], ".html") {
		response.Header().Set("content-type", "text/html")
		fmt.Fprint(response, getHtmlFile(path[1:]))
	} else if strings.Contains(path[1:], ".css") {
		response.Header().Set("content-type", "text/css")
		fmt.Fprint(response, getHtmlFile(path[1:]))
	} else if strings.Contains(path[1:], ".js") {
		response.Header().Set("content-type", "text/javascript")
		fmt.Fprint(response, getHtmlFile(path[1:]))
	} else {
		fmt.Fprint(response, getHtmlFile("UI/index.html"))
	}

}

func getHtmlFile(path string) string {
	if data, err := ioutil.ReadFile(path); err != nil {
		return fmt.Sprintf("Error: %v", err)
	} else {
		fmt.Println(string(data))
		return string(data)
	}
}
