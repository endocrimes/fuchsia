package main

import (
	"fmt"
	"os"

	"github.com/endocrimes/fuchsia/pkg/registry"
	"github.com/endocrimes/fuchsia/pkg/registry/consul"
	hclog "github.com/hashicorp/go-hclog"
	"github.com/urfave/cli"
)

var Version = ""

func main() {
	logger := hclog.New(&hclog.LoggerOptions{
		Name:  "fuchsia",
		Level: hclog.Debug,
	})

	app := cli.NewApp()

	app.Name = "fuchsia"
	app.Usage = "an experimental connect ingress load balancer"
	app.Version = Version

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "config",
			Usage:  "path to the configuration file",
			Value:  ".fuchsia.toml",
			EnvVar: "FUCHSIA_CONFIG",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "start",
			Usage: "start the load balancer",
			Action: func(c *cli.Context) error {
				return runAgent(logger)
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		logger.Error("failed", "error", err)
		os.Exit(2)
	}
}

func runAgent(logger hclog.Logger) error {
	logger.Info("Starting", "version", Version)

	providers := []registry.Registry{
		&consul.Registry{},
	}

	for _, p := range providers {
		err := p.Init()
		if err != nil {
			return fmt.Errorf("failed to init provider (%s): %v", p.Name(), err)
		}
	}

	return nil
}
