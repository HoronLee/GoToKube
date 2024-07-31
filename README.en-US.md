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

## Key Features:
- [x] View Docker and Kubernetes information via the console
- [x] Supports multiple databases (SQLite, MySQL)
- [x] Manipulate Kubernetes cluster resources using YAML files via various requests
- [x] Query, create, and delete Docker containers

## Usage

⚠️: Requires Docker Client API Version >= 1.45
The environment variable `JWT_SECRET_KEY=JWT_TOKEN` must be set, otherwise the web service will not work properly.
Most of the software's features are provided by the API. The best way is to check the API documentation: https://documenter.getpostman.com/view/34220703/2sA3e5d86S

## Building

### Build using make
1. Enter the project directory and open Makefile
2. Edit variables 1-6 to your needs. Generally, you only need to change GOOS (your system) and GOARCH (system architecture)
3. Execute `make` in the current directory to generate the binary file
4. Grant executable permissions: `sudo chmod +x GoToKube`

### Build using Go
1. Enter the project directory and execute `go build`
2. Obtain the `GoToKube` binary file and grant executable permissions: `sudo chmod +x GoToKube`

> Build using Docker
1. Build using the Dockerfile in the project: `docker build -t gotokube:dev .`
2. It is recommended to start the container using Docker Compose: `docker-compose up -d`
    1. The Docker sock file must be mapped to the container, otherwise the software cannot be started
   ```yaml
   volumes:
     - /var/run/docker.sock:/var/run/docker.sock
   ```

## Configuration File
> The configuration file will be generated in the same directory as the program after the first run, and can be modified later. Case is **insensitive**

```toml
[auth]
pass = 'gotokube'  # Not used yet

[database]
addr = ''  # Database address
name = ''  # Database name
password = ''  # Database password
path = 'data.db'  # Sqlite Database file path, defaults to data.db in the current directory
type = 'sqlite'  # Database type, defaults to sqlite, currently supports sqlite and mysql
user = ''  # Database username

[kube]
configpath = '' # Kubernetes configuration file path, defaults to $HOME/.kube/config
enable = true  # Whether to enable Kubernetes function

[log]
dir = ''  # Log file storage path

[web]
enable = true  # Whether to enable web function
listeningaddr = ':8080' # Web service listening address
```

## Environment Variables
> Similar to the configuration file, use `configuration unit`_`configuration item` = `configuration value` to set environment variables (variable values must be uppercase), which will override the values in the configuration file

Example:
- LOG_DIR='/var/log/gotokub' Log file storage path
- WEB_LISTENINGADDR=":9090" Web service listening addressw