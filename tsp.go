package main
  
import (
    "fmt"
    "math/rand"
)

type Node struct{
    name string
    X,Y int
}

type Vertex struct {
    first, second *Node
}

func generateNodes(n, field_size int) []Node{
    rand.Seed(0)
    var nodes = make([]Node, n)
    for i := 0; i < n; i++ {
        nodes[i] = Node{string('A' - 1 + i+1), rand.Intn(field_size), rand.Intn(field_size)}
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
    n = func(testStr []rune, p []string) []string{
        if len(testStr) == 0 {
            return p
        }else {
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


  
func main() {

    nodes:=generateNodes(10, 100)
    for _, node := range nodes {
        fmt.Print(node)
    }
    d := permutations("ABCD")
    fmt.Print(d)
}
