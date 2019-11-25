package main

import (
	"common"
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
	"server"
)

func main() {
	app := cli.NewApp()
	app.Name = "tran"
	app.Version = common.Version
	app.Description = "简繁转换"
	app.Usage = ""

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config,c",
			Usage: "load configuration from `FILE`",
		},
	}

	app.Action = func(c *cli.Context) error {
		configFile := c.String("config")
		if len(configFile) == 0 {
			cli.ShowAppHelp(c)
			return nil
		}

		log.Printf("run with config file %s", configFile)

		if _, err := common.ParseXmlConfig(configFile); err != nil {
			return fmt.Errorf("not found config file %v", configFile)
		}

		return server.Run()
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Print(err)
	}
}