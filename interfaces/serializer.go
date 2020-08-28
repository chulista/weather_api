package interfaces

import "github.com/chulista/weather_api/domain"
/**
metodos que se van a implementar para ser utilizados con Json y MsgPack
*/
type RedirectSerializer interface {
	Decode(input []byte) (*domain.Redirect, error)
	Encode(input *domain.Redirect) ([]byte, error)
}
