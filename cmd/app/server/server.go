package server

import (
	"github.com/google/wire"
)

// ProviderSet is app providers.
var ProviderSet = wire.NewSet(NewGRPCServer, NewHTTPServer)
