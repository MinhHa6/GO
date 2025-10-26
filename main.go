package main

import (
	"bufio"
	"fmt"
	"os"
)
var (
	name ="Go lang "
	age=21
)

func main() {
	// In ra "helo" và xuống dòng (lưu ý: \nhelo sẽ in một dòng trống rồi in helo)
	fmt.Print("\nhelo\n")

	// 1. Khai báo biến bằng var (khai báo tường minh)
	var a int = 10

	// 2. Khai báo biến bằng toán tử ngắn gọn :=
	b := 20

	// In b, sau đó in a
	// Lưu ý: fmt.Print(b) và fmt.Print(a) sẽ in liền nhau: 2010
	fmt.Print(b)
	fmt.Print(a)
    
    // In giá trị của biến name và xuống dòng
	fmt.Println(name)

    // In giá trị của biến age bằng fmt.Printf, cần một format string (ví dụ: %d cho số nguyên)
    // Và xuống dòng để kết quả đẹp hơn
	fmt.Printf("Tuoi: %d\n", age)
	// 1. In ra lời nhắc người dùng
	fmt.Println("Tên tôi là :")

	// 2. Khai báo biến để lưu tên (dù không dùng trực tiếp trong bufio.Scanner nhưng vẫn nên có)
	// var hoten string 

	// 3. Khởi tạo Scanner để đọc từ Standard Input (bàn phím)
	scan := bufio.NewScanner(os.Stdin)

	// 4. THỰC HIỆN LỆNH ĐỌC DỮ LIỆU
	// scan.Scan() đọc một token (mặc định là một dòng) cho đến khi gặp ký tự xuống dòng (Enter)
	scan.Scan()

	// 5. LẤY KẾT QUẢ ĐÃ ĐỌC
	// scan.Text() trả về chuỗi dữ liệu đã được đọc
	hoten := scan.Text()

	// 6. In kết quả để kiểm tra
	fmt.Printf("Xin chào, bạn tên là: %s\n", hoten)
}
