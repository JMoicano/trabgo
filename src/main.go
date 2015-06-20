package main

import(
	"./revista"
)

func main() {
	e := CriarEdicao(1, 2, time.Now())

	fmt.Println(e)
}