package main

import (
	"fmt"
	"os"

	"github.com/satori/go.uuid"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app = kingpin.New("guery", "distributed SQL engine")

	master        = app.Command("master", "Start a master")
	masterAddress = master.Flag("address", "host:port").Default(":1234").String()

	executor        = app.Command("executor", "Start a executor")
	executorMaster  = executor.Flag("master", "host:port").Default(":1234").String()
	executorDC      = executor.Flag("datacenter", "data center name").Default("default").String()
	executorRack    = executor.Flag("rack", "rack name").Default("default").String()
	executorAddress = executor.Flag("address", "host:port").Default(":4321").String()
	executorName    = executor.Flag("name", "executor name").Default("executor_" + uuid.Must(uuid.NewV4())).String()
)

func main() {
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case master.FullCommand():

	case executor.FullCommand():
	default:

	}
}
