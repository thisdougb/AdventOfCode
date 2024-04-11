package main

import (
	"log"
	"strconv"
	"strings"
)

/*

Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11

*/

type Card struct {
	id             int
	winningNumbers []int
	myNumbers      []int
	copies         int
}

func NewScratchcard(id int, winning []int, my []int) *Card {

	card := &Card{}

	card.id = id
	card.winningNumbers = winning
	card.myNumbers = my

	return card
}

func (c *Card) score() int {

	total := 0
	wins := 0

	for _, n := range c.myNumbers {
		for _, m := range c.winningNumbers {
			if n == m {
				wins++
			}
		}
	}

	if wins > 0 {
		total = 1 // first match point
		for i := 1; i < wins; i++ {
			total *= 2
		}
	}

	return total
}

func (c *Card) score2() int {

	wins := 0

	for _, n := range c.myNumbers {
		for _, m := range c.winningNumbers {
			if n == m {
				wins++
			}
		}
	}

	return wins
}

func day4() {

	totalPoints := 0

	cards := loadCards(readFileContents("day4_input.txt"))

	for _, c := range cards {
		//log.Println("card:", c.id, c.score())
		totalPoints += c.score()
	}
	log.Println("total:", totalPoints)

	part2("day4_input.txt")
}

func part2(filepath string) {

	totalCards := 0

	cards := loadCards(readFileContents(filepath))

	// process each card
	for i, c := range cards {

		log.Println("process card", c.id, "score", c.score2(), "copies", c.copies)

		// for every winning number we roll down the card list
		for j := 0; j < c.score2(); j++ {

			// add the initial win to subsequent card
			cards[i+j+1].copies++

			//for every copy, add another win to subsequent card
			for k := 0; k < c.copies; k++ {
				cards[i+j+1].copies++
			}
			log.Println("card", cards[c.id+j].id, "copies", cards[c.id+j].copies)
		}
	}

	for _, c := range cards {
		totalCards += 1 + c.copies
	}
	log.Println("total", totalCards)

}

func strToIntArray(in []string) []int {

	out := []int{}

	for _, i := range in {
		if i == "" {
			continue
		}
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		out = append(out, j)
	}
	return out
}

// loadCards takes input lines of card data and returns
// an array of cards
func loadCards(input []string) []*Card {

	cards := []*Card{}

	for i, line := range input {

		// first split
		meta := strings.Split(line, ":")

		// second split
		cardNumbers := strings.Split(meta[1], "|")

		winningNumbers := strings.Split(cardNumbers[0], " ")
		myNumbers := strings.Split(cardNumbers[1], " ")

		c := NewScratchcard(i+1, strToIntArray(winningNumbers), strToIntArray(myNumbers))
		cards = append(cards, c)
	}

	return cards
}
