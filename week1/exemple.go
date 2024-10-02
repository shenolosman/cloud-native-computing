package main

import "fmt"

func FindBiggestNumber() {
	var numbers [4]int // Array för att spara fyra tal

	// Fråga användaren att mata in fyra tal
	for i := 0; i < 4; i++ {
		fmt.Printf("Mata in tal #%d: ", i+1)
		fmt.Scan(&numbers[i])
	}

	// Hitta det största talet
	storsta := numbers[0] // Anta att första talet är störst till en början
	for i := 1; i < len(numbers); i++ {
		if numbers[i] > storsta {
			storsta = numbers[i]
		}
	}

	// Skriv ut det största talet
	fmt.Printf("Det största talet är: %d\n", storsta)
}

func LongestName() {
	var names []string // Slice för att spara namn
	var continueInput string

	for {
		// Be användaren mata in ett namn
		var name string
		fmt.Print("Mata in ett namn: ")
		fmt.Scan(&name)

		// Lägg till namnet i slicen
		names = append(names, name)

		// Fråga om användaren vill fortsätta
		fmt.Print("Vill du fortsätta? (y/n): ")
		fmt.Scan(&continueInput)

		// Om användaren svarar 'n', avsluta loopen
		if continueInput == "n" || continueInput == "N" {
			break
		}
	}

	// Hitta det namn som har flest tecken
	longestName := ""
	for _, name := range names {
		if len(name) > len(longestName) {
			longestName = name
		}
	}

	// Skriv ut det längsta namnet
	fmt.Printf("Det längsta namnet är: %s\n", longestName)
}
