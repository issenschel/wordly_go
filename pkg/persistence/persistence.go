package persistence

import (
	"bufio"
	"math/rand"
	"mygame/pkg/constants"
	"os"
)

func IsWordValid(word, dictionaryFile string) bool {
	file, err := os.Open(dictionaryFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == word {
			return true
		}
	}
	return false
}

// Жёсткий рандомайзер слов
func GetRandomWord(dictionaryFile string) string {
	var currentWord string

	file, err := os.Open(dictionaryFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	randomIndex := rand.Intn(constants.WordQuantity)

	scanner := bufio.NewScanner(file)
	for i := 0; i <= randomIndex; i++ {
		scanner.Scan()
		currentWord = scanner.Text()
	}
	return currentWord
}
