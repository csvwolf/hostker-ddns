#!/bin/bash

# 检测操作系统类型
unameOut="$(uname -s)"
case "${unameOut}" in
    Linux*)     machine=Linux;;
    Darwin*)    machine=Mac;;
    CYGWIN*)    machine=Cygwin;;
    MINGW*)     machine=MinGw;;
    *)          machine="UNKNOWN:${unameOut}"
esac

# 检测机器架构
machineArch="$(uname -m)"
case "${machineArch}" in
    armv6l)   arch=arm64;;
    armv7l)   arch=arm64;;
    aarch64)  arch=arm64;;
    x86)      arch=x86_64;;
    x86_64)   arch=x86_64;;
    i386)     arch=i386;;
    *)        arch="UNKNOWN:${machineArch}"
esac

# 根据操作系统和架构选择对应的下载链接
fileName=""
case "${machine}" in
    Linux)
        case "${arch}" in
            arm64)     fileName="hostker-ddns_Linux_arm64.tar.gz";;
            x86_64)    fileName="hostker-ddns_Linux_x86_64.tar.gz";;
            i386)      fileName="hostker-ddns_Linux_i386.tar.gz";;
            *)       echo "不支持的架构：${arch}" && exit 1;;
        esac
        ;;
    Mac)
        case "${arch}" in
            arm64)     fileName="hostker-ddns_Darwin_arm64.tar.gz";;
            x86_64)    fileName="hostker-ddns_Darwin_x86_64.tar.gz";;
            *)       echo "不支持的架构：${arch}" && exit 1;;
        esac
        ;;
    *)
        echo "不支持的操作系统：${machine}" && exit 1;;
esac

# 下载安装包
echo "正在下载安装包...${fileName}"
mkdir -p tmp
curl -sL "https://github.com/csvwolf/hostker-ddns/releases/latest/download/${fileName}"  | tar -xzf - -C ./tmp
chmod +x tmp/hostker-ddns  # 赋予可执行权限
mv ./tmp/hostker-ddns /usr/local/bin/
rm -rf tmp
# 完成安装
echo "安装完成！"