package main

import (
    "fmt"
)

const NMAX = 100

type PolusiUdara struct {
    ID     int
    Tahun  int
    Bulan  int 
    Daerah string
    PM25   float64
    PM10   float64
    CO2    float64
    Status string
}

var dataPolusi [NMAX]PolusiUdara
var jumlahData int

func main() {
    initSampleData()

    for {
        fmt.Println("\n=== Selamat Datang di Aplikasi Pemantauan Polusi Udara Lokal Bandung ===")
        fmt.Println("1. Start")
        fmt.Println("2. Keluar")
        fmt.Print("Pilih menu: ")

        var awal int
        fmt.Scan(&awal)

        if awal == 1 {
            menuUtama()
            break
        } else if awal == 2 {
            fmt.Println("Terima kasih. Sampai jumpa!")
            return
        } else {
            fmt.Println("Pilihan tidak valid. Coba lagi.")
        }
    }
}



func initSampleData() {
    dataPolusi[0] = PolusiUdara{ID: 1, Tahun: 2024, Bulan: 1, Daerah: "Ciparay", PM25: 56.2, PM10: 85.1, CO2: 402.3, Status: "Sedang"}
    dataPolusi[1] = PolusiUdara{ID: 2, Tahun: 2024, Bulan: 2, Daerah: "Rancamanyar", PM25: 78.4, PM10: 122.7, CO2: 489.1, Status: "Buruk"}
    dataPolusi[2] = PolusiUdara{ID: 3, Tahun: 2024, Bulan: 3, Daerah: "Bojongsoang", PM25: 42.9, PM10: 69.5, CO2: 384.0, Status: "Baik"}
    jumlahData = 3
}

func tampilkanData() {
    fmt.Println("\n=== Daftar Data Polusi Udara Tahun 2024 ===")
    for i := 0; i < jumlahData; i++ {
        d := dataPolusi[i]
        fmt.Printf("%d. [%s %d] %s - PM2.5: %.1f, PM10: %.1f, CO2: %.1f, Status: %s\n",
            d.ID, namaBulan(d.Bulan), d.Tahun, d.Daerah, d.PM25, d.PM10, d.CO2, d.Status)
    }
}

func tambahData() {
    var daerah string
    var pm25, pm10, co2 float64
    var tahun, bulan int
    var status string

    fmt.Print("Tahun: ")
    fmt.Scan(&tahun)
    fmt.Print("Bulan (1-12): ")
    fmt.Scan(&bulan)
    fmt.Print("Nama Daerah: ")
    fmt.Scan(&daerah)
    fmt.Print("PM2.5: ")
    fmt.Scan(&pm25)
    fmt.Print("PM10: ")
    fmt.Scan(&pm10)
    fmt.Print("CO2: ")
    fmt.Scan(&co2)

   if pm25 <= 12 && pm10 <= 50 && co2 <= 450 {
    status = "Baik"
	} else if pm25 <= 35 && pm10 <= 150 && co2 <= 1000 {
		status = "Sedang"
	} else {
		status = "Berbahaya"
	}

    fmt.Println("Status:", status)

    dataPolusi[jumlahData] = PolusiUdara{
        ID:     jumlahData + 1,
        Tahun:  tahun,
        Bulan:  bulan,
        Daerah: daerah,
        PM25:   pm25,
        PM10:   pm10,
        CO2:    co2,
        Status: status,
    }
    jumlahData++
    fmt.Println("Data polusi berhasil ditambahkan!")
}

func editData() {
    tampilkanData()
    fmt.Print("Masukkan ID data yang ingin diedit: ")
    var id int
    fmt.Scan(&id)

    for i := 0; i < jumlahData; i++ {
        if dataPolusi[i].ID == id {
            fmt.Print("Nama Daerah Baru: ")
            fmt.Scan(&dataPolusi[i].Daerah)
            fmt.Print("PM2.5 Baru: ")
            fmt.Scan(&dataPolusi[i].PM25)
            fmt.Print("PM10 Baru: ")
            fmt.Scan(&dataPolusi[i].PM10)
            fmt.Print("CO2 Baru: ")
            fmt.Scan(&dataPolusi[i].CO2)

            // Hitung ulang status berdasarkan nilai PM dan CO2
            pm25 := dataPolusi[i].PM25
            pm10 := dataPolusi[i].PM10
            co2 := dataPolusi[i].CO2

            if pm25 < 13 && pm10 < 54 && co2 < 450 {
				dataPolusi[i].Status = "Baik"
			} else if pm25 < 36 && pm10 < 154 && co2 > 451 {
				dataPolusi[i].Status = "Sedang"
			} else {
				dataPolusi[i].Status = "Berbahaya"
			}

            fmt.Println("Data berhasil diperbarui!")
            return
        }
    }
    fmt.Println("ID tidak ditemukan.")
}


func hapusData() {
    tampilkanData()
    fmt.Print("Masukkan ID data yang ingin dihapus: ")
    var id int
    fmt.Scan(&id)

    for i := 0; i < jumlahData; i++ {
        if dataPolusi[i].ID == id {
            for j := i; j < jumlahData-1; j++ {
                dataPolusi[j] = dataPolusi[j+1]
            }
            jumlahData--
            fmt.Println("Data berhasil dihapus.")
            return
        }
    }
    fmt.Println("ID tidak ditemukan.")
}

func namaBulan(bulan int) string {
    bulanNama := [12]string{"Januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli", "Agustus", "September", "Oktober", "November", "Desember"}
    if bulan >= 1 && bulan <= 12 {
        return bulanNama[bulan-1]
    }
    return "Bulan Tidak Valid"
}

func cariBerdasarkanBulan() {
    fmt.Print("Masukkan bulan yang ingin dicari: ")
    var bulan int
    fmt.Scan(&bulan)

    fmt.Printf("\n=== Data Polusi Udara Bulan %s ===\n", namaBulan(bulan))
    found := false
    for i := 0; i < jumlahData; i++ {
        if dataPolusi[i].Bulan == bulan {
            d := dataPolusi[i]
            fmt.Printf("%d. [%s %d] %s - PM2.5: %.1f, PM10: %.1f, CO2: %.1f, Status: %s\n",
                d.ID, namaBulan(d.Bulan), d.Tahun, d.Daerah, d.PM25, d.PM10, d.CO2, d.Status)
            found = true
        }
    }
    if !found {
        fmt.Println("Tidak ada data untuk bulan tersebut.")
    }
}

func cariBerdasarkanTahun() {
    fmt.Print("Masukkan tahun yang ingin dicari: ")
    var tahun int
    fmt.Scan(&tahun)

    fmt.Printf("\n=== Data Polusi Udara Tahun %d ===\n", tahun)
    found := false
    for i := 0; i < jumlahData; i++ {
        if dataPolusi[i].Tahun == tahun {
            d := dataPolusi[i]
            fmt.Printf("%d. [%s %d] %s - PM2.5: %.1f, PM10: %.1f, CO2: %.1f, Status: %s\n",
                d.ID, namaBulan(d.Bulan), d.Tahun, d.Daerah, d.PM25, d.PM10, d.CO2, d.Status)
            found = true
        }
    }
    if !found {
        fmt.Println("Tidak ada data untuk tahun tersebut.")
    }
}

func sortDataBerdasarkanTahun() {
    for i := 0; i < jumlahData-1; i++ {
        for j := i + 1; j < jumlahData; j++ {
            if dataPolusi[i].Tahun > dataPolusi[j].Tahun {
                dataPolusi[i], dataPolusi[j] = dataPolusi[j], dataPolusi[i]
            }
        }
    }
    fmt.Println("Data berhasil diurutkan berdasarkan tahun (ascending).")
    tampilkanData()
}
