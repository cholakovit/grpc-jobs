package main

import (
	"io"
	"log"
	"net"

	pb "grpc-jobs/proto"
	"grpc-jobs/server/queries"

	"google.golang.org/grpc"
)

//define the port
const (
	port = ":8080"
)

//this is the struct to be created, pb is imported upstairs
type JobsServer struct {
	pb.JobServiceServer
}

func main() {
	//listen on the port
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to start server %v", err)
	}

	// create a new gRPC server
	grpcServer := grpc.NewServer()
	// register the greet service
	pb.RegisterJobServiceServer(grpcServer, &JobsServer{})
	log.Printf("Server started at %v", lis.Addr())
	//list is the port, the grpc server needs to start there
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start: %v", err)
	}

	
}

func (s *JobsServer) JobsBiStreaming(stream pb.JobService_JobsBiStreamingServer) error {

		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		
		queries.PostJobQuery(req.Message)
		res := &pb.JobsResponse{
			Message: "Success!!!",
		}
		if err := stream.Send(res); err != nil {
			return err
		}

		return nil
}