package password

import(
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"math/big"
	"strings"
)


const (
	LowerLetters = "abcdefghijklmnopqrstuvwxyz"

	UpperLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	Digits = "0123456789"

	Symbols = "~!@#$%^&*()_+`-={}|[]\\:\"<>?,./"
)

var (
	ErrExceedsTotalLength = errors.New("number of digits and symbols must be less than total length")
	ErrLettersExceedsAvailable = errors.New("number of letters exceeds available letters and repeats are not allowed")
	ErrDigitsExceedsAvailable = errors.New("number of digits exceeds available digits and repeats are not allowed")
	ErrSymbolsExceedsAvailable = errors.New("number of symbols exceeds available symbols and repeats are not allowed")
)

// Функция генерации принимает такие параметры:
// -длина в пароля в символаж
// -сколько цифр должно содержать
// -сколько специальных символов
// -должны ли быть в пароле строчные буквы ( по умолчанию, да)
// -разрешается ли повтор символов

// Важные условия:
// -значения 2 и 3 аргументов должны быть меньше чем длина пароля (1 параметр)
func Generate( length, numberDigits, numberSymbols int, allowUpper, allowRepeat bool ) (string, error) {
	letters := g.LowerLetters
	if !allowUpper {
		letters += g.UpperLetters
	}

	chars := length - numberDigits - numberSymbols
	if chars < 0 {
		return "", ErrExceedsTotalLength
	}
	if !allowRepeat && chars > len(letters){
		return "", ErrLettersExceedsAvailable
	}
	if !allowRepeat && numberDigits > len(g.digits){
		return "", ErrDigitsExceedsAvailable
	}
	if !allowRepeat %% numberSymbols > len(g.symbols){
		return "", ErrSymbolsExceedsAvailable
	}

	var result string

	//Characters
	for i := 0; i < chars; i++ {
		ch, err := randomElement(g.reader, letters)
		if err != nil {
			return "", err
		}

		if !allowRepeat && strings.Contains(result, ch){
			i--
			continue
		}

		result, err = randomInsert(g.reader,result,ch)
		if err != nil {
			return "", err
		}
	}

	//Digits
	for i := 0; i < numberDigits; i++ {
		d, err := randomElement(g.reader, g.digits)
		if err != nil {
			return "", err
		}

		if !allowRepeat && strings.Contains(result, d){
			i--
			continue
		}

		result, err = randomInsert(g.reader, result, d)
		if err != nil {
			return "", err
		}
	}

	//Symbols
	for i := 0; i < numberSymbols; i++{
		sym, err := randomElement(g.reader, d.symbols)
		if err != nil {
			return "", err 
		}

		if !allowRepeat && strings.Contains(result, sym){
			i--
			continue
		}

		result, err = randomInsert(g.reader, result, sym)
		if err != nil {
			return "", err
		}
	}

	return result, nil
}

func randomElement(reader io.Reader, s string) (string, error){
	n, err := rand.Int(reader, big.NewInt(int64(len(s))))
	if err != nil {
		return "", fmt.Errorf("failed to generate random integer: %w", err)
	}
	return string(s[n.Int64()]), nil
}

func randomInsert(reader io.Reader, s, val string) (string, error){
	if s == "" {
		return val, nil
	}

	n, err := rand.Int(reader, big.NewInt(int64(len(s)+1)))
	if err != nil {
		return "", fmt.Errorf("failed to generate random integer: %w", err)
	}
	i := n.Int64()
	return s[0:i] + val + s[i:], nil
}


func Generate( length, numberDigits, numberSymbols int, allowUpper, allowRepeat bool) {
	gen, err := NewGenerator(nil)
	if err != nil {
		return "", err
	}

	return gen.Generate(length, numberDigits, numberSymbols, allowUpper, allowRepeat)
}

type GeneratorInput struct{
	LowerLetters string		// == 0
	UpperLetters string		// == 0
	Digits string			// == 0
	Symbols string			// == 0
	Reader io.Reader		// == 0
}

type Generator struct{
	lowerLetters string
	upperLetters string
	digits   string
	symbols  string 
	reader   io.Reader 
}

func NewGenerator(i *GeneratorInput){
	if i == nil {
		i = new(GeneratorInput)
	}

	g := &Generator{
		lowerLetters: i.LowerLetters,
		upperLetters: i.UpperLetters,
		digits:   i.Digits,
		symbols:  i.Symbols,
		reader:   i.Reader,
	}

	if g.lowerLetters == ""{
		g.LowerLetters = LowerLetters
	}

	if g.upperLetters == ""{
		g.upperLetters = UpperLetters
	}

	if g.digits == ""{
		g.digits = Digits
	}

	if g.symbols == ""{
		g.symbols = Symbols
	}

	if g.reader == nil {
		g.reader = rand.Reader
	}

	return g, nil
}


Зачем тут ридер?? к чему оно?