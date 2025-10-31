package main

import (
	"image/color"
	"math"
	"strings"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// Room represents a room in the colony
type Room struct {
	Name string
	X, Y float32
}

// Link represents a connection between two rooms
type Link struct {
	A, B string
}

// Move represents an ant moving to a room
type Move struct {
	Ant, Room string
}

// Game holds the state of the visualization
type Game struct {
	rooms    map[string]Room
	links    []Link
	moves    [][]Move
	ants     map[string]Room
	step     int
	lastStep time.Time
	scale    float32
	offsetX  float32
	offsetY  float32
}

// Embedded ant-farm data
const antFarmData = `
3
2 5 0
##start
0 1 2
##end
1 9 2
3 5 4
0-2
0-3
2-1
3-1
2-3

L1-2 L2-3
L1-1 L2-1 L3-2
L3-1`

// Parse ant-farm data
func parseAntFarm(data string) (map[string]Room, []Link, [][]Move) {
	lines := strings.Split(data, "\n")
	rooms := make(map[string]Room)
	var links []Link
	var moves [][]Move

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		if strings.Contains(line, "-") && !strings.HasPrefix(line, "L") {
			parts := strings.Split(line, "-")
			if len(parts) == 2 {
				links = append(links, Link{A: parts[0], B: parts[1]})
			}
		} else if strings.HasPrefix(line, "L") {
			parts := strings.Split(line, " ")
			var stepMoves []Move
			for _, m := range parts {
				moveParts := strings.Split(m, "-")
				if len(moveParts) == 2 {
					stepMoves = append(stepMoves, Move{Ant: moveParts[0], Room: moveParts[1]})
				}
			}
			if len(stepMoves) > 0 {
				moves = append(moves, stepMoves)
			}
		} else {
			parts := strings.Fields(line)
			if len(parts) == 3 {
				x := float32(parseInt(parts[1]))
				y := float32(parseInt(parts[2]))
				rooms[parts[0]] = Room{Name: parts[0], X: x, Y: y}
			}
		}
	}
	return rooms, links, moves
}

func parseInt(s string) int {
	v := 0
	negative := false
	for i := 0; i < len(s); i++ {
		if s[i] == '-' {
			negative = true
			continue
		}
		if s[i] >= '0' && s[i] <= '9' {
			v = v*10 + int(s[i]-'0')
		}
	}
	if negative {
		return -v
	}
	return v
}

// NewGame creates a new game instance
func NewGame() *Game {
	rooms, links, moves := parseAntFarm(antFarmData)
	return &Game{
		rooms:    rooms,
		links:    links,
		moves:    moves,
		ants:     make(map[string]Room),
		step:     0,
		lastStep: time.Now(),
		scale:    60,  // Scale factor
		offsetX:  100, // Left margin
		offsetY:  100, // Top margin
	}
}

// Update moves ants every second
func (g *Game) Update() error {
	if g.step < len(g.moves) && time.Since(g.lastStep) > 2000*time.Millisecond {
		for _, mv := range g.moves[g.step] {
			if room, ok := g.rooms[mv.Room]; ok {
				g.ants[mv.Ant] = room
			}
		}
		g.step++
		g.lastStep = time.Now()
	}
	return nil
}

// drawFilledCircle draws a filled circle using triangulation
func drawFilledCircle(screen *ebiten.Image, cx, cy, radius float32, clr color.Color) {
	segments := 32
	for i := 0; i < segments; i++ {
		angle1 := float32(i) * 2 * math.Pi / float32(segments)
		angle2 := float32(i+1) * 2 * math.Pi / float32(segments)

		// Create triangle from center to edge
		x1 := cx
		y1 := cy
		x2 := cx + radius*float32(math.Cos(float64(angle1)))
		y2 := cy + radius*float32(math.Sin(float64(angle1)))
		x3 := cx + radius*float32(math.Cos(float64(angle2)))
		y3 := cy + radius*float32(math.Sin(float64(angle2)))

		// Draw lines to form filled triangle
		vector.StrokeLine(screen, x1, y1, x2, y2, 1, clr, false)
		vector.StrokeLine(screen, x2, y2, x3, y3, 1, clr, false)
		vector.StrokeLine(screen, x3, y3, x1, y1, 1, clr, false)
		
		// Fill with additional lines
		for j := float32(0); j < radius; j += 0.5 {
			scale := j / radius
			ix2 := cx + scale*radius*float32(math.Cos(float64(angle1)))
			iy2 := cy + scale*radius*float32(math.Sin(float64(angle1)))
			ix3 := cx + scale*radius*float32(math.Cos(float64(angle2)))
			iy3 := cy + scale*radius*float32(math.Sin(float64(angle2)))
			vector.StrokeLine(screen, ix2, iy2, ix3, iy3, 1, clr, false)
		}
	}
}

// Draw renders rooms, links, and ants
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)

	// Draw links between rooms
	for _, l := range g.links {
		a, okA := g.rooms[l.A]
		b, okB := g.rooms[l.B]
		if okA && okB {
			x1 := a.X*g.scale + g.offsetX
			y1 := a.Y*g.scale + g.offsetY
			x2 := b.X*g.scale + g.offsetX
			y2 := b.Y*g.scale + g.offsetY
			vector.StrokeLine(screen, x1, y1, x2, y2, 2, color.White, false)
		}
	}

	// Draw rooms as blue filled circles
	for _, r := range g.rooms {
		cx := r.X*g.scale + g.offsetX
		cy := r.Y*g.scale + g.offsetY
		drawFilledCircle(screen, cx, cy, 20, color.RGBA{0, 180, 255, 255})
	}

	// Draw ants as yellow filled circles
	for _, pos := range g.ants {
		cx := pos.X*g.scale + g.offsetX
		cy := pos.Y*g.scale + g.offsetY
		drawFilledCircle(screen, cx, cy, 10, color.RGBA{255, 255, 0, 255})
	}
}

// Layout defines the window size
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 1000, 600
}

// Main entry point
func main() {
	game := NewGame()
	ebiten.SetWindowSize(1000, 600)
	ebiten.SetWindowTitle("Lem-in Ant Farm 🐜")
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}