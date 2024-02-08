package main

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/h2non/bimg"
	"github.com/urfave/cli/v2"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	app := &cli.App{
		Name:  "favicon-generator",
		Usage: "Command-line to generate favicons from a image.",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "input", Aliases: []string{"i"}, Required: true, Usage: "Input image"},
			&cli.StringFlag{Name: "output", Aliases: []string{"o"}, Value: "favicons", Usage: "Folder to output the favicons"},
			&cli.IntSliceFlag{Name: "size", Aliases: []string{"s"}, Value: cli.NewIntSlice(32, 57, 76, 96, 120, 128, 144, 152, 180, 195, 196, 228, 270, 558)},
		},
		Action: func(cCtx *cli.Context) error {
			return generateImages(cCtx.String("input"), cCtx.String("output"), cCtx.IntSlice("size"))
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func generateImages(input, output string, sizes []int) error {
	inputBuffer, err := bimg.Read(input)
	if err != nil {
		return err
	}

	inputSize, err := bimg.NewImage(inputBuffer).Size()
	if err != nil {
		return err
	}

	minSize := min(inputSize.Width, inputSize.Height)

	squareBuffer, err := bimg.NewImage(inputBuffer).Resize(minSize, minSize)
	if err != nil {
		return err
	}

	err = os.MkdirAll(output, 0755)
	if err != nil {
		return err
	}

	for _, size := range sizes {
		if size > inputSize.Width {
			continue
		}

		newIcon, err := bimg.NewImage(squareBuffer).Resize(size, size)
		if err != nil {
			return err
		}

		iconPath := path.Join(output, fmt.Sprintf("favicon-%d.png", size))

		err = bimg.Write(iconPath, newIcon)
		if err != nil {
			return err
		}

		fmt.Printf("%s created\n", iconPath)
	}
	return nil
}
