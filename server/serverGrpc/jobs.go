package servergrpc

import (
	"grpc-jobs/server/queries"
	"io"

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