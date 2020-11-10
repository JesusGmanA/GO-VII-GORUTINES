package main

import (
	"fmt"
	"os"
	"os/exec"

	"./proceso"
)

const AGREGAR_PROC = 1
const MOSTRAR_PROC = 2
const ELIMINAR_PROC = 3
const EXIT = 4

func main() {
	lp := proceso.ListaProcesos{Procesos: map[uint64]proceso.Proceso{}, ContinueRunning: map[uint64]bool{}}
	var opc, id uint64
	for opc != EXIT {
		limpiarPantalla()
		opc = obtenerOpcMenu(&lp, &id)
	}
}

func limpiarPantalla() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func obtenerOpcMenu(lp *proceso.ListaProcesos, id *uint64) uint64 {
	fmt.Println("1) Agregar proceso")
	fmt.Println("2) Mostrar procesos")
	fmt.Println("3) Eliminar proceso")
	fmt.Println("4) Salir")
	fmt.Print("Selecciona una opcion: ")
	var opc uint64
	fmt.Scanln(&opc)
	switch opc {
	case AGREGAR_PROC:
		p := proceso.Proceso{Id: *id, Value: uint64(1)}
		lp.AgregarProceso(p)
		*id++
	case MOSTRAR_PROC:
		proceso.CambiarDisplayProc()
	case ELIMINAR_PROC:
		var idDel uint64
		fmt.Print("Dame ID a eliminar: ")
		fmt.Scanln(&idDel)
		lp.EliminarProceso(idDel)
	}
	if opc != EXIT {
		fmt.Print("Presiona 'Enter' para continuar...")
		fmt.Scanln() //El primero se come el "Enter" atorado en el buffer cuando se lee algo desde consola
	}
	return opc
}
