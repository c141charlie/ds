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

type CallCenter struct {
    agents int
    queue chan *Call
    clock_in chan bool
    clock_out chan bool
}

func (cc *CallCenter) Open() {
    fmt.Println("Call center opening")
    for i:=0; i < cc.agents; i++ {
        go agent(i, cc.clock_in, cc.queue, cc.clock_out)
    }
    for b:=0; b < cc.agents; b++ {
        <- cc.clock_in
    }
    fmt.Println("Call center open")
}

func (cc *CallCenter) Close() {
    fmt.Println("Call center closing")
    
    for i:=0; i < cc.agents; i++ {
        cc.queue <- &Call{-1, 1, time.Nanoseconds()}
    }

    for d := 0; d < cc.agents; d++ {
        <- cc.clock_out
    }

    fmt.Println("Call center closed")
}

func agent(id int, clock_in chan bool, queue chan *Call, clock_out chan bool) {
    fmt.Printf("Agent %d logging in\n", id)
    clock_in <- true
    for {
        call := <- queue
        if call.id == -1 {
            fmt.Printf("Agent %d going home\n", id)
            clock_out <- true
            break
        }
        time.Sleep(int64(call.duration * 1000000))
        fmt.Printf("Agent %d, answering Call %d\n", id, call.id)
        fmt.Printf("Call %d answered; waited %d milliseconds\n", call.id, (time.Nanoseconds() - call.start_time)/1000000)
    }
}

func main() {
    runtime.GOMAXPROCS(2) //Dual core
    args := flag.Args()
    
    if len(args) < 4 {
        fmt.Println("Usage: CallGenerator <numberOfAgents> <numberOfCalls> <maxCallDuration> <maxCallInterval>")
        return
    }
    num_agents, _ := strconv.Atoi(args[0])
    num_calls, _ := strconv.Atoi(args[1])
    max_dur, _ := strconv.Atoi(args[2])
    max_int, _ := strconv.Atoi(args[3])

    queue := make(chan *Call, 10000)
    clock_in := make(chan bool)
    clock_out := make(chan bool)

    call_center := &CallCenter{num_agents, queue, clock_in, clock_out}
    
    begin := time.Nanoseconds()

    call_center.Open()
    
    for i:=0; i<num_calls; i++ {
        time.Sleep(int64(max_int))
        queue <- &Call{i,rand.Int() * max_dur, time.Nanoseconds()}
    }

    call_center.Close()

    end := time.Nanoseconds()

    fmt.Printf("Simulation elapsed time: %d seconds\n", (end - begin)/1000000000)
}
