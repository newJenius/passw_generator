



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
	LowerLetters string
	UpperLetters string
	Digits string
	Symbols string
	Reader io.Reader
}

func NewGenerator(i *GeneratorInput){
	if i == nil {
		i = new(GeneratorInput)
	}
}