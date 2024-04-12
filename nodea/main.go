package main

import (
	"flag"
	"log"
	"os"

	"github.com/keshu12345/guardianlink/nodea/config"
	"github.com/keshu12345/guardianlink/nodea/db"
	"github.com/keshu12345/guardianlink/nodea/server/router"
	"github.com/keshu12345/guardianlink/nodea/service"
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
