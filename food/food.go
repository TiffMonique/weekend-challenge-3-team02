package food

type Food interface {
	GetName() string
	GetTaste() string
	GetEffects() Effects
}
type Effects struct {
	Hunger     int
	Stamina    int
	Sleepiness int
}
type FoodItem struct {
	name    string
	taste   string
	effects Effects
}
type Dessert struct {
	FoodItem
}

type Meal struct {
	FoodItem
}

type Energizer struct {
	FoodItem
}

func (d FoodItem) GetName() string {
	return d.name
}

func (d FoodItem) GetTaste() string {
	return d.taste
}

func (d FoodItem) GetEffects() Effects {
	return d.effects
}

func NewFoodItem(name, taste string, hunger, stamina, sleepiness int) FoodItem {
	return FoodItem{
		name:  name,
		taste: taste,
		effects: Effects{
			Hunger:     hunger,
			Stamina:    stamina,
			Sleepiness: sleepiness,
		},
	}
}

func NewDessert(name, taste string) Dessert {
	return Dessert{FoodItem: NewFoodItem(name, taste, -20, 10, 0)}
}

func NewMeal(name, taste string) Meal {
	return Meal{FoodItem: NewFoodItem(name, taste, -50, 25, 10)}
}

func NewEnergizer(name, taste string) Energizer {
	return Energizer{FoodItem: NewFoodItem(name, taste, 10, 50, -50)}
}
