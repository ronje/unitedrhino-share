package application

import (
	"gitee.com/i-Things/share/def"
	"gitee.com/i-Things/share/devices"
)

// 连接和断连消息信息
type ConnectMsg struct {
	Device    devices.Core   `json:"device"`
	Status    def.ConnStatus `json:"status"`
	Timestamp int64          `json:"timestamp,string"` //毫秒时间戳
}

func (c ConnectMsg) GenSerial() string {
	return ""
}
