package main

import (
	"fmt"
	"strings"
)

func main() {
	// var message string = "Hello Världen From Markaryd!"
	// var age int = 36

	// if strings.Contains(message, "Hello") {
	// 	fmt.Printf("my message is: %v and age is:%v\n", message, age)
	// }
	// stringarna := []string{"hello", "world"}
	// for _, stringen := range stringarna {
	// 	fmt.Println(stringen)
	// }
	// fmt.Printf("%v !\n", message)

	// print("2021\n")

	// print(2021)

	// print("2021-12-24\n")

	// print(2021 - 12 - 24)

	// 	4. För och efternamn
	// Skapa en applikation där användaren matar in för och efternamn.
	// var name string = ""
	// var lastName string = ""
	// var myAge int8 = 0
	// fmt.Print("Write your first name:")
	// fmt.Scan(&name)
	// fmt.Print("Write your last name:")
	// fmt.Scan(&lastName)
	// fmt.Print("Write your age name:")
	// fmt.Scan(&myAge)

	// fmt.Printf("You name is %v %v and %v years old\n", name, lastName, myAge)

	// 	7. Strängoperatorer
	// Be användaren mata sitt namn. Lägg det i en variabel som heter namn
	// Skapa en ny variabel namnUpprepat med namn fem gånger efter varann . Skriv ut den nya variabeln
	// 	name := ""
	// 	fmt.Print("Write your name :")
	// 	fmt.Scan(&name)

	// 	for i := 0; i < 5; i++ {
	// 		fmt.Print(name)
	// 	}
	// 	fmt.Print("\n")

	//1. IF #1
	// rightNumber := true
	// for rightNumber {
	// 	number := 0
	// 	fmt.Print("Write in a number :")
	// 	fmt.Scan(&number)
	// 	if number > 10 {
	// 		fmt.Print("Talet är större än 10!\n")
	// 		rightNumber = true
	// 		continue
	// 	} else if number == 10 {
	// 		fmt.Print("yay YOU hit the 10!\n")
	// 		rightNumber = false
	// 		break
	// 	} else {
	// 		fmt.Print("Talet är mindre än 10!\n")
	// 		rightNumber = true
	// 		continue
	// 	}
	// }

	//2. IF #2
	// paketMilk := 0
	// fmt.Print("Write in how many milk left :")
	// fmt.Scan(&paketMilk)

	// if paketMilk < 10 {
	// 	fmt.Print("Har beställt 30 paket mjölk!")
	// } else if paketMilk > 10 || paketMilk < 20 {
	// 	fmt.Print("Har beställt 20 paket mjölk!")
	// } else {
	// 	fmt.Print("Du behöver inte beställa någon mjölk")
	// }

	//3. if #3
	// var firstno, secondno int = 0, 0
	// fmt.Print("Mata in två nummer :")
	// fmt.Scan(&firstno, &secondno)
	// störstaNo := 0
	// if firstno > secondno {
	// 	störstaNo = firstno
	// } else {
	// 	störstaNo = secondno
	// }
	// fmt.Printf("Största talet är %v\n", störstaNo)

	// 4. IF #4
	// var firstno, secondno, tredjeno int = 0, 0, 0
	// fmt.Print("Mata in tre nummer :")
	// fmt.Scan(&firstno, &secondno, &tredjeno)
	// störstaNo := firstno // Assume first number is the largest

	// if secondno > störstaNo {
	// 	störstaNo = secondno
	// }

	// if tredjeno > störstaNo {
	// 	störstaNo = tredjeno
	// }
	// fmt.Printf("Största talet är %v\n", störstaNo)

	// 5. IF #5
	// väljKategori := [...]string{"vuxen", "pensionär", "student"}
	// var userInput string
	// fmt.Print("Skriva in vilken kategori tillhör du :")
	// found := false
	// fmt.Scan(&userInput)
	// for _, val := range väljKategori {
	// 	if strings.EqualFold(val, userInput) { // Case-insensitive comparison
	// 		fmt.Printf("Du har valt %v\n", val)
	// 		found = true
	// 		break
	// 	}
	// }

	// if !found {
	// 	fmt.Print("Du valde fel!\n")
	// }

	//6. IF #6
	// var födelsedag uint16 = 0
	// fmt.Print("Skriva in ditt födelsedag :")
	// fmt.Scan(&födelsedag)
	// if födelsedag >= 1980 && födelsedag < 1990 {
	// 	fmt.Print("Du är född på 1980-talet\n")
	// } else if födelsedag >= 1990 && födelsedag < 2000 {
	// 	fmt.Print("Du är född på 1990-talet\n")
	// } else if födelsedag < 1980 || födelsedag > 2000 {
	// 	fmt.Print("Du är inte född på 1990 eller 1980-talen\n")
	// }

	//7. IF #7
	// land := ""
	// fmt.Print("Mata in ditt land :")
	// fmt.Scan(&land)
	// scandinavian := [...]string{"Danmark", "Sweden", "Norge"}
	// for _, selectedLand := range scandinavian {
	// 	if strings.EqualFold(land, selectedLand) {
	// 		fmt.Print("Du bor i Skandinavien\n")
	// 		break
	// 	} else {
	// 		fmt.Print("Du är utanför Skandinavien\n")
	// 		break
	// 	}
	// }

	//8. IF #8
	var summa float64
	var harRabatt string

	fmt.Print("Mata in summan :")
	fmt.Scan(&harRabatt)
	fmt.Print("Har du rabat?y/n :")
	fmt.Scan(&harRabatt)

	harRabatt = strings.ToLower(harRabatt)

	if summa >= 15 && summa <= 25 {
		if harRabatt == "ja" {
			fmt.Println("Du kan köpa en liten hamburgare och en pommes frites.")
		} else {
			fmt.Println("Du kan köpa en liten hamburgare.")
		}
	} else if summa > 25 && summa <= 50 {
		if harRabatt == "ja" {
			fmt.Println("Du kan köpa en stor hamburgare och pommes frites.")
		} else {
			fmt.Println("Du kan köpa en stor hamburgare.")
		}
	} else if summa > 50 && summa <= 60 && harRabatt == "ja" {
		fmt.Println("Du kan köpa ett meal med dryck.")
	} else if summa > 60 {
		fmt.Println("Du kan köpa ett meal med dryck.")
	} else {
		fmt.Println("Tyvärr, du har inte tillräckligt med pengar.")
	}
}
