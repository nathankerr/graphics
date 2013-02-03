package graphics

// #cgo pkg-config: cairo
// #include <cairo/cairo.h>
import "C"

type Point struct {
	x float32
	y float32
}

// Paths need to be manually destroyed
// type Path C.cairo_path_t
// func (p *Path) Destroy() {
// 	cr_path := p.(C.cairo_path_t)
// 	C.cairo_path_destroy(cr_path)
// }

// func (g *Graphic) CopyPath() error {
// 	path := C.cairo_copy_path(g.cairo.cr)

// 	status := C.cairo_status(g.cairo.cr)
// 	err := checkCairoStatus(status)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return path, nil
// }

// func (g *Graphic) CopyPathFlat() error {
// 	path := C.cairo_copy_path_flat(g.cairo.cr)

// 	status := C.cairo_status(g.cairo.cr)
// 	err := checkCairoStatus(status)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return path, nil
// }

// func (g *Graphic) AppendPath(p *Path) error {
// 	C.cairo_append_path(g.cairo.cr, p)

// 	status := C.cairo_status(g.cairo.cr)
// 	err := checkCairoStatus(status)
// 	if err != nil {
// 		return err
// 	}

// 	return nil	
// }

func (g *Graphic) Save() {
	C.cairo_save(g.cairo.cr)
}

func (g *Graphic) Restore() {
	C.cairo_restore(g.cairo.cr)
}

func (g *Graphic) NewPath() {
	C.cairo_new_path(g.cairo.cr)
}

func (g *Graphic) ClosePath() {
	C.cairo_close_path(g.cairo.cr)
}

func (g *Graphic) Arc(center Point, radius float32, angle1 float32, angle2 float32) {
	C.cairo_arc(
		g.cairo.cr,
		C.double(center.x),
		C.double(center.y),
		C.double(radius),
		C.double(angle1),
		C.double(angle2),
	)
}

func (g *Graphic) ArcNegative(center Point, radius float32, angle1 float32, angle2 float32) {
	C.cairo_arc_negative(
		g.cairo.cr,
		C.double(center.x),
		C.double(center.y),
		C.double(radius),
		C.double(angle1),
		C.double(angle2),
	)
}

func (g *Graphic) CurveTo(p1 Point, p2 Point, p3 Point) {
	C.cairo_curve_to(
		g.cairo.cr,
		C.double(p1.x),
		C.double(p1.y),
		C.double(p2.x),
		C.double(p2.y),
		C.double(p3.x),
		C.double(p3.y),
	)
}

func (g *Graphic) LineTo(p Point) {
	C.cairo_line_to(
		g.cairo.cr,
		C.double(p.x),
		C.double(p.y),
	)
}

func (g *Graphic) MoveTo(p Point) {
	C.cairo_move_to(
		g.cairo.cr,
		C.double(p.x),
		C.double(p.y),
	)
}

func (g *Graphic) Rectangle(topLeft Point, width float32, height float32) {
	C.cairo_rectangle(
		g.cairo.cr,
		C.double(topLeft.x),
		C.double(topLeft.y),
		C.double(width),
		C.double(height),
	)
}

func (g *Graphic) RelCurveTo(dp1 Point, dp2 Point, dp3 Point) {
	C.cairo_rel_curve_to(
		g.cairo.cr,
		C.double(dp1.x),
		C.double(dp1.y),
		C.double(dp2.x),
		C.double(dp2.y),
		C.double(dp3.x),
		C.double(dp3.y),
	)
}

func (g *Graphic) RelLineTo(dp Point) {
	C.cairo_rel_line_to(
		g.cairo.cr,
		C.double(dp.x),
		C.double(dp.y),
	)
}

func (g *Graphic) RelMoveTo(dp Point) {
	C.cairo_rel_move_to(
		g.cairo.cr,
		C.double(dp.x),
		C.double(dp.y),
	)
}
