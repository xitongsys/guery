package Master

import (
	"io"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/soheilhy/cmux"
	"github.com/xitongsys/guery/Logger"
	"github.com/xitongsys/guery/Scheduler"
	"github.com/xitongsys/guery/Topology"
	"github.com/xitongsys/guery/pb"
	"google.golang.org/grpc"
)

var masterServer *Master

type Master struct {
	Topology  *Topology.Topology
	Scheduler *Scheduler.Scheduler
	StartTime time.Time
}

func NewMaster() *Master {
	m := &Master{
		Topology:  Topology.NewTopology(),
		StartTime: time.Now(),
	}
	m.Scheduler = Scheduler.NewScheduler(m.Topology)
	return m
}

func (self *Master) SendHeartbeat(stream pb.GueryMaster_SendHeartbeatServer) error {
	var hb *pb.AgentHeartbeat
	for {
		hbc, err := stream.Recv()
		if err == nil {
			if hb == nil {
				hb = hbc
				Logger.Infof("Add executor %v", hb.Location)
			}

		} else {
			if hb != nil {
				self.Topology.DropAgentInfo(hb)
			}
			if err == io.EOF {
				return nil
			}
			if err != nil {
				return err
			}
		}
		self.Topology.UpdateAgentInfo(hb)
		self.Scheduler.KillErrorTasks(hb)
	}
	return nil
}

///////////////////////////
func RunMaster(address string) {
	masterServer = NewMaster()

	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Master failed to start on %v: %v", address, err)
	}
	defer listener.Close()

	m := cmux.New(listener)

	grpcL := m.Match(cmux.HTTP2HeaderField("content-type", "application/grpc"))
	httpL := m.Match(cmux.Any())

	grpcS := grpc.NewServer()
	pb.RegisterGueryMasterServer(grpcS, masterServer)

	r := mux.NewRouter()
	r.HandleFunc("/", masterServer.UIHandler)
	r.HandleFunc("/UI/{dir}/{file}", masterServer.UIHandler)

	r.HandleFunc("/query", masterServer.QueryHandler)
	r.HandleFunc("/getinfo", masterServer.GetInfoHandler)
	r.HandleFunc("/control", masterServer.ControlHandler)
	httpS := &http.Server{Handler: r}

	go grpcS.Serve(grpcL)
	go httpS.Serve(httpL)

	masterServer.Scheduler.AutoFresh()

	if err := m.Serve(); err != nil {
		log.Fatalf("Master failed to serve: %v", err)

	}

}
