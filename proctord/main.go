package main

import (
	"github.com/getsentry/raven-go"
	"github.com/gojektech/proctor/proctord/config"
	"os"

	"github.com/gojektech/proctor/proctord/logger"
	"github.com/gojektech/proctor/proctord/scheduler"
	"github.com/gojektech/proctor/proctord/server"
	"github.com/gojektech/proctor/proctord/storage/postgres"

	"github.com/urfave/cli"
)

func main() {
	logger.Setup()
	raven.SetDSN(config.SentryDSN())

	proctord := cli.NewApp()
	proctord.Name = "proctord"
	proctord.Usage = "Handle executing jobs and maintaining their configuration"
	proctord.Version = "0.2.0"
	proctord.Commands = []cli.Command{
		{
			Name:        "migrate",
			Description: "Run database migrations for proctord",
			Action: func(c *cli.Context) {
				err := postgres.Up()
				if err != nil {
					panic(err.Error())
				}
				logger.Info("Migration successful")
			},
		},
		{
			Name:        "rollback",
			Description: "Rollback database migrations by one step for proctord",
			Action: func(c *cli.Context) {
				err := postgres.DownOneStep()
				if err != nil {
					panic(err.Error())
				}
				logger.Info("Rollback successful")
			},
		},
		{
			Name:    "start",
			Aliases: []string{"s"},
			Usage:   "starts server",
			Action: func(c *cli.Context) error {
				return server.Start()
			},
		},
		{
			Name:  "start-scheduler",
			Usage: "starts scheduler",
			Action: func(c *cli.Context) error {
				return scheduler.Start()
			},
		},
	}

	proctord.Run(os.Args)
}
