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
	cards:= New()
	// 13 ranks 4 suits
	if len(cards) != 13*4 {
		t.Error("Wrong number of cards in a new deck.")
	}
}
