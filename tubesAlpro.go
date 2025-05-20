package main
import "fmt"
type layanan_berlangganan struct {
	no                int
	nama_layanan      string
	biaya             int
	metode_pembayaran string
	tgl_pembayaran    string
	status            string
}
const NMAX int = 7
type tabInt [NMAX]layanan_berlangganan

func tambahLayanan(A *tabInt, no *int, tabungan, jumlah int) {
	A[*no].no = *no + 1
	fmt.Printf("%-35s: ", "Nama Layanan")
	fmt.Scan(&A[*no].nama_layanan)
	fmt.Printf("%-35s: ", "Harga")
	fmt.Scan(&A[*no].biaya)
	_, utang := hitUtang(*A, tabungan, jumlah)
	if utang+A[*no].biaya > tabungan {
		fmt.Println("Saldo anda tidak cukup")
	} else {
		fmt.Printf("%-35s: ", "Metode Pembayaran (cash/transfer)")
		fmt.Scan(&A[*no].metode_pembayaran)
		for A[*no].metode_pembayaran != "cash" && A[*no].metode_pembayaran != "transfer" {
			fmt.Println("Input tidak valid!")
			fmt.Println("Mohon lakukan pengisian ulang sesuai aturan pilihan yang tertera")
			fmt.Printf("%-35s: ", "Metode Pembayaran (cash/transfer)")
			fmt.Scan(&A[*no].metode_pembayaran)
		}
		fmt.Printf("%-35s: ", "Tanggal Pembayaran (tgl bln thn)")
		fmt.Scan(&A[*no].tgl_pembayaran)
		fmt.Printf("%-35s: ", "Status (Lunas/Belum)")
		fmt.Scan(&A[*no].status)
		for A[*no].status != "Lunas" && A[*no].status != "Belum" {
			fmt.Println("Input tidak valid!")
			fmt.Println("Mohon lakukan pengisian ulang sesuai aturan pilihan yang tertera")
			fmt.Printf("%-35s: ", "Status (Lunas/Belum)")
			fmt.Scan(&A[*no].status)
		}
		fmt.Println("=======================================")
		*no++
		fmt.Println("Selamat! Anda berhasil menambahkan data")
	}
}

func tampilkanArray(A *tabInt, no int) {
	fmt.Printf("%-4s %-20s %-12s %-10s %-20s %-10s\n", "No", "Nama Layanan", "Biaya", "Metode", "Tanggal Pembayaran", "Status")
	fmt.Println("------------------------------------------------------------------------------------")
	for i := 0; i < no; i++ {
		fmt.Printf("%-4d %-20s Rp%-10d %-10s %-20s %-10s\n", A[i].no, A[i].nama_layanan, A[i].biaya, A[i].metode_pembayaran, A[i].tgl_pembayaran, A[i].status)
	}
	fmt.Println()
}

func editArray(A *tabInt, idx int) {
	fmt.Println("======== EDIT ========")
	idx = idx - 1
	fmt.Print("Nama Layanan: ")
	fmt.Scan(&A[idx].nama_layanan)
	fmt.Print("biaya : ")
	fmt.Scan(&A[idx].biaya)
	fmt.Print("Metode Pembayaran : ")
	fmt.Scan(&A[idx].metode_pembayaran)
	fmt.Print("Tanggal Pembayaran : ")
	fmt.Scan(&A[idx].tgl_pembayaran)
	fmt.Print("Status : ")
	fmt.Scan(&A[idx].status)
	fmt.Println()
}

func cekIsiArray(A *tabInt) int {
	var jmlh int
	for i := 0; i < NMAX && A[i].no != 0; i++ {
		jmlh++
	}
	return jmlh
}

func sortArray(A *tabInt, jumlah int) {
	// DESCENDING (Besar ke Kecil)
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
}

func cekJatuhTempo(A *tabInt, jumlah int) {
	var hariIni string
	// fmt.Print("Masukkan tanggal hari ini (YYYY-MM-DD): ")
	// fmt.Scan(&hariIni)
	hariIni = "20-05-2025"

	fmt.Println("⏰ Pengeluaran mendekati jatuh tempo: ")
	for i := 0; i < jumlah; i++ {
		if (A[i].status) != "lunas" {
			if A[i].tgl_pembayaran <= hariIni {
				fmt.Println("⚠️ Sudah lewat:", A[i].nama_layanan)
			} else {
				fmt.Println("⏰ Mendekati:", A[i].nama_layanan)
			}
		}
	}
	
}

func rekomendasiPengeluaran(A *tabInt, jumlah int) {
	fmt.Println("📉 Rekomendasi pengeluaran yang bisa dikurangi:")
	var i int = 0
	for i < jumlah {
		if A[i].biaya > 100000 {
			fmt.Println("- ", A[i].nama_layanan, "(", A[i].biaya, ")")
		}
		i = i + 1
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
	i = 0
	for i < *jumlah {
		A[i].no = i + 1
		i = i + 1
	}
	fmt.Println("Layanan berhasil dihapus!")
}

func searchData(A *tabInt, keyword string, jumlah int) {
	var ketemu int
	for i := 0; i < jumlah; i++ {
		if A[i].status == keyword || A[i].nama_layanan == keyword || A[i].tgl_pembayaran == keyword {
			fmt.Println(A[i])
			ketemu++
		}
	}
	if ketemu == 0 {
		fmt.Println("Data tidak ditemukan")
	}
}

func menu(A tabInt) {
	var pil, isiArr int
	var tabungan int
	tabungan = 50000
	pil = 0
	for pil != 10 {
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
		loadData(&A)
		isiArr = cekIsiArray(&A)
		fmt.Println()
		fmt.Println()
		switch pil {
		case 1:
			fmt.Println("============ DATA PAGE ===============")
			tampilkanArray(&A, isiArr)
		case 2:
			fmt.Println("============ ADD PAGE ===============")
			ok, _ := hitUtang(A, tabungan, isiArr)
			_, utang := hitUtang(A, tabungan, isiArr)
			if ok {
				tambahLayanan(&A, &isiArr, tabungan, isiArr)
			} else {
				fmt.Println("Saldo anda tidak cukup!")
				fmt.Println("Uang anda saat ini: ", tabungan)
				fmt.Println("Layanan yang belum dibayar: ", utang)
			}
		case 3:
			var index int
			fmt.Println("============ EDIT PAGE ===============")
			fmt.Print("Baris nomor berapa yang ingin diubah? ")
			fmt.Scan(&index)
			fmt.Println("Sebelum edit ")
			fmt.Println(A)
			editArray(&A, index)
			fmt.Println("Sesudah edit ")
			fmt.Println(A)
		case 4:
			fmt.Println("============ SORT PAGE ===============")
			sortArray(&A, isiArr)
			tampilkanArray(&A, isiArr)
		case 5:
			fmt.Println("============ CEK JATUH TEMPO PAGE ===============")
			cekJatuhTempo(&A, isiArr)
		case 6:
			fmt.Println("============ REKOMENDASI PAGE ===============")
			rekomendasiPengeluaran(&A, isiArr)
		case 7:
			var index int
			fmt.Println("============ DELETE PAGE ===============")
			tampilkanArray(&A, isiArr)
			fmt.Print("Baris nomor berapa yang ingin dihapus? ")
			fmt.Scan(&index)
			hapusLayanan(&A, index, &isiArr)
			fmt.Println("Setelah hapus:")
			tampilkanArray(&A, isiArr)
			fmt.Println("data skrg = ", isiArr)
		case 8:
			var keyword string
			fmt.Println("============ SEARCH PAGE ===============")
			fmt.Print("Cari data berdasarkan nama layanan / tanggal pembayaran / status : ")
			fmt.Scan(&keyword)
			searchData(&A, keyword, isiArr)
		}
	}
}

func loadData(data *tabInt) {
	*data = tabInt{
		{1, "aaa", 10000, "ww", "20-302-0", "lunas"},
		{2, "bbb", 8000, "ww", "19-05-2025", "belum"},
		{3, "ccc", 30000, "ww", "20-302-0", "belum"},
		{4, "ddd", 1100, "ww", "20-302-1", "lunas"},
		{5, "eee", 5000, "ww", "20-302-0", "lunas"},
	}
}

func hitUtang(A tabInt, tabungan, jumlah int) (bool, int) {
	var utang int
	for i := 0; i < jumlah; i++ {
		if A[i].status != "lunas" {
			utang += A[i].biaya
		}
	}
	total := utang + 10000
	return total < tabungan, total
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
	menu(data)
}
