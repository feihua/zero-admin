package tag

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteMemberTagLogic 删除会员标签
/*
Author: LiuFeiHua
Date: 2024/5/23 9:20
*/
type DeleteMemberTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteMemberTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMemberTagLogic {
	return &DeleteMemberTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteMemberTag 删除会员标签
func (l *DeleteMemberTagLogic) DeleteMemberTag(req *types.DeleteMemberTagReq) (resp *types.DeleteMemberTagResp, err error) {
	_, err = l.svcCtx.MemberTagService.DeleteMemberTag(l.ctx, &umsclient.DeleteMemberTagReq{
		Ids: req.Ids,
	})

	if err != nil {
		logc.Errorf(l.ctx, "根据Id: %+v,删除会员标签异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("删除会员标签失败")
	}
	return &types.DeleteMemberTagResp{
		Code:    "000000",
		Message: "删除会员标签成功",
	}, nil
}
