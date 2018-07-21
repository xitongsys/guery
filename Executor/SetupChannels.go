package Executor

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

func (self *Executor) SetupWriters(ctx context.Context, empty *pb.Empty) (*pb.Empty, error) {
	Logger.Infof("SetupWriters start")
	var err error

	ip := strings.Split(self.Address, ":")[0]

	for i := 0; i < len(self.OutputLocations); i++ {
		pr, pw := io.Pipe()
		self.Writers = append(self.Writers, pw)
		listener, err := net.Listen("tcp", ip+":0")
		if err != nil {
			Logger.Errorf("failed to open listener: %v", err)
			return nil, fmt.Errorf("failed to open listener: %v", err)
		}

		self.OutputChannelLocations = append(self.OutputChannelLocations,
			&pb.Location{
				Name:    self.Name,
				Address: Util.GetHostFromAddress(listener.Addr().String()),
				Port:    Util.GetPortFromAddress(listener.Addr().String()),
			},
		)

		go func() {
			for {
				select {
				case <-self.DoneChan:
					listener.Close()
					return

				default:
					conn, err := listener.Accept()
					if err != nil {
						Logger.Errorf("failed to accept: %v", err)
						continue
					}
					Logger.Infof("connect %v", conn)

					go func(w io.Writer) {
						err := Util.CopyBuffer(pr, w)
						if err != nil && err != io.EOF {
							Logger.Errorf("failed to CopyBuffer: %v", err)
						}
						if wc, ok := w.(io.WriteCloser); ok {
							wc.Close()
						}
					}(conn)
				}
			}
		}()
	}
	Logger.Infof("SetupWriters Input=%v, Output=%v", self.InputChannelLocations, self.OutputChannelLocations)
	return empty, err
}

func (self *Executor) SetupReaders(ctx context.Context, empty *pb.Empty) (*pb.Empty, error) {
	var err error
	Logger.Infof("SetupReaders start")

	for i := 0; i < len(self.InputLocations); i++ {
		pr, pw := io.Pipe()
		self.Readers = append(self.Readers, pr)

		conn, err := grpc.Dial(self.InputLocations[i].GetURL(), grpc.WithInsecure())

		if err != nil {
			Logger.Errorf("failed to connect to %v: %v", self.InputLocations[i], err)
			return empty, err
		}
		client := pb.NewGueryAgentClient(conn)
		inputChannelLocation, err := client.GetOutputChannelLocation(context.Background(), self.InputLocations[i])

		if err != nil {
			Logger.Errorf("failed to connect %v: %v", self.InputLocations[i], err)
			return empty, err
		}

		conn.Close()

		self.InputChannelLocations = append(self.InputChannelLocations, inputChannelLocation)
		cconn, err := net.Dial("tcp", inputChannelLocation.GetURL())
		if err != nil {
			Logger.Errorf("failed to connect to input channel %v: %v", inputChannelLocation, err)
			return empty, err
		}
		Logger.Infof("connect to %v", inputChannelLocation)

		go func(r io.Reader) {
			err := Util.CopyBuffer(r, pw)
			if err != nil && err != io.EOF {
				Logger.Errorf("failed to CopyBuffer: %v", err)
			}
			pw.Close()
			if rc, ok := r.(io.ReadCloser); ok {
				rc.Close()
			}
		}(cconn)
	}
	Logger.Infof("SetupReaders Input=%v, Output=%v", self.InputChannelLocations, self.OutputChannelLocations)
	return empty, err
}
