package main
import ("fmt" 
"strings")
func main(){
	fmt.Print("input a word to check it is palindrome \n")
	var word string
	fmt.Scanln(&word)
	word = strings.ToLower(word) 

	reverse:=""
	for i := len(word) - 1; i >= 0; i-- {
        reverse += string(word[i])
    }
	if word == reverse {
        fmt.Println("palindrome")
    } else {
        fmt.Println("not palindrome")
    }
}

	// if word==[:-1]word{
	// 	fmt.print("it is palindrome")
	// else{
	// 	fmt.print("it is not palindrome")
	// }
	// }
