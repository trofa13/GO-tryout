package main


import (
	"fmt"
	"sync"
	"math"
)



// use in your shell `time GOMAXPROCS=1 go run main.go` to measure time on a single thread
// use in your shell `time go run main.go` to feel the power of concurrency :)
func main() {
	names := []string{"Alex", "Valera", "Antoha"}
	
	var wg sync.WaitGroup
	wg.Add(len(names))
	for _, name := range names{
		go printName(name, &wg)
	}
	wg.Wait()
}

func printName(n string, wg *sync.WaitGroup){
	result := 0.0
	for i:= 0; i < 100000000; i++ {
		result += math.Pi * math.Sin(float64(len(n)))
	}
	fmt.Println("Name: ", n)
	wg.Done()
}



//real time - 9.153 (no concurrency)
//uncomment everything below to see results with no concurrency on your machine
// use in your shell `time go run main.go` to measure time
/*
func main() {
	names := []string{"Alex", "Valera", "Antoha"}

	for _, name := range names{
		printName(name)
	}
}



func printName(n string){
	result := 0.0
	for i:= 0; i < 100000000; i++ {
		result += math.Pi * math.Sin(float64(len(n)))
	}
	fmt.Println("Name: ", n)
}
*/