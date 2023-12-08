package y2023d08p02

import (
	"fmt"
	"math"
	"regexp"
	"strings"
)

type Options struct {
	L string
	R string
}

func ParseNode(input string) (string, Options) {

	inputParts := strings.Split(input, " = ")
	nodeName := inputParts[0]

	optionsRegex := regexp.MustCompile(`\(([0-9A-Za-z]+), ([0-9A-Za-z]+)\)`)

	options := optionsRegex.FindAllStringSubmatch(inputParts[1], -1)

	return nodeName, Options{L: options[0][1], R: options[0][2]}

}

func followPath(network map[string]Options, path []string, startingNode string, c chan int) {

	pathIndex := 0
	currentNode := startingNode
	stepCount := 0
	for zFoundCount := 0; zFoundCount < 1; {
		if path[pathIndex] == "R" {
			nextNode := network[currentNode].R
			currentNode = nextNode
		} else {
			nextNode := network[currentNode].L
			currentNode = nextNode
		}
		pathIndex = (pathIndex + 1) % len(path)

		if zFoundCount == 0 {
			stepCount++
		}

		if strings.HasSuffix(currentNode, "Z") {
			zFoundCount++
		}

	}

	c <- stepCount
}

func gcd(a int, b int) int {
	// The steps
	//    gcd(0, v) = v, because everything divides zero, and v is the largest number that divides v. Similarly, gcd(u, 0) = u.
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	//    gcd(2u, 2v) = 2·gcd(u, v)
	if (a%2 == 0) && (b%2 == 0) {
		return 2 * gcd(a/2, b/2)
	}
	//    gcd(2u, v) = gcd(u, v), if v is odd (2 is not a common divisor). Similarly, gcd(u, 2v) = gcd(u, v) if u is odd.
	if (a%2 == 0) && (b%2 != 0) {
		return gcd(a/2, b)
	}
	if (a%2 != 0) && (b%2 == 0) {
		return gcd(a, b/2)
	}
	//    gcd(u, v) = gcd(|u − v|, min(u, v)), if u and v are both odd.
	return gcd(int(math.Abs(float64(a-b))), int(math.Min(float64(a), float64(b))))
}

func lcm(a int, b int) int {
	return a * (b / gcd(a, b))
}

func Calculate(input string) int {

	rows := strings.Split(input, "\n")

	path := strings.Split(rows[0], "")

	network := map[string]Options{}

	for i := 2; i < len(rows); i++ {
		nodeName, options := ParseNode(rows[i])
		network[nodeName] = options
	}

	var currentNodes []string

	for nodeName := range network {
		if strings.HasSuffix(nodeName, "A") {
			currentNodes = append(currentNodes, nodeName)
		}
	}

	// This entire thing relies on the fact that the 'ghosts' obey the following rules. Which it does for my input
	// 1) Each unique A visits a unique Z value. It never visits another Z value
	// 2) Once it has visited a Z value once it loops. And the loop length is the same the second time around.
	//    (This was tested by doing it twice and inspecting the lengths of the loop)
	var channels []chan int
	for i := range currentNodes {
		channels = append(channels, make(chan int))
		go followPath(network, path, currentNodes[i], channels[i])
	}

	var stepCounts []int
	for i := 0; i < len(channels); i++ {
		stepCount := <-channels[i]
		stepCounts = append(stepCounts, stepCount)
	}

	fmt.Printf("Step Counts: %+v\n", stepCounts)

	// I think what we now need to do is find the lowest number that is a multiple of each step count....
	// How the heck do I work out a lowest common multiple again?

	for i := 0; i < len(stepCounts)-1; i++ {
		a := stepCounts[i]
		b := stepCounts[i+1]
		lcmResult := lcm(a, b)
		stepCounts[i+1] = lcmResult
	}

	return stepCounts[len(stepCounts)-1]
}
