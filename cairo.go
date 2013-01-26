package graphics

// #cgo pkg-config: cairo
// #include "cairo.h"
import "C"

import (
	"errors"
	"path/filepath"
)

type cairo struct {
	format   string
	filename string // needs to be kept for image surfaces
	surface  *C.cairo_surface_t
	cr       *C.cairo_t
}

func newCairo(filename string, width float32, height float32) (*cairo, error) {
	filename = filepath.Clean(filename)
	filename, err := filepath.Abs(filename)
	if err != nil {
		return nil, err
	}

	c := &cairo{
		format:   filepath.Ext(filename)[1:],
		filename: filename,
	}

	// create the surface, error checking is the same for all
	switch c.format {
	case "pdf":
		c.surface = C.cairo_pdf_surface_create(
			C.CString(filename),
			C.double(width),
			C.double(height),
		)
	case "png":
		c.surface = C.cairo_image_surface_create(
			C.CAIRO_FORMAT_ARGB32,
			C.int(width),
			C.int(height),
		)
	case "ps":
		c.surface = C.cairo_ps_surface_create(
			C.CString(filename),
			C.double(width),
			C.double(height),
		)
	case "svg":
		c.surface = C.cairo_svg_surface_create(
			C.CString(filename),
			C.double(width),
			C.double(height),
		)
	default:
		return nil, errors.New("cairo: unsupported format: " + c.format)
	}

	// error checking for all surface creations
	status := C.cairo_surface_status(c.surface)
	err = checkCairoStatus(status)
	if err != nil {
		return nil, err
	}

	// create the cairo context
	c.cr = C.cairo_create(c.surface)
	status = C.cairo_status(c.cr)
	err = checkCairoStatus(status)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *cairo) Close() error {
	// cr needs to be destroyed before the surface
	// and the status needs to be checked before that
	status := C.cairo_status(c.cr)
	err := checkCairoStatus(status)
	if err != nil {
		return err
	}
	C.cairo_destroy(c.cr)

	// write the surface to file
	switch c.format {
	case "pdf", "ps", "svg":
		// written when the surface is destroyed
	case "png":
		// TODO: use the go image libraries to handle
		// image output as cairo's png api is a "toy"
		status := C.cairo_surface_write_to_png(
			c.surface,
			C.CString(c.filename),
		)
		err := checkCairoStatus(status)
		if err != nil {
			return err
		}
	default:
		return errors.New("cairo: unsupported format: " + c.format)
	}

	C.cairo_surface_destroy(c.surface)
	status = C.cairo_surface_status(c.surface)
	err = checkCairoStatus(status)
	if err != nil {
		return err
	}

	return nil
}

func checkCairoStatus(status C.cairo_status_t) error {
	if status != C.CAIRO_STATUS_SUCCESS {
		return errors.New(C.GoString(C.cairo_status_to_string(status)))
	}
	return nil
}

func (c *cairo) save() {
	C.cairo_save(c.cr)
}

func (c *cairo) restore() {
	C.cairo_restore(c.cr)
}
