package main
import "fmt"



type Animal interface {
	Says() string
	NoOfLegs() int
}
type Dog struct {
	Name  string
	Breed string
}
func (d Dog) Says() string {
	return "woof"
}
func (d Dog) NoOfLegs() int {
	return 4
}
type Monkey struct {
	Name      string
	Color     string
	NoOfTeeth int
}
func (m Monkey) Says() string {
	return "khi khi khi"
}
func (m Monkey) NoOfLegs() int {
	return 2
}
func main() {
	dog := Dog{
		Name:  "Tomy",
		Breed: "GermanShepherd",
	}
	monkey := Monkey{
		Name:      "Chen",
		Color:     "brown",
		NoOfTeeth: 25,
	}
	PrintInfo(dog)
	PrintInfo(monkey)
	log.Println("hi")
	
	
}
func PrintInfo(a Animal) {
	fmt.Println("this animal says", a.Says(), "having", a.NoOfLegs(), "legs")
}
