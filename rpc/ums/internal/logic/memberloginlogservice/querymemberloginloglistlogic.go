package memberloginlogservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryMemberLoginLogListLogic 查询会员登录记录列表
/*
Author: LiuFeiHua
Date: 2024/6/11 14:15
*/
type QueryMemberLoginLogListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryMemberLoginLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberLoginLogListLogic {
	return &QueryMemberLoginLogListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryMemberLoginLogList 查询会员登录记录列表
func (l *QueryMemberLoginLogListLogic) QueryMemberLoginLogList(in *umsclient.QueryMemberLoginLogListReq) (*umsclient.QueryMemberLoginLogListResp, error) {
	loginLog := query.UmsMemberLoginLog
	q := loginLog.WithContext(l.ctx).Where(loginLog.MemberID.Eq(in.MemberId))

	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询会员登录记录列表失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询会员登录记录列表失败")
	}

	var list []*umsclient.MemberLoginLogListData
	for _, item := range result {
		list = append(list, &umsclient.MemberLoginLogListData{
			Id:         item.ID,                              //
			MemberId:   item.MemberID,                        // 会员id
			MemberIp:   item.MemberIP,                        // 登录ip
			City:       item.City,                            // 登录城市
			LoginType:  item.LoginType,                       // 登录类型：0->PC；1->android;2->ios;3->小程序
			Province:   item.Province,                        // 登录省份
			CreateTime: time_util.TimeToStr(item.CreateTime), // 登录时间
		})
	}

	return &umsclient.QueryMemberLoginLogListResp{
		Total: count,
		List:  list,
	}, nil

}
