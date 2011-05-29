package ternary

const WILDCARD = "?"

type TernaryTree struct {
    root *Node
}

func NewTernaryTree() *TernaryTree {
    return &TernaryTree{}
}

func (t *TernaryTree) Contains(word string) bool {
    if word == "" { return false } 
    node := search(t.root, word, 0)
    return node != nil && node.isEndOfWord()
}

func search(node *Node, word string, index int) *Node {
    if word == "" { return }
}

func (t *TernaryTree) Add(word string) {
    if word == "" { return }
    node := insert(t.root, word, 0)

    if t.root == nil {
        t.root = node
    }
}

func insert(node *Node, word string, index int) *Node {
    runes := []int(word)
    rune := runes[index]

    if node == nil {
        node = NewNode(rune)
    }

    if rune == node.Rune {
        if (index + 1) < len(runes) {
            node.Child = insert(node.Child, word, index + 1)
        } else {
            node.Word = runes
        }
    } else if rune < node.Rune {
        node.Smaller = insert(node.Smaller, word, index)
    } else {
        node.Larger = insert(node.Larger, word, index)
    }
    return node
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

func (n *Node) isEndOfWord() {
    return n.Word != nil
}

