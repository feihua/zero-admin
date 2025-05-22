package membertagservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/pointerprocess"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

// QueryMemberTagDetailLogic 查询用户标签详情
/*
Author: LiuFeiHua
Date: 2025/05/22 10:44:59
*/
type QueryMemberTagDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryMemberTagDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberTagDetailLogic {
	return &QueryMemberTagDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryMemberTagDetail 查询用户标签详情
func (l *QueryMemberTagDetailLogic) QueryMemberTagDetail(in *umsclient.QueryMemberTagDetailReq) (*umsclient.QueryMemberTagDetailResp, error) {
	item, err := query.UmsMemberTag.WithContext(l.ctx).Where(query.UmsMemberTag.ID.Eq(in.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "用户标签不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("用户标签不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询用户标签异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询用户标签异常")
	}

	data := &umsclient.QueryMemberTagDetailResp{
		Id:                item.ID,                                          // 主键ID
		TagName:           item.TagName,                                     // 标签名称
		Description:       item.Description,                                 // 标签描述
		FinishOrderCount:  item.FinishOrderCount,                            // 自动打标签完成订单数量
		FinishOrderAmount: float32(item.FinishOrderAmount),                  // 自动打标签完成订单金额
		Status:            item.Status,                                      // 状态：0-禁用，1-启用
		CreateBy:          item.CreateBy,                                    // 创建人ID
		CreateTime:        item.CreateTime.Format("2006-01-02 15:04:05"),    // 创建时间
		UpdateBy:          pointerprocess.DefaltData(item.UpdateBy).(int64), // 更新人ID
		UpdateTime:        time_util.TimeToString(item.UpdateTime),          // 更新时间
	}

	return data, nil
}
