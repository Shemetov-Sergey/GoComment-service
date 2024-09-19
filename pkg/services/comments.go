package services

import (
	"context"
	"net/http"

	"github.com/Shemetov-Sergey/GoComment-service/pkg/db"
	"github.com/Shemetov-Sergey/GoComment-service/pkg/models"
	"github.com/Shemetov-Sergey/GoComment-service/pkg/pb"
)

type Server struct {
	H db.Handler
}

func (s *Server) CreateComment(ctx context.Context, req *pb.CreateCommentRequest) (*pb.CreateCommentResponse, error) {
	c := &models.Comment{
		NewsId:   req.NewsId,
		ParentId: req.ParentId,
		Text:     req.Text,
		Censored: req.Censored,
	}

	if result := s.H.DB.Create(c); result.Error != nil {
		return &pb.CreateCommentResponse{
			Status: http.StatusConflict,
			Error:  result.Error.Error(),
		}, nil
	}

	return &pb.CreateCommentResponse{
		Status: http.StatusCreated,
		Id:     uint64(c.ID),
	}, nil
}

func (s *Server) CommentsByNews(ctx context.Context, req *pb.CommentsByNewsRequest) (*pb.CommentsByNewsResponse, error) {
	commentsSlice := make([]*pb.Comment, 0)
	if result := s.H.DB.Where(&models.Comment{NewsId: req.NewsId}).Find(&commentsSlice); result.Error != nil {
		return &pb.CommentsByNewsResponse{
			Status: http.StatusNotFound,
			Error:  result.Error.Error(),
		}, nil
	}

	return &pb.CommentsByNewsResponse{
		Status:   http.StatusOK,
		Comments: commentsSlice,
	}, nil
}
