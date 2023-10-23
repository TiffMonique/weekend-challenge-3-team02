package main

import (
	"fmt"

	"trucode.com/interactions/food"
	humans "trucode.com/interactions/person"
	"trucode.com/interactions/screen"
)

func main() {
	choice := 0

	fmt.Println(screen.PrintMenu())
	if _, err := fmt.Scanln(&choice); err != nil {
		choice = 0
	}

	var person *humans.Person

	fmt.Println("Enter the name of the person: ")
	var name string
	fmt.Scanln(&name)

	if choice == 1 {

		person = humans.NewPerson(name)
		fmt.Println(screen.PrintStatus(name))
		InteractWithPerson(person)

	} else if choice == 2 {
		child := humans.NewChild(name)

		fmt.Println(screen.PrintStatus(name))
		InteractWithPerson(child)
	} else {
		println("Action not allowed, exiting")
	}

}

func InteractWithPerson(person humans.Human) {
	option := 0
	fmt.Scanln(&option)

	switch option {
	case 1:
		fmt.Println(screen.PrintExercise())

		var intensity int
		fmt.Scanln(&intensity)
		switch intensity {
		case 1:
			person.Exercising(humans.LOW_INTENSITY)
		case 2:
			person.Exercising(humans.MEDIUM_INTENSITY)
		case 3:
			person.Exercising(humans.HIGH_INTENSITY)
		default:
			fmt.Println("Invalid intensity choice.")
		}
	case 2:
		person.Sleep()
	case 3:
		person.Study()
	case 4:
		fmt.Println("Which thing should I eat.")
		fmt.Println("1. dessert")
		fmt.Println("2. meal")
		fmt.Println("3. energizer")
		var foodChoice int
		fmt.Scanln(&foodChoice)

		if foodChoice < 1 || foodChoice > 3 {
			fmt.Println("Invalid food choice.")
			return
		}

		fmt.Println("Enter the name of the food: ")
		var foodName string
		fmt.Scanln(&foodName)

		fmt.Println("Enter the taste of the food: ")
		var foodTaste string
		fmt.Scanln(&foodTaste)

		var foodItem food.Food

		if foodChoice == 1 {
			foodItem = food.NewDessert(foodName, foodTaste)
		} else if foodChoice == 2 {
			foodItem = food.NewMeal(foodName, foodTaste)
		} else {
			foodItem = food.NewEnergizer(foodName, foodTaste)
		}
		person.Eat(foodItem)

	case 5:
		if p, ok := person.(*humans.Person); ok {
			humans.ShowStatus(p)
		} else {
			fmt.Println("Invalid type for ShowStatus")
		}
	case 6:
		fmt.Println("Goodbye!")
		return
	default:
		fmt.Println("Invalid choice. Try again.")
	}

}
