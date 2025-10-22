package main

import (
	"fmt"
	"strings"
)

func num_to_words(num int) string {
	units := []string{"", "один", "два", "три", "четыре", "пять", "шесть", "семь", "восемь", "девять",
		"десять", "одиннадцать", "двенадцать", "тринадцать", "четырнадцать", "пятнадцать",
		"шестнадцать", "семнадцать", "восемнадцать", "девятнадцать"}
	tens := []string{"", "", "двадцать", "тридцать", "сорок", "пятьдесят", "шестьдесят", "семьдесят", "восемьдесят", "девяносто"}
	hundreds := []string{"", "сто", "двести", "триста", "четыреста", "пятьсот", "шестьсот", "семьсот", "восемьсот", "девятьсот"}

	if num == 0 {
		return "ноль"
	}

	parts := make([]string, 0)
	add_part := func(part string) {
		if part != "" {
			parts = append(parts, strings.TrimSpace(part))
		}
	}

	thousands_switch := func(num int) {
		switch num {
		case 1:
			add_part("одна тысяча")
		case 2:
			add_part("две тысячи")
		case 3, 4:
			add_part(units[num])
			add_part("тысячи")
		default:
			add_part(units[num])
			add_part("тысяч")
		}
	}

	add_thousands := func(num int) {
		t_hund := num / 100
		t_rem := num % 100
		if t_hund > 0 {
			add_part(hundreds[t_hund])
		}
		if t_rem < 20 {
			thousands_switch(t_rem)
		} else {
			t_tens := t_rem / 10
			t_units := t_rem % 10
			add_part(tens[t_tens])
			thousands_switch(t_units)
		}
	}

	thousand := num / 1000
	remainder := num % 1000
	if thousand > 0 {
		if thousand > 999 {
			return "number too large (>999999)"
		}
		add_thousands(thousand)
	}

	hundred := remainder / 100
	remainder %= 100
	if hundred > 0 {
		add_part(hundreds[hundred])
	}

	if remainder < 20 {
		add_part(units[remainder])
	} else {
		ten := remainder / 10
		unit := remainder % 10
		add_part(tens[ten])
		add_part(units[unit])
	}

	return strings.Join(parts, " ")
}

func cycle(num, count *int, done *bool) {
	for *num < 12307 {
		*count++
		if *num < 0 {
			*num *= -1
		} else if *num%7 == 0 {
			*num *= 39
		} else if *num%9 == 0 {
			*num = *num*13 + 1
			continue
		} else {
			*num = (*num + 2) * 3
		}
		if *num%13 == 0 && *num%9 == 0 {
			fmt.Println("service error")
			*done = true
			break
		} else {
			*num++
		}
	}
}

func main() {
	var num, count int
	var done bool
	for !done {
		fmt.Print("Input number: ")
		fmt.Scan(&num)
		if num >= 12307 {
			fmt.Println("Number should be (n < 12307)")
		} else {
			cycle(&num, &count, &done)
			if !done {
				text := num_to_words(num)
				fmt.Println("Result num:", num)
				fmt.Println("Result text:", text)
				done = true
			}
			fmt.Println("Iteration:", count)
		}
	}
}
