# hostker-ddns

Hostker DDNS 方便用户使用 DDNS 到 Hostker 域名解析服务

## 使用

### 安装
现在可以基于 shell 脚本安装依赖了，支持 Mac / Linux；Windows 有 exe，不确定是否可用
```sh
curl -o /tmp/install.sh -L https://github.com/csvwolf/hostker-ddns/raw/master/build.sh && \
    chmod +x /tmp/install.sh && \
    /tmp/install.sh
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

