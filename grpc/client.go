package main

import (
	"awesomeProject/grpc/pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/resolver"
	"log"
	"time"
)

const address = "dns:///dns-record-name:443"

func main() {
	// The secret sauce
	resolver.SetDefaultScheme("dns")
	// Set up a connection to the server.

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	c, err := grpc.DialContext(ctx, address, grpc.WithInsecure(), grpc.WithBalancerName(roundrobin.Name))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer c.Close()
	demoClient := pb.NewDemoClient(c)
	resp, err := demoClient.SayHello(context.TODO(), &pb.HelloRequest{Greeting: "Hello"})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp.GetReply())

	client, err := demoClient.ReportName(context.TODO(), &pb.ReportNameRequest{})
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		resp := new(pb.ReportNameResponse)
		if err := client.RecvMsg(resp); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(resp.String())
	}
}
