package graphics

import (
	"errors"
	"github.com/ungerik/go-cairo"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

type cairoWrapper struct {
	format   string
	filename string // needs to be kept for image surfaces
	surface  *cairo.Surface
}

func newCairo(filename string, width float64, height float64) (*cairoWrapper, error) {
	filename = filepath.Clean(filename)
	filename, err := filepath.Abs(filename)
	if err != nil {
		return nil, err
	}

	c := &cairoWrapper{
		format:   filepath.Ext(filename)[1:],
		filename: filename,
	}

	// create the surface, error checking is the same for all
	switch c.format {
	case "pdf":
		c.surface = cairo.NewPDFSurface(filename, width, height, cairo.PDF_VERSION_1_5)
	case "png", "jpeg":
		c.surface = cairo.NewSurface(cairo.FORMAT_ARGB32, int(width), int(height))
	case "ps":
		c.surface = cairo.NewPSSurface(filename, width, height, cairo.PS_LEVEL_3)
	case "svg":
		c.surface = cairo.NewSVGSurface(filename, width, height, cairo.SVG_VERSION_1_2)
	default:
		return nil, errors.New("cairo: unsupported format: " + c.format)
	}

	err = c.status()
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *cairoWrapper) Close() error {
	err := c.status()
	if err != nil {
		return err
	}

	c.surface.Finish()

	// write the surface to file
	switch c.format {
	case "pdf", "ps", "svg":
		// written when the surface is destroyed
	case "png":
		img := c.surface.GetImage()

		file, err := os.Create(c.filename)
		if err != nil {
			return err
		}
		defer file.Close()

		err = png.Encode(file, img)
		if err != nil {
			return err
		}
	case "jpeg":
		img := c.surface.GetImage()

		file, err := os.Create(c.filename)
		if err != nil {
			return err
		}
		defer file.Close()

		err = jpeg.Encode(file, img, nil)
		if err != nil {
			return err
		}
	default:
		return errors.New("cairo: unsupported format: " + c.format)
	}

	err = c.status()
	if err != nil {
		return err
	}

	return nil
}

func (c *cairoWrapper) status() error {
	status := c.surface.GetStatus()

	switch status {
	case cairo.STATUS_SUCCESS:
		return nil
	case cairo.STATUS_NO_MEMORY:
		return errors.New("Cairo: No Memory")
	case cairo.STATUS_INVALID_RESTORE:
		return errors.New("Cairo: Invalid Restore")
	case cairo.STATUS_INVALID_POP_GROUP:
		return errors.New("Cairo: Invalid Pop Group")
	case cairo.STATUS_NO_CURRENT_POINT:
		return errors.New("Cairo: No Current Point")
	case cairo.STATUS_INVALID_MATRIX:
		return errors.New("Cairo: Invalid Matrix")
	case cairo.STATUS_INVALID_STATUS:
		return errors.New("Cairo: Invalid Status")
	case cairo.STATUS_NULL_POINTER:
		return errors.New("Cairo: Null Pointer")
	case cairo.STATUS_INVALID_STRING:
		return errors.New("Cairo: Invalid String")
	case cairo.STATUS_INVALID_PATH_DATA:
		return errors.New("Cairo: Invalid Path Data")
	case cairo.STATUS_READ_ERROR:
		return errors.New("Cairo: Read Error")
	case cairo.STATUS_WRITE_ERROR:
		return errors.New("Cairo: Write Error")
	case cairo.STATUS_SURFACE_FINISHED:
		return errors.New("Cairo: Surface Finished")
	case cairo.STATUS_SURFACE_TYPE_MISMATCH:
		return errors.New("Cairo: Surface Type Mismatch")
	case cairo.STATUS_PATTERN_TYPE_MISMATCH:
		return errors.New("Cairo: Pattern Type Mismatch")
	case cairo.STATUS_INVALID_CONTENT:
		return errors.New("Cairo: Invalid Content")
	case cairo.STATUS_INVALID_FORMAT:
		return errors.New("Cairo: Invalid Format")
	case cairo.STATUS_INVALID_VISUAL:
		return errors.New("Cairo: Invalid Visual")
	case cairo.STATUS_FILE_NOT_FOUND:
		return errors.New("Cairo: File Not Found")
	case cairo.STATUS_INVALID_DASH:
		return errors.New("Cairo: Invalid Dash")
	case cairo.STATUS_INVALID_DSC_COMMENT:
		return errors.New("Cairo: Invalid DSC Comment")
	case cairo.STATUS_INVALID_INDEX:
		return errors.New("Cairo: Invalid Index")
	case cairo.STATUS_CLIP_NOT_REPRESENTABLE:
		return errors.New("Cairo: Clip Not Representable")
	case cairo.STATUS_TEMP_FILE_ERROR:
		return errors.New("Cairo: Temp File Error")
	case cairo.STATUS_INVALID_STRIDE:
		return errors.New("Cairo: Invalid Stride")
	case cairo.STATUS_FONT_TYPE_MISMATCH:
		return errors.New("Cairo: Font Type Mismatch")
	case cairo.STATUS_USER_FONT_IMMUTABLE:
		return errors.New("Cairo: User Font Immutable")
	case cairo.STATUS_USER_FONT_ERROR:
		return errors.New("Cairo: User Font Error")
	case cairo.STATUS_NEGATIVE_COUNT:
		return errors.New("Cairo: Negative Count")
	case cairo.STATUS_INVALID_CLUSTERS:
		return errors.New("Cairo: Invalid Clusters")
	case cairo.STATUS_INVALID_SLANT:
		return errors.New("Cairo: Invalid Slant")
	case cairo.STATUS_INVALID_WEIGHT:
		return errors.New("Cairo: Invalid Weight")
	case cairo.STATUS_INVALID_SIZE:
		return errors.New("Cairo: Invalid Size")
	case cairo.STATUS_USER_FONT_NOT_IMPLEMENTED:
		return errors.New("Cairo: User Font Not Implemented")
	case cairo.STATUS_DEVICE_TYPE_MISMATCH:
		return errors.New("Cairo: Device Type Mismatch")
	case cairo.STATUS_DEVICE_ERROR:
		return errors.New("Cairo: Device Error")
	}

	return errors.New("Unknown Cairo Error")
}
