package main

import(
    "fmt"
    "time"
    "rand"
    "flag"
    "strconv"
    "runtime"
)

type Call struct {
    id int
    duration int
    start_time int64
}

func agent(id int, call *Call) {
    time.Sleep(int64(call.duration * 1000000))
    fmt.Printf("Agent %d, answering Call %d\n", id, call.id)
    fmt.Printf("Call %d answered; waited %d milliseconds\n", 
    call.id, (time.Nanoseconds() - call.start_time)/1000000)
}

func main() {
    runtime.GOMAXPROCS(2) //Dual core
    args := flag.Args()
    
    if len(args) < 3 {
        fmt.Println("Usage: CallGenerator <numberOfCalls> <maxCallDuration> <maxCallInterval>")
        return
    }
    num_calls, _ := strconv.Atoi(args[0])
    max_dur, _ := strconv.Atoi(args[1])
    max_int, _ := strconv.Atoi(args[2])

    for i:=0; i<num_calls; i++ {
        time.Sleep(int64(max_int))
        go agent(i, &Call{i,rand.Int() * max_dur, time.Nanoseconds()})
    }
}
