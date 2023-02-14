package main

import "fmt"

// secara default golang itu passing by value, bukan by reference

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{
		City:     "LEBAK",
		Province: "BANTEN",
		Country:  "INDONESIA",
	}
	// INI ADALAH KONDISI DIMANA KITA BIKIN PASS BY VALUE
	// DIMANA KITA BUAT VARIABLE ADDRESS 2 Berasal dari address1
	// dan ketika address2 diubah datanya, address1 tidak berubah
	address2 := address1
	address2.City = "SERANG"
	fmt.Println(address1)
	fmt.Println(address2)

	// POINTER adalah untuk kemampuan pass by reference
	adrdress3 := Address{
		City:     "CIREBON",
		Province: "JAWA BARAT",
		Country:  "INDONESIA",
	}
	adrdress4 := &adrdress3
	adrdress4.City = "BANDUNG"
	fmt.Println(adrdress3)
	fmt.Println(adrdress4)

	// ANOTHER WAY UNTUK SEBUAH POINTER
	var adrdress5 Address = Address{
		City:     "TEGAL",
		Province: "JAWA TENGAH",
		Country:  "INDONESIA",
	}
	var adrdress6 *Address = &adrdress5 // TANDA BINTANG DISINI SEBAGAI POINTER
	adrdress6.City = "SEMARANG"
	fmt.Println(adrdress5)
	fmt.Println(adrdress6)

	// jika kita ubah keynya saja yang di struct. maka dia akan berubah karena pointer. tapi kalo kita reasign itu semua maka address 1 tidak akan berubah
	// sebagai contoh
	// kondisi code dibawah ini sebagai berikut
	// kita bikin address 7, dan address8 pointer dari address 7. address8 kita ubah citynya menjadi ubud. maka, karena masih dalam satu pointer,
	// maka address 7 juga akan berubah menjadi ubud.
	// selanjutnya kita re-asign address 8 ek struct pointer address. maka dia akan memisakan diri dari address 7 membuat address yang baru. sehingga
	// address7 akan tetap di ubud, tapi address 8 menjadi struct baru
	var address7 Address = Address{
		City:     "DENPASAR",
		Province: "BALI",
		Country:  "INDONESIA",
	}
	var address8 *Address = &address7
	address8.City = "UBUD"
	address8 = &Address{
		City:     "GILIMANUK",
		Province: "BALI",
		Country:  "INDONESIA",
	}
	fmt.Println(address7)
	fmt.Println(address8)

	// Bagaimana jika kita ingin memaksa semuanya pindah. maka dilakukan dengan code pointer berikut ini
	var address9 Address = Address{
		City:     "SAMARINDA",
		Province: "KALIMANTAN TIMUR",
		Country:  "INDONESIA",
	}
	var address10 *Address = &address9
	address10.City = "BERAU"
	*address10 = Address{
		City:     "BALIKPAPAN",
		Province: "BALI",
		Country:  "INDONESIA",
	}
	fmt.Println(address9)
	fmt.Println(address10)

	// cara lain
	var address11 *Address = &Address{"JAKSEL", "JAKARTA", "INDONESIA"}
	var address12 *Address = new(Address)

	address12.Province = "JAKARTA"

	fmt.Println(address11)
	fmt.Println(address12)
}
