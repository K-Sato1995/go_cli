package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()

	app.Name = "sampleApp"
	app.Usage = "This app echo input arguments"
	app.Version = "0.0.1"

	app.Action = func(context *cli.Context) error {
		if context.Bool("cat") {
			fmt.Println(context.Args().Get(0) + "だにゃん♡")
		} else {
			fmt.Println(context.Args().Get(0))
		}
		return nil
	}

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "cat, c",
			Usage: "Echo with cat",
		},
	}

	app.Run(os.Args)
}
