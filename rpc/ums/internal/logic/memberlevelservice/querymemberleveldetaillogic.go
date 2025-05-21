package memberlevelservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/pointerprocess"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/zeromicro/go-zero/core/logc"
	"gorm.io/gorm"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryMemberLevelDetailLogic 查询会员等级详情
/*
Author: LiuFeiHua
Date: 2025/5/20 15:16
*/
type QueryMemberLevelDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryMemberLevelDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberLevelDetailLogic {
	return &QueryMemberLevelDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryMemberLevelDetail 查询会员等级详情
func (l *QueryMemberLevelDetailLogic) QueryMemberLevelDetail(in *umsclient.QueryMemberLevelDetailReq) (*umsclient.QueryMemberLevelDetailResp, error) {
	item, err := query.UmsMemberLevel.WithContext(l.ctx).Where(query.UmsMemberLevel.ID.Eq(in.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "会员等级不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("会员等级不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询会员等级异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询会员等级异常")
	}

	data := &umsclient.QueryMemberLevelDetailResp{
		Id:           item.ID,                                          // 主键ID
		Name:         item.Name,                                        // 等级名称
		Level:        item.Level,                                       // 等级
		GrowthPoint:  item.GrowthPoint,                                 // 升级所需成长值
		DiscountRate: float32(item.DiscountRate),                       // 折扣率(0-100)
		FreeFreight:  item.FreeFreight,                                 // 是否免运费
		CommentExtra: item.CommentExtra,                                // 是否可评论获取奖励
		Privileges:   item.Privileges,                                  // 会员特权JSON
		Remark:       item.Remark,                                      // 备注
		IsEnabled:    item.IsEnabled,                                   // 是否启用
		CreateBy:     item.CreateBy,                                    // 创建人ID
		CreateTime:   time_util.TimeToStr(item.CreateTime),             // 创建时间
		UpdateBy:     pointerprocess.DefaltData(item.UpdateBy).(int64), // 更新人ID
		UpdateTime:   time_util.TimeToString(item.UpdateTime),          // 更新时间
		IsDeleted:    item.IsDeleted,                                   // 是否删除
	}

	return data, nil
}
