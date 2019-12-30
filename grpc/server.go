package main

//go:generate protoc -I=pb/ pb/server.proto --go_out=plugins=grpc:pb/

import (
	"awesomeProject/grpc/pb"
	"context"
	"fmt"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"net"
	"time"
)

type ConsulService struct {
	IP   string
	Port int
	Tag  []string
	Name string
}

//health
type HealthImpl struct{}

// Check implements the health check interface, which directly returns to health status. There are also more complex health check strategies, such as returning based on server load.
func (h *HealthImpl) Check(ctx context.Context, req *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	fmt.Print("health checking\n")
	return &grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	}, nil
}

func (h *HealthImpl) Watch(req *grpc_health_v1.HealthCheckRequest, w grpc_health_v1.Health_WatchServer) error {
	return nil
}

type Demo struct{}

func (d *Demo) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Reply: "hello"}, nil
}

func (d *Demo) ReportName(req *pb.ReportNameRequest, resp pb.Demo_ReportNameServer) error {
	for {
		if err := resp.Send(&pb.ReportNameResponse{Name: "alvin"}); err != nil {
			return err
		}
		time.Sleep(time.Second)
	}
}

func RegitserService(ca string, cs *ConsulService) {

	//register consul
	consulConfig := api.DefaultConfig()
	consulConfig.Address = ca
	client, err := api.NewClient(consulConfig)
	if err != nil {
		fmt.Printf("NewClient error\n%v", err)
		return
	}
	agent := client.Agent()
	interval := time.Duration(10) * time.Second
	deregister := time.Duration(1) * time.Minute

	reg := &api.AgentServiceRegistration{
		ID:      fmt.Sprintf("%v-%v-%v", cs.Name, cs.IP, cs.Port), // name of service node
		Name:    cs.Name,                                          // Service Name
		Tags:    cs.Tag,                                           // tag, can be empty
		Port:    cs.Port,                                          // Service Port
		Address: cs.IP,                                            //Service IP
		Check: &api.AgentServiceCheck{ // Health Examination
			Interval: interval.String(),                                // health check interval
			GRPC:     fmt.Sprintf("%v:%v/%v", cs.IP, cs.Port, cs.Name), // grpc support, address to perform health check, service will be passed to Health. Check function
			DeregisterCriticalServiceAfter: deregister.String(), // logout time, equivalent to expiration time
		},
	}

	fmt.Printf("registing to %v\n", ca)
	if err := agent.ServiceRegister(reg); err != nil {
		fmt.Printf("Service Register error\n%v", err)
		return
	}

}

func RegisterToConsul() {
	RegitserService("127.0.0.1:8500", &ConsulService{
		Name: "helloworld",
		Tag:  []string{"helloworld"},
		IP:   "127.0.0.1",
		Port: 50051,
	})
}

func main() {
	l, e := net.Listen("tcp", "0.0.0.0:10000")
	if e != nil {
		fmt.Println(e)
		return
	}

	rpcServer := grpc.NewServer()

	pb.RegisterDemoServer(rpcServer, new(Demo))

	grpc_health_v1.RegisterHealthServer(rpcServer, &HealthImpl{})

	RegisterToConsul()

	if e := rpcServer.Serve(l); e != nil {
		fmt.Println(e)
	}
	return
}
