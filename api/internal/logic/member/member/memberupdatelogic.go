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

type MemberUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberUpdateLogic {
	return MemberUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberUpdateLogic) MemberUpdate(req types.UpdateMemberReq) (*types.UpdateMemberResp, error) {
	_, err := l.svcCtx.Ums.MemberUpdate(l.ctx, &umsclient.MemberUpdateReq{
		Id:                    req.Id,
		MemberLevelId:         req.MemberLevelId,
		Username:              req.Username,
		Password:              req.Password,
		Nickname:              req.Nickname,
		Phone:                 req.Phone,
		Status:                req.Status,
		CreateTime:            req.CreateTime,
		Icon:                  req.Icon,
		Gender:                req.Gender,
		Birthday:              req.Birthday,
		City:                  req.City,
		Job:                   req.Job,
		PersonalizedSignature: req.PersonalizedSignature,
		SourceType:            req.SourceType,
		Integration:           req.Integration,
		Growth:                req.Growth,
		LuckeyCount:           req.LuckeyCount,
		HistoryIntegration:    req.HistoryIntegration,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("更新会员信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("更新会员信息失败")
	}

	return &types.UpdateMemberResp{
		Code:    "000000",
		Message: "更新会员信息成功",
	}, nil
}
