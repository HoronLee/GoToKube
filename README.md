# VisualDocker

> 让你更加方便快捷的使用docker

目前进度：
1. 可以通过控制台查看docker的信息
2. 控制台检测到Docker异常会终止程序
3. 通过 Web 界面展示信息
4. soon...

## 使用方法

1. 进入项目目录执行`go build`
2. 得到`VDController`二进制文件，给予可执行权限`sudo chmod +x VDController`
3. 将`VDController`放到独立文件夹
4. 执行`./VDController`即可开启程序

## 网页端用法

1. `IP:8080` 是一个默认主页（什么都没有）
2. `IP:8080/json` 返回 docker 版本
3. `IP:8080/search?image=xxxxx` 返回指定镜像对应在运行的容器