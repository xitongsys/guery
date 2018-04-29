package main

import (
	"fmt"
	"log"
	"os"

	"github.com/satori/go.uuid"
	"github.com/xitongsys/guery/Executor"
	"github.com/xitongsys/guery/Master"
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
	executorAddress = executor.Flag("address", "host:port").Default("127.0.0.1:0").String()
	executorName    = executor.Flag("name", "executor name").Default("executor_" + uuid.Must(uuid.NewV4()).String()).String()
)

func main() {
	fmt.Println("Welcome to use Guery !")
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case master.FullCommand():
		Master.RunMaster(*masterAddress)

	case executor.FullCommand():
		Executor.RunExecutor(*executorMaster, *executorDC, *executorRack, *executorAddress, *executorName)

	default:
		log.Fatalf("Guery failed to start: command error")
	}
}
