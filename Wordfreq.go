package main
import (
    "fmt"
)
func main() {
    fmt.Print("Input a string: ")
    var word string
    fmt.Scanln(&word)
    dic := make(map[string]int)
    for _, ch := range word {
        char := string(ch)
        dic[char]++
    }

    fmt.Println("Character freq ")
    for letter, freq := range dic {
        fmt.Printf("%s: %d\n", letter, freq)
    }
}
