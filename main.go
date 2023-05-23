package main

import (
	"log"
	"os"

	"github.com/appinventories/generator-packages/internal/generator"
	"github.com/urfave/cli/v2"
)

func main() {
	cfgFileFlag := &cli.StringFlag{
		Name:     "config-file",
		Usage:    "Path to .yaml config file. If config file is used, the program will ignore other parameters.",
		Aliases:  []string{"c"},
		Required: true,
	}

	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "GenLocaleKeys",
				Usage: "Generate locale keys",
				Flags: []cli.Flag{
					cfgFileFlag,
				},
				Action: func(c *cli.Context) error {
					cfgFile := c.String("config-file")
					return generator.GenerateLocaleKey(cfgFile)
				},
			},
			{
				Name:  "CsvToJson",
				Usage: "csv to json",
				Flags: []cli.Flag{
					cfgFileFlag,
				},
				Action: func(c *cli.Context) error {
					cfgFile := c.String("config-file")
					return generator.CsvToJson(cfgFile)
				},
			},
			{
				Name:  "GenAppPath",
				Usage: "Generate app path",
				Flags: []cli.Flag{
					cfgFileFlag,
				},
				Action: func(c *cli.Context) error {
					cfgFile := c.String("config-file")

					return generator.AppPath(cfgFile)
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
