package main
import (
	"log"
	"net"
	"google.golang.org/grpc"
	pb "./Users"
)

func (s *server) ChatService(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main(){
	lis,err :=net.listen("tcp",":9090")
	if err!=nil{
		log.Fatalf("Failed to listen on port 9090")
	}
	grpcServer := grpc.newServe(lis); err!=nil{
		log.Fatalf("Failed to serve gRPC server over port 9090 %v", err)
	}
}