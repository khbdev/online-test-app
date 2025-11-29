package handler

import (
	"context"
	"test-section-service/internal/models"
	"test-section-service/internal/service"

	pb "github.com/khbdev/proto-online-test/proto/test"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
)

type TestHandler struct {
	pb.UnimplementedTestServiceServer
	SectionService  *service.SectionService
	QuestionService *service.QuestionService
	OptionService   *service.OptionService
}

func NewTestHandler(sectionService *service.SectionService, questionService *service.QuestionService, optionService *service.OptionService) *TestHandler {
	return &TestHandler{
		SectionService:  sectionService,
		QuestionService: questionService,
		OptionService:   optionService,
	}
}

// --- Section Methods ---

func (h *TestHandler) CreateSection(ctx context.Context, req *pb.CreateSectionRequest) (*pb.Section, error) {
	section := models.Section{Name: req.Name}
	if err := h.SectionService.CreateSection(&section); err != nil {
		return nil, err
	}
	return &pb.Section{
		Id:   uint64(section.ID),
		Name: section.Name,
	}, nil
}

func (h *TestHandler) GetSectionById(ctx context.Context, req *pb.SectionID) (*pb.GetSectionResponse, error) {
	section, err := h.SectionService.GetSectionByID(uint(req.Id))
	if err != nil {
		return nil, err
	}
	return &pb.GetSectionResponse{
		Section: &pb.Section{
			Id:   uint64(section.ID),
			Name: section.Name,
		},
	}, nil
}

func (h *TestHandler) GetAllSections(ctx context.Context, _ *emptypb.Empty) (*pb.GetAllSectionsResponse, error) {
	sections, err := h.SectionService.GetAllSections()
	if err != nil {
		return nil, err
	}
	var pbSections []*pb.Section
	for _, s := range sections {
		pbSections = append(pbSections, &pb.Section{
			Id:   uint64(s.ID),
			Name: s.Name,
		})
	}
	return &pb.GetAllSectionsResponse{Sections: pbSections}, nil
}

func (h *TestHandler) UpdateSection(ctx context.Context, req *pb.UpdateSectionRequest) (*pb.Section, error) {
	section, err := h.SectionService.GetSectionByID(uint(req.Id))
	if err != nil {
		return nil, err
	}
	section.Name = req.Name
	if err := h.SectionService.UpdateSection(section); err != nil {
		return nil, err
	}
	return &pb.Section{
		Id:   uint64(section.ID),
		Name: section.Name,
	}, nil
}

func (h *TestHandler) DeleteSection(ctx context.Context, req *pb.SectionID) (*emptypb.Empty, error) {
	if err := h.SectionService.DeleteSection(uint(req.Id)); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (h *TestHandler) GetFullSectionStructure(ctx context.Context, req *pb.GetFullSectionRequest) (*pb.GetFullSectionResponse, error) {
	section, err := h.SectionService.GetFullSectionStructure(uint(req.SectionId))
	if err != nil {
		return nil, err
	}
	var pbQuestions []*pb.Question
	for _, q := range section.Questions {
		var pbOptions []*pb.Option
		for _, o := range q.Options {
			pbOptions = append(pbOptions, &pb.Option{
				Id:         uint64(o.ID),
				Text:       o.Text,
				IsCorrect:  o.IsCorrect,
				QuestionId: uint64(o.QuestionID),
			})
		}
		pbQuestions = append(pbQuestions, &pb.Question{
			Id:        uint64(q.ID),
			Text:      q.Text,
			SectionId: uint64(q.SectionID),
			Options:   pbOptions,
		})
	}
	return &pb.GetFullSectionResponse{
		Section: &pb.Section{
			Id:        uint64(section.ID),
			Name:      section.Name,
			Questions: pbQuestions,
		},
	}, nil
}

// --- Question Methods ---

func (h *TestHandler) CreateQuestion(ctx context.Context, req *pb.CreateQuestionRequest) (*pb.Question, error) {
	question := &models.Question{
		SectionID: uint(req.SectionId),
		Text:      req.Text,
	}
	created, err := h.QuestionService.Create(question)
	if err != nil {
		return nil, err
	}
	return &pb.Question{
		Id:        uint64(created.ID),
		SectionId: uint64(created.SectionID),
		Text:      created.Text,
	}, nil
}

func (h *TestHandler) GetQuestionById(ctx context.Context, req *pb.QuestionID) (*pb.GetQuestionResponse, error) {
	q, err := h.QuestionService.GetByID(uint(req.Id))
	if err != nil {
		return nil, err
	}
	return &pb.GetQuestionResponse{
		Question: &pb.Question{
			Id:        uint64(q.ID),
			SectionId: uint64(q.SectionID),
			Text:      q.Text,
		},
	}, nil
}

func (h *TestHandler) GetAllQuestions(ctx context.Context, _ *emptypb.Empty) (*pb.GetAllQuestionsResponse, error) {
	questions, err := h.QuestionService.GetAll()
	if err != nil {
		return nil, err
	}
	var list []*pb.Question
	for _, q := range questions {
		list = append(list, &pb.Question{
			Id:        uint64(q.ID),
			SectionId: uint64(q.SectionID),
			Text:      q.Text,
		})
	}
	return &pb.GetAllQuestionsResponse{Questions: list}, nil
}

func (h *TestHandler) UpdateQuestion(ctx context.Context, req *pb.UpdateQuestionRequest) (*pb.Question, error) {
	q := &models.Question{
		Model:     gorm.Model{ID: uint(req.Id)},
		SectionID: uint(req.SectionId),
		Text:      req.Text,
	}
	updated, err := h.QuestionService.Update(q)
	if err != nil {
		return nil, err
	}
	return &pb.Question{
		Id:        uint64(updated.ID),
		SectionId: uint64(updated.SectionID),
		Text:      updated.Text,
	}, nil
}

func (h *TestHandler) DeleteQuestion(ctx context.Context, req *pb.QuestionID) (*emptypb.Empty, error) {
	if err := h.QuestionService.Delete(uint(req.Id)); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

// --- Option Methods ---

func (h *TestHandler) CreateOption(ctx context.Context, req *pb.CreateOptionRequest) (*pb.Option, error) {
	option := &models.Option{
		Text:       req.Text,
		IsCorrect:  req.IsCorrect,
		QuestionID: uint(req.QuestionId),
	}
	if err := h.OptionService.CreateOption(option); err != nil {
		return nil, err
	}
	return &pb.Option{
		Id:         uint64(option.ID),
		Text:       option.Text,
		IsCorrect:  option.IsCorrect,
		QuestionId: uint64(option.QuestionID),
	}, nil
}

func (h *TestHandler) GetOptionById(ctx context.Context, req *pb.OptionID) (*pb.GetOptionResponse, error) {
	option, err := h.OptionService.GetOptionByID(uint(req.Id))
	if err != nil {
		return nil, err
	}
	return &pb.GetOptionResponse{
		Option: &pb.Option{
			Id:         uint64(option.ID),
			Text:       option.Text,
			IsCorrect:  option.IsCorrect,
			QuestionId: uint64(option.QuestionID),
		},
	}, nil
}

func (h *TestHandler) GetAllOptions(ctx context.Context, _ *emptypb.Empty) (*pb.GetAllOptionsResponse, error) {
	options, err := h.OptionService.GetAllOptions()
	if err != nil {
		return nil, err
	}
	var res []*pb.Option
	for _, o := range options {
		res = append(res, &pb.Option{
			Id:         uint64(o.ID),
			Text:       o.Text,
			IsCorrect:  o.IsCorrect,
			QuestionId: uint64(o.QuestionID),
		})
	}
	return &pb.GetAllOptionsResponse{Options: res}, nil
}

func (h *TestHandler) UpdateOption(ctx context.Context, req *pb.UpdateOptionRequest) (*pb.Option, error) {
	option := &models.Option{
		Model:      gorm.Model{ID: uint(req.Id)},
		Text:       req.Text,
		IsCorrect:  req.IsCorrect,
		QuestionID: uint(req.QuestionId),
	}
	if err := h.OptionService.UpdateOption(option); err != nil {
		return nil, err
	}
	return &pb.Option{
		Id:         uint64(option.ID),
		Text:       option.Text,
		IsCorrect:  option.IsCorrect,
		QuestionId: uint64(option.QuestionID),
	}, nil
}

func (h *TestHandler) DeleteOption(ctx context.Context, req *pb.OptionID) (*emptypb.Empty, error) {
	if err := h.OptionService.DeleteOption(uint(req.Id)); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}