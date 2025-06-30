package consumer

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logc"
)

// OrderCancelByUser 用户自己取消订单通知
func OrderCancelByUser(ctx context.Context, body []byte) {
	logc.Infof(ctx, "用户自己取消订单通知mq消息: %s", body)
	var orderInfo map[string]int64
	err := json.Unmarshal(body, &orderInfo)
	if err != nil {
		logc.Errorf(ctx, "序列化 JSON 失败: %v", err)
		return
	}

}
