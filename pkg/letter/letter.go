package letter

import (
	"fmt"
)

// Letter представляет собой структуру для буквы со стилем отображения.
type Letter struct {
	Char  rune
	Color string
}

// NewLetter создает новую букву с заданным символом и цветом.
func NewLetter(char rune, color string) *Letter {
	return &Letter{Char: char, Color: color}
}

// String возвращает строковое представление буквы с учетом цвета.
func (l *Letter) String() string {
	return fmt.Sprintf(l.Color, string(l.Char))
}
