package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

func init() {
}

// GetSeed ...
func GetSeed(ctx context.Context) (string, error) {
	fmt.Println("GetSeed")
	return "Seed Retrieved", nil
}

func main() {
	lambda.Start(GetSeed)
}
