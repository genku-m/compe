package main

import (
	"fmt"
	"os"

	"github.com/genku-m/competition_file_generator/generator"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "competition_tools"
	app.Usage = "for competition tools"
	app.Version = "0.0.1"

	app.Commands = []cli.Command{
		{
			Name:  "generate, g",
			Usage: "generate [title]",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "name, n",
					Value: "atcoder",
					Usage: "generate file for atcoder",
				},
				cli.StringFlag{
					Name:  "type, t",
					Value: "abc",
					Usage: "generate file for atcoder abc",
				},
			},
			Action: func(c *cli.Context) {
				competition := c.String("name")
				ctype := c.String("type")
				title := c.Args().First()
				if title == "" {
					fmt.Println("title is empty")
					return
				}
				g, err := generator.NewGenerator(competition, ctype, title)
				if err != nil {
					fmt.Println(err.Error())
					return
				}

				if err = g.Generate(); err != nil {
					fmt.Println(err.Error())
					return
				}

				fmt.Printf("generate file for %s %s %s", competition, ctype, title)
				return
			},
		},
	}
	app.Run(os.Args)
}
