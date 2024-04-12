package service

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Invoke(RegisterEndpoint),
	fx.Provide(NewAuthService),
	fx.Provide(NewNodeAService),
	fx.Provide(ProvideMutex),
)
