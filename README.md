# ginblog
gin + vue 博客学习

## 创建项目库并本地拉取

## 添加README

## 初始化项目目录

1. 拉取项目之后，配置GO SDK
2. 需要配置go的proxy的就去配置
3. 初始化mod
   ```bash
   go mod init ginblog
   ```
4. 安装go框架gin
   文档地址：https://gin-gonic.com/zh-cn/docs/quickstart/
    ```bash
   go get -u github.com/gin-gonic/gin 
   ```
5. 添加主入口文件`main.go`
6. 配置工程目录
    ```text
    .
    ├── LICENSE
    ├── README.md
    ├── api
    │   └── v1
    ├── config
    ├── go.mod
    ├── go.sum
    ├── main.go
    ├── middleware
    ├── model
    ├── routes
    ├── upload
    ├── utils
    └── web
    
    ```

7. 添加配置文件
   安装ini配置包
   ```bash
   go get gopkg.in/ini.v1
   ```
   文档地址：https://ini.unknwon.io/docs
