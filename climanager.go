package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func (g *GhostManager) parseArgs() {

	app := &cli.App{
		Name:     "tgug",
		HelpName: "tgug",
		Usage:    "Typora-Ghost image Uploader in Go",
		Flags: []cli.Flag{
			&cli.StringSliceFlag{
				Name:     "file",
				Aliases:  []string{"f"},
				Usage:    "path to files, separated by commas \",\"",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			var respUrls []string
			for _, path := range c.StringSlice("file") {
				resp, _ := g.UploadImage(path)
				respUrls = append(respUrls, resp)
			}
			
			fmt.Println("🎉 Upload Successful!")
			for _, url := range respUrls{
				fmt.Println(url)
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
