package main

import (
	"fmt"
	"log"
	"os"
	"image"
	_ "image/png"
	_ "image/jpeg"
)
const length uint32 = 10
const usage string = "Usage: ./ascii <image path> <dir path>"

// verifies and retrieves the image at the specified path
func getSource() *os.File {
	path := os.Args[1]

	if path == "" {
		log.Fatal("No image specified!\n" + usage)
	}

	img, err := os.Open(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}
	return img
}

// verifies and retrieves the directory path specified
func getDestination() string {
	path := os.Args[2]

	if path == "" {
		log.Fatal("No directory path specified!\n" + usage)
	}

	_, err := os.Open(os.Args[2])

	if err != nil {
		log.Fatal(err)
	}
	
	return os.Args[2]
}

func main() {
	// ascii colour mapping, from darkest to brightest
	opacity := [length + 1]rune{' ', '.', ',', '+', '*', ':', 'o', '&', '8', '#', '@'}

	// make sure we actually have the image & save directory path
	if len(os.Args) < 3 {
		log.Fatal("Not enough arguments specified!\n" + usage)
	}

	file := getSource()
	path := getDestination() + "/output.txt"

	img, _, err := image.Decode(file)

	if err != nil {
		log.Fatal(err)
	}
	
	// determine image dimensions
	bounds := img.Bounds()
	var text string

	// iterate pixels from top corner to bottom corner of image
	for y, yl := bounds.Min.Y, bounds.Max.Y; y <= yl; y++ {
		for x, xl := bounds.Min.X, bounds.Max.X; x <= xl; x++ {
			r,g,b,a := img.At(x, y).RGBA()
			// get the average of RGB values to find monochromatic equilivent
			greyscale := (r + g + b)/3
			// use pixel opacity as a multiplier for resulting colour
			ratio := (greyscale * (a/65535))
			// multiply the ratio by ASCII map length to get required char
			shade := (length * ratio) / 65535
			// convert ASCII decimal value to string
			char := string(opacity[shade])
			// double character to scale image correctly
			text += char + char
		}
		text += "\n" // new Y value, new line
	}

	output, err := os.Create(path)

	// file couldn't be created
	if err != nil {
		log.Fatal(err)
	}

	output.WriteString(text)
	fmt.Println("Successfully wrote ASCII image to", path)
}
