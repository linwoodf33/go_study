package main
import (
	 "fmt"
	"linwood/calculator"
)

func main(){
	total :=calculator.Sum(4,5)
	fmt.Println(total)
	fmt.Println("Version:",calculator.logMessage)
}
