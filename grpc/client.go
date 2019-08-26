package main

import (
	"awesomeProject/grpc/pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
)

func main() {
	c,e := grpc.Dial("localhost:10000",grpc.WithInsecure())
	if e != nil {
		fmt.Println(e)
		return
	}
	demoClient := pb.NewDemoClient(c)
	resp,err := demoClient.SayHello(context.TODO(),&pb.HelloRequest{Greeting:"Hello"})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp.GetReply())

	client,err := demoClient.ReportName(context.TODO(),&pb.ReportNameRequest{})
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		resp := new(pb.ReportNameResponse)
		if err := client.RecvMsg(resp);err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(resp.String())
	}
}
