package services

import (
	"context"
)

type Service struct {
	service Emp
}

func New(s Emp) Service {
	return Service{service: s}
}

func (s Service) Get(ctx context.Context, id int) bool {
	emp, err := s.service.GetById(ctx, id)

	if err != nil {
		return false
	}
	if emp.Age < 25 {
		return false
	}
	return true
}
