package main

//go:generate protoc -I=pb/ pb/server.proto --go_out=plugins=grpc:pb/

import (
	"awesomeProject/grpc/pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"time"
)


type Demo struct {}

func (d *Demo) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse,  error) {
	return &pb.HelloResponse{Reply:"hello"},nil
}

func (d *Demo) ReportName(req *pb.ReportNameRequest,resp pb.Demo_ReportNameServer) error {
	for {
		if err := resp.Send(&pb.ReportNameResponse{Name:"alvin"});err !=nil {
			return err
		}
		time.Sleep(time.Second)
	}
}

func main() {
	l,e := net.Listen("tcp","0.0.0.0:10000")
	if e != nil {
		fmt.Println(e)
		return
	}

	rpcServer := grpc.NewServer()

	pb.RegisterDemoServer(rpcServer,new(Demo))

	if e := rpcServer.Serve(l);e != nil {
		fmt.Println(e)
	}
	return
}