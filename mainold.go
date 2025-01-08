package automiesegnali

import (
	"fmt"
	"sort"
)

// Punto rappresenta una posizione nella griglia 2D
type Punto struct {
	X, Y int
}

// Automa rappresenta un singolo automa
type Automa struct {
	Nome      string
	Posizione Punto
}

// Ostacolo rappresenta un ostacolo rettangolare
type Ostacolo struct {
	AngoloInferioreSinistro Punto
	AngoloSuperioreDestro   Punto
}

// Piano rappresenta l'intero sistema
type Piano struct {
	automa   map[string]*Automa
	ostacoli []*Ostacolo
}

// Crea crea un piano vuoto
func Crea() *Piano {
	return &Piano{
		automa:   make(map[string]*Automa),
		ostacoli: make([]*Ostacolo, 0),
	}
}

// Stato stampa cosa c'è nella posizione data
func (p *Piano) Stato(x, y int) {
	punto := Punto{X: x, Y: y}

	// Controlla se il punto è in un ostacolo
	if p.isPuntoInOstacolo(punto) {
		fmt.Println("O")
		return
	}

	// Controlla se il punto contiene un automa
	for _, aut := range p.automa {
		if aut.Posizione == punto {
			fmt.Println("A")
			return
		}
	}

	// Se nessuna delle condizioni precedenti è soddisfatta, il punto è vuoto
	fmt.Println("E")
}

// TODO
// Stampa stampa le liste degli automi e degli ostacoli
func (p *Piano) Stampa() {
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

// isPuntoInOstacolo controlla se un punto è dentro un ostacolo
func (p *Piano) isPuntoInOstacolo(punto Punto) bool {
	for _, ost := range p.ostacoli {
		if punto.X >= ost.AngoloInferioreSinistro.X && punto.X <= ost.AngoloSuperioreDestro.X &&
			punto.Y >= ost.AngoloInferioreSinistro.Y && punto.Y <= ost.AngoloSuperioreDestro.Y {
			return true
		}
	}
	return false
}

// Automa aggiunge o sposta un automa
func (p *Piano) Automa(x, y int, nome string) {
	punto := Punto{X: x, Y: y}

	// Controlla se il punto è in un ostacolo
	if p.isPuntoInOstacolo(punto) {
		return
	}

	// Crea o sposta automa
	p.automa[nome] = &Automa{
		Nome:      nome,
		Posizione: punto,
	}
}

// Ostacolo aggiunge un ostacolo se possibile
func (p *Piano) Ostacolo(x0, y0, x1, y1 int) {
	// Controlla se un automa è nell'area proposta per l'ostacolo
	for _, aut := range p.automa {
		if aut.Posizione.X >= x0 && aut.Posizione.X <= x1 &&
			aut.Posizione.Y >= y0 && aut.Posizione.Y <= y1 {
			return
		}
	}

	// Aggiungi ostacolo
	p.ostacoli = append(p.ostacoli, &Ostacolo{
		AngoloInferioreSinistro: Punto{X: x0, Y: y0},
		AngoloSuperioreDestro:   Punto{X: x1, Y: y1},
	})
}

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
