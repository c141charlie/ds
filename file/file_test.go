package file

import (
    "testing"
)

func TestFileReader(t *testing.T) {
    
    file, err := Open("/Users/c141charlie/Code/go/reco/data/data.txt", "r")

    if err != nil {
        t.Errorf("err should be nil, but is: %s\n", err)
    }


    for file.IsDone() != true {
        
        line := file.ReadLine()
        if line != "field1  field2  field3" {
            t.Errorf("line should be \"field1 field2 filed3\", but is \"%s\"\n", line)
        }
    }
    
    file.Close()
}

func TestFileReaderWithEmptyFile(t *testing.T) {

    file, err := Open("/Users/c141charlie/Code/go/reco/data/empty_file.txt", "r")

    if err != nil {
        t.Errorf("err should be nil, but is: %s\n", err)
    }

    if file.IsDone() != true {
        t.Errorf("file.IsDone() should be true.")
    }

    if file.ReadLine() != "" {
        t.Errorf("file.ReadLine should be \"\"\n")
    }

    file.Close()
}
