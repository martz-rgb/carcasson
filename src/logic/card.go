package logic

// counter-clockwise
type Rotation uint8

type Card struct {
	Tile     TileID
	Rotation Rotation

	Emblem bool // -> modificators
}
