package helpers

import "fmt"

func printBiodata(p Participant, id int) {
	fmt.Printf("ID: %+v \n", id)
	fmt.Printf("Nama: %+v \n", p.name)
	fmt.Printf("Alamat: %+v \n", p.address)
	fmt.Printf("Pekerjaan: %+v \n", p.job)
	fmt.Printf("Alasan: %+v \n", p.reason)
}

func notFound() {
	fmt.Printf("ID tidak ditemukan")
}


