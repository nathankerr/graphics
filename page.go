package graphics

// #cgo pkg-config: cairo
// #include <cairo/cairo.h>
import "C"

func (g *Graphic) ShowPage() {
	C.cairo_show_page(g.cairo.cr)
}
