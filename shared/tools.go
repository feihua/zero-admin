package shared

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/tal-tech/go-zero/rest/httpx"
)

type (
	Returns struct {
		Code int         `json:"code"`
		Msg  string      `json:"msg"`
		Data interface{} `json:"data"`
	}
)

func OkJson(w http.ResponseWriter, v interface{}) {

	httpx.OkJson(w, Returns{
		Code: 0,
		Data: v,
	})
}

func GetDate() string {
	return time.Now().Format("2006-01-02")
}

/**
 * 获取 微信登陆的url
 */

func GetWxLoginUrl(appid, fillwebsite string) string {
	return fmt.Sprintf(`https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=code&scope=snsapi_userinfo`, appid, fillwebsite)
}

//字符串转int

func StrToInt64(str string) (id int64) {
	id, _ = strconv.ParseInt(str, 10, 64)
	return
}

func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func Int64ToStr(id int64) (str string) {
	str = fmt.Sprintf("%d", id)
	return
}

/**
 * 字符 转换
 */
func ToLimitOffset(pageIn, pagesizeIn string) (offset int64, limit int64) {
	page := StrToInt64(pageIn)
	limit = StrToInt64(pagesizeIn)
	if page <= 0 {
		page = 1
	}
	offset = limit * (page - 1)
	if limit == 0 {
		limit = 10
	}
	return
}

func HttpPostForm(gourl string, form url.Values) ([]byte, error) {
	resp, err := http.PostForm(gourl, form)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

// 得到一个随机数
func RandInt(max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	if max < 1 {
		return r.Int()
	} else {
		return r.Intn(max)
	}
}
