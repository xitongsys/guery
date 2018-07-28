package main

import (
	"log"
	"os"

	"github.com/satori/go.uuid"
	"github.com/xitongsys/guery/agent"
	"github.com/xitongsys/guery/config"
	"github.com/xitongsys/guery/executor"
	"github.com/xitongsys/guery/logger"
	"github.com/xitongsys/guery/master"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app = kingpin.New("guery", "distributed SQL engine")

	master        = app.Command("master", "Start a master")
	masterAddress = master.Flag("address", "host:port").Default(":1234").String()
	masterConfig  = master.Flag("config", "config file").Default("./config.json").String()

	agent        = app.Command("agent", "Start a agent")
	agentMaster  = agent.Flag("master", "host:port").Default("127.0.0.1:1234").String()
	agentAddress = agent.Flag("address", "host:port").Default("127.0.0.1:0").String()
	agentName    = agent.Flag("name", "agent name").Default("agent_" + uuid.Must(uuid.NewV4()).String()).String()
	agentConfig  = agent.Flag("config", "config file").Default("./config.json").String()

	executor        = app.Command("executor", "Start a executor")
	executorAgent   = executor.Flag("agent", "host:port").Default("127.0.0.1:1234").String()
	executorAddress = executor.Flag("address", "host:port").Default("127.0.0.1:0").String()
	executorName    = executor.Flag("name", "executor name").Default("executor_" + uuid.Must(uuid.NewV4()).String()).String()
	executorConfig  = executor.Flag("config", "config file").Default("./config.json").String()
)

func main() {
	Logger.Infof("Welcome to use Guery !")
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case master.FullCommand():
		Config.LoadConfig(*masterConfig)
		Master.RunMaster(*masterAddress)

	case agent.FullCommand():
		Config.LoadConfig(*agentConfig)
		Agent.RunAgent(*agentMaster, *agentAddress, *agentName)

	case executor.FullCommand():
		Config.LoadConfig(*executorConfig)
		Executor.RunExecutor(*executorAgent, *executorAddress, *executorName)

	default:
		log.Fatalf("Guery failed to start: command error")
	}
}
