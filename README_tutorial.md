# 需求
- 脑子
- 云服务器（1核心1G内存起步）
- 域名（解析国内云服务器IP需备案）
- PostgreSQL
- golang
- Nginx
- 宝塔面板（建议）
- clash/v2ray/ssr（可选）
- Docker (推荐)

## 准备工作
- 以下请在root用户下执行，如不是root用户，请执行`sudo su`切换到root用户

1.更新系统环境
- centos
```
yum update -y
```
- Ubuntu & debian
```
apt update && apt upgrade -y
```
2. 安装必要组件
- centos
```
yum install git -y
```
- Ubuntu & debian
```
apt install git -y
```
3.（可选）安装并设置 [宝塔操作面板](https://www.bt.cn/bbs/thread-79460-1-1.html)
- 配置账户密码以及安全入口即可
## 正式开始
### 拉取源码
```
git clone https://github.com/JasonKhew96/biliroaming-go-server.git
```
### 修改[config.example.yml](config.example.yml)文件名为 config.yml ，根据内容按需设置

###编译二进制文件

- 下载golang
```
wget -O go.tar.gz https://go.dev/dl/go1.18.linux-amd64.tar.gz
tar -zxvf go.tar.gz -C /usr/local/bin/
```
- 配置临时golang环境
```
export GOROOT=/usr/local/bin/go
export GOPATH=/home/golang
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOROOT/bin
export PATH=$PATH:$GOPATH/bin
```
- 进入源码目录使用golang进行编译
```
cd biliroaming-go-server
go build
```
###安装并启用 [PostgreSQL](https://www.postgresql.org/download/)

- 根据系统安装对应的PostgreSQL
- 更改数据库默认用户的密码
