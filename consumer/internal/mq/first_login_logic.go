package consumer

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/consumer/internal/types"
	"github.com/zeromicro/go-zero/core/logc"
)

// FirstLogin 处理会员第一次登录的时候,发放新手优惠券
func FirstLogin(ctx context.Context, body []byte) {
	logc.Infof(ctx, "收到消息: %s", body)
	var msg types.MemberInfo
	err := json.Unmarshal(body, &msg)
	if err != nil {
		logc.Errorf(context.Background(), "序列化 JSON 失败: %v", err)
	}

	logc.Infof(ctx, "消息内容: %s", msg.Nickname)

	return
}
