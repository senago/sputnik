package media

import (
	"encoding/base64"
	"image"
	"strings"

	"github.com/senago/sputnik/internal/domain"
)

var (
	earthImage     image.Image
	kanopusImage   image.Image
	resourcePImage image.Image
	kondorImage    image.Image
	meteorMImage   image.Image
	electroLImage  image.Image
)

func init() {
	earthImage = mustDecodeBase64Image(rawEarthImage)
	kanopusImage = mustDecodeBase64Image(rawKanopusImage)
	resourcePImage = mustDecodeBase64Image(rawResourcePImage)
	kondorImage = mustDecodeBase64Image(rawKondorImage)
	meteorMImage = mustDecodeBase64Image(rawMeteorMImage)
	electroLImage = mustDecodeBase64Image(rawElectroLImage)
}

func GetEarthImage() image.Image {
	return earthImage
}

func GetSatelliteImage(satelliteType domain.SatelliteType) image.Image {
	switch satelliteType {
	case domain.SatelliteTypeResourceP:
		return resourcePImage
	case domain.SatelliteTypeKanopus:
		return kanopusImage
	case domain.SatelliteTypeKondor:
		return kondorImage
	case domain.SatelliteTypeMeteorM:
		return meteorMImage
	case domain.SatelliteTypeElectroL:
		return electroLImage
	default:
		return resourcePImage
	}
}

func mustDecodeBase64Image(s string) image.Image {
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(s))

	img, _, err := image.Decode(reader)
	if err != nil {
		panic(err)
	}

	return img
}
