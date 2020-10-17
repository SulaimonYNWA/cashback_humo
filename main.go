package main

	import(
		"fmt"
	)
func main() {
	purchases := [...]int{1,2,108,107}
	percent := 10
	limitCashback := 30
	cashback := 0
	sumPurchases := 0 
	for _, value := range purchases {
		//sum all purchases
		sumPurchases = value + sumPurchases
		} 
		cashback = sumPurchases * percent / 100
		if cashback > limitCashback{
			fmt.Println(limitCashback)
		} else if cashback == limitCashback{
			fmt.Println("cashback sum =", limitCashback)
		} else{
			fmt.Println("cashback sum =", cashback)
		}

	
}