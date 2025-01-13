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
	x, y int
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
	automi   map[string]*automa
	ostacoli []*ostacolo
}
// Definizioni:1 ends here

// [[file:main.org::*Stato][Stato:1]]
// Stato stampa cosa c'è nella posizione data
func (p *piano) stato(x, y int) {
	punto := punto{x: x, y: y}

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
func (p *piano) stampaAutomi(filter func(string)) {
	fmt.Println("(")
	for nome, posizione := range p.automa {
		if filter(nome) {
			fmt.Printf("%v:%v,%v", nome, posizione.x, posizione.y)
		}
	}
	fmt.Println(")")

}
// Stampa stampa le liste degli automi e degli ostacoli
func (p *piano) stampa() {
	// Stampa automi
	p.stampaAutomi(true)

	// Stampa ostacoli
	for _, ost := range p.ostacoli {
		fmt.Printf("ostacolo %d %d %d %d\n",
			ost.AngoloInferioreSinistro.x, ost.AngoloInferioreSinistro.y,
			ost.AngoloSuperioreDestro.x, ost.AngoloSuperioreDestro.y)
	}
}
// Stampa:1 ends here

// [[file:main.org::*Automa][Automa:1]]
// Automa aggiunge o sposta un automa
func (p *piano) automa(x, y int, nome string) {
	punto := punto{x: x, y: y}

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
func (p *piano) ostacolo(x0, y0, x1, y1 int) {
	ostacolo := ostacolo{
		AngoloInferioreSinistro: punto{x: x0, y: y0}
		AngoloSuperioreDestro: punto{x: x1, y: y1}
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
func distanza(a, b punto) int {
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
func (p *piano) richiamo(x, y int, alpha string) {
	sorgente := punto{X: x, Y: y}

	// Se la sorgente è in un ostacolo, nessun movimento possibile
	if p.isPuntoInOstacolo(sorgente) {
		return
	}

	for _, aut := range p.automa {
		aut.gestisciRichiamo(sorgente)
	}
}
// Richiamo:1 ends here

// [[file:main.org::*Richiamo][Richiamo:2]]
func (a *automa) gestisciRichiamo(sorgente punto) {
	// Il punto è che faccio bfs ma aggiungo soltanto i punti la cui distanza  dalla sorgente è diminuita di 1.
}
// Richiamo:2 ends here

// [[file:main.org::*Posizioni][Posizioni:1]]
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
// Posizioni:1 ends here

// [[file:main.org::*Esiste Percorso][Esiste Percorso:1]]
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

	//corrente := da

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
// Esiste Percorso:1 ends here

// [[id:4805c5fe-d8e8-4cad-a8bb-0ed6bce1c769][Utilities:1]]
// isPuntoInOstacolo controlla se un punto è dentro un ostacolo
func (p *piano) isPuntoInOstacolo(punto punto) bool {
	for _, ost := range p.ostacoli {
		if punto.x >= ost.AngoloInferioreSinistro.x && punto.x <= ost.AngoloSuperioreDestro.x &&
			punto.y >= ost.AngoloInferioreSinistro.y && punto.y <= ost.AngoloSuperioreDestro.y {
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
