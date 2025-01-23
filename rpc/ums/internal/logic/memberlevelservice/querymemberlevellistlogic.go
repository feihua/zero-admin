package memberlevelservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryMemberLevelListLogic 查询会员等级表列表
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

// QueryMemberLevelList 查询会员等级表列表
func (l *QueryMemberLevelListLogic) QueryMemberLevelList(in *umsclient.QueryMemberLevelListReq) (*umsclient.QueryMemberLevelListResp, error) {
	q := query.UmsMemberLevel.WithContext(l.ctx)
	if len(in.LevelName) > 0 {
		q = q.Where(query.UmsMemberLevel.LevelName.Like("%" + in.LevelName + "%"))
	}

	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询会员等级列表信息失败,参数:%+v,异常:%s", in, err.Error())
		return nil, err
	}
	var list []*umsclient.MemberLevelListData
	for _, level := range result {

		list = append(list, &umsclient.MemberLevelListData{
			Id:                 level.ID,
			LevelName:          level.LevelName,
			GrowthPoint:        level.GrowthPoint,
			DefaultStatus:      level.DefaultStatus,
			FreeFreightPoint:   level.FreeFreightPoint,
			CommentGrowthPoint: level.CommentGrowthPoint,
			IsFreeFreight:      level.IsFreeFreight,
			IsSignIn:           level.IsSignIn,
			IsComment:          level.IsComment,
			IsPromotion:        level.IsPromotion,
			IsMemberPrice:      level.IsMemberPrice,
			IsBirthday:         level.IsBirthday,
			Remark:             level.Remark,
		})
	}

	return &umsclient.QueryMemberLevelListResp{
		Total: count,
		List:  list,
	}, nil

}
