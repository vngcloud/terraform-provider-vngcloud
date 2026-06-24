package vdbv2

import (
	"encoding/json"
	"encoding/xml"
	"strings"
)

// Kafka APIs trả Content-Type "text/plain;charset=UTF-8" cho body JSON, nên
// APIClient.decode mặc định trả "undefined response type". Helper này dùng
// riêng cho các file api_kafka_*.go: bỏ qua Content-Type, parse JSON trực tiếp.
func decodeKafkaResponse(v interface{}, body []byte, contentType string) error {
	if len(body) == 0 {
		return nil
	}
	if strings.Contains(contentType, "application/xml") {
		return xml.Unmarshal(body, v)
	}
	return json.Unmarshal(body, v)
}
