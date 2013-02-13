// Coordinated wrapper for cairo based graphics handling
package graphics

// #cgo pkg-config: cairo
// #include <cairo-pdf.h>
// #include <cairo-ps.h>
// #include <cairo-svg.h>
import "C"

import (
	"errors"
	"fmt"
	"github.com/ungerik/go-cairo/extimage"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"unsafe"
)

const (
	A5_WIDTH  = 419.5276
	A5_HEIGHT = 595.2756
)

type Graphic struct {
	filename string
	format   string
	surface  *C.cairo_surface_t
	cr       *C.cairo_t
}

// Format is determined from filename extension
// Supported formats: jpeg, pdf, png, ps, svg
//
// Width and height are in pts for pdf, ps, and svg; pixels for jpeg and png.
// Pixel measures will be truncated into integers
//
// Close the graphic to write the file
func Create(filename string, width float64, height float64) (*Graphic, error) {
	g := &Graphic{}

	var err error
	filename = filepath.Clean(filename)
	g.filename, err = filepath.Abs(filename)
	if err != nil {
		return nil, err
	}

	g.format = filepath.Ext(filename)[1:]

	// create the surface
	switch g.format {
	case "pdf":
		g.surface = C.cairo_pdf_surface_create(
			C.CString(filename),
			C.double(width),
			C.double(height),
		)
	case "png", "jpeg":
		g.surface = C.cairo_image_surface_create(
			C.CAIRO_FORMAT_ARGB32,
			C.int(width),
			C.int(height),
		)
	case "ps":
		g.surface = C.cairo_ps_surface_create(
			C.CString(filename),
			C.double(width),
			C.double(height),
		)
	case "svg":
		g.surface = C.cairo_svg_surface_create(
			C.CString(filename),
			C.double(width),
			C.double(height),
		)
	default:
		return nil, errors.New("cairo: unsupported format: " + g.format)
	}
	err = g.cairoSurfaceStatus()
	if err != nil {
		return nil, err
	}

	// create the cairo context
	g.cr = C.cairo_create(g.surface)
	err = g.cairoStatus()
	if err != nil {
		return nil, err
	}

	return g, nil
}

// completes and closes the file being written to
func (g *Graphic) Close() error {
	// destroy the context
	err := g.cairoStatus()
	if err != nil {
		return err
	}
	C.cairo_destroy(g.cr)

	// finishing the surface writes pdf, ps, and svg files
	C.cairo_surface_finish(g.surface)
	err = g.cairoSurfaceStatus()
	if err != nil {
		return err
	}

	// write other formats
	switch g.format {
	case "pdf", "ps", "svg":
		// cairo_surface_finish writes the surface to file
	case "png":
		img, err := g.Image()
		if err != nil {
			return err
		}

		file, err := os.Create(g.filename)
		if err != nil {
			return err
		}
		defer file.Close()

		err = png.Encode(file, img)
		if err != nil {
			return err
		}
	case "jpeg":
		img, err := g.Image()
		if err != nil {
			return err
		}

		file, err := os.Create(g.filename)
		if err != nil {
			return err
		}
		defer file.Close()

		err = jpeg.Encode(file, img, nil)
		if err != nil {
			return err
		}
	default:
		panic("unsupported format: " + g.format)
	}

	// destroy the surface
	C.cairo_surface_destroy(g.surface)
	err = g.cairoSurfaceStatus()
	if err != nil {
		return err
	}

	return nil
}

// pdf, ps, and svg surfaces use the current surface (page)
func (g *Graphic) Image() (image.Image, error) {
	var surface *C.cairo_surface_t
	switch g.format {
	case "pdf", "ps", "svg":
		// map vector surfaces to an image surface
		surface = C.cairo_surface_map_to_image(g.surface, nil)
		defer C.cairo_surface_unmap_image(g.surface, surface)

		status := C.cairo_surface_status(surface)
		err := statusToError(status)
		if err != nil {
			return nil, err
		}
	case "png", "jpeg":
		// no conversion needed
		surface = g.surface
	}

	width := int(C.cairo_image_surface_get_width(surface))
	height := int(C.cairo_image_surface_get_height(surface))
	stride := int(C.cairo_image_surface_get_stride(surface))
	format := C.cairo_image_surface_get_format(surface)
	dataPtr := C.cairo_image_surface_get_data(surface)
	data := C.GoBytes(unsafe.Pointer(dataPtr), C.int(stride*height))

	var img image.Image
	switch format {
	case C.CAIRO_FORMAT_ARGB32:
		img = &extimage.ARGB{
			Pix:    data,
			Stride: stride,
			Rect:   image.Rect(0, 0, width, height),
		}
	case C.CAIRO_FORMAT_RGB24:
		img = &extimage.RGB{
			Pix:    data,
			Stride: stride,
			Rect:   image.Rect(0, 0, width, height),
		}
	case C.CAIRO_FORMAT_A8:
		img = &image.Alpha{
			Pix:    data,
			Stride: stride,
			Rect:   image.Rect(0, 0, width, height),
		}
	default:
		// known unsupported formats:
		// CAIRO_FORMAT_INVALID   = -1,
		// CAIRO_FORMAT_A1        = 3,
		// CAIRO_FORMAT_RGB16_565 = 4,
		// CAIRO_FORMAT_RGB30     = 5
		panic(fmt.Sprintf("unsupported cairo image surface format: %d", int(format)))
	}

	return img, nil
}
