package Executor

import (
	"context"
	"io"
	"net"
	"strings"

	"github.com/xitongsys/guery/Logger"
	"github.com/xitongsys/guery/Util"
	"github.com/xitongsys/guery/pb"
	"google.golang.org/grpc"
)

func (self *Executor) SetupChannels() error {
	ip := strings.Split(self.Address, ":")[0]

	for i := 0; i < len(self.OutputLocations); i++ {
		pr, pw := io.Pipe()
		self.Writers = append(self.Writers, pw)
		go func() {
			listener, err := net.Listen("tcp", ip+":0")
			if err != nil {
				Logger.Errorf("failed to open listener: %v", err)
				return
			}

			self.OutputChannelLocations = append(self.OutputChannelLocations,
				&pb.Location{
					DataCenter: self.DataCenter,
					Rack:       self.Rack,
					Address:    listener.Addr().String(),
					Name:       self.Name,
				},
			)

			for {
				conn, err := listener.Accept()
				if err != nil {
					Logger.Errorf("failed to accept: %v", err)
					continue
				}
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
		}()
	}

	for i := 0; i < len(self.InputLocations); i++ {
		pr, pw := io.Pipe()
		self.Readers = append(self.Readers, pr)

		conn, err := grpc.Dial(self.InputLocations[i].Address, grpc.WithInsecure())
		if err != nil {
			Logger.Errorf("failed to connect to %v: %v", self.InputLocations[i], err)
			return err
		}
		client := pb.NewGueryExecutorClient(conn)
		inputChannelLocation, err := client.GetOutputChannelLocation(context.Background(), self.InputChannelLocations[i])
		if err != nil {
			Logger.Errorf("failed to connect %v: %v", self.InputLocations[i], err)
			return err
		}
		conn.Close()

		self.InputChannelLocations = append(self.InputChannelLocations, inputChannelLocation)
		cconn, err := net.Dial("tcp", inputChannelLocation.Address)
		if err != nil {
			Logger.Errorf("failed to connect to input channel %v: %v", inputChannelLocation, err)
			return err
		}

		go func(r io.Reader) {
			err := Util.CopyBuffer(r, pw)
			if err != nil && err != io.EOF {
				Logger.Errorf("failed to CopyBuffer: %v", err)
			}
			pw.Close()
		}(cconn)
	}
	return nil
}
