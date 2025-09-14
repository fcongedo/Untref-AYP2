# Guía: Pilas y Colas

## Tipos Abstractos de Datos

1. Implementar `Stack` sobre _slices_ ([stack/stack.go](./stack/stack.go)).
2. Implementar `Queue` sobre _slices_ ([queue/queue.go](./queue/queue.go)).
3. Implementar una cola (`StackQueue`) pero usando internamente dos pilas (`Stack`) para almacenar los datos.
   - Analizar el orden de las operaciones de la cola implementada.

### Usos de pilas y colas

1. Escribir una función `InvertirCadena` que reciba una cadena de caracteres y devuelva la cadena invertida. Analizar el orden.
2. Escribir una función `EsPalindromo` que verifique si una palabra es palíndromo (es decir que una cadena es igual a su inversa. Por ejemplo: las cadenas `"1456541"` y `"145541"` son palíndromos). Analizar el orden.
3. Escribir una función que evalúe si una cadena de paréntesis, corchetes y llaves está bien balanceada y devuelve `true` si está bien balanceada y `false` si está mal balanceada. Por ejemplo: `[()]{}{[()()]()}` debe devolver `true`, mientras que `[(])` debe devolver `false`. Analizar el orden.
4. Escribir una función, tal que, dadas dos colas, construya una cola con el resultado de poner una a continuación de la otra. Por ejemplo: si `q1 = [1, 2, 3]` (con 1 al frente y 3 al final) y `q2 = [5, 7]`, el resultado es `[1, 2, 3, 5, 7]` (con 1 al frente y 7 al final). Analizar el orden.
5. Escribir una función que recibe una cadena de caracteres que representa una expresión en [notación polaca inversa (RPN)](https://es.wikipedia.org/wiki/Notación_polaca_inversa) y devuelve el resultado de la expresión. Por ejemplo: `"2 3 + 5 *"` debe devolver `25`. La expresión se evalúa de izquierda a derecha. Los operadores válidos son `+`, `-`, `*` y `/`. Analizar el orden.

   Tip: se utiliza una pila para ir almacenando los operandos y los resultado.
