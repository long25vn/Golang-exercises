package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// wg được sử dụng để đợi các goroutine kết thúc
var wg sync.WaitGroup

// Hàm này chạy trước khi các lệnh hàm main được thực thi
func init() {
	rand.Seed(time.Now().UnixNano()) // Tạo nhân gieo giá trị ngẫu nhiên
}

func main() {
	court := make(chan int)

	wg.Add(2)

	// Tạo 2 người chơi
	go player("Federer", court)
	go player("Djokovic", court)

	// Bắt đầu phát bóng cho một ván đấu
	court <- 1

	wg.Wait() // Đợi ván đấu kết thúc
	fmt.Println("Ván đấu kết thúc!")
}

// Hàm mô tả một người chơi quần vợt
func player(name string, court chan int) {
	// Thông báo người này chơi xong ván đấu
	defer wg.Done()

	for {
		// Đợi banh từ đối thủ
		ball, ok := <-court
		if !ok {
			// Thắng nếu kênh đã đóng
			fmt.Printf("%s thắng!\n", name)
			return
		}

		// Lấy một giá trị ngẫu nhiên 0 - 99
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("%s đánh hỏng ở lượt đánh thứ %d!\n", name, ball)

			// Đóng kênh khi đánh hỏng
			close(court)
			return
		}

		fmt.Printf("Lượt đánh bóng thành công thứ %d: %s\n", ball, name)
		ball++
		court <- ball
	}
}
