package main
import ."fmt"

const NMAX int = 1000
const nPaket int = 48  

type Identitas struct {
    no_pasien int
    nama_pasien string
    umur int
    golongan_darah string
    paket string
    jenis_cek_out [NMAX] string
    qty_out [NMAX] int
    kondisi_out [NMAX] string
    tgl int
    bln int
    tahun int
}

type Paket_MCU struct {
    no_paket int                //indeks dimulai dari 1 dan indeks ke 0 tidak memiliki nilai 
    jenis_check string          //array indeks ke 0 diisi sebagai nama paket dan indeks seterusnya digunakan sebagai string dari jenis pengecekan
    hasil int                   //indeks dimulai dari 1 dan indeks ke 0 tidak memiliki nilai
    kondisi string              //indeks dimulai dari 1 dan indeks ke 0 tidak memiliki nilai
    count_jenis_check int       //Setiap array tipe Bentukan Paket_MCU hanya menyimpan 1 array count_jenis_check pada indeks 0 sebagai penentu berapa banyak jenis cek yang ada
}


// Menampung field dari tipe bentukan Identitas
type arrIdentitas [NMAX] Identitas

// Menampung field dari tipe bentukan Identitas dan berfungsi untuk menampilkan Hasil data yang nantinya akan diurut agar data utama tidak terganggu
type arrTemp [NMAX] Identitas

// Menampung field dari tipe bentukan Paket_MCU
type arrPaketMCU [nPaket][NMAX] Paket_MCU        //array berkonstanta nPaket berfungsi untuk membedakan setiap data paket yang akan ditampung karena di dalam 1 Paket terdapat jenis pengecekan yang berbeda

//Menyimpan jumlah data pasien yang ada
var count_pasien int

//menyimpan jumlah data paket yang ada
var count_paket int

var pasien arrIdentitas
var Paket arrPaketMCU
var Temp arrTemp

func main_menu(pilihan int, pasien *arrIdentitas, paketMCU *arrPaketMCU) {
    Println("=======================================================================================")
    Println(" 1. Daftar Pasien")
    Println(" 2. Daftar Paket Medical Check-up")
    Println(" 3. Rekap Pasien")
    Println(" 99. Keluar Aplikasi")
    Println("=======================================================================================")
    Print("Pilih menu: ");Scan(&pilihan)
    if pilihan == 1 {
        if count_paket == 0 {
            Println("=======================================================================================")
            Println("\n              PAKET MEDICAL CHECKUP BELUM TERSEDIA\n")
            main_menu(pilihan, pasien, &Paket)
        }else {
            Daftar_Pasien(pilihan)
        }
    }else if pilihan == 2 {
        PaketMedicalCheckUp(pilihan)
    }else if pilihan == 3 {
        RekapPasien(pasien)
    }else if pilihan == 99 {
        Println("-------------------    APLIKASI BERAKHIR    ---------------------")
        Println("TERIMA KASIH TELAH MENGGUNAKAN APLIKASI MEDICAL CHECKUP SEDERHANA")
        Println("        Made by MUHAMMAD FADJAR AL FARISYI                       ")
    }
}

//function untuk cek apakah tanggal yang diinput valid atau tidak
func ValidityDate(tgl int, bln int, tahun int)bool{
    if tahun%4 == 0  && tahun > 0 {
        if bln == 2 {
            if tgl > 0 && tgl <= 29 {
                return true
            }else {
                return false
            }
        }else if bln == 1 || bln == 3 || bln == 5 || bln == 7 || bln == 9 || bln == 11 {
            if tgl > 0 && tgl <= 31 {
                return true
            }else {
                return false
            }
        }else if bln == 4 || bln == 6 || bln == 8 || bln == 10 || bln == 12 {
            if tgl > 0 && tgl <= 30 {
                return true
            }else {
                return false
            }
        }else{
            return false
        }
    }else if tahun%4 != 0 && tahun > 0 {
        if bln == 2 {
            if tgl > 0 && tgl <=28 {
                return true
            }else {
                return false
            }
        }else if bln == 1 || bln == 3 || bln == 5 || bln == 7 || bln == 9 || bln == 11 {
            if tgl > 0 && tgl <= 31 {
                return true
            }else {
                return false
            }
        }else if bln == 4 || bln == 6 || bln == 8 || bln == 10 || bln == 12 {
            if tgl > 0 && tgl <= 30 {
                return true
            }else {
                return false
            }
        }else{
            return false
        }
    }else{
        return false
    }
}

//Section Untuk Mendaftar,Mengedit dan Menghapus Pasien
func Daftar_Pasien(pilihan int) {
    Println("1. Tambah Pasien")
    Println("2. Edit Pasien")
    Println("3. Hapus Pasien")
    Println("4. List Pasien")
    Println("99. Kembali")
    Print("Pilih menu: ")
    Scan(&pilihan)
    Println("=======================================================================================")
    if pilihan == 1 {
       Tambah_Pasien(&pasien, Paket)
    }else if pilihan == 2 {
        Edit_Pasien(&pasien) 
    }else if pilihan == 3 {
        Hapus_Pasien(&pasien)
    }else if pilihan == 4 {
        List_Pasien(&pasien)
    }else if pilihan == 99{
        main_menu(pilihan,&pasien, &Paket)
    }else {
        Println("\nDATA INPUT TIDAK VALID\n")
        Daftar_Pasien(pilihan)
    }
}

//Procedure untuk menambahkan pasien
func Tambah_Pasien(pasien *arrIdentitas, paketMCU arrPaketMCU){
    // var count_pasien adalah global
    var pilihan int
    var pilih_paket int
    Print("Masukkan waktu Diagnosa: (DD M/MM YYYY)");Scan(&pasien[count_pasien].tgl, &pasien[count_pasien].bln, &pasien[count_pasien].tahun)
    if ValidityDate(pasien[count_pasien].tgl, pasien[count_pasien].bln, pasien[count_pasien].tahun) {
        Print("Masukkan Nama Pasien: "); Scanln(&pasien[count_pasien].nama_pasien)
        Print("Masukkan Umur Pasien: "); Scanln(&pasien[count_pasien].umur)
        Print("Masukkan Golongan Darah Pasien: "); Scanln(&pasien[count_pasien].golongan_darah)
        Println("---------------------------------------------------------------------------------------")
        for i:=0; i<count_paket; i++ {
               Println(paketMCU[i][0].no_paket,".", paketMCU[i][0].jenis_check)
          }
        Print("Pilih paket yang diinginkan: ");Scan(&pilih_paket)
        if pilih_paket <= count_paket {
            pasien[count_pasien].paket = paketMCU[pilih_paket - 1][0].jenis_check
            Println("Silahkan input hasil diagnosa")
            for k:=1; k<paketMCU[pilih_paket - 1][0].count_jenis_check; k++{
                Println(paketMCU[pilih_paket - 1][k].jenis_check)
                Print("Jumlah:");Scan(&paketMCU[pilih_paket - 1][k].hasil)
                Print("Keterangan:");Scan(&paketMCU[pilih_paket - 1][k].kondisi)
            }
            Println("=======================================================================================")
            Println("\nNama: ", pasien[count_pasien].nama_pasien)
            Println("Umur: ", pasien[count_pasien].umur)
            Println("Golongan Darah: ", pasien[count_pasien].golongan_darah)
            Println("Paket yang dipilih:",pasien[count_pasien].paket)
            Println("----------HASIL DIAGNOSA----------")
            for l:=1; l<paketMCU[pilih_paket - 1][0].count_jenis_check; l++{
                Print(paketMCU[pilih_paket - 1][l].jenis_check, "  ")
                Print(paketMCU[pilih_paket - 1][l].hasil, "  ")
                Print(paketMCU[pilih_paket - 1][l].kondisi)
                Print("\n")
            }
            Println("=======================================================================================")
            Println("Apakah Data sudah benar?")
            Println("1. Sudah")
            Println("2. Belum")
            Print("Pilihan: ");Scan(&pilihan)
            if pilihan == 1 {
                pasien[count_pasien].no_pasien = count_pasien + 1
                for j:=1; j<paketMCU[pilih_paket - 1][0].count_jenis_check; j++{
                    pasien[count_pasien].jenis_cek_out[j - 1] = paketMCU[pilih_paket - 1][j].jenis_check
                    pasien[count_pasien].qty_out[j - 1] = paketMCU[pilih_paket - 1][j].hasil
                    pasien[count_pasien].kondisi_out[j - 1] = paketMCU[pilih_paket - 1][j].kondisi
                }
                count_pasien++
                Daftar_Pasien(pilihan)
            }else if pilihan == 2 {
                Tambah_Pasien(pasien,Paket)
            }else {
                Println("\nINPUT TIDAK VALID, LAKUKAN INPUT ULANG\n")
                Tambah_Pasien(pasien,Paket)
            }
        }else {
            Println("\nINPUT TIDAK VALID, LAKUKAN INPUT ULANG\n")
            Tambah_Pasien(pasien,Paket)
        }
    }else {
        Println("\nTANGGAL TIDAK VALID")
        Tambah_Pasien(pasien,Paket)
    }
}

//Procedure untuk mengubah data pasien
func Edit_Pasien(pasien *arrIdentitas){
    var choose,i, nomor int
    var pilihan,update_umur int
    var update_nama,update_goldar string
    if count_pasien > 0 {
        nomor = 1
        for i=0; i<count_pasien; i++ {
            Println(pasien[i].no_pasien,".", pasien[i].nama_pasien, pasien[i].umur, pasien[i].golongan_darah)
            nomor++
        }
        Print("Data Mana yang ingin diedit?: ");Scan(&choose)
        if Searh_urutan_pasien(choose,pasien) == -1 {
            Println("\n DATA TIDAK DITEMUKAN \n")
            Edit_Pasien(pasien)
        }else {
            Println("=======================================================================================")
            Println("Nama Pasien Sebelumnya adalah: ", pasien[Searh_urutan_pasien(choose,pasien)].nama_pasien)
            Println("Umur Pasien Sebelumnya adalah: ", pasien[Searh_urutan_pasien(choose,pasien)].umur)
            Println("Umur Pasien Sebelumnya adalah: ", pasien[Searh_urutan_pasien(choose,pasien)].golongan_darah)
            Print("Silahkan Masukkan nama Pasien yang Baru: ");Scan(&update_nama)
            pasien[Searh_urutan_pasien(choose,pasien)].nama_pasien = update_nama
            Print("Silahkan Masukkan data umur Pasien yang Baru: ");Scan(&update_umur)
            pasien[Searh_urutan_pasien(choose,pasien)].umur = update_umur
            Print("Silahkan Masukkan data golongan darah Pasien yang Baru: ");Scan(&update_goldar)
            pasien[Searh_urutan_pasien(choose,pasien)].golongan_darah = update_goldar
            Println("NAMA PASIEN BERHASIL DIUBAH")
            Daftar_Pasien(pilihan)
        }
    }else{
        Println("\nDATA BELUM ADA\n")
        Daftar_Pasien(pilihan)
    }
}


//Procedure untuk menghapus data pasien yang ada
func Hapus_Pasien(pasien *arrIdentitas){
    var pilihan,choose int
    var nomor, i int
    if count_pasien > 0 {
    i= 0
    nomor = 1
    for i<count_pasien {
        Println(pasien[i].no_pasien,".",pasien[i].nama_pasien, pasien[i].umur, pasien[i].golongan_darah)
        i++
        nomor++
    }
    Println("Data mana yang ingin dihapus? ")
    Print("Pilihan: ");Scan(&pilihan)
    for pilihan < count_pasien {
        pasien[pilihan - 1].nama_pasien, pasien[pilihan - 1].umur, pasien[pilihan - 1].golongan_darah, pasien[pilihan - 1].tgl,pasien[pilihan - 1].bln, pasien[pilihan - 1].tahun, pasien[pilihan - 1].paket = pasien[pilihan].nama_pasien, pasien[pilihan].umur, pasien[pilihan].golongan_darah, pasien[pilihan].tgl, pasien[pilihan].bln, pasien[pilihan].tahun, pasien[pilihan].paket
        pilihan++
    }
        count_pasien--
    Println("\nData berhasil dihapus!")
    Daftar_Pasien(choose)
    }else {
        Println("DATA BELUM ADA\n")
        Daftar_Pasien(choose)
    }
}

//Procedure untuk menampilkan list pasien yang sudah terdaftar
func List_Pasien(pasien *arrIdentitas){
    var i, nomor int
    var pilihan int
    if count_pasien > 0 {
        i = 0
        nomor = 1
        for i<count_pasien {
           Println(nomor,". ",pasien[i].nama_pasien, pasien[i].umur, pasien[i].golongan_darah, pasien[i].paket)
            i++
            nomor++
        }
        Println("=======================================================================================")
        Daftar_Pasien(pilihan)
    }else {
        Println("DATA MASIH KOSONG SILAHKAN ISI DATA PASIEN!!!")
        Daftar_Pasien(pilihan)
    }
}

//Funtion yang mencari indeks array ke berapa dari sebuah tipe bentukan
func Searh_urutan_pasien(choose int, pasien *arrIdentitas)int {
    var idx,left,right,mid int 
    idx = -1
    left = 0
    right = count_pasien
    mid = (left + right) / 2
    for left <= right && idx == -1 && count_pasien >= choose{
        if pasien[mid].no_pasien == choose {
            idx = mid
        }else if pasien[mid].no_pasien > choose {
            right = mid - 1
        }else if pasien[mid].no_pasien < choose {
            left = mid + 1
        }
        mid = (left + right) / 2
    }    
    return idx 
}


//Section untuk menu dari Paket Medical Checkup 
func PaketMedicalCheckUp(pilihan int) {
    if count_paket <= 0 {
        Println("=======================================================================================")
        Println("Paket belum Tersedia\n")
    }else {
        Println("=======================================================================================")
        Println("Paket yang tersedia:", count_paket,"\n")
    }
    Println("1. Tambah Paket")
    Println("2. Edit nama Paket")
    Println("3. Hapus Paket")
    Println("99. Kembali")
    Println("=======================================================================================")
    Print("Pilihan anda: ");Scan(&pilihan)
    if pilihan == 1 {
        TambahPaket(&Paket)
    }else if pilihan == 2 {
        EditPaket(&Paket)
    }else if pilihan == 3 {
        HapusPaket(&Paket)
    }else if pilihan == 99 {
        main_menu(pilihan,&pasien, &Paket)
    }else {
        Println("=======================================================================================")
        Println("Inputan tidak valid")
        PaketMedicalCheckUp(pilihan)
    }
}

//Procedure untuk menambahkan paket 
func TambahPaket(paketMCU *arrPaketMCU){
    var pilihan,i int
    i=0
    Println("=======================================================================================")
    Print("Masukkan nama Paket: "); Scan(&paketMCU[count_paket][0].jenis_check)
    Println("Masukkan Kriteria yang akan di cek pada Paket: (Ketik DONE untuk mengakhiri input) ")
    for paketMCU[count_paket][i].jenis_check != "DONE"  {
        i++
        Scan(&paketMCU[count_paket][i].jenis_check)
        paketMCU[count_paket][0].count_jenis_check++
    }
    count_paket++
    paketMCU[count_paket - 1][0].no_paket = count_paket
    Println("=======================================================================================")
    Println("\n               PAKET BERHASIL DITAMBAHKAN\n")
    PaketMedicalCheckUp(pilihan)
}

//Procedure untuk mengubah nama paket
func EditPaket(paketMCU *arrPaketMCU){
    var choose,i int
    var pilihan int
    var update string
    if count_paket > 0 {
        Println("=======================================================================================")
        for i=0; i<count_paket; i++ {
            Println(paketMCU[i][0].no_paket,".", paketMCU[i][0].jenis_check)
        }
        Println("=======================================================================================")
        Print("Paket Mana yang ingin diedit?: ");Scan(&choose)
        Println("=======================================================================================")
        if Searh_urutan(choose,paketMCU) == -1 {
            Println("\n Silahkan masukkan angka yang valid \n")
            EditPaket(paketMCU)
        }else {
            Println("Nama Paket Sebelumnya adalah: ", paketMCU[Searh_urutan(choose,paketMCU)][0].jenis_check)
            Print("Silahkan Masukkan nama Paket yang Baru: ");Scan(&update)
            paketMCU[Searh_urutan(choose,paketMCU)][0].jenis_check = update
            Println("=======================================================================================")
            Println("\n               NAMA PAKET BERHASIL DIUBAH\n")
            PaketMedicalCheckUp(pilihan)
        }
    }else{
        Println("=======================================================================================")
        Println("\nPAKET BELUM TERSEDIA SILAHKAN MENAMBAHKAN PAKET!!!\n")
        PaketMedicalCheckUp(pilihan)
    }
}

//Procedure untuk menghapus paket yang ada
func HapusPaket(paketMCU *arrPaketMCU){
    var choose,i,j, nomor int
    var pilihan int
    nomor = 1
    if count_paket > 0 {
        Println("=======================================================================================")
        for i=0; i<count_paket; i++ {
            Println(paketMCU[i][0].no_paket,".", paketMCU[i][0].jenis_check)
            nomor++
        }
        Print("Paket Mana yang ingin dihapus?: ");Scan(&choose)
        Println("=======================================================================================")
        if Searh_urutan(choose,paketMCU) == -1 {
            Println("\nSilahkan Masukkan angka yang valid\n")
            HapusPaket(paketMCU)
        }else {
            for j=choose - 1; j<=count_paket; j++{
                paketMCU[j], paketMCU[j][0].no_paket = paketMCU[j + 1], j+1
            }
            Println("Paket Berhasil di update")
            count_paket--
            PaketMedicalCheckUp(pilihan)
        }
    }else {
        Println("\nPAKET BELUM TERSEDIA SILAHKAN MENAMBAHKAN PAKET!!!\n")
        PaketMedicalCheckUp(pilihan)
    }    
}

//Fungsi mencari urutan paket yang akan dicari menggunakan Binary Searh
func Searh_urutan(choose int, paketMCU *arrPaketMCU)int {
    var idx,left,right,mid int 
    idx = -1
    left = 0
    right = count_paket
    mid = (left + right) / 2
    for left <= right && idx == -1 && count_paket >= choose{
        if paketMCU[mid][0].no_paket == choose {
            idx = mid
        }else if paketMCU[mid][0].no_paket > choose {
            right = mid - 1
        }else if paketMCU[mid][0].no_paket < choose {
            left = mid + 1
        }
        mid = (left + right) / 2
    }    
    return idx 
}

//Procedure untuk menampilkan 
func RekapPasien(pasien *arrIdentitas){
    var pilihan int
    var choose int
    var pilih int
    var index_paket int
    Println("=======================================================================================")
    //arrTemp menampung dari semua data dari arrIdentitas yang berguna agar membedakan data yang akan di Sorting dan data Default
    for i:=0; i<=count_pasien; i++{
        Temp[i].no_pasien = pasien[i].no_pasien
        Temp[i].nama_pasien = pasien[i].nama_pasien
        Temp[i].umur = pasien[i].umur
        Temp[i].golongan_darah = pasien[i].golongan_darah
        Temp[i].paket = pasien[i].paket
        Temp[i].tgl = pasien[i].tgl
        Temp[i].bln = pasien[i].bln
        Temp[i].tahun = pasien[i].tahun
    }
    Println("---------------------------------------------------------------------------------------")
    if count_pasien == 0 {
        Println("\n                  REKAP DATA PASIEN MASIH KOSONG\n")
    }
    //Yang ditampilkan disini adalah data Default
    for i:=0; i<=count_pasien - 1; i++ {
        Println(pasien[i].no_pasien,".","NAMA:", pasien[i].nama_pasien,pasien[i].golongan_darah,"UMUR:", pasien[i].umur,"Paket yang dipilih:" , pasien[i].paket ,"WAKTU PEMERIKSAAN:", pasien[i].tgl, "/" ,pasien[i].bln, "/", pasien[i].tahun)
    }
    Println("---------------------------------------------------------------------------------------")
    Println("1. Cari berdasarkan Nama")
    Println("2. Cari berdasarkan waktu")
    Println("3. Cari berdasarkan periode")
    Println("4. Cari berdasarkan paket")
    Println("5. Urutkan berdasarkan Umur")
    Println("6. Urutkan berdasarkan Waktu Pengecekan")
    Println("7. Cek Detail")
    Println("99. Kembali")
    Println("=======================================================================================")
    Print("Pilihan anda:");Scan(&choose)
    if choose == 1 {
        SearchNAME(pasien)
    }else if choose == 2 {
        SearchTIME(pasien)
    }else if choose == 3 {
        SearchPeriod(pasien)
    }else if choose == 4 {
        SearchPaket(pasien)
    }else if choose == 5 {
        Println("1.ASCENDING")
        Println("2.DESCENDING\n")
        Print("Pilihan:");Scan(&choose)
        if choose == 1 {
            SortUmurASCENDING(pasien)
        }else if choose == 2{
            SortUmurDESCENDING(pasien)
        }else{
            RekapPasien(pasien)
        }
    }else if choose == 6 {
        Println("1. ASCENDING")
        Println("2. DESCENDING\n")
        Print("Pilihan:");Scan(&choose)
        if choose == 1 {
            SortTimeASCENDING(pasien)
        }else if choose == 2{
            SortTimeDESCENDING(pasien)
        }else {
            RekapPasien(pasien)
        }
    }else if choose == 7 {
        Print("Pilih Data:");Scan(&pilih)
        Println("=======================================================================================")
        Println("NAMA:",pasien[pilih-1].nama_pasien)
        Println("UMUR:",pasien[pilih-1].umur)
        Println("GOLONGAN DARAH:",pasien[pilih-1].golongan_darah)
        Println("PAKET YANG DIPILIH:",pasien[pilih-1].paket)
        Println("WAKTU PENGECEKAN:",pasien[pilih-1].tgl, "/", pasien[pilih-1].bln, "/", pasien[pilih-1].tahun)
        index_paket = CariIndexPaket(pasien[pilih-1].paket, &Paket)
        for i:=0; i<Paket[index_paket][0].count_jenis_check - 1; i++ {
            Println("  *  ", pasien[pilih - 1].jenis_cek_out[i], pasien[pilih - 1].qty_out[i], pasien[pilih - 1].kondisi_out[i])
        }
        Println("=======================================================================================")
        RekapPasien(pasien)
    }else if choose == 99 {
        main_menu(pilihan,pasien,&Paket)
    }else {
        Println("INPUT TIDAK VALID")
        RekapPasien(pasien)
    }
}

func CariIndexPaket(x string, Paket *arrPaketMCU)int {
    var idx,i int
    idx = -1
    i = 0
    for idx == -1 {
        if Paket[i][0].jenis_check == x {
            idx = i
        }
        i++
    }
    return idx
}

//Procedure untuk mencari data berdsarkan Nama
func SearchNAME(pasien *arrIdentitas){
    var search string
    Print("Masukkan nama yang ingin dicari:");Scan(&search)
    if SearchbyName(search) == -1 {
        Println("=======================================================================================")
        Println("\nDATA TIDAK DITEMUKAN\n")
        Println("=======================================================================================")
    }else {
        Println("=======================================================================================")
        Println("\nDATA DITEMUKAN\n")
        Println(pasien[SearchbyName(search)].no_pasien,".","NAMA:", pasien[SearchbyName(search)].nama_pasien,"UMUR:", pasien[SearchbyName(search)].umur,"Paket yang dipilih:" , pasien[SearchbyName(search)].paket ,"WAKTU PEMERIKSAAN:", pasien[SearchbyName(search)].tgl, "/" ,pasien[SearchbyName(search)].bln, "/", pasien[SearchbyName(search)].tahun)
        Println("=======================================================================================")
    }
    RekapPasien(pasien)    
}


func SearchbyName(search string)int{
    var idx,i int 
    idx = -1
    i = 0
    for idx == -1 && i<count_pasien {
        if pasien[i].nama_pasien == search {
            idx = i
        }
        i++
    }
    return idx
}

func SearchTIME(pasien *arrIdentitas){
    var day,month,tahun int
    Println("Masukkan waktu (DD M/MM YYYY):");Scan(&day, &month, &tahun)
    if SearchbyTime(day,month,tahun) == -1 {
        Println("=======================================================================================")
        Println("\nDATA TIDAK DITEMUKAN\n")
        Println("=======================================================================================")
    }else {
        Println("=======================================================================================")
        Println("DATA DITEMUKAN\n")
        Println(pasien[SearchbyTime(day,month,tahun)].no_pasien,".","NAMA:", pasien[SearchbyTime(day,month,tahun)].nama_pasien,"UMUR:", pasien[SearchbyTime(day,month,tahun)].umur,"Paket yang dipilih:" , pasien[SearchbyTime(day,month,tahun)].paket ,"WAKTU PEMERIKSAAN:", pasien[SearchbyTime(day,month,tahun)].tgl, "/" ,pasien[SearchbyTime(day,month,tahun)].bln, "/", pasien[SearchbyTime(day,month,tahun)].tahun)
        Println("=======================================================================================")
    }
    RekapPasien(pasien)
}

func SearchbyTime(day,month,tahun int)int {
    var idx,i int
    idx = -1
    i = 0
    for idx == -1 && i<count_pasien {
        if pasien[i].tgl == day && pasien[i].bln == month && pasien[i].tahun == tahun {
            idx = i 
        }
        i++
    }
    return idx
}

func SearchPeriod(pasien *arrIdentitas){
    var search_bln, search_thn int
    var i,found int
    var choose int
    i = 0 
    Print("Berdasarkan apa? 1. Tahun/2. Bulan");Scan(&choose)
    if choose == 1 {
        Print("Masukkan Tahun: ");Scan(&search_thn)
        Println("=======================================================================================")
        for i<count_pasien {
            if pasien[i].tahun == search_thn {
                Println(pasien[i].no_pasien,".", pasien[i].nama_pasien,pasien[i].golongan_darah,"UMUR:", pasien[i].umur,"PAKET YANG DIPILIH:", pasien[i].paket, "WAKTU PENGECEKAN: ", pasien[i].tgl,"/", pasien[i].bln, "/", pasien[i].tahun)
            }
            i++
            found++
        }
        if found == 0 {
        Println("\nDATA TIDAK DITEMUKAN!!!\n")
    }
    Println("=================================================================")
    RekapPasien(pasien)
    }else if choose == 2 {
        Print("Masukkan Tahun: ");Scan(&search_thn)
        Print("Masukkan Bulan: ");Scan(&search_bln)
        Println("=======================================================================================")
        for i<count_pasien {
            if pasien[i].tahun == search_thn && pasien[i].bln == search_bln {
                Println(pasien[i].no_pasien, pasien[i].nama_pasien, pasien[i].umur, pasien[i].paket, "WAKTU PENGECEKAN: ", pasien[i].tgl,"/", pasien[i].bln, "/", pasien[i].tahun)
            }
            i++
            found++
        }
        if found == 0 {
        Println("\nDATA TIDAK DITEMUKAN!!!\n")
    }
    Println("=======================================================================================")
    RekapPasien(pasien)
    }else {
        Println("\nINPUT TIDAK VALID\n")
        RekapPasien(pasien)
    }
}

func SearchPaket(pasien *arrIdentitas){
    var search string
    var i,found int
    i=0
    Print("Input nama paket:");Scan(&search)
    Println("=======================================================================================")
    for i<count_pasien {
        if pasien[i].paket == search {
            Println(pasien[i].no_pasien, pasien[i].nama_pasien, pasien[i].umur, pasien[i].paket)
            }
        i++
        found++
    }
    if found == 0 {
        Println("\nDATA TIDAK DITEMUKAN!!!\n")
    }
    Println("=======================================================================================")
    RekapPasien(pasien)
}

//Sort dt berdasarkan Umur menggunakan Insertion Sort
func SortUmurASCENDING(pasien *arrIdentitas){
    var pass,i int
    var temp_nama, temp_paket, temp_gol_darah string
    var temp_umur, temp_tgl, temp_bln, temp_tahun, temp_urut int
    pass = 0 
    for pass < count_pasien - 1 {
        i = pass + 1
        temp_urut = Temp[pass + 1].no_pasien
        temp_nama = Temp[pass + 1].nama_pasien
        temp_umur = Temp[pass + 1].umur
        temp_paket = Temp[pass + 1].paket
        temp_gol_darah = Temp[pass + 1].golongan_darah
        temp_tgl = Temp[pass + 1].tgl
        temp_bln = Temp[pass + 1].bln
        temp_tahun = Temp[pass + 1].tahun
        for i>0 && temp_umur < Temp[i - 1].umur {
            Temp[i].no_pasien = Temp[i - 1].no_pasien
            Temp[i].nama_pasien = Temp[i - 1].nama_pasien
            Temp[i].umur = Temp[i - 1].umur
            Temp[i].paket = Temp[i - 1].paket
            Temp[i].golongan_darah = Temp[i - 1].golongan_darah
            Temp[i].tgl = Temp[i - 1].tgl
            Temp[i].bln = Temp[i - 1].bln
            Temp[i].tahun = Temp[i - 1].tahun
            i--
        }
        Temp[i].no_pasien = temp_urut
        Temp[i].nama_pasien = temp_nama
        Temp[i].umur = temp_umur
        Temp[i].paket = temp_paket
        Temp[i].golongan_darah = temp_gol_darah
        Temp[i].tgl = temp_tgl
        Temp[i].bln = temp_bln
        Temp[i].tahun = temp_tahun
        pass++
    }
    Println("=======================================================================================")
    Println("\nHASIL SORT ASCENDING by UMUR\n")
    for j:=0; j<=count_pasien - 1; j++ {
        Println(Temp[j].no_pasien,".","NAMA:", Temp[j].nama_pasien,Temp[i].golongan_darah,"UMUR:", Temp[j].umur,"Paket yang dipilih:" ,Temp[j].paket, "WAKTU PEMERIKSAAN", Temp[j].tgl,"/", Temp[j].bln, "/", Temp[j].tahun)
    }
    Println("=======================================================================================")
    RekapPasien(pasien)
}
//Insertion SORT
func SortUmurDESCENDING(pasien *arrIdentitas){
    var pass,i int
    var temp_nama, temp_paket, temp_gol_darah string
    var temp_umur, temp_tgl, temp_bln, temp_tahun,temp_urut int
    pass = 0 
    for pass < count_pasien-1 {
        i = pass + 1
        temp_urut = Temp[pass + 1].no_pasien
        temp_nama = Temp[pass + 1].nama_pasien
        temp_umur = Temp[pass + 1].umur
        temp_paket = Temp[pass + 1].paket
        temp_gol_darah = Temp[pass + 1].golongan_darah
        temp_tgl = Temp[pass + 1].tgl
        temp_bln = Temp[pass + 1].bln
        temp_tahun = Temp[pass + 1].tahun
        for i>0 && temp_umur > Temp[i - 1].umur {
            Temp[i].no_pasien = Temp[i - 1].no_pasien
            Temp[i].nama_pasien = Temp[i - 1].nama_pasien
            Temp[i].umur = Temp[i - 1].umur
            Temp[i].paket = Temp[i - 1].paket
            Temp[i].golongan_darah = Temp[i - 1].golongan_darah
            Temp[i].tgl = Temp[i - 1].tgl
            Temp[i].bln = Temp[i - 1].bln
            Temp[i].tahun = Temp[i - 1].tahun
            i--
        }
        Temp[i].no_pasien = temp_urut
        Temp[i].nama_pasien = temp_nama
        Temp[i].umur = temp_umur
        Temp[i].paket = temp_paket
        Temp[i].golongan_darah = temp_gol_darah
        Temp[i].tgl = temp_tgl
        Temp[i].bln = temp_bln
        Temp[i].tahun = temp_tahun
        pass++
    }
    Println("=======================================================================================")
    Println("\nHASIL SORT DESCENDING by UMUR\n")
    for j:=0; j<count_pasien; j++ {
        Println(Temp[j].no_pasien,".","NAMA:", Temp[j].nama_pasien,Temp[j].golongan_darah,"UMUR:", Temp[j].umur,"Paket yang dipilih:" ,Temp[j].paket, "WAKTU PEMERIKSAAN", Temp[j].tgl,"/", Temp[j].bln, "/", Temp[j].tahun)
    }
    Println("=======================================================================================")
    RekapPasien(pasien)
}

//Procedure untuk mengurutkan data pasien menggunakan Selection sort
func SortTimeASCENDING(pasien *arrIdentitas){
    // var idx,idx2,idx3,pass,pass2,pass3 int
    var idx,idx2,idx3,pass,pass2,pass3,i,j,k int
    var temp_urut,temp_umur, temp_tgl, temp_bln, temp_tahun int
    // var i,j,k int
    var temp_gol_darah, temp_paket, temp_nama string
    pass = 0 
    for pass < count_pasien - 1 {
        idx = pass 
        i = pass + 1
        for i < count_pasien {
            if Temp[i].tgl < Temp[idx].tgl {
                idx = i
            }
            i++
        }
        temp_urut = Temp[idx].no_pasien
        temp_nama = Temp[idx].nama_pasien
        temp_umur = Temp[idx].umur
        temp_gol_darah = Temp[idx].golongan_darah
        temp_paket = Temp[idx].paket
        temp_tgl = Temp[idx].tgl
        temp_bln = Temp[idx].bln
        temp_tahun = Temp[idx].tahun
        
        Temp[idx].no_pasien = Temp[pass].no_pasien
        Temp[idx].nama_pasien = Temp[pass].nama_pasien
        Temp[idx].umur = Temp[pass].umur
        Temp[idx].golongan_darah = Temp[pass].golongan_darah
        Temp[idx].paket = Temp[pass].paket
        Temp[idx].tgl = Temp[pass].tgl
        Temp[idx].bln = Temp[pass].bln
        Temp[idx].tahun = Temp[pass].tahun
        
        Temp[pass].no_pasien = temp_urut
        Temp[pass].nama_pasien = temp_nama
        Temp[pass].umur = temp_umur
        Temp[pass].golongan_darah = temp_gol_darah
        Temp[pass].paket = temp_paket
        Temp[pass].tgl = temp_tgl
        Temp[pass].bln = temp_bln
        Temp[pass].tahun = temp_tahun
        pass++
    }
    
    pass2 = 0 
    for pass2 < count_pasien - 1 {
        idx2 = pass2 
        j = pass2 + 1
        for j < count_pasien {
            if Temp[j].bln < Temp[idx2].bln {
                idx2 = j
            }
            j++
        }
        temp_urut = Temp[idx2].no_pasien
        temp_nama = Temp[idx2].nama_pasien
        temp_umur = Temp[idx2].umur
        temp_gol_darah = Temp[idx2].golongan_darah
        temp_paket = Temp[idx2].paket
        temp_tgl = Temp[idx2].tgl
        temp_bln = Temp[idx2].bln
        temp_tahun = Temp[idx2].tahun
        
        Temp[idx2].no_pasien = Temp[pass2].no_pasien
        Temp[idx2].nama_pasien = Temp[pass2].nama_pasien
        Temp[idx2].umur = Temp[pass2].umur
        Temp[idx2].golongan_darah = Temp[pass2].golongan_darah
        Temp[idx2].paket = Temp[pass2].paket
        Temp[idx2].tgl = Temp[pass2].tgl
        Temp[idx2].bln = Temp[pass2].bln
        Temp[idx2].tahun = Temp[pass2].tahun
        
        Temp[pass2].no_pasien = temp_urut
        Temp[pass2].nama_pasien = temp_nama
        Temp[pass2].umur = temp_umur
        Temp[pass2].golongan_darah = temp_gol_darah
        Temp[pass2].paket = temp_paket
        Temp[pass2].tgl = temp_tgl
        Temp[pass2].bln = temp_bln
        Temp[pass2].tahun = temp_tahun
        pass2++
    }
    
    pass3 = 0 
    for pass3 < count_pasien - 1 {
        idx3 = pass3 
        k = pass3 + 1
        for k < count_pasien {
            if Temp[k].tahun < Temp[idx3].tahun {
                idx3 = k
            }
            k++
        }
        temp_urut = Temp[idx3].no_pasien
        temp_nama = Temp[idx3].nama_pasien
        temp_umur = Temp[idx3].umur
        temp_gol_darah = Temp[idx3].golongan_darah
        temp_paket = Temp[idx3].paket
        temp_tgl = Temp[idx3].tgl
        temp_bln = Temp[idx3].bln
        temp_tahun = Temp[idx3].tahun
        
        Temp[idx3].no_pasien = Temp[pass3].no_pasien
        Temp[idx3].nama_pasien = Temp[pass3].nama_pasien
        Temp[idx3].umur = Temp[pass3].umur
        Temp[idx3].golongan_darah = Temp[pass3].golongan_darah
        Temp[idx3].paket = Temp[pass3].paket
        Temp[idx3].tgl = Temp[pass3].tgl
        Temp[idx3].bln = Temp[pass3].bln
        Temp[idx3].tahun = Temp[pass3].tahun
        
        Temp[pass3].no_pasien = temp_urut
        Temp[pass3].nama_pasien = temp_nama
        Temp[pass3].umur = temp_umur
        Temp[pass3].golongan_darah = temp_gol_darah
        Temp[pass3].paket = temp_paket
        Temp[pass3].tgl = temp_tgl
        Temp[pass3].bln = temp_bln
        Temp[pass3].tahun = temp_tahun
        
        pass3++
    }
    Println("=======================================================================================")
    Println("HASIL SORT ASCENDING by TIME\n")
    for l:=0; l<count_pasien; l++ {
        Println(Temp[l].no_pasien,".","NAMA:", Temp[l].nama_pasien,Temp[l].golongan_darah,"UMUR:", Temp[l].umur,"Paket yang dipilih:" ,Temp[l].paket, "WAKTU PEMERIKSAAN", Temp[l].tgl,"/", Temp[l].bln, "/", Temp[l].tahun)

    }
    Println("=======================================================================================")
    RekapPasien(pasien)
}

func SortTimeDESCENDING(pasien *arrIdentitas){
    var idx,idx2,idx3,pass,pass2,pass3 int
    var temp_urut,temp_umur, temp_tgl, temp_bln, temp_tahun,i,j,k int
    var temp_gol_darah, temp_paket, temp_nama string
    pass = 0 
    for pass < count_pasien - 1 {
        idx = pass 
        i = pass + 1
        for i < count_pasien {
            if Temp[i].tgl > Temp[idx].tgl {
                idx = i
            }
            i++
        }
        temp_urut = Temp[idx].no_pasien
        temp_nama = Temp[idx].nama_pasien
        temp_umur = Temp[idx].umur
        temp_gol_darah = Temp[idx].golongan_darah
        temp_paket = Temp[idx].paket
        temp_tgl = Temp[idx].tgl
        temp_bln = Temp[idx].bln
        temp_tahun = Temp[idx].tahun
        
        Temp[idx].no_pasien = Temp[pass].no_pasien
        Temp[idx].nama_pasien = Temp[pass].nama_pasien
        Temp[idx].umur = Temp[pass].umur
        Temp[idx].golongan_darah = Temp[pass].golongan_darah
        Temp[idx].paket = Temp[pass].paket
        Temp[idx].tgl = Temp[pass].tgl
        Temp[idx].bln = Temp[pass].bln
        Temp[idx].tahun = Temp[pass].tahun
        
        Temp[pass].no_pasien = temp_urut
        Temp[pass].nama_pasien = temp_nama
        Temp[pass].umur = temp_umur
        Temp[pass].golongan_darah = temp_gol_darah
        Temp[pass].paket = temp_paket
        Temp[pass].tgl = temp_tgl
        Temp[pass].bln = temp_bln
        Temp[pass].tahun = temp_tahun
        pass++
    }
    
    pass2 = 0 
    for pass2 < count_pasien - 1 {
        idx2 = pass2 
        j = pass2 + 1
        for j < count_pasien {
            if Temp[j].bln > Temp[idx2].bln {
                idx2 = j
            }
            j++
        }
        temp_urut = Temp[idx2].no_pasien
        temp_nama = Temp[idx2].nama_pasien
        temp_umur = Temp[idx2].umur
        temp_gol_darah = Temp[idx2].golongan_darah
        temp_paket = Temp[idx2].paket
        temp_tgl = Temp[idx2].tgl
        temp_bln = Temp[idx2].bln
        temp_tahun = Temp[idx2].tahun
        
        Temp[idx2].no_pasien = Temp[pass2].no_pasien
        Temp[idx2].nama_pasien = Temp[pass2].nama_pasien
        Temp[idx2].umur = Temp[pass2].umur
        Temp[idx2].golongan_darah = Temp[pass2].golongan_darah
        Temp[idx2].paket = Temp[pass2].paket
        Temp[idx2].tgl = Temp[pass2].tgl
        Temp[idx2].bln = Temp[pass2].bln
        Temp[idx2].tahun = Temp[pass2].tahun
        
        Temp[pass2].no_pasien = temp_urut
        Temp[pass2].nama_pasien = temp_nama
        Temp[pass2].umur = temp_umur
        Temp[pass2].golongan_darah = temp_gol_darah
        Temp[pass2].paket = temp_paket
        Temp[pass2].tgl = temp_tgl
        Temp[pass2].bln = temp_bln
        Temp[pass2].tahun = temp_tahun
        pass2++
    }
    
    pass3 = 0 
    for pass3 < count_pasien - 1 {
        idx3 = pass3 
        k = pass3 + 1
        for k < count_pasien {
            if Temp[k].tahun > Temp[idx3].tahun {
                idx3 = k
            }
            k++
        }
        temp_urut = Temp[idx3].no_pasien
        temp_nama = Temp[idx3].nama_pasien
        temp_umur = Temp[idx3].umur
        temp_gol_darah = Temp[idx3].golongan_darah
        temp_paket = Temp[idx3].paket
        temp_tgl = Temp[idx3].tgl
        temp_bln = Temp[idx3].bln
        temp_tahun = Temp[idx3].tahun
        
        Temp[idx3].no_pasien = Temp[pass3].no_pasien
        Temp[idx3].nama_pasien = Temp[pass3].nama_pasien
        Temp[idx3].umur = Temp[pass3].umur
        Temp[idx3].golongan_darah = Temp[pass3].golongan_darah
        Temp[idx3].paket = Temp[pass3].paket
        Temp[idx3].tgl = Temp[pass3].tgl
        Temp[idx3].bln = Temp[pass3].bln
        Temp[idx3].tahun = Temp[pass3].tahun
        
        Temp[pass3].no_pasien = temp_urut
        Temp[pass3].nama_pasien = temp_nama
        Temp[pass3].umur = temp_umur
        Temp[pass3].golongan_darah = temp_gol_darah
        Temp[pass3].paket = temp_paket
        Temp[pass3].tgl = temp_tgl
        Temp[pass3].bln = temp_bln
        Temp[pass3].tahun = temp_tahun
        
        pass3++
    }
    Println("=======================================================================================")
    Println("HASIL SORT DESCENDING by TIME\n")
    for l:=0; l<count_pasien; l++ {
        Println(Temp[l].no_pasien,".","NAMA:", Temp[l].nama_pasien,Temp[l].golongan_darah,"UMUR:", Temp[l].umur,"Paket yang dipilih:" ,Temp[l].paket, "WAKTU PEMERIKSAAN", Temp[l].tgl,"/", Temp[l].bln, "/", Temp[l].tahun)

    }
    Println("=======================================================================================")
    RekapPasien(pasien)
}

func main(){
    var n int
    Println("\n            SELAMAT DATANG DI APLIKASI MEDICAL CHECK-UP SEDERHANA\n   ")
    main_menu(n,&pasien, &Paket)
}
