package proceso

import (
	"fmt"
	"time"
)

var display bool = false

type ListaProcesos struct {
	Procesos        map[uint64]Proceso
	ContinueRunning map[uint64]bool
}

type Proceso struct {
	Id    uint64
	Value uint64
}

func (lp *ListaProcesos) AgregarProceso(p Proceso) {
	lp.Procesos[p.Id] = p
	lp.ContinueRunning[p.Id] = true
	go lp.StartProceso(p.Id)
}

func (lp *ListaProcesos) EliminarProceso(id uint64) {
	if lp.existeIdProceso(id) {
		lp.StopProceso(id)

		fmt.Printf("Proceso #%d fue detenido y eliminado\n", id)
	} else {
		fmt.Printf("Proceso #%d no encontrado\n", id)
	}

}

func CambiarDisplayProc() {
	display = !display
}

func (lp *ListaProcesos) StartProceso(pId uint64) {
	p := lp.getProceso(pId)
	for {
		if lp.ContinueRunning[p.Id] {
			p.Value++
			if display {
				fmt.Printf("ID Proc: %d Contador: %d\n", p.Id, p.Value)
			}
			time.Sleep(time.Millisecond * 500)
		} else {
			delete(lp.Procesos, p.Id)
			delete(lp.ContinueRunning, p.Id)
			break
		}
	}
}

func (lp *ListaProcesos) getProceso(pId uint64) *Proceso {
	for k, p := range lp.Procesos {
		if k == pId {
			return &p
		}
	}
	return nil
}

func (lp *ListaProcesos) StopProceso(pid uint64) {
	lp.ContinueRunning[pid] = false

}

func (lp *ListaProcesos) existeIdProceso(id uint64) bool {
	_, found := lp.Procesos[id]
	if found {
		return true
	}
	return false
}
