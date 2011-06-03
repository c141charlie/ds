package ternary

import "testing"

var tree *TernaryTree = NewTernaryTree()

func setup() {
    tree.Add("car")
}

func TestContains(t *testing.T) {
    setup()
    if tree.Contains("car") == false {
        t.Errorf("tree should contain \"car\"")
    }
}






/*
func setup() {
    tree.Add("prefabricate")
    tree.Add("presume")
    tree.Add("prejudice")
    tree.Add("preliminary")
    tree.Add("apple")
    tree.Add("ape")
    tree.Add("appeal")
    tree.Add("car")
    tree.Add("dog")
    tree.Add("cat")
    tree.Add("mouse")
    tree.Add("mince")
    tree.Add("minty")
}

func TestContains(t *testing.T) {
    
    setup()

    if tree.Contains("prefabricate") == false {
        t.Errorf("tree.Contains(\"prefabricate\") should be true\n")
    }
    if tree.Contains("presume") == false {
        t.Errorf("tree.Contains(\"presume\") should be true\n")
    }
    if tree.Contains("prejudice") == false {
        t.Errorf("tree.Contains(\"prejudice\") should be true\n")
    }
    if tree.Contains("preliminary") == false {
        t.Errorf("tree.Contains(\"preliminary\") should be true\n")
    }
    if tree.Contains("apple") == false {
        t.Errorf("tree.Contains(\"apple\") should be true\n")
    }
    if tree.Contains("ape") == false {
        t.Errorf("tree.Contains(\"ape\") should be true\n")
    }
    if tree.Contains("appeal") == false {
        t.Errorf("tree.Contains(\"appeal\") should be true\n")
    }
    if tree.Contains("car") == false {
        t.Errorf("tree.Contains(\"car\") should be true\n")
    }
    if tree.Contains("dog") == false {
        t.Errorf("tree.Contains(\"dog\") should be true\n")
    }
    if tree.Contains("cat") == false {
        t.Errorf("tree.Contains(\"cat\") should be true\n")
    }
    if tree.Contains("mouse") == false {
        t.Errorf("tree.Contains(\"mouse\") should be true\n")
    }
    if tree.Contains("mince") == false {
        t.Errorf("tree.Contains(\"mince\") should be true\n")
    }
    if tree.Contains("minty") == false {
        t.Errorf("tree.Contains(\"minty\") should be true\n")
    }
}

func TestPrefixSearch(t *testing.T) {
    fmt.Println("hello")
}

*/
