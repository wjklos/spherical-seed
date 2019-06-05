package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

func init() {
}

// AddSeed ...
func AddSeed(ctx context.Context) (string, error) {
	fmt.Println("AddSeed")
	return "Seed Added", nil
}

func main() {
	lambda.Start(AddSeed)
}
