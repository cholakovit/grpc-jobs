$ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go get github.com/githubnemo/CompileDaemon
go get go.mongodb.org/mongo-driver/mongo@latest

Starting commands:
client/     gin --appPort 3000 --immediate
server/     go run main.go

protoc --go_out=. --go-grpc_out=. proto/jobs.proto
go mod tidy

$ CompileDaemon -command="./server"

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