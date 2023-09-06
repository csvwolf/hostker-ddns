# hostker-ddns

Hostker DDNS 方便用户使用 DDNS 到 Hostker 域名解析服务

## 使用

### 安装
```sh
go get github.com/csvwolf/hostker-ddns
```

### 初始化
```sh
hostker-ddns init # 输入必要的信息
```

Token 请在：<https://console.hostker.net/index.html#/Account> 获取 API 密钥

### 测试配置
```sh
hostker-ddns run # 单次运行，可以在平台中确认
```

### 运行
```sh
hostker-ddns start # 会每分钟跑一次，可以配合 Dockerfile 使用
```

## DDNS 是个啥
家里的 IP 经常变，想绑个域名的话________，这么说你应该懂了。

