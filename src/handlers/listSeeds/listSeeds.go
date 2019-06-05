package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

func init() {
}

// ListSeeds ...
func ListSeeds(ctx context.Context) (string, error) {
	fmt.Println("ListSeeds")
	return "Seeds Listed", nil
}

func main() {
	lambda.Start(ListSeeds)
}
