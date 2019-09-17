package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()

	app.Name = "K-Sato's CLI"
	app.Usage = "Return inputs"
	app.Version = "0.0.1"

	app.Action = func(context *cli.Context) error {
		if context.Bool("y") {
			// return YO + input
			// go run -y main.go
			fmt.Println("YO " + context.Args().Get(0))
		} else {
			// just return the input
			fmt.Println(context.Args().Get(0))
		}
		return nil
	}

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "yo, y",
			Usage: "Echo with Yo",
		},
	}

	app.Run(os.Args)
}
