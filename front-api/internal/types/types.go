// Code generated by goctl. DO NOT EDIT.
package types

type AddressListReq struct {
	Current  int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
}

type AddressListResp struct {
	Code int64           `json:"code"`
	Data AddressListData `json:"data"`
	Msg  string          `json:"msg"`
}

type AddressListData struct {
	Total int64         `json:"total"` //总数
	Pages int64         `json:"pages"` //总页数
	Limit int64         `json:"limit"` //分页数量
	Page  int64         `json:"page"`  //当前页
	List  []AddressList `json:"list"`  //地址列表
}

type AddressList struct {
	ID            int64  `json:"id"`            //id
	Name          string `json:"name"`          //用户名
	UserID        int64  `json:"userId"`        //用户Id
	Province      string `json:"province"`      //省
	City          string `json:"city"`          //市
	Region        string `json:"region"`        //区
	AddressDetail string `json:"addressDetail"` //详情地址
	AreaCode      string `json:"areaCode"`      //地区编码
	Tel           string `json:"tel"`           //电话
	IsDefault     bool   `json:"isDefault"`     //是否为默认地址
	AddTime       string `json:"addTime"`       //添加时间
	UpdateTime    string `json:"updateTime"`    //更新时间
	Deleted       bool   `json:"deleted"`       //是否删除
}

type AddressSaveReq struct {
	Name          string `json:"name"`
	Tel           string `json:"tel"`
	Province      string `json:"province"`
	City          string `json:"city"`
	Region        string `json:"region"`
	AddressDetail string `json:"addressDetail"`
	PostalCode    string `json:"postalCode"`
}

type AddressSaveResp struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}

type AddressUpdateReq struct {
	AddressID     int64  `json:"addressID"`
	Name          string `json:"name"`
	Tel           string `json:"tel"`
	Province      string `json:"province"`
	City          string `json:"city"`
	Region        string `json:"region"`
	AddressDetail string `json:"addressDetail"`
	PostalCode    string `json:"postalCode"`
}

type AddressUpdateResp struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}

type AddressDeleteReq struct {
	AddressID int64 `json:"addressID"`
}

type AddressDeleteResp struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}

type AddressDetailReq struct {
	AddressID int64 `json:"addressID"`
}

type AddressDetailResp struct {
	Code int64             `json:"code"`
	Data AddressDetailData `json:"data"`
	Msg  string            `json:"msg"`
}

type AddressDetailData struct {
	IsDeleted     bool   `json:"IS_DELETED"`
	NotDeleted    bool   `json:"NOT_DELETED"`
	ID            int64  `json:"id"`            //id
	Name          string `json:"name"`          //用户名
	UserID        int64  `json:"userId"`        //用户Id
	Province      string `json:"province"`      //省
	City          string `json:"city"`          //市
	County        string `json:"county"`        //国家
	AddressDetail string `json:"addressDetail"` //详情地址
	AreaCode      string `json:"areaCode"`      //地区编码
	PostalCode    string `json:"postalCode"`
	Tel           string `json:"tel"`        //电话
	IsDefault     bool   `json:"isDefault"`  //是否为默认地址
	AddTime       string `json:"addTime"`    //添加时间
	UpdateTime    string `json:"updateTime"` //更新时间
	Deleted       bool   `json:"deleted"`    //是否删除
}

type LoginReq struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type LoginAndRegisterResp struct {
	Errno  int64                `json:"errno"`
	Data   LoginAndRegisterData `json:"data"`
	Errmsg string               `json:"errmsg"`
}

type LoginAndRegisterData struct {
	UserInfo     UserInfo `json:"userInfo"` //用户基本信息
	Token        string   `json:"token"`    //服务端返回的token值
	AccessExpire int64    `json:"accessExpire"`
	RefreshAfter int64    `json:"refreshAfter"`
}

type UserInfo struct {
	NickName string `json:"nickName"` //昵称
	Phone    string `json:"phone"`
	Icon     string `json:"icon"` //头像地址
}

type RegisterReq struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type WXMiniAuthReq struct {
	Code          string `json:"code"`
	IV            string `json:"iv"`
	EncryptedData string `json:"encryptedData"`
	PNCode        string `json:"pncode"`
}

type CartListReq struct {
	UserId int64 `path:"userId"`
}

type CartListResp struct {
	Errno  int64        `json:"errno"`
	Data   CartListData `json:"data"`
	Errmsg string       `json:"errmsg"`
}

type CartListData struct {
	CartTotal CartTotal  `json:"cartTotal"` //统计数据
	CartList  []CartList `json:"cartList"`  //商品数据列表
}

type CartList struct {
	ID             int64    `json:"id"`             //id
	UserID         int64    `json:"userId"`         //用户Id
	GoodsID        int64    `json:"goodsId"`        //商品Id
	GoodsSn        string   `json:"goodsSn"`        //商品编号
	GoodsName      string   `json:"goodsName"`      //商品名称
	ProductID      int64    `json:"productId"`      //产品规格Id
	Price          int64    `json:"price"`          //价格
	Number         int64    `json:"number"`         //数量
	Specifications []string `json:"specifications"` //规格值
	Checked        bool     `json:"checked"`        //是否选择
	PicURL         string   `json:"picUrl"`         //图片
	AddTime        string   `json:"addTime"`        //添加时间
	UpdateTime     string   `json:"updateTime"`     //更新时间
	Deleted        bool     `json:"deleted"`        //是否删除
}

type CartTotal struct {
	GoodsCount        int64 `json:"goodsCount"`        //商品总数
	CheckedGoodsCount int64 `json:"checkedGoodsCount"` //选择的商品总数
}

type CartAddReq struct {
	IsDeleted      bool     `json:"IS_DELETED"`
	NotDeleted     bool     `json:"NOT_DELETED"`
	ID             int64    `json:"id"`             //id
	UserID         int64    `json:"userId"`         //用户Id
	GoodsID        int64    `json:"goodsId"`        //商品Id
	GoodsSn        string   `json:"goodsSn"`        //商品编号
	GoodsName      string   `json:"goodsName"`      //商品名称
	ProductID      int64    `json:"productId"`      //产品规格Id
	Price          int64    `json:"price"`          //价格
	Number         int64    `json:"number"`         //数量
	Specifications []string `json:"specifications"` //规格值
	Checked        bool     `json:"checked"`        //是否选择
	PicURL         string   `json:"picUrl"`         //图片
	AddTime        string   `json:"addTime"`        //添加时间
	UpdateTime     string   `json:"updateTime"`     //更新时间
	Deleted        bool     `json:"deleted"`        //是否删除
}

type CartAddResp struct {
	Errno  int64  `json:"errno"`
	Errmsg string `json:"errmsg"`
}

type CartFastAddReq struct {
	IsDeleted      bool     `json:"IS_DELETED"`
	NotDeleted     bool     `json:"NOT_DELETED"`
	ID             int64    `json:"id"`             //id
	UserID         int64    `json:"userId"`         //用户Id
	GoodsID        int64    `json:"goodsId"`        //商品Id
	GoodsSn        string   `json:"goodsSn"`        //商品编号
	GoodsName      string   `json:"goodsName"`      //商品名称
	ProductID      int64    `json:"productId"`      //产品规格Id
	Price          int64    `json:"price"`          //价格
	Number         int64    `json:"number"`         //数量
	Specifications []string `json:"specifications"` //规格值
	Checked        bool     `json:"checked"`        //是否选择
	PicURL         string   `json:"picUrl"`         //图片
	AddTime        string   `json:"addTime"`        //添加时间
	UpdateTime     string   `json:"updateTime"`     //更新时间
	Deleted        bool     `json:"deleted"`        //是否删除
}

type CartFastAddResp struct {
	Errno  int64             `json:"errno"`
	Data   []CartFastAddData `json:"data"`
	Errmsg string            `json:"errmsg"`
}

type CartFastAddData struct {
}

type CartCheckedReq struct {
	UserID     int64   `json:"userId"`
	ProductIds []int64 `json:"productIds"`
	IsChecked  int64   `json:"isChecked"`
}

type CartCheckOutReq struct {
	UserID     int64   `json:"userId"`
	ProductIds []int64 `json:"productIds"`
}

type CartCheckOutResp struct {
	Errno  int64            `json:"errno"`
	Data   CartCheckOutData `json:"data"`
	Errmsg string           `json:"errmsg"`
}

type CartCheckOutData struct {
	CheckedAddress   CheckedAddress     `json:"checkedAddress"`
	ActualPrice      int64              `json:"actualPrice"`
	OrderTotalPrice  int64              `json:"orderTotalPrice"`
	CartID           int64              `json:"cartId"`
	FreightPrice     int64              `json:"freightPrice"`
	CheckedGoodsList []CheckedGoodsList `json:"checkedGoodsList"`
	GoodsTotalPrice  int64              `json:"goodsTotalPrice"`
	AddressID        int64              `json:"addressId"`
}

type CheckedAddress struct {
	ID int64 `json:"id"`
}

type CheckedGoodsList struct {
	ID             int64    `json:"id"`
	UserID         int64    `json:"userId"`
	GoodsID        int64    `json:"goodsId"`
	GoodsSn        string   `json:"goodsSn"`
	GoodsName      string   `json:"goodsName"`
	ProductID      int64    `json:"productId"`
	Price          int64    `json:"price"`
	Number         int64    `json:"number"`
	Specifications []string `json:"specifications"`
	Checked        bool     `json:"checked"`
	PicURL         string   `json:"picUrl"`
	AddTime        string   `json:"addTime"`
	UpdateTime     string   `json:"updateTime"`
	Deleted        bool     `json:"deleted"`
}

type CartDeleteReq struct {
	UserId     int64   `json:"userId"`
	ProductIds []int64 `json:"productIds"`
}

type CartUpdateReq struct {
	IsDeleted      bool   `json:"IS_DELETED"`
	NotDeleted     bool   `json:"NOT_DELETED"`
	ID             int64  `json:"id"`
	UserID         int64  `json:"userId"`
	GoodsID        int64  `json:"goodsId"`
	GoodsSn        string `json:"goodsSn"`
	GoodsName      string `json:"goodsName"`
	ProductID      int64  `json:"productId"`
	Price          string `json:"price"`
	Number         string `json:"number"`
	Specifications string `json:"specifications"`
	Checked        bool   `json:"checked"`
	PicURL         string `json:"picUrl"`
	AddTime        string `json:"addTime"`
	UpdateTime     string `json:"updateTime"`
	Deleted        bool   `json:"deleted"`
}

type CartUpdateResp struct {
	Errno  int64  `json:"errno"`
	Errmsg string `json:"errmsg"`
}

type CardListReq struct {
	Page  int64 `form:"page"`
	Limit int64 `form:"limit"`
}

type CardListResp struct {
	Errno  int64        `json:"code"`
	Data   CardListData `json:"data"`
	Errmsg string       `json:"msg"`
}

type CardListData struct {
	CardList []CardList `json:"cardList"` //商品数据列表
}

type CardList struct {
	ID             int64   `json:"id"`             //id
	Name           string  `json:"name"`           //名称
	FaceValue      float64 `json:"faceValue"`      //面值
	ExpiredDay     int64   `json:"expiredDay"`     //失效周期（天）
	Status         int64   `json:"status"`         //状态
	CommissionRate float64 `json:"commissionRate"` //佣金率
	DiscountRate   float64 `json:"discountRate"`   //折扣率
	Note           string  `json:"note"`           //备注
}

type CardBuyReq struct {
	CardId int64 `form:"cardId"` //卡号
}

type CardBuyResp struct {
	Errno  int64  `json:"code"`
	Errmsg string `json:"msg"`
}

type CardMyListReq struct {
	Page  int64 `form:"page"`
	Limit int64 `form:"limit"`
}

type CardMyListResp struct {
	Errno  int64          `json:"code"`
	Data   CardMyListData `json:"data"`
	Errmsg string         `json:"msg"`
}

type CardMyListData struct {
	CardMyList []CardMyList `json:"cardMyList"` //商品数据列表
}

type CardMyList struct {
	ID            int64   `json:"id"`
	PrepaidCardId int64   `json:"prepaidCardId"`
	PrepaidCardSn string  `json:"prepaidCardSn"`
	Balance       float64 `json:"balance"`
	Status        int64   `json:"status"`
	CreateTime    string  `json:"createTime"` //状态
	ExpiredTime   string  `json:"expiredTime"`
	UpdateTime    string  `json:"updateTime"`
	Note          string  `json:"note"` //备注
}

type CategoryResp struct {
	Errno  int64          `json:"code"`
	Data   []CategoryData `json:"data"`
	Errmsg string         `json:"msg"`
}

type CategoryData struct {
	ID         int64  `json:"id"`        //分类id
	Name       string `json:"name"`      //名称
	Keywords   string `json:"keywords"`  //关键字
	Desc       string `json:"desc"`      //描述
	PID        int64  `json:"pid"`       //父id
	IconURL    string `json:"iconUrl"`   //图标地址
	PicURL     string `json:"picUrl"`    //图片地址
	Level      string `json:"level"`     //等级
	SortOrder  int64  `json:"sortOrder"` //排序规则
	AddTime    string `json:"addTime"`
	UpdateTime string `json:"updateTime"`
	Deleted    bool   `json:"deleted"`
}

type SecondCategoryReq struct {
	Id int64 `form:"id"`
}

type CollectListReq struct {
	MemberId int64 `json:"memberId"`
}

type CollectListResp struct {
	Errno  int64             `json:"errno"`
	Data   []CollectListData `json:"data"`
	Errmsg string            `json:"errmsg"`
}

type CollectListData struct {
	ID          int32  `json:"id"`          //id
	Type        int32  `json:"type"`        //类型
	ValueID     int32  `json:"valueId"`     //商品Id
	Name        string `json:"name"`        //名称
	Brief       string `json:"brief"`       //描述
	PicUrl      string `json:"picUrl"`      //图片
	RetailPrice string `json:"retailPrice"` //价格
}

type AddOrDeleteReq struct {
	Type     int32 `json:"type"`    //类型
	ValueID  int32 `json:"valueId"` //商品Id
	MemberId int32 `json:"memberId"`
}

type AddOrDeleteResp struct {
	Errno  int64  `json:"errno"`
	Errmsg string `json:"errmsg"`
}

type HomeResp struct {
	Errno  int64    `json:"errno"`
	Data   HomeData `json:"data"`
	Errmsg string   `json:"errmsg"`
}

type HomeData struct {
	NewGoodsList []GoodsList `json:"newGoodsList"` //新品列表
	Channel      []Channel   `json:"channel"`      //分类列表
	Banner       []Banner    `json:"banner"`       //轮播图列表
	HotGoodsList []GoodsList `json:"hotGoodsList"` //热卖商品列表
}

type Banner struct {
	ID         int64  `json:"id"`       //轮播图id
	Name       string `json:"name"`     //名称
	Link       string `json:"link"`     //轮播图链接
	URL        string `json:"url"`      //轮播图地址
	Position   int64  `json:"position"` //位置
	Content    string `json:"content"`  //内容
	Enabled    bool   `json:"enabled"`  //是否启用
	AddTime    string `json:"addTime"`
	UpdateTime string `json:"updateTime"`
	Deleted    bool   `json:"deleted"`
}

type Channel struct {
	ID      int64  `json:"id"`      //分类id
	Name    string `json:"name"`    //分类名称
	IconURL string `json:"iconUrl"` //分类图标地址
}

type GoodsList struct {
	ID           int64  `json:"id"`           //商品id
	Name         string `json:"name"`         //商品名称
	Brief        string `json:"brief"`        //商品简介
	PicURL       string `json:"picUrl"`       //商品图片
	IsNew        bool   `json:"isNew"`        //是否为新品
	IsHot        bool   `json:"isHot"`        //是否为热卖商品
	CounterPrice int64  `json:"counterPrice"` //专柜价格
	RetailPrice  int64  `json:"retailPrice"`  //零售价格
}

type Good struct {
	Id  int64 `json:"id"`
	Num int64 `json:"num"`
}

type OrderCreateReq struct {
	Goods     []Good `json:"goods"`
	AddressId int64  `json:"addressId"`
	PaymentId int64  `json:"paymentId"`
	Note      string `json:"note"`
}

type OrderCreateResp struct {
	Errno  int64     `json:"code"`
	Data   OrderData `json:"data"`
	Errmsg string    `json:"msg"`
}

type OrderData struct {
	OrderId int64  `json:"orderId"`
	OrderSn string `json:"orderSn"`
}

type OrderDetailReq struct {
	OrderId int64 `form:"orderId"`
}

type OrderDetailResp struct {
	Errno  int64           `json:"code"`
	Data   OrderDetailData `json:"data"`
	Errmsg string          `json:"msg"`
}

type OrderDetailData struct {
	Id              int32           `json:"id"`
	OrderSn         string          `json:"orderSn"`
	ActualPrice     string          `json:"actualPrice"`
	OrderStatusText string          `json:"orderStatusText"`
	HandleOption    string          `json:"handleOption"`
	CreateTime      string          `json:"createTime"`
	GoodsList       []GoodsListData `json:"goodsList"`
	IsDelivery      int64           `json:"isDelivery"`
}

type OrderListReq struct {
	Page   int64 `form:"page"`
	Limit  int64 `form:"limit"`
	Status int64 `form:"status"`
}

type OrderListResp struct {
	Errno  int64           `json:"code"`
	Data   []OrderListData `json:"data"`
	Errmsg string          `json:"msg"`
}

type OrderListData struct {
	Id              int32           `json:"id"`
	OrderSn         string          `json:"orderSn"`
	ActualPrice     string          `json:"actualPrice"`
	OrderStatusText string          `json:"orderStatusText"`
	HandleOption    string          `json:"handleOption"`
	CreateTime      string          `json:"createTime"`
	GoodsList       []GoodsListData `json:"goodsList"`
	IsDelivery      int64           `json:"isDelivery"`
}

type GoodsListData struct {
	Id             int32  `json:"id"`
	GoodsName      string `json:"goodsName"`
	Number         string `json:"number"`
	PicUrl         string `json:"picUrl"`
	Specifications string `json:"specifications"`
	Price          string `json:"price"`
	Quantity       int64  `json:"quantity"`
}

type OrderCancelReq struct {
	OrderId int64 `form:"orderId"`
}

type OrderCancelResp struct {
	Errno  int64  `json:"code"`
	Errmsg string `json:"msg"`
}

type OrderRefundReq struct {
	OrderId int64 `form:"orderId"`
}

type OrderRefundResp struct {
	Errno  int64  `json:"code"`
	Errmsg string `json:"msg"`
}

type OrderConfirmReq struct {
	OrderId int64 `form:"orderId"`
}

type OrderConfirmResp struct {
	Errno  int64  `json:"code"`
	Errmsg string `json:"msg"`
}

type OrderDeleteReq struct {
	UserId  int64 `json:"userId"`
	OrderId int64 `json:"orderId"`
}

type OrderDeleteResp struct {
	Errno  int64  `json:"code"`
	Errmsg string `json:"msg"`
}

type OrderGoodsReq struct {
	OrderId int64 `form:"orderId"`
	GoodsId int64 `form:"goodsId"`
}

type OrderGoodsResp struct {
	Errno  int64            `json:"code"`
	Data   []OrderGoodsData `json:"data"`
	Errmsg string           `json:"msg"`
}

type OrderGoodsData struct {
	IsDeleted      bool   `json:"IS_DELETED"`
	NotDeleted     bool   `json:"NOT_DELETED"`
	ID             int64  `json:"id"`        //id
	OrderID        int64  `json:"orderId"`   //订单Id
	GoodsID        int64  `json:"goodsId"`   //商品Id
	GoodsName      string `json:"goodsName"` //商品名称
	GoodsSn        string `json:"goodsSn"`   //商品编号
	ProductID      int64  `json:"productId"` //产品规格Id
	Number         string `json:"number"`    //商品数量
	Price          string `json:"price"`     //商品价格
	Specifications string `json:"specifications"`
	PicURL         string `json:"picUrl"`  //商品图片
	Comment        int64  `json:"comment"` //评论
	AddTime        string `json:"addTime"`
	UpdateTime     string `json:"updateTime"`
	Deleted        bool   `json:"deleted"`
}

type OrderCommentReq struct {
	UserId       int64    `json:"userId"`
	OrderGoodsId int64    `json:"orderGoodsId"`
	Content      string   `json:"content"`
	PicUrls      []string `json:"picUrls"`
}

type OrderCommentResp struct {
	Errno  int64  `json:"code"`
	Errmsg string `json:"msg"`
}

type GoodsDetailReq struct {
	Id int64 `form:"id"`
}

type GoodsDetailResp struct {
	Errno  int64           `json:"code"`
	Data   GoodsDetailData `json:"data"`
	Errmsg string          `json:"msg"`
}

type GoodsDetailData struct {
	SpecificationList []SpecificationList `json:"specificationList"`
	Issue             []Issue             `json:"issue"` //常见问题
	UserHasCollect    int64               `json:"userHasCollect"`
	ShareImage        string              `json:"shareImage"`
	Share             bool                `json:"share"`
	Attribute         []interface{}       `json:"attribute"`   //商品属性
	ProductList       []ProductList       `json:"productList"` //商品规格
	Info              Info                `json:"info"`        //商品信息
}

type Info struct {
	ID           int64    `json:"id"`           //id
	GoodsSn      string   `json:"goodsSn"`      //商品编号
	Name         string   `json:"name"`         //名称
	CategoryID   int64    `json:"categoryId"`   //分类Id
	Gallery      []string `json:"gallery"`      //图片
	Keywords     string   `json:"keywords"`     //关键字
	Brief        string   `json:"brief"`        //简介
	IsOnSale     bool     `json:"isOnSale"`     //是否在售
	SortOrder    int64    `json:"sortOrder"`    //排序规则
	PicURL       string   `json:"picUrl"`       //图片地址
	ShareURL     string   `json:"shareUrl"`     //分享地址
	IsNew        bool     `json:"isNew"`        //是否为新品
	IsHot        bool     `json:"isHot"`        //是否为热卖商品
	Unit         string   `json:"unit"`         //单位
	CounterPrice int64    `json:"counterPrice"` //专柜价格
	RetailPrice  int64    `json:"retailPrice"`  //零售价格
	AddTime      string   `json:"addTime"`
	UpdateTime   string   `json:"updateTime"`
	Deleted      bool     `json:"deleted"`
	Detail       string   `json:"detail"`
}

type Issue struct {
	ID         int64  `json:"id"`       //id
	Question   string `json:"question"` //问
	Answer     string `json:"answer"`   //答
	AddTime    string `json:"addTime"`
	UpdateTime string `json:"updateTime"`
	Deleted    bool   `json:"deleted"`
}

type ProductList struct {
	ID             int64    `json:"id"`             //id
	GoodsID        int64    `json:"goodsId"`        //商品Id
	Specifications []string `json:"specifications"` //规格值
	Price          int64    `json:"price"`          //价格
	Number         int64    `json:"number"`         //数量
	URL            string   `json:"url"`            //图片路径
	AddTime        string   `json:"addTime"`
	UpdateTime     string   `json:"updateTime"`
	Deleted        bool     `json:"deleted"`
}

type SpecificationList struct {
	Name      string      `json:"name"`
	ValueList []ValueList `json:"valueList"`
}

type ValueList struct {
	ID            int64  `json:"id"`
	GoodsID       int64  `json:"goodsId"`
	Specification string `json:"specification"`
	Value         string `json:"value"`
	PicURL        string `json:"picUrl"`
	AddTime       string `json:"addTime"`
	UpdateTime    string `json:"updateTime"`
	Deleted       bool   `json:"deleted"`
}

type GoodsCategoryReq struct {
	Id int64 `form:"id"`
}

type GoodsCategoryResp struct {
	Errno  int64             `json:"code"`
	Data   GoodsCategoryData `json:"data"`
	Errmsg string            `json:"msg"`
}

type GoodsCategoryData struct {
	Total int64               `json:"total"`
	Pages int64               `json:"pages"`
	Limit int64               `json:"limit"`
	Page  int64               `json:"page"`
	List  []GoodsCategoryList `json:"list"`
}

type GoodsCategoryList struct {
	ID           int64   `json:"id"`           //商品id
	Name         string  `json:"name"`         //商品名称
	Brief        string  `json:"brief"`        //商品简介
	PicURL       string  `json:"picUrl"`       //商品图片
	IsNew        bool    `json:"isNew"`        //是否为新品
	IsHot        bool    `json:"isHot"`        //是否为热卖商品
	CounterPrice float64 `json:"counterPrice"` //专柜价格
	RetailPrice  float64 `json:"retailPrice"`  //零售价格
	Sale         int64   `json:"sale"`
}
