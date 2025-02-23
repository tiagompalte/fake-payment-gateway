package main

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/tiagompalte/fake-payment-gateway/application"
	"github.com/tiagompalte/fake-payment-gateway/internal/app/job"
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatalf("arg job name is empty")
	}

	jobName := os.Args[1]
	if strings.TrimSpace(jobName) == "" {
		log.Fatalf("arg job name is empty")
	}

	app, err := application.Build()
	if err != nil {
		log.Fatalf("failed to build the application (error: %v)", err)
	}

	jobGroup := job.NewJobGroup(app)
	err = jobGroup.Execute(context.Background(), jobName, os.Args[2:])
	if err != nil {
		log.Fatalf("error to execute job: %v", err)
	}
}
