package services

import context "context"

// Server represents the gRPC server
type Server struct {
}

// GenerateServiceCode fuctions generates service code
func (s *Server) GenerateServiceCode(ctx context.Context, req *GenerateServiceCodeRequest) (*GenerateServiceCodeResponse, error) {

	return &GenerateServiceCodeResponse{Success: "generated service"}, nil
}
