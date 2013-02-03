package graphics

// #cgo pkg-config: cairo
// #include <cairo/cairo-pdf.h>
// #include <cairo/cairo-ps.h>
// #include <cairo/cairo-svg.h>
import "C"

import (
	"image/jpeg"
	"os"
)

func (g *Graphic) PlaceImage(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	image, err := jpeg.Decode(file)
	if err != nil {
		return err
	}

	// The damn pixels are in the wrong order!
	// cairo expects ARGB, but go uses RGBA!
	bounds := image.Bounds()
	width := bounds.Max.X - bounds.Min.X
	height := bounds.Max.Y - bounds.Min.Y
	buffer := make([]uint8, 0, 4*width*height)
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			r, g, b, a := image.At(x, y).RGBA()
			buffer = append(buffer, uint8(a))
			buffer = append(buffer, uint8(r))
			buffer = append(buffer, uint8(g))
			buffer = append(buffer, uint8(b))
		}
	}

	// imageSurface := C.cairo_image_surface_create_for_data(
	// 	C.unsigned_char(&buffer),
	// 	C.CAIRO_FORMAT_ARGB32,
	// 	width,
	// 	height,
	// 	32*width,
	// )

	return nil
}
