package main

import (
	"fmt"
)

type Word struct {
	letters []*Letter
}

// Equals проверяет, совпадает ли слово с другим словом.
func (w *Word) Equals(other *Word) bool {
	for i, letter := range w.letters {
		if letter.char != other.letters[i].char {
			return false
		}
	}
	return true
}

// ChangeColor изменяет цвет буквы в слове по индексу.
func (w *Word) ChangeColor(index int, color string) {
	w.letters[index].color = color
}

// Print выводит слово в консоль.
func (w *Word) Print() {
	for _, letter := range w.letters {
		fmt.Print(letter)
	}
	fmt.Println()
}

func NewWord(word string) *Word {
	w := &Word{}
	for _, char := range word {
		w.letters = append(w.letters, NewLetter(char, gray))
	}
	return w
}
