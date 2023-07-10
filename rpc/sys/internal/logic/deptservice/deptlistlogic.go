package deptservicelogic

import (
	"context"
	"encoding/json"
	"strconv"
	"strings"
	"zero-admin/rpc/sys/sysclient"

	"zero-admin/rpc/sys/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeptListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeptListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeptListLogic {
	return &DeptListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeptListLogic) DeptList(in *sysclient.DeptListReq) (*sysclient.DeptListResp, error) {
	count, _ := l.svcCtx.DeptModel.Count(l.ctx)
	all, err := l.svcCtx.DeptModel.FindAll(l.ctx, 1, count)

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询机构列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	var list []*sysclient.DeptListData

	for _, dept := range *all {
		var parentIds []int64
		if len(dept.ParentIds) > 0 {
			for _, i := range strings.Split(dept.ParentIds, ",") {
				p, _ := strconv.ParseInt(i, 10, 64)
				parentIds = append(parentIds, p)
			}
		}

		list = append(list, &sysclient.DeptListData{
			Id:             dept.Id,
			Name:           dept.Name,
			ParentId:       dept.ParentId,
			OrderNum:       dept.OrderNum,
			CreateBy:       dept.CreateBy,
			CreateTime:     dept.CreateTime.Format("2006-01-02 15:04:05"),
			LastUpdateBy:   dept.UpdateBy.String,
			LastUpdateTime: dept.UpdateTime.Time.Format("2006-01-02 15:04:05"),
			DelFlag:        dept.DelFlag,
			ParentIds:      parentIds,
		})
	}

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(list)
	logx.WithContext(l.ctx).Infof("查询机构列表信息,参数：%s,响应：%s", reqStr, listStr)
	return &sysclient.DeptListResp{
		Total: count,
		List:  list,
	}, nil

}
