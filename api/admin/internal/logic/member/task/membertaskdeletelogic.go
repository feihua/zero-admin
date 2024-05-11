package logic

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberTaskDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberTaskDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberTaskDeleteLogic {
	return MemberTaskDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberTaskDeleteLogic) MemberTaskDelete(req types.DeleteMemberTaskReq) (*types.DeleteMemberTaskResp, error) {
	_, err := l.svcCtx.MemberTaskService.MemberTaskDelete(l.ctx, &umsclient.MemberTaskDeleteReq{
		Ids: req.Ids,
	})

	if err != nil {
		logc.Errorf(l.ctx, "根据Id: %+v,删除会员任务异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("删除会员任务失败")
	}

	return &types.DeleteMemberTaskResp{
		Code:    "000000",
		Message: "",
	}, nil
}
