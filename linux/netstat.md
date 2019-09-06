## linux netstat

### 统计80端口连接数
```
netstat -nat|grep -i "80"|wc -l
```

### 统计httpd协议连接数
```
ps -ef|grep httpd|wc -l
```

### 统计已连接上的，状态为"established"
```
netstat -na|grep ESTABLISHED|wc -l
```

### 查出哪个ip地址连接最多，将其封了
```
netstat -na|grep ESTABLISHED|awk {print $5}|awk -F: {print $1}|sort|uniq -c|sort -r +0n

netstat -na|grep SYN|awk {print $5}|awk -F: {print $1}|sort|uniq -c|sort -r +0n
```

### 查看apache当前并发访问数
```
netstat -an | grep ESTABLISHED | wc -l
```

### 查看有多少个进程数
```
ps aux|grep httpd|wc -l
```

### 显示所有活动的网络连接
```
netstat -na
```

### 查看同时连接到哪个服务器IP比较多
```
netstat -an|awk  '{print $4}'|sort|uniq -c|sort -nr|head
```

### 查看哪些IP连接到服务器连接多，可以查看连接异常IP
```
netstat -an|awk -F: '{print $2}'|sort|uniq -c|sort -nr|head
```

### 列出所有连接过的IP地址
```
netstat -n -p | grep SYN_REC | sort -u
```

### 列出所有发送SYN_REC连接节点的IP地址
```
netstat -n -p | grep SYN_REC | awk '{print $5}' | awk -F: '{print $1}'
```

### 列出所有连接到本机的UDP或者TCP连接的IP数量
```
netstat -anp |grep 'tcp|udp' | awk '{print $5}' | cut -d: -f1 | sort | uniq -c | sort -n
```