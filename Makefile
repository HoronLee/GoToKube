# 定义变量
BINARY_NAME := GoToKube
SRC_DIR := .
VERSION := $(shell git describe --tags --abbrev=0 --match 'v*')
COMMIT := $(shell git rev-parse --short HEAD)
EXTERNAL_VERSION ?= $(VERSION)
GOOS ?= windows
GOARCH ?= amd64
OUTPUT_DIR := bin
OUTPUT_NAME := $(BINARY_NAME)-$(EXTERNAL_VERSION)-$(GOOS)-$(GOARCH)

# 如果是 Windows 平台，添加 .exe 后缀
ifeq ($(GOOS), windows)
    OUTPUT_NAME := $(OUTPUT_NAME).exe
endif

# 默认目标
.PHONY: all
all: build

# 构建二进制文件
.PHONY: build
build:
	CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(OUTPUT_DIR)/$(OUTPUT_NAME) $(SRC_DIR)

all:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o $(OUTPUT_DIR)/$(BINARY_NAME)-$(EXTERNAL_VERSION)-windows-amd64.exe $(SRC_DIR)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(OUTPUT_DIR)/$(BINARY_NAME)-$(EXTERNAL_VERSION)-linux-amd64 $(SRC_DIR)
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o $(OUTPUT_DIR)/$(BINARY_NAME)-$(EXTERNAL_VERSION)-linux-arm64 $(SRC_DIR)
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o $(OUTPUT_DIR)/$(BINARY_NAME)-$(EXTERNAL_VERSION)-darwin-amd64 $(SRC_DIR)
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o $(OUTPUT_DIR)/$(BINARY_NAME)-$(EXTERNAL_VERSION)-darwin-arm64 $(SRC_DIR)

# 清理构建文件
.PHONY: clean
clean:
	rm -f $(OUTPUT_DIR)/$(BINARY_NAME)-*

# 运行二进制文件 (仅适用于当前平台构建)
.PHONY: run
run: build
	./$(OUTPUT_DIR)/$(OUTPUT_NAME)

# 安装依赖
.PHONY: deps
deps:
	go mod tidy

# 开发构建
.PHONY: dev
dev:
	go build -o $(OUTPUT_DIR)/$(BINARY_NAME)-dev
	./$(OUTPUT_DIR)/$(BINARY_NAME)-dev