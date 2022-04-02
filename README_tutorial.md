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

- centos请不要选择centos8，不受PostgreSQL支持

1.更新系统环境
- centos7
```
yum update -y
```
- Ubuntu & debian
```
apt update && apt upgrade -y
```
2. 安装必要组件
- centos7
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
### 编译二进制文件

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
### 修改[config.example.yml](config.example.yml)文件名为 config.yml ，根据内容按需设置
- 可用宝塔面板直接重命名并编辑

### 安装并启用 [PostgreSQL](https://www.postgresql.org/download/)

1.根据系统安装对应的PostgreSQL
- centos7
```
sudo yum install -y https://download.postgresql.org/pub/repos/yum/reporpms/EL-7-x86_64/pgdg-redhat-repo-latest.noarch.rpm
sudo yum install -y postgresql14-server
sudo /usr/pgsql-14/bin/postgresql-14-setup initdb
sudo systemctl enable postgresql-14
sudo systemctl start postgresql-14
```
- Ubuntu & debian
```
sudo sh -c 'echo "deb http://apt.postgresql.org/pub/repos/apt $(lsb_release -cs)-pgdg main" > /etc/apt/sources.list.d/pgdg.list'
wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc | sudo apt-key add -
sudo apt-get update
sudo apt-get -y install postgresql
```
2.更改数据库默认用户的密码
- 切换到数据库用户
```
sudo -u postgres psql
```
- 修改密码，引号内是需要修改的密码，可根据喜好修改，命令最后有分号
```
ALTER USER postgres WITH PASSWORD 'postgres';
```
- 退出PostgreSQL
```
\q
```
3.修改linux系统postgres用户的密码
- 删除用户postgres的密码
```
sudo  passwd -d postgres
```
- 设置用户postgres的密码
```
sudo -u postgres passwd
```
- 此时系统会提示输入新的密码，输入两次密码并回车即可
