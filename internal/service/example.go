// Package service contains application business logic.
package service

import "context"

// ExampleRepository is an infrastructure boundary. Implement it in a repository package.
type ExampleRepository interface {
	Find(ctx context.Context, id string) (Example, error)
}

type Example struct {
	ID   string
	Name string
}

type ExampleService struct {
	repository ExampleRepository
}

func NewExampleService(repository ExampleRepository) *ExampleService {
	return &ExampleService{repository: repository}
}

func (s *ExampleService) Get(ctx context.Context, id string) (Example, error) {
	return s.repository.Find(ctx, id)
}
