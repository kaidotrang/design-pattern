package main

import "fmt"

// Đối tượng thật
type Document struct {
	Title   string
	Content string
}

func (d *Document) Display() {
	fmt.Printf("Tài liệu: %s\nNội dung: %s\n", d.Title, d.Content)
}

func main() {
	doc := &Document{
		Title:   "Kế hoạch tuyệt mật",
		Content: "Chi tiết chiến lược...",
	}

	// Người dùng truy cập trực tiếp
	doc.Display()
}
