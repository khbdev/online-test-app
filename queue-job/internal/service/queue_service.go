package service

import (
	"encoding/json"
	"fmt"
	"log"
	"queue-job-service/internal/config"
	"queue-job-service/internal/util/redisGet"

	rabbitmq "queue-job-service/internal/util/rabbitMq"
	"queue-job-service/internal/client"

	userpb "github.com/khbdev/proto-online-test/proto/user"
)

// JobRequest - handler’dan keladigan request modeli
type JobRequest struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	TgUsername string `json:"tg_username"`
	Key        string `json:"key"`
	Sections   []struct {
		SectionID int `json:"section_id"`
		Questions []struct {
			QuestionID int   `json:"question_id"`
			OptionIDs  []int `json:"option_ids"`
		} `json:"questions"`
	} `json:"sections"`
}

// 🔹 PublishJob - handler chaqiradigan function
func PublishJob(req *JobRequest) error {
	rmq := config.NewRabbitMQ()
	defer rmq.Conn.Close()
	defer rmq.Channel.Close()

	body, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("marshal xatolik: %v", err)
	}

	exchange := "direct_exchange"
	routingKey := "queue_key"

	if err := rabbitmq.PublishMessage(rmq.Channel, exchange, routingKey, string(body)); err != nil {
		return fmt.Errorf("publish xatolik: %v", err)
	}

	log.Printf("📨 Job queue’ga yuborildi: %s\n", req.Key)
	return nil
}

// 🔹 ConsumeJobs - queue’dan xabarlarni o‘qiydi va Redis bilan solishtirib, natijani user-service ga yuboradi
func ConsumeJobs() {
	rmq := config.NewRabbitMQ()
	defer rmq.Conn.Close()
	defer rmq.Channel.Close()

	queue := "task_queue"
	msgs, err := rmq.Channel.Consume(
		queue,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("❌ Queue consume xato: %v", err)
	}

	log.Println("👂 Job consumer ishlamoqda...")

	for msg := range msgs {
		var data JobRequest
		if err := json.Unmarshal(msg.Body, &data); err != nil {
			log.Println("❌ JSON parse xato:", err)
			continue
		}

		fmt.Printf("\n📥 Kelgan job key: %s\n", data.Key)

		// Redis’dan test natijasini olish
		testData, err := redisGet.GetByKey(data.Key)
		if err != nil {
			log.Printf("❌ Redis’dan ma’lumot topilmadi: %v\n", err)
			continue
		}

		fmt.Println("🔑 Redis’dan topildi:", testData)

		// ✅ Javoblarni solishtirish
		totalCorrect, totalWrong := CompareAnswers(&data, testData)

		// 🔢 Har bir to‘g‘ri javob uchun 10 ball
		score := totalCorrect * 10

		// 🔹 1. Bo‘limlarni tayyorlash (faqat ID va nom)
		var bolimlar []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		}

		// Redis ma’lumotini pars qilish
		var redisData RedisData
		if err := json.Unmarshal([]byte(testData), &redisData); err != nil {
			log.Printf("❌ Redis JSON parse xato: %v", err)
			continue
		}

		for _, sec := range redisData.Sections {
			bolimlar = append(bolimlar, struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			}{ID: sec.ID, Name: sec.Name})
		}
		bolimlarJSON, _ := json.Marshal(bolimlar)

		// 🔹 2. Savollarni tayyorlash (ID va matn)
		var savollar []struct {
			ID   int    `json:"id"`
			Text string `json:"text"`
		}

		// Foydalanuvchi javob bergan barcha savollar uchun
		for _, userSection := range data.Sections {
			for _, userQ := range userSection.Questions {
				// Redis’dan savol matnini topish
				for _, rSec := range redisData.Sections {
					for _, rQ := range rSec.Questions {
						if rQ.ID == userQ.QuestionID {
							savollar = append(savollar, struct {
								ID   int    `json:"id"`
								Text string `json:"text"`
							}{ID: rQ.ID, Text: rQ.Text})
							break
						}
					}
				}
			}
		}
		savollarJSON, _ := json.Marshal(savollar)

		// 🔹 3. Javoblarni tayyorlash (faqat foydalanuvchi tanlagan option ID’lari)
		var javoblar []struct {
			QuestionID int   `json:"question_id"`
			OptionIDs  []int `json:"option_ids"`
		}

		for _, userSection := range data.Sections {
			for _, userQ := range userSection.Questions {
				javoblar = append(javoblar, struct {
					QuestionID int   `json:"question_id"`
					OptionIDs  []int `json:"option_ids"`
				}{
					QuestionID: userQ.QuestionID,
					OptionIDs:  userQ.OptionIDs,
				})
			}
		}
		javoblarJSON, _ := json.Marshal(javoblar)

		// 🛰 User-service ga yuborish
		userClient := client.NewUserClient()
		defer userClient.Close()

		req := &userpb.CreateUserRequest{
			FirstName:       data.FirstName,
			LastName:        data.LastName,
			Phone:           data.Phone,
			Email:           data.Email,
			TgUsername:      data.TgUsername,
			Bolimlar:        string(bolimlarJSON),
			Savollar:        string(savollarJSON),
			Javoblar:        string(javoblarJSON),
			TogriJavoblar:   int32(totalCorrect),
			NatogriJavoblar: int32(totalWrong),
			ScorePercent:    int32(score), // ✅ har bir to‘g‘ri = 10 ball
			Description:     fmt.Sprintf("User test yakunladi. To‘g‘ri: %d, Noto‘g‘ri: %d, Ball: %d", totalCorrect, totalWrong, score),
		}

		res, err := userClient.CreateUser(req)
		if err != nil {
			log.Printf("❌ gRPC CreateUser xato: %v\n", err)
			continue
		}

		log.Printf("✅ User-service ga yuborildi. Javob: %v\n", res)
	}
}
