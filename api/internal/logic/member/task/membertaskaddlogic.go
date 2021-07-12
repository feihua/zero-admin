package logic

import (
	"context"
	"encoding/json"
	"go-zero-admin/api/internal/common/errorx"
	"go-zero-admin/rpc/ums/umsclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
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
		logx.Errorf("添加会员任务参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("添加会员任务失败")
	}

	return &types.AddMemberTaskResp{
		Code:    "000000",
		Message: "添加会员任务成功",
	}, nil
}
