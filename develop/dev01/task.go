package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
	"time"
)

/*
=== Базовая задача ===

Создать программу, печатающую точное время с использованием NTP библиотеки. Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу, печатающую текущее время / точное время с использованием этой библиотеки.

Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и golint.`1
*/
const ntpServerURL = "0.beevik-ntp.pool.ntp.org"

func main() {
	ntpTime, err := ntp.Time(ntpServerURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "NTP error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Current NTP time:", ntpTime.Format(time.RFC850))

}
