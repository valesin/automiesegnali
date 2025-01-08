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

// [[file:main.org::*Richiamo][Richiamo:1]]
// Calcolo della distanza di Manhattan
func distanza(a, b Punto) int {
	return abs(b.X-a.X) + abs(b.Y-b.Y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// isPrefisso controlla se α è un prefisso di η
func isPrefisso(alpha, eta string) bool {
	if len(alpha) > len(eta) {
		return false
	}
	return eta[:len(alpha)] == alpha
}

// Richiamo emette un segnale dal punto (x,y)
func (p *Piano) Richiamo(x, y int, alpha string) {
	sorgente := Punto{X: x, Y: y}

	// Se la sorgente è in un ostacolo, nessun movimento possibile
	if p.isPuntoInOstacolo(sorgente) {
		return
	}

	// Trova gli automi rispondenti e la distanza minima
	automiRispondenti := make([]*Automa, 0)
	minDist := -1

	for _, aut := range p.automa {
		if isPrefisso(alpha, aut.Nome) {
			dist := distanza(sorgente, aut.Posizione)
			if minDist == -1 || dist < minDist {
				minDist = dist
				automiRispondenti = automiRispondenti[:0]
				automiRispondenti = append(automiRispondenti, aut)
			} else if dist == minDist {
				automiRispondenti = append(automiRispondenti, aut)
			}
		}
	}

	// Muovi gli automi qualificati
	for _, aut := range automiRispondenti {
		if p.esistePercorsoLibero(aut.Posizione, sorgente) {
			aut.Posizione = sorgente
		}
	}
}
#+end_src go

// Posizioni stampa le posizioni degli automi che corrispondono al prefisso
func (p *Piano) Posizioni(alpha string) {
	// Ottieni automi corrispondenti ordinati per nome
	corrispondenti := make([]string, 0)
	for nome := range p.automa {
		if isPrefisso(alpha, nome) {
			corrispondenti = append(corrispondenti, nome)
		}
	}
	sort.Strings(corrispondenti)

	// Stampa posizioni
	for _, nome := range corrispondenti {
		aut := p.automa[nome]
		fmt.Printf("%s %d %d\n", nome, aut.Posizione.X, aut.Posizione.Y)
	}
}

// EsistePercorso controlla se esiste un percorso libero
func (p *Piano) EsistePercorso(x, y int, nome string) {
	dest := Punto{X: x, Y: y}

	// Controlla se l'automa esiste e se la destinazione è valida
	aut, esiste := p.automa[nome]
	if !esiste || p.isPuntoInOstacolo(dest) {
		fmt.Println("NO")
		return
	}

	if p.esistePercorsoLibero(aut.Posizione, dest) {
		fmt.Println("SI")
	} else {
		fmt.Println("NO")
	}
}

// esistePercorsoLibero controlla se esiste un percorso libero di lunghezza minima
func (p *Piano) esistePercorsoLibero(da, a Punto) bool {
	// Implementazione dell'algoritmo di ricerca del percorso
	// Questa è una versione semplificata che controlla se un punto
	// nel percorso di distanza minima interseca con gli ostacoli

	// Ottieni la direzione del movimento
	dx := segno(a.X - da.X)
	dy := segno(a.Y - da.Y)

	corrente := da

	// Prima tenta il movimento orizzontale
	for corrente.X != a.X {
		corrente.X += dx
		if p.isPuntoInOstacolo(corrente) {
			return false
		}
	}

	// Poi il movimento verticale
	for corrente.Y != a.Y {
		corrente.Y += dy
		if p.isPuntoInOstacolo(corrente) {
			return false
		}
	}

	return true
}

func segno(x int) int {
	if x < 0 {
		return -1
	}
	if x > 0 {
		return 1
	}
	return 0
}
// Richiamo:1 ends here
