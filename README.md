<h1 align="center">
<a href="https://blog.horonlee.com">Visual Docker</a>
</h1>

<p align="center">
🐳 Manage Containers with Ease (Beta).
</p>

<pre align="center">
🧪 developing
</pre>

- **English** | [简体中文](./README.zh-CN.md)

> Makes it easier and faster to use docker
> Required Docker Client API Version >= 1.45

Current progress:
1. You can view docker information through the console.
2. The console will terminate the application if it detects a Docker exception.
3. Displaying information through the web interface
4. Connect to Kubernetes cluster and show all pods through console.
5. soon...

## How to build

1. Go to the project directory and execute `go build`. 2.
2. Get the `VDController` binary and give it executable permissions `sudo chmod +x VDController`. 3.
3. Put `VDController` into a separate folder and put it into the project's webSrc folder.
4. Execute `. /VDController` to start the application

## Configuration files

> The configuration file is created in the same directory as the application after the first run and can be changed later.

- `WebEnable = true&false` Whether or not to enable the web function automatically after starting the application.
- `ListeningPort = '0.0.0.0:8080'` Listening address and port for the web feature.
- `KubeEnable = true&false` Whether or not to enable kubernetes automatically when the application is started.
- `KubeconfigPath = '.kube/config file path'` Path to the configuration file for the kubernetes feature.
    - If this field is not filled in, `$HOME/.kube/config` will be used by default.

Example:
```toml
WebEnable = true
ListeningPort = '127.0.0.1:1024'
KubeEnable = true
KubeconfigPath = '/Users/horonlee/Downloads/k8s/config'
```

## Web usage

1. `IP:8080` is a default homepage (nothing)
2. `IP:8080/json/*` returns a variety of json information.
   1. `IP:8080/json/docker` docker
   2. `IP:8080/json/kube` kubernetes
3. `IP:8080/search?image=$IMAGE_NAME` Returns the running container for the specified image.

## Environment variable

- LOG_DIR Path to the log file `/var/log/vdcontroller`.

## Startup parameters

> Support to configure software settings via startup parameters, e.g.: `./VDController -kubeconfig="/home/user/document/k8s/config"

- `-kubeconfig` Kubernetes configuration file path