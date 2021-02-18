package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

const (
	entrada         = 9
	camino          = 0
	rastro          = 1
	pared           = -1
	minotauroVivo   = 100
	minotauroMuerto = -100
	salida          = 1000
	splitRow        = "\n"
	splitCol        = ";"
)

//Indica los cuatro posibles movimientos que se pueden realizar en el laberinto
var movimientos = []Punto{
	{-1, 0}, //Arriba
	{0, 1},  //Derecha
	{1, 0},  //Abajo
	{0, -1}, //Izquierda
}

//Punto indica una posición dentro del laberinto
type Punto struct {
	fila int
	col  int
}

//Laberingo estructura del laberinto
type Laberingo struct {
	mapa    [][]int
	entrada Punto
	caminos [][]Punto
}

func main() {
	var url string
	fmt.Print("Ingresa la URL: ")
	fmt.Scanln(&url)

	l := Cargar(url)
	camino := []Punto{l.entrada}

	fmt.Println("Recorriendo...")
	l.mover(l.entrada, camino, true)
	l.hallarCorto()
}

// Cargar desde una URL
func Cargar(url string) Laberingo {
	fmt.Println("Cargando...")

	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	archivo, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	plano := string(archivo)

	filas := strings.Split(plano, splitRow)
	mapa := make([][]int, len(filas))

	laberingo := new(Laberingo)

	for fila := range filas {
		columnas := strings.Split(filas[fila], splitCol)
		mapa[fila] = make([]int, len(columnas))

		for col := range columnas {
			numero, err := strconv.Atoi(columnas[col])

			if err != nil {
				panic(err)
			}

			if numero == entrada {
				laberingo.entrada = Punto{fila, col}
			}

			mapa[fila][col] = numero
		}
	}

	laberingo.mapa = mapa

	fmt.Println(plano)
	return *laberingo
}

// Mover a cada una de las distintas posibilidades en el laberinto
func (l *Laberingo) mover(punto Punto, recorrido []Punto, vivo bool) {
	coordenada := &l.mapa[punto.fila][punto.col]

	if *coordenada == salida && !vivo {
		l.caminos = append(l.caminos, recorrido)
		return
	}

	if *coordenada == camino {
		*coordenada = rastro
	}

	if *coordenada == minotauroVivo {
		vivo = false
		*coordenada = minotauroMuerto
	}

	for _, m := range movimientos {
		siguiente := Punto{punto.fila + m.fila, punto.col + m.col}

		if l.esCamino(siguiente, vivo) {
			l.paso(siguiente, recorrido, vivo)
		}
	}
}

// Dar un paso a una coordenada en específico
func (l *Laberingo) paso(punto Punto, recorrido []Punto, vivo bool) {
	recorrido = append(recorrido, punto)
	l.mover(punto, recorrido, vivo)

	//Haciendo "backtracking"
	coordenada := &l.mapa[punto.fila][punto.col]

	if *coordenada == rastro {
		*coordenada = camino
	}

	if *coordenada == minotauroMuerto {
		*coordenada = minotauroVivo
	}
}

// Comprobar si en el punto hay un camino válido
func (l *Laberingo) esCamino(punto Punto, vivo bool) bool {
	if punto.fila < 0 || punto.fila >= len(l.mapa) {
		return false
	}

	if punto.col < 0 || punto.col >= len(l.mapa[punto.fila]) {
		return false
	}

	numero := l.mapa[punto.fila][punto.col]

	return numero == camino ||
		numero == minotauroVivo ||
		(numero == salida && !vivo)
}

func (l *Laberingo) hallarCorto() {

	if len(l.caminos) == 0 {
		fmt.Println("No se encontró un camino")
		return
	}

	menor := l.caminos[0]

	for i := 0; i < len(l.caminos); i++ {
		c := l.caminos[i]

		if len(c) < len(menor) {
			menor = c
		}
	}

	fmt.Println("El camino más corto es: ", menor)
}
