package comandos

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func DoRmdisk(params []string) {

	ParamPath := ""

	// Recorriendo cada elemento del array
	for _, item := range params {
		fmt.Println(item)
		// variables
		intPosIgual := strings.Index(item, "=")
		intLenItem := len(item)

		// datos que me interesan
		// strParametro := strings.Replace(item[0:intPosIgual+1], "-", "", 5)
		strParametro := item[0 : intPosIgual+1]
		strParametro = strings.TrimLeft(strParametro, "–")
		strParametro = strings.TrimLeft(strParametro, "-")
		strParametro = strings.TrimRight(strParametro, "=")
		strParametro = strings.TrimSpace(strParametro)

		strDato := item[intPosIgual+1 : intLenItem]

		// Guardando informacion de cada parametro
		// y verificando que no esten repetidos
		if strings.ToUpper(strParametro) == "PATH" {
			if ParamPath == "" {
				ParamPath = strDato
			} else {
				fmt.Println("")
				fmt.Println("----------------------------------------------------------------")
				fmt.Println("**!PATH repetido")
				fmt.Println("----------------------------------------------------------------")
				fmt.Println("")
			}
		}

		//fmt.Println(intPosIgual, "///", strParametro, "///", strDato)

	}

	// Validaciones que no falte ningun parametro
	if ParamPath == "" {
		fmt.Println("")
		fmt.Println("----------------------------------------------------------------")
		fmt.Println("**!PATH necesario")
		fmt.Println("----------------------------------------------------------------")
		fmt.Println("")

	} else {
		deleteDisk(ParamPath)
	}

}

func deleteDisk(ppath string) {

	comillaDer := strings.TrimRight(ppath, "\"")
	comillaIzq := strings.TrimLeft(comillaDer, "\"")
	cambio_ruta := comillaIzq

	var opcion string
	fmt.Println("")
	fmt.Println("¿Desea eliminar el disco? S/N : ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	opcion = text

	if opcion == "S\n" || opcion == "s\n" {
		err := os.Remove(cambio_ruta)
		if err != nil {
			fmt.Println("")
			fmt.Println("----------------------------------------------------------------")
			fmt.Println("**!!Disco no encontrado")
			fmt.Println("----------------------------------------------------------------")
			fmt.Println("")
		} else {
			fmt.Println("")
			fmt.Println("--Disco eliminado exitosamente")
			fmt.Println("")
		}
	} else if opcion == "N\n" || opcion == "n\n" {
		fmt.Println("")
		fmt.Println("Eliminacion cancelada")
	} else {
		fmt.Println("")
		fmt.Println("----------------------------------------------------------------")
		fmt.Println("**!Ingrese 's' o 'n'")
		fmt.Println("----------------------------------------------------------------")
		fmt.Println("")
	}
}
