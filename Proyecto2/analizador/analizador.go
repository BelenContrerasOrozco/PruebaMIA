package analizador

import (
	"Proyecto2/analizador/comandos"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Start_Analizer() {

	intContador := 0
	for { // simulacion de un while infinito

		EncabezadoInfo()
		// leyendo entrada por consola
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		entrada := scanner.Text()

		arrEntrada := strings.Split(entrada, " ")

		if strings.ToUpper(arrEntrada[0]) == "SALIR" {
			fmt.Println("Si lees esto me debes 10 puntos del proyecto")
			break
		}

		fmt.Println("\n-------------------------------------------")
		Comandos_indentify(arrEntrada[0], arrEntrada, entrada)
		fmt.Println("---------------------------------------------")
		fmt.Println(" ")

		intContador++

	}
}

var lista *comandos.Lista = comandos.New_Lista()

// Identifica cada comando
func Comandos_indentify(command string, params []string, entrada string) {
	// fmt.Println("Comando -> ", command)
	// fmt.Println("**************************************")
	// fmt.Println(params)
	// fmt.Println(params[1])
	// fmt.Println(params[2])

	if strings.ToUpper(command) == "MKDISK" {
		fmt.Println("- - - - - - - - - - -MKDISK- - - - - - - - - - -")
		// enviando los parametros
		//mkdisk.DoMkdisk(params)
		comandos.DoMkdisk(params)

	} else if strings.ToUpper(command) == "RMDISK" {
		fmt.Println("- - - - - - - - - - -RMKDISK- - - - - - - - - - -")
		comandos.DoRmdisk(params)

	} else if strings.ToUpper(command) == "FDISK" {
		fmt.Println("- - - - - - - - - - -FDISK- - - - - - - - - - -")
		comandos.DoFdisk(entrada)
	} else if strings.ToUpper(command) == "MOUNT" {
		fmt.Println("- - - - - - - - - - -MOUNT- - - - - - - - - - -")
		comandos.DoMount(entrada, lista)
	} else if strings.ToUpper(command) == "MKFS" {
		fmt.Println("- - - - - - - - - - -MKFS- - - - - - - - - - -")
	} else if strings.ToUpper(command) == "LOGIN" {
		fmt.Println("- - - - - - - - - - -LOGIN- - - - - - - - - - -")
	} else if strings.ToUpper(command) == "LOGOUT" {
		fmt.Println("- - - - - - - - - - -LOGOUT- - - - - - - - - - -")
	} else if strings.ToUpper(command) == "MKGRP" {
		fmt.Println("- - - - - - - - - - -MKGRP- - - - - - - - - - -")
	} else if strings.ToUpper(command) == "RMGRP" {
		fmt.Println("- - - - - - - - - - -RMGRP- - - - - - - - - - -")
	} else if strings.ToUpper(command) == "MKUSER" {
		fmt.Println("- - - - - - - - - - -MKUSER- - - - - - - - - - -")
	} else if strings.ToUpper(command) == "RMUSR" {
		fmt.Println("- - - - - - - - - - -RMUS- - - - - - - - - - -")
	} else if strings.ToUpper(command) == "MKFILE" {
		fmt.Println("- - - - - - - - - - -MKFILE- - - - - - - - - - -")
	} else if strings.ToUpper(command) == "MKDIR" {
		fmt.Println("- - - - - - - - - - -MKDIR- - - - - - - - - - -")
	} else if strings.ToUpper(command) == "PAUSE" {
		fmt.Println("- - - - - - - - - - -PAUSE- - - - - - - - - - -")
	} else if strings.ToUpper(command) == "EXEC" {
		fmt.Println("- - - - - - - - - - -EXEC- - - - - - - - - - -")
		DoExec(entrada)
	} else if strings.ToUpper(command) == "REP" {
		fmt.Println("- - - - - - - - - - -REP- - - - - - - - - - -")
		comandos.DoRep(entrada, lista)
	}

}

// Muestra informacion personal
func EncabezadoInfo() {
	fmt.Println("")
	fmt.Println("-  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -  -")
	fmt.Println("-----------------------[MIA] Proyecto 2-----------------------")
	fmt.Println("------------------Ana Belén Contreras Orozco------------------")
	fmt.Println("---------------------------201901604--------------------------")
	fmt.Println("")
	fmt.Println("----Ingrese comando: ")
}
