package main

import (
	"fmt"
)

type Item struct {
	ID    int
	Name  string
	Price float64
	Stock int
}

type Transaction struct {
	ItemName string
	Qty      int
	Total    float64
}

type MiniMart struct {
	Items        [100]Item        // Array statis
	Transactions [100]Transaction // Array statis
	ItemCount    int              // Untuk lacak jumlah item aktif
	TransCount   int              // Untuk lacak jumlah transaksi aktif
}

// Search item by name (sequential) -> returns index or -1
func SearchItemByName(m *MiniMart, name string) int {
	for i := 0; i < m.ItemCount; i++ {
		if m.Items[i].Name == name {
			return i
		}
	}
	return -1
}

// Binary search by ID -> returns index or -1
// PENTING: Array Items harus sudah terurut berdasarkan ID agar fungsi ini bekerja.
func BinarySearchByID(m *MiniMart, id int) int {
	low, high := 0, m.ItemCount-1
	for low <= high {
		mid := (low + high) / 2
		if m.Items[mid].ID == id {
			return mid
		} else if m.Items[mid].ID < id {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func InsertionSortByName(m *MiniMart) {
	for i := 1; i < m.ItemCount; i++ {
		key := m.Items[i]
		j := i - 1
		for j >= 0 && m.Items[j].Name > key.Name {
			m.Items[j+1] = m.Items[j]
			j--
		}
		m.Items[j+1] = key
	}
}

func SelectionSortByPrice(m *MiniMart) {
	for i := 0; i < m.ItemCount; i++ {
		minIdx := i
		for j := i + 1; j < m.ItemCount; j++ {
			if m.Items[j].Price < m.Items[minIdx].Price {
				minIdx = j
			}
		}
		// m.Items[i], m.Items[minIdx] = m.Items[minIdx], m.Items[i]

		temp := m.Items[i]
		m.Items[i] = m.Items[minIdx]
		m.Items[minIdx] = temp
	}
}

func AddItem(m *MiniMart, id int, name string, price float64, stock int) {
	if m.ItemCount >= len(m.Items) {
		fmt.Println("Kapasitas barang penuh!")
		return
	}
	m.Items[m.ItemCount] = Item{
		ID:    id,
		Name:  name,
		Price: price,
		Stock: stock,
	} // tidak sesuai modul
	m.ItemCount++
	fmt.Printf("Barang %s berhasil ditambahkan dengan ID %d\n", name, id)
}

func EditItem(m *MiniMart, id int, name string, price float64, stock int) {
	// Memanggil BinarySearchByID untuk menemukan indeks barang
	index := BinarySearchByID(m, id)

	if index != -1 {
		m.Items[index].Name = name
		m.Items[index].Price = price
		m.Items[index].Stock = stock
		fmt.Printf("Barang dengan ID %d berhasil diubah\n", id)
	} else {
		fmt.Println("Barang dengan ID tersebut tidak ditemukan")
	}
}

func DisplayItems(m *MiniMart) {
	if m.ItemCount == 0 {
		fmt.Println("Tidak ada barang di minimart")
		return
	}
	fmt.Println("\nDaftar Barang di Minimart:")
	fmt.Println("ID\tNama\t\tHarga\t\tStok")
	fmt.Println("==================================================")
	for i := 0; i < m.ItemCount; i++ {
		item := m.Items[i]
		fmt.Printf("%d\t%s\t\t%.2f\t\t%d\n", item.ID, item.Name, item.Price, item.Stock)
	}
	fmt.Println("==================================================")
}

func AddTransaction(m *MiniMart, name string, qty int) {
	// Menggunakan sequential search di sini karena barang mungkin tidak terurut berdasarkan nama
	index := SearchItemByName(m, name)

	if index != -1 && m.Items[index].Stock >= qty {
		total := float64(qty) * m.Items[index].Price
		m.Items[index].Stock -= qty

		if m.TransCount >= len(m.Transactions) {
			fmt.Println("Kapasitas transaksi penuh!")
			return
		}

		m.Transactions[m.TransCount] = Transaction{
			ItemName: name,
			Qty:      qty,
			Total:    total,
		}
		m.TransCount++
		fmt.Printf("Transaksi berhasil: %s x%d = Rp%.2f\n", name, qty, total)
		return
	}

	fmt.Println("Barang tidak tersedia atau stok kurang")
}

func DisplayTransactions(m *MiniMart) {
	if m.TransCount == 0 {
		fmt.Println("Belum ada transaksi hari ini.")
		return
	}
	fmt.Println("\nDaftar Transaksi Hari Ini:")
	fmt.Println("==================================================")
	totalOmzet := 0.0
	for i := 0; i < m.TransCount; i++ {
		t := m.Transactions[i]
		fmt.Printf("%s x%d = Rp%.2f\n", t.ItemName, t.Qty, t.Total)
		totalOmzet += t.Total
	}
	fmt.Println("==================================================")
	fmt.Printf("Total Omzet Hari Ini: Rp%.2f\n", totalOmzet)
}

func main() {
	var minimart MiniMart
	var choice int
	for {
		fmt.Println("\n=== Aplikasi Kasir Minimart ===")
		fmt.Println("1. Tambah Barang")
		fmt.Println("2. Ubah Barang")
		fmt.Println("3. Hapus Barang")
		fmt.Println("4. Tampilkan Semua Barang")
		fmt.Println("5. Catat Transaksi")
		fmt.Println("6. Tampilkan Transaksi dan Omzet")
		fmt.Println("7. Urutkan Barang (Nama - Insertion Sort)")
		fmt.Println("8. Urutkan Barang (Harga - Selection Sort)")
		fmt.Println("9. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var id int
			var name string
			var price float64
			var stock int
			fmt.Print("ID Barang: ")
			fmt.Scan(&id)
			fmt.Print("Nama Barang: ")
			fmt.Scan(&name)
			fmt.Print("Harga Barang: ")
			fmt.Scan(&price)
			fmt.Print("Stok Barang: ")
			fmt.Scan(&stock)
			AddItem(&minimart, id, name, price, stock)
		case 2:
			var id int
			var name string
			var price float64
			var stock int
			fmt.Print("ID Barang yang akan diubah: ")
			fmt.Scan(&id)
			fmt.Print("Nama Baru: ")
			fmt.Scan(&name)
			fmt.Print("Harga Baru: ")
			fmt.Scan(&price)
			fmt.Print("Stok Baru: ")
			fmt.Scan(&stock)
			EditItem(&minimart, id, name, price, stock)
		case 4:
			DisplayItems(&minimart)
		case 5:
			var name string
			var qty int
			fmt.Print("Nama Barang: ")
			fmt.Scan(&name)
			fmt.Print("Jumlah: ")
			fmt.Scan(&qty)
			AddTransaction(&minimart, name, qty)
		case 6:
			DisplayTransactions(&minimart)
		case 7:
			InsertionSortByName(&minimart)
			fmt.Println("Barang telah diurutkan berdasarkan nama.")
		case 8:
			SelectionSortByPrice(&minimart)
			fmt.Println("Barang telah diurutkan berdasarkan harga.")
		case 9:
			fmt.Println("Terima kasih telah menggunakan aplikasi kasir minimart!")
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}
