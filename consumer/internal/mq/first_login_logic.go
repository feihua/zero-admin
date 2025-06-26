package consumer

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/consumer/internal/types"
	"github.com/feihua/zero-admin/rpc/sms/client/couponrecordservice"
	"github.com/feihua/zero-admin/rpc/sms/client/couponservice"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/feihua/zero-admin/rpc/ums/client/memberinfoservice"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"
)

// FirstLogin 处理会员第一次登录的时候,发放新手优惠券
func FirstLogin(ctx context.Context, body []byte, memberInfoService memberinfoservice.MemberInfoService, CouponService couponservice.CouponService, CouponRecordService couponrecordservice.CouponRecordService) {
	logc.Infof(ctx, "收到消息: %s", body)
	var memberInfo types.MemberInfo
	err := json.Unmarshal(body, &memberInfo)
	if err != nil {
		logc.Errorf(context.Background(), "序列化 JSON 失败: %v", err)
	}

	logc.Infof(ctx, "消息内容: %+v", memberInfo)

	if memberInfo.FirstLoginStatus != 1 {
		logc.Errorf(context.Background(), "用户: %s,不是首次登录了", memberInfo.Nickname)
		return
	}

	req := &smsclient.QueryCouponByCodeReq{
		Code: "XINREN",
	}
	resp, err := CouponService.QueryCouponByCode(ctx, req)
	if err != nil {
		logc.Errorf(ctx, "查询优惠券列表失败,参数：%s,响应：%s", req, err.Error())
		return
	}

	if len(resp.List) == 0 {
		logc.Errorf(ctx, "暂无可发放的新人优惠券,会员信息：%+v", memberInfo)
		return
	}
	for _, i := range resp.List {
		recordReq := smsclient.AddCouponRecordReq{
			MemberId: memberInfo.MemberId, // 会员ID
			CouponId: i.Id,                // 优惠券ID
			GetType:  0,                   // 获取类型：0->后台赠送；1->主动获取
		}
		_, err = CouponRecordService.AddCouponRecord(ctx, &recordReq)
		if err != nil {
			logc.Errorf(ctx, "给用户: %s,发放优惠券失败：%+v", memberInfo.Nickname, err)
			return
		}
		logc.Infof(ctx, "给用户: %s,发放优惠券成功：%+v", memberInfo.Nickname, i)
	}

	_, err = memberInfoService.UpdateFirstLoginStatus(ctx, &umsclient.UpdateFirstLoginStatusReq{
		MemberId:    memberInfo.MemberId,
		CouponCount: int32(len(resp.List)),
	})

	if err != nil {
		logc.Errorf(ctx, "更新用户: %s,首次登录状态失败：%+v", memberInfo.Nickname, err)
		return
	}

}
