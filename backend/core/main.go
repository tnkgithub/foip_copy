package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"

	// pkg
	"foip/core/config"
	"foip/core/pkg/server"
	"foip/core/version"
)

func main() {
	app := cli.App{
		Name:  "server",
		Usage: "application",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "config-path",
				Usage:    "config file path",
				Required: false,
			},
		},
		Action:  startServer,
		Version: version.Version,
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}

func startServer(cxt *cli.Context) error {
	cfg, err := config.New(cxt)
	if err != nil {
		return err
	}

	server, err := server.New(cfg)
	if err != nil {
		return err
	}

	return server.Start(cfg.Server.Port)
}
