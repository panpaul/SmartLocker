package main

import (
	"SmartLocker/cmd/server/router"
	"SmartLocker/config"
	"SmartLocker/logger"
	"SmartLocker/model"
	"SmartLocker/service/auth"
	"github.com/go-playground/log"
	"github.com/urfave/cli"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// init the helpers
	logger.Setup()
	var daba string
	// create an app instance
	app := cli.NewApp()
	app.Name = "SmartLocker"
	app.Usage = "the backend of SmartLocker"
	app.Version = getVersion()
	app.Commands = []cli.Command{
		{
			Name:    "server",
			Aliases: []string{"s"},
			Usage:   "Start the Server",
			Flags  : []cli.Flag{
				cli.StringFlag{
					Name		: "database, d",
					Usage		: "Input the database",
					Value		: "sqlite",
					Destination	: &daba,
				},
			},
			Action:  func (c *cli.Context) error {
				b := false
				if (daba == "sqlite" || daba == "SQLite") || daba == "Sqlite" {
					config.Setup("config.yaml")
					b = true
				}
				if (daba == "mysql" || daba == "Mysql") || daba == "MySQL" {
					config.Setup("config_mysql.yaml")
					b = true
				}
				if b == false {
					log.Fatal("Unknown database, please input \"-d SQLite\" or \"-d MySQL\"")
				}
				
				model.Setup()
				auth.JwtSetup()
				StartServer(c)
				return nil
			},
		},
		}

	err := app.Run(os.Args)

	if err != nil {
		log.WithError(err).Fatal("Couldn't startup")
	}

}

func StartServer(ctx *cli.Context) {
	// init the route
	routersInit := router.InitRouters()
	addr := config.Conf.WebServer.Address + ":" + config.Conf.WebServer.Port

	// setup a http server instance
	server := &http.Server{
		Addr:    addr,
		Handler: routersInit,
	}

	// start the server
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.WithError(err).Fatal("Couldn't start server")
		}
	}()

	// shutdown
	// Handle SIGINT and SIGTERM.
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Debug(<-ch)

	// Stop the service gracefully.
	model.CloseDB()
	err := server.Shutdown(nil)
	if err != nil {
		log.WithError(err).Warn("Couldn't shutdown the server")
	}
}

var (
	BuildTimeStamp = "None"
	GitHash        = "None"
	Version        = "0.1"
)

func getVersion() string {
	return Version + "+" + GitHash + "+" + BuildTimeStamp
}
