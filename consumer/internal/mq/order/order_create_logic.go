package order

import (
	"context"
	"github.com/bytedance/sonic"
	"github.com/zeromicro/go-zero/core/logc"
)

// OrderCreate 创建订单通知
func OrderCreate(ctx context.Context, body []byte) {
	logc.Infof(ctx, "创建订单通知mq消息: %s", body)
	var orderInfo map[string]int64
	err := sonic.Unmarshal(body, &orderInfo)
	if err != nil {
		logc.Errorf(ctx, "序列化 JSON 失败: %v", err)
		return
	}

}
