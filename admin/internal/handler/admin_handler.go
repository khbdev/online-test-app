package handler

import (
	"context"

	pb "github.com/khbdev/proto-online-test/proto/admin"

	"admin-service/internal/service"
	
)

type AdminHandler struct {
	pb.UnimplementedAdminServiceServer
	service *service.AdminService
}

func NewAdminHandler(s *service.AdminService) *AdminHandler {
	return &AdminHandler{service: s}
}


// CreateAdmin
func (h *AdminHandler) CreateAdmin(ctx context.Context, req *pb.CreateAdminRequest) (*pb.CreateAdminResponse, error) {
	admin, err := h.service.Create(req.Username, req.Password)
	if err != nil {
		return nil, err
	}

	return &pb.CreateAdminResponse{
		Admin: &pb.Admin{
			Id:        uint64(admin.ID),
			Username:  admin.Username,
			Password:  admin.Password,
			CreatedAt: admin.CreatedAt.String(),
			UpdatedAt: admin.UpdatedAt.String(),
		},
	}, nil
}

// GetAdminList
func (h *AdminHandler) GetAdminList(ctx context.Context, _ *pb.GetAdminListRequest) (*pb.GetAdminListResponse, error) {
	admins, err := h.service.GetAll()
	if err != nil {
		return nil, err
	}

	var pbAdmins []*pb.Admin
	for _, a := range admins {
		pbAdmins = append(pbAdmins, &pb.Admin{
			Id:        uint64(a.ID),
			Username:  a.Username,
			Password:  a.Password,
			CreatedAt: a.CreatedAt.String(),
			UpdatedAt: a.UpdatedAt.String(),
		})
	}

	return &pb.GetAdminListResponse{Admins: pbAdmins}, nil
}


func (h *AdminHandler) GetAdminByID(ctx context.Context, req *pb.GetAdminByIDRequest) (*pb.GetAdminByIDResponse, error) {
	admin, err := h.service.GetByID(uint(req.Id))
	if err != nil {
		return nil, err
	}
	if admin == nil {
		return &pb.GetAdminByIDResponse{}, nil
	}

	return &pb.GetAdminByIDResponse{
		Admin: &pb.Admin{
			Id:        uint64(admin.ID),
			Username:  admin.Username,
			Password:  admin.Password,
			CreatedAt: admin.CreatedAt.String(),
			UpdatedAt: admin.UpdatedAt.String(),
		},
	}, nil
}

func (h *AdminHandler) UpdateAdmin(ctx context.Context, req *pb.UpdateAdminRequest) (*pb.UpdateAdminResponse, error) {
	admin, err := h.service.Update(uint(req.Id), req.Username, req.Password)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateAdminResponse{
		Admin: &pb.Admin{
			Id:        uint64(admin.ID),
			Username:  admin.Username,
			Password:  admin.Password,
			CreatedAt: admin.CreatedAt.String(),
			UpdatedAt: admin.UpdatedAt.String(),
		},
	}, nil
}


func (h *AdminHandler) DeleteAdmin(ctx context.Context, req *pb.DeleteAdminRequest) (*pb.DeleteAdminResponse, error) {
	err := h.service.Delete(uint(req.Id))
	if err != nil {
		return &pb.DeleteAdminResponse{Success: false}, err
	}
	return &pb.DeleteAdminResponse{Success: true}, nil
}


func (h *AdminHandler) VerifyAdmin(ctx context.Context, req *pb.VerifyAdminRequest) (*pb.VerifyAdminResponse, error) {
	valid, err := h.service.VerifyAdmin(req.Username, req.Password)
	if err != nil {
		return nil, err
	}
	return &pb.VerifyAdminResponse{Valid: valid}, nil
}
