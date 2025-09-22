package divisionyconquista

func BusquedaBinaria(arreglo []int, elemento int) bool {
	// O(log2n) equivalente a O(logn)
	medio := len(arreglo) / 2
	if len(arreglo) == 0 {
		return false
	}
	if arreglo[medio] == elemento {
		return true
	}
	if elemento > arreglo[medio] {
		return BusquedaBinaria(arreglo[medio+1:], elemento)
	} else {
		return BusquedaBinaria(arreglo[:medio], elemento)
	}
}

func BusquedaTernariaRecursiva(arr []int, x int) int {
	return busquedaTernaria(arr, 0, len(arr)-1, x)
}
func busquedaTernaria(arr []int, inicio, fin, x int) int {
	// O(log3n) equivalente a O(logn)
	if inicio > fin {
		return -1
	}
	tercio1 := inicio + (fin-inicio)/3
	tercio2 := inicio + 2*(fin-inicio)/3

	if arr[tercio1] == x {
		return tercio1
	}
	if arr[tercio2] == x {
		return tercio2
	}
	if x < arr[tercio1] {
		return busquedaTernaria(arr, inicio, tercio1-1, x)
	} else if x > arr[tercio2] {
		return busquedaTernaria(arr, tercio2+1, fin, x)
	} else {
		return busquedaTernaria(arr, tercio1+1, tercio2-1, x)
	}
}

func Pico(arreglo []int) int {
	return picoRec(arreglo, 0, len(arreglo)-1)
}
func picoRec(arreglo []int, inicio, fin int) int {
	medio := (inicio + fin) / 2

	if medio == 0 || medio == len(arreglo)-1 ||
		(arreglo[medio] > arreglo[medio-1] && arreglo[medio] > arreglo[medio+1]) {
		return medio
	}

	if arreglo[medio] < arreglo[medio+1] {
		return picoRec(arreglo, medio+1, fin)
	} else {
		return picoRec(arreglo, inicio, medio-1)
	}

}

func SubsecuenciaSumaMaxima(arreglo []int) (int, int, int) {
	// Implementar
	return -1, -1, -1
}

func EsArregloMagico(arreglo []int) bool {
	// Implementar
	return false
}
