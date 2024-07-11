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
- [x] You can view docker information through the console.
- [x] The console will terminate the application if it detects a Docker exception.
- [x] Displaying information through the web interface
- [x] Connect to Kubernetes cluster and show all pods through console.
- [x] Multi-database support(SQLite MySQL)
- [x] The YAML file is used to manage resources within a Kubernetes cluster through various requests.

## How to build

> Required Docker Client API Version >= 1.45

1. Go to the project directory and execute `go build`. 2.
2. Get the `VDController` binary and give it executable permissions `sudo chmod +x VDController`. 3.
3. Put `VDController` into a separate folder and put it into the project's webSrc folder.
4. Execute `. /VDController` to start the application

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

**Docker Operations**

   > The URL prefix is $IP/docker, followed by the addresses below.

- `/search?ctr=$ImageName` View all Docker containers created with the specified image name.

**Kubernetes Operations**

   > The URL prefix is $IP/kube, followed by the addresses below.

- GET `/deployments/$Namespace` Get all Deployments in the specified namespace.
- GET `/deployment/$Namespace/$DeployName` Get detailed information about the specified Deployment in the specified namespace.
- GET `/services/$Namespace` Get all Services in the specified namespace.
- GET `/pods/$Namespace` Get all Pods in the specified namespace.
    - `/pod/$Namespace/$PodName` Get detailed information about the specified Pod.
- GET `/namespaces` Get all namespaces.
- POST `/uploadYaml` to upload a YAML file
  - Usage:
    ```bash
    curl -X POST http://127.0.0.1:1024/kube/uploadYaml \
    -F "file=@/Users/horonlee/code/kubernetes/nginx.yaml" \
    -H "Content-Type: multipart/form-data"
    ```
- GET `/listYaml` to get all uploaded YAML files
- DELETE `/deleteYaml/$YamlName` to delete a YAML file

## Environment variable

- LOG_DIR Path to the log file `/var/log/vdcontroller`.

## Startup parameters

> Support to configure software settings via startup parameters, e.g.: `./VDController -kubeconfig="/home/user/document/k8s/config"

- `-kubeconfig` Kubernetes configuration file path