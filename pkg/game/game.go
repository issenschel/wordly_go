package game

import (
	"bufio"
	"fmt"
	"mygame/pkg/constants"
	"mygame/pkg/persistence"
	"mygame/pkg/word"
	"os"
	"strings"
)

// Ну а что тут объяснять?
type Game struct {
	dictionaryFile string
	attempts       []*word.Word
}

// NewGame создает новую игру и загружает словарь.
func NewGame(dictionaryFile string) *Game {
	return &Game{dictionaryFile: dictionaryFile}
}

// Start запускает игру.
func (g *Game) Start() {
	fmt.Print("\033[H\033[2J")
	attemptsCount := 0
	randomWord := persistence.GetRandomWord(g.dictionaryFile)

	fmt.Println("\033[1;35mПриветствую в Wordly! Попробуйте отгадать загаданное слово.\033")
	for attemptsCount < constants.AttemptsNumber {
		fmt.Printf("\033[1;34mПопытка %d из %d\n", attemptsCount+1, constants.AttemptsNumber)
		fmt.Print("\033[1;37mВведите слово: ")
		input := getUserInput()
		fmt.Print("\033[H\033[2J")

		if !persistence.IsWordValid(input, g.dictionaryFile) {
			fmt.Printf("\033[1;31mСлово '%s' не найдено в словаре или не соответствует длине в %d букв!\033\n", input, constants.WordLength)
			g.printAttempts()
			continue
		}

		current := word.NewWord(input)
		compare(current, word.NewWord(randomWord))
		g.attempts = append(g.attempts, current)
		g.printAttempts()

		if current.Equals(word.NewWord(randomWord)) {
			fmt.Printf("\033[1;32mПоздравляем, вы отгадали слово: %s\033\n", randomWord)
			break
		}

		attemptsCount++
		if attemptsCount == constants.AttemptsNumber {
			fmt.Printf("\033[1;31mСлово не угадано. Загаданное слово было: %s\033\n", randomWord)
		}
	}
}

// Выводим попытки
func (g *Game) printAttempts() {
	fmt.Println("\033[1;36mВаши попытки:\033")
	for _, attempt := range g.attempts {
		attempt.Print()
	}
	fmt.Println()
}

// Ввод пользователя
func getUserInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// compare сравнивает текущее слово с правильным и изменяет цвет букв.
func compare(current, correct *word.Word) {
	usedIndices := make(map[int]bool)

	// Сначала отмечаем зелёные буквы
	for i, letter := range current.Letters {
		if letter.Char == correct.Letters[i].Char {
			current.ChangeColor(i, constants.Green)
			usedIndices[i] = true
		}
	}

	// Затем отмечаем жёлтые буквы
	for i, letter := range current.Letters {
		if current.Letters[i].Color == constants.Gray {
			for j, correctLetter := range correct.Letters {
				if letter.Char == correctLetter.Char && !usedIndices[j] {
					current.ChangeColor(i, constants.Yellow)
					usedIndices[j] = true
					break
				}
			}
		}
	}
}
