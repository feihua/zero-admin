package operatelogservicelogic

import (
	"context"
	"errors"

	"github.com/bytedance/sonic"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/sys/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryOperateLogListLogic 查询操作日志列表
/*
Author: LiuFeiHua
Date: 2023/12/18 17:09
*/
type QueryOperateLogListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryOperateLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryOperateLogListLogic {
	return &QueryOperateLogListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryOperateLogList 查询操作日志列表
func (l *QueryOperateLogListLogic) QueryOperateLogList(in *sysclient.QueryOperateLogListReq) (*sysclient.QueryOperateLogListResp, error) {
	operateLog := query.SysOperateLog
	q := operateLog.WithContext(l.ctx)

	if len(in.OperateName) > 0 {
		q = q.Where(operateLog.OperateName.Like("%" + in.OperateName + "%"))
	}

	if len(in.OperateUrl) > 0 {
		q = q.Where(operateLog.OperateURL.Like("%" + in.OperateUrl + "%"))
	}
	if len(in.OperateIp) > 0 {
		q = q.Where(operateLog.OperateIP.Like("%" + in.OperateIp + "%"))
	}

	if in.Status != 2 {
		q = q.Where(operateLog.Status.Eq(in.Status))
	}

	result, count, err := q.Order(operateLog.ID.Desc()).FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询操作日志列表失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询操作日志列表失败")
	}
	var list = make([]*sysclient.OperateLogListData, 0, len(result))

	for _, item := range result {
		data := &sysclient.OperateLogListData{
			Id:           item.ID,                               // 操作日志id
			Title:        item.Title,                            // 模块标题
			OperateName:  item.OperateName,                      // 操作人员
			OperateUrl:   item.OperateURL,                       // 请求URL
			OperateIp:    item.OperateIP,                        // 主机地址
			OperateParam: item.OperateParam,                     // 请求参数
			JsonResult:   item.JSONResult,                       // 返回参数
			Status:       item.Status,                           // 操作状态(0:异常,正常)
			OperateTime:  time_util.TimeToStr(item.OperateTime), // 操作时间
			CostTime:     item.CostTime,                         // 消耗时间
		}

		_ = sonic.Unmarshal(item.Extra, data.Extra)
		list = append(list, data)
	}

	return &sysclient.QueryOperateLogListResp{
		Total: count,
		List:  list,
	}, nil

}
