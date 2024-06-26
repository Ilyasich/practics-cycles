package syracuse

import (
	"fmt"
)

/*Гипотеза Сиракуз:
    Гипотеза гласит, что если взять любое число и применить к нему алгоритм ниже, то через несколько шагов получиться `1`.
    Алгоритм:
    - Если число четное, делим на `2`
    - Если не четное, умножаем на `3` и прибавляем `1`, затем делим на `2`
    - Теперь снова делаем все сначала, но уже с полученным числом
    **Задача**: Написать функцию принимающую входное число и выводящую на экран все шаги алгоритма пока не получится `1`.
    **Задача 2**: Проверить работу функции для чисел от `20` до `30` */

func Syracuse(n int) {
	fmt.Println(n)
	if n == 1 {
		return
	}
	if n%2 == 0 {
		Syracuse(n / 2)
	} else {
		Syracuse((n*3 + 1) / 2)
	}
}

func Syr(n int) {
	for i := 20; i <= 30; i++ {
		fmt.Printf("Syracuse sequence for %d:\n", i)
		Syracuse(i)
		fmt.Println()
	}
}
