package servergrpc

import (
	"context"
	"grpc-jobs/server/queries"
	"io"
	"log"

	pb "grpc-jobs/proto"
)

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

func (s *JobsServer) ReturnJobList(ctx context.Context, req *pb.NoParam) (*pb.JobListResponse, error) {

	jobs, err := queries.GetJobsQuery()
	if err != nil {
		log.Fatal(err)
	}

	pbJobsSlice, err := queries.ConvertModelSliceToProtoSlice(jobs)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("pbJobsSlice: ", pbJobsSlice)

	return &pb.JobListResponse{
		Message: pbJobsSlice,
	}, nil
}