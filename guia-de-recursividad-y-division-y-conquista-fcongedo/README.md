# Guía: Recursividad y División y Conquista

## Recursividad

1. Escribir una función recursiva que tome un entero `n` y devuelva la suma de los primeros `n` números.

2. Escribir una función recursiva que tome un entero `n` y devuelva su factorial.

3. Escribir una función recursiva que devuelva la cantidad de unos en la representación binaria de un número `n`. Use el hecho de que es igual al número de unos en la representación binaria de `n/2`, mas 1 si es impar.

4. Escribir una función recursiva que tome una cadena y devuelva como resultado la cadena invertida.

5. Escribir una función recursiva que reciba una cola e invierta todos los elementos de la cola sin utilizar ninguna estructura de datos adicional.

6. Escribir una función recursiva que calcule el máximo común divisor entre dos números enteros.

   Nota: Se puede usar el algoritmo de Euclides:

   $$
   \begin{align*}
   \text{mcd}(a, 0) &= a \\
   \text{mcd}(a, b) &= \text{mcd}(b, a \bmod b)
   \end{align*}
   $$

7. Escribir una función recursiva que tome dos números enteros y calcule la multiplicación entre ellos, usando sólo sumas.

8. Escribir una función recursiva que devuelva el cociente y el resto de la división entera mediante restas sucesivas.

9. Escribir una función recursiva que tome un array de números enteros y devuelva la suma de todos sus elementos.

10. Escribir una función recursiva que dado un número entero decimal devuelva su equivalente en binario.

11. Escribir una función recursiva que reciba 2 enteros `n` y `b` y devuelva true si `n` es potencia de `b`.

    Por ejemplo: `esPotencia(8, 2)` devuelve `true`.

12. Escribir una función recursiva que calcule Fibonacci de `n`.

13. El triángulo de Pascal es un triángulo numérico en el que cada número es la suma de los dos números directamente encima de él. Es decir, el número en la fila `n` y columna `k` es igual a la suma del número en la fila `n-1` y columna `k-1` y el número en la fila `n-1` y columna `k`.

    Escribir una función recursiva que calcule el valor en la fila `n` y columna `k` del triángulo de Pascal. Por ejemplo, el triángulo de Pascal se ve así:

    ```text
            1
          1   1
        1   2   1
      1   3   3   1
    1   4   6   4   1
    ```

    La función debe devolver el valor en la fila `n` y columna `k`. Por ejemplo, `pascal(4, 2)` devuelve `6`.

14. Escribir una función recursiva que resuelva el juego de las [Torres de Hanoi](https://es.wikipedia.org/wiki/Torres_de_Hanói) ([Ver animacion](https://youtu.be/8XQmuPKOgy8?t=38)):

    > El juego, en su forma más tradicional, consiste en tres postes verticales. En uno de los postes se apila un número indeterminado de discos perforados por su centro (elaborados de madera), que determinará la complejidad de la solución. Por regla general se consideran siete discos. Los discos se apilan sobre uno de los postes en tamaño decreciente de abajo arriba. No hay dos discos iguales, y todos ellos están apilados de mayor a menor radio -desde la base del poste hacia arriba- en uno de los postes, quedando los otros dos postes vacíos. El juego consiste en pasar todos los discos desde el poste ocupado (es decir, el que posee la torre) a uno de los otros postes vacíos. Para realizar este objetivo, es necesario seguir tres simples reglas:
    >
    > 1. Solo se puede mover un disco cada vez y para mover otro los demás tienen que estar en postes.
    > 2. Un disco de mayor tamaño no puede estar sobre uno más pequeño que él mismo.
    > 3. Solo se puede desplazar el disco que se encuentre arriba en cada poste.

    En [`hanoi/hanoi/tablero.go`](hanoi/hanoi/tablero.go) se encuentra una implementación que valida que los movimientos que hacen sobre las torres son válidas. En [`hanoi/hanoi/resolver.go`](hanoi/hanoi/resolver.go) deberán implementa de forma recursiva la solución de Torres de Hanoi, moviendo todos los discos desde la torre origen a la torre destino. Para ejecutar el juego, deben ejecutar el siguiente comando:

    ```bash
    go run hanoi/main.go
    ```

    Realizar un movimiento inválido, causará un error irrecuperable. ¡Buena suerte!

## División y Conquista

1. Escribir una función recursiva para determinar si un elemento está en un arreglo utilizando búsqueda binaria. Calcule su eficiencia.

2. Programar la busqueda ternaria recursiva. Donde en lugar de dividir el arreglo en dos partes iguales, se divide en tres partes iguales. Calcular el orden.

3. Se tiene un arreglo de `len(n) >= 3` elementos en forma de pico, esto es: estrictamente creciente hasta una determinada posición `p`, y estrictamente decreciente a partir de ella (con `0 < p < n-1`). Por ejemplo, en el arreglo `[1, 2, 3, 1, 0, -2]` la posición del pico es `p=2`.

   Se pide implementar un algoritmo de División y Conquista de orden $O(\log{n})$ que encuentre la posición p del pico.

4. Resolver el problema de la subsecuencia de suma máxima aplicando la técnica de División y Conquista.

5. Un arreglo es mágico cuando `a[i] = i` para algún `i` entre `0` y `n-1`, siendo `n` el tamaño del arreglo. Los datos están ordenados en forma ascendente y no hay elementos repetidos. Resolver en $O(\log{n})$. Es una variante de la busqueda binaria.
