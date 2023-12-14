package imagegenerator

type ImageGenerator interface {
	GenerateAsciiImages(imagePath string) (string, error)
}
