/* Program Aplikasi Quiz 'Who Wants To Be A Millionaire'
Deskripsi: Aplikasi ini digunakan untuk mengolah data bank soal dan juga pelaksanaan kuis, yang mana peserta yang akan menjawab soal pilihan ganda yang akan dipilih sebanyak N secara acak dari bank soal. Pengguna aplikasi ini adalah admin aplikasi dan peserta kuis.
Spesifikasi :
a.Admin bisa menambah, mengubah (edit), dan menghapus data soal beserta kunci jawabannya.
b.Calon peserta bisa mendaftar sebagai peserta melalui aplikasi.
c.Peserta bisa mengikuti kuis dengan menjawab pertanyaan yang dipilih secara acak dari bank soal. Solusi akan selalu ditampilkan setiap kali peserta menjawab pertanyaan. Selain itu juga menampilkan skor akhir apabila telah menjawab semua pertanyaan.
d.Peserta boleh mengikuti kuis sebanyak 3 kali untuk memperbaharui skor.
e.Aplikasi bisa menampilkan daftar peserta terurut berdasarkan skor yang diperoleh.
f.Admin bisa melihat 10 daftar soal yang paling banyak benar atau salah dijawab oleh peserta.*/

package main
//Mengimpor paket-paket yang diperlukan untuk menjalankan program seperti input/output, format teks, random number generator, dan manipulasi string.
import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

//Mendefinisikan dua tipe struct: Question untuk menyimpan pertanyaan dan jawabannya, Participant untuk menyimpan data peserta seperti nama, skor, dan jumlah usaha.
type Question struct {
	question string
	answers  [4]string
	correct  string
}

type Participant struct {
	name     string
	score    int
	attempts int
}

//Mendefinisikan variabel global untuk menyimpan daftar pertanyaan, daftar peserta, serta penghitung jumlah pertanyaan dan peserta.
var questions [23]Question
var participants [100]Participant
var participantCount int
var questionCount int

//Fungsi init menginisialisasi array questions dengan 23 pertanyaan awal dan menetapkan questionCount menjadi 23
func init() {
	questions = [23]Question{
		{"Apa nama ibukota Indonesia?", [4]string{"A. London", "B. Jakarta", "C. Rome", "D. Madrid"}, "B"},
		{"Apa nama planet yang memiliki julukan 'Planet Merah'?", [4]string{"A. Mars", "B. Venus", "C. Jupiter", "D. Saturnus"}, "A"},
		{"Siapa penulis buku 'Sapiens'?", [4]string{"A. J.K. Rowling", "B. Harper Lee", "C. Yuval Noah Harari", "D. F. Scott Fitzgerald"}, "C"},
		{"Apa nama mata uang negara Rusia?", [4]string{"A. Poundsterling", "B. Euro", "C. Yen", "D. Rubel"}, "D"},
		{"Siapa Raja di Inggris satu-satunya yang memiliki gelar 'The Great'?", [4]string{"A. Raja Alfred", "B. Raja George V", "C. Raja Æthelwulf", "D. Raja Egbert"}, "A"},
		{"Tahun berapa Perang Dunia I terjadi?", [4]string{"A. 1912", "B. 1913", "C. 1914", "D. 1915"}, "C"},
		{"Siapakah nama orang yang disebut 'Bapak Sosialisme'?", [4]string{"A. Sigmund Freud", "B. Herbert Spencer", "C. August Comte", "D. Karl Marx"}, "D"},
		{"Siapa nama Kaisar Rusia terakhir?", [4]string{"A. Bjorn 'Ironside'", "B. Tsar Nicholas II", "C. Baldwin IV", "D. Erik 'Si Merah'"}, "B"},
		{"Apa nama merek mobil pertama di dunia?", [4]string{"A. Ford", "B. Mercedes Benz", "C. Fiat", "D. Renault"}, "B"},
		{"Apa negara penghasil nikel terbesar di dunia?", [4]string{"A. Indonesia", "B. China", "C. Rusia", "D. Filipina"}, "A"},
		{"Apa nama ibukota Australia?", [4]string{"A. Sydney", "B. Melbourne", "C. Tomsk", "D. Canberra"}, "D"},
		{"Siapakah penemu kereta api?", [4]string{"A. George Stephenson", "B. Alexander Graham Bell", "C. Luis Amstrong", "D. Galileo Galilei"}, "A"},
		{"Siapakah Raja Jerusalem yang memiliki julukan 'Si Lepra'?", [4]string{"A. Richard I", "B. Godfrey", "C. Baldwin IV", "D. Guy De Lusignan"}, "C"},
		{"Tahun berapa Perang Dunia II berakhir?", [4]string{"A. 1950", "B. 1945", "C. 1948", "D. 1955"}, "B"},
		{"Di negara manakah teh pertama kali ditemukan?", [4]string{"A. Jepang", "B. China", "C. Korea", "D. Turki"}, "B"},
		{"Siapakah penemu teh?", [4]string{"A. Zhuge Liang", "B. Cao Cao", "C. Sun Tzu", "D. Shen Nong"}, "D"},
		{"Siapakah penulis buku 'Art Of War'?", [4]string{"A. Sun Tzu", "B. Zhuge Liang", "C. Cao Cao", "D. Shen Nong"}, "A"},
		{"Siapakah tokoh Jerman pada perang dunia II yang gagal menjadi seniman dan menjadi Kanselir Jerman paling kontroversial?", [4]string{"A. Alfred Von Tirpitz", "B. Otto Von Bismarck", "C. Alois Hitler", "D. Adolf Hitler"}, "D"},
		{"Apa negara penghasil kopi terbesar di dunia?", [4]string{"A. Indonesia", "B. Columbia", "C. Turki", "D. Brazil"}, "D"},
		{"Apa nama ibu kota Amerika Serikat?", [4]string{"A. Washington, D.C.", "B. New York", "C. Birmingham", "D. Canberra"}, "A"},
		{"Siapakah Raja Inggris yang pertama kali memimpin Kerajaan Inggris Bersatu setelah berhasil menyatukan kerajaan Anglo-Saxon?", [4]string{"A. Edward 'The Elder'", "B. William I", "C. Æthelstan", "D. Æthelflæd"}, "C"},
		{"Siapakah pendiri Google?", [4]string{"A. Larry Page", "B. Larry Ellison", "C. Paul Allen", "D. Brendan Eich"}, "A"},
	}
	questionCount = 23
}

//Menambahkan pertanyaan baru ke dalam array questions.
func addQuestion() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan pertanyaan: ")
	question, _ := reader.ReadString('\n')
	question = strings.TrimSpace(question)

	var answers [4]string
	for i := 0; i < 4; i++ {
		fmt.Printf("Masukkan pilihan jawaban %d: ", i+1)
		answers[i], _ = reader.ReadString('\n')
		answers[i] = strings.TrimSpace(answers[i])
	}

	fmt.Print("Masukkan jawaban yang benar (huruf besar): ")
	correct, _ := reader.ReadString('\n')
	correct = strings.TrimSpace(correct)
	correct = strings.ToUpper(correct)

	questions[questionCount] = Question{
		question: question,
		answers:  answers,
		correct:  correct,
	}
	questionCount++
}

//Memperbarui pertanyaan yang sudah ada pada indeks tertentu dalam array questions
func updateQuestion(index int) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan pertanyaan baru: ")
	question, _ := reader.ReadString('\n')
	question = strings.TrimSpace(question)

	var answers [4]string
	for i := 0; i < 4; i++ {
		fmt.Printf("Masukkan pilihan jawaban baru %d: ", i+1)
		answers[i], _ = reader.ReadString('\n')
		answers[i] = strings.TrimSpace(answers[i])
	}

	fmt.Print("Masukkan jawaban yang benar baru (huruf besar): ")
	correct, _ := reader.ReadString('\n')
	correct = strings.TrimSpace(correct)
	correct = strings.ToUpper(correct)

	questions[index] = Question{
		question: question,
		answers:  answers,
		correct:  correct,
	}
}

//Menghapus pertanyaan dari array questions pada indeks tertentu
func deleteQuestion(index int) {
	for i := index; i < questionCount-1; i++ {
		questions[i] = questions[i+1]
	}
	questionCount--
}

//Mengacak urutan pertanyaan dalam array questions.
func shuffleQuestions() [23]Question {
	shuffled := questions
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(questionCount, func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})
	return shuffled
}

//Mencari peserta berdasarkan nama dengan metode pencarian sekuensial
func sequentialSearch(name string) int {
	for i := 0; i < participantCount; i++ {
		if participants[i].name == name {
			return i
		}
	}
	return -1
}

//Mencari peserta berdasarkan nama dengan metode pencarian biner
func binarySearch(name string) int {
	low, high := 0, participantCount-1
	for low <= high {
		mid := (low + high) / 2
		if participants[mid].name == name {
			return mid
		} else if participants[mid].name < name {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

//Mengurutkan peserta berdasarkan skor menggunakan algoritma selection sort.
func selectionSort() {
	for i := 0; i < participantCount-1; i++ {
		minIdx := i
		for j := i + 1; j < participantCount; j++ {
			if participants[j].score < participants[minIdx].score {
				minIdx = j
			}
		}
		participants[i], participants[minIdx] = participants[minIdx], participants[i]
	}
}

//Mengurutkan peserta berdasarkan skor menggunakan algoritma insertion sort
func insertionSort() {
	for i := 1; i < participantCount; i++ {
		key := participants[i]
		j := i - 1
		for j >= 0 && participants[j].score > key.score {
			participants[j+1] = participants[j]
			j--
		}
		participants[j+1] = key
	}
}

//Menjalankan kuis untuk peserta tertentu. Mengacak pertanyaan dan meminta jawaban dari peserta.
func runQuiz(participant *Participant) {
	reader := bufio.NewReader(os.Stdin)
	shuffledQuestions := shuffleQuestions()

	questionCount := 10
	if len(questions) < questionCount {
		questionCount = len(questions)
	}

	for i, question := range shuffledQuestions[:questionCount] {
		fmt.Printf("Pertanyaan %d: %s\n", i+1, question.question)
		fmt.Println("Pilihan jawaban:")
		for _, answer := range question.answers {
			fmt.Println(answer)
		}
		fmt.Print("Jawaban Anda: ")
		userAnswer, _ := reader.ReadString('\n')
		userAnswer = strings.TrimSpace(userAnswer)
		userAnswer = strings.ToUpper(userAnswer)

		if userAnswer == question.correct {
			participant.score++
			fmt.Println("Benar!")
		} else {
			fmt.Println("Salah!")
		}
	}

	participant.attempts++
	fmt.Printf("Skor Anda: %d dari %d\n", participant.score, questionCount)
}

//Mendaftarkan peserta baru dan mengembalikan objek Participant.
func registerParticipant() Participant {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan nama Anda: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	return Participant{name: name, score: 0, attempts: 0}
}

//Menampilkan daftar peserta yang diurutkan berdasarkan skor
func displayParticipants() {
	selectionSort()

	fmt.Println("Daftar peserta berdasarkan skor:")
	for i := 0; i < participantCount; i++ {
		participant := participants[i]
		fmt.Printf("%d. %s - Skor: %d, Upaya: %d\n", i+1, participant.name, participant.score, participant.attempts)
	}
}

//Menampilkan 10 soal yang paling banyak benar atau salah dijawab (tidak sepenuhnya diimplementasikan dalam kode ini).
func displayTopQuestions() {
	fmt.Println("10 soal yang paling banyak benar atau salah dijawab:")
	for i := 0; i < 10 && i < questionCount; i++ {
		fmt.Printf("%d. %s\n", i+1, questions[i].question)
	}
}

//Menyediakan menu utama untuk interaksi dengan pengguna. Berisi opsi untuk mengerjakan kuis, melihat hasil, mengedit pertanyaan, dan keluar dari program
func mainMenu() {
	var currentParticipant *Participant
	for {
		fmt.Println("\nMain Menu:")
		fmt.Println("1. Kerjakan Quiz")
		fmt.Println("2. Coba Ulang Quiz (Max 3 Attempt)")
		fmt.Println("3. Lihat Hasil Quiz")
		fmt.Println("4. Edit Pertanyaan(Admin)")
		fmt.Println("5. Keluar")

		var pilihan int
		fmt.Print("Pilih menu: ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			participant := registerParticipant()
			currentParticipant = &participant
			runQuiz(currentParticipant)
			participants[participantCount] = *currentParticipant
			participantCount++
		case 2:
			if currentParticipant != nil && currentParticipant.attempts < 3 {
				runQuiz(currentParticipant)
			} else {
				fmt.Println("Anda telah mencoba sebanyak 3 kali atau belum mendaftar. Silakan pilih menu lain.")
			}
		case 3:
			displayParticipants()
			displayTopQuestions()
		case 4:
			fmt.Println("Edit Pertanyaan:")
			fmt.Println("1. Tambah Pertanyaan")
			fmt.Println("2. Ubah Pertanyaan")
			fmt.Println("3. Hapus Pertanyaan")
			fmt.Println("4. Kembali ke Menu Utama")

			var editPilihan int
			fmt.Print("Pilih menu: ")
			fmt.Scanln(&editPilihan)

			switch editPilihan {
			case 1:
				addQuestion()
			case 2:
				fmt.Print("Masukkan indeks pertanyaan yang ingin diubah: ")
				var index int
				fmt.Scanln(&index)
				if index >= 0 && index < questionCount {
					updateQuestion(index)
				} else {
					fmt.Println("Indeks tidak valid.")
				}
			case 3:
				fmt.Print("Masukkan indeks pertanyaan yang ingin dihapus: ")
				var index int
				fmt.Scanln(&index)
				if index >= 0 && index < questionCount {
					deleteQuestion(index)
				} else {
					fmt.Println("Indeks tidak valid.")
				}
			case 4:
				continue
			default:
				fmt.Println("Pilihan tidak valid")
			}
		case 5:
			fmt.Println("Terima Kasih Telah Mencoba!")
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}

//Fungsi main yang menjalankan fungsi mainMenu untuk memulai program
func main() {
	mainMenu()
}
