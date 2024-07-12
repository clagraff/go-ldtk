package goldtk

import (
	"fmt"
	"goldtk/quicktype"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io/fs"
	"log"
)

// Tileset interface defines the methods for working with a tileset, including retrieving tiles and tileset properties.
type Tileset interface {
	Identifier() Identifier
	Uid() Uid

	Image() image.Image
	Tile(tileId int) image.Image

	GridWidth() int
	GridHeight() int

	TileGridSize() int // How does this differ from GridSizeInpx
	Padding() int
	Spacing() int

	Width() int
	Height() int
	Tags() []string
}

// tileset struct implements the Tileset interface and holds the tileset definition and image source.
type tileset struct {
	def    quicktype.TilesetDefinition
	source image.Image
}

// Tile returns the subimage of the specified tileId from the tileset image.
// It calculates the tile's position considering padding and spacing and extracts the subimage.
func (t tileset) Tile(tileId int) image.Image {
	gridWidth := t.GridWidth()
	gridSize := t.TileGridSize()
	padding := t.Padding()
	spacing := t.Spacing()

	// Calculate grid-based coordinates of the tileId
	gridTileX := tileId % gridWidth
	gridTileY := tileId / gridWidth

	// Calculate the atlas pixel coordinates
	px := padding + gridTileX*(gridSize+spacing)
	py := padding + gridTileY*(gridSize+spacing)

	// Calculate the subimage rectangle
	subImageRect := image.Rect(px, py, px+gridSize, py+gridSize)

	// Ensure the subImageRect is within the bounds of the source image
	if !subImageRect.In(t.source.Bounds()) {
		log.Fatalf("subImageRect %v is out of bounds of source image %v", subImageRect, t.source.Bounds())
	}

	// Extract the subimage
	subImage := t.Image().(interface {
		SubImage(r image.Rectangle) image.Image
	}).SubImage(subImageRect)

	return subImage
}

// Identifier returns the identifier of the tileset.
func (t tileset) Identifier() Identifier {
	return Identifier(t.def.Identifier)
}

// Uid returns the unique ID of the tileset.
func (t tileset) Uid() Uid {
	return Uid(t.def.Uid)
}

// Image returns the source image of the tileset.
func (t tileset) Image() image.Image {
	return t.source
}

// GridWidth returns the number of tiles per row in the tileset grid.
func (t tileset) GridWidth() int {
	return int(t.def.CWid)
}

// GridHeight returns the number of tiles per column in the tileset grid.
func (t tileset) GridHeight() int {
	return int(t.def.CHei)
}

// TileGridSize returns the size of each tile in the grid.
func (t tileset) TileGridSize() int {
	return int(t.def.TileGridSize)
}

// Padding returns the padding around the tiles in the tileset image.
func (t tileset) Padding() int {
	return int(t.def.Padding)
}

// Spacing returns the spacing between the tiles in the tileset image.
func (t tileset) Spacing() int {
	return int(t.def.Spacing)
}

// Width returns the width of the tileset image in pixels.
func (t tileset) Width() int {
	return int(t.def.PxWid)
}

// Height returns the height of the tileset image in pixels.
func (t tileset) Height() int {
	return int(t.def.PxHei)
}

// Tags returns the tags associated with the tileset.
func (t tileset) Tags() []string {
	return t.def.Tags
}

// NewTileset creates a new Tileset instance from a tileset definition and file system.
func NewTileset(def quicktype.TilesetDefinition, sys fs.FS) (Tileset, error) {
	if def.RelPath == nil {
		return nil, fmt.Errorf("tileset definition requires relative path")
	}

	// Open the tileset image file
	f, err := sys.Open(*def.RelPath)
	if err != nil {
		return nil, fmt.Errorf("opening tileset %s: %w", *def.RelPath, err)
	}
	defer func(f fs.File) {
		err := f.Close()
		if err != nil {
			log.Fatalf("failed to close tileset file: %v", err)
		}
	}(f)

	// Decode the image
	src, _, err := image.Decode(f)
	if err != nil {
		return nil, fmt.Errorf("decoding tileset image %s: %w", *def.RelPath, err)
	}

	return tileset{
		def:    def,
		source: src,
	}, nil
}

// Ensure that tileset implements the Tileset interface.
var _ Tileset = tileset{}

// Tile interface defines methods for working with individual tiles, including their properties and image representation.
type Tile interface {
	Opacity() float64

	X() int
	Y() int

	FlipY() bool
	FlipX() bool
	FlipBoth() bool

	Image() image.Image
}

// tile struct implements the Tile interface and holds the tile instance and the associated tileset.
type tile struct {
	inst quicktype.TileInstance
	ts   Tileset
}

// Opacity returns the opacity of the tile.
func (t tile) Opacity() float64 {
	return t.inst.A
}

// X returns the X coordinate of the tile.
func (t tile) X() int {
	return int(t.inst.Px[0])
}

// Y returns the Y coordinate of the tile.
func (t tile) Y() int {
	return int(t.inst.Px[1])
}

// FlipX returns true if the tile is flipped along the X axis.
func (t tile) FlipX() bool {
	return t.inst.F&1 != 0
}

// FlipY returns true if the tile is flipped along the Y axis.
func (t tile) FlipY() bool {
	return t.inst.F&2 != 0
}

// FlipBoth returns true if the tile is flipped along both the X and Y axes.
func (t tile) FlipBoth() bool {
	return t.inst.F&3 == 3
}

// Image returns the image of the tile, extracted from the associated tileset.
func (t tile) Image() image.Image {
	return t.ts.Tile(int(t.inst.T))
}

// NewTile creates a new Tile instance from a tile instance, associated tileset, and image.
func NewTile(inst quicktype.TileInstance, ts Tileset) Tile {
	return tile{
		inst: inst,
		ts:   ts,
	}
}

// Ensure that tile implements the Tile interface.
var _ Tile = tile{}
