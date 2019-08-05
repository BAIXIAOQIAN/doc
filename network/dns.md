## DNS是什么
> DNS(Domain Name System的缩写)的作用非常简单,就是根据域名查出IP地址。

### 查询过程

虽然只需要返回一个IP地址，但是DNS的查询过程非常复杂，分成多个步骤。

工具软件`dig`可以显示整个查询过程

```
$ dig www.baidu.com
```

上面的命令会输出四段信息

```
; <<>> DiG 9.11.3-1ubuntu1.8-Ubuntu <<>> www.baidu.com
;; global options: +cmd
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 5746
;; flags: qr rd ra; QUERY: 1, ANSWER: 3, AUTHORITY: 0, ADDITIONAL: 1

;; OPT PSEUDOSECTION:
; EDNS: version: 0, flags:; udp: 65494
;; QUESTION SECTION:
;www.baidu.com.			IN	A

;; ANSWER SECTION:
www.baidu.com.		1077	IN	CNAME	www.a.shifen.com.
www.a.shifen.com.	176	IN	A	183.232.231.174
www.a.shifen.com.	176	IN	A	183.232.231.172

;; Query time: 33 msec
;; SERVER: 127.0.0.53#53(127.0.0.53)
;; WHEN: Mon Aug 05 19:50:08 CST 2019
;; MSG SIZE  rcvd: 101

```

第一段是查询参数和统计
```
; <<>> DiG 9.11.3-1ubuntu1.8-Ubuntu <<>> www.baidu.com
;; global options: +cmd
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 5746
;; flags: qr rd ra; QUERY: 1, ANSWER: 3, AUTHORITY: 0, ADDITIONAL: 1
```

第二段是查询内容
```
;; QUESTION SECTION:
;www.baidu.com.			IN	A
```
上面的结果表明，查询域名`www.baidu.com`的`A`记录，`A`是address的缩写

第三段是DNS服务器的答复
```
;; ANSWER SECTION:
www.baidu.com.		1077	IN	CNAME	www.a.shifen.com.
www.a.shifen.com.	176	IN	A	183.232.231.174
www.a.shifen.com.	176	IN	A	183.232.231.172
```
上面的结果表明`www.baidu.com`的别名CNAME`www.a.shifen.com`,CNAME通常称别名指向
`www.a.shifen.com`有两个`A`记录，即两个IP地址,`176`是TTL值，表示缓存时间，即176秒之内不用重新查询

第四段是DNS服务器的一些传输信息
```
;; Query time: 33 msec
;; SERVER: 127.0.0.53#53(127.0.0.53)
;; WHEN: Mon Aug 05 19:50:08 CST 2019
;; MSG SIZE  rcvd: 101
```

### DNS服务器

我们根据上面的例子，一步步还原，本机到底怎么得到域名`www.baidu.com`的IP地址
首先，本机一定要知道DNS服务器的IP地址，否则上不了网。通过DNS服务器，才能知道某个域名的IP地址到底是什么。

```
DNS服务器的IP地址，有可能是动态的，每次上网时由网关分配，这叫做DHCP机制；也有可能是事先指定的固定地址。Linux系统中，DNS服务器的IP地址保存在/etc/resolv.conf文件。
本机只向自己的DNS服务器查询，`dig`命令有一个`@`参数，显示向其他DNS服务器查询的结果。
```

### 域名的层级

DNS服务器怎么会知道每个域名的IP地址呢，答案是分级查询。
```
;; ANSWER SECTION:
www.baidu.com.		1077	IN	CNAME	www.a.shifen.com.
www.a.shifen.com.	176	IN	A	183.232.231.174
www.a.shifen.com.	176	IN	A	183.232.231.172
```

我们可以看到，每个域名的尾部都多了一个点，`www.baidu.com`显示为`www.baidu.com.`。这不是疏忽，而是所有域名的尾部，实际上都有一个根域名。
举例来说，`www.example.com`真正的域名是`www.example.com.root`,简写为`www.example.com.`。因为根域名`.root`对于所有域名都是一样的，所以平时是省略的。

根域名的下一级，叫做`顶级域名`(top-level domain,缩写为TLD),比如`.com`,`.net`;再下一级域名叫做`次级域名`(second-level domain,缩写为SLD)，比如`www.exmaple.com`里面的`.example`,这一级
域名是用户可以注册的；再下一级是主机名`host`,比如`www.example.com`里面的`www`，又称为`三级域名`,这是用户在自己的域里面为服务器分配的名称，是用户可以任意分配的。

总结一下，域名的层级结构如下：
```
主机名.次级域名.顶级域名.根域名

即: host.sld.tld.root
```

### 根域名服务器

> DNS服务器根据域名的层级，进行分级查询。

需要明确的是，每一级域名都有自己的NS(name server)记录，NS记录指向该级域名的域名服务器。这些服务器知道下一级域名的各种记录。所谓`分级查询`，就是从根域名开始，依次查询每一级域名的NS记录，直到查到最终
的IP地址，过程大致如下:
```
- 1. 从`根域名服务器`查到`顶级域名服务器`的NS记录和A记录(IP地址)
- 2. 从`顶级域名服务器`查到`次级域名服务器`的NS记录和A记录(IP地址)
- 3. 从`次级域名服务器`查出`主机名`的IP地址
```

上面的过程中，并没有提到DNS服务器是怎么知道`根域名服务器`的IP地址，这是因为`根域名服务器`的NS记录和IP地址一般是不会变化的，所以内置在DNS服务器里面。
