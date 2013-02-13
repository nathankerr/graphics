// Coordinated wrapper for:
// - cairo
// - pango
// - poppler

package graphics

const (
	A5_WIDTH  = 419.5276
	A5_HEIGHT = 595.2756
)

type Graphic struct {
	cairo *cairoWrapper
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
	g.cairo, err = newCairo(filename, width, height)
	if err != nil {
		return nil, err
	}

	return g, nil
}

// completes and closes the file being written to
func (g *Graphic) Close() error {
	err := g.cairo.Close()
	if err != nil {
		return err
	}

	return nil
}
