package bank

func Withdraw(currentBalance int, belopp int) (newBalance int, errorText string) {
	errorText = ""
	newBalance = currentBalance
	if belopp > currentBalance {
		errorText = "Så mycket pengar har du inte"
	} else if belopp > 3000 {
		errorText = "Maxbelopp 3000 per uttag"
	} else {
		newBalance = currentBalance - belopp
	}
	return
}
