# Changelog
## v0.1.0 2019-11-15
### added
- viper配置文件读取
- mysql数据库连接池
- logger日志模块
- cookie用户登录状态管理
- student模型的增删改查
## v0.2.0 2019-12-02
### added
- 丢失一卡通的登记和展示
- 给一卡通失主自动发短信通知
- 通过配置文件选择gorm的模式
### changed
- 用户系统修改为一张表
### fixed
- 请求小美接口后挥手
- bjutRegister通过学号来判断是否重复注册
### security
- 修改密码时需要输入原密码
## v0.2.1 2019-12-16
### added
- unique 字段显示的创建索引
### changed
- gin v1.4.0 -> v1.5.0
- 数据校验充分利用 validator tag
- database合入model包中
- 删除多余的err.error()
### fixed
- sms请求url中random参数
- sms返回值序列化的错误处理
