# ClouDisk
## 用户模块
1. 邮箱注册
   1. 用户拿到验证码
   2. 判断用户在mysql中是否已经注册
   3. 已注册，用户停止发送验证码
   4. 未注册，将邮箱和验证码存储在redis
   5. 用户发送验证码，判断验证码和reids中是否一致
   6. 若不一致，注册失败
   7. 一致，将用户信息存储在mysql中
2. 密码登录
   1. 能否成功查询数据库
   2. 用户名和密码是否存在数据库
   3. 登录成功，返回一个token(id,idnetity,name)和refresh_token 用于刷新 token
3. 个人资料详情
   1. 发送用户id，判断mysl中是否存在
## 存储池模块
1. 用户身份验证中间件middleware Auth
   1. 判断能否成功解析请求头中的Authorization ,不成功：身份验证失败，成功：将其解析出user_id,user_identity,user_name到请求头，
2. 中心存储池资源管理
   1. 文件上传
      1. 先经过Auth中间件验证
      2. 根据文件基本信息判断在mysql中是否已存在
      3. 上传腾讯云
         1. 分片上传
         2. 文件秒传
      4. 返回文件基本信息
      5. 不存在，将文件基本信息插入数据库，返回得到repository_identity,ext,name
   
3. 个人存储池资源管理
   1. 文件关联存储
      1. 先经过Auth中间件验证
      2. 得到请求头中的user_identity,
      3. 存储文件的描述信息 包含identity,user_identity，repository_identity,ext,parentid,name,
      4. 返回identity,name,ext
   2. 文件列表
      1. 先经过Auth中间件验证
      2. 得到请求头中userIdentity
      3. 发送id和分页参数：page,size(默认0,1,20)
      4. 返回UserFile数组，统计条目数cnt
   3. 文件名称修改
      1. 传identity,name参数，判断name在当前层级下是否已经存在
      2. 不存在-> 使用Update修改文件名name
   4. 文件删除
      1. xorm是软删除，需要在列表处两表连接时，自定义删除时间为空
   5. 文件移动
4. 文件共享模块
   1. 文件分享
   2. 获取分享资源详情
   3. 资源保存
## 命令和三方包
```
   goctl api new core
   goctl api go -api core.api -dir . -style=goZero

   github.com/jordan-wright/email
   github.com/satori/go.uuid
   github.com/dgrijalva/jwt-go
```
## 163授权码
NELNPMNEBPAETNTZ 