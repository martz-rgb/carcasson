package logic

type Position struct {
	X, Y int
}

// Zero-valued Field is ready to use
type Field struct {
	DownLeft, RightUp Position
	Cards             map[Position]Card

	Meeple map[Position]Meeple
}

func (f *Field) AddTile(pos Position, card Card)

func (f *Field) AddMeeple(pos Position, meeple Meeple)

func (f *Field) IsPossible(tile TileID) bool

func (f *Field) IsCorrect(pos Position, card Card) bool

func (f *Field) Score()
