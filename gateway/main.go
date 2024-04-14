package main

import (
	"flag"
	"log"
	"os"

	"github.com/keshu12345/guardianlink/gateway/config"
	"github.com/keshu12345/guardianlink/gateway/db"
	"github.com/keshu12345/guardianlink/gateway/server/router"
	"github.com/keshu12345/guardianlink/gateway/service"
	"go.uber.org/fx"
)

var configDirPath = flag.String("config", "", "path for config dir")

func main() {

	flag.Parse()
	log.New(os.Stdout, "", 0)
	app := fx.New(
		config.NewFxModule(*configDirPath, ""),
		router.Module,
		db.Module,
		service.Module,
	)
	app.Run()
}
