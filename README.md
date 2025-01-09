# kv-auth-db

bitcask 内嵌数据库+auth 用户权限管理+可视化页面+docker compose 部署

## 项目结构

1. backend go后端
2. web react 前端
3. docker-configs docker-compose部署，基础组件数据和配置映射。
    1. mysql 数据存储数据库
    2. redis 缓存
    3. 其他必要基础组件

### Backend 组织架构

* app 应用层
    * common 服务请求request和响应response 定义、封装
    * middleware 中间件 jwt认证中间件等
    * models 模型层
    * services 服务层/逻辑层
* bootstrap 项目启动相关。1、基础组件初始化（配置文件加载，数据库，redis，自定义验证等）2、路由组
* config 系统配置
* controllers 控制层。处理入参（app/common/*下定义request），调用服务层功能（app/services/*）,处理响应。概括为定义gin.HandleFunc
* global 系统全局变量
* routes 路由层
* utils 系统全局工具
* config.yaml 系统配置文件
* main.go 系统入口/启动文件
