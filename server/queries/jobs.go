package queries

import (
	"context"
	"grpc-jobs/server/db"
	"grpc-jobs/server/model"
	"sync"
	"time"

	pb "grpc-jobs/proto"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	collection *mongo.Collection = db.OpenCollection(db.Client, "jobs")
	wg 				 sync.WaitGroup
)

func PostJobQuery(job *pb.Jobs) error {
	var wg sync.WaitGroup // To wait for goroutine to complete
	var resultErr error   // To store the error from the database operation

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	modelJob := ConvertToModelJob(job)

	wg.Add(1)
	go func() {
			defer wg.Done()
			_, err := collection.InsertOne(ctx, modelJob)
			if err != nil {
					resultErr = err
			}
	}()

	wg.Wait() // Wait for the goroutine to finish
	return resultErr
}

func ConvertToModelJob(pbJob *pb.Jobs) model.Job {
	return model.Job{
		// Convert individual fields as needed
		Title:             pbJob.Title,
		Description:       pbJob.Description,
		Company:           pbJob.Company,
		Location:          pbJob.Location,
		Employment_Type:   pbJob.EmploymentType,
		Salary:            pbJob.Salary,
		Requirements:      pbJob.Requirements,
		Responsibilities:  pbJob.Responsibilities,
		Contact_Information: pbJob.ContactInformation,
		Application_Process: pbJob.ApplicationProcess,
	}
}