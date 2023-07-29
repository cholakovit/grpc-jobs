
go get go.mongodb.org/mongo-driver/mongo@latest

protoc --go_out=. --go-grpc_out=. proto/jobs.proto
go mod tidy

Creating Job JSON: 
{
  "id": "1",
  "title": "qqq",
  "description": "qqq",
  "company": "qqq",
  "location": "qq",
  "employment_type": "qqqq",
  "salary": "qqq",
  "requirements": "qqq",
  "responsibilities": "qqq",
  "posted_date": "2006-01-02T15:04:05Z",
  "expiry_date": "2006-01-02T15:04:05Z",
  "contact_information": "qqq",
  "application_process": "qqq"
}