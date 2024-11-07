package logic

type Deck struct {
	Start TileID

	Count int
	Cards []Card
}

func StandardDeck() *Deck {
	cards := make([]Card, standart_count)

	index := 0
	for tile, amount := range standard_plain_cards {
		for i := 0; i < amount; i++ {
			cards[index].Tile = tile
			index++
		}
	}

	for tile, amount := range standard_emblem_cards {
		for i := 0; i < amount; i++ {
			cards[index].Tile = tile
			cards[index].Emblem = true
			index++
		}
	}

	return &Deck{
		Start: standard_start,
		Count: standart_count,
		Cards: cards,
	}
}

// O(n) is okay for 72 cards, O(n) is okay for 72 cards, O(n) is oka-
func (d *Deck) TakeStart() TileID

func (d *Deck) RandomCard() TileID

func (d *Deck) Take(tile TileID)

func (d *Deck) IsOver() bool
