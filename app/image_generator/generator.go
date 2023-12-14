package imagegenerator

// import (
// 	_ "image/jpeg"
// 	_ "image/png"

// 	"github.com/qeesung/image2ascii/convert"
// )

// func GetImageToPrint(filepath string) (string, error) {
// 	convertOptions := convert.DefaultOptions
// 	convertOptions.Ratio = 0.1

// 	converter := convert.NewImageConverter()
// 	asciiArt := converter.ImageFile2ASCIIString(filepath, &convertOptions)
// 	return asciiArt, nil
// }

import (
	"github.com/TheZoraiz/ascii-image-converter/aic_package"
)

func GetImageToPrint(filepath string) (string, error) {
	flags := aic_package.DefaultFlags()
	flags.Width = 60
	flags.Threshold = 80
	flags.Colored = true
	flags.Braille = true
	flags.Dither = true
	asciiArt, err := aic_package.Convert(filepath, flags)
	if err != nil {
		return "", err
	}

	return asciiArt, nil
}
