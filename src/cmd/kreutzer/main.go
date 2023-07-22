package main

import (
	r "github.com/leryn1122/kreutzer/v2/kreutzer/router"
	"github.com/leryn1122/kreutzer/v2/lib/webserver"
	"github.com/leryn1122/kreutzer/v2/support/version"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Version = version.CurrentVersion()
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:  "listen-port",
			Usage: "Listen to the specified web server port",
			Value: 8080,
		},
	}
	app.Action = func(ctx *cli.Context) error {
		return run(ctx)
	}
	app.ExitErrHandler = func(ctx *cli.Context, err error) {
		logrus.Fatal(err)
	}

	_ = app.Run(os.Args)
}

func run(ctx *cli.Context) error {
	logrus.Infof("Kreutzer version %s is running.", version.CurrentVersion())
	logrus.Infof("Kreutzer arguments: %s", ctx.Args())

	router := webserver.CreateRoute()

	r.RegisterRoutes(router)

	webserver.StartRestfulWebServer(router)
	return nil
}
