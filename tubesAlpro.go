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

const NMAX int = 20

type tabLayanan [NMAX]layanan_berlangganan

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

// TAMPILKAN ISI ARRAY
func tampilkanArray(A tabLayanan, jumlah int) {
	fmt.Println("---------------------------------------------------------------------------------------------")
	fmt.Printf("|%-4s | %-20s | %-12s | %-10s | %-20s | %-10s | \n", "No", "Nama Layanan", "Biaya", "Metode", "Tanggal Pembayaran", "Status")
	fmt.Println("---------------------------------------------------------------------------------------------")
	for i := 0; i < jumlah; i++ {
		fmt.Printf("|%-4d | %-20s | Rp%-10d | %-10s | %-20s | %-10s | \n", A[i].no, A[i].nama_layanan, A[i].biaya, A[i].metode_pembayaran, A[i].tgl_pembayaran, A[i].status)
	}
	fmt.Println("---------------------------------------------------------------------------------------------")
	fmt.Println()
	fmt.Println("Jumlah Data : ", jumlah)
}

func tambahLayanan(A *tabLayanan, tabungan int, jumlah *int) {
	var idx int
	idx = *jumlah
	fmt.Printf("%-35s: ", "Nama Layanan")
	fmt.Scan(&A[idx].nama_layanan)
	fmt.Printf("%-35s: ", "Harga")
	fmt.Scan(&A[idx].biaya)
	_, utang := hitUtang(*A, tabungan, *jumlah)
	if utang+A[idx].biaya > tabungan {
		fmt.Println("Saldo anda tidak cukup")
		fmt.Println()
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

func editStrip(A *tabLayanan, tabungan int, idx int, jumlah *int) bool {
	if idx < 0 || idx >= *jumlah {
		fmt.Println("Nomor layanan tidak valid!")
		return false
	}
	dataAsli := A[idx]
	var namaLayanan, metodePembayaran, tglPembayaran, status string
	var biaya int
	var biayaStr string
	var konfirmasi string
	fmt.Printf("Data no-%d saat ini: \n", idx+1)
	fmt.Println("---------------------------------------------------------------------------------------------")
	fmt.Printf("|%-4s | %-20s | %-12s | %-10s | %-20s | %-10s | \n", "No", "Nama Layanan", "Biaya", "Metode", "Tanggal Pembayaran", "Status")
	fmt.Println("---------------------------------------------------------------------------------------------")
	fmt.Printf("|%-4d | %-20s | Rp%-10d | %-10s | %-20s | %-10s | \n", A[idx].no, A[idx].nama_layanan, A[idx].biaya, A[idx].metode_pembayaran, A[idx].tgl_pembayaran, A[idx].status)
	fmt.Println("---------------------------------------------------------------------------------------------")
	fmt.Printf("%-35s: ", "Nama Layanan")
	fmt.Scan(&namaLayanan)
	if namaLayanan == "back" {
		fmt.Println("Kembali ke menu...")
		return true
	}
	if namaLayanan == "-" {
		namaLayanan = dataAsli.nama_layanan
	}
	fmt.Printf("%-35s: ", "Biaya")
	fmt.Scan(&biayaStr)
	if biayaStr == "back" {
		fmt.Println("Kembali ke menu...")
		return true
	}
	if biayaStr == "-" {
		biaya = dataAsli.biaya
	}
	fmt.Scanf(biayaStr, "%d", &biaya)
	fmt.Printf("%-35s: ", "Metode Pembayaran (cash/transfer)")
	fmt.Scan(&metodePembayaran)
	if metodePembayaran == "back" {
		fmt.Println("Kembali ke menu...")
		return true
	}
	if metodePembayaran == "-" {
		metodePembayaran = dataAsli.metode_pembayaran
	}
	for metodePembayaran != "cash" && metodePembayaran != "transfer" {
		fmt.Println("Input tidak valid!")
		fmt.Println("Mohon lakukan pengisian ulang sesuai aturan pilihan yang tertera")
		fmt.Printf("%-35s: ", "Metode Pembayaran (cash/transfer)")
		fmt.Scan(&metodePembayaran)
		if metodePembayaran == "back" {
			fmt.Println("Kembali ke menu...")
			return true
		}
		if metodePembayaran == "-" {
			metodePembayaran = dataAsli.metode_pembayaran
		}
	}
	fmt.Printf("%-35s: ", "Tanggal Pembayaran (dd-mm-yyyy)")
	fmt.Scan(&tglPembayaran)
	if tglPembayaran == "back" {
		fmt.Println("Kembali ke menu...")
		return true
	}
	if tglPembayaran == "-" {
		tglPembayaran = dataAsli.tgl_pembayaran
	}
	for !IsDateValid(tglPembayaran) {
		fmt.Println("Input tidak valid!")
		fmt.Printf("%-35s: ", "Tanggal Pembayaran (dd-mm-yyyy)")
		fmt.Scan(&tglPembayaran)
		if tglPembayaran == "back" {
			fmt.Println("Kembali ke menu...")
			return true
		}
		if tglPembayaran == "-" {
			tglPembayaran = dataAsli.tgl_pembayaran
		}
	}
	fmt.Printf("%-35s: ", "Status (Lunas/Belum)")
	fmt.Scan(&status)
	if status == "back" {
		fmt.Println("Kembali ke menu...")
		return true
	}
	if status == "-" {
		status = dataAsli.status
	}
	for status != "Lunas" && status != "Belum" {
		fmt.Println("Input tidak valid!")
		fmt.Println("Mohon lakukan pengisian ulang sesuai aturan pilihan yang tertera")
		fmt.Printf("%-35s: ", "Status (Lunas/Belum)")
		fmt.Scan(&status)
		if status == "back" {
			fmt.Println("Kembali ke menu...")
			return true
		}
		if status == "-" {
			status = dataAsli.status
		}
	}
	totalUtangLain := 0
	for i := 0; i < *jumlah; i++ {
		if i != idx && (A[i].status == "Belum" || A[i].status == "belum") {
			totalUtangLain += A[i].biaya
		}
	}
	totalUtangBaru := totalUtangLain
	if status == "Belum" || status == "belum" {
		totalUtangBaru += biaya
	}

	if totalUtangBaru > tabungan {
		fmt.Println("Saldo anda tidak cukup dengan perubahan ini!")
		fmt.Printf("Total utang akan menjadi: Rp%d, sedangkan saldo: Rp%d\n", totalUtangBaru, tabungan)
		fmt.Println("Edit dibatalkan!")
		return false
	}
	fmt.Println("\n========== KONFIRMASI PERUBAHAN ==========")
	fmt.Println("Data Lama: ")
	fmt.Println("---------------------------------------------------------------------------------------------")
	fmt.Printf("|%-4s | %-20s | %-12s | %-10s | %-20s | %-10s | \n", "No", "Nama Layanan", "Biaya", "Metode", "Tanggal Pembayaran", "Status")
	fmt.Println("---------------------------------------------------------------------------------------------")
	fmt.Printf("|%-4d | %-20s | Rp%-10d | %-10s | %-20s | %-10s | \n", A[idx].no, dataAsli.nama_layanan, dataAsli.biaya, dataAsli.metode_pembayaran, dataAsli.tgl_pembayaran, dataAsli.status)
	fmt.Println("---------------------------------------------------------------------------------------------")
	// fmt.Printf("%-20s %-12s %-10s %-20s %-10s\n", "Nama Layanan", "Biaya", "Metode", "Tanggal Pembayaran", "Status")
	// fmt.Println("------------------------------------------------------------------------------------")
	// fmt.Printf("%-20s Rp%-10d %-10s %-20s %-10s\n",
	// 	dataAsli.nama_layanan, dataAsli.biaya, dataAsli.metode_pembayaran,
	// 	dataAsli.tgl_pembayaran, dataAsli.status)
	// fmt.Printf("Data Baru: %s - Rp%d - %s - %s - %s\n",
	// 	namaLayanan, biaya, metodePembayaran, tglPembayaran, status)
	fmt.Println("Data Baru: ")
	fmt.Println("---------------------------------------------------------------------------------------------")
	fmt.Printf("|%-4s | %-20s | %-12s | %-10s | %-20s | %-10s | \n", "No", "Nama Layanan", "Biaya", "Metode", "Tanggal Pembayaran", "Status")
	fmt.Println("---------------------------------------------------------------------------------------------")
	fmt.Printf("|%-4d | %-20s | Rp%-10d | %-10s | %-20s | %-10s | \n", A[idx].no, namaLayanan, biaya, metodePembayaran, tglPembayaran, status)
	fmt.Println("---------------------------------------------------------------------------------------------")
	fmt.Print("Apakah Anda yakin ingin menyimpan perubahan? (ya/tidak): ")
	fmt.Scan(&konfirmasi)

	if konfirmasi == "ya" || konfirmasi == "Ya" || konfirmasi == "YES" || konfirmasi == "yes" {
		A[idx].nama_layanan = namaLayanan
		A[idx].biaya = biaya
		A[idx].metode_pembayaran = metodePembayaran
		A[idx].tgl_pembayaran = tglPembayaran
		A[idx].status = status

		fmt.Println("=======================================")
		fmt.Println("Selamat! Anda berhasil mengedit data")
	} else {
		fmt.Println("Edit dibatalkan!")
	}
	return false
}

func cekIsiArray(A *tabLayanan) int {
	var jmlh int
	for i := 0; i < NMAX && A[i].no != 0; i++ {
		jmlh++
	}
	return jmlh
}

func sortArray(A *tabLayanan, jumlah int) {
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
	reorderNumbers(A, jumlah)
}

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

func cekJatuhTempo(A *tabLayanan, jumlah int) {
	var hariIni string
	fmt.Print("Hari ini : ")
	fmt.Scan(&hariIni)
	fmt.Println("⏰ Status jatuh tempo layanan: ")

	found := false
	for i := 0; i < jumlah; i++ {
		if A[i].status != "Lunas" && A[i].status != "lunas" {
			selisihHari := hitungSelisihHari(hariIni, A[i].tgl_pembayaran)

			if selisihHari < 0 {
				fmt.Printf("🔴 Sudah lewat: %s (Tanggal: %s) - Terlambat %d hari\n", A[i].nama_layanan, A[i].tgl_pembayaran, -selisihHari)
				found = true
			} else if selisihHari <= 7 {
				fmt.Printf("🟡 Mendekati: %s (Tanggal: %s) - %d hari lagi\n", A[i].nama_layanan, A[i].tgl_pembayaran, selisihHari)
				found = true
			} else {
				fmt.Printf("🟢 Masih lama: %s (Tanggal: %s) - %d hari lagi\n", A[i].nama_layanan, A[i].tgl_pembayaran, selisihHari)
				found = true
			}
		}
	}

	if !found {
		fmt.Println("Tidak ada layanan yang belum lunas")
	}
}

func rekomendasiPengeluaran(A *tabLayanan, jumlah int) {
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

func reorderNumbers(A *tabLayanan, jumlah int) {
	for i := 0; i < jumlah; i++ {
		A[i].no = i + 1
	}
}

func toLower(input string) string {
	result := ""
	for _, ch := range input {
		if ch >= 'A' && ch <= 'Z' {
			result += string(ch + ('a' - 'A'))
		} else {
			result += string(ch)
		}
	}
	return result
}


func hapusLayanan(A *tabLayanan, idx int, jumlah *int) {
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

func searchData(A *tabLayanan, keyword string, jumlah int) {
	var ketemu int
	fmt.Println("---------------------------------------------------------------------------------------------")
	fmt.Printf("|%-4s | %-20s | %-12s | %-10s | %-20s | %-10s | \n", "No", "Nama Layanan", "Biaya", "Metode", "Tanggal Pembayaran", "Status")
	fmt.Println("---------------------------------------------------------------------------------------------")
	for i := 0; i < jumlah; i++ {
		if A[i].status == keyword || A[i].nama_layanan == keyword || A[i].tgl_pembayaran == keyword {
			fmt.Printf("|%-4d | %-20s | Rp%-10d | %-10s | %-20s | %-10s | \n", A[i].no, A[i].nama_layanan, A[i].biaya, A[i].metode_pembayaran, A[i].tgl_pembayaran, A[i].status)
			fmt.Println("---------------------------------------------------------------------------------------------")
			ketemu++
		}
	}
	if ketemu == 0 {
		fmt.Println("Data tidak ditemukan")
	} else {
		fmt.Printf("\nDitemukan %d data\n", ketemu)
	}
}

func menu(A *tabLayanan) {
	var pil, isiArr int
	var tabungan int
	fmt.Print("masukan nominal tabungan: ")
	fmt.Scan(&tabungan)
	pil = 0
	dummyData(A)
	isiArr = cekIsiArray(A)

	for pil != 9 {
		fmt.Println()
		fmt.Printf("=============================================================================================\n")
		fmt.Printf("					DATA MENU\n")
		fmt.Printf("=============================================================================================\n")
		fmt.Println("1. Tampilkan Daftar Layanan")
		fmt.Println("2. Tambahkan Daftar Layanan")
		fmt.Println("3. Edit Daftar Layanan")
		fmt.Println("4. Sort Daftar Layanan")
		fmt.Println("5. Cek Jatuh Tempo")
		fmt.Println("6. Rekomendasi Pengeluaran")
		fmt.Println("7. Hapus Layanan")
		fmt.Println("8. Search Daftar Layanan")
		fmt.Println("9. Keluar")
		// fmt.Printf("=============================================================================================\n")
		fmt.Print("\nPilih : ")
		fmt.Scan(&pil)
		fmt.Println()

		switch pil {
		case 1:
			fmt.Printf("=============================================================================================\n")
			fmt.Printf("					DATA PAGE\n")
			fmt.Printf("=============================================================================================\n")
			tampilkanArray(*A, isiArr)
		case 2:
			fmt.Printf("=============================================================================================\n")
			fmt.Printf("					ADD PAGE\n")
			fmt.Printf("=============================================================================================\n")
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
			fmt.Printf("=============================================================================================\n")
			fmt.Printf("					EDIT PAGE\n")
			fmt.Printf("               (Gunakan tanda '-' jika tidak ada data yang ingin diisi)\n")
			fmt.Printf("=============================================================================================\n")
			fmt.Print("Baris nomor berapa yang ingin diubah? ")
			fmt.Scan(&index)
			backToMenu := editStrip(A, tabungan, index-1, &isiArr)
			if backToMenu {
			} else {
				tampilkanArray(*A, isiArr)
			}
		case 4:
			fmt.Printf("=============================================================================================\n")
			fmt.Printf("					SORT PAGE\n")
			fmt.Printf("=============================================================================================\n")
			sortArray(A, isiArr)
			tampilkanArray(*A, isiArr)
		case 5:
			fmt.Printf("=============================================================================================\n")
			fmt.Printf("					CEK JATUH TEMPO PAGE\n")
			fmt.Printf("=============================================================================================\n")
			cekJatuhTempo(A, isiArr)
		case 6:
			fmt.Printf("=============================================================================================\n")
			fmt.Printf("					REKOMENDASI PAGE\n")
			fmt.Printf("=============================================================================================\n")
			rekomendasiPengeluaran(A, isiArr)
		case 7:
			var index int
			tampilkanArray(*A, isiArr)
			fmt.Printf("=============================================================================================\n")
			fmt.Printf("					DELETE PAGE\n")
			fmt.Printf("=============================================================================================\n")
			fmt.Print("Baris nomor berapa yang ingin dihapus? ")
			fmt.Scan(&index)
			hapusLayanan(A, index, &isiArr)
			fmt.Println("Setelah hapus:")
			tampilkanArray(*A, isiArr)
		case 8:
			var keyword string
			fmt.Printf("=============================================================================================\n")
			fmt.Printf("					SEARCH PAGE\n")
			fmt.Printf("=============================================================================================\n")
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

func dummyData(A *tabLayanan) {
	A[0] = layanan_berlangganan{1, "Netflix", 54000, "transfer", "25-05-2025", "Belum"}
	A[1] = layanan_berlangganan{2, "Spotify", 54000, "cash", "19-05-2025", "Belum"}
	A[2] = layanan_berlangganan{3, "YouTube", 59000, "transfer", "30-05-2025", "Belum"}
	A[3] = layanan_berlangganan{4, "Disney+", 39000, "cash", "15-06-2025", "Lunas"}
	A[4] = layanan_berlangganan{5, "ChatGPT", 5900, "transfer", "20-06-2025", "Lunas"}
	A[5] = layanan_berlangganan{6, "HBO", 7900, "cash", "25-06-2025", "Lunas"}
	A[6] = layanan_berlangganan{7, "Canva", 15000, "transfer", "10-07-2025", "Belum"}
	A[7] = layanan_berlangganan{8, "Adobe", 28700, "transfer", "05-07-2025", "Belum"}
	A[8] = layanan_berlangganan{9, "Microsoft", 11900, "cash", "12-07-2025", "Lunas"}
	A[9] = layanan_berlangganan{10, "Zoom", 14900, "transfer", "18-07-2025", "Belum"}
}

func hitUtang(A tabLayanan, tabungan, jumlah int) (bool, int) {
	var utang int
	for i := 0; i < jumlah; i++ {
		if A[i].status != "Lunas" && A[i].status != "lunas" {
			utang += A[i].biaya
		}
	}
	return utang < tabungan, utang
}

func hitPengeluaran(A tabLayanan, jumlah int) int {
	var pengeluaran int
	for i := 0; i < jumlah; i++ {
		pengeluaran += A[i].biaya
	}
	return pengeluaran
}

func main() {
	var data tabLayanan
	menu(&data)
}
// fitur back minus
// lower up case to