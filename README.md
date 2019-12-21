# Go-bjut
# 项目规范
## git commit
一个 commit 尽可能只做一件事
1. added: 添加，一般在添加了新功能时使用
2. improved: 改进，一般在优化和改进代码时使用
3. refactored：重构，一般在优化代码结构和设计时使用
4. fixbug: 修复bug，一般在修复bug时使用
## 命名规范
1. 文件名统一用小驼峰法，例如：bjutRegister
2. 路由命名统一用"-"，例如：bjut-register
3. 包名统一用小写字母，例如：student
## 版本约定
1. 采用主版本号.子版本号.修正版本号，比如：V1.2.1
2. 修正版本号一般在修复bug时使用
3. 子版本号一般在添加了新功能时使用
4. 主版本号一般是在积累了较多新功能且代码稳定时使用
## 更新日志
1. added: 添加的新功能
2. changed: 功能的变更
3. fixed: 修改的bug
4. security: 修改的关于安全的bug
5. removed: 删除的功能
## 前后端数据交互
1. json
2. 时间统一用rfc3339格式
