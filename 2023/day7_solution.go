package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"

	"golang.org/x/exp/slices"
)

func day7() {
	//day7_part1("day7_input.txt")
	day7_part1("day7_input.txt")
}

func day7_part2(filepath string) {

	input := readFileContents(filepath)

	hands := []string{}
	bids := map[string]int{}

	for _, s := range input {

		hand := s[:5]
		bid := s[6:]
		bids[hand], _ = strconv.Atoi(bid)

		//fmt.Printf("hand: %s, bid: %d\n", hand, bids[hand])
		hands = addHandToHands2(hand, hands)

	}

	var winnings int

	for i, h := range hands {
		winnings += bids[h] * (i + 1)
		fmt.Printf("%2d %s %d\n", i+1, h, bids[h])
	}

	fmt.Println(hands)
	fmt.Printf("winnings: %d\n", winnings)

}

func cardValue2(card rune) int {
	// A, K, Q, J, T, 9, 8, 7, 6, 5, 4, 3, or 2
	switch card {
	case 'A':
		return 14
	case 'K':
		return 13
	case 'Q':
		return 12
	case 'J':
		return 1
	case 'T':
		return 10
	default:
		val, _ := strconv.Atoi(string(card))
		return val
	}
}

/*
	hand {
	  cards = AAK44
	  groupings =  []{
		  {value: A, count: 2}
		  {value: 4, count: 2}
		  {value: K, count: 1}
	  }
*/
func getType2(hand string) int {

	cards := map[rune]int{}

	for _, c := range hand {
		cards[c]++
	}

	counts := []int{}
	for _, v := range cards {
		counts = append(counts, v)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(counts)))

	// five of a kind - AAAAA
	if len(cards) == 1 {
		return 7
	}

	if len(cards) == 2 {

		log.Println("cards", cards)
		log.Println("counts", counts)

		// four of a kind - AA8AA
		if counts[0] == 4 {
			return 6
		}

		// fullhouse - 23332
		if counts[0] == 3 {
			return 5
		}
	}

	if len(cards) == 3 {

		// three of a kind - TTT98
		if counts[0] == 3 && counts[1] == 1 && counts[2] == 1 {
			return 4
		}

		// two pair - 23432
		if counts[0] == 2 && counts[1] == 2 && counts[2] == 1 {
			return 3
		}

	}

	if len(cards) == 4 {
		// one pair - A23A4
		return 2
	}

	// high card - 23456
	return 1
}

func getStrongestHandType(hand string) string {

	return ""
}
func firstBeatsSecond2(first string, second string) bool {

	//fmt.Printf("test if %s beats %s\n", first, second)

	if getType2(first) > getType2(second) {
		//fmt.Printf("%s beats %s as a type\n", first, second)
		return true
	}

	if getType2(first) == getType2(second) {

		//fmt.Printf("%s equals %s as a type\n", first, second)

		// loop each char in the strings
		for i, c := range first {

			secondChar := []rune(second)[i]

			if cardValue2(c) < cardValue2(secondChar) {
				//fmt.Printf("%c doesn't beat %c as a char at %d\n", c, secondChar, i)
				return false
			}
			if cardValue2(c) > cardValue2(secondChar) {
				//fmt.Printf("%c beats %c as a char at %d\n", c, secondChar, i)
				return true
			}
		}
	}

	return false
}

func addHandToHands2(hand string, hands []string) []string {

	if len(hands) == 0 {
		hands = append(hands, hand)
		//fmt.Printf("hand: %s, bid: %d first hand\n", hand, bids[hand])
		return hands
	}

	for i, h := range hands {

		// if the hand to insert beats this one, move on
		if firstBeatsSecond2(hand, h) {

			//fmt.Printf("first: %s, beats second: %s at %d\n", hand, h, i)

			if i == len(hands)-1 {
				//fmt.Printf("first: %s, beats second: %s at %d append last item\n", hand, h, i)
				hands = append(hands, hand)
				return hands
			} else {
				continue
			}
		}
		//fmt.Printf("first: %s, does not beat second: %s at %d\n", hand, h, i)
		hands = slices.Insert(hands, i, hand)
		//fmt.Printf("insert %s into %v at %d\n", hand, hands, i)
		return hands
	}

	return hands
}

func day7_part1(filepath string) {

	input := readFileContents(filepath)

	hands := []string{}
	bids := map[string]int{}

	for _, s := range input {

		hand := s[:5]
		bid := s[6:]
		bids[hand], _ = strconv.Atoi(bid)

		//fmt.Printf("hand: %s, bid: %d\n", hand, bids[hand])
		hands = addHandToHands(hand, hands)

	}

	var winnings int

	for i, h := range hands {
		winnings += bids[h] * (i + 1)
		fmt.Printf("%2d %s %d\n", i+1, h, bids[h])
	}

	fmt.Println(hands)
	fmt.Printf("winnings: %d\n", winnings)

}

func addHandToHands(hand string, hands []string) []string {

	if len(hands) == 0 {
		hands = append(hands, hand)
		//fmt.Printf("hand: %s, bid: %d first hand\n", hand, bids[hand])
		return hands
	}

	for i, h := range hands {

		// if the hand to insert beats this one, move on
		if firstBeatsSecond(hand, h) {

			//fmt.Printf("first: %s, beats second: %s at %d\n", hand, h, i)

			if i == len(hands)-1 {
				//fmt.Printf("first: %s, beats second: %s at %d append last item\n", hand, h, i)
				hands = append(hands, hand)
				return hands
			} else {
				continue
			}
		}
		//fmt.Printf("first: %s, does not beat second: %s at %d\n", hand, h, i)
		hands = slices.Insert(hands, i, hand)
		//fmt.Printf("insert %s into %v at %d\n", hand, hands, i)
		return hands
	}

	return hands
}

func firstBeatsSecond(first string, second string) bool {

	//fmt.Printf("test if %s beats %s\n", first, second)

	if getType(first) > getType(second) {
		//fmt.Printf("%s beats %s as a type\n", first, second)
		return true
	}

	if getType(first) == getType(second) {

		//fmt.Printf("%s equals %s as a type\n", first, second)

		// loop each char in the strings
		for i, c := range first {

			secondChar := []rune(second)[i]

			if cardValue(c) < cardValue(secondChar) {
				//fmt.Printf("%c doesn't beat %c as a char at %d\n", c, secondChar, i)
				return false
			}
			if cardValue(c) > cardValue(secondChar) {
				//fmt.Printf("%c beats %c as a char at %d\n", c, secondChar, i)
				return true
			}
		}
	}

	return false
}

func cardValue(card rune) int {
	// A, K, Q, J, T, 9, 8, 7, 6, 5, 4, 3, or 2
	switch card {
	case 'A':
		return 14
	case 'K':
		return 13
	case 'Q':
		return 12
	case 'J':
		return 11
	case 'T':
		return 10
	default:
		val, _ := strconv.Atoi(string(card))
		return val
	}
}

func getType(hand string) int {

	cards := map[rune]int{}

	for _, c := range hand {
		cards[c]++
	}

	counts := []int{}
	for _, v := range cards {
		counts = append(counts, v)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(counts)))

	// five of a kind - AAAAA
	if len(cards) == 1 {
		return 7
	}

	if len(cards) == 2 {

		// four of a kind - AA8AA
		if counts[0] == 4 {
			return 6
		}

		// fullhouse - 23332
		if counts[0] == 3 {
			return 5
		}
	}

	if len(cards) == 3 {

		// three of a kind - TTT98
		if counts[0] == 3 && counts[1] == 1 && counts[2] == 1 {
			return 4
		}

		// two pair - 23432
		if counts[0] == 2 && counts[1] == 2 && counts[2] == 1 {
			return 3
		}

	}

	if len(cards) == 4 {
		// one pair - A23A4
		return 2
	}

	// high card - 23456
	return 1
}
