package order

import (
	"context"
	"github.com/bytedance/sonic"
	"github.com/zeromicro/go-zero/core/logc"
)

// OrderConfirm 用户确认订单通知
func OrderConfirm(ctx context.Context, body []byte) {
	logc.Infof(ctx, "用户确认订单通知mq消息: %s", body)
	var orderInfo map[string]int64
	err := sonic.Unmarshal(body, &orderInfo)
	if err != nil {
		logc.Errorf(ctx, "序列化 JSON 失败: %v", err)
		return
	}

}
