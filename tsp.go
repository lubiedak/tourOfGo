package main

import (
	"fmt"
	"math/rand"
	"math"
	"strings"
	"time"
)

type Node struct {
	name string
	X, Y int
}

type Vertex struct {
	first, second *Node
}

func generateNodes(n, field_size int) []Node {
	rand.Seed(0)
	var nodes = make([]Node, n)
	for i := 0; i < n; i++ {
		nodes[i] = Node{string('A' - 1 + i + 1), rand.Intn(field_size), rand.Intn(field_size)}
	}
	return nodes
}

func join(ins []rune, c rune) (result []string) {
	for i := 0; i <= len(ins); i++ {
		result = append(result, string(ins[:i])+string(c)+string(ins[i:]))
	}
	return
}

func permutations(testStr string) []string {
	var n func(testStr []rune, p []string) []string
	n = func(testStr []rune, p []string) []string {
		if len(testStr) == 0 {
			return p
		} else {
			result := []string{}
			for _, e := range p {
				result = append(result, join([]rune(e), testStr[0])...)
			}
			return n(testStr[1:], result)
		}
	}

	output := []rune(testStr)
	return n(output[1:], []string{string(output[0])})
}

func getNodeByNameRef(nodes *[]Node, name string) *Node {
	for _, node := range *nodes {
		if node.name == name {
			return &node
		}
	}
	return nil
}

func getNodeByName(nodes []Node, name string) Node {
	for _, node := range nodes {
		if node.name == name {
			return node
		}
	}
	return nodes[0]
}

func distance(a, b Node) float64 {
	return math.Sqrt(math.Pow(float64(a.X-b.X), 2.0) + math.Pow(float64(a.Y-b.Y), 2.0))
}

func distanceRef(a, b *Node) float64 {
	return math.Sqrt(math.Pow(float64(a.X-b.X), 2.0) + math.Pow(float64(a.Y-b.Y), 2.0))
}

func calculateLength(nodes []Node, sequence string) float64 {
	nodeNames := strings.Split(sequence, "")
	length := distance(getNodeByName(nodes, nodeNames[0]), getNodeByName(nodes, nodeNames[len(nodeNames)-1]))
	for i := 0; i < len(nodes)-1; i++ {
		length += distance(getNodeByName(nodes, nodeNames[i]), getNodeByName(nodes, nodeNames[i+1]))
	}
	return length
}

func calculateLengthRef(nodes *[]Node, sequence string) float64 {
	nodeNames := strings.Split(sequence, "")
	length := distanceRef(getNodeByNameRef(nodes, nodeNames[0]), getNodeByNameRef(nodes, nodeNames[len(nodeNames)-1]))
	for i := 0; i < len(*nodes)-1; i++ {
		length += distanceRef(getNodeByNameRef(nodes, nodeNames[i]), getNodeByNameRef(nodes, nodeNames[i+1]))
	}
	return length
}

func findMTSP(nodes []Node, allPerms []string) string {
	minimumLength := 9999999.0
	bestPerm := allPerms[0]
	for _, perm := range allPerms {
		length := calculateLengthRef(&nodes, perm)
		if length < minimumLength {
			fmt.Printf("New minimum: %f\n", length)
			minimumLength = length
			bestPerm = perm
		}
	}
	return bestPerm
}

func main() {
	nodes := generateNodes(11, 100)

	t1 := time.Now()
	allPerms := permutations("ABCDEFGHIJK")
	fmt.Println(len(allPerms))
	fmt.Println(time.Now().Sub(t1))
	t1 = time.Now()

	bestPerm := findMTSP(nodes, allPerms)
	fmt.Printf("Minimum length %f\n", calculateLength(nodes, bestPerm))
	fmt.Println(bestPerm)

	fmt.Println(time.Now().Sub(t1))
}
