package interfaces

import "github.com/chulista/weather_api/domain"

type RedirectRepository interface {
	Find(code string) (*domain.Redirect, error)
	Store(redirect *domain.Redirect) error
}