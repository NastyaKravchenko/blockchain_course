package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

func main() {
	var n float64
	fmt.Print("Введіть довжину ключа (біт): ")
	fmt.Scan(&n)                                                       //n - довжина бітової послідовності
	nk1 := big.NewInt(0).Exp(big.NewInt(2), big.NewInt(int64(n)), nil) //розрахунок кількості варіантів ключів
	fmt.Println("Кількість ключів: ", nk1)
	nn3 := big.NewInt(0).Exp(big.NewInt(2), big.NewInt(int64(n-5)), nil)  //допоміжні розрахунки - 3.125%
	nn25 := big.NewInt(0).Exp(big.NewInt(2), big.NewInt(int64(n-2)), nil) //25%
	nn50 := big.NewInt(0).Exp(big.NewInt(2), big.NewInt(int64(n-1)), nil) //50%
	nn75 := big.NewInt(0).Sub(nk1, nn25)                                  //75% від кількості ключів
	fmt.Println(nn25, nn50, nn75)
	key, err := rand.Int(rand.Reader, nk1) //генерація рандомного ключа з діапазону
	fmt.Println("Kлюч: ", key)             //та його вивід у двох виглядах
	fmt.Println("      ", fmt.Sprintf("0x%x", key))

	t0 := time.Now()
	i := big.NewInt(0)
	j := big.NewInt(1)

	for {
		if fmt.Sprintf("%v", i) == fmt.Sprintf("%v", nn3) {
			tt1 := time.Now()
			fmt.Printf("\nрозрахований (дууже приблизний) час для перебору усіх ключів: %v м\n\n", tt1.Sub(t0).Minutes()*32)
		} //фрагмент для розрахунку приблизного часу очікування, має сенс для великих послідовностей
		if fmt.Sprintf("0x%x", i) == fmt.Sprintf("0x%x", key) {
			fmt.Println("\nЗнайдений ключ:", i)
			break
			fmt.Println("   ", err)
		} //реалізація брутфорсу значень

		i = big.NewInt(0).Add(i, j)
		//наступні цикли - щоб не дуже сумно було чекати:)
		if fmt.Sprintf("%v", i) == fmt.Sprintf("%v", nn25) {
			fmt.Println("оброблено 25% усіх ключів...")
		}
		if fmt.Sprintf("%v", i) == fmt.Sprintf("%v", nn50) {
			fmt.Println("оброблено 50% усіх ключів...")
		}
		if fmt.Sprintf("%v", i) == fmt.Sprintf("%v", nn75) {
			fmt.Println("оброблено 75% усіх ключів...")
		}
	}
	t1 := time.Now()
	fmt.Printf("Витрачено часу: %v мс", t1.Sub(t0).Milliseconds()) //вивід затраченого часу

}
