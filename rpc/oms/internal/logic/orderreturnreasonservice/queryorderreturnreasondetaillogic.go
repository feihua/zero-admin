package orderreturnreasonservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/pointerprocess"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

// QueryOrderReturnReasonDetailLogic 查询退货原因详情
/*
Author: LiuFeiHua
Date: 2025/05/26 15:21:44
*/
type QueryOrderReturnReasonDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryOrderReturnReasonDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryOrderReturnReasonDetailLogic {
	return &QueryOrderReturnReasonDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryOrderReturnReasonDetail 查询退货原因详情
func (l *QueryOrderReturnReasonDetailLogic) QueryOrderReturnReasonDetail(in *omsclient.QueryOrderReturnReasonDetailReq) (*omsclient.QueryOrderReturnReasonDetailResp, error) {
	item, err := query.OmsOrderReturnReason.WithContext(l.ctx).Where(query.OmsOrderReturnReason.ID.Eq(in.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "退货原因不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("退货原因不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询退货原因异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询退货原因异常")
	}

	data := &omsclient.QueryOrderReturnReasonDetailResp{
		Id:         item.ID,                                          // 主键ID
		Name:       item.Name,                                        // 退货类型
		Sort:       item.Sort,                                        // 排序
		Status:     item.Status,                                      // 状态：0->不启用；1->启用
		CreateBy:   item.CreateBy,                                    // 创建者
		CreateTime: time_util.TimeToStr(item.CreateTime),             // 创建时间
		UpdateBy:   pointerprocess.DefaltData(item.UpdateBy).(int64), // 更新人ID
		UpdateTime: time_util.TimeToString(item.UpdateTime),          // 更新时间
	}

	return data, nil
}
