package clientgrpc

import (
	"context"
	"log"

	pb "grpc-jobs/proto"
)

func CallCreateJob() {

}

func CallJobBiStream(client pb.JobServiceClient, job *pb.Jobs) {
	log.Printf("Bidirectional Streaming started")
	stream, err := client.JobsBiStreaming(context.Background())
	if err != nil {
		log.Fatalf("Could not send products: %v", err)
	}

	waitc := make(chan struct{})

	go func() {
		message, err := stream.Recv()
		if err != nil {
			log.Fatalf("Error while streaming %v", err)
		}
		log.Println("Comming from the server", message)
		close(waitc)
	}()

	message := &pb.Jobs{
		Id:                 job.Id,
		Title:              job.Title,
		Description:        job.Description,
		Company:            job.Company,
		Location:           job.Location,
		Salary:             job.Salary,
		Requirements:       job.Requirements,
		Responsibilities:   job.Responsibilities,
		ContactInformation: job.ContactInformation,
		ApplicationProcess: job.ApplicationProcess,
	}

	req := &pb.JobsRequest{
		Message: message,
	}

	if err := stream.Send(req); err != nil {
		log.Fatalf("Error while sending %v", err)
	}

	stream.CloseSend()
	<-waitc
	log.Printf("Bidirectional Streaming finished")
}