package regiservice

import (
	"github.com/CharlesWong/darkpassenger/config"
	"github.com/CharlesWong/darkpassenger/model"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"math/rand"
	"net"
	"time"
)

// server is used to implement helloworld.GreeterServer.
type regiServer struct{}

// SayHello implements helloworld.GreeterServer
func (s *regiServer) ListAllWorkerServices(ctx context.Context, in *model.NullMsg) (*model.ServiceInfos, error) {
	return &model.ServiceInfos{
		ServiceInfo: []*model.ServiceInfo{
			&model.ServiceInfo{
				WorkerAddr: config.FrontEndAddr,
				Version:    config.Version,
			}}}, nil
}

func StartRegiService() {
	lis, err := net.Listen("tcp", config.RegiServiceAddr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	model.RegisterRegistrationServiceServer(s, &regiServer{})
	s.Serve(lis)
}

func SelectWorkerAddr() *model.ServiceInfo {
	conn, err := grpc.Dial(config.RegiServiceAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	c := model.NewRegistrationServiceClient(conn)

	r, err := c.ListAllWorkerServices(context.Background(), &model.NullMsg{})
	if err != nil {
		log.Fatalf("could not reach regi service: %v", err)
	}
	if len(r.ServiceInfo) == 0 {
		log.Fatalln("No worker found.")
	}

	rand.Seed(time.Now().UTC().UnixNano())
	return r.ServiceInfo[rand.Int()%len(r.ServiceInfo)]

}
