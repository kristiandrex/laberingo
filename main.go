package main

import "github.com/kristiandrex/laberingo/util"

func main() {
	laberingo := util.Cargar("https://gist.githubusercontent.com/kristiandrex/eb5d0303e52b2c5c35040df135647275/raw/a268a0606cfecb0f8572b289619a643ef2368402/laberingo.txt")
	laberingo.Iniciar()
}
