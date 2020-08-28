package application

import (
	"errors"
	"github.com/chulista/weather_api/domain"
	"github.com/chulista/weather_api/interfaces"
	errs "github.com/pkg/errors"
	"github.com/teris-io/shortid"
	"gopkg.in/dealancer/validate.v2"
	"time"
)

var (
	ErrRedirectNotFound = errors.New("Redirect Not Found")
	ErrRedirectInvalid  = errors.New("Redirect Invalid")
)

type redirectService struct {
	redirectRepo interfaces.RedirectRepository
}

func NewRedirectService(redirectRepo interfaces.RedirectRepository) interfaces.RedirectService {
	return &redirectService{
		redirectRepo,
	}
}

func (r *redirectService) Find(code string) (*domain.Redirect, error) {
	return r.redirectRepo.Find(code)
}

func (r *redirectService) Store(redirect *domain.Redirect) error {
	if err := validate.Validate(redirect); err != nil {
		return errs.Wrap(ErrRedirectInvalid, "service.Redirect.Store")
	}
	redirect.Code = shortid.MustGenerate()
	redirect.CreatedAt = time.Now().UTC().Unix()
	return r.redirectRepo.Store(redirect)
}
