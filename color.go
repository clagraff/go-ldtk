package goldtk

import (
	"fmt"
	"image/color"
	"log"
	"strconv"
	"strings"
)

// Color provides an interface for working with colors.
type Color interface {
	// Hex returns the hex code of the current color. It will be
	// prefixed with a octothorpe.
	// The alpha value of the underlying color is completely ignored when creating
	// the hexcode.
	Hex() string
	RGBA() color.RGBA
}

func ColorFromInt64(value int64) Color {
	return clr{
		value: color.RGBA{
			R: uint8(value >> 16),
			G: uint8(value >> 8),
			B: uint8(value),
		},
	}
}

func ColorFromColor(c color.Color) Color {
	return clr{
		value: c,
	}
}

func ColorFromHex(hex string) Color {
	// Remove the '#' if it exists
	hex = strings.TrimPrefix(hex, "#")

	// Ensure the hex string is 6 characters long
	if len(hex) != 6 {
		log.Println("invalid hex color format:", hex)
		return clr{value: color.Black}
	}

	// Parse the hex string
	r, err := strconv.ParseUint(hex[0:2], 16, 8)
	if err != nil {
		log.Println("invalid hex color format:", hex)
		return clr{value: color.Black}
	}
	g, err := strconv.ParseUint(hex[2:4], 16, 8)
	if err != nil {
		log.Println("invalid hex color format:", hex)
		return clr{value: color.Black}
	}
	b, err := strconv.ParseUint(hex[4:6], 16, 8)
	if err != nil {
		log.Println("invalid hex color format:", hex)
		return clr{value: color.Black}
	}

	return clr{
		value: color.RGBA{
			R: uint8(r),
			G: uint8(g),
			B: uint8(b),
			A: 255, // assuming full opacity since the hex code does not include alpha
		},
	}
}

type clr struct {
	value color.Color
}

func (c clr) Hex() string {
	r, g, b, _ := c.value.RGBA()
	return fmt.Sprintf("#%02x%02x%02x", r, g, b)
}

func (c clr) RGBA() color.RGBA {
	r, g, b, a := c.value.RGBA()

	return color.RGBA{
		R: uint8(r),
		G: uint8(g),
		B: uint8(b),
		A: uint8(a),
	}
}

var _ Color = clr{}
