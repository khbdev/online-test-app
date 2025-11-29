package handler

import (
	"context"
	"log"

	"queue-job-service/internal/service"

	jobpb "github.com/khbdev/proto-online-test/proto/job"
)

// ✅ Shu qatorda embedding qo‘shiladi
type JobServiceServer struct {
	jobpb.UnimplementedJobServiceServer
}

func (s *JobServiceServer) SubmitTest(ctx context.Context, req *jobpb.SubmitTestRequest) (*jobpb.SubmitTestResponse, error) {
	log.Printf("📥 Yangi test natijasi keldi: key=%s, user=%s %s", req.Key, req.FirstName, req.LastName)

	// JobRequest yaratish
	jobReq := &service.JobRequest{
		FirstName:  req.FirstName,
		LastName:   req.LastName,
		Phone:      req.Phone,
		Email:      req.Email,
		TgUsername: req.TgUsername,
		Key:        req.Key,
	}

	// Sections ni to'ldirish
	for _, sec := range req.Sections {
		section := struct {
			SectionID int `json:"section_id"`
			Questions []struct {
				QuestionID int   `json:"question_id"`
				OptionIDs  []int `json:"option_ids"`
			} `json:"questions"`
		}{
			SectionID: int(sec.SectionId),
		}

		for _, q := range sec.Questions {
			optionIDs := make([]int, len(q.OptionIds))
			for i, id := range q.OptionIds {
				optionIDs[i] = int(id)
			}
			question := struct {
				QuestionID int   `json:"question_id"`
				OptionIDs  []int `json:"option_ids"`
			}{
				QuestionID: int(q.QuestionId),
				OptionIDs:  optionIDs,
			}
			section.Questions = append(section.Questions, question)
		}

		jobReq.Sections = append(jobReq.Sections, section)
	}

	// 🔹 Faqat publish qilish
	if err := service.PublishJob(jobReq); err != nil {
		log.Printf("❌ PublishJob xato: %v", err)
		return &jobpb.SubmitTestResponse{
			Status:  "error",
			Message: "Test natijasi queue’ga yuborilmadi",
			Key:     req.Key,
		}, nil
	}

	return &jobpb.SubmitTestResponse{
		Status:  "success",
		Message: "Test natijasi muvaffaqiyatli qabul qilindi",
		Key:     req.Key,
	}, nil
}
