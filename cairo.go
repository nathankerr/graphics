package graphics

// #cgo pkg-config: pangocairo
// #include "cairo.h"
import "C"

import (
	"errors"
)

type cairo struct {
	surface *C.cairo_surface_t
	cr      *C.cairo_t
}

func (g *Graphic) cairoInit() error {
	g.cairo = cairo{}

	// create the surface, error checking is the same for all
	switch g.format {
	case "pdf":
		g.cairo.surface = C.cairo_pdf_surface_create(
			C.CString(g.filename),
			C.double(g.width),
			C.double(g.height),
		)
	case "png":
		g.cairo.surface = C.cairo_image_surface_create(
			C.CAIRO_FORMAT_ARGB32,
			C.int(g.width),
			C.int(g.height),
		)
	case "ps":
		g.cairo.surface = C.cairo_ps_surface_create(
			C.CString(g.filename),
			C.double(g.width),
			C.double(g.height),
		)
	case "svg":
		g.cairo.surface = C.cairo_svg_surface_create(
			C.CString(g.filename),
			C.double(g.width),
			C.double(g.height),
		)
	default:
		return errors.New("cairo: unsupported format: " + g.format)
	}

	// error checking for all surface creations
	status := C.cairo_surface_status(g.cairo.surface)
	err := checkCairoStatus(status)
	if err != nil {
		return err
	}

	// create the cairo context
	g.cairo.cr = C.cairo_create(g.cairo.surface)
	status = C.cairo_status(g.cairo.cr)
	err = checkCairoStatus(status)
	if err != nil {
		return err
	}

	return nil
}

func (g *Graphic) cairoClose() error {
	// cr needs to be destroyed before the surface
	// and the status needs to be checked before that
	status := C.cairo_status(g.cairo.cr)
	err := checkCairoStatus(status)
	if err != nil {
		return err
	}
	C.cairo_destroy(g.cairo.cr)

	// write the surface to file
	switch g.format {
	case "pdf", "ps", "svg":
		// written when the surface is destroyed
	case "png":
		// TODO: use the go image libraries to handle
		// image output as cairo's png api is a "toy"
		status := C.cairo_surface_write_to_png(
			g.cairo.surface,
			C.CString(g.filename),
		)
		err := checkCairoStatus(status)
		if err != nil {
			return err
		}
	default:
		return errors.New("cairo: unsupported format: " + g.format)
	}

	C.cairo_surface_destroy(g.cairo.surface)
	status = C.cairo_surface_status(g.cairo.surface)
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
