package main

import (
	"fmt"
	"week1/bank"
	//"week1/bank"
)

func main() {
	//--------------------------- dag1--------------------------
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
	// var summa float64
	// var harRabatt string

	// fmt.Print("Mata in summan :")
	// fmt.Scan(&summa)
	// fmt.Print("Har du rabat?y/n :")
	// fmt.Scan(&harRabatt)

	// harRabatt = strings.ToLower(harRabatt)

	// if summa >= 15 && summa <= 25 {
	// 	if harRabatt == "ja" {
	// 		fmt.Println("Du kan köpa en liten hamburgare och en pommes frites.")
	// 	} else {
	// 		fmt.Println("Du kan köpa en liten hamburgare.")
	// 	}
	// } else if summa > 25 && summa <= 50 {
	// 	if harRabatt == "ja" {
	// 		fmt.Println("Du kan köpa en stor hamburgare och pommes frites.")
	// 	} else {
	// 		fmt.Println("Du kan köpa en stor hamburgare.")
	// 	}
	// } else if summa > 50 && summa <= 60 && harRabatt == "ja" {
	// 	fmt.Println("Du kan köpa ett meal med dryck.")
	// } else if summa > 60 {
	// 	fmt.Println("Du kan köpa ett meal med dryck.")
	// } else {
	// 	fmt.Println("Tyvärr, du har inte tillräckligt med pengar.")
	// }

	// 	------------LOOPAR
	// 1. LOOP #1
	// Skapa ett program som skriver ut talen 0-10 på skärmen. Lös detta med en loop.
	// for i := 0; i < 11; i++ {
	// 	fmt.Println(i)
	// }

	// 2. LOOP #2
	// Skapa ett program där användaren får mata in två tal. Låt sedan programmet skriva ut alla
	// tal som finns mellan dessa två tal på skärmen. Lös detta med en loop.
	// var no1, no2 int
	// fmt.Println("---Mata in två tal---")
	// fmt.Print("Första tal :")
	// fmt.Scan(&no1)
	// fmt.Print("Andra tal :")
	// fmt.Scan(&no2)

	// if no1 < no2 {
	// 	for i := no1 + 1; i < no2; i++ {
	// 		fmt.Println(i)
	// 	}
	// } else if no1 > no2 {
	// 	for i := no1 - 1; i > no2; i-- {
	// 		fmt.Println(i)
	// 	}
	// } else {
	// 	fmt.Println("Talet är lika!")
	// }

	// 3. LOOP #3
	// Skapa ett program där användaren
	// Får mata in två tal.
	// Skriv sedan till skärmen summan av de två talen.
	// Skriv sedan en fråga- Vill du fortsätta(J/N)?.
	// Svarar användaren J återupprepas punkt a-c
	// Svarar användaren N avbryts programmet

	// var fortsätta bool = true
	// for fortsätta {
	// 	var no1, no2 int
	// 	fmt.Println("---Mata in två tal---")
	// 	fmt.Print("Första tal :")
	// 	fmt.Scan(&no1)
	// 	fmt.Print("Andra tal :")
	// 	fmt.Scan(&no2)
	// 	fmt.Printf("Summan av två talet är %v+%v=%v\n", no1, no2, no1+no2)
	// 	var svaret string
	// 	fmt.Print("Vill du fortsätta programmet?y/n :")
	// 	fmt.Scan(&svaret)
	// 	svaret = strings.ToLower(svaret)
	// 	if svaret == "y" {
	// 		fortsätta = true
	// 		continue
	// 	} else {
	// 		fortsätta = false
	// 		break
	// 	}
	// }

	// 4. LOOP #4
	// Be användaren mata in ett tal. Spara värdet i en variabel. Upprepa detta 10 gånger. För
	// varje gång lägg till det inmatade värdet till variabelns värde. När det är klart skriv ut-
	// Summan av det du matat in är: summan.
	// var no int
	// fmt.Print("Mata in ett tal :")
	// fmt.Scan(&no)
	// var summa int
	// for i := 0; i < 10; i++ {
	// 	summa += no
	// 	fmt.Println(summa)
	// }
	// fmt.Printf("Summan av det du matat in är: %v\n", summa)

	// 5. LOOP #5
	// Skapa ett program där användaren får mata in ett tal. Låt sedan programmet skriva ut
	// alla siffor som är mindre än det inmatade talet men större än 0. Lös detta med en
	// loop.
	// var no int
	// fmt.Print("Mata in ett tal :")
	// fmt.Scan(&no)
	// for i := no - 1; i > 0; i-- {
	// 	fmt.Println(i)
	// }

	// 6. LOOP #7 (Överkurs)
	// Rolling the dice
	// Kasta två tärningar” och visa resultatet enligt skärmdump ända tills man INTE svarar ”y” eller ”yes” på frågan om igen
	// var fortsätta bool = true
	// fmt.Println("Welcome to Rooling the dice")
	// for fortsätta {
	// 	randomNo1 := rand.IntN(7)
	// 	if randomNo1 == 0 {
	// 		randomNo1 = 1
	// 	}
	// 	randomNo2 := rand.IntN(7)
	// 	if randomNo2 == 0 {
	// 		randomNo2 = 1
	// 	}
	// 	fmt.Printf("\n\nYour first dice is %v and second dice is %v\n\n", randomNo1, randomNo2)
	// 	if randomNo1 == randomNo2 {
	// 		fmt.Println("You are lucky! You have found 2 same number!")
	// 	}
	// 	if randomNo1 == 6 && randomNo2 == 6 {
	// 		fmt.Println("You found 6x6! Lets rock!")
	// 	}
	// 	var svaret string
	// 	fmt.Print("Vill du fortsätta programmet?y/n :")
	// 	fmt.Scan(&svaret)
	// 	svaret = strings.ToLower(svaret)
	// 	if svaret == "y" {
	// 		fortsätta = true
	// 		continue
	// 	} else {
	// 		fortsätta = false
	// 		break
	// 	}
	// }

	//--------------------------- dag2--------------------------
	//funkitoner

	// add(5, 5)

	// fmt.Println("Area: ", rectangle(20, 30))

	// var a, p int
	// a, p = rectangle(20, 30)
	// fmt.Println("Area:", a)
	// fmt.Println("Parameter:", p)

	balance := 6000
	belopp := 1555
	errorText := ""
	balance, errorText = bank.Withdraw(balance, belopp)
	if len(errorText) > 0 {
		fmt.Println(errorText)
	}
	fmt.Println(balance)

	// arr1 := [6]int{10, 11, 12, 13, 14, 15}
	// myslice1 := make([]int, 5, 10)
	// fmt.Println(arr1)
	// fmt.Println(myslice1)
	//FindBiggestNumber()
	LongestName()
}

//	func add(x int, y int) {
//		total := 0
//		total = x + y
//		fmt.Println(total)
//	}
//
//	func rectangle(l int, b int) (area int) {
//		var parameter int
//		parameter = 2 * (l + b)
//		fmt.Println("Parameter: ", parameter)
//		area = l * b
//		return // Return statement without specify variable name
//	}
//
//	func rectangle(l int, b int) (area int, parameter int) {
//		parameter = 2 * (l + b)
//		area = l * b
//		return // Return statement without specify variable name
//	}

// func printMenu() {
// 	fmt.Println("\n1. Skapa device :")
// 	fmt.Println("2. Lista Alla :")
// 	fmt.Println("3. Ändra device :")
// 	fmt.Println("4. Sök :")
// 	fmt.Println("5. Avsluta :")
// }

// func createDevice(name string) (errorText string) {

// }
