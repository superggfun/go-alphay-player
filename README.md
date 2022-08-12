# go-alphay-player
超仿Bilibili播放器Go版，开箱即用，免安装环境

# go-alphay-player
超防Bilibili弹幕播放器，开箱即用
![index.png](https://tva1.sinaimg.cn/large/008c6yxSgy1h537zlyqc2j31h60pvtx7.jpg)
![index2.png](https://tva1.sinaimg.cn/large/008c6yxSgy1h5380v99acj30rc0ir7ar.jpg)

## 优点
* 开箱即用
* 免安装运行环境
* 免配置弹幕数据库，内置数据库
* 高性能（使用Go开发）

## 配置文件config.json
```json
{
    "username": "user",     // 登录名
    "password": "123123",   // 登录密码
    "duration": 600,        // 登录持续时间
    "admin": "/admin",      // 后台地址
    "player": "/",          // 播放器地址
    "domain": "localhost",  // 域名，必须填写。不然后台无法登录
    "port": 8001,           // 端口
    "tip": {
      "time": "6",
      "color": "#fb7299",
      "text": "请大家遵守弹幕礼仪，文明发送弹幕"
    },
    "Limit_time": 60,       // 限制弹幕发送时间
    "Limit_requence": 20,   // 限制发送弹幕次数
    "allow_url": [],        // 限定域名访问
    "paras": 0              // 是否开启解析，0为关闭，1为打开
}
```
## 部署
#### Nginx(推荐)
1. 解压
```bash
tar zxvf alphay-player1.0.2.tar.gz
```
2. 运行
```bash
nohup ./main &
exit
```
3. Nginx反向代理
* 使用vim修改nginx配置文件
```bash
vim usr/local/nginx/conf/nginx.conf
```
* 在http段中添加
```vim
server {
        listen       80;
        server_name  域名;

        location / {
                proxy_pass  http://localhost:8001;

                }
   }
```
* vim保存
```vim
Esc:wq
```
5. 如何停止播放器
```bash
lsof -i:8001
kill PID数字
```
#### Linux宝塔面板
在Linux宝塔面板中部署Go。
![baota.png](https://tva1.sinaimg.cn/large/008c6yxSgy1h54a3l9dowj31aq066q50.jpg)

#### Docker部署
```bash
docker build -t alphay-player .
docker run --name player -p 8001:8001 -v /player:/player -d alphay-player
```


