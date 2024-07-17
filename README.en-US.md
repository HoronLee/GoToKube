<h1 align="center">
<a href="https://blog.horonlee.com">GoToKube</a>
</h1>

<p align="center">
🐳 Easily manage Kubernetes clusters (Beta).
</p>

<pre align="center">
Make Kubernetes usage more convenient and faster
🧪 In development
</pre>

- [English](./README.en-US.md) | **简体中文**

## Main Features:
- [x] View Docker and Kubernetes information through the console

- [x] Multi-database support (SQLite, MySQL)

- [x] Manipulate resources within the Kubernetes cluster using YAML files through various requests

- [x] Query, create, and delete Docker containers

  ⚠️: Requires Docker Client API Version >= 1.45

## Build Methods

### Build with Make

1. Enter the project directory and open the Makefile
2. Edit the variables in lines 1-6 to your desired values; generally, you only need to change GOOS (your system) and GOARCH (system architecture)
3. Run `make` in the current directory to generate the binary file
4. Grant executable permission with `sudo chmod +x GoToKube`
5. Run `./GoToKube` to start the program

### Build with Go

1. Enter the project directory and run `go build`
2. Get the `GoToKube` binary file and grant executable permission with `sudo chmod +x GoToKube`
3. Run `./GoToKube` to start the program

> Build with Docker

1. Use the Dockerfile in the project to build with `docker build -t gotokube:dev .`
2. It is recommended to start the container using DockerCompose with `docker-compose up -d`
   1. The Docker sock file must be mapped into the container; otherwise, the software cannot start
      ```yml
      volumes:
        - /var/run/docker.sock:/var/run/docker.sock
      ```

## Configuration File

> The configuration file will be generated in the same directory as the program after the first run and can be modified later

- `WebEnable = true&false` Whether to automatically enable the web function after starting the program
- `ListeningAddr = '0.0.0.0:8080'` Listening address and port for the web function
- `TermEnable = true&false` Whether to enable the interactive terminal
- `KubeEnable = true&false` Whether to automatically enable the Kubernetes function after starting the program
- `KubeConfigPath = '.kube/config file path'` Configuration file path for the Kubernetes function
  - If not specified, the default is `$HOME/.kube/config`
- `DBType = 'sqlite&mysql'` Database type, default is sqlite, currently only supports sqlite and mysql
- `DBPath = 'data.db'` Database file path, default is `data.db` in the current directory of the program
- `DBAddr = '127.0.0.1:3306'` Database address
- `DBUser = 'root'` Database username
- `DBPass = 'password'` Database password
- `DBName = 'test'` Database name

Example:

```toml
WebEnable = true
ListeningPort = '127.0.0.1:1024'
KubeEnable = true
KubeConfigPath = '/Users/horonlee/Downloads/k8s/config'
```

## Web Usage

Most of the software's functions are provided by the API. The best way is to check the API documentation: https://documenter.getpostman.com/view/34220703/2sA3e5d86S

## Environment Variables
- LOG_DIR Log file storage path `/var/log/vdcontroller`

## Startup Parameters

Support configuration of software settings through startup parameters, such as: `./GoToKube -kubeconfig="/home/user/document/k8s/config"`

- `-KubeConfig` Kubernetes configuration file path