package main

import(
	"fmt"
	"strconv"
	"os"
)

type Student struct{
	ID, Name, Alamat, Pekerjaan, Alasan string
}

func main(){
	data := StudentData()
	if len(os.Args) == 2{
 	getData(os.Args[1],data)
	} else{
		fmt.Println("Salah Memasukkan Argument")
	}
}

func StudentData()[]Student{
	listData := [][]string{			//Memasukkan data
		{"1","MUHAMMAD ZUNAN ALFIKRI", "INDONESIA","MAHASISWA" ,"MEMILIKI CITA - CITA SEBAGAI BACKEND"},
		{"2","ARFAN JADULHAQ", "INDONESIA","MAHASISWA", "MEMILIKI CITA - CITA SEBAGAI BACKEND"},
		{"3","TRIYONO", "INDONESIA", "MAHASISWA", "MEMILIKI CITA - CITA SEBAGAI BACKEND"},
		{"4","ADITYA RIZKI PRATAMA", "INDONESIA", "MAHASISWA", "MEMILIKI CITA - CITA SEBAGAI BACKEND"},
		{"5","YULYANO THOMAS DJAYA", "INDONESIA", "MAHASISWA", "MEMILIKI CITA - CITA SEBAGAI BACKEND"},
		{"6","ARIFINAL", "INDONESIA", "MAHASISWA", "MEMILIKI CITA - CITA SEBAGAI BACKEND"},
		{"7","FELIX YANGSEN", "INDONESIA", "MAHASISWA", "MEMILIKI CITA - CITA SEBAGAI BACKEND"},
		{"8","WAHYU DWI RAMADHAN", "INDONESIA", "MAHASISWA", "MEMILIKI CITA - CITA SEBAGAI BACKEND"},
		{"9","MUHAMMAD HANIF NAUFAL EKA", "INDONESIA", "MAHASISWA", "MEMILIKI CITA - CITA SEBAGAI BACKEND"},
		{"10","THOBIB KHOIRUL ANNAS", "INDONESIA", "MAHASISWA", "MEMILIKI CITA - CITA SEBAGAI BACKEND"},
	}
	var singleData Student
	var result []Student
	for _,value := range listData{	//Memasukkan daftar data ke []Student
		singleData.ID = value[0]
		singleData.Name = value[1]
		singleData.Alamat = value[2]
		singleData.Pekerjaan = value[3]
		singleData.Alasan = value[4]
		result = append(result,singleData)
	}
	return result
}

func getData(n string,data []Student){
	m,err := strconv.Atoi(n)
	m--
	if err == nil && m >=0 && m < len(data){
		fmt.Printf(
			"ID \t\t\t: %v\nName \t\t\t: %v\nAlamat \t\t\t: %v\nPekerjaan \t\t: %v\nAlasan Belajar Go \t: %v\n",data[m].ID, data[m].Name, data[m].Alamat, data[m].Pekerjaan, data[m].Alasan)
	} else {
		fmt.Printf("Data yang anda masukkan TIDAK ADA")
	}

}