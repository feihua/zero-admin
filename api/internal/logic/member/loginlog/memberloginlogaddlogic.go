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

type MemberLoginLogAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberLoginLogAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberLoginLogAddLogic {
	return MemberLoginLogAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberLoginLogAddLogic) MemberLoginLogAdd(req types.AddMemberLoginLogReq) (*types.AddMemberLoginLogResp, error) {
	_, err := l.svcCtx.Ums.MemberLoginLogAdd(l.ctx, &umsclient.MemberLoginLogAddReq{
		MemberId:   req.MemberId,
		CreateTime: req.CreateTime,
		Ip:         req.Ip,
		City:       req.City,
		LoginType:  req.LoginType,
		Province:   req.Province,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("添加会员登录记录信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("添加会员登录记录失败")
	}

	return &types.AddMemberLoginLogResp{
		Code:    "000000",
		Message: "添加员登录记录成功",
	}, nil
}
