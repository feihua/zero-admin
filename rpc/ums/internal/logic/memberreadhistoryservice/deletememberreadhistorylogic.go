package memberreadhistoryservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteMemberReadHistoryLogic 清空浏览记录/删除浏览记录
/*
Author: LiuFeiHua
Date: 2024/6/11 14:07
*/
type DeleteMemberReadHistoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteMemberReadHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMemberReadHistoryLogic {
	return &DeleteMemberReadHistoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteMemberReadHistory 清空浏览记录/删除浏览记录
func (l *DeleteMemberReadHistoryLogic) DeleteMemberReadHistory(in *umsclient.DeleteMemberReadHistoryReq) (*umsclient.DeleteMemberReadHistoryResp, error) {
	q := query.UmsMemberReadHistory
	historyDo := q.WithContext(l.ctx).Where(q.MemberID.Eq(in.MemberId))
	if len(in.Ids) > 0 {
		historyDo = historyDo.Where(q.ID.In(in.Ids...))
	}
	_, err := historyDo.Delete()
	if err != nil {
		logc.Errorf(l.ctx, "删除浏览记录失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除浏览记录失败")
	}

	return &umsclient.DeleteMemberReadHistoryResp{}, nil
}
