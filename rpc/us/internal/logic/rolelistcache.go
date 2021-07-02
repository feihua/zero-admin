package logic

import (
	"errors"
	"github.com/tal-tech/go-zero/core/stores/redis"
	"go-zero-admin/rpc/model/usmodel"
	"go-zero-admin/rpc/us/internal/svc"
	"strconv"
)

//type Content struct {
//	Id         string    `json:"id"`
//	Title      string    `json:"title"`
//	Content    string    `json:"content"`
//	CreateTime time.Time `json:"create_time"`
//}

const bizContentCacheKey = `biz#role#cache`

//func RoleRedisCacheExists(r *redis.Redis) (bool, error) {
//	return r.Exists(bizContentCacheKey)
//}

// AddContent 提供内容存储
func AddRole(r *redis.Redis, c *usmodel.UsRoles) error {
	v := compress(c)
	_, err := r.Zadd(bizContentCacheKey, c.Id, v)
	return err
}

//
func AddRoleId(r *redis.Redis, id int64) error {
	s := strconv.FormatInt(id, 10)
	_, err := r.Zadd(bizContentCacheKey,id,s)
	return err
}

// DelContent 提供内容删除
func DelRole(r *redis.Redis, c *usmodel.UsRoles) error {
	v := compress(c)
	_, err := r.Zrem(bizContentCacheKey, v)
	return err
}

func DelRoleId(r *redis.Redis, id int64) error {
	s := strconv.FormatInt(id, 10)
	_, err := r.Zrem(bizContentCacheKey, s)
	return err
}

func DelAllRole(r *redis.Redis) error {
	_, err := r.Zremrangebyscore(bizContentCacheKey, 0,20)
	return err
}

func SetRoleBizCacheExpire(r *redis.Redis,seconds int){
	r.Expire(bizContentCacheKey,seconds)
}
// 内容压缩
//func compress(c *usmodel.UsRoles) string {
//	// todo: do it yourself
//	var ret string
//	return ret
//}

func compress(c *usmodel.UsRoles) string {
	return strconv.Itoa(int(c.Id))
}

// 内容解压
//func uncompress(v string) *Content {
//	// todo: do it yourself
//	var ret Content
//	return &ret
//}

func uncompress(v string) string {
	// todo: do it yourself
	return v
}

func redisListRolesByScoreId(svcCtx *svc.ServiceContext, start, end int64) (*[]usmodel.UsRoles, error) {
	kvs, err := svcCtx.RedisConn.ZrangebyscoreWithScores(bizContentCacheKey, start, end)
	if err != nil {
		return nil, err
	}

	usRoles := make([]usmodel.UsRoles, 0)
	//usRoleModel := usmodel.NewUsRolesModel(svcCtx.Config, c)
	for _, kv := range kvs {
		//fmt.Println("kv.Ke:" + kv.Key)

		idStr := uncompress(kv.Key)
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			continue
		}
		usRole, err := svcCtx.UsRolesModel.FindOne(id)
		if err != nil {
			continue
		}
		itemRole := usmodel.UsRoles{
			Id:       usRole.Id,
			RoleName: usRole.RoleName,
			Remark:   usRole.Remark,
			CreateTime: usRole.CreateTime,
			UpdateTime: usRole.UpdateTime,
			DeleteTime: usRole.DeleteTime,
		}
		usRoles = append(usRoles, itemRole)
		//ids = append(ids,usRole.Id)
	}
	if len(usRoles) == 0{
		return &usRoles, errors.New("Redis 找不到角色")
	}

	return &usRoles, nil
}

func RedisListAllRoles(svcCtx *svc.ServiceContext) (*[]usmodel.UsRoles, error) {
	return redisListRolesByScoreId(svcCtx, 0, 10)
}

// ListByRangeTime提供根据时间段进行数据查询
//func ListByRangeTime(r redis.Redis, start, end time.Time) ([]*Content, error) {
//	kvs, err := r.ZrangebyscoreWithScores(bizContentCacheKey, start.UnixNano()/1e6, end.UnixNano()/1e6)
//	if err != nil {
//		return nil, err
//	}
//
//	var list []*Content
//	for _, kv := range kvs {
//		data := uncompress(kv.Key)
//		list = append(list, data)
//	}
//
//	return list, nil
//}
