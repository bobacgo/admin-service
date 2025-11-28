package basic

import "context"

type Service struct{}

func NewService() *Service {
	return &Service{}
}

// Get /health 健康检查
func (svc *Service) GetHealth(_ context.Context, _ any) (any, error) {
	return nil, nil
}
