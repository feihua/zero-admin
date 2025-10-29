package memberconsumesettingservicelogic

import (
	"context"
	"errors"

	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteMemberConsumeSettingLogic 删除积分消费设置
/*
Author: LiuFeiHua
Date: 2025/05/23 11:32:02
*/
type DeleteMemberConsumeSettingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteMemberConsumeSettingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMemberConsumeSettingLogic {
	return &DeleteMemberConsumeSettingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteMemberConsumeSetting 删除积分消费设置
func (l *DeleteMemberConsumeSettingLogic) DeleteMemberConsumeSetting(in *umsclient.DeleteMemberConsumeSettingReq) (*umsclient.DeleteMemberConsumeSettingResp, error) {
	q := query.UmsMemberConsumeSetting

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		logc.Errorf(l.ctx, "删除积分消费设置失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除积分消费设置失败")
	}

	return &umsclient.DeleteMemberConsumeSettingResp{}, nil
}
