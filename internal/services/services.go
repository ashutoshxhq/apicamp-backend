package services

import context "context"

// Server represents the gRPC server
type Server struct {
}

// GetService ...
func (s *Server) GetService(ctx context.Context, req *GetServiceRequest) (*GetServiceResponse, error) {

	return &GetServiceResponse{}, nil
}

// GetServices ...
func (s *Server) GetServices(ctx context.Context, req *GetServicesRequest) (*GetServicesResponse, error) {

	return &GetServicesResponse{}, nil
}

// CreateService ...
func (s *Server) CreateService(ctx context.Context, req *CreateServiceRequest) (*CreateServiceResponse, error) {

	return &CreateServiceResponse{}, nil
}

// CreateServices ...
func (s *Server) CreateServices(ctx context.Context, req *CreateServicesRequest) (*CreateServicesResponse, error) {

	return &CreateServicesResponse{}, nil
}

// UpdateService ...
func (s *Server) UpdateService(ctx context.Context, req *UpdateServiceRequest) (*UpdateServiceResponse, error) {

	return &UpdateServiceResponse{}, nil
}

// UpdateServices ...
func (s *Server) UpdateServices(ctx context.Context, req *UpdateServicesRequest) (*UpdateServicesResponse, error) {

	return &UpdateServicesResponse{}, nil
}

// DeleteService ...
func (s *Server) DeleteService(ctx context.Context, req *DeleteServiceRequest) (*DeleteServiceResponse, error) {

	return &DeleteServiceResponse{}, nil
}

// DeleteServices ...
func (s *Server) DeleteServices(ctx context.Context, req *DeleteServicesRequest) (*DeleteServicesResponse, error) {

	return &DeleteServicesResponse{}, nil
}
