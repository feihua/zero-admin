package level

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteMemberLevelLogic 会员等级
/*
Author: LiuFeiHua
Date: 2024/5/13 13:38
*/
type DeleteMemberLevelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteMemberLevelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMemberLevelLogic {
	return &DeleteMemberLevelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteMemberLevel 删除会员等级
func (l *DeleteMemberLevelLogic) DeleteMemberLevel(req *types.DeleteMemberLevelReq) (resp *types.DeleteMemberLevelResp, err error) {
	_, err = l.svcCtx.MemberLevelService.MemberLevelDelete(l.ctx, &umsclient.MemberLevelDeleteReq{
		Ids: req.Ids,
	})

	if err != nil {
		logc.Errorf(l.ctx, "根据Id: %+v,删除会员等级异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("删除会员等级失败")
	}
	return &types.DeleteMemberLevelResp{
		Code:    "000000",
		Message: "添加会员等级成功",
	}, nil
}
