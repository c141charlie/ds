package ternary

import "fmt"

const WILDCARD = "?"

type TernaryTree struct {
    root *Node
}

func NewTernaryTree() *TernaryTree {
    return &TernaryTree{}
}

func (t *TernaryTree) Contains(word string) bool {
    if word == "" { return false } 
    node := t.search(t.root, word, 0)
   
    return node != nil && node.isEndOfWord()
}

func (t *TernaryTree) search(node *Node, word string, index int) *Node {
    
    if node == nil || word == "" { return nil }

    runes := []int(word)

    if runes[index] == node.Rune {
        if index + 1 < len(runes) {
            node = t.search(node.Child, word, index + 1)
        }
    
    } else if runes[index] < node.Rune {
        node = t.search(node.Smaller, word, index)
    
    } else {
        node = t.search(node.Larger, word, index)
    }

    return node
}

func (t *TernaryTree) Add(word string) {
    if word == "" { return }
    node := t.insert(t.root, word, 0)

    if t.root == nil {
        t.root = node
    }
}

func (t *TernaryTree) insert(node *Node, word string, index int) *Node {
    if word == "" || index < 0 { return nil }

    runes := []int(word)
    rune := runes[index]

    if node == nil {
        node = &Node{rune, nil, nil, nil, nil}
    }

    if node.Rune == rune {
        if index + 1 < len(runes) {
            node.Child = t.insert(node.Child, word, index +1)
        } else {
            node.Word = runes
        }
    } else if rune < node.Rune {
        node.Smaller = t.insert(node.Smaller, word, index)
    } else {
        node.Larger = t.insert(node.Larger, word, index)
    }

    return node
}

func (t *TernaryTree) BreadthFirstTraversal(node *Node) {
    if node.Smaller != nil {
        t.BreadthFirstTraversal(node.Smaller)
    }
    if node.isEndOfWord() == true {
        fmt.Println(string(node.Word))
    }
    if node.Child != nil {
        t.BreadthFirstTraversal(node.Child)
    }
    if node.Larger != nil {
        t.BreadthFirstTraversal(node.Larger)
    }
}










type Node struct {
    
    Rune int
    Smaller *Node
    Larger *Node
    Child *Node
    Word []int
}

func NewNode(rune int) *Node {
    return &Node{rune, nil, nil, nil, nil}
}

func (n *Node) isEndOfWord() bool {
    return n.Word != nil
}

