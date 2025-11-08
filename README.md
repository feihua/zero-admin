# 将会在flutter星球更新getx版本的flutter_mall  🎉🎉🎉
# 将会在element plus星球更新element plus版本的的zero-vue-admin  🎉🎉🎉

# Zero-Admin 电商系统

> 注：ORM持久层已经整体切换成gorm
> 后端接口改动较大,前端正在重新适配中

> 基础代码生成工具

## goctl-helper is out! 🎉🎉🎉
**goland可视化插件，基于database生成api和protobuf文件(用于goctl官方插件使用)**

[goctl-helper](https://plugins.jetbrains.com/plugin/25693-goctl-helper)

## 安装
```shell
go install github.com/feihua/generate-code@latest

generate-code golang zero --dsn "root:123456@tcp(127.0.0.1:3306)/demo" --tableNames sys_ --prefix sys_  --rpcClient sysclient --author liufeihua
```
Zero-Admin 是一套基于 go-zero 框架实现的电商系统，采用 Docker 容器化部署，包含前台商城系统和后台管理系统。


## 前台商城系统

### 模块介绍

1. **首页门户**: 提供用户访问网站的入口，展示热门商品和推荐信息。

2. **商品推荐**: 根据用户的历史行为和个人喜好，推荐个性化商品。

3. **商品搜索**: 强大的商品搜索功能，支持关键字搜索、筛选等。

4. **商品展示**: 以优雅的方式展示商品信息，包括详细描述、价格、评价等。

5. **购物车**: 用户可以将喜欢的商品添加到购物车，方便批量购买。

6. **订单流程**: 提供完整的订单流程，包括下单、支付、发货、收货等环节。

7. **会员中心**: 用户可以管理个人信息、查看订单状态、积分等。

8. **帮助中心**: 提供用户常见问题解答、售后政策等信息。

### 技术栈

- go-zero 框架实现，高性能、易扩展。
- 前端采用现代化的前端框架，例如 React 或 Vue。
- Docker 容器化部署，方便快捷。

## 后台管理系统

### 模块介绍

1. **商品管理**: 管理商品信息，包括添加、编辑、删除商品。

2. **订单管理**: 实时监控订单状态，支持订单发货、取消等操作。

3. **会员管理**: 管理用户信息，包括注册用户、会员等级等。

4. **促销管理**: 管理营销活动，例如满减、打折等。

5. **运营管理**: 管理广告、推广等运营活动。

6. **内容管理**: 管理网站内容，包括公告、资讯等。

7. **权限管理**: 管理系统用户权限，确保安全性。

8. **设置**: 系统配置，包括支付方式、物流信息等。

### 技术栈

- go-zero 框架提供后台接口支持。
- 使用现代化的前端框架进行界面开发。
- 数据库采用 mysql。
- Docker 容器化部署，方便管理和维护。

# 文档地址
[https://feihua.github.io/](https://feihua.github.io/) 正在完善
#
[zero-admin-ui是后台的pc管理端](https://github.com/feihua/zero-admin-ui)是一个基于react实现的管理后台

[flutter_mall是zero-admin的app端](https://github.com/feihua/flutter_mall)是一个Flutter的电商实战项目，包括首页、列表页、详细页、购物车页、会员中心和支付(支付对接的是支付宝)

[zero-pc-web 是 zero-admin 的网页端](https://github.com/feihua/zero-pc-web)zero-pc-web 是一个基于 React 框架实现的 web
端电商系统(预览地址[http://110.41.179.89/pc/](http://110.41.179.89/pc/))



# android版本
android版本体验地址 [flutter-mall-app](https://www.pgyer.com/OoW2Zy)


# 项目模板
[zero-admin-template](https://github.com/feihua/zero-admin-template)(只包含基础的rbac权限)

# 1.项目预览

**预览地址**http://129.204.203.29/mall <span  style="color: red;"> 账号：admin 密码: 123456</span>
**预览地址（vue）**http://129.204.203.29/vue/login <span  style="color: red;"> 账号：admin 密码: 123456</span>
> 注：演示账号部分功能修改删除权限未开放。


# 2.感谢
[go-zero](https://github.com/zeromicro/go-zero)
<p></p>

[mall](https://github.com/macrozheng/mall)

## 许可证

本项目采用 Apache License 2.0 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情

如果您觉得这个项目对您有帮助，请给我们一个 ⭐️，这将鼓励我们持续改进！
