package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	userpb "github.com/khbdev/proto-online-test/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Komanda flaglari
	action := flag.String("action", "", "create|get|show|update|delete")
	id := flag.Uint("id", 0, "User ID (for get/show/update/delete)")
	firstName := flag.String("first", "", "First name")
	lastName := flag.String("last", "", "Last name")
	phone := flag.String("phone", "", "Phone number")
	email := flag.String("email", "", "Email")
	tgUsername := flag.String("tg", "", "Telegram username")

	flag.Parse()

	// gRPC serverga ulanamiz
	conn, err := grpc.Dial("localhost:50053", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("❌ Serverga ulanishda xato: %v", err)
	}
	defer conn.Close()

	client := userpb.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 🟢 CREATE
	if *action == "create" {
		req := &userpb.CreateUserRequest{
			FirstName:       *firstName,
			LastName:        *lastName,
			Phone:           *phone,
			Email:           *email,
			TgUsername:      *tgUsername,
			Bolimlar:        `["Backend", "DevOps"]`,
			Savollar:        `["Q1", "Q2"]`,
			Javoblar:        `["A", "B"]`,
			TogriJavoblar:   2,
			NatogriJavoblar: 0,
		}
		res, err := client.CreateUser(ctx, req)
		if err != nil {
			log.Fatalf("Create xato: %v", err)
		}
		fmt.Printf("✅ User yaratildi: %+v\n", res.User)
		return
	}

	// 🟡 GET ALL
	if *action == "get" {
		res, err := client.GetAllUsers(ctx, &userpb.GetAllUsersRequest{})
		if err != nil {
			log.Fatalf("GetAll xato: %v", err)
		}
		for _, u := range res.Users {
			fmt.Printf("👤 ID:%d | %s %s | %s\n", u.Id, u.FirstName, u.LastName, u.Phone)
		}
		return
	}

	// 🔵 SHOW BY ID
	if *action == "show" {
		res, err := client.GetUserByID(ctx, &userpb.GetUserByIDRequest{Id: uint64(*id)})
		if err != nil {
			log.Fatalf("Show xato: %v", err)
		}
		fmt.Printf("📄 User: %+v\n", res.User)
		return
	}

	// 🟠 UPDATE
	if *action == "update" {
		req := &userpb.UpdateUserRequest{
			Id:              uint64(*id),
			FirstName:       *firstName,
			LastName:        *lastName,
			Phone:           *phone,
			Email:           *email,
			TgUsername:      *tgUsername,
			Bolimlar:        `["Frontend"]`,
			Savollar:        `["Q3"]`,
			Javoblar:        `["C"]`,
			TogriJavoblar:   1,
			NatogriJavoblar: 0,
		}
		res, err := client.UpdateUser(ctx, req)
		if err != nil {
			log.Fatalf("Update xato: %v", err)
		}
		fmt.Printf("✏️ Yangilandi: %+v\n", res.User)
		return
	}

	// 🔴 DELETE
	if *action == "delete" {
		res, err := client.DeleteUser(ctx, &userpb.DeleteUserRequest{Id: uint64(*id)})
		if err != nil {
			log.Fatalf("Delete xato: %v", err)
		}
		fmt.Printf("🗑️ O‘chirildi: %v\n", res.Success)
		return
	}

	log.Println("❗ Iltimos --action flagini kiriting: create|get|show|update|delete")
}
