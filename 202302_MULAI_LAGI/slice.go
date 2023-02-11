package main

import "fmt"

// slice mirip dengan array.
// yang membedakan adalah ukuran yang bisa berubah

// slice dan array selalu terkoneksi. dimana slice adalah data yang mengakses sebagian atau seluruh data di array

// slice: pointer, length, capacity
// pointer adalah petunjuku data pertama di array pada slice
// length adalah panjang dari slice, dan
// capacity adalah kapisatas dari slice. dimana length tidak boleh lebih dari capcity

// cara membuat slice
// 1. array[low:high] index low. sampe dengan index sebelum high
// 2. array[low:] dari index low sampe dengan index akhir
// 3. array[:high] dari index 0 sampe dengan index sebelum high
// 4. array[:] slice dari semua data array

// saat kita ubah semua data di array, slicenya akan berubah
// begitupun ketika ubah slicenya, array akan berubah

func main() {
	var months = [...]string{
		"januari",
		"februari",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1)) // cek length
	fmt.Println(cap(slice1)) // cek capacity

	slice2 := months[10:]
	fmt.Println(slice2)

	slice3 := append(slice2, "DJALAL") // mereplace. kalo masih muat kapasitas, akan masih jadi satu reference. kalo sudah engga muat bikin array baru
	fmt.Println(slice3)
	slice3[1] = "Bukan Desember"
	fmt.Println(slice3)

	// bikin slice secara langsung
	newSlice := make([]string, 2, 5) // (arraynya, panjangnya, capacity)
	newSlice[0] = "Djalal"
	newSlice[1] = "Kurnia"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	// copy. pastikan lengthnya sama
	copySlice := make([]string, 2, 5)
	copy(copySlice, newSlice)
	fmt.Println(copySlice)
}
