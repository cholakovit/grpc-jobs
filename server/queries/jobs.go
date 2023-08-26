package queries

import (
	"context"
	"errors"
	"grpc-jobs/server/db"
	"grpc-jobs/server/model"
	"sync"
	"time"

	pb "grpc-jobs/proto"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	collection *mongo.Collection = db.OpenCollection(db.Client, "jobs")
	wg 				 sync.WaitGroup
)

func GetJobsQuery() ([]model.Job, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// Create a channel to receive the fetched jobs from goroutines
	jobChan := make(chan []model.Job)

	go func() {
		var fetchedJobs []model.Job
		for cursor.Next(ctx){
			var job model.Job
			err := cursor.Decode(&job)
			if err != nil {
				break
			}
			fetchedJobs = append(fetchedJobs, job)
		}
		jobChan <- fetchedJobs
	}()

	// Wait for the jobs to be fetched from the goroutine
	fetchedJobs := <- jobChan

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	if len(fetchedJobs) == 0 {
		return nil, errors.New("jobs not found")
	}

	return fetchedJobs, nil
}

func PostJobQuery(job *pb.Job) error {
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

func ConvertToModelJob(pbJob *pb.Job) model.Job {
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

func ConvertModelToProto(jobModel *model.Job) (*pb.Job, error) {
	// Create a new pb.Jobs instance
	jobPB := &pb.Job{}

	// Convert ID from primitive.ObjectID to string
	jobPB.Id = jobModel.ID.Hex()

	// Copy other fields directly
	jobPB.Title = jobModel.Title
	jobPB.Description = jobModel.Description
	jobPB.Company = jobModel.Company
	jobPB.Location = jobModel.Location
	jobPB.EmploymentType = jobModel.Employment_Type
	jobPB.Salary = jobModel.Salary
	jobPB.Requirements = jobModel.Requirements
	jobPB.Responsibilities = jobModel.Responsibilities
	jobPB.ContactInformation = jobModel.Contact_Information
	jobPB.ApplicationProcess = jobModel.Application_Process

	return jobPB, nil
}

func ConvertModelSliceToProtoSlice(jobs []model.Job) ([]*pb.Job, error) {
	var pbJobsSlice []*pb.Job

	for _, job := range jobs {
		pbJob, err := ConvertModelToProto(&job)
		if err != nil {
			return nil, err
		}
		pbJobsSlice = append(pbJobsSlice, pbJob)
	}

	return pbJobsSlice, nil
}