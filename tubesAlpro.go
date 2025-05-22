package main

import (
	"fmt"
	"time"
)

type layanan_berlangganan struct {
	no                int
	nama_layanan      string
	biaya             int
	metode_pembayaran string
	tgl_pembayaran    string
	status            string
}

const NMAX int = 10

type tabInt [NMAX]layanan_berlangganan

var jumlah int = 0

func ValidateDate(dateStr string, layout string) bool {
	_, err := time.Parse(layout, dateStr)
	return err == nil
}

func IsDateValid(dateStr string) bool {
	layout := "02-01-2006"
	t, err := time.Parse(layout, dateStr)
	if err != nil {
		return false
	}
	checkStr := t.Format(layout)
	return checkStr == dateStr
}

func tambahLayanan(A *tabInt, tabungan int, jumlah *int) {
	var idx int
	idx = *jumlah
	fmt.Printf("%-35s: ", "Nama Layanan")
	fmt.Scan(&A[idx].nama_layanan)
	fmt.Printf("%-35s: ", "Harga")
	fmt.Scan(&A[idx].biaya)
	_, utang := hitUtang(*A, tabungan, *jumlah)
	if utang+A[idx].biaya > tabungan {
		fmt.Println("Saldo anda tidak cukup")
	} else {
		fmt.Printf("%-35s: ", "Metode Pembayaran (cash/transfer)")
		fmt.Scan(&A[idx].metode_pembayaran)
		for A[idx].metode_pembayaran != "cash" && A[idx].metode_pembayaran != "transfer" {
			fmt.Println("Input tidak valid!")
			fmt.Println("Mohon lakukan pengisian ulang sesuai aturan pilihan yang tertera")
			fmt.Printf("%-35s: ", "Metode Pembayaran (cash/transfer)")
			fmt.Scan(&A[idx].metode_pembayaran)
		}
		fmt.Printf("%-35s: ", "Tanggal Pembayaran (dd-mm-yyyy)")
		fmt.Scan(&A[idx].tgl_pembayaran)
		for !IsDateValid(A[idx].tgl_pembayaran) {
			fmt.Println("Input tidak valid!")
			fmt.Printf("%-35s: ", "Tanggal Pembayaran (dd-mm-yyyy)")
			fmt.Scan(&A[idx].tgl_pembayaran)
		}
		fmt.Printf("%-35s: ", "Status (Lunas/Belum)")
		fmt.Scan(&A[idx].status)
		for A[idx].status != "Lunas" && A[idx].status != "Belum" {
			fmt.Println("Input tidak valid!")
			fmt.Println("Mohon lakukan pengisian ulang sesuai aturan pilihan yang tertera")
			fmt.Printf("%-35s: ", "Status (Lunas/Belum)")
			fmt.Scan(&A[idx].status)
		}
		A[idx].no = idx + 1
		*jumlah = *jumlah + 1
		fmt.Println("=======================================")
		fmt.Println("Selamat! Anda berhasil menambahkan data")
	}
}

func tampilkanArray(A tabInt, jumlah int) {
	fmt.Printf("%-4s %-20s %-12s %-10s %-20s %-10s\n", "No", "Nama Layanan", "Biaya", "Metode", "Tanggal Pembayaran", "Status")
	fmt.Println("------------------------------------------------------------------------------------")
	for i := 0; i < jumlah; i++ {
		fmt.Printf("%-4d %-20s Rp%-10d %-10s %-20s %-10s\n", A[i].no, A[i].nama_layanan, A[i].biaya, A[i].metode_pembayaran, A[i].tgl_pembayaran, A[i].status)
	}
	fmt.Println()
	fmt.Println("Jumlah Data : ", jumlah)
}

func editArray(A *tabInt, tabungan int, idx int, jumlah *int) {
	fmt.Println("======== EDIT ========")
	idx = idx - 1
	if idx < 0 || idx >= *jumlah {
		fmt.Println("Nomor layanan tidak valid!")
		return
	}
	
	fmt.Printf("%-35s: ", "Nama Layanan")
	fmt.Scan(&A[idx].nama_layanan)
	fmt.Print("Biaya : ")
	fmt.Scan(&A[idx].biaya)
	
	// Hitung utang tanpa item yang sedang diedit
	utangLain := 0
	for i := 0; i < *jumlah; i++ {
		if i != idx && A[i].status != "Lunas" && A[i].status != "lunas" {
			utangLain += A[i].biaya
		}
	}
	
	if utangLain+A[idx].biaya > tabungan {
		fmt.Println("Saldo anda tidak cukup")
		return
	}
	
	fmt.Printf("%-35s: ", "Metode Pembayaran (cash/transfer)")
	fmt.Scan(&A[idx].metode_pembayaran)
	for A[idx].metode_pembayaran != "cash" && A[idx].metode_pembayaran != "transfer" {
		fmt.Println("Input tidak valid!")
		fmt.Println("Mohon lakukan pengisian ulang sesuai aturan pilihan yang tertera")
		fmt.Printf("%-35s: ", "Metode Pembayaran (cash/transfer)")
		fmt.Scan(&A[idx].metode_pembayaran)
	}
	fmt.Printf("%-35s: ", "Tanggal Pembayaran (dd-mm-yyyy)")
	fmt.Scan(&A[idx].tgl_pembayaran)
	for !IsDateValid(A[idx].tgl_pembayaran) {
		fmt.Println("Input tidak valid!")
		fmt.Printf("%-35s: ", "Tanggal Pembayaran (dd-mm-yyyy)")
		fmt.Scan(&A[idx].tgl_pembayaran)
	}
	fmt.Printf("%-35s: ", "Status (Lunas/Belum)")
	fmt.Scan(&A[idx].status)
	for A[idx].status != "Lunas" && A[idx].status != "Belum" {
		fmt.Println("Input tidak valid!")
		fmt.Println("Mohon lakukan pengisian ulang sesuai aturan pilihan yang tertera")
		fmt.Printf("%-35s: ", "Status (Lunas/Belum)")
		fmt.Scan(&A[idx].status)
	}
	
	fmt.Println("=======================================")
	fmt.Println("Selamat! Anda berhasil mengedit data")
}

func cekIsiArray(A *tabInt) int {
	var jmlh int
	for i := 0; i < NMAX && A[i].no != 0; i++ {
		jmlh++
	}
	return jmlh
}

func sortArray(A *tabInt, jumlah int) {
	var i int = 0
	for i < jumlah-1 {
		var maxIdx int = i
		var j int = i + 1
		for j < jumlah {
			if A[j].biaya > A[maxIdx].biaya {
				maxIdx = j
			}
			j = j + 1
		}
		var temp layanan_berlangganan = A[i]
		A[i] = A[maxIdx]
		A[maxIdx] = temp
		i = i + 1
	}
	// Setelah sort, perbaiki nomor urut
	reorderNumbers(A, jumlah)
}

// Fungsi untuk menghitung selisih hari antara dua tanggal
func hitungSelisihHari(tanggal1, tanggal2 string) int {
	layout := "02-01-2006"
	t1, err1 := time.Parse(layout, tanggal1)
	t2, err2 := time.Parse(layout, tanggal2)
	
	if err1 != nil || err2 != nil {
		return 0
	}
	
	selisih := t2.Sub(t1)
	return int(selisih.Hours() / 24)
}

func cekJatuhTempo(A *tabInt, jumlah int) {
	var hariIni string
	hariIni = "20-05-2025"
	fmt.Println("⏰ Status jatuh tempo layanan: ")
	
	found := false
	for i := 0; i < jumlah; i++ {
		if A[i].status != "Lunas" && A[i].status != "lunas" {
			selisihHari := hitungSelisihHari(hariIni, A[i].tgl_pembayaran)
			
			if selisihHari < 0 {
				fmt.Printf("⚠️ Sudah lewat: %s (Tanggal: %s) - Terlambat %d hari\n", A[i].nama_layanan, A[i].tgl_pembayaran, -selisihHari)
				found = true
			} else if selisihHari <= 7 {
				fmt.Printf("⏰ Mendekati: %s (Tanggal: %s) - %d hari lagi\n", A[i].nama_layanan, A[i].tgl_pembayaran, selisihHari)
				found = true
			} else if selisihHari <= 30 {
				fmt.Printf("📅 Dalam waktu dekat: %s (Tanggal: %s) - %d hari lagi\n", A[i].nama_layanan, A[i].tgl_pembayaran, selisihHari)
				found = true
			} else {
				fmt.Printf("📆 Masih lama: %s (Tanggal: %s) - %d hari lagi\n", A[i].nama_layanan, A[i].tgl_pembayaran, selisihHari)
				found = true
			}
		}
	}
	
	if !found {
		fmt.Println("Tidak ada layanan yang belum lunas")
	}
}

func rekomendasiPengeluaran(A *tabInt, jumlah int) {
	fmt.Println("📉 Rekomendasi layanan yang dapat dihentikan berdasarkan harga termahal:")
	if jumlah == 0 {
		fmt.Println("Tidak ada layanan yang tersedia")
		return
	}
	var hargaTermahal int = A[0].biaya
	var i int = 1
	for i < jumlah {
		if A[i].biaya > hargaTermahal {
			hargaTermahal = A[i].biaya
		}
		i = i + 1
	}
	fmt.Printf("Layanan dengan biaya termahal (Rp%d):\n", hargaTermahal)
	i = 0
	for i < jumlah {
		if A[i].biaya == hargaTermahal {
			fmt.Printf("- %s (Rp%d)\n", A[i].nama_layanan, A[i].biaya)
		}
		i = i + 1
	}
}

// Fungsi untuk mengurutkan ulang nomor setelah hapus/edit
func reorderNumbers(A *tabInt, jumlah int) {
	for i := 0; i < jumlah; i++ {
		A[i].no = i + 1
	}
}

func hapusLayanan(A *tabInt, idx int, jumlah *int) {
	var index int = idx - 1
	if index < 0 || index >= *jumlah {
		fmt.Println("Nomor layanan tidak valid!")
		return
	}
	var i int = index
	for i < *jumlah-1 {
		A[i] = A[i+1]
		i = i + 1
	}
	A[*jumlah-1] = layanan_berlangganan{}
	*jumlah = *jumlah - 1
	
	// Perbaiki nomor urut setelah hapus
	reorderNumbers(A, *jumlah)
	
	fmt.Println("Layanan berhasil dihapus!")
}

func searchData(A *tabInt, keyword string, jumlah int) {
	var ketemu int
	fmt.Printf("%-4s %-20s %-12s %-10s %-20s %-10s\n", "No", "Nama Layanan", "Biaya", "Metode", "Tanggal Pembayaran", "Status")
	fmt.Println("------------------------------------------------------------------------------------")
	for i := 0; i < jumlah; i++ {
		if A[i].status == keyword || A[i].nama_layanan == keyword || A[i].tgl_pembayaran == keyword {
			fmt.Printf("%-4d %-20s Rp%-10d %-10s %-20s %-10s\n", A[i].no, A[i].nama_layanan, A[i].biaya, A[i].metode_pembayaran, A[i].tgl_pembayaran, A[i].status)
			ketemu++
		}
	}
	if ketemu == 0 {
		fmt.Println("Data tidak ditemukan")
	} else {
		fmt.Printf("\nDitemukan %d data\n", ketemu)
	}
}

func menu(A *tabInt) {
	var pil, isiArr int
	var tabungan int
	tabungan = 50000
	pil = 0
	loadData(A)
	isiArr = cekIsiArray(A)
	
	for pil != 9 {
		fmt.Println()
		fmt.Println("===================== MENU ========================")
		fmt.Println("1. Tampilkan Daftar Layanan")
		fmt.Println("2. Tambahkan Daftar Layanan")
		fmt.Println("3. Edit Daftar Layanan")
		fmt.Println("4. Sort Daftar Layanan")
		fmt.Println("5. Cek Jatuh Tempo")
		fmt.Println("6. Rekomendasi Pengeluaran")
		fmt.Println("7. Hapus Layanan")
		fmt.Println("8. Search Daftar Layanan")
		fmt.Println("9. Keluar")
		fmt.Println("====================================================")
		fmt.Print("Pilih : ")
		fmt.Scan(&pil)
		fmt.Println()
		
		switch pil {
		case 1:
			fmt.Println("============ DATA PAGE ===============")
			tampilkanArray(*A, isiArr)
		case 2:
			fmt.Println("============ ADD PAGE ===============")
			var move string
			var index int
			if isiArr >= NMAX {
				fmt.Println("Array sudah penuh!")
				fmt.Println("Hapus salah satu data jika Anda tetap ingin menambahkan data")
				fmt.Print("Mau hapus data? (Ya/Tidak) ")
				fmt.Scan(&move)
				if move == "Ya" {
					fmt.Print("Baris nomor berapa yang ingin dihapus? ")
					fmt.Scan(&index)
					hapusLayanan(A, index, &isiArr)
				}
			}
			ok, utang := hitUtang(*A, tabungan, isiArr)
			if ok {
				tambahLayanan(A, tabungan, &isiArr)
			} else {
				fmt.Println("Saldo anda tidak cukup!")
				fmt.Println("Uang anda saat ini: ", tabungan)
				fmt.Println("Layanan yang belum dibayar: ", utang)
			}
			tampilkanArray(*A, isiArr)
		case 3:
			var index int
			tampilkanArray(*A, isiArr)
			fmt.Println("============ EDIT PAGE ===============")
			fmt.Print("Baris nomor berapa yang ingin diubah? ")
			fmt.Scan(&index)
			editArray(A, tabungan, index, &isiArr)
			tampilkanArray(*A, isiArr)
		case 4:
			fmt.Println("============ SORT PAGE ===============")
			sortArray(A, isiArr)
			tampilkanArray(*A, isiArr)
		case 5:
			fmt.Println("============ CEK JATUH TEMPO PAGE ===============")
			cekJatuhTempo(A, isiArr)
		case 6:
			fmt.Println("============ REKOMENDASI PAGE ===============")
			rekomendasiPengeluaran(A, isiArr)
		case 7:
			var index int
			tampilkanArray(*A, isiArr)
			fmt.Println("============ DELETE PAGE ===============")
			fmt.Print("Baris nomor berapa yang ingin dihapus? ")
			fmt.Scan(&index)
			hapusLayanan(A, index, &isiArr)
			fmt.Println("Setelah hapus:")
			tampilkanArray(*A, isiArr)
		case 8:
			var keyword string
			fmt.Println("============ SEARCH PAGE ===============")
			fmt.Print("Cari data berdasarkan nama layanan / tanggal pembayaran / status : ")
			fmt.Scan(&keyword)
			searchData(A, keyword, isiArr)
		case 9:
			fmt.Println("Terima kasih telah menggunakan program ini!")
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

func loadData(data *tabInt) {
	*data = tabInt{
		{1, "Netflix", 10000, "transfer", "25-05-2025", "Belum"},
		{2, "Spotify", 8000, "cash", "19-05-2025", "Belum"},
		{3, "YouTube Premium", 30000, "transfer", "30-05-2025", "Belum"},
		{4, "Disney+", 1100, "cash", "15-06-2025", "Lunas"},
		{5, "Amazon Prime", 5000, "transfer", "20-06-2025", "Lunas"},
		{6, "HBO Max", 5000, "cash", "25-06-2025", "Lunas"},
	}
}

func hitUtang(A tabInt, tabungan, jumlah int) (bool, int) {
	var utang int
	for i := 0; i < jumlah; i++ {
		if A[i].status != "Lunas" && A[i].status != "lunas" {
			utang += A[i].biaya
		}
	}
	return utang < tabungan, utang
}

func hitPengeluaran(A tabInt, jumlah int) int {
	var pengeluaran int
	for i := 0; i < jumlah; i++ {
		pengeluaran += A[i].biaya
	}
	return pengeluaran
}

func main() {
	var data tabInt
	menu(&data)
	//sss
}