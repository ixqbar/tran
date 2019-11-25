### 简繁转换


#### 依赖
```
go get -u github.com/liuzl/gocc
go get -u github.com/tidwall/redcon
go get -u github.com/urfave/cli
```

#### 安装
```
make
```

#### 配置
```
<?xml version="1.0" encoding="utf-8" ?>
<config>
    <address>0.0.0.0:8686</address>
    <data>/data/server/tran/data</data>
</config>
```

#### 启动服务
```
./bin/tran_linux --config=server.xml
```

#### 使用
```
[Will]# redis-cli -p 8686 --raw
127.0.0.1:8686> get 中华
中華
127.0.0.1:8686>
```

