package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/h2non/bimg"
)

func main() {
	var input string
	var output string

	flag.StringVar(&input, "i", "logo.png", "Input image")
	flag.StringVar(&output, "o", "favicons", "Folder to output the favicons")
	flag.Usage = func() {
		fmt.Printf("Usage\n")
		fmt.Printf("  $ favicon-generator [options]\n")
		fmt.Printf("\n")
		fmt.Printf("Options\n")
		flag.PrintDefaults()
		fmt.Printf("\n")
		fmt.Printf("Examples\n")
		fmt.Printf("  $ favicon-generator\n")
		fmt.Printf("  $ favicon-generator -i path/to/logo.jpg\n")
		fmt.Printf("  $ favicon-generator -o /tmp/favicons\n")
		fmt.Printf("  $ favicon-generator -i path/to/logo.jpg -o /tmp/favicons\n")
	}

	flag.Parse()

	inputBuffer, err := bimg.Read(input)
	if err != nil {
		log.Fatal(err)
	}

	inputSize, err := bimg.NewImage(inputBuffer).Size()
	if err != nil {
		log.Fatal(err)
	}

	if inputSize.Width != inputSize.Height {
		log.Fatal("Input image is not a square")
	}

	err = os.MkdirAll(output, 0755)
	if err != nil {
		log.Fatal(err)
	}

	sizes := []int{32, 57, 76, 96, 120, 128, 144, 152, 180, 195, 196, 228, 270, 558}

	for _, size := range sizes {
		if size > inputSize.Width {
			continue
		}

		newIcon, err := bimg.NewImage(inputBuffer).Resize(size, size)
		if err != nil {
			log.Fatal(err)
		}

		iconPath := path.Join(output, fmt.Sprintf("favicon-%d.png", size))

		err = bimg.Write(iconPath, newIcon)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s created\n", iconPath)
	}
}
