package memberloginlogservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

// MemberLoginLogListLogic 会员登录日志
/*
Author: LiuFeiHua
Date: 2024/5/21 17:49
*/
type MemberLoginLogListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberLoginLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberLoginLogListLogic {
	return &MemberLoginLogListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// MemberLoginLogList 分页查询会员登录日志
func (l *MemberLoginLogListLogic) MemberLoginLogList(in *umsclient.MemberLoginLogListReq) (*umsclient.MemberLoginLogListResp, error) {
	loginLog := query.UmsMemberLoginLog
	q := loginLog.WithContext(l.ctx).Where(loginLog.MemberID.Eq(in.MemberId))
	offset := (in.Current - 1) * in.PageSize
	result, err := q.Offset(int(offset)).Limit(int(in.PageSize)).Find()
	count, err := q.Count()

	if err != nil {
		logc.Errorf(l.ctx, "查询会员登录记录列表信息失败,参数:%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*umsclient.MemberLoginLogListData
	for _, item := range result {
		list = append(list, &umsclient.MemberLoginLogListData{
			Id:         item.ID,
			MemberId:   item.MemberID,
			CreateTime: item.CreateTime.Format("2006-01-02 15:04:05"),
			MemberIP:   item.MemberIP,
			City:       item.City,
			LoginType:  item.LoginType,
			Province:   item.Province,
		})
	}

	logc.Infof(l.ctx, "查询会员登录记录列表信息,参数：%+v,响应：%+v", in, list)
	return &umsclient.MemberLoginLogListResp{
		Total: count,
		List:  list,
	}, nil

}
