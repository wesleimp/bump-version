package cmd

import (
	"errors"
	"fmt"

	"github.com/urfave/cli/v2"
	"github.com/wesleimp/bump-version/internal/semver"
)

type Cli struct {
	Version   string
	Fragement string
}

func (c Cli) Validate() error {
	if c.Version == "" {
		return errors.New("the parameter `VERSION` must be specified")
	}

	if c.Fragement == "" {
		return errors.New("flag `fragment` not set. See --help for more information")
	}

	if c.Fragement != "major" && c.Fragement != "feature" && c.Fragement != "bug" &&
		c.Fragement != "alpha" && c.Fragement != "beta" && c.Fragement != "rc" {
		return errors.New("invalid fragment option. Possible options [major | feature | bug | alpha | beta | rc]")
	}

	return nil
}

func Execute(version string, args []string) error {
	app := cli.App{
		Name:            "bump-version",
		Usage:           "Bump a semantic version, following a given version fragment",
		UsageText:       "bump-version [options...] <version>",
		Version:         version,
		HideHelpCommand: true,
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

func run(ctx *cli.Context) error {
	c := Cli{
		Version:   ctx.Args().First(),
		Fragement: ctx.String("fragment"),
	}

	err := c.Validate()
	if err != nil {
		return err
	}

	current, err := semver.Parse(c.Version)
	if err != nil {
		return err
	}

	next := semver.Bump(current, c.Fragement)

	fmt.Printf("create %s-version: %v -> %v\n", c.Fragement, current.Print(), next.Print())

	return nil
}
