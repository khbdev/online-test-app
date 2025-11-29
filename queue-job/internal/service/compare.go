package service

import (
	"encoding/json"
	"fmt"
)

// Redis’dagi strukturani model qilib olish
type RedisOption struct {
	ID        int    `json:"id"`
	Text      string `json:"text"`
	IsCorrect bool   `json:"is_correct"`
}

type RedisQuestion struct {
	ID       int           `json:"id"`
	Text     string        `json:"text"`
	Options  []RedisOption `json:"options"`
}

type RedisSection struct {
	ID        int             `json:"id"`
	Name      string          `json:"name"`
	Questions []RedisQuestion `json:"questions"`
}

type RedisData struct {
	Name     string         `json:"name"`
	Sections []RedisSection `json:"sections"`
}

// CompareAnswers - foydalanuvchi javoblarini Redis’dagi to‘g‘ri javoblar bilan solishtiradi
func CompareAnswers(req *JobRequest, redisJSON string) (int, int) {
	var redisData RedisData
	if err := json.Unmarshal([]byte(redisJSON), &redisData); err != nil {
		fmt.Println("❌ Redis JSON parse xato:", err)
		return 0, 0
	}

	totalCorrect := 0
	totalWrong := 0

	for _, userSection := range req.Sections {
		for _, userQuestion := range userSection.Questions {
			var correctOptionIDs []int

			for _, redisSection := range redisData.Sections {
				for _, redisQuestion := range redisSection.Questions {
					if redisQuestion.ID == userQuestion.QuestionID {
						for _, opt := range redisQuestion.Options {
							if opt.IsCorrect {
								correctOptionIDs = append(correctOptionIDs, opt.ID)
							}
						}
					}
				}
			}

			if len(correctOptionIDs) == 0 {
				fmt.Printf("⚠️  Savol %d uchun Redis’da javob topilmadi\n", userQuestion.QuestionID)
				continue
			}

			if compareSlices(userQuestion.OptionIDs, correctOptionIDs) {
				totalCorrect++
			} else {
				totalWrong++
			}
		}
	}

	return totalCorrect, totalWrong
}

// compareSlices - 2 ta int slice bir xilmi (tartibdan qat’i nazar)
func compareSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	m := make(map[int]int)
	for _, v := range a {
		m[v]++
	}
	for _, v := range b {
		if m[v] == 0 {
			return false
		}
		m[v]--
	}
	return true
}
