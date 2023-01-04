package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Введите математическое выражение")
	vvod := bufio.NewScanner(os.Stdin) // Сканирование строки и перемещение ее в буфер
	vvod.Scan()
	line := vvod.Text() // Запись строки из буфера в переменную
	razbit := strings.Split(line, " ")
	if strings.Contains(razbit[0], "-") || strings.Contains(razbit[2], "-") { //Проверка на наличие в введенной строке знака "-" у первого и второго операнда
		fmt.Println("Нельзя использовать отрицательные значения")
		return
	}

	first_operand := razbit[0]
	second_operand := razbit[2]
	operation := razbit[1]

	// Перевод строчного типа в целочисленный первого операнда
	var first_operand_int int
	var flag_eror int
	flag_eror = 0
	switch first_operand {
	case "I":
		first_operand_int = 1
		flag_eror = 1
		_ = flag_eror
	case "II":
		first_operand_int = 2
		flag_eror = 1
		_ = flag_eror
	case "III":
		first_operand_int = 3
		flag_eror = 1
		_ = flag_eror
	case "IV":
		first_operand_int = 4
		flag_eror = 1
		_ = flag_eror
	case "V":
		first_operand_int = 5
		flag_eror = 1
		_ = flag_eror
	case "VI":
		first_operand_int = 6
		flag_eror = 1
		_ = flag_eror
	case "VII":
		first_operand_int = 7
		flag_eror = 1
		_ = flag_eror
	case "VIII":
		first_operand_int = 8
		flag_eror = 1
		_ = flag_eror
	case "IX":
		first_operand_int = 9
		flag_eror = 1
		_ = flag_eror
	case "X":
		first_operand_int = 10
		flag_eror = 1
		_ = flag_eror
	case "1", "2", "3", "4", "5", "6", "7", "8", "9", "10":
		first_operand_int, _ = strconv.Atoi(first_operand)
	default:
		fmt.Println("Первый операнд не удовлетворяет условию") // Проверка того, что операнд имеет не верный формат (>10 или вообще не то что нужно написанно)
		return
	}

	// Перевод строчного типа в целочисленный второго операнда
	var second_operand_int int
	var flag_eror_2 int
	flag_eror_2 = 0
	switch second_operand {
	case "I":
		second_operand_int = 1
		flag_eror_2 = 1
		_ = flag_eror_2
	case "II":
		second_operand_int = 2
		flag_eror_2 = 1
		_ = flag_eror_2
	case "III":
		second_operand_int = 3
		flag_eror_2 = 1
		_ = flag_eror_2
	case "IV":
		second_operand_int = 4
		flag_eror_2 = 1
		_ = flag_eror_2
	case "V":
		second_operand_int = 5
		flag_eror_2 = 1
		_ = flag_eror_2
	case "VI":
		second_operand_int = 6
		flag_eror_2 = 1
		_ = flag_eror_2
	case "VII":
		second_operand_int = 7
		flag_eror_2 = 1
		_ = flag_eror_2
	case "VIII":
		second_operand_int = 8
		flag_eror_2 = 1
		_ = flag_eror_2
	case "IX":
		second_operand_int = 9
		flag_eror_2 = 1
		_ = flag_eror_2
	case "X":
		second_operand_int = 10
		flag_eror_2 = 1
		_ = flag_eror_2
	case "1", "2", "3", "4", "5", "6", "7", "8", "9", "10":
		if flag_eror == 1 {
			fmt.Println("Нельзя одновременно использовать разные системы счисления") //При переводе первого операнда в целочисленное значение и наличие у него флага, программа проводит проверку
			return
		} else {
			second_operand_int, _ = strconv.Atoi(second_operand) //Конвертация строчного значения в целочисленный
		}
	default:
		fmt.Println("Второй операнд не удовлетворяет условию") // Проверка того, что операнд имеет не верный формат (>10 или вообще не то что нужно написанно)
		return
	}

	if (flag_eror == 0) && (flag_eror_2 == 1) { //Повторная проверка на случчай того, что первый операнд - арабская цифра , второй - римская
		fmt.Println("Нельзя одновременно использовать разные системы счисления")
		return
	}

	var rezult int
	switch operation { //Произведение операции над операндами
	case "+":
		rezult = first_operand_int + second_operand_int

	case "-":
		rezult = first_operand_int - second_operand_int
		if (rezult < 0) && (flag_eror == 1) { //Проверка, что римские цифры используются и результат меньше 0
			fmt.Println("В римской системе счисления нет отрицательных чисел")
			return
		}
	case "*":
		rezult = first_operand_int * second_operand_int

	case "/":
		rezult = first_operand_int / second_operand_int

	}

	var hundreds, ten, one int

	if flag_eror == 1 && flag_eror_2 == 1 { //Начало перевода из арабских в римские , если были использованны римские числа
		hundreds = rezult / 100 // Результат разбивает на сотни , десятки и единицы
		if hundreds == 1 {
			ten = 0
		} else {
			ten = rezult / 10
		}
		if hundreds == 1 && ten == 0 {
			one = 0
		} else {
			one = rezult % 10
		}

	}

	var hundredsString string // Если ответ 100 (может получиться в единственно случаи)
	if hundreds == 1 {
		hundredsString = "C"
		_ = hundredsString
	} else {
		hundredsString = ""
	}

	var tenString string // Перевод десятков в римские цифры
	if ten > 0 {
		switch ten {
		case 1:
			tenString = "X"
		case 2:
			tenString = "XX"
		case 3:
			tenString = "XXX"
		case 4:
			tenString = "XL"
		case 5:
			tenString = "L"
		case 6:
			tenString = "LX"
		case 7:
			tenString = "LXX"
		case 8:
			tenString = "LXXX"
		case 9:
			tenString = "XC"
		}
	} else {
		tenString = ""
	}

	var onestring string // Перевод единиц в римские цифры
	switch one {
	case 1:
		onestring = "I"
	case 2:
		onestring = "II"
	case 3:
		onestring = "III"
	case 4:
		onestring = "IV"
	case 5:
		onestring = "V"
	case 6:
		onestring = "VI"
	case 7:
		onestring = "VII"
	case 8:
		onestring = "VIII"
	case 9:
		onestring = "IX"
	default:
		onestring = ""
	}

	if flag_eror == 1 && flag_eror_2 == 1 { // Проверка, что были введены римские цифры, тогда вывод римскими цифрами
		fmt.Println("Ответ:", hundredsString+tenString+onestring)
	} else {
		fmt.Println("Ответ:", rezult) // Если вводили арабские, то выводим арабские
	}
}
