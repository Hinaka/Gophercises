package deck

import (
	"fmt"
	"testing"
)

func ExampleCard_String() {
	fmt.Println(Card{Rank: Ace, Suit: Heart})
	fmt.Println(Card{Rank: Three, Suit: Spade})
	fmt.Println(Card{Rank: Jack, Suit: Diamond})
	fmt.Println(Card{Rank: King, Suit: Club})
	fmt.Println(Card{Suit: Joker})

	// Output:
	// Ace of Hearts
	// Three of Spades
	// Jack of Diamonds
	// King of Clubs
	// Joker
}

func TestNew(t *testing.T) {
	cards := New()
	// 13 ranks 4 suits
	if len(cards) != 13*4 {
		t.Error("Wrong number of cards in a new deck.")
	}
}

func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)
	exp := Card{Rank: Ace, Suit: Spade}
	if cards[0] != exp {
		t.Error("Expected Ace of Spades as first card. Received:", cards[0])
	}
}

func TestSort(t *testing.T) {
	cards := New(Sort(Less))
	exp := Card{Rank: Ace, Suit: Spade}
	if cards[0] != exp {
		t.Error("Expected Ace of Spades as first card. Received:", cards[0])
	}
}

func TestJokers(t *testing.T) {
	cards := New(Jokers(3))
	count := 0
	for _, c := range cards {
		if c.Suit == Joker {
			count++
		}
	}
	if count != 3 {
		t.Error("Expected 3 Jokers, received:", count)
	}
}

func TestFilter(t *testing.T) {
	filter := func(card Card) bool {
		return card.Rank == Two || card.Rank == Three
	}
	cards := New(Filter(filter))
	for _, c := range cards {
		if c.Rank == Two || c.Rank == Three {
			t.Error("Expected all twos and threes to be filtered out.")
		}
	}
}
