package deck

import "fmt"

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
