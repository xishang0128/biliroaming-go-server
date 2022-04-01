# 需求
- 脑子
- 云服务器（1核心1G内存起步）
- 域名（解析国内云服务器IP需备案）
- golang
- PostgreSQL
- 宝塔面板（建议）
- clash/v2ray/ssr（可选）
- Nginx (可选)
- Docker (推荐)

## 准备工作
1.更新系统环境

centos
```
yum update -y
```
Ubuntu & debian
```
apt update _y && apt upgrade -y
```
2.安装必须的环境
centos
```
yum install epel-release -y
yum install git golang -y
```
Ubuntu & debian
```
apt install git golang -y
```
