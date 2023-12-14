package imagegenerator

import (
	"github.com/TheZoraiz/ascii-image-converter/aic_package"
)

type defaultGenerator struct {
	defaultFlags aic_package.Flags
}

func NewDefaultGenerator() *defaultGenerator {
	flags := aic_package.DefaultFlags()
	flags.Width = 60
	flags.Threshold = 80
	flags.Colored = true
	flags.Braille = true
	flags.Dither = true

	generator := &defaultGenerator{
		defaultFlags: flags,
	}

	return generator
}

func (d *defaultGenerator) GenerateAsciiImages(imagePath string) (string, error) {
	asciiArt, err := aic_package.Convert(imagePath, d.defaultFlags)
	if err != nil {
		return "", err
	}

	return asciiArt, nil
}
