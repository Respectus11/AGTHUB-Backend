package main
import (
    "fmt" 
)
func main() {
    var name string
    var nosubjects int
    fmt.Print("what is your name? ")
    fmt.Scanln(&name) 
    fmt.Print("how many subjects do u have taken")
    fmt.Scanln(&nosubjects) 
    subjectGrades := make(map[string]float64)
    for i := 0; i < nosubjects; i++ {
        var sub string
        var grade float64
        fmt.Printf("Enter the name of subject #%d ", i+1)
        fmt.Scanln(&sub) 
        for {
            fmt.Printf("Enter the grade  ")
            _, err := fmt.Scanln(&grade) 
            if err != nil || grade < 1 || grade > 100 {
                fmt.Println("plese insert correct grade 1-100")
            } else {
                break
            }
        }
        subjectGrades[sub] = grade
    }
    average := calculateAverage(subjectGrades)
    fmt.Printf("\nGrade Report for %s:\n", name)
    for subject, grade := range subjectGrades {
        fmt.Printf("- %s: %.2f\n", subject, grade)
    }
    fmt.Printf("Average Grade: %.2f\n", average)
}
func calculateAverage(grades map[string]float64) float64 {
    var total float64
    for _, grade := range grades {
        total += grade
    }
    return total / float64(len(grades))
}
