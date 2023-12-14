package imagegenerator

import (
	_ "image/jpeg"
	_ "image/png"

	"github.com/qeesung/image2ascii/convert"
)

type anotherGenerator struct {
	defaultOptions *convert.Options
	converter      *convert.ImageConverter
}

func NewAnotherGenerator() *anotherGenerator {
	convertOptions := &convert.DefaultOptions
	convertOptions.Ratio = 0.1

	generator := &anotherGenerator{
		defaultOptions: convertOptions,
		converter:      convert.NewImageConverter(),
	}
	return generator
}

func (g *anotherGenerator) GenerateAsciiImages(imagePath string) (string, error) {
	asciiArt := g.converter.ImageFile2ASCIIString(imagePath, g.defaultOptions)
	return asciiArt, nil
}
