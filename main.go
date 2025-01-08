package automiesegnali

import (
	"fmt"
	"sort"
)

// Point represents a position in the 2D grid
type Point struct {
	X, Y int
}

// Automaton represents a single automaton
type Automaton struct {
	Name     string
	Position Point
}

// Obstacle represents a rectangular obstacle
type Obstacle struct {
	BottomLeft Point
	TopRight   Point
}

// Plane represents the entire system
type Plane struct {
	automata  map[string]*Automaton
	obstacles []*Obstacle
}

// Create creates an empty plane
func Create() *Plane {
	return &Plane{
		automata:  make(map[string]*Automaton),
		obstacles: make([]*Obstacle, 0),
	}
}

// State prints what is at the given location
func (p *Plane) State(x, y int) {
	point := Point{X: x, Y: y}

	// Check if point is in any obstacle
	if p.isPointInObstacle(point) {
		fmt.Println("O")
		return
	}

	// Check if point contains an automaton
	for _, aut := range p.automata {
		if aut.Position == point {
			fmt.Println("A")
			return
		}
	}

	// If none of the previous conditions is met, the point is empty
	fmt.Println("E")
}

// TODO
// Print prints the lists of automata and obstacles
func (p *Plane) Print() {
	// Get sorted list of automaton names for consistent output
	names := make([]string, 0, len(p.automata))
	for name := range p.automata {
		names = append(names, name)
	}
	sort.Strings(names)

	// Print automata
	for _, name := range names {
		a := p.automata[name]
		fmt.Printf("automaton %d %d %s\n", a.Position.X, a.Position.Y, a.Name)
	}

	// Print obstacles
	for _, obs := range p.obstacles {
		fmt.Printf("obstacle %d %d %d %d\n",
			obs.BottomLeft.X, obs.BottomLeft.Y,
			obs.TopRight.X, obs.TopRight.Y)
	}
}

// isPointInObstacle checks if a point is inside any obstacle
func (p *Plane) isPointInObstacle(point Point) bool {
	for _, obs := range p.obstacles {
		if point.X >= obs.BottomLeft.X && point.X <= obs.TopRight.X &&
			point.Y >= obs.BottomLeft.Y && point.Y <= obs.TopRight.Y {
			return true
		}
	}
	return false
}

// Automaton adds or moves an automaton
func (p *Plane) Automaton(x, y int, name string) {
	point := Point{X: x, Y: y}

	// Check if point is in obstacle
	if p.isPointInObstacle(point) {
		return
	}

	// Create or move automaton
	p.automata[name] = &Automaton{
		Name:     name,
		Position: point,
	}
}

// Obstacle adds an obstacle if possible
func (p *Plane) Obstacle(x0, y0, x1, y1 int) {
	// Check if any automaton is in the proposed obstacle area
	for _, aut := range p.automata {
		if aut.Position.X >= x0 && aut.Position.X <= x1 &&
			aut.Position.Y >= y0 && aut.Position.Y <= y1 {
			return
		}
	}

	// Add obstacle
	p.obstacles = append(p.obstacles, &Obstacle{
		BottomLeft: Point{X: x0, Y: y0},
		TopRight:   Point{X: x1, Y: y1},
	})
}

// Manhattan distance calculation
func distance(a, b Point) int {
	return abs(b.X-a.X) + abs(b.Y-b.Y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// isPrefix checks if α is a prefix of η
func isPrefix(alpha, eta string) bool {
	if len(alpha) > len(eta) {
		return false
	}
	return eta[:len(alpha)] == alpha
}

// Recall emits a signal from point (x,y)
func (p *Plane) Recall(x, y int, alpha string) {
	source := Point{X: x, Y: y}

	// If source is in obstacle, no movement possible
	if p.isPointInObstacle(source) {
		return
	}

	// Find responding automata and minimum distance
	respondingAutomata := make([]*Automaton, 0)
	minDist := -1

	for _, aut := range p.automata {
		if isPrefix(alpha, aut.Name) {
			dist := distance(source, aut.Position)
			if minDist == -1 || dist < minDist {
				minDist = dist
				respondingAutomata = respondingAutomata[:0]
				respondingAutomata = append(respondingAutomata, aut)
			} else if dist == minDist {
				respondingAutomata = append(respondingAutomata, aut)
			}
		}
	}

	// Move qualifying automata
	for _, aut := range respondingAutomata {
		if p.existsFreePath(aut.Position, source) {
			aut.Position = source
		}
	}
}

// Positions prints positions of automata matching prefix
func (p *Plane) Positions(alpha string) {
	// Get matching automata sorted by name
	matching := make([]string, 0)
	for name := range p.automata {
		if isPrefix(alpha, name) {
			matching = append(matching, name)
		}
	}
	sort.Strings(matching)

	// Print positions
	for _, name := range matching {
		aut := p.automata[name]
		fmt.Printf("%s %d %d\n", name, aut.Position.X, aut.Position.Y)
	}
}

// ExistsPath checks if a free path exists
func (p *Plane) ExistsPath(x, y int, name string) {
	dest := Point{X: x, Y: y}

	// Check if automaton exists and destination is valid
	aut, exists := p.automata[name]
	if !exists || p.isPointInObstacle(dest) {
		fmt.Println("NO")
		return
	}

	if p.existsFreePath(aut.Position, dest) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

// existsFreePath checks if there exists a free path of minimum length
func (p *Plane) existsFreePath(from, to Point) bool {
	// Implementation of path finding algorithm
	// This is a simplified version that checks if any point in the
	// minimum distance path intersects with obstacles

	// Get the direction of movement
	dx := sign(to.X - from.X)
	dy := sign(to.Y - from.Y)

	current := from

	// First try horizontal movement
	for current.X != to.X {
		current.X += dx
		if p.isPointInObstacle(current) {
			return false
		}
	}

	// Then vertical movement
	for current.Y != to.Y {
		current.Y += dy
		if p.isPointInObstacle(current) {
			return false
		}
	}

	return true
}

func sign(x int) int {
	if x < 0 {
		return -1
	}
	if x > 0 {
		return 1
	}
	return 0
}
