// Coordinated wrapper for:
// - cairo
// - pango
// - poppler

package graphics

import ()

const (
	A5_WIDTH  = 419.5276
	A5_HEIGHT = 595.2756
)

type Graphic struct {
	cairo *cairo
}

// Format is determined from filename extension
// Supported formats: pdf, png, ps
//
// Width and height are in pts for pdf, ps; pixels for png.
// Pixel measures will be truncated into integers
//
// Close the graphic to write the file
func NewGraphic(filename string, width float32, height float32) (*Graphic, error) {
	g := &Graphic{}

	var err error
	g.cairo, err = newCairo(filename, width, height)
	if err != nil {
		return nil, err
	}

	return g, nil
}

// completes the file being written to
func (g *Graphic) Close() error {
	err := g.cairo.Close()
	if err != nil {
		return err
	}

	return nil
}

// save the drawing context
func (g *Graphic) Save() {
	g.cairo.save()
}

// restore the drawing context
func (g *Graphic) Restore() {
	g.cairo.restore()
}
