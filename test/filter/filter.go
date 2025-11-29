package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	filterpb "github.com/khbdev/proto-online-test/proto/filter"
)

func main() {
	// 1️⃣ gRPC serverga ulanamiz
	conn, err := grpc.Dial("localhost:50057", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("❌ Serverga ulana olmadim: %v", err)
	}
	defer conn.Close()

	client := filterpb.NewUserServiceClient(conn)

	// 2️⃣ So‘rov tayyorlaymiz
	req := &filterpb.FilterRequest{
		FirstName: "Azizbek",
		LastName:  "Xasanov",
		Phone:     "+998901234567",
		Year:      "2025",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 3️⃣ Serverdan ma’lumot olamiz
	resp, err := client.GetUsers(ctx, req)
	if err != nil {
		log.Fatalf("❌ RPC xato: %v", err)
	}

	// 4️⃣ JSON formatda chiqaramiz
	for _, user := range resp.Users {
		data, err := json.MarshalIndent(user, "", "  ")
		if err != nil {
			log.Fatalf("❌ JSON marshal xato: %v", err)
		}

		fmt.Println("───────────────────────────────")
		fmt.Println(string(data))
	}
}
