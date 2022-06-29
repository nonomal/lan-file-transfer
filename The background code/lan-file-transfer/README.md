# lan-file-transfer
局域网传输文件

### 编译步骤：



1、将前端的dist 文件夹拷贝到后台项目目录中;

2、执行命令  ,将前端文件生成二进制文件;

```shell
go-bindata -o asset/asset.go dist/...
```

3、将 asset 里的包名改成asset （生成默认的时候是main）

4、编译后台go项目

```shell
go build .\main.go
```

