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

  ⚠️：需要 Docker Client API Version >= 1.45

## 构建方法

### 使用 make 构建

1. 进入项目目录，打开 Makefile
2. 编辑1-6 行的变量为自己需要的内容，一般只需要更改GOOS（你的系统）和GOARCH（系统架构）
3. 在当前目录执行 `make`即可生成二进制文件
4. 给予可执行权限`sudo chmod +x GoToKube`
5. 执行`./GoToKube`即可开启程序

### 使用 Go 构建

1. 进入项目目录执行`go build`
2. 得到`GoToKube`二进制文件，给予可执行权限`sudo chmod +x GoToKube`
3. 执行`./GoToKube`即可开启程序

> 使用 Docker 构建

1. 使用项目中的 Dockerfile 进行构建`docker build -t gotokube:dev .`
2. 推荐使用 DockerCompose 启动容器`docker-compose up -d`
   1. 其中，Docker 的 sock 文件必须映射到容器内，否则无法开启软件
      ```yml
      volumes:
        - /var/run/docker.sock:/var/run/docker.sock
      ```

## 配置文件

> 配置文件在第一次运行后会在程序同级目录生成，随后可自行更改

- `WebEnable = true&false` 开启程序后是否自动开启网页功能
- `ListeningAddr = '0.0.0.0:8080'` 网页功能的监听地址以及端口
- `TermEnable = true&false`是否启用可交互终端
- `KubeEnable = true&false` 开启程序后是否自动开启 kubernetes 功能
- `KubeConfigPath = '.kube/config 文件路径'` kubernetes 功能的配置文件路径
  - 如果不填写此项，则默认会使用 `$HOME/.kube/config`
- `DBType = 'sqlite&mysql'` 数据库类型，默认为 sqlite，目前仅支持 sqlite和mysql
- `DBPath = 'data.db'` 数据库文件路径，默认为程序当前目录的`data.db`
- `DBAddr = '127.0.0.1:3306'` 数据库地址
- `DBUser = 'root'` 数据库用户名
- `DBPass = 'password'` 数据库密码
- `DBName = 'test'` 数据库名称

示例：

```toml
WebEnable = true
ListeningPort = '127.0.0.1:1024'
KubeEnable = true
KubeConfigPath = '/Users/horonlee/Downloads/k8s/config'
```

## 网页端用法

本软件大多功能由 API 提供，最好的方式是前往查看 API 文档：https://documenter.getpostman.com/view/34220703/2sA3e5d86S

## 环境变量
- LOG_DIR 日志文件存放路径`/var/log/vdcontroller`

## 启动参数

支持通过启动参数来配置软件的设置，如：`./GoToKube -kubeconfig="/home/user/document/k8s/config"

- `-KubeConfig` Kubernetes配置文件路径