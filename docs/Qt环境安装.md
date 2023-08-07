# Qt工具文档

## 在GOPATH（全局安装）模式
```shell
export GO111MODULE=off  \
&& go get -v github.com/therecipe/qt/cmd/... \
&& go install -v -tags=no_env github.com/therecipe/qt/cmd/... \
&& $(go env GOPATH)/bin/qtsetup -test=false
```

## 在模块（每个项目）模式下（要解决 GFW 问题，您可以使用如下代理GOPROXY=https://goproxy.io：）
```shell
export GO111MODULE=on; go get -v github.com/therecipe/qt && go install -v -tags=no_env github.com/therecipe/qt/cmd/... && go mod vendor && rm -rf vendor/github.com/therecipe/env_linux_amd64_513 && git clone https://github.com/therecipe/env_linux_amd64_513.git vendor/github.com/therecipe/env_linux_amd64_513 && $(go env GOPATH)/bin/qtsetup
```

## qt编译
```shell
go mod vendor

GOWORK=off qtdeploy -fast test desktop
```

## qt编译2
```shell
go mod vendor

GOWORK=off qtdeploy -qt_dir=../vendor/github.com/therecipe/env_linux_amd64_513 -qt_version=5.13.0 build desktop main.go
```

## qt编译2
```shell
go mod vendor

git clone https://github.com/therecipe/env_linux_amd64_513.git vendor/github.com/therecipe/env_linux_amd64_513
GOWORK=off qtdeploy -qt_dir=../vendor/github.com/therecipe/env_linux_amd64_513 build desktop
```

## qt编译2
```shell
go mod vendor

GOWORK=off qtdeploy build desktop
```