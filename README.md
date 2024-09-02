<h1 align="center">
<a href="https://blog.horonlee.com">GoToKube</a>
</h1>

<p align="center">
🐳 轻松管理 Kubernetes 集群 (Beta).
</p>

<pre align="center">
让你更加方便快捷的使用 Kubernetes
🧪 开发中
</pre>

- [English](./README.en-US.md) | **简体中文**

## 主要功能：

- [x] 可以通过控制台查看 Docker 和 Kubernetes 的信息
- [x] 多数据库支持(SQLite MySQL)
- [x] 通过各种请求来使用 yaml 文件对 kubernetes 集群内的资源进行操控
- [x] 查询、创建和删除 Docker 容器

## 用法

⚠️：需要 Docker Client API Version >= 1.45
环境变量JWT密钥`JWT_SECRET_KEY=JWT_TOKEN`和根用户密码`AUTH_PASS`必须设置，否则 Web 服务无法正常使用
本软件大多功能由 API 提供，最好的方式是前往查看 API 文档：https://documenter.getpostman.com/view/34220703/2sA3e5d86S

## 构建方法
### 使用 make 构建

1. 进入项目目录，打开 Makefile
2. 编辑1-6 行的变量为自己需要的内容，一般只需要更改GOOS（你的系统）和GOARCH（系统架构）
3. 在当前目录执行 `make`即可生成二进制文件
4. 给予可执行权限`sudo chmod +x GoToKube`

### 使用 Go 构建

1. 进入项目目录执行`go build`
2. 得到`GoToKube`二进制文件，给予可执行权限`sudo chmod +x GoToKube`

> 使用 Docker 构建

1. 使用项目中的 Dockerfile 进行构建`docker build -t gotokube:dev .`
2. 推荐使用 DockerCompose 启动容器`docker-compose up -d`
   1. 其中，Docker 的 sock 文件必须映射到容器内，否则无法开启软件
      ```yml
      volumes:
        - /var/run/docker.sock:/var/run/docker.sock
      ```

## 配置文件

> 配置文件在第一次运行后会在程序同级目录生成，随后可自行更改，其中大小写**不敏感**

```toml
[auth]  # 管理员认证信息，默认为 root:123456，运行过一次程序后可删除该配置的值
pass = '114514'
user = 'root'

[database]
addr = ''   # 数据库地址
name = ''   # 数据库名称
password = ''   # 数据库密码
path = 'data.db'    # sqlite数据库文件路径，默认为当前目录下的 data.db
type = 'sqlite'    # 数据库类型，默认为 sqlite，目前支持sqlite和mysql
user = ''   # 数据库用户名

[kube]
configpath = '' # kubernetes 配置文件路径，默认为 $HOME/.kube/config
enable = true   # 是否启用 kubernetes 功能

[common]
dir = ''    # 日志文件存放路径
termenable = false  # 是否开启控制台

[web]
enable = true   # 是否启用 web 功能
listeningaddr = ':8080' # web 服务监听地址
```

## 环境变量

> 同配置文件，使用 `配置单元`_`配置项` = `配置值` 来设定环境变量（变量值必须大写），这将会覆盖配置文件的值

示例：

- LOG_DIR='/var/log/gotokub' 日志文件存放路径
- WEB_LISTENINGADDR=":9090" web 服务监听地址
