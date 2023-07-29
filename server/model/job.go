package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Job struct {
	ID           					primitive.ObjectID `bson:"_id,omitempty"`
	Title        					string             `json:"title" 								bson:"title"`
	Description						string             `json:"description" 					bson:"description"`
	Company								string             `json:"company" 							bson:"company"`
	Location							string             `json:"location" 						bson:"location"`
	Employment_Type				string             `json:"employment_type" 			bson:"employment_type"`
	Salary								string             `json:"salary" 							bson:"salary"`
	Requirements					string             `json:"requirements" 				bson:"requirements"`
	Responsibilities			string             `json:"responsibilities" 		bson:"responsibilities"`
	Posted_Date						time.Time          `json:"posted_date" 					bson:"posted_date"`
	Expiry_Date						time.Time          `json:"expiry_date" 					bson:"expiry_date"`
	Contact_Information		string             `json:"contact_information" 	bson:"contact_information"`
	Application_Process		string             `json:"application_process" 	bson:"application_process"`
}