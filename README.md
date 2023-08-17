#  轻量级网盘项目 自己学习练手go 项目学习使用


## 1.使用Gin + Xorm +Go-redis


### Xorm官网
[Xorm](https://xorm.io/zh)
### Go-redis
[Go-redis](https://redis.uptrace.dev/zh)


## 开放端口
```txt
开放80端口
firewall-cmd --zone=public --add-port=80/tcp --permanent
重启防火墙
systemctl restart firewalld.service

```

## 修改环境
```txt
go env -w GOOS=windows
go env GOOS

```














