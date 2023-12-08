package imagegenerator

import (
	"github.com/TheZoraiz/ascii-image-converter/aic_package"
)

func GetImageToPrint(filepath string) (string, error) {
	flags := aic_package.DefaultFlags()
	flags.Height = 30
	flags.Threshold = 95
	flags.Colored = true
	flags.Braille = true
	asciiArt, err := aic_package.Convert(filepath, flags)
	if err != nil {
		return "", err
	}

	return asciiArt, nil
}
