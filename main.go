package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/genku-m/compe/generator"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "competition_tools"
	app.Usage = "for competition tools"
	app.Version = "0.0.1"

	app.Commands = []*cli.Command{
		{
			Name:    "generate",
			Aliases: []string{"g"},
			Usage:   "generate [title]",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "name",
					Aliases: []string{"n"},
					Value:   "atcoder",
					Usage:   "generate file for atcoder",
				},
				&cli.StringFlag{
					Name:    "type",
					Aliases: []string{"t"},
					Value:   "abc",
					Usage:   "generate file for atcoder abc",
				},
			},
			Action: func(c *cli.Context) error {
				competition := c.String("name")
				ctype := c.String("type")
				title := c.Args().First()
				if title == "" {
					return errors.New("title is empty")
				}
				g, err := generator.NewGenerator(competition, ctype, title)
				if err != nil {
					return err
				}

				if err = g.Generate(); err != nil {
					return err
				}

				fmt.Printf("generate file for %s %s %s", competition, ctype, title)
				return nil
			},
		},
	}
	app.Run(os.Args)
}
