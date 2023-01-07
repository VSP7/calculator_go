package main // имя текущего файла

/* ТЗ:
Создай консольное приложение «Калькулятор». Приложение должно читать из консоли введенные пользователем строки, числа, арифметические операции, проводимые между ними,
и выводить в консоль результат их выполнения.
Калькулятор можно реализовать обычными функциями либо использовать структуру с методами, здесь это не принципиально.
Требования:

    Калькулятор умеет выполнять операции сложения, вычитания, умножения и деления с двумя числами: a + b, a - b, a * b, a / b.
	Данные передаются в одну строку (смотри пример ниже). Решения, в которых каждое число и арифметическая операция передаются с новой строки, считаются неверными.
    Калькулятор умеет работать как с арабскими (1,2,3,4,5…), так и с римскими (I,II,III,IV,V…) числами.
    Калькулятор должен принимать на вход числа от 1 до 10 включительно, не более. На выходе числа не ограничиваются по величине и могут быть любыми.
    Калькулятор умеет работать только с целыми числами.
    Калькулятор умеет работать только с арабскими или римскими цифрами одновременно, при вводе пользователем строки вроде 3 + II калькулятор должен указать на ошибку
	и прекратить работу.
    При вводе римских чисел ответ должен быть выведен римскими цифрами, соответственно, при вводе арабских — ответ ожидается арабскими.
    При вводе пользователем не подходящих чисел приложение выводит ошибку в терминал и завершает работу.
    При вводе пользователем строки, не соответствующей одной из вышеописанных арифметических операций, приложение выводит ошибку и завершает работу.
    Результатом операции деления является целое число, остаток отбрасывается.
    Результатом работы калькулятора с арабскими числами могут быть отрицательные числа и ноль. Результатом работы калькулятора с римскими числами
	могут быть только положительные числа, если результат работы меньше единицы, программа должна указать на исключение.

Пример работы программы:

Input:
1 + 2

Output:
3

Input:
VI / III

Output:
II

Input:
I - II

Output:
Вывод ошибки, так как в римской системе нет отрицательных чисел.

Input:
I + 1

Output:
Вывод ошибки, так как используются одновременно разные системы счисления.

Input:
1

Output:
Вывод ошибки, так как строка не является математической операцией.

Input:
1 + 2 + 3

Output:
Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).
*/


import ( // импорт нескольких пакетов
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var numbers = map[int]string{
	1: "I",
	2: "II",
	3: "III",
	4: "IV",
	5: "V",
	6: "VI",
	7: "VII",
	8: "VIII",
	9: "IX",
	10: "X",
	20: "XX",
	30: "XXX",
	40: "XL",
	50: "L",
	60: "LX",
	70: "LXX",
	80: "LXXX",
	90: "XC",
	100: "C",
}

func main() { // главная функция запуска точка входа в программу

	var number1, number2, operator string // создадим переменные для принимаемых чисел
	var number11, number22 int

	// доступные операторы
	operators := [4]string{"+", "-", "*", "/"} // массив

    // получим данные от юзера
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    str := scanner.Text() // запишем в строку

	// удалим случайные пробелы вначале и вконце строки
	str=strings.TrimSpace(str)

	number1_Ok := false // флаг того, что первое число по символам закончилось

    for i := 0; i < len(str); i++ { // переберём строку посимвольно и запишем числа в переменные
		if string(str[i])!=" " && !stringInSlice(string(str[i]), operators) && !number1_Ok{ // если символ не равен пробелу и не равен оператору и первое число ещё не заполнено
			number1+=string(str[i])
		}else if (string(str[i])==" " || stringInSlice(string(str[i]), operators)) && !number1_Ok{ // если символ равен пробелу или оператору и флаг заполнения первого числа false
			number1_Ok = true
			if string(str[i])!=" "{ operator=string(str[i]) } // если символ не равен пробелу, значит это оператор
		}else if stringInSlice(string(str[i]), operators) && number1_Ok && operator==""{ // если символ = оператору из массива и первое число заполнено, а оператор нет
			operator=string(str[i])
		}else if string(str[i])!=" " && number1_Ok{ // если символ не равен пробелу и первое число с оператором заполнены
			number2+=string(str[i])
		}else if string(str[i])==" " && number1_Ok && number2!=""{ // если символ равен пробелу и флаг заполнения первого числа true и второе число заполнено
			fmt.Println("Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
			return
		}
    }

	// проверяем, что оператор существует:
	if operator==""{
		fmt.Println("Вывод ошибки, так как строка не является математической операцией.")
		return
	}

	// флаг того, что 1ая цифра римская
	rome1:=false
	// флаг того, что 2ая цифра римская
	rome2:=false

	// конвертируем числа в интовый тип из стрингового
	number11,err := strconv.Atoi(number1) // 1 число интового типа
	if err != nil { // если получаем ошибку, то ищем римскую цифру в мапе и переводим в арабские
		if toArabic(number1)>0{
			number11 = toArabic(number1)
			rome1 = true // ставим флаг, что 1ая цифра римская
		}else{ // если в мапе нет данной цифры, то выводим ошибку
			fmt.Println("Вывод ошибки, так как разрешено использовать только целые числа от 1 до 10 включительно.")
			return
		}
	}

	number22,err = strconv.Atoi(number2)  // 2 число интового типа
	if err != nil { // если получаем ошибку, то ищем римскую цифру в мапе и переводим в арабскую
		if toArabic(number2)>0{
			number22 = toArabic(number2)
			rome2 = true // ставим флаг, что 1ая цифра римская
		}else{ // если в мапе нет данной цифры, то выводим ошибку
			fmt.Println("Вывод ошибки, так как разрешено использовать только целые числа от 1 до 10 включительно.")
			return
		}
	}

	if rome1!=rome2{ // если системы двух чисел не совпадают, то выводим ошибку
		fmt.Println("Вывод ошибки, так как используются одновременно разные системы счисления.")
		return
	}

	// проверяем, что числа в диапазоне от 1 до 10, иначе выводим ошибку
	if number11<1 || number11>10 || number22<1 || number22>10{
		fmt.Println("Вывод ошибки, так как разрешено использовать только целые числа от 1 до 10 включительно.")
		return
	}

	//определим оператор, выполним действие и выведем результат
	if operator=="+"{
		if rome1 && rome2{
			fmt.Println(toRome(number11+number22))
		}else{
			fmt.Println(number11+number22)
		}
	}else if operator=="-"{
		if rome1 && rome2 && number11-number22<1{
			fmt.Println("Вывод ошибки, так как в римской системе только положительные числа.")
		}else if rome1 && rome2 && number11-number22>0{
			fmt.Println(toRome(number11-number22))
		}else{
			fmt.Println(number11-number22)
		}
	}else if operator=="/"{
		if rome1 && rome2{
			if number11/number22>0{
				fmt.Println(toRome(number11/number22))
			}else{
				fmt.Println("Вывод ошибки, так как в римской системе только положительные числа.")
			}
		}else{
			fmt.Println(number11/number22)
		}
	}else if operator=="*"{
		if rome1 && rome2{
			fmt.Println(toRome(number11*number22))
		}else{
			fmt.Println(number11*number22)
		}
	}
}

// функция для конвертации арабских в римские
func toRome(arab int)string{
	roman := ""
	if val, ok := numbers[arab]; ok{
		roman = val
	}else{
		ten:=arab/10*10
		if val, ok := numbers[ten]; ok{
			roman += val
		}
		remains:=arab%10
		if remains>0{
			if val, ok := numbers[remains]; ok{
				roman += val
			}
		}
	}
	return roman
}

// функция для конвертации римских в арабские
func toArabic(roman string)int{
	arab:=0
	for key, value := range numbers {
		if value==roman{
			arab=key
		}
	}

	return arab
}

// функция для проверки наличия элемента в массиве
func stringInSlice(a string, arr [4]string) bool {
    for _, b := range arr {
        if b == a {
            return true
        }
    }
    return false
}

// калькулятор написан нейронкой:
// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// // map that maps Roman numerals to their corresponding integer values
// var romanNumerals = map[string]int{
// 	"I": 1,
// 	"V": 5,
// 	"X": 10,
// 	"L": 50,
// 	"C": 100,
// 	"D": 500,
// 	"M": 1000,
// }

// func main() {
// 	fmt.Println("Welcome to the calculator!")
// 	fmt.Println("Enter an operator (+, -, *, /) and two operands: ")
// 	var operator string
// 	var operand1Str string
// 	var operand2Str string
// 	fmt.Scanln(&operator, &operand1Str, &operand2Str)

// 	// convert operand1Str and operand2Str to float64 if they are not Roman numerals
// 	operand1 := 0.0
// 	if _, ok := romanNumerals[operand1Str]; !ok {
// 		operand1, _ = strconv.ParseFloat(operand1Str, 64)
// 	} else {
// 		operand1 = float64(romanNumerals[operand1Str])
// 	}
// 	operand2 := 0.0
// 	if _, ok := romanNumerals[operand2Str]; !ok {
// 		operand2, _ = strconv.ParseFloat(operand2Str, 64)
// 	} else {
// 		operand2 = float64(romanNumerals[operand2Str])
// 	}

// 	result := 0.0
// 	switch operator {
// 	case "+":
// 		result = operand1 + operand2
// 	case "-":
// 		result = operand1 - operand2
// 	case "*":
// 		result = operand1 * operand2
// 	case "/":
// 		result = operand1 / operand2
// 	default:
// 		fmt.Println("Invalid operator")
// 	}

// 	fmt.Println("Result: " + strconv.FormatFloat(result, 'f', 2, 64))
// }