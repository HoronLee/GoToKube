# VisualDocker

> 让你更加方便快捷的使用docker

目前进度：
1. 可以通过控制台查看docker的信息
2. 控制台检测到Docker异常会终止程序
3. 通过 Web 界面展示信息
4. 对接 Kubernetes 集群，可以通过控制台显示所有 Pod
5. soon...

## 使用方法

1. 进入项目目录执行`go build`
2. 得到`VDController`二进制文件，给予可执行权限`sudo chmod +x VDController`
3. 将`VDController`放到独立文件夹，并且放入项目的 webSrc 文件夹
4. 执行`./VDController`即可开启程序

## 配置文件

> 配置文件在第一次运行后会在程序同级目录生成，随后可自行更改

- WebEnable = true&false 开启程序后是否自动开启网页功能
- ListeningPort = '0.0.0.0:8080' 网页功能的监听地址以及端口
- KubeEnable = true&false 开启程序后是否自动开启 kubernetes 功能
- KubeconfigPath = '.kube/config 文件路径' kubernetes 功能的配置文件路径
  - 如果不填写此项，则默认会使用 $HOME/.kube/config`''`

示例：
```toml
WebEnable = true
ListeningPort = '0.0.0.0:1024'
KubeEnable = true
KubeconfigPath = '/Users/horonlee/Downloads/k8s/config'
```

## 网页端用法

1. `IP:8080` 是一个默认主页（什么都没有）
2. `IP:8080/json/*` 返回各种 json 信息
   1. `IP:8080/json/docker` docker
   2. `IP:8080/json/kube` kubernetes
3. `IP:8080/search?image=xxxxx` 返回指定镜像对应在运行的容器