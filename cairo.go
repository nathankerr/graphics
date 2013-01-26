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

	switch g.format {
	case "pdf":
		g.cairo.surface = C.cairo_pdf_surface_create(
			C.CString(g.filename),
			C.double(g.width),
			C.double(g.height),
		)
		err := g.cairoStatus()
		if err != nil {
			return err
		}
	default:
		return errors.New("cairo: unsupported format: " + g.format)
	}

	return nil
}

func (g *Graphic) cairoClose() error {
	C.cairo_surface_destroy(g.cairo.surface)
	err := g.cairoStatus()
	if err != nil {
		return err
	}

	return nil
}

func (g *Graphic) cairoStatus() error {
	status := C.cairo_surface_status(g.cairo.surface)
	if status != C.CAIRO_STATUS_SUCCESS {
		return errors.New(C.GoString(C.cairo_status_to_string(status)))
	}
	return nil
}
