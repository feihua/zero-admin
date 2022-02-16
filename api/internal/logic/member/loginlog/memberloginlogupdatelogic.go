package logic

import (
	"context"
	"encoding/json"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/rpc/ums/umsclient"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberLoginLogUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberLoginLogUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberLoginLogUpdateLogic {
	return MemberLoginLogUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberLoginLogUpdateLogic) MemberLoginLogUpdate(req types.UpdateMemberLoginLogReq) (*types.UpdateMemberLoginLogResp, error) {
	_, err := l.svcCtx.Ums.MemberLoginLogUpdate(l.ctx, &umsclient.MemberLoginLogUpdateReq{
		Id:         req.Id,
		MemberId:   req.MemberId,
		CreateTime: req.CreateTime,
		Ip:         req.Ip,
		City:       req.City,
		LoginType:  req.LoginType,
		Province:   req.Province,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("更新会员登录记录信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("更新员登录记录失败")
	}

	return &types.UpdateMemberLoginLogResp{
		Code:    "000000",
		Message: "更新员登录记录成功",
	}, nil
}
