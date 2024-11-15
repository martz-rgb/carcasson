package logic

type Position struct {
	X, Y int
}

// Zero-valued Field is ready to use
type Field struct {
	DownLeft, RightUp Position
	Cards             map[Position]Card

	Meeples map[Position]Meeple
}

func (f *Field) AddTile(pos Position, card Card)

func (f *Field) AddTileWithMeeple(pos Position, card Card, meeple Meeple)

func (f *Field) addTile()

func (f *Field) PossibleSlots(tile TileID) map[Position][]Rotation

func (f *Field) IsCorrect(pos Position, card Card) bool

func (f *Field) Scores()
