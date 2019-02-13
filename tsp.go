package main
  
import (
    "fmt"
    "math/rand"
    "math"
    "strings"
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

func getNodeByName(nodes []Node, name string ) *Node{
    for _, node := range nodes{
        if node.name == name {
            return &node
        }
    }
    return nil
}

func distance(a,b Node) float64{
    return math.Sqrt(math.Pow(float64(a.X-b.X),2.0) + math.Pow(float64(a.Y-b.Y),2.0))
}

func distanceRef(a,b *Node) float64{
    return math.Sqrt(math.Pow(float64(a.X-b.X),2.0) + math.Pow(float64(a.Y-b.Y),2.0))
}


func calculateLength(nodes []Node, sequence string) float64{
    nodeNames := strings.Split(sequence, "")
    length:= distanceRef(getNodeByName(nodes, nodeNames[0]), getNodeByName(nodes, nodeNames[len(nodeNames)-1]))
    for i:=0; i<len(nodes)-1;i++{
        length+=distanceRef(getNodeByName(nodes, nodeNames[i]), getNodeByName(nodes, nodeNames[i+1]))
    }
    return length
}


func main() {

    nodes:=generateNodes(10, 100)
    for _, node := range nodes {
        fmt.Println(node)
    }

    fmt.Println(getNodeByName(nodes, "B"))
    fmt.Println(getNodeByName(nodes, "W"))


    allPerms := permutations("ABCDEFGHIJ")
    fmt.Println(len(allPerms))
    minimumLength:= 9999999.0
    bestPerm:= allPerms[0]
    for _, perm := range allPerms {
        length:=calculateLength(nodes, perm)
        if(length < minimumLength){
            fmt.Printf("New minimum: %f\n", length)
            minimumLength = length
            bestPerm = perm
        }
    }
    fmt.Println(minimumLength)
    fmt.Println(bestPerm)
}
