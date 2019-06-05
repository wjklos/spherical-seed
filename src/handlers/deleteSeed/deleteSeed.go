package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

func init() {
}

// DeleteSeed ...
func DeleteSeed(ctx context.Context) (string, error) {
	fmt.Println("DeleteSeed")
	return "Seed Deleted", nil
}

func main() {
	lambda.Start(DeleteSeed)
}
