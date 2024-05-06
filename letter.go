package main

import (
	"fmt"
)

// Letter представляет собой структуру для буквы со стилем отображения.
type Letter struct {
	char  rune
	color string
}

// NewLetter создает новую букву с заданным символом и цветом.
func NewLetter(char rune, color string) *Letter {
	return &Letter{char: char, color: color}
}

// String возвращает строковое представление буквы с учетом цвета.
func (l *Letter) String() string {
	return fmt.Sprintf(l.color, string(l.char))
}
