package route

import (
	"melody-io/global-es/pkg/rsocket-route/internal/handle"

	"github.com/rsocket/rsocket-go"
)

func GetHandlers() []rsocket.OptAbstractSocket {
	return handle.Methods
}

func Add(routes map[string]interface{}) error {
	return handle.Paths(routes)
}
