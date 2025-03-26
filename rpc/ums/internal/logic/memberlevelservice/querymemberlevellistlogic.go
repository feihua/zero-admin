package memberlevelservicelogic

import (
	"context"
	"errors"
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
	q := query.UmsMemberLevel.WithContext(l.ctx)
	if len(in.LevelName) > 0 {
		q = q.Where(query.UmsMemberLevel.LevelName.Like("%" + in.LevelName + "%"))
	}

	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询会员等级列表失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询会员等级列表失败")
	}
	var list []*umsclient.MemberLevelListData
	for _, item := range result {

		list = append(list, &umsclient.MemberLevelListData{
			Id:                 item.ID,                 //
			LevelName:          item.LevelName,          // 等级名称
			GrowthPoint:        item.GrowthPoint,        // 成长点
			DefaultStatus:      item.DefaultStatus,      // 是否为默认等级：0->不是；1->是
			FreeFreightPoint:   item.FreeFreightPoint,   // 免运费标准
			CommentGrowthPoint: item.CommentGrowthPoint, // 每次评价获取的成长值
			IsFreeFreight:      item.IsFreeFreight,      // 是否有免邮特权
			IsSignIn:           item.IsSignIn,           // 是否有签到特权
			IsComment:          item.IsComment,          // 是否有评论获奖励特权
			IsPromotion:        item.IsPromotion,        // 是否有专享活动特权
			IsMemberPrice:      item.IsMemberPrice,      // 是否有会员价格特权
			IsBirthday:         item.IsBirthday,         // 是否有生日特权
			Remark:             item.Remark,             // 备注
		})
	}

	return &umsclient.QueryMemberLevelListResp{
		Total: count,
		List:  list,
	}, nil

}
