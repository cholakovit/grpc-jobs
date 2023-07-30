package servergrpc

import (
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"

	pb "grpc-jobs/proto"
)

//this is the struct to be created, pb is imported upstairs
type JobsServer struct {
	pb.JobServiceServer
}

func Init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")

	//listen on the port
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("Failed to start server %v", err)
	}

	// create a new gRPC server
	grpcServer := grpc.NewServer()
	// register the job service
	pb.RegisterJobServiceServer(grpcServer, &JobsServer{})
	log.Printf("Server started at %v", lis.Addr())
	//list is the port, the grpc server needs to start there
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start: %v", err)
	}
}