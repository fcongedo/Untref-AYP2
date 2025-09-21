package main

import "untref-ayp2/guia-recursividad-division-conquista/hanoi/hanoi"

func main() {
	n := 5
	th := hanoi.NewTorresDeHanoi(n)
	th.Mostrar()
	hanoi.ResolverTorresDeHanoi(th, n, hanoi.Torres[0], hanoi.Torres[1], hanoi.Torres[2])
}
