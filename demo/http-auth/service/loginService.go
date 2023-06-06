package service

import (
	"learning-http-auth/domain"
	"learning-http-auth/errs"
	"learning-http-auth/logger"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type DefaultAuthService struct {
	repo domain.AuthRepositoryDb
}

const HMAC_SAMPLE_SECRET = "hmacSampleSecret"

func (s *DefaultAuthService) Login(req domain.Login) (string, *errs.AppError) {
	if appErr := s.repo.FindBy(req.UserId, req.Password); appErr != nil {
		return "", appErr
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": req.UserId,
		"exp":    jwt.NewNumericDate(time.Now().Add(time.Minute * 5)),
	})

	signedToken, err := token.SignedString([]byte(HMAC_SAMPLE_SECRET))
	if err != nil {
		logger.Info("not able to sign token, " + err.Error())
		return "", errs.NewAuthenticationError("Not able to sign token")
	}

	return signedToken, nil
}

func (s *DefaultAuthService) Verify(token string) *errs.AppError {
	// TODO
	return nil
}

func NewLoginService(repo domain.AuthRepositoryDb) DefaultAuthService {
	return DefaultAuthService{repo}
}
