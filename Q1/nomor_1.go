package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	struck, err := os.Create("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Input Resto Name : ")
	resto, _ := reader.ReadString('\n')
	resto = strings.TrimSuffix(resto, "\n")
	// fmt.Println(resto)
	dateNow := time.Now()
	// fmt.Println(dateNow)
	fmt.Print("Enter Name Casier : ")
	casier, _ := reader.ReadString('\n')
	casier = strings.TrimSuffix(casier, "\n")
	// fmt.Println(casier)

	var pilih bool = true

	var itemPrice map[string]int64
	itemPrice = make(map[string]int64)
	for pilih {
		fmt.Print("Input name item : ")
		item, _ := reader.ReadString('\n')
		item = strings.TrimSuffix(item, "\n")
		fmt.Print("Input Price : ")
		var price int64
		_, err := fmt.Scanf("%d", &price)
		if err != nil {
			fmt.Println(err)
		}
		itemPrice[item] = price

		fmt.Println("Pesan Lagi (Y/N) : ")
		pesan, _ := reader.ReadString('\n')
		pesan = strings.TrimSuffix(pesan, "\n")
		if strings.ToLower(pesan) == "n" {
			pilih = false
		}
	}
	var titik string = "............"
	var tot int64 = 0
	restoText := fmt.Sprintf("%10s", resto)
	_, err = struck.Write([]byte(restoText + "\n"))
	if err != nil {
		log.Fatal(err)
	}
	var tanggal string = "Tanggal : " + dateNow.String()
	_, err = struck.Write([]byte(tanggal + "\n"))
	if err != nil {
		log.Fatal(err)
	}
	kasir := "Nama Kasir : " + casier
	_, err = struck.Write([]byte(kasir + "\n"))
	if err != nil {
		log.Fatal(err)
	}
	_, err = struck.Write([]byte("=================" + "\n"))
	if err != nil {
		log.Fatal(err)
	}
	for k, m := range itemPrice {
		batas := k + titik + "Rp " + strconv.FormatInt(m, 10)
		if len(batas) > 30 {
			_, err = struck.Write([]byte(k + titik + "\n"))
			if err != nil {
				log.Fatal(err)
			}
			for i := 0; i < 25; i++ {
				_, err = struck.Write([]byte("."))
				if err != nil {
					log.Fatal(err)
				}
			}
			_, err = struck.Write([]byte("Rp " + strconv.FormatInt(m, 10) + "\n"))
			if err != nil {
				log.Fatal(err)
			}
		} else {
			_, err = struck.Write([]byte(batas + "\n"))
			if err != nil {
				log.Fatal(err)
			}
		}
		tot += m
	}
	batas := "total" + titik + "Rp " + strconv.FormatInt(tot, 10)
	if len(batas) > 30 {
		_, err = struck.Write([]byte("total" + titik + "\n\n"))
		if err != nil {
			log.Fatal(err)
		}
		for i := 0; i < 25; i++ {
			_, err = struck.Write([]byte("."))
			if err != nil {
				log.Fatal(err)
			}
		}
		_, err = struck.Write([]byte("Rp" + strconv.FormatInt(tot, 10) + "\n"))
		if err != nil {
			log.Fatal(err)
		}
	} else {
		_, err = struck.Write([]byte(batas + "\n"))
		if err != nil {
			log.Fatal(err)
		}
	}
}
