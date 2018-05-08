package main

import (
	"log"
	"os"

	"github.com/satori/go.uuid"
	"github.com/xitongsys/guery/Executor"
	"github.com/xitongsys/guery/Logger"
	"github.com/xitongsys/guery/Master"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app = kingpin.New("guery", "distributed SQL engine")

	master        = app.Command("master", "Start a master")
	masterAddress = master.Flag("address", "host:port").Default(":1234").String()

	executor        = app.Command("executor", "Start a executor")
	executorMaster  = executor.Flag("master", "host:port").Default("127.0.0.1:1234").String()
	executorAddress = executor.Flag("address", "host:port").Default("127.0.0.1:0").String()
	executorName    = executor.Flag("name", "executor name").Default("executor_" + uuid.Must(uuid.NewV4()).String()).String()
)

func main() {
	Logger.Infof("Welcome to use Guery !")
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case master.FullCommand():
		Master.RunMaster(*masterAddress)

	case executor.FullCommand():
		Executor.RunExecutor(*executorMaster, *executorAddress, *executorName)

	default:
		log.Fatalf("Guery failed to start: command error")
	}
}
