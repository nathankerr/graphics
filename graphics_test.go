package graphics

import (
	"testing"
)

func TestNewGraphicPdf(t *testing.T) {
	g, err := NewGraphic("test.pdf", A5_WIDTH, A5_HEIGHT)
	if err != nil {
		t.Error(err)
	}
	g.Close()
}
