package interfaces

import (
	"github.com/chulista/weather_api/domain"
	//	"github.com/mercadolibre/fury_mpcs-preferences-api/internal/domain/model""
)

type RedirectService interface {
	Find(code string) (*domain.Redirect, error)
	Store(redirect *domain.Redirect) error
}
