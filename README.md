# Laberingo

Laberinto recursivo en Go.

Halla el camino más corto de un laberinto de "Teseo ⚔ y el Minotauro 🐮" a partir de una matriz cargada desde una URL.

---

## Valores de la matriz

| Número | Descripción           |
| ------ | --------------------- |
| 9      | Entrada del laberinto |
| 0      | Camino sin recorrer   |
| 1      | Camino recorrido      |
| -1     | Pared                 |
| 100    | Minotauro vivo        |
| -100   | Minotauro muerto      |
| 1000   | Salida del laberinto  |

Matriz de ejemplo en este [Gist](https://gist.githubusercontent.com/kristiandrex/eb5d0303e52b2c5c35040df135647275/raw/a268a0606cfecb0f8572b289619a643ef2368402/laberingo.txt
)