package main

import (
	"context"
	"log"
	"time"

	jobpb "github.com/khbdev/proto-online-test/proto/job"
	"google.golang.org/grpc"
)

func main() {
	// gRPC serverga ulanish
	conn, err := grpc.Dial("localhost:50056", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("❌ gRPC serverga ulanish xato: %v", err)
	}
	defer conn.Close()

	client := jobpb.NewJobServiceClient(conn)

	// Request yaratish
	req := &jobpb.SubmitTestRequest{
	FirstName:  "Azizbek",
	LastName:   "Xasanov",
	Phone:      "+998901234567",
	Email:      "azizbek@example.com",
	TgUsername: "@azizbek",
	Key:        "10-ta-oldimi-yoqmi-2_1761244224091540466",
	Sections: []*jobpb.SectionAnswer{
		{
			SectionId: 2, // backend bo‘limi
			Questions: []*jobpb.QuestionAnswer{
				{QuestionId: 7, OptionIds: []uint64{20}},    // JSON va XML farqi — to‘g‘ri javob 20
				{QuestionId: 21, OptionIds: nil},            // Firewall nima — javob yo‘q
				{QuestionId: 18, OptionIds: nil},            // Proxy server — javob yo‘q
				{QuestionId: 20, OptionIds: nil},            // Ping buyrug‘i — javob yo‘q
				{QuestionId: 15, OptionIds: nil},            // WebSocket — javob yo‘q
				{QuestionId: 13, OptionIds: nil},            // Cookie va Session — javob yo‘q
				{QuestionId: 11, OptionIds: []uint64{36}},   // API versiyalash — to‘g‘ri javob 36
				{QuestionId: 28, OptionIds: nil},            // RPC — javob yo‘q
				{QuestionId: 10, OptionIds: []uint64{32}},   // Header va Body — to‘g‘ri javob 32
				{QuestionId: 29, OptionIds: nil},            // Network latency — javob yo‘q
			},
		},
	},
}


	// gRPC chaqirish
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := client.SubmitTest(ctx, req)
	if err != nil {
		log.Fatalf("❌ SubmitTest xato: %v", err)
	}

	log.Printf("✅ Javob keldi: status=%s, message=%s, key=%s", res.Status, res.Message, res.Key)
}
