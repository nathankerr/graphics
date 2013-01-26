package graphics

import (
	"testing"
)

func TestNewGraphic(t *testing.T) {
	tests := []string{} // format

	tests = append(tests, "pdf")

	for _, test := range tests {
		filename := "test." + test
		g, err := NewGraphic(filename, A5_WIDTH, A5_HEIGHT)
		if err != nil {
			t.Error(err)
		}
		g.Close()
	}
}
