package person

import (
	"fmt"

	food "trucode.com/interactions/food"
)

type Human interface {
	Sleep()
	Study()
	Eat(food.Food)
	Exercising(intensity string)
}
type Person struct {
	Name       string
	Hunger     int
	Sleepiness int
	Stamina    int
}
type Child struct {
	Person
}

func NewPerson(name string) *Person {
	p := &Person{Name: name, Hunger: 80, Sleepiness: 10, Stamina: 80}
	return p
}

func NewChild(name string) *Child {
	c := &Child{Person: Person{Name: name, Hunger: 100, Sleepiness: 10, Stamina: 80}}
	return c
}

func ShowStatus(p *Person) {
	fmt.Printf("My name is %s, this is my status:\n", p.Name)
	fmt.Printf("Hunger: %d\nSleepiness: %d\nStamina: %d\n", p.Hunger, p.Sleepiness, p.Stamina)
}

const (
	LOW_INTENSITY    = "low"
	MEDIUM_INTENSITY = "medium"
	HIGH_INTENSITY   = "high"
)

func (p *Person) Sleep() {
	fmt.Println("I have slept.")
	p.Hunger += 20
	p.Sleepiness = 0
	p.Stamina = 100
}

func (p *Person) Study() {
	fmt.Println("I have studied.")
	p.Hunger += 25
	p.Sleepiness += 30
	p.Stamina -= 10
}

func (p *Person) Eat(f food.Food) {
	foodEffects := f.GetEffects()
	p.Hunger += foodEffects.Hunger
	p.Sleepiness += foodEffects.Sleepiness
	p.Stamina += foodEffects.Stamina
	fmt.Printf("I have eaten. That %s tasted %s!\n", f.GetName(), f.GetTaste())
}

func (p *Person) Exercising(intensity string) {
	switch intensity {
	case LOW_INTENSITY:
		fmt.Println("I have exercised at low intensity.")
		p.Stamina -= 10
		p.Hunger += 10
	case MEDIUM_INTENSITY:
		fmt.Println("I have exercised at medium intensity.")
		p.Stamina -= 25
		p.Hunger += 30
	case HIGH_INTENSITY:
		fmt.Println("I have exercised at high intensity.")
		p.Stamina -= 50
		p.Hunger += 60
	default:
		fmt.Println("Invalid intensity level.")
	}
}
