package screen

import (
	"fmt"
	"strings"
)

func PrintStatus(name string) string {
	var menuText strings.Builder

	fmt.Fprintln(&menuText, "\nMy name is", name)
	fmt.Fprintln(&menuText, "What do you want me to do?")
	fmt.Fprintln(&menuText, "1. Exercise")
	fmt.Fprintln(&menuText, "2. Sleep")
	fmt.Fprintln(&menuText, "3. Study")
	fmt.Fprintln(&menuText, "4. Eat")
	fmt.Fprintln(&menuText, "5. Show status")
	fmt.Fprintln(&menuText, "6. Exit")

	return menuText.String()
}
