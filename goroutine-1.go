// Chương trình in ra 3 lần bảng chữ cái viết thường và viết hoa
package main

import (
	"fmt"
	//"time"
)

func main() {
	fmt.Println("Bắt đầu Goroutines")

	// Khai báo hàm vô danh và tạo một goroutine với từ khóa go
	go func() {
		// Hiển thị bảng chữ cái viết thường 3 lần
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
				//time.Sleep(time.Millisecond * 10)
			}
		}
	}()

	go func() {
		// Hiển thị bảng chữ cái viết hoa 3 lần
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
				//time.Sleep(time.Millisecond * 10)
			}
		}
	}()

	// Đợi các goroutine kết thúc
	fmt.Println("Đợi kết thúc các gorountine")
	var input string
	fmt.Scanln(&input)
	fmt.Println("\nKết thúc chương trình")
}
