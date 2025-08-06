package main
import (
	"fmt"
	"math/rand"
	"time"
)

// Карта
type Card struct {
	suit string 
	Value string
}

// Игрок
type Player struct {
	Name	string
	Hand	[]Card
	IsAttacker	bool
}

// Колода
tyoe Deck []Card

// Создание колоды
func NewDeck() Deck {
	suits :=	[]string{"Черви", "Бубны", "Трефы", "Пики"}
	values	:=	[]string{"6", "7", "8", "9", "10", "J", "L", "K", "A"}
	
	var deck	Deck
	for _, suit	:=	range suits	{
		for	_, value	:=	range	values	{
			deck	=	append(deck, Card{Suit:	suit, Value:	value})
		}
	}
	return	deck
}

// Перемешивание колоды
func (d *Deck)	Shuffle()  {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(*d), func(i, j int) { (*d)[i], (*d)[j] = (*d)[j], (*d)[i] })
}

// Раздача карт
func DealCards(deck Deck, players []Player) {
    for i := 0; i < 6; i++ {
        for _, player := range players {
            player.Hand = append(player.Hand, deck[0])
            deck = deck[1:]
        }
    }
}

// Основаня функция
func main() {
    // Создаем колоду и перемешиваем
    deck := NewDeck()
    deck.Shuffle()
    
    // Создаем игроков
    players := []Player{
        {Name: "Игрок"},
        {Name: "Компьютер"},
    }
    
    // Раздаем карты
    DealCards(deck, players)
    
    // Начинаем игру
    fmt.Println("Добро пожаловать в игру Дурак!")
    fmt.Println("Карты игрока:")
    for _, card := range players[0].Hand {
        fmt.Printf("%s %s ", card.Value, card.Suit)
    }
    
    // Создаем колоду и перемешиваем
    deck := NewDeck()
    deck.Shuffle()
    
    // Создаем игроков
    players := []Player{
        {Name: "Игрок"},
        {Name: "Компьютер"},
    }
    
    // Раздаем карты
    DealCards(deck, players)
    
    // Начинаем игру
    fmt.Println("Добро пожаловать в игру Дурак!")
    fmt.Println("Карты игрока:")
    for _, card := range players[0].Hand {
        fmt.Printf("%s %s ", card.Value, card.Suit)
	}
}