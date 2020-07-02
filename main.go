package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := &cli.App{
		Name:  "f2i",
		Usage: "Encode/decode data to BMP images",
		Commands: []cli.Command{
			{
				Name:        "encode",
				Usage:       "Encode specified file into BMP images",
				UsageText:   "encode [file] [output_dir]",
				Description: "Encode specified file into BMP images",
				Action:      CmdEncode,
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:  "width",
						Usage: "",
						Value: 1000,
					},
					&cli.IntFlag{
						Name:  "height",
						Usage: "",
						Value: 1000,
					},
				},
			},
			{
				Name:        "decode",
				Usage:       "Decode BMP images in directory into a file",
				UsageText:   "decode [images_dir] [output_file]",
				Description: "Decode BMP images in directory into a file",
				Action:      CmdDecode,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func CmdEncode(ctx *cli.Context) error {
	args := ctx.Args()
	if len(args) != 2 {
		return fmt.Errorf(ctx.Command.UsageText)
	}
	return Encode(args[0], args[1], ctx.Int("width"), ctx.Int("height"))
}

func CmdDecode(ctx *cli.Context) error {
	args := ctx.Args()
	if len(args) != 2 {
		return fmt.Errorf(ctx.Command.UsageText)
	}
	return Decode(args[1], args[0])
}
