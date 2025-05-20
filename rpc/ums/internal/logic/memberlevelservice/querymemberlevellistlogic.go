package memberlevelservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/pointerprocess"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryMemberLevelListLogic 查询会员等级列表
/*
Author: LiuFeiHua
Date: 2024/6/11 14:17
*/
type QueryMemberLevelListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryMemberLevelListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberLevelListLogic {
	return &QueryMemberLevelListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryMemberLevelList 查询会员等级列表
func (l *QueryMemberLevelListLogic) QueryMemberLevelList(in *umsclient.QueryMemberLevelListReq) (*umsclient.QueryMemberLevelListResp, error) {
	memberLevel := query.UmsMemberLevel
	q := memberLevel.WithContext(l.ctx)
	if len(in.Name) > 0 {
		q = q.Where(memberLevel.Name.Like("%" + in.Name + "%"))
	}
	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询会员等级列表失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询会员等级列表失败")
	}

	var list []*umsclient.MemberLevelListData

	for _, item := range result {
		list = append(list, &umsclient.MemberLevelListData{
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

		})
	}

	return &umsclient.QueryMemberLevelListResp{
		Total: count,
		List:  list,
	}, nil

}
