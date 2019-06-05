package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

func init() {
}

// UpdateSeed ...
func UpdateSeed(ctx context.Context) (string, error) {
	fmt.Println("UpdateSeed")
	return "Seed Updated", nil
}

func main() {
	lambda.Start(UpdateSeed)
}
