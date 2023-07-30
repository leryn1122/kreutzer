package main

import (
	"github.com/leryn1122/kreutzer/v2/kreutzer/config"
	r "github.com/leryn1122/kreutzer/v2/kreutzer/router"
	"github.com/leryn1122/kreutzer/v2/lib/db"
	"github.com/leryn1122/kreutzer/v2/lib/grpc"
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
		cli.StringFlag{
			Name:   "config-file",
			EnvVar: "CONFIG_FILE",
			Value:  "conf/config.toml",
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

	configPath := ctx.String("config-file")
	config_, err := config.FromFile(configPath)
	if err != nil {
		return err
	}

	err = InitializeDatabase(config_.Database)
	if err != nil {
		return err
	}

	router := webserver.CreateRoute()

	r.RegisterRoutes(router)

	grpc.StartGrpcServer()

	webserver.StartRestfulWebServer(router)

	return nil
}

func InitializeDatabase(database config.Database) error {
	dbConfig := db.NewConfig(database.Host, database.Port, database.Username, database.Password)
	_, err := db.InitialDatabase(dbConfig)
	return err
}
