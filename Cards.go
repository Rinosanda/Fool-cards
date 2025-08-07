package main

import (
    "fmt"
    "math/rand"
    "time"
)


type Card struct {
    suit  string
    value string
}


type Player struct {
    name       string
    hand       []Card
    isAttacker bool
}


type Deck []Card


func NewDeck() Deck {
    suits := []string{"Черви", "Бубны", "Трефы", "Пики"}
    values := []string{"6", "7", "8", "9", "10", "В", "Д", "К", "Т"}
    
    var deck Deck
    for _, suit := range suits {
        for _, value := range values {
            deck = append(deck, Card{suit: suit, value: value})
        }
    }
    return deck
}


func (d *Deck) Shuffle() {
    rand.Seed(time.Now().UnixNano())
    rand.Shuffle(len(*d), func(i, j int) { (*d)[i], (*d)[j] = (*d)[j], (*d)[i] })
}


func DealCards(deck Deck, players []Player) {
    for i := 0; i < 6; i++ {
        for j, player := range players {
            players[j].hand = append(player.hand, deck[0])
            deck = deck[1:]
        }
    }
}


func removeCard(hand []Card, card Card) []Card {
    newHand := []Card{}
    for _, c := range hand {
        if c.value != card.value || c.suit != card.suit {
            newHand = append(newHand, c)
        }
    }
    return newHand
}


func canPlayCard(attacker *Player, defender *Player, card Card) bool {
    
    return true
}


func playerTurn(player *Player, opponent *Player) {
    fmt.Printf("Ход игрока %s\n", player.name)
    
    chosenCard := player.hand[0] 
    fmt.Printf("Выбранная карта: %s %s\n", chosenCard.value, chosenCard.suit)
    
    
    if canPlayCard(player, opponent, chosenCard) {
        
        player.hand = removeCard(player.hand, chosenCard)
        opponent.hand = append(opponent.hand, chosenCard)
    } else {
        fmt.Println("Некорректный ход!")
    }
}


func computerTurn(computer *Player, human *Player) {
    fmt.Println("Ход компьютера")
    
    chosenCard := computer.hand[0] 
    fmt.Printf("Компьютер выбрал карту: %s %s\n", chosenCard.value, chosenCard.suit)
    
    
    computer.hand = removeCard(computer.hand, chosenCard)
    human.hand = append(human.hand, chosenCard)
}


func gameLoop(players []Player) {
    attacker := &players[0] 
    defender := &players[1]
    
    for {
        
        if len(attacker.hand) == 0 {
            fmt.Printf("Победил %s!\n", defender.name)
            break
        }
        if len(defender.hand) == 0 {
            fmt.Printf("Победил %s!\n", attacker.name)
            break
        }
        
        
        if attacker.name == "Игрок" {
            playerTurn(attacker, defender)
        } else {
            computerTurn(attacker, defender)
        }
        
        
        attacker, defender = defender, attacker
    }
}

func main() {
    
    deck := NewDeck()
	deck.Shuffle()  
    
    
    players := []Player{
        {name: "Игрок", hand: []Card{}, isAttacker: true},
        {name: "Компьютер", hand: []Card{}, isAttacker: false},
    }
    
   
    DealCards(deck, players)
    
    
    gameLoop(players)
}


