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

// Interface chung
type DocumentViewer interface {
	Display()
}

// Proxy kiểm soát truy cập
type DocumentProxy struct {
	realDocument *Document
	userRole     string
}

func (p *DocumentProxy) Display() {
	if p.userRole != "admin" {
		fmt.Println("Truy cập bị từ chối: bạn không có quyền xem tài liệu này.")
		return
	}
	p.realDocument.Display()
}

func main() {
	secretDoc := &Document{
		Title:   "Kế hoạch tuyệt mật",
		Content: "Chi tiết chiến lược...",
	}

	// Người dùng thường
	userProxy := &DocumentProxy{
		realDocument: secretDoc,
		userRole:     "guest",
	}
	userProxy.Display()

	// Quản trị viên
	adminProxy := &DocumentProxy{
		realDocument: secretDoc,
		userRole:     "admin",
	}
	adminProxy.Display()
}
