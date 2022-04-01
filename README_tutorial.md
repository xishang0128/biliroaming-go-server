# 需求
- 脑子
- 云服务器（1核心1G内存起步）
- 域名（解析国内云服务器IP需备案）
- PostgreSQL
- 宝塔面板（建议）
- golang（可选）
- clash/v2ray/ssr（可选）
- Nginx (可选)
- Docker (推荐)

## 准备工作
1.更新系统环境
- centos
```
yum update -y
```
- Ubuntu & debian
```
apt update _y && apt upgrade -y
```
2.安装并设置 [宝塔操作面板](https://www.bt.cn/bbs/thread-79460-1-1.html)
- 配置账户密码即可
3.安装并启用 [PostgreSQL](https://www.postgresql.org/download/)
- 根据系统安装对应的PostgreSQL
- 更改数据库默认用户的密码
