package util

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// Laberingo ...
type Laberingo struct {
	coordenadas [][]int
	entrada     [2]int
	caminos     []string
}

// Cargar un laberinto desde una URL
func Cargar(url string) *Laberingo {
	laberingo := Laberingo{}

	fmt.Println("Cargando laberinto...")

	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	archivo, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	data := string(archivo)

	filas := strings.Split(data, "\n")
	coordenadas := make([][]int, len(filas))

	for i := 0; i < len(filas); i++ {
		columnas := strings.Split(filas[i], ";")
		coordenadas[i] = make([]int, len(columnas))

		for j := 0; j < len(columnas); j++ {
			numero, err := strconv.Atoi(columnas[j])

			if err != nil {
				panic(err)
			}

			if numero == 9 {
				laberingo.entrada = [2]int{i, j}
			}

			coordenadas[i][j] = numero
		}
	}

	laberingo.coordenadas = coordenadas

	fmt.Println("laberinto cargado: \n", data)
	return &laberingo
}

// Iniciar el recorrido del laberinto
func (l *Laberingo) Iniciar() {
	fmt.Println("Recorriendo...")
	camino := fmt.Sprint("(", l.entrada[0], ",", l.entrada[1], ")")

	l.recorrer(l.entrada[0], l.entrada[1], camino, true)
	l.hallarCorto()
}

// Recorrer recursivamente las coordenadas del laberinto
func (l *Laberingo) recorrer(fila int, columna int, camino string, vivo bool) {
	coordenada := &l.coordenadas[fila][columna]

	if *coordenada == 1000 && !vivo {
		l.caminos = append(l.caminos, camino)
		return
	}

	if *coordenada == 0 {
		*coordenada = 1
	}

	if *coordenada == 100 {
		vivo = false
		*coordenada = -100
	}

	if l.irArriba(fila, columna, vivo) {
		l.mover(fila-1, columna, camino, vivo)
	}

	if l.irAbajo(fila, columna, vivo) {
		l.mover(fila+1, columna, camino, vivo)
	}

	if l.irDerecha(fila, columna, vivo) {
		l.mover(fila, columna+1, camino, vivo)
	}

	if l.irIzquierda(fila, columna, vivo) {
		l.mover(fila, columna-1, camino, vivo)
	}

}

// Mover a una coordenada en específico
func (l *Laberingo) mover(fila int, columna int, camino string, vivo bool) {
	camino += fmt.Sprint("-(", fila, ",", columna, ")")
	l.recorrer(fila, columna, camino, vivo)

	coordenada := &l.coordenadas[fila][columna]

	if *coordenada == 1 {
		*coordenada = 0
	}

	if *coordenada == -100 {
		*coordenada = 100
	}
}

func (l *Laberingo) irArriba(fila int, columna int, vivo bool) bool {
	return l.esCamino(fila-1, columna, vivo)
}

func (l *Laberingo) irAbajo(fila int, columna int, vivo bool) bool {
	return l.esCamino(fila+1, columna, vivo)
}

func (l *Laberingo) irDerecha(fila int, columna int, vivo bool) bool {
	return l.esCamino(fila, columna+1, vivo)
}

func (l *Laberingo) irIzquierda(fila int, columna int, vivo bool) bool {
	return l.esCamino(fila, columna-1, vivo)
}

// Comprobar si en las coordenadas hay un camino
func (l *Laberingo) esCamino(fila int, columna int, vivo bool) bool {
	if fila < 0 || fila >= len(l.coordenadas) {
		return false
	}

	if columna < 0 || columna >= len(l.coordenadas[fila]) {
		return false
	}

	valor := l.coordenadas[fila][columna]

	return valor == 0 || valor == 100 || valor == 1000 && !vivo
}

// Hallar el camino más corto
func (l *Laberingo) hallarCorto() {
	var menor string

	for _, camino := range l.caminos {
		if len(menor) == 0 || len(camino) < len(menor) {
			menor = camino
		}
	}

	fmt.Println("El camino más corto es: ", menor)
}
