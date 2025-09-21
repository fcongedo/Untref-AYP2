package hanoi

import "fmt"

type torre struct {
	discos []int
}

func newTorre(n int) *torre {
	t := &torre{discos: []int{}}
	for i := range n {
		t.poner(n - i)
	}
	return t
}

// poner agrega un disco a la torre.
// Si el disco es más grande que el disco en la parte superior de la torre, lanza un error.
func (t *torre) poner(disco int) {
	if len(t.discos) > 0 && disco > t.discos[len(t.discos)-1] {
		panic("No se puede poner un disco más grande sobre uno más pequeño")
	}
	t.discos = append(t.discos, disco)
}

// sacar quita el disco de la parte superior de la torre y lo devuelve.
// Si la torre está vacía, lanza un error.
func (t *torre) sacar() int {
	if len(t.discos) == 0 {
		panic("No hay discos para sacar")
	}
	disco := t.discos[len(t.discos)-1]
	t.discos = t.discos[:len(t.discos)-1]
	return disco
}

// TorresDeHanoi representa el juego de Torres de Hanoi con n discos y tres torres.
// Las torres se llaman "Origen", "Auxiliar" y "Destino".
type TorresDeHanoi struct {
	n      int
	torres map[string]*torre
}

var Torres = [...]string{"Origen", "Auxiliar", "Destino"}

// NewTorresDeHanoi crea un nuevo juego de Torres de Hanoi con n discos
// y las torres "Origen", "Auxiliar" y "Destino".
// Inicialmente, la torre "Origen" tiene n discos y las otras dos torres están vacías.
func NewTorresDeHanoi(n int) *TorresDeHanoi {
	th := &TorresDeHanoi{n: n, torres: make(map[string]*torre)}
	th.torres["Origen"] = newTorre(n)
	th.torres["Auxiliar"] = newTorre(0)
	th.torres["Destino"] = newTorre(0)
	return th
}

// Mover mueve un disco de la torre "desde" a la torre "hasta".
// Si la torre de origen o destino no es válida, lanza un error.
func (th *TorresDeHanoi) Mover(desde string, hasta string) {
	if desde == hasta {
		panic("No se puede mover un disco a la misma torre")
	}
	if _, ok := th.torres[desde]; !ok {
		panic("Torre de origen no válida")
	}
	if _, ok := th.torres[hasta]; !ok {
		panic("Torre de destino no válida")
	}
	disco := th.torres[desde].sacar()
	th.torres[hasta].poner(disco)
	fmt.Printf("Mover disco %d de %s a %s\n", disco, desde, hasta)
	th.Mostrar()
}

// Mostrar imprime el estado actual de las torres y los discos en la consola.
// Cada torre se representa como una columna y los discos se muestran en orden descendente.
func (th *TorresDeHanoi) Mostrar() {
	str := ""
	for i := range th.n {
		str += "    "
		for _, torre := range Torres {
			if len(th.torres[torre].discos) > (th.n - i - 1) {
				str += fmt.Sprintf("[ %v ] ", th.torres[torre].discos[th.n-i-1])
			} else {
				str += "  |   "
			}
		}
		str += "\n"
	}
	str += "  =====================\n"
	fmt.Println(str)
}
