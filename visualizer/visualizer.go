package main

import (
	"bufio"
	"fmt"
	"image/color"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Room represents a node in the graph
type Room struct {
	Name string
	X, Y int
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
	rooms map[string]Room
	links []Link
	moves [][]Move
	ants  map[string]Room
	step  int
	lastStep time.Time
}

func parseInput() *Game {
	scanner := bufio.NewScanner(os.Stdin)
	rooms := make(map[string]Room)
	var links []Link
	var moves [][]Move

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		if strings.Contains(line, "-") && !strings.HasPrefix(line, "L") {
			// link
			parts := strings.Split(line, "-")
			links = append(links, Link{parts[0], parts[1]})
		} else if strings.HasPrefix(line, "L") {
			// move line
			moveParts := strings.Split(line, " ")
			var stepMoves []Move
			for _, m := range moveParts {
				parts := strings.Split(m, "-")
				stepMoves = append(stepMoves, Move{Ant: parts[0], Room: parts[1]})
			}
			moves = append(moves, stepMoves)
		} else {
			// room
			parts := strings.Fields(line)
			if len(parts) == 3 {
				x, _ := strconv.Atoi(parts[1])
				y, _ := strconv.Atoi(parts[2])
				rooms[parts[0]] = Room{parts[0], x, y}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return &Game{
		rooms: rooms,
		links: links,
		moves: moves,
		ants:  make(map[string]Room),
		step:  0,
		lastStep: time.Now(),
	}
}

// Update moves ants at a fixed interval
func (g *Game) Update() error {
	// move every 500ms
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

// Draw the rooms, links, and ants
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 0, 255}) 

	// Draw links
	for _, l := range g.links {
		a := g.rooms[l.A]
		b := g.rooms[l.B]
		ebitenutil.DrawLine(screen, float64(a.X*30), float64(a.Y*30),
			float64(b.X*30), float64(b.Y*30), color.RGBA{255, 255, 255, 255})
	}

	// Draw rooms
	for _, r := range g.rooms {
		ebitenutil.DrawCircle(screen, float64(r.X*30+2), float64(r.Y*30+2), 20, color.RGBA{0, 180, 255, 255})
		ebitenutil.DebugPrintAt(screen, r.Name, r.X*30, r.Y*30)
	}

	// Draw ants
	for _, pos := range g.ants {
		ebitenutil.DrawCircle(screen, float64(pos.X*30), float64(pos.Y*30), 10, color.RGBA{255, 255, 0, 255})
	}
}

// Layout window size
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 1000, 600
}

// Entry point
func main() {
	game := parseInput()
	ebiten.SetWindowSize(1000, 600)
	ebiten.SetWindowTitle("Lem-in Ant Farm 🐜")
	if err := ebiten.RunGame(game); err != nil {
		fmt.Println("Error:", err)
	}
}
