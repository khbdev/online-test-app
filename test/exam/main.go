package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	exampb "github.com/khbdev/proto-online-test/proto/exam"
	"google.golang.org/grpc"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [create|get|getall|update|delete]")
		return
	}
	action := os.Args[1]

	// gRPC server
	conn, err := grpc.Dial("localhost:50058", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()
	client := exampb.NewExamServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	reader := bufio.NewReader(os.Stdin)

	switch strings.ToLower(action) {
	case "create":
		fmt.Print("Title: ")
		title, _ := reader.ReadString('\n')
		title = strings.TrimSpace(title)

		fmt.Print("Description: ")
		desc, _ := reader.ReadString('\n')
		desc = strings.TrimSpace(desc)

		fmt.Print("Section IDs (comma separated): ")
		secStr, _ := reader.ReadString('\n')
		secStr = strings.TrimSpace(secStr)
		secIDs := parseIDs(secStr)

		req := &exampb.CreateExamRequest{
			Title:       title,
			Description: desc,
			SectionIds:  secIDs,
		}
		resp, err := client.CreateExam(ctx, req)
		if err != nil {
			log.Fatalf("CreateExam failed: %v", err)
		}
		fmt.Println("Created Exam:")
		printExam(resp.Exam)

	case "get":
		fmt.Print("Exam ID: ")
		idStr, _ := reader.ReadString('\n')
		idStr = strings.TrimSpace(idStr)
		id, _ := strconv.ParseUint(idStr, 10, 64)

		resp, err := client.GetExamByID(ctx, &exampb.GetExamByIDRequest{Id: id})
		if err != nil {
			log.Fatalf("GetExamByID failed: %v", err)
		}
		printExam(resp.Exam)

	case "getall":
		resp, err := client.GetExamList(ctx, &exampb.GetExamListRequest{})
		if err != nil {
			log.Fatalf("GetExamList failed: %v", err)
		}
		fmt.Println("All Exams:")
		for _, e := range resp.Exams {
			printExam(e)
		}

	case "update":
		fmt.Print("Exam ID: ")
		idStr, _ := reader.ReadString('\n')
		idStr = strings.TrimSpace(idStr)
		id, _ := strconv.ParseUint(idStr, 10, 64)

		fmt.Print("New Title: ")
		title, _ := reader.ReadString('\n')
		title = strings.TrimSpace(title)

		fmt.Print("New Description: ")
		desc, _ := reader.ReadString('\n')
		desc = strings.TrimSpace(desc)

		fmt.Print("New Section IDs (comma separated): ")
		secStr, _ := reader.ReadString('\n')
		secStr = strings.TrimSpace(secStr)
		secIDs := parseIDs(secStr)

		req := &exampb.UpdateExamRequest{
			Id:          id,
			Title:       title,
			Description: desc,
			SectionIds:  secIDs,
		}
		resp, err := client.UpdateExam(ctx, req)
		if err != nil {
			log.Fatalf("UpdateExam failed: %v", err)
		}
		fmt.Println("Updated Exam:")
		printExam(resp.Exam)

	case "delete":
		fmt.Print("Exam ID: ")
		idStr, _ := reader.ReadString('\n')
		idStr = strings.TrimSpace(idStr)
		id, _ := strconv.ParseUint(idStr, 10, 64)

		resp, err := client.DeleteExam(ctx, &exampb.DeleteExamRequest{Id: id})
		if err != nil {
			log.Fatalf("DeleteExam failed: %v", err)
		}
		fmt.Printf("Deleted Exam success: %v\n", resp.Success)

	default:
		fmt.Println("Unknown action. Available: create, get, getall, update, delete")
	}
}

func parseIDs(input string) []uint64 {
	// space yoki comma bilan ajratish
	parts := strings.FieldsFunc(input, func(r rune) bool {
		return r == ',' || r == ' ' || r == '\t'
	})

	var ids []uint64
	for _, p := range parts {
		if id, err := strconv.ParseUint(strings.TrimSpace(p), 10, 64); err == nil {
			ids = append(ids, id)
		}
	}
	return ids
}

func printExam(e *exampb.Exam) {
	if e == nil {
		fmt.Println("Exam not found")
		return
	}
	fmt.Printf("ID: %d\nTitle: %s\nDescription: %s\nSection IDs: %v\n\n",
		e.Id, e.Title, e.Description, e.SectionIds)
}
