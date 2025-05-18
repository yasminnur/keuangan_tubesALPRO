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
func tambahLayanan(A *tabInt, no *int) {
	A[*no].no = *no + 1
	fmt.Print("Nama Layanan: ")
	fmt.Scan(&A[*no].nama_layanan)
	fmt.Print("biaya : ")
	fmt.Scan(&A[*no].biaya)
	fmt.Print("Metode Pembayaran : ")
	fmt.Scan(&A[*no].metode_pembayaran)
	fmt.Print("Tanggal Pembayaran : ")
	fmt.Scan(&A[*no].tgl_pembayaran)
	fmt.Print("Status : ")
	fmt.Scan(&A[*no].status)
	fmt.Println("======================")
	fmt.Println(A)
	cekIsiArray(A)
	*no++
}
func tampilkanArray(A *tabInt, no int) {
	fmt.Printf("%-4s %-20s %-12s %-10s %-13s %-10s\n", "no", "nama Layanan", "Biaya", "metode", "tgl", "status")
	for i := 0; i < no; i++ {
		fmt.Printf("%-4d %-20s Rp%-10d %-10s %-13s %-10s\n", A[i].no, A[i].nama_layanan, A[i].biaya, A[i].metode_pembayaran, A[i].tgl_pembayaran, A[i].status)
	}
	fmt.Println()
}

func menu(A tabInt){
	var pil, isiArr int
	pil = 0
	for pil != 7{
	fmt.Println("===================== MENU ========================")
	fmt.Println("1. Tampilkan Daftar Layanan")
	fmt.Println("2. Tambahkan Daftar Layanan")
	fmt.Println("3. Edit Daftar Layanan")
	fmt.Println("4. Sort Daftar Layanan")
	fmt.Println("5. Cek Jatuh Tempo")
	fmt.Println("6. Rekomendasi Pengeluaran")
	fmt.Println("7. Keluar")
	fmt.Scan(&pil)
	loadData(&A)
	isiArr = cekIsiArray(&A)
	switch pil{
	case 1 :
		fmt.Println("============ DATA PAGE ===============")
		tampilkanArray(&A, isiArr);
		fmt.Println("data skrg = ", isiArr)
	case 2 : 
	fmt.Println("============ ADD PAGE ===============")
		cekIsiArray(&A)
		tambahLayanan(&A, &isiArr)
		fmt.Println("data skrg = ", isiArr)
	case 3 :
		var index int
		fmt.Println("============ EDIT PAGE ===============")
		fmt.Print("Baris nomor berapa yang ingin diubah? ")
		fmt.Scan(&index)
		fmt.Println("Sebelum edit ")
		fmt.Println(A)
		editArray(&A, index)
		fmt.Println("Sesudah edit ")
		fmt.Println(A)
	case 4 :
		fmt.Println("============ SORT PAGE ===============")
		sortArray(&A, isiArr)
		tampilkanArray(&A, isiArr)
	case 5 :
		fmt.Println("============ CEK JATUH TEMPO PAGE ===============")
		cekJatuhTempo(&A, isiArr)
	case 6 :
		fmt.Println("============ REKOMENDASI PAGE ===============")
		rekomendasiPengeluaran(&A, isiArr)
	}
}
}
func loadData(data *tabInt){
	*data = tabInt{
		{1, "aaa", 1, "ww", "20-302-0", "lunas"},
		{2, "bbb", 8, "ww", "20-302-0", "lunas"},
		{3, "ccc", 3, "ww", "20-302-0", "belum"},
		{4, "ddd", 11, "ww", "20-302-0", "lunas"},
		{5, "eee", 5, "ww", "20-302-0", "lunas"},
	}
}
func main() {
	var data tabInt
	menu(data)
}