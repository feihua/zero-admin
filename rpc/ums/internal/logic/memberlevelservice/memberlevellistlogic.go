package memberlevelservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberLevelListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberLevelListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberLevelListLogic {
	return &MemberLevelListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberLevelListLogic) MemberLevelList(in *umsclient.MemberLevelListReq) (*umsclient.MemberLevelListResp, error) {
	q := query.UmsMemberLevel.WithContext(l.ctx)
	if len(in.Name) > 0 {
		q = q.Where(query.UmsMemberLevel.LevelName.Like("%" + in.Name + "%"))
	}
	offset := (in.Current - 1) * in.PageSize
	result, err := q.Offset(int(offset)).Limit(int(in.PageSize)).Find()
	count, err := q.Count()

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
			FreeFreightPoint:   float32(level.FreeFreightPoint),
			CommentGrowthPoint: level.CommentGrowthPoint,
			IsFreeFreight:      level.IsFreeFreight,
			IsSignIn:           level.IsSignIn,
			IsComment:          level.IsComment,
			IsPromotion:        level.IsPromotion,
			IsMemberPrice:      level.IsMemberPrice,
			IsBirthday:         level.IsBirthday,
			Remark:             *level.Remark,
		})
	}

	logc.Infof(l.ctx, "查询会员等级列表信息,参数：%+v,响应：%+v", in, list)
	return &umsclient.MemberLevelListResp{
		Total: count,
		List:  list,
	}, nil

}
