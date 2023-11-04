package service

import (
	"context"

	"project/internal/models"
)

func (s *Service) ViewJobById(ctx context.Context, jid uint64) (models.Jobs, error) {
	jobData, err := s.UserRepo.Jobbyjid(ctx, jid)
	if err != nil {
		return models.Jobs{}, nil
	}
	return jobData, nil
}

func (s *Service) ViewAllJobs(ctx context.Context) ([]models.Jobs, error) {
	jobDatas, err := s.UserRepo.FetchAllJobs(ctx)
	if err != nil {
		return nil, err
	}
	return jobDatas, nil

}

func (s *Service) AddJobDetails(ctx context.Context, jobData models.Jobs, cid uint64) (models.Jobs, error) {
	jobData.Cid = uint(cid)
	jobData, err := s.UserRepo.CreateUserJob(ctx, jobData)
	if err != nil {
		return models.Jobs{}, err
	}
	return jobData, nil
}

func (s *Service) ViewJob(ctx context.Context, cid uint64) ([]models.Jobs, error) {
	jobData, err := s.UserRepo.Jobbycid(ctx, cid)
	if err != nil {
		return nil, err
	}
	return jobData, nil
}
