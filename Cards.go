package main

import (
    "fmt"
    "math/rand"
    "time"
)

// Карта
type Card struct {
    suit string
    value string
}

// Игрок
type Player struct {
    name string
    hand []Card
    isAttacker bool
}

// Колода
type Deck []Card

// Создание колоды
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

// Перемешивание колоды
func (d *Deck) Shuffle() {
    rand.Seed(time.Now().UnixNano())
    rand.Shuffle(len(*d), func(i, j int) { (*d)i, (*d)j = (*d)j, (*d)i })
}

// Раздача карт
func DealCards(deck Deck, players []Player) {
    for i := 0; i < 6; i++ {
        for j, player := range players {
            players[j].hand = append(player.hand, deck[0])
            deck = deck[1:]
        }
    }
}

// Функция для удаления карты из руки
func removeCard(hand []Card, card Card) []Card {
    newHand := []Card{}
    for _, c := range hand {
        if c.value != card.value || c.suit != card.suit {
            newHand = append(newHand, c)
        }
    }
    return newHand
}

// Функция для проверки возможности хода
func canPlayCard(attacker *Player, defender *Player, card Card) bool {
    // Простая проверка - можно ходить любой картой
    return true
}

// Функция для выполнения хода игрока
func playerTurn(player *Player, opponent *Player) {
    fmt.Printf("Ход игрока %s\n", player.name)
    // Здесь логика выбора карты игроком
    chosenCard := player.hand[0] // Временная реализация
    fmt.Printf("Выбранная карта: %s %s\n", chosenCard.value, chosenCard.suit)
    
    // Проверка возможности хода
    if canPlayCard(player, opponent, chosenCard) {
        // Выполнение хода
        player.hand = removeCard(player.hand, chosenCard)
        opponent.hand = append(opponent.hand, chosenCard)
    } else {
        fmt.Println("Некорректный ход!")
    }
}

// Функция для выполнения хода компьютера
func computerTurn(computer *Player, human *Player) {
    fmt.Println("Ход компьютера")
    // Простая реализация ИИ
    chosenCard := computer.hand[0] // Компьютер ходит первой картой
    fmt.Printf("Компьютер выбрал карту: %s %s\n", chosenCard.value, chosenCard.suit)
    
    // Выполнение хода
    computer.hand = removeCard(computer.hand, chosenCard)
    human.hand = append(human.hand, chosenCard)
}

// Основной игровой цикл
func gameLoop(players []Player) {
    attacker := &players[0] // Первый атакующий
    defender := &players[1]
    
    for {
        // Проверка на победу
        if len(attacker.hand) == 0 {
            fmt.Printf("Победил %s!\n", defender.name)
            break
        }
        if len(defender.hand) == 0 {
            fmt.Printf("Победил %s!\n", attacker.name)
            break
        }
        
        // Ход атакующего
        if attacker.name == "Игрок" {
            playerTurn(attacker, defender)
        } else {
            computerTurn(attacker, defender)
        }
        
        // Смена ролей
        attacker, defender = defender, attacker
    }
}
