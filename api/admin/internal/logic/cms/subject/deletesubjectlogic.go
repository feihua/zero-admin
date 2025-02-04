package subject

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteSubjectLogic 删除专题表
/*
Author: 刘飞华
Date: 2025/02/04 15:04:17
*/
type DeleteSubjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteSubjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSubjectLogic {
	return &DeleteSubjectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteSubject 删除专题表
func (l *DeleteSubjectLogic) DeleteSubject(req *types.DeleteSubjectReq) (resp *types.DeleteSubjectResp, err error) {
	_, err = l.svcCtx.SubjectService.DeleteSubject(l.ctx, &cmsclient.DeleteSubjectReq{
		Ids: req.Ids,
	})

	if err != nil {
		logc.Errorf(l.ctx, "删除专题表失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.DeleteSubjectResp{
		Code:    "000000",
		Message: "删除专题表成功",
	}, nil
}
