<h1 align="center">
<a href="https://blog.horonlee.com">GoToKube</a>
</h1>

<p align="center">
🐳 Manage Kubernetes Cluster with Ease (Beta).
</p>

<pre align="center">
Makes it easier and faster to use Kubernetes
🧪 developing
</pre>

- **English** | [简体中文](./README.md)

## Key Features:
- [x] Ability to view Docker and Kubernetes information via the console
- [x] Multiple database support (SQLite MySQL).
- [x] Manipulate resources within a kubernetes cluster through various requests using yaml files.
- [x] Querying, creating and deleting Docker containers

## How to build

> Required Docker Client API Version >= 1.45

1. Go to the project directory and execute `go build`. 2.
2. Get the `VDController` binary and give it executable permissions `sudo chmod +x VDController`. 3.
3. Execute `. /VDController` to start the application

## Configuration files

> The configuration file is created in the same directory as the application after the first run and can be changed later.

- `WebEnable = true&false` Whether to automatically enable the web function after starting the program
- `ListeningPort = '0.0.0.0:8080'` The listening address and port for the web function
- `KubeEnable = true&false` Whether to automatically enable the Kubernetes function after starting the program
- `KubeconfigPath = '.kube/config file path'` The configuration file path for the Kubernetes function
    - If not specified, the default will be `$HOME/.kube/config`
- `DBType = 'sqlite&mysql'` Database type, defaults to sqlite. Currently, only sqlite and mysql are supported
- `DBPath = 'data.db'` Database file path, defaults to `data.db` in the current directory of the program
- `DBAddr = '127.0.0.1:3306'` Database address
- `DBUser = 'root'` Database username
- `DBPass = 'password'` Database password
- `DBName = 'test'` Database name

Example:

```toml
WebEnable = true
ListeningPort = '127.0.0.1:1024'
KubeEnable = true
KubeconfigPath = '/Users/horonlee/Downloads/k8s/config'
```

## Web usage

Most of the software's features are provided by the API, the best way to get started is to check out the API documentation.: https://documenter.getpostman.com/view/34220703/2sA3e5d86S

## Environment variable

- LOG_DIR Path to the log file `/var/log/vdcontroller`.

## Startup parameters

> Support to configure software settings via startup parameters, e.g.: `./VDController -kubeconfig="/home/user/document/k8s/config"

- `-kubeconfig` Kubernetes configuration file path