package memberloginlogservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

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

func (l *MemberLoginLogListLogic) MemberLoginLogList(in *umsclient.MemberLoginLogListReq) (*umsclient.MemberLoginLogListResp, error) {
	q := query.UmsMemberLoginLog.WithContext(l.ctx)
	if in.MemberId != 0 {
		q = q.Where(query.UmsMemberLoginLog.MemberID.Eq(in.MemberId))
	}
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
			Ip:         item.IP,
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
