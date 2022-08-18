package main

import (
	"encoding/json"
	"fmt"
	"github.com/dustin/go-humanize"
	"log"
	"strings"
)

func main() {
	var amount int64
	fmt.Print("Masukan jumlah uang: ")
	fmt.Scanf("%d", &amount)

	countCurrency(amount)
}

func countCurrency(amount int64) {
	notes := []int64{100000, 50000, 20000, 10000, 5000, 2000, 1000, 500, 200, 100}
	noteDeno := map[string]int64{}
	for _, note := range notes {
		if amount >= note {
			denoCount := amount / note
			value := humanize.Comma(note)
			res := strings.Replace(value, ",", ".", -1)
			noteDeno[fmt.Sprintf("Rp. %s", res)] = denoCount
			amount = amount % note
			if amount >= 1 && amount <= 99 {
				amount = 100
			}
		}
	}

	jsonBody, err := json.Marshal(noteDeno)
	if err != nil {
		log.Println(err)
	}

	log.Println(string(jsonBody))

}
