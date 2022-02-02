package cmd

import (
	"errors"
	"fmt"

	"github.com/urfave/cli/v2"
)

func Execute(version string, args []string) error {
	app := cli.App{
		Name:        "bump-version",
		Usage:       "Bumps the specified version",
		Description: "Bumps the specified version accordingly semantic versioning specification",
		UsageText:   "bump-version [options...] <version>",
        Version: version,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "fragment",
				Aliases: []string{"f"},
				Usage:   "The versions fragment you want to increment. Possible options: [major | feature | bug | alpha | beta | rc]",
			},
		},

		Action: run,
	}

	return app.Run(args)
}

func run(c *cli.Context) error {
    version := c.Args().First()
    if version == "" {
        return errors.New("the version params must be specified")
    }

	fmt.Println(version)

	return nil
}
