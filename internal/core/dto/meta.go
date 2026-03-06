package dto

import "encoding/json"

type Meta struct {
	status     int
	message    string
	total      int64
	page       int
	pageSize   int
	totalPages int
}

// Constructor: Cách duy nhất để tạo ra Meta từ bên ngoài package
func NewMeta(status int, message string) *Meta {
	return &Meta{
		status:  status,
		message: message,
	}
}

// Getters: Đảm bảo tính đóng gói, chỉ đọc, không ghi đè lung tung
func (m *Meta) Status() int     { return m.status }
func (m *Meta) Message() string { return m.message }

// Setters (nếu cần logic kiểm tra)
func (m *Meta) SetPagination(total int64, page, size int) {
	m.total = total
	m.page = page
	m.pageSize = size
	if size > 0 {
		m.totalPages = int((total + int64(size) - 1) / int64(size))
	}
}

// MarshalJSON: Bắt buộc phải có để thư viện JSON thấy được các field private
func (m Meta) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Status     int    `json:"status"`
		Message    string `json:"message"`
		Total      int64  `json:"total,omitempty"`
		Page       int    `json:"page,omitempty"`
		PageSize   int    `json:"page_size,omitempty"`
		TotalPages int    `json:"total_pages,omitempty"`
	}{
		Status:     m.status,
		Message:    m.message,
		Total:      m.total,
		Page:       m.page,
		PageSize:   m.pageSize,
		TotalPages: m.totalPages,
	})
}
