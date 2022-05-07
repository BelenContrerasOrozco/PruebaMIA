package comandos

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"unsafe"
)

func DoRep(comando string, lista *Lista) {
	sinsalto := strings.TrimRight(comando, "\n")
	coman := strings.Split(sinsalto, " ")
	lista_simple := lista
	var bandera_error bool = false
	var bandera_path bool = false
	var bandera_name bool = false
	var bandera_id bool = false
	var valor_path string = ""
	var valor_name string = ""
	var valor_id string = ""
	var valor_ruta string = ""
	for i := 0; i < len(coman); i++ {
		param := strings.Split(coman[i], "=")
		for j := 0; j < len(param); j++ {
			if param[j] == "-path" || param[j] == "-PATH" {
				bandera_path = true
				valor_path = param[j+1]
			} else if param[j] == "-name" || param[j] == "-NAME" {
				bandera_name = true
				valor_name = param[j+1]
			} else if param[j] == "-id" || param[j] == "-ID" {
				bandera_id = true
				valor_id = param[j+1]
			} else if param[j] == "-ruta" || param[j] == "-RUTA" {
				valor_ruta = param[j+1]
			}
		}
	}
	if !bandera_name {
		bandera_error = true
		fmt.Println("")
		fmt.Println("----------------------------------------------------------------")
		fmt.Println("**!NAME necesario")
		fmt.Println("----------------------------------------------------------------")
		fmt.Println("")
	}
	if !bandera_path {
		bandera_error = true
		fmt.Println("")
		fmt.Println("----------------------------------------------------------------")
		fmt.Println("**!PATH necesario")
		fmt.Println("----------------------------------------------------------------")
		fmt.Println("")
	}
	if !bandera_id {
		bandera_error = true
		fmt.Println("")
		fmt.Println("----------------------------------------------------------------")
		fmt.Println("**!ID necesario")
		fmt.Println("----------------------------------------------------------------")
		fmt.Println("")
	}
	if !bandera_error {
		ejecutar_rep(valor_name, valor_path, valor_id, valor_ruta, lista_simple)
	}
}

func ejecutar_rep(pname string, ppath string, pid string, pruta string, lista *Lista) {
	/*fmt.Println("El valor de pname: ", pname)
	fmt.Println("El valor de ppath: ", ppath)
	fmt.Println("El valor de pid: ", pid)
	fmt.Println("El valor de pruta: ", pruta)*/

	comillaDer := strings.TrimRight(ppath, "\"")
	comillaIzq := strings.TrimLeft(comillaDer, "\"")
	cambio_ruta := comillaIzq

	path := GetDireccion(pid, lista)
	ext := obtener_extension(cambio_ruta)

	if path != "null" {
		crearRuta(cambio_ruta)
		if pname == "disk" {
			graficarDisco(path, cambio_ruta, ext)
		}
	}
}

func graficarDisco(ppaht string, destino string, extension string) {
	fmt.Println("")
	fmt.Println("---path: ", ppaht)
	fmt.Println("---destino: ", destino)
	fmt.Println("---extension: ", extension)
	fmt.Println("")
	var destinoDot string

	for i := 0; i < len(destino); i++ {
		destinoDot += string(destino[i])
		if string(destino[i]) == "." {
			break
		}
	}

	destinoDot += "dot"

	_, err := os.Stat(destinoDot)
	if os.IsNotExist(err) {
		var archivo, err = os.Create(destinoDot)
		if err != nil {
			fmt.Println("")
			fmt.Println("----------------------------------------------------------------")
			fmt.Println("**!Ruta no creada")
			fmt.Println("----------------------------------------------------------------")
			fmt.Println("")
		}

		_, err = archivo.WriteString(generarReporteDisk(ppaht))
		if err != nil {
			fmt.Println("")
			fmt.Println("----------------------------------------------------------------")
			fmt.Println("**!Escritura no lograda")
			fmt.Println("----------------------------------------------------------------")
			fmt.Println("")
		}

		err = archivo.Sync()
		if err != nil {
			return
		}

		defer archivo.Close()
		fmt.Println("")
		fmt.Println("--Archivo creado")
		fmt.Println("")
	}

	comando1 := "-T" + extension

	out, err := exec.Command("dot", comando1, destinoDot, "-o", destino).Output()
	if err != nil {
		fmt.Println("")
		fmt.Println("----------------------------------------------------------------")
		fmt.Println("**!Comando no ejecutado")
		fmt.Println("----------------------------------------------------------------")
		fmt.Println("")
	}

	err = os.Remove(destinoDot)
	if err != nil {
		fmt.Println("")
		fmt.Println("----------------------------------------------------------------")
		fmt.Println("**!Error al eliminar el .dot")
		fmt.Println("----------------------------------------------------------------")
		fmt.Println("")
	}
	fmt.Println(string(out))
}

func generarReporteDisk(ppath string) string {
	file, _ := os.OpenFile(ppath, os.O_RDWR, 0777)
	defer file.Close()

	rama := "digraph G{\n\n"
	rama += "tbl [shape=\"" + "box" + "\"label=<\n"
	rama += "<table border=\"" + "0" + "\" cellborder=\"" + "2" + "\" width=\"" + "600" + "\" height=\"" + "200" + "\" color=\"" + "black" + "\">\n"
	rama += "<tr>\n"
	rama += "<td height=\"" + "200" + "\" width=\"" + "100" + "\"> MBR </td>\n"

	mbr := MBR{}

	//nos posicionamos al inicio del archivo usando la funcion Seek
	file.Seek(0, 0)

	//obtenemor el size del MBR para empezar a leer desde ahi
	var sizeMbr int64 = int64(unsafe.Sizeof(mbr))

	file.Seek(0, 0)
	dataControl := leerBytes(file, int(sizeMbr))
	bufferControl := bytes.NewBuffer(dataControl)
	err := binary.Read(bufferControl, binary.BigEndian, &mbr)
	if err != nil {
		log.Fatal("binary.Read failed", err)
	}

	var total int = int(mbr.Mbr_tamano)
	var espacioUsado float64 = 0

	for i := 0; i < 4; i++ {
		parcial := mbr.Mbr_partition[i].Part_size
		if mbr.Mbr_partition[i].Part_start != -1 {
			var porcentaje_real float64 = (float64(parcial) * 100) / float64(total)
			var porcentaje_aux float64 = (porcentaje_real * 500) / 100
			espacioUsado += porcentaje_real
			if mbr.Mbr_partition[i].Part_status != '1' {
				if mbr.Mbr_partition[i].Part_type == 'P' {
					rama += "<td height=\"" + "200" + "\" width=\"" + fmt.Sprintf("%v", porcentaje_aux) + "\"> PRIMARIA <br/>" + fmt.Sprintf("%v", porcentaje_real) + "</td>\n"
					if i != 3 {
						var p1 int = int(mbr.Mbr_partition[i].Part_start) + int(mbr.Mbr_partition[i].Part_size)
						var p2 int = int(mbr.Mbr_partition[i+1].Part_start)
						if mbr.Mbr_partition[i+1].Part_start != -1 {
							if (p2 - p1) != 0 {
								var fragmentacion int = p2 - p1
								var porcentaje_real float64 = (float64(fragmentacion) * 100) / float64(total)
								var porcentaje_aux float64 = (porcentaje_real * 500) / 100
								rama += "<td height=\"" + "200" + "\" width=\"" + fmt.Sprintf("%v", porcentaje_aux) + "\">LIBRE<br/>" + fmt.Sprintf("%v", porcentaje_real) + "</td>\n"
							}
						}
					} else {
						var p1 int = int(mbr.Mbr_partition[i].Part_start) + int(mbr.Mbr_partition[i].Part_size)
						var mbr_size int = total + int(sizeMbr)
						if (mbr_size - p1) != 0 {
							var libre float64 = (float64(mbr_size) - float64(p1)) + float64(sizeMbr)
							var porcentaje_real float64 = (libre * 100) / float64(total)
							var porcentaje_aux float64 = (porcentaje_real * 500) / 100
							rama += "<td height=\"" + "200" + "\" width=\"" + fmt.Sprintf("%v", porcentaje_aux) + "\">LIBRE<br/>" + fmt.Sprintf("%v", porcentaje_real) + "</td>\n"
						}
					}
				} else {
					ebr := EBR{}
					rama += "<td height=\"" + "200" + "\" width=\"" + fmt.Sprintf("%v", porcentaje_real) + "\">\n <table border=\"" + "0" + "\" height=\"" + "200" + "\" width=\"" + fmt.Sprintf("%v", porcentaje_real) + "\" cellborder=\"" + "1" + "\">\n"
					rama += "<tr><td height=\"" + "60" + "\" colspan=\"" + "15" + "\">EXTENDIDA</td> </tr>\n <tr>\n"
					file.Seek(mbr.Mbr_partition[i].Part_start, 0)
					var sizeEbr int64 = int64(unsafe.Sizeof(ebr))
					dataControl := leerBytes(file, int(sizeEbr))
					bufferControl := bytes.NewBuffer(dataControl)
					err := binary.Read(bufferControl, binary.BigEndian, &ebr)
					if err != nil {
						log.Fatal("binary.Read failed", err)
					}
					if ebr.Part_size != 0 {
						file.Seek(mbr.Mbr_partition[i].Part_start, 0)
						offset, _ := file.Seek(0, os.SEEK_CUR)

						for {
							if sizeEbr != 0 && (offset < (mbr.Mbr_partition[i].Part_size + mbr.Mbr_partition[i].Part_start)) {
								parcial = ebr.Part_size
								fmt.Println("--Valor parcial", parcial)
								porcentaje_real = (float64(parcial) * 100) / float64(total)
								fmt.Println("--Porcentaje real", porcentaje_real)
								if porcentaje_real != 0 {
									fmt.Println("--Valor de ebr.Part_status:", ebr.Part_status)
									if ebr.Part_status != '1' {
										rama += "<td height=\"" + "140" + "\">EBR</td>\n"
										rama += "<td height=\"" + "140" + "\">LOGICA<br/>" + fmt.Sprintf("%v", porcentaje_real) + "</td>\n"
									} else {
										rama += "<td height=\"" + "150" + "\">LIBRE<br/>" + fmt.Sprintf("%v", porcentaje_real) + "</td>\n"
									}
									fmt.Println("--Valor de ebr.Part_next:", ebr.Part_next)
									if ebr.Part_next == -1 {
										parcial = (mbr.Mbr_partition[i].Part_start + mbr.Mbr_partition[i].Part_size) - (ebr.Part_start + ebr.Part_size)
										porcentaje_real = (float64(parcial) * 100) / float64(total)
										fmt.Println("--Valor de porcentaje real", porcentaje_real)
										if porcentaje_real != 0 {
											rama += "<td height=\"" + "150" + "\">LIBRE <br/>" + fmt.Sprintf("%v", porcentaje_real) + "</td>\n"
										}
										break
									} else {
										file.Seek(ebr.Part_next, 0)
									}
								}
							}
						}
					} else {
						rama += "<td height=\"" + "140" + "\"> Ocupado " + fmt.Sprintf("%v", porcentaje_real) + "</td>"
					}
					rama += "</tr>\n </table>\n </td>\n"
					//--Opcional
					if i != 3 {
						var p1 int = int(mbr.Mbr_partition[i].Part_start) + int(mbr.Mbr_partition[i].Part_size)
						var p2 int = int(mbr.Mbr_partition[i+1].Part_start)
						if mbr.Mbr_partition[i+1].Part_start != -1 {
							if (p2 - p1) != 0 {
								var fragmentacion = p2 - p1
								var porcentaje_real float64 = (float64(fragmentacion) * 100) / float64(total)
								var porcentaje_aux float64 = (porcentaje_real * 500) / 100
								rama += "<td height=\"" + "200" + "\" width=\"" + fmt.Sprintf("%v", porcentaje_aux) + "\">LIBRE<br/>" + fmt.Sprintf("%v", porcentaje_real) + "</td>\n"
							}
						}
					} else {
						var p1 int = int(mbr.Mbr_partition[i].Part_start) + int(mbr.Mbr_partition[i].Part_size)
						var mbr_size int = total + int(sizeMbr)
						if (mbr_size - p1) != 0 { //Libre
							var libre float64 = (float64(mbr_size) - float64(p1)) + float64(sizeMbr)
							var porcentaje_real float64 = (libre * 100) / float64(total)
							var porcentaje_aux float64 = (porcentaje_real * 500) / 100
							rama += "<td height=\"" + "200" + "\" width=\"" + fmt.Sprintf("%v", porcentaje_aux) + "\">LIBRE<br/>" + fmt.Sprintf("%v", porcentaje_real) + "</td>\n"
						}
					}
				}
			} else {
				rama += "<td height=\"" + "200" + "\" width=\"" + fmt.Sprintf("%v", porcentaje_aux) + "\">LIBRE <br/>" + fmt.Sprintf("%v", porcentaje_real) + "</td>\n"
			}
		}
	}
	rama += "</tr> \n </table> \n>];\n\n}"

	//En esta parte esta el uso correcto de graphviz
	/*rama := "digraph G{\n\n"
	rama += "node[shape=\"" + "box" + "\",style=\"" + "filled" + "\",fillcolor=\"" + "#EEEEE" + "\",color=\"" + "#EEEEE" + "\"];"
	rama += "	node1" + "1" + "[label=\"" + "1" + "\"];"
	rama += "}\n"*/

	return rama
}

func crearRuta(path string) {
	directorio := obtener_ruta(path)
	if _, err := os.Stat(directorio); os.IsNotExist(err) {
		err = os.MkdirAll(directorio, 0777)
		if err != nil {
			panic(err)
		}
	}
}

func obtener_ruta(path string) string {
	var aux_path int
	var aux_ruta string

	for i := len(path) - 1; i >= 0; i-- {
		aux_path++
		if string(path[i]) == "/" {
			break
		}
	}

	for i := 0; i < ((len(path)) - (aux_path - 1)); i++ {
		aux_ruta += string(path[i])
	}

	return aux_ruta
}

func obtener_extension(path string) string {
	var aux_path int
	var aux_delimitador string

	for i := len(path) - 1; i >= 0; i-- {
		aux_path++
		if string(path[i]) == "." {
			break
		}
	}

	for i := (len(path) - (aux_path - 1)); i <= len(path)-1; i++ {
		aux_delimitador += string(path[i])
	}

	return aux_delimitador
}
