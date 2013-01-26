// Coordinated wrapper for:
// - cairo
// - pango
// - poppler

package graphics

import (
	"errors"
	"path/filepath"
)

const (
	A5_WIDTH  = 419.5276
	A5_HEIGHT = 595.2756
)

type Graphic struct {
	filename string
	format   string
	width    float32
	height   float32
	cairo    cairo
}

// format is automatically determined from filename extension
// supported formats: pdf
// width and height are in pts for pdf; pixels for ...
func NewGraphic(filename string, width float32, height float32) (*Graphic, error) {
	filename = filepath.Clean(filename)
	filename, err := filepath.Abs(filename)
	if err != nil {
		return nil, err
	}

	g := &Graphic{
		filename: filename,
		format:   filepath.Ext(filename)[1:],
		width:    width,
		height:   height,
	}

	switch g.format {
	case "pdf":
		// supported format types
	default:
		return nil, errors.New("unsupported format: " + g.format)
	}

	err = g.cairoInit()
	if err != nil {
		return nil, err
	}

	return g, nil
}

func (g *Graphic) Close() error {
	err := g.cairoClose()
	if err != nil {
		return err
	}

	return nil
}
