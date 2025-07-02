package order

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logc"
)

// OrderDelivery 订单发货通知
func OrderDelivery(ctx context.Context, body []byte) {
	logc.Infof(ctx, "订单发货通知mq消息: %s", body)
	var orderInfo map[string]int64
	err := json.Unmarshal(body, &orderInfo)
	if err != nil {
		logc.Errorf(ctx, "序列化 JSON 失败: %v", err)
		return
	}

}
