package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	pb "github.com/khbdev/proto-online-test/proto/admin"
	"google.golang.org/grpc"
)

func main() {
	action := flag.String("action", "", "create|get|show|update|delete|verify")
	id := flag.Uint64("id", 0, "Admin ID")
	username := flag.String("username", "", "Username")
	password := flag.String("password", "", "Password")
	flag.Parse()

	conn, err := grpc.Dial("admin-service:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("❌ Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewAdminServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	switch *action {
	case "create":
		res, err := client.CreateAdmin(ctx, &pb.CreateAdminRequest{
			Username: *username,
			Password: *password,
		})
		checkErr(err)
		fmt.Printf("✅ Created Admin: %+v\n", res.Admin)

	case "get":
		res, err := client.GetAdminList(ctx, &pb.GetAdminListRequest{})
		checkErr(err)
		fmt.Println("✅ Admin List:")
		for _, a := range res.Admins {
			fmt.Printf("- ID:%d Username:%s CreatedAt:%s\n", a.Id, a.Username, a.CreatedAt)
		}

	case "show":
		res, err := client.GetAdminByID(ctx, &pb.GetAdminByIDRequest{Id: *id})
		checkErr(err)
		fmt.Printf("✅ Admin Found: %+v\n", res.Admin)

	case "update":
		res, err := client.UpdateAdmin(ctx, &pb.UpdateAdminRequest{
			Id:       *id,
			Username: *username,
			Password: *password,
		})
		checkErr(err)
		fmt.Printf("✅ Updated Admin: %+v\n", res.Admin)

	case "delete":
		res, err := client.DeleteAdmin(ctx, &pb.DeleteAdminRequest{Id: *id})
		checkErr(err)
		if res.Success {
			fmt.Println("✅ Admin deleted successfully.")
		} else {
			fmt.Println("⚠️ Delete failed.")
		}

	case "verify":
		res, err := client.VerifyAdmin(ctx, &pb.VerifyAdminRequest{
			Username: *username,
			Password: *password,
		})
		checkErr(err)
		if res.Valid {
			fmt.Println("✅ Username & password are valid.")
		} else {
			fmt.Println("❌ Invalid credentials.")
		}

	default:
		fmt.Println(`⚙️ Usage examples:
  go run main.go -action=create -username=admin -password=123
  go run main.go -action=get
  go run main.go -action=show -id=1
  go run main.go -action=update -id=1 -username=new -password=abc
  go run main.go -action=delete -id=1
  go run main.go -action=verify -username=admin -password=123`)
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatalf("❌ Error: %v", err)
	}
}
