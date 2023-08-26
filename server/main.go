package main

import (
//	"grpc-jobs/server/queries"
	servergrpc "grpc-jobs/server/serverGrpc"
//	"log"
)

func main() {

	// jobs, err := queries.GetJobsQuery()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Println(jobs)

	servergrpc.Init()


}