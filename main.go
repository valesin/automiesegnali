// [[file:main.org::*Inizio][Inizio:1]]
package automiesegnali


import (
	"fmt"
	"sort"
)
// Inizio:1 ends here

// [[file:main.org::*Definizioni][Definizioni:1]]
// punto rappresenta una posizione nella griglia 2D
type punto struct {
	X, Y int
}

// automa rappresenta un singolo automa
type automa struct {
	Nome      string
	Posizione punto
}

// ostacolo rappresenta un ostacolo rettangolare
type ostacolo struct {
	AngoloInferioreSinistro punto
	AngoloSuperioreDestro   punto
}

// piano rappresenta l'intero sistema
type piano struct {
	automa   map[string]*automa
	ostacoli []*ostacolo
}
// Definizioni:1 ends here

// [[file:main.org::*Utilities][Utilities:1]]

// newPiano crea un piano vuoto
func Crea() *piano {
	return &piano{
		automa:   make(map[string]*automa),
		ostacoli: make([]*ostacolo, 0),
	}
}

// isPuntoInOstacolo controlla se un punto è dentro un ostacolo
func (p *piano) isPuntoInOstacolo(punto punto) bool {
	for _, ost := range p.ostacoli {
		if punto.X >= ost.AngoloInferioreSinistro.X && punto.X <= ost.AngoloSuperioreDestro.X &&
			punto.Y >= ost.AngoloInferioreSinistro.Y && punto.Y <= ost.AngoloSuperioreDestro.Y {
			return true
		}
	}
	return false
}

// isPuntoUnAutoma checks if a point is an automaton
func (p *piano) isPuntoUnAutoma(punto punto) bool {
    for _, aut := range p.automati {
        if aut.Posizione == punto {
            return true
        }
    }
    return false
}

// isAutomaInOstacolo checks if there is any automaton within the given obstacle
func (p *piano) isOstacoloSuPunto(ostacolo ostacolo) bool {
    for _, aut := range p.automati {
        if aut.Posizione.X >= ostacolo.BottomLeft.X && aut.Posizione.X <= ostacolo.TopRight.X &&
            aut.Posizione.Y >= ostacolo.BottomLeft.Y && aut.Posizione.Y <= ostacolo.TopRight.Y {
            return true
        }
    }
    return false
}
// Utilities:1 ends here

// [[file:main.org::*Stato][Stato:1]]
// Stato stampa cosa c'è nella posizione data
func (p *piano) Stato(x, y int) {
	punto := punto{X: x, Y: y}

	// Controlla se il punto è in un ostacolo
	if p.isPuntoInOstacolo(punto) {
		fmt.Println("O")
		return
	}

	// Controlla se il punto contiene un automa
	if p.isPuntoUnAutoma(punto) {
		fmt.Println("A")
		return
	}

	// Se nessuna delle condizioni precedenti è soddisfatta, il punto è vuoto
	fmt.Println("E")
}
// Stato:1 ends here

// [[file:main.org::*Stampa][Stampa:1]]
// Stampa stampa le liste degli automi e degli ostacoli
func (p *piano) Stampa() {
	// Ottieni l'elenco ordinato dei nomi degli automi per output coerente
	nomi := make([]string, 0, len(p.automa))
	for nome := range p.automa {
		nomi = append(nomi, nome)
	}
	sort.Strings(nomi)

	// Stampa automi
	for _, nome := range nomi {
		a := p.automa[nome]
		fmt.Printf("automa %d %d %s\n", a.Posizione.X, a.Posizione.Y, a.Nome)
	}

	// Stampa ostacoli
	for _, ost := range p.ostacoli {
		fmt.Printf("ostacolo %d %d %d %d\n",
			ost.AngoloInferioreSinistro.X, ost.AngoloInferioreSinistro.Y,
			ost.AngoloSuperioreDestro.X, ost.AngoloSuperioreDestro.Y)
	}
}
// Stampa:1 ends here

// [[file:main.org::*Automa][Automa:1]]
// Automa aggiunge o sposta un automa
func (p *piano) Automa(x, y int, nome string) {
	punto := punto{X: x, Y: y}

	// Controlla se il punto è in un ostacolo
	if p.isPuntoInOstacolo(punto) {
		return
	}

	// Crea o sposta automa
	p.automa[nome] = &automa{
		Nome:      nome,
		Posizione: punto,
	}
}
// Automa:1 ends here

// [[file:main.org::*Ostacolo][Ostacolo:1]]
// Ostacolo aggiunge un ostacolo se possibile
func (p *piano) Ostacolo(x0, y0, x1, y1 int) {
	ostacolo := ostacolo{
		AngoloInferioreSinistro: punto{X: x0, Y: y0}
		AngoloSuperioreDestro: punto{X: x1, Y: y1}
		}
	// Controlla se un automa è nell'area proposta per l'ostacolo
	p.isOstacoloSuPunto(ostacolo) {
		return
	}

	// Aggiungi ostacolo
	p.ostacoli = append(p.ostacoli, &Ostacolo{
		AngoloInferioreSinistro: Punto{X: x0, Y: y0},
		AngoloSuperioreDestro:   Punto{X: x1, Y: y1},
	})
}
// Ostacolo:1 ends here
