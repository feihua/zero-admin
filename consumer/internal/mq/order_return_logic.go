package consumer

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logc"
)

// OrderReturn 申请退款通知
func OrderReturn(ctx context.Context, body []byte) {
	logc.Infof(ctx, "申请退款通知mq消息: %s", body)
	var orderInfo map[string]int64
	err := json.Unmarshal(body, &orderInfo)
	if err != nil {
		logc.Errorf(ctx, "序列化 JSON 失败: %v", err)
		return
	}

}
