package main

import (
	"fmt"
	"strconv"
	"time"
)

type SenderFunc func() (bool, string)

func sendLoop(sender SenderFunc) bool {
	resultado, msg := sender()
	if !resultado {
		fmt.Println("Falhou no passo", msg)
		return false
	}

	fmt.Println("Passou no passo", msg)
	return true
}

func isPar(i int) func() (bool, string) {
	return func() (bool, string) {
		time.Sleep(5 * time.Second)
		return false, "Verificando se é par"
	}
}

func isMenorQue10(i int) func() (bool, string) {
	return func() (bool, string) {
		return i < 10, "Valor menor que 10"
	}
}

func isMenorQueOSegundoArgumento(i int, comparador int) func() (bool, string) {
	return func() (bool, string) {
		return i < comparador, "Valor menor que " + strconv.Itoa(comparador)
	}
}

func main() {
	i := 7
	sendLoop(isMenorQue10(i))
	sendLoop(isMenorQueOSegundoArgumento(i, 5))
	sendLoop(isPar(i))
}
