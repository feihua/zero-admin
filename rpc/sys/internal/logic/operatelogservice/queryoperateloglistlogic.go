package operatelogservicelogic

import (
	"context"
	"errors"
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
	if len(in.OperationIp) > 0 {
		q = q.Where(operateLog.OperationIP.Like("%" + in.OperationIp + "%"))
	}
	if len(in.DeptName) > 0 {
		q = q.Where(operateLog.DeptName.Like("%" + in.DeptName + "%"))
	}
	if len(in.OperationName) > 0 {
		q = q.Where(operateLog.OperationName.Like("%" + in.OperationName + "%"))
	}
	if in.OperationStatus != 0 {
		q = q.Where(operateLog.OperationStatus.Eq(in.OperationStatus))
	}
	if len(in.OperationType) > 0 {
		q = q.Where(operateLog.OperationType.Like("%" + in.OperationType + "%"))
	}
	if len(in.OperationUrl) > 0 {
		q = q.Where(operateLog.OperationURL.Like("%" + in.OperationUrl + "%"))
	}
	if len(in.Title) > 0 {
		q = q.Where(operateLog.Title.Like("%" + in.Title + "%"))
	}

	result, count, err := q.Order(operateLog.ID.Desc()).FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询操作日志列表失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询操作日志列表失败")
	}
	var list []*sysclient.OperateLogListData
	for _, item := range result {
		opTime := item.OperationTime.Format("2006-01-02 15:04:05")
		list = append(list, &sysclient.OperateLogListData{
			Id:                item.ID,                // 编号
			Title:             item.Title,             // 系统模块
			OperationType:     item.OperationType,     // 操作类型
			OperationName:     item.OperationName,     // 操作人员
			RequestMethod:     item.RequestMethod,     // 请求方式
			OperationUrl:      item.OperationURL,      // 操作方法
			OperationParams:   item.OperationParams,   // 请求参数
			OperationResponse: item.OperationResponse, // 响应参数
			OperationStatus:   item.OperationStatus,   // 操作状态
			DeptName:          item.DeptName,          // 部门名称
			UseTime:           item.UseTime,           // 执行时长(毫秒)
			Browser:           item.Browser,           // 浏览器
			Os:                item.Os,                // 操作信息
			OperationIp:       item.OperationIP,       // 操作地址
			OperationTime:     opTime,                 // 操作时间
		})
	}

	return &sysclient.QueryOperateLogListResp{
		Total: count,
		List:  list,
	}, nil

}
