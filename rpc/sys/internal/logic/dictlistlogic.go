package logic

import (
	"context"
	"encoding/json"
	"fmt"

	"zero-admin/rpc/sys/internal/svc"
	"zero-admin/rpc/sys/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type DictListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDictListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DictListLogic {
	return &DictListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DictListLogic) DictList(in *sys.DictListReq) (*sys.DictListResp, error) {
	all, err := l.svcCtx.DictModel.FindAll(in.Current, in.PageSize)
	count, _ := l.svcCtx.DictModel.Count()

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询字典列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}
	var list []*sys.DictListData
	for _, dict := range *all {
		fmt.Println(dict)
		list = append(list, &sys.DictListData{
			Id:             dict.Id,
			Value:          dict.Value,
			Label:          dict.Label,
			Type:           dict.Type,
			Description:    dict.Description,
			Sort:           int64(dict.Sort),
			Remarks:        dict.Remarks,
			CreateBy:       dict.CreateBy,
			CreateTime:     dict.CreateTime.Format("2006-01-02 15:04:05"),
			LastUpdateBy:   dict.LastUpdateBy,
			LastUpdateTime: dict.LastUpdateTime.Format("2006-01-02 15:04:05"),
			DelFlag:        dict.DelFlag,
		})
	}

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(list)
	logx.WithContext(l.ctx).Infof("查询字典列表信息,参数：%s,响应：%s", reqStr, listStr)
	return &sys.DictListResp{
		Total: count,
		List:  list,
	}, nil

}
