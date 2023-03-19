说明：
1. 选择 startWebApp 执行的是 web 模式
2. 选择 startDaemon 执行的是 daemon 模式
3. 选择 startCron   执行的是 cron 模式


涉及组件：
1. 路由 mux
2. 数据库管理 gorm


目录结构：
1. api层 接口
2. models层 结构体统一定义位置
3. resposity 层 外部依赖组件
4. route 层 路由配置

启动方式可以参考 Dockerfile 里的命令

TODO：
1.登录注册还有权限管理还有验证码 登录需要用token token产生设计
  注册功能：
     前端传用户名密码   密码Sha3-256 加盐 （盐后台随机生成） 数据库里存加盐后密码 以及盐用于后期校验 
  登录
     前端传用户名密码  后端用密码+盐然后Sha3-256 加密后与数据库内的密码进行比对，一致后确认登录成功。这个时候可以设计token系统了。
  登录成功后，后端返回token。 token 默认六小时，续期时间为按照最后一次请求进行续期，如果最后一次不请求了，就不续期了。

  验证码功能：
     前端验证码
  token系统设计：
      1.token 存redis 中。 过期时间可以由redis自带的失效功能实现。前端拿到token后，请求每个接口时均需要带上，后端做token校验。
      2.用户重新登录时，token也需要重置
      3.token在登录后可以获得。后续前端每次请求都需要带上此token，token失效后，后端进行校验提醒，前端再做

  jwt+token 登录 JsonWebToken
2.需要给产品prd
3.缺失redis 组件
