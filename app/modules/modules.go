package modules

import (
	"robbo-assets/package/assets"
)

type HandlerModule struct {
	AssetsHandler assets.Handler
}

func SetupHandler() HandlerModule {
	return HandlerModule{}
}
