package deck

import (
	"testing"
)

type data struct {
	want Card
	get  string
}

func TestCard(t *testing.T) {
	var tests []data

	tests = append(tests, data{want: Card{Rank: Ace, Suit: Heart}, get: "Ace of Hearts"})
	tests = append(tests, data{want: Card{Rank: Jack, Suit: Club}, get: "Jack of Clubs"})
	tests = append(tests, data{want: Card{Rank: Seven, Suit: Diamond}, get: "Seven of Diamonds"})
	tests = append(tests, data{want: Card{Rank: King, Suit: Spade}, get: "King of Spades"})
	tests = append(tests, data{want: Card{Suit: Joker}, get: "Joker"})

	for _, check := range tests {
		if check.want.String() != check.get {
			t.Errorf("want: %s\t get: %s", check.want.String(), check.get)
		}
	}
}

func TestNew(t *testing.T) {
	cards := New()
	if len(cards) != 13*4 {
		t.Error("Wrong number of cards in new Deck")
	}
}

func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)
	exp := Card{Spade, Ace}
	if cards[0] != exp {
		t.Error("Expected Ace of Spades as first card. Received: ", cards[0])
	}
}

func TestSort(t *testing.T) {
	cards := New(Sort(Less))
	exp := Card{Spade, Ace}
	if cards[0] != exp {
		t.Error("Expected Ace of Spades as first card. Received: ", cards[0])
	}
}

func TestShuffle(t *testing.T) {
	cards := New()

	newCards := Shuffle(cards)

	if newCards[0] == cards[0] {
		t.Errorf("got same card %v", cards[0])
	}
}

func TestJokers(t *testing.T) {
	n := 3
	cards := New(Jokers(n))
	cnt := 0

	for _, card := range cards {
		if card.Suit == Joker {
			cnt++
		}
	}

	if cnt != n {
		t.Errorf("want jokers: %d \t for jokers: %d", n, cnt)
	}
}

func TestFilter(t *testing.T) {
	f := func(card Card) bool {
		return card.Rank == Two || card.Rank == Three
	}
	cards := New(Filter(f))

	for _, card := range cards {
		if card.Rank == Two || card.Rank == Three {
			t.Errorf("Filter failed: %v", card)
		}
	}
}

func TestDeck(t *testing.T) {
	n := 3
	cards := New(Deck(n))
	if len(cards) != 13*4*n {
		t.Errorf("Wrong number of cards in Deck. Wanted: %d\t Get: %d\n", 13*4*n, len(cards))
	}
}
