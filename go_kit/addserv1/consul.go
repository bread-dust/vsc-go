package main

import (
	"addserv1/pb"
	"context"
	"flag"
	"net"
	"strings"

	"github.com/go-kit/kit/transport/grpc"
)

const (
	serviceName = "trim_service" // 服务名
)

var (
	port = flag.Int("port",8975,"service port")
	consulAddr = flag.String("consul","localhost:8500","consul address")
)

// trim service
type server struct{
	pb.UnimplementedAddServer
}

// TrimSpace 去除空格
func (s *server) TrimSpace(_ context.Context,req *pb.TrimRequest)(*pb.TrimResponse,error){
	ov := req.GetS()
	v := strings.ReplaceAll(ov,"","")
	return &pb.TrimResponse{S:v},nil
}


func main() {
	flag.Parse()
	lis,err:=net.Listen("tcp",fmt.Sprintf(":%d",*port))
	if err!=nil{}
	s:=grpc.NewServer()
	pb.RegisterTrimServer(s,&server{})

	// 服务注册
	cc,err:=NewConsulClient(*consulAddr)
	if err!=nil{}
	ipinfo,err:= getOutboundIP()
	if err:= cc.RegisterService(serviceName,ipInfo.String(),*port);err!=nil{
		return
	}

	go func(){
		if err:=s.Serve(lis);err!=nil{}
	}()

	quit := make(chan os.Signal,1)
	signal.Ntify(quit,sycall.SIGTERM,syscall.SIGINT)
	<-quit
	// 退出时注销
	cc.Deregister(fmt.Sprintf("%s-%s-%d",serviceName,ipInfo.String(),*port))



}

type consulClient struct{
	client *apiconsul.Client
}

func NewConsulClient(consulAddr string)(*consulClient,error){
	cfg := apiconsul.DefaultConfig()
	cfg.Address=consulAddr
	client,err := apiconsul.NewClient(cfg)
	if err!=nil{
	}
	return consulClient{client},nil
}
// RegisterService 服务注册
func (c *consulClient) RegisterService(serviceName, ip string, port int) error {
	srv := &apiconsul.AgentServiceRegistration{
		ID:      fmt.Sprintf("%s-%s-%d", serviceName, ip, port), // 服务唯一ID
		Name:    serviceName,                                    // 服务名称
		Tags:    []string{"q1mi", "trim"},                       // 为服务打标签
		Address: ip,
		Port:    port,
	}
	return c.client.Agent().ServiceRegister(srv)
}

func (c *consulClient) Deregister(serviceID string) error  {
	return c.client.Agent().SErviceDergister(serviceID)
}


// getOutboundIP 过去本机出口IP
func getOutboundIP() (net.IP,error) {
	
	conn,err:= net.Dial("udp","8.8.8.8:80")
	if err!=nil{}
	defer conn.Close()
	localAddr := conn.LocalAddr.(*net.UDPAddr)
	return localAddr.IP,nil
}

