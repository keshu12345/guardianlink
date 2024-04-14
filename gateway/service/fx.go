package service

import (
	"github.com/keshu12345/guardianlink/gateway/service/auth"
	"github.com/keshu12345/guardianlink/gateway/service/ratelimitor"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Invoke(RegisterEndpoint),
	fx.Provide(ratelimitor.NewService),
	fx.Provide(auth.NewService),
)
