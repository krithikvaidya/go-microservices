package service

import (
	"learning-http-auth/domain"
	"learning-http-auth/errs"
)

type DefaultAuthService struct {
	repo domain.AuthRepositoryDb
}

const HMAC_SAMPLE_SECRET = "hmacSampleSecret"

func (s *DefaultAuthService) Login(req domain.Login) (string, *errs.AppError) {

	return "", nil
}

func NewLoginService(repo domain.AuthRepositoryDb) DefaultAuthService {
	return DefaultAuthService{repo}
}
