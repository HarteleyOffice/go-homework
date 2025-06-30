package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	for {
		harga := inputHarga()
		lama_pinjaman := inputLama()

		sewa := harga * 0.1

		admin := harga * 0.05
		if admin < 20000 {
			admin = 20000
		}
		admin = bulatkanAdmin(admin)

		fmt.Println()

		fmt.Println("===Detail Biaya===")

		fmt.Printf("Harga: %.0f\n", harga)

		fmt.Printf("Admin: %.0f\n", admin)

		fmt.Println("Pinjaman", lama_pinjaman, "hari")
		total_sewa := float64(0)

		sewa_ke := lama_pinjaman / 30
		remainder := lama_pinjaman % 30
		for i := 1; i <= sewa_ke; i++ {
			fmt.Printf("Sewa ke-%d: %.0f\n", i, sewa)
			total_sewa = total_sewa + sewa
		}
		if remainder > 0 {
			if remainder < 16 {
				fmt.Printf("Sewa ke-%d: %.0f\n", sewa_ke+1, sewa*0.5)
				total_sewa = total_sewa + (sewa * 0.5)
			} else {
				fmt.Printf("Sewa ke-%d: %.0f\n", sewa_ke+1, sewa)
				total_sewa = total_sewa + sewa
			}
		}
		fmt.Printf("Akumulasi sewa: %.0f\n", total_sewa)

		potongan := 0.0
		useVoucher := inputUseVoucher()
		if useVoucher {
			potongan = total_sewa * 0.5
			if potongan > 30000 {
				potongan = 30000
			}
			fmt.Printf("Terdapat potongan sewa: %.0f\n", potongan)
		}
		fmt.Println()
	}
}

func inputHarga() float64 {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Input Harga: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		harga, err := strconv.ParseFloat(input, 64)
		if err != nil || harga < 0 {
			fmt.Println("Invalid input harga")
			continue
		}

		return harga
	}
}

func inputLama() int {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Input Lama Pinjaman: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		lama_pinjaman, err := strconv.Atoi(input)
		if err != nil || lama_pinjaman < 0 {
			fmt.Println("Invalid input lama pinjaman")
			continue
		}

		return lama_pinjaman
	}
}

func bulatkanAdmin(admin float64) float64 {
	x := int(admin) / 10000
	remainder := int(admin) % 10000

	if remainder < 2500 {
		remainder = 0
	} else if remainder < 7500 {
		remainder = 5000
	} else {
		remainder = 10000
	}

	admin = float64(x*10000 + remainder)

	return admin
}

func inputUseVoucher() bool {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Voucher potongan Tarif Sewa: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))

		if input == "true" {
			return true
		} else if input == "false" {
			return false
		} else {
			fmt.Println("Input tidak valid")
		}
	}
}
