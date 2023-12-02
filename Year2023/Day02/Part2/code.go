package y2023d02p02

import (
	"strconv"
	"strings"
)

type Draw struct {
	r int
	g int
	b int
}

type maxBalls struct {
	r int
	g int
	b int
}

type Game struct {
	draws []Draw
}

func (g Game) getMaxBalls() maxBalls {
	r := 0
	for d := range g.draws {
		if g.draws[d].r > r {
			r = g.draws[d].r
		}
	}
	green := 0
	for d := range g.draws {
		if g.draws[d].g > green {
			green = g.draws[d].g
		}
	}
	b := 0
	for d := range g.draws {
		if g.draws[d].b > b {
			b = g.draws[d].b
		}
	}
	return maxBalls{r: r, g: green, b: b}
}

func parseDraw(draw string) Draw {

	ballSets := strings.Split(draw, ", ")

	r := 0
	g := 0
	b := 0

	for _, ballSet := range ballSets {
		ballQty := strings.Split(ballSet, " ")
		qty, _ := strconv.Atoi(ballQty[0])
		if ballQty[1] == "red" {
			if qty > r {
				r = qty
			}
		} else if ballQty[1] == "green" {
			if qty > g {
				g = qty
			}
		} else if ballQty[1] == "blue" {
			if qty > b {
				b = qty
			}
		}
	}

	parsedDraw := Draw{r: r, g: g, b: b}

	return parsedDraw
}

func parseGame(game string) Game {

	draws := strings.Split(game, "; ")
	thisGame := Game{draws: make([]Draw, len(draws))}
	for i := range draws {
		thisGame.draws = append(thisGame.draws, parseDraw(draws[i]))
	}

	return thisGame

}

func Calculate(input []string) int {

	result := 0
	for i := range input {
		gameData := strings.Split(input[i], ": ")
		game := parseGame(gameData[1])
		maxBallsInGame := game.getMaxBalls()

		result = result + (maxBallsInGame.r * maxBallsInGame.g * maxBallsInGame.b)
	}
	return result
}
