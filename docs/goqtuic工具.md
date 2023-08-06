# goqtuic工具
ui 转 golang 工具

## 安装工具
```shell
go get -u -v github.com/stephenlyu/goqtuic

goqtuic -h
```

## 将xx.ui文件转化成golang文件并放到目录uitogo下
```shell
goqtuic -go-ui-dir="uitogo" -ui-file=xx.ui 
```