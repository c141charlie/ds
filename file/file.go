package file

import (
    "os"
    "bufio"
    "strings"
)

type FileReader struct {
    file        *os.File
    reader      *bufio.Reader
    line        string
    is_done     bool
}

func Open(filename string, options string) (*FileReader, os.Error) {
    if options == "r" {

        file, err := load(filename)
        if err != nil {
            return nil, err 
        }
        
        rdr := bufio.NewReader(file)
        line := readstring(rdr)
        is_done := line == ""

        return &FileReader{file,rdr,line, is_done}, nil
    }
    new_error := os.NewError("invalid option")
    return nil, new_error

}

func (r *FileReader) ReadLine() string {
    
    line := r.line
    
    next_line := readstring(r.reader)
    if next_line == "" {
        r.is_done = true
    }
    r.line = next_line
    return line
}

func (r *FileReader) IsDone() bool {
    return r.is_done
}

func (r *FileReader) Close() {
    r.file.Close()
}

func load(filename string) (*os.File, os.Error) {

    file, err := os.Open(filename, os.O_RDONLY, 0666)
    if err != nil {
        return nil, os.NewError("File did not open.")
    }

    return file, nil
}

func readstring(r *bufio.Reader) string {
    str, _ := r.ReadString('\n')
    str = strings.Trim(str, "\n")
    return str
}
