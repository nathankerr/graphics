package graphics

import (
	"fmt"
	"image/png"
	"os"
	"testing"
)

func TestNewGraphic(t *testing.T) {
	tests := []string{} // format

	tests = append(tests, "pdf")
	tests = append(tests, "png")
	tests = append(tests, "jpeg")
	tests = append(tests, "ps")
	tests = append(tests, "eps")
	tests = append(tests, "svg")

	for _, test := range tests {
		filename := "test." + test
		g, err := Create(filename, A5_WIDTH, A5_HEIGHT)
		if err != nil {
			t.Fatal(err)
		}

		img, err := g.Image()
		imageFilename := fmt.Sprintf("test.%s.png", test)
		imageFile, err := os.Create(imageFilename)
		if err != nil {
			t.Fatal(err)
		}
		err = png.Encode(imageFile, img)
		if err != nil {
			t.Fatal(err)
		}
		imageFile.Close()

		err = g.Close()
		if err != nil {
			t.Error(err)
		}
	}
}
