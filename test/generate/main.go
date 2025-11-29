package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	testpb "github.com/khbdev/proto-online-test/proto/generate"
	"google.golang.org/grpc"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [get|post]")
		return
	}

	action := os.Args[1]

	// gRPC connection
	conn, err := grpc.Dial("127.0.0.1:50055", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := testpb.NewTestServiceClient(conn)

	switch action {
	case "get":
		key := "Example_1761491567510010691" // Redis key, test yaratgandan key oling
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		resp, err := client.GetTest(ctx, &testpb.GetTestRequest{Key: key})
		if err != nil {
			log.Fatalf("Failed to get test: %v", err)
		}

		fmt.Println("Test JSON:", resp.TestJson)

	case "post":
		// 2 section yuboriladi (ID=2 faqat mavjud)
		sectionIDs := []int32{2}
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		resp, err := client.GenerateTest(ctx, &testpb.GenerateTestRequest{
			Name:       "10-ta-oldimi-yoqmi-2",
			SectionIds: sectionIDs,
		})
		if err != nil {
			log.Fatalf("Failed to generate test: %v", err)
		}

		fmt.Println("Test link:", resp.Link)

	default:
		fmt.Println("Unknown action. Use get or post")
	}
}
