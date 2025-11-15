package greetings

// fmt là một package tiêu chuẩn trong Go cung cấp các hàm để định dạng đầu vào và đầu ra.
import "fmt"

// Hello return a greeting for the named person.
// Hello: Function name
// the first string: Parameter type
// the second string: Return type
func Hello(name string) string {
	// Sử dụng Sprintf để định dạng chuỗi với tên được cung cấp.
	// %v là một verb trong fmt để in giá trị ở dạng mặc định.
	// Hàm Sprintf trả về một chuỗi đã được định dạng.
	// Trong Go ":=" operator là một cách ngắn gọn để khai báo và khởi tạo biến chỉ trong 1 dòng.
	// Nó tự động suy luận kiểu dữ liệu từ giá trị bên phải.
	// Ở đây, biến message được khai báo và khởi tạo với chuỗi đã định dạng.
	// var message string
	// message = fmt.Sprintf("Hi, %v. Welcome!", name)
	// Hoặc có thể viết gọn lại như sau:
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	// Trả về chuỗi đã định dạng.
	return message
}
