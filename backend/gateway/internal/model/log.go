package model

type ApiLog struct {
	Level        string  `json:"level"`         // 日誌等級
	StatusCode   int32   `json:"status_code"`   // 請求狀態碼
	RequestAt    int64   `json:"request_at"`    // 請求時間
	ResponseTime float64 `json:"response_time"` // 響應時間
	ClientIp     string  `json:"client_ip"`     // 客戶端IP
	Method       string  `json:"method"`        // 請求方法
	Path         string  `json:"path"`          // 請求路徑
}
