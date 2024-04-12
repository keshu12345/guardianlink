package main

import (
	"flag"
	"log"
	"os"

	"github.com/keshu12345/guardianlink/nodeb/config"
	"github.com/keshu12345/guardianlink/nodeb/db"
	"github.com/keshu12345/guardianlink/nodeb/server/router"
	"github.com/keshu12345/guardianlink/nodeb/service"
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
		// notesService.Module,
	)
	app.Run()
}
