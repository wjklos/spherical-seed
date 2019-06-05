package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

func init() {
}

// LoadSeeds ...
func LoadSeeds(ctx context.Context) (string, error) {
	fmt.Println("LiadSeeds")
	return "Seeds Loaded", nil
}

func main() {
	lambda.Start(LoadSeeds)
}
