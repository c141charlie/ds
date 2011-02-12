package main

import "fmt"

func SayHello() {
	fmt.Println("Hello")
}

func SayHelloToSomeone(name string) {
	fmt.Println("Hello " + name + ".")
}

/*
func GetGreeting (name string) string {
    greeting := "Hello " + name + "."
    return greeting
}
*/

func GetGreeting (name string) (greeting string) {
    greeting = "Hello " + name + "."
    return
}

type TormentList struct {
    people []string
}

func NewTormentList(people []string) *TormentList {
    return &TormentList{people}
}

func (g *TormentList) Map(f(func(string))) {
    for _, val := range(g.people) {
        f(val)
    }
}

func Screen(people []string) func(string) bool {
    return func(name string) bool {
        for _, person := range people {
            if person == name {
                return true
            }
        }
        return false
    }
}

func main() {
    greeting := GetGreeting("Bob")
    fmt.Println(greeting) //outputs "Hello Bob."

    people := []string{"Anand", "David", "Ivan", "JoJo", "Jin", "Mon", "Peter", "Sachin"}
    tl := NewTormentList(people)
    tl.Map(SayHelloToSomeone)

    those_who_bought_sushi := []string{"Anand", "JoJo", "Jin", "Mon", "Peter", "Sachin"}

    bought_sushi := Screen(those_who_bought_sushi)

    fmt.Println(bought_sushi("Anand")) // true
    fmt.Println(bought_sushi("Alex")) // false

}
