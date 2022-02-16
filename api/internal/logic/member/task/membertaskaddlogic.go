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

type MemberTaskAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberTaskAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberTaskAddLogic {
	return MemberTaskAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberTaskAddLogic) MemberTaskAdd(req types.AddMemberTaskReq) (*types.AddMemberTaskResp, error) {
	_, err := l.svcCtx.Ums.MemberTaskAdd(l.ctx, &umsclient.MemberTaskAddReq{
		Name:         req.Name,
		Growth:       req.Growth,
		Intergration: req.Intergration,
		Type:         req.Type,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("添加会员任务信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("添加会员任务失败")
	}

	return &types.AddMemberTaskResp{
		Code:    "000000",
		Message: "添加会员任务成功",
	}, nil
}
