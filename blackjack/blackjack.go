package main

import "fmt"

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten", "jack", "queen", "king":
		return 10
	case "ace":
		return 11
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	card1Value, card2Value, dealerCardValue := ParseCard(card1), ParseCard(card2), ParseCard(dealerCard)
	cardsSum := card1Value + card2Value
	isBlackjack := cardsSum == 21
	cardsSumBetween17and20 := cardsSum >= 17 && cardsSum <= 20
	cardsSumBetween12and16 := cardsSum >= 12 && cardsSum <= 16
	switch {
	case card1 == "ace" && card2 == "ace":
		return "P" // Split
	case isBlackjack && dealerCardValue < 10:
		return "W" // Automatically win
	case isBlackjack && dealerCardValue >= 10, cardsSumBetween17and20, cardsSumBetween12and16 && dealerCardValue < 7:
		return "S" // Stand
	case cardsSumBetween12and16 && dealerCardValue >= 7, cardsSum <= 11:
		return "H" // Hit
	default:
		return ""
	}
}

func main() {
	fmt.Println(ParseCard("ace"))
	fmt.Println(FirstTurn("ace", "ace", "jack"))
	fmt.Println(FirstTurn("ace", "king", "ace"))
	fmt.Println(FirstTurn("king", "ace", "queen"))
}
