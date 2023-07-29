package clientgrpc

import (
	"grpc-jobs/server/model"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "grpc-jobs/proto"
)

const (
	port = ":8080"
)

func ConnectServer(postedJob *model.Job) {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewJobServiceClient(conn)

	job := &pb.Jobs{
		Title:              postedJob.Title,
		Description:        postedJob.Description,
		Company:            postedJob.Company,
		Location:           postedJob.Location,
		EmploymentType:     postedJob.Employment_Type,
		Salary:             postedJob.Salary,
		Requirements:			  postedJob.Requirements,
		Responsibilities:   postedJob.Responsibilities,
		ContactInformation: postedJob.Responsibilities,
		ApplicationProcess: postedJob.Application_Process,
	}
	CallJobBiStream(client, job)
}