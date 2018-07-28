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

	masterFlag    = app.Command("master", "Start a master")
	masterAddress = masterFlag.Flag("address", "host:port").Default(":1234").String()
	masterConfig  = masterFlag.Flag("config", "config file").Default("./config.json").String()

	agentFlag    = app.Command("agent", "Start a agent")
	agentMaster  = agentFlag.Flag("master", "host:port").Default("127.0.0.1:1234").String()
	agentAddress = agentFlag.Flag("address", "host:port").Default("127.0.0.1:0").String()
	agentName    = agentFlag.Flag("name", "agent name").Default("agent_" + uuid.Must(uuid.NewV4()).String()).String()
	agentConfig  = agentFlag.Flag("config", "config file").Default("./config.json").String()

	executorFlag    = app.Command("executor", "Start a executor")
	executorAgent   = executorFlag.Flag("agent", "host:port").Default("127.0.0.1:1234").String()
	executorAddress = executorFlag.Flag("address", "host:port").Default("127.0.0.1:0").String()
	executorName    = executorFlag.Flag("name", "executor name").Default("executor_" + uuid.Must(uuid.NewV4()).String()).String()
	executorConfig  = executorFlag.Flag("config", "config file").Default("./config.json").String()
)

func main() {
	logger.Infof("Welcome to use Guery !")
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case masterFlag.FullCommand():
		config.LoadConfig(*masterConfig)
		master.RunMaster(*masterAddress)

	case agentFlag.FullCommand():
		config.LoadConfig(*agentConfig)
		agent.RunAgent(*agentMaster, *agentAddress, *agentName)

	case executorFlag.FullCommand():
		config.LoadConfig(*executorConfig)
		executor.RunExecutor(*executorAgent, *executorAddress, *executorName)

	default:
		log.Fatalf("Guery failed to start: command error")
	}
}
