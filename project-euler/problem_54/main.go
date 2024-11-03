package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func elapsed(what string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", what, time.Since(start))
	}
}

const (
	Diamonds byte	= 'D'
	Clubs			= 'C'
	Hearts			= 'H'
	Spades			= 'S'
)

const (
	One byte = iota
	Two
	Three
	Four
	Five
	Six	
	Seven
	Eight
	Nine
	Ten	
	Jack
	Queen
	King
	Ace	
	IMPOSSIBLE
)

const (
	HighCard byte = iota
	OnePair
	TwoPairs
	ThreeOfAKind
	Straight
	Flush
	FullHouse
	FourOfAKind
	StraightFlush
	RoyalFlush
)

var card_ranks_labels map[byte]string
var combination_ranks_labels map[byte]string

func init() {
	card_ranks_labels = map[byte]string {
		One			:	"1",
		Two			:	"2",
		Three		:	"3",
		Four		:	"4",
		Five		:	"5",
		Six			:	"6",
		Seven		:	"7",
		Eight		:	"8",
		Nine		:	"9",
		Ten			:	"T",
		Jack		:	"J",
		Queen		:	"Q",
		King		:	"K",
		Ace			:	"A",
		IMPOSSIBLE	:	"IMPOSSIBLE",
	}

	combination_ranks_labels = map[byte]string {
		HighCard		:	"High Card",
		OnePair			:	"One Pair",
		TwoPairs		:	"Two Pairs",
		ThreeOfAKind	:	"Three Of A Kind",
		Straight		:	"Straight",
		Flush			:	"Flush",
		FullHouse		:	"Full House",
		FourOfAKind		:	"Four Of A Kind",
		StraightFlush	:	"Straight Flush",
		RoyalFlush		:	"Royal Flush",
	}
}

const (
	Tie byte = iota
	AWins
	BWins
)

type PokerCard struct {
	Rank byte
	Suit byte
}

func (x PokerCard) toString() string {
	return card_ranks_labels[x.Rank] + string(x.Suit)
}

func NewPokerCard(card string) PokerCard {
	rank := One
	if card[0] >= byte('1') && card[0] <= byte('9') {
		rank = card[0] - byte('1')
	} else {
		rank_map := map[byte]byte { 'T': Ten, 'J': Jack, 'Q': Queen, 'K': King, 'A': Ace }
		rank = rank_map[card[0]]
	}
	return PokerCard{rank, card[1]}
}

type PokerCardSlice []PokerCard

func (x PokerCardSlice) Len() int 			{ return len(x) }
func (x PokerCardSlice) Less(i, j int) bool { return x[i].Rank > x[j].Rank || (x[i].Rank == x[j].Rank && x[i].Suit > x[j].Suit)}
func (x PokerCardSlice) Swap(i, j int) 		{ x[i], x[j] = x[j], x[i]}

type PokerHand struct {
	Cards []PokerCard
	CardGroups map[byte][]PokerCard
	Rank byte
}

func (x PokerHand) toString() string {
	card_strings := []string {}
	for _, card := range x.Cards {
		card_strings = append(card_strings, card.toString())
	}
	return strings.Join(card_strings, ", ") + " Rank: " + combination_ranks_labels[x.Rank]
}

type RankRule struct {
	Rank byte
	IsApplicable func(PokerHand) bool
}

func NewPokerHand(cards []PokerCard) PokerHand {
	poker_cards := PokerCardSlice(cards)
	sort.Sort(poker_cards)

	card_groups := map[byte][]PokerCard {}
	for _, card := range cards {
		card_groups[card.Rank] = append(card_groups[card.Rank], card)
	}

	poker_hand := PokerHand{ 
		Cards: cards, 
		CardGroups: card_groups,
	}

	ranks := []RankRule { 
		{ RoyalFlush, is_royal_flush },
		{ StraightFlush, is_straight_flush },
		{ FourOfAKind, is_four_of_a_kind },
		{ FullHouse, is_full_house },
		{ Flush, is_flush },
		{ Straight, is_straight },
		{ ThreeOfAKind, is_three_of_kind },
		{ TwoPairs, is_two_pair },
		{ OnePair, is_pair },
		{ HighCard, func (PokerHand) bool { return true } },
	}

	for _, rank := range ranks {
		if rank.IsApplicable(poker_hand) {
			poker_hand.Rank = rank.Rank
			break
		}
	}
	
	return poker_hand
}

func read_hands() [][]PokerHand {
	file, err := os.Open("p054_poker.txt")
	if err != nil {
		log.Fatal(err)
		return nil
	}

	res := [][]PokerHand {}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		hands := strings.Split(scanner.Text(), " ")
		cards := []PokerCard {}
		for _, card := range hands {
			cards = append(cards, NewPokerCard(card))
		}
		player_one_hand := NewPokerHand(cards[:5])
		player_two_hand := NewPokerHand(cards[5:])
		res = append(res, []PokerHand { player_one_hand, player_two_hand })
	}
	return res
}

func is_pair(hand PokerHand) bool {
	cnt := 0
	for _, cards := range hand.CardGroups {
		if len(cards) == 2 {
			cnt++
		}
	}
	return cnt == 1
}

func get_pair_rank(hand PokerHand) byte {
	for rank, cards := range hand.CardGroups {
		if len(cards) == 2 {
			return rank
		}
	}
	return IMPOSSIBLE
}

func is_two_pair(hand PokerHand) bool {
	cnt := 0
	for _, cards := range hand.CardGroups {
		if len(cards) == 2 {
			cnt++
		}
	}
	return cnt == 2
}

func get_two_pair_ranks(hand PokerHand) []byte {
	ranks := []byte{}
	for rank, cards := range hand.CardGroups {
		if len(cards) == 2 {
			ranks = append(ranks, rank)
		}
	}
	if ranks[0] < ranks[1] {
		ranks[0], ranks[1] = ranks[1], ranks[0]
	}
	return ranks
}

func is_three_of_kind(hand PokerHand) bool {
	for _, cards := range hand.CardGroups {
		if len(cards) == 3 {
			return true
		}
	}
	return false
}

func get_three_of_a_kind_rank(hand PokerHand) byte {
	for rank, cards := range hand.CardGroups {
		if len(cards) == 3 {
			return rank
		}
	}
	return IMPOSSIBLE
}

func is_straight(hand PokerHand) bool {
	for i := 1; i < len(hand.Cards); i++ {
		if hand.Cards[i].Rank != hand.Cards[i - 1].Rank - 1 {
			return false
		}
	}
	return true
}

func is_flush(hand PokerHand) bool {
	for i := 1; i < len(hand.Cards); i++ {
		if hand.Cards[i].Suit != hand.Cards[i - 1].Suit {
			return false
		}
	}
	return true
}

func is_full_house(hand PokerHand) bool {
	if len(hand.CardGroups) == 2 {
		for _, cards := range hand.CardGroups {
			if len(cards) == 2 || len(cards) == 3 {
				return true
			}
		}
	}
	return false
}

func get_full_house_rank(hand PokerHand) []byte {
	ranks := []byte { IMPOSSIBLE, IMPOSSIBLE }
	for rank, cards := range hand.CardGroups {
		if len(cards) == 3 {
			ranks[0] = rank
		}
		if len(cards) == 2 {
			ranks[1] = rank
		}
	}
	return ranks
}

func is_four_of_a_kind(hand PokerHand) bool {
	if len(hand.CardGroups) == 2 {
		for _, cards := range hand.CardGroups {
			if len(cards) == 1 || len(cards) == 4 {
				return true
			}
		}
	}
	return false
}

func get_four_of_a_kind_rank(hand PokerHand) byte {
	for rank, cards := range hand.CardGroups {
		if len(cards) == 4 {
			return rank
		}
	}
	return IMPOSSIBLE
}

func is_straight_flush(hand PokerHand) bool {
	return is_straight(hand) && is_flush(hand)
}

func is_royal_flush(hand PokerHand) bool {
	return is_straight_flush(hand) && hand.Cards[0].Rank == Ace
}

// returns 0 - tie, 1 - a wins, 2 - b wins
func play_ranks(rank_a byte, rank_b byte) byte {
	if rank_a > rank_b {
		return AWins
	} else if rank_b > rank_a {
		return BWins
	}
	return Tie
}

// returns 0 - tie, 1 - a wins, 2 - b wins
func play_same_rank(a PokerHand, b PokerHand, rank byte) byte {
	switch rank {
	case HighCard:
		for i := 0; i < len(a.Cards); i++ {
			win := play_ranks(a.Cards[i].Rank, b.Cards[i].Rank)
			if win != Tie {
				return win
			}
		}
		return Tie
	case OnePair:
		a_rank := get_pair_rank(a)
		b_rank := get_pair_rank(b)
		win := play_ranks(a_rank, b_rank)
		if win != Tie {
			return win
		}
		return play_same_rank(a, b, HighCard)
	case TwoPairs:
		a_ranks := get_two_pair_ranks(a)
		b_ranks := get_two_pair_ranks(b)
		for i := 0; i < 2; i++ {
			win := play_ranks(a_ranks[0], b_ranks[0])
			if win != Tie {
				return win
			}
		}
		return play_same_rank(a, b, HighCard)
	case ThreeOfAKind:
		a_rank := get_three_of_a_kind_rank(a)
		b_rank := get_three_of_a_kind_rank(b)
		win := play_ranks(a_rank, b_rank)
		if win != Tie {
			return win
		}
		return play_same_rank(a, b, HighCard)
	case Straight:
		a_rank := a.Cards[0].Rank
		b_rank := b.Cards[0].Rank
		return play_ranks(a_rank, b_rank)
	case Flush:
		return play_same_rank(a, b, HighCard)
	case FullHouse:
		a_ranks := get_full_house_rank(a)
		b_ranks := get_full_house_rank(b)
		for i := 0; i < 2; i++ {
			win := play_ranks(a_ranks[0], b_ranks[0])
			if win != Tie {
				return win
			}
		}
		return Tie
	case FourOfAKind:
		a_rank := get_four_of_a_kind_rank(a)
		b_rank := get_four_of_a_kind_rank(b)
		win := play_ranks(a_rank, b_rank)
		if win != Tie {
			return win
		}
		return play_same_rank(a, b, HighCard)
	case StraightFlush:
		a_rank := a.Cards[0].Rank
		b_rank := b.Cards[0].Rank
		return play_ranks(a_rank, b_rank)
	case RoyalFlush:
		return Tie
	}
	return Tie
}

// returns 0 - tie, 1 - a wins, 2 - b wins
func play(a PokerHand, b PokerHand) byte {
	win := play_ranks(a.Rank, b.Rank)
	if win != Tie {
		return win
	}
	return play_same_rank(a, b, a.Rank)
}

func solve() int64 {
	all_hands := read_hands()
	res := 0
	file, err := os.OpenFile("output.txt", os.O_RDWR|os.O_CREATE, 0755)
	defer file.Close()

	if err != nil {
		log.Fatal(err)
		return -1
	}
	writer := bufio.NewWriter(file)
	for id, hands := range all_hands {
		writer.WriteString("______________________________Game:" + strconv.Itoa(id + 1) + "\n")
		writer.WriteString("Player 1: " + hands[0].toString() + "\n")
		writer.WriteString("Player 2: " + hands[1].toString() + "\n")
		win := play(hands[0], hands[1])
		if win == AWins {
			writer.WriteString("Player 1 wins\n")
			res++
		} else if win == BWins {
			writer.WriteString("Player 2 wins\n")
		} else {
			writer.WriteString("Tie\n")
		}
	}
	return int64(res)
}

func main() {
	defer elapsed("main")()
	fmt.Println(solve())
}