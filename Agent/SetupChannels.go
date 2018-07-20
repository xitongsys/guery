package Agent

import (
	"context"
	"fmt"
	"io"
	"net"
	"strings"

	"github.com/xitongsys/guery/Logger"
	"github.com/xitongsys/guery/Util"
	"github.com/xitongsys/guery/pb"
	"google.golang.org/grpc"
)

func (self *Agent) SetupWriters(ctx context.Context, empty *pb.Empty) (*pb.Empty, error) {
	Logger.Infof("[Agent] SetupWriters start")
	var err error
	return empty, err
}

func (self *Agent) SetupReaders(ctx context.Context, empty *pb.Empty) (*pb.Empty, error) {
	Logger.Infof("[Agent] SetupReaders start")
	var err error
	return empty, err
}
