## protobuf协议

### proto文件生成go代码

#### 单个文件
```
protoc --go_out=plugins=grpc:. hello.proto
```

#### 多个文件
```
protoc --go_out=plugins=grpc:. *.proto
```

#### 多个文件并且依赖其他proto文件
```
protoc --go_out=plugins=grpc:. P2pTrans.proto --proto_path=. --proto_path=../common/
```