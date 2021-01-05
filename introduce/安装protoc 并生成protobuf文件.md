### 1.安装 protoc （或者github相关仓库https://github.com/protocolbuffers/protobuf/releases）
`yum -y install protobuf-compiler`

### 2.安装相关套件
`go get -u github.com/golang/protobuf/{proto,protoc-gen-go}`
`go get github.com/micro/micro/v2/cmd/protoc-gen-micro@master`

### 3.相关文件夹下执行命令行
`/usr/local/protobuf/bin/protoc -I ./ --go_out=./ --micro_out=./ product.proto`