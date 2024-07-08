<h1 align="center">
<a href="https://blog.horonlee.com">Visual Docker</a>
</h1>

<p align="center">
🐳 轻松管理容器 (Beta).
</p>

<pre align="center">
让你更加方便快捷的使用docker
🧪 开发中
</pre>

- [English](./README.en-US.md) | **简体中文**

## TODO：
- [x] 可以通过控制台查看docker的信息
- [x] 控制台检测到Docker异常会终止程序
- [x] 通过 Web 界面展示信息
- [x] 对接 Kubernetes 集群，可以通过控制台显示所有 Pod
- [x] 多数据库支持(SQLite MySQL)

## 构建方法

> 需要的 Docker Client API Version >= 1.45

1. 进入项目目录执行`go build`
2. 得到`VDController`二进制文件，给予可执行权限`sudo chmod +x VDController`
3. 将`VDController`放到独立文件夹，并且放入项目的 webSrc 文件夹
4. 执行`./VDController`即可开启程序

## 配置文件

> 配置文件在第一次运行后会在程序同级目录生成，随后可自行更改

- `WebEnable = true&false` 开启程序后是否自动开启网页功能
- `ListeningPort = '0.0.0.0:8080'` 网页功能的监听地址以及端口
- `KubeEnable = true&false` 开启程序后是否自动开启 kubernetes 功能
- `KubeconfigPath = '.kube/config 文件路径'` kubernetes 功能的配置文件路径
  - 如果不填写此项，则默认会使用 $HOME/.kube/config`''`
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
KubeconfigPath = '/Users/horonlee/Downloads/k8s/config'
```

## 网页端用法

**Docker** 操作

   > URL 的前缀都是$IP/docker，后面跟随下方的地址

- `/search?ctr=$ImageName` 根据镜像名查看所有使用该镜像创建的Docker容器

**Kubernetes** 操作

   > URL 的前缀都是$IP/kube，后面跟随下方的地址

- `/deployments/$Namespace` 获得该命名空间下的所有 Deployment
- `/deployment/$Namespace/$DeployName` 获得该命名空间该 Deployment 的详细信息
- `/services/$Namespace` 获得该命名空间下的所有 Service
- `/pods/$Namespace` 获得该命名空间下的所有 Pod
  - `/pod/$Namespace/$PodName` 获得该 Pod 的详细信息 
- `/namespaces` 获得所有命名空间


## 环境变量
- LOG_DIR 日志文件存放路径`/var/log/vdcontroller`

## 启动参数

支持通过启动参数来配置软件的设置，如：`./VDController -kubeconfig="/home/user/document/k8s/config"

- `-kubeconfig` Kubernetes配置文件路径