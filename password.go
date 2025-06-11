


const (
	LowerLetters = "abcdefghijklmnopqrstuvwxyz"

	UpperLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	Digits = "0123456789"

	Symbols = "~!@#$%^&*()_+`-={}|[]\\:\"<>?,./"
)

Функция генерации принимает такие параметры:
-длина в пароля в символаж
-сколько цифр должно содержать
-сколько специальных символов
-должны ли быть в пароле строчные буквы ( по умолчанию, да)
-разрешается ли повтор символов

Важные условия:
-значения 2 и 3 аргументов должны быть меньше чем длина пароля (1 параметр)
func Generate( length, numberDigits, numberSymbols int, allowUpper, allowRepeat bool ) {

}


func Generate( length, numberDigits, numberSymbols int, allowUpper, allowRepeat bool) {
	gen, err := NewGenerator(nil)
	if err != nil {
		return "", err
	}
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

	if lowerLetters == ""{
		g.LowerLetters = LowerLetters
	}

}


Зачем тут ридер?? к чему оно?