// [[file:main.org::*Backbone][Backbone:1]]
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// punto rappresenta una posizione nella griglia 2D
type punto struct {
	x, y int
}

// automa rappresenta un singolo automa
type automa = punto

// ostacolo rappresenta un ostacolo rettangolare
type ostacolo struct {
	angoloInferioreSinistro punto
	angoloSuperioreDestro   punto
}

// piano rappresenta l'intero sistema
type piano struct {
	automi            map[string]automa
	ostacoli          *List[ostacolo]
	daReinizializzare *bool
}

func main() {
	p := newPiano()

	reader := bufio.NewReader(os.Stdin)
	for {
		if *p.daReinizializzare {
			p = newPiano()
			*p.daReinizializzare = false
		}
		riga, _ := reader.ReadString('\n')
		esegui(p, riga)
	}
}

// Backbone:1 ends here

// [[file:main.org::newPiano][newPiano]]
// newPiano crea un piano vuoto
func newPiano() piano {
	return piano{
		automi:            make(map[string]automa),
		ostacoli:          NewList[ostacolo](),
		daReinizializzare: new(bool),
	}
}

// newPiano ends here

// [[file:main.org::esegui][esegui]]
func esegui(p piano, s string) {
	parti := strings.Fields(s)
	comando := parti[0][0]

	switch comando {
	case 'c':
		*p.daReinizializzare = true
	case 'S':
		p.stampaAutomi("")
		p.stampaOstacoli()
	case 's':
		a, _ := strconv.Atoi(parti[1])
		b, _ := strconv.Atoi(parti[2])
		p.stampaStato(a, b)
	case 'a':
		a, _ := strconv.Atoi(parti[1])
		b, _ := strconv.Atoi(parti[2])
		omega := parti[3]
		p.aggiungiAutoma(a, b, omega)
	case 'o':
		a, _ := strconv.Atoi(parti[1])
		b, _ := strconv.Atoi(parti[2])
		c, _ := strconv.Atoi(parti[3])
		d, _ := strconv.Atoi(parti[4])
		p.aggiungiOstacolo(a, b, c, d)
	case 'r':
		x, _ := strconv.Atoi(parti[1])
		y, _ := strconv.Atoi(parti[2])
		alpha := parti[3]
		p.emettiRichiamo(x, y, alpha)
	case 'p':
		omega := parti[1]
		p.stampaAutomi(omega)
	case 'e':
		x, _ := strconv.Atoi(parti[1])
		y, _ := strconv.Atoi(parti[2])
		nome := parti[3]
		p.stampaEsistenzaPercorso(x, y, nome)
	case 'f':
		os.Exit(0)
	}

}

// esegui ends here

// [[file:main.org::*Stato][Stato:1]]
// Stato stampa cosa c'è nella posizione data
func (p *piano) stampaStato(x, y int) {
	punto := punto{x: x, y: y}

	// Controlla se il punto è in un ostacolo
	if p.isPuntoInQualcheOstacolo(punto) {
		fmt.Println("O")
		return
	}

	// Controlla se il punto contiene un automa
	if p.isPuntoAutoma(punto) {
		fmt.Println("A")
		return
	}

	// Se nessuna delle condizioni precedenti è soddisfatta, il punto è vuoto
	fmt.Println("E")
}

// Stato:1 ends here

// [[file:main.org::*Stampa][Stampa:1]]
func (p *piano) stampaAutomi(filtro string) {
	fmt.Println("(")
	for nome, automa := range p.automi {
		if filtro == "" || verificaPrefisso(filtro, nome) {
			fmt.Printf("%v: %v,%v\n", nome, automa.x, automa.y)
		}
	}
	fmt.Println(")")
}
func (p *piano) stampaOstacoli() {
	fmt.Println("[")
	current := p.ostacoli.Head
	for current != nil {
		fmt.Printf("(%v,%v)(%v,%v)\n",
			current.Value.angoloInferioreSinistro.x,
			current.Value.angoloInferioreSinistro.y,
			current.Value.angoloSuperioreDestro.x,
			current.Value.angoloSuperioreDestro.y)
		current = current.Next
	}
	fmt.Println("]")
}

// Stampa:1 ends here

// [[file:main.org::automa][automa]]
// Automa aggiunge o sposta un automa
func (p *piano) aggiungiAutoma(x, y int, nome string) {
	punto := punto{x: x, y: y}

	// Controlla se il punto è in un ostacolo
	if p.isPuntoInQualcheOstacolo(punto) {
		return
	}

	// Crea o sposta automa
	p.automi[nome] = automa{
		x: x,
		y: y,
	}
}

// automa ends here

// [[file:main.org::*Ostacolo][Ostacolo:1]]
// Ostacolo aggiunge un ostacolo se possibile
func (p *piano) aggiungiOstacolo(x0, y0, x1, y1 int) {

	ostacolo := ostacolo{
		angoloInferioreSinistro: punto{x: x0, y: y0},
		angoloSuperioreDestro:   punto{x: x1, y: y1},
	}

	// Controlla se un automa è nell'area proposta per l'ostacolo
	if p.isOstacoloSuQualchePunto(ostacolo) {
		return
	}

	// Aggiungi ostacolo
	p.ostacoli.AddNewNode(ostacolo)
}

// Ostacolo:1 ends here

// [[file:main.org::*Richiamo][Richiamo:1]]
// Calcolo della calcolaDistanzaManhattan di Manhattan
func calcolaDistanzaManhattan(a, b punto) int {
	return valoreAssoluto(b.x-a.x) + valoreAssoluto(b.y-a.y)
}

func valoreAssoluto(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// verificaPrefisso controlla se α è un prefisso di η
func verificaPrefisso(alpha, eta string) bool {
	if len(alpha) > len(eta) {
		return false
	}
	return eta[:len(alpha)] == alpha
}

// Richiamo emette un segnale dal punto (x,y)
func (p *piano) emettiRichiamo(x, y int, alpha string) {
	sorgente := punto{x: x, y: y}

	// Se la sorgente è in un ostacolo, nessun movimento possibile
	if p.isPuntoInQualcheOstacolo(sorgente) {
		return
	}

	for nome, automa := range p.automi {
		if verificaPrefisso(alpha, nome) && p.esistePercorsoMinimoLibero(sorgente, automa) {
			automa.x = sorgente.x
			automa.y = sorgente.y
		}
	}
}

// Richiamo:1 ends here

// [[file:main.org::*Esiste Percorso][Esiste Percorso:1]]
// EsistePercorso controlla se esiste un percorso libero
func (p *piano) stampaEsistenzaPercorso(x, y int, nome string) {
	dest := punto{x: x, y: y}

	// Controlla se l'automa esiste e se la destinazione è valida
	aut, esiste := p.automi[nome]

	if !esiste || p.isPuntoInQualcheOstacolo(dest) {
		fmt.Println(!esiste, p.isPuntoInQualcheOstacolo(dest))
		fmt.Println("NO")
		return
	}

	if p.esistePercorsoMinimoLibero(aut, dest) {
		fmt.Println("SI")
	} else {
		fmt.Println("NO")
	}
}

// trovaDirezioni returns the immediate direction from current to target.
// For instance, if target.x > current.x, direction's x-component is +1 (moving right).
func trovaPassiUnitariAmmessi(da, a punto) [][2]int {

	passiUnitariAmmessi := [][2]int{}

	// Compare x-coordinates to see if we should move left or right
	if a.x > da.x {
		passiUnitariAmmessi = append(passiUnitariAmmessi, [2]int{1, 0})
	} else if a.x < da.x {
		passiUnitariAmmessi = append(passiUnitariAmmessi, [2]int{-1, 0})
	}

	// Compare y-coordinates to see if we should move up or down
	if a.y > da.y {
		passiUnitariAmmessi = append(passiUnitariAmmessi, [2]int{0, 1})
	} else if a.y < da.y {
		passiUnitariAmmessi = append(passiUnitariAmmessi, [2]int{0, -1})
	}

	return passiUnitariAmmessi
}

// esistePercorsoLibero controlla se esiste un percorso libero di lunghezza minima
// non ci serve controllare se un punto è visitato o meno ( adifferenzza della BFS) perchè per rivisitare un punto bisognerebbe allontanarsi dalla destinazione e questo non succede
func (p *piano) esistePercorsoMinimoLibero(aut automa, dest punto) bool {

	direzioni := trovaPassiUnitariAmmessi(aut, dest)

	// Initialize queue with the starting point
	var queue Queue[punto]

	queue.Push(aut)

	for current, ok := queue.Pop(); ok; current, ok = queue.Pop() {

		if current == dest {
			return true
		}
		// Explore neighbors using the directions vector
		for _, dir := range direzioni {
			prossimoPunto := punto{
				x: current.x + dir[0],
				y: current.y + dir[1],
			}

			if calcolaDistanzaManhattan(prossimoPunto, dest) < calcolaDistanzaManhattan(current, dest) && !p.isPuntoInQualcheOstacolo(prossimoPunto) {
				queue.Push(prossimoPunto)
			}
		}
	}

	// If we exhaust all possibilities without reaching the end, no path exists
	return false
}

// Esiste Percorso:1 ends here

// [[id:4805c5fe-d8e8-4cad-a8bb-0ed6bce1c769][Utilities:1]]
// isPuntoInQualcheOstacolo controlla se un punto è dentro un ostacolo
func (p *piano) isPuntoInQualcheOstacolo(punto punto) bool {
	current := p.ostacoli.Head
	for current != nil {
		if isPuntoInOstacolo(punto, current.Value) {
			return true
		}
		current = current.Next
	}
	return false
}

// isPuntoAutoma checks if a point is an automaton
func (p *piano) isPuntoAutoma(punto punto) bool {
	for _, aut := range p.automi {
		if aut == punto {
			return true
		}
	}
	return false
}

// isOstacoloSuQualchePunto checks if there is any automaton within the given obstacle
func (p *piano) isOstacoloSuQualchePunto(ostacolo ostacolo) bool {
	for _, aut := range p.automi {
		if isPuntoInOstacolo(aut, ostacolo) {
			return true
		}
	}
	return false
}

func isPuntoInOstacolo(p punto, o ostacolo) bool {
	return p.x >= o.angoloInferioreSinistro.x &&
		p.x <= o.angoloSuperioreDestro.x &&
		p.y >= o.angoloInferioreSinistro.y &&
		p.y <= o.angoloSuperioreDestro.y
}

// Utilities:1 ends here
