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

	var configName string

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
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "config",
					Usage:       "the position of the config",
					Value:       "config.yaml",
					Destination: &configName,
				},
			},
			Action: func(c *cli.Context) error {
				config.Setup(configName)
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
