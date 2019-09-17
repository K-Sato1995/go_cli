package main

import (
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()

	app.Name = "sampleApp"
	app.Usage = "This app echo input arguments"
	app.Version = "0.0.1"

	app.Run(os.Args)
}
