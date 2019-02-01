package main

import (
	"os"

	"github.com/urfave/cli"
	"github.com/vikebot/vbmapg/cmd/gen"
	"go.uber.org/zap"
)

func main() {
	log, err := zap.NewDevelopment()
	if err != nil {
		println("main: Unable to create a logger: " + err.Error())
		os.Exit(-1)
	}
	defer log.Sync()

	app := cli.NewApp()
	app.Version = "1.0.0"
	app.Copyright = "(c) 2018 vikebot"

	app.Commands = []cli.Command{
		{
			Name:    "gen",
			Aliases: []string{"g"},
			Usage:   "generates a map",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "height, mh",
					Value: 0,
					Usage: "map height",
				},
				cli.IntFlag{
					Name:  "width, mw",
					Value: 0,
					Usage: "map width",
				},
			},
			Action: func(c *cli.Context) {
				log.Info("Create map")
				gen.Create(c.Int("height"), c.Int("width"), log.Named("Gen"))
			},
		},
		{
			Name:    "approve",
			Aliases: []string{"a"},
			Usage:   "approves a map",
			Action: func(c *cli.Context) error {
				log.Info("Approve map")
				return nil
			},
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal("main: Unable to start cli", zap.Error(err))
	}

	// if len(os.Args) != 2 {
	// 	log.Error("To few arguments. Usage: mapgen <width> <height>")
	// 	return
	// }

	// width, err := strconv.Atoi(os.Args[1])
	// if err != nil {
	// 	println(err.Error())
	// 	os.Exit(-1)
	// }
	// height, err := strconv.Atoi(os.Args[2])
	// if err != nil {
	// 	println(err.Error())
	// 	os.Exit(-1)
	// }

}
