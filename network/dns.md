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
目前，世界上一共有十三组根域名服务器，从`A.ROOT-SERVERS.NET`一直到`M.ROOT-SERVERS.NET`。

### 分级查询的实例

> `dig`命令的`+trace`参数可以显示DNS的整个分级查询过程

上面命令的第一段列出根域名`.`的所有NS记录，即所有根域名服务器。

根据内置的根域名服务器，DNS服务器向所有这些IP地址发出查询请求，询问顶级域名服务器的NS记录，最先回复的根域名服务器将被缓存，以后只向这台服务器发请求。
然后，DNS服务器向这些顶级域名服务器发出查询请求，询问次级域名的NS记录。
最后，DNS服务器向NS服务器查询主机名。

### NS记录的查询

> `dig`命令可以单独查看每一级域名的NS记录。


```
$ dig ns com


; <<>> DiG 9.11.3-1ubuntu1.8-Ubuntu <<>> ns com
;; global options: +cmd
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 39851
;; flags: qr rd ra; QUERY: 1, ANSWER: 13, AUTHORITY: 0, ADDITIONAL: 1

;; OPT PSEUDOSECTION:
; EDNS: version: 0, flags:; udp: 65494
;; QUESTION SECTION:
;com.				IN	NS

;; ANSWER SECTION:
com.			6025	IN	NS	c.gtld-servers.net.
com.			6025	IN	NS	b.gtld-servers.net.
com.			6025	IN	NS	h.gtld-servers.net.
com.			6025	IN	NS	d.gtld-servers.net.
com.			6025	IN	NS	g.gtld-servers.net.
com.			6025	IN	NS	m.gtld-servers.net.
com.			6025	IN	NS	f.gtld-servers.net.
com.			6025	IN	NS	i.gtld-servers.net.
com.			6025	IN	NS	k.gtld-servers.net.
com.			6025	IN	NS	l.gtld-servers.net.
com.			6025	IN	NS	j.gtld-servers.net.
com.			6025	IN	NS	a.gtld-servers.net.
com.			6025	IN	NS	e.gtld-servers.net.

;; Query time: 1790 msec
;; SERVER: 127.0.0.53#53(127.0.0.53)
;; WHEN: Tue Aug 06 09:11:05 CST 2019
;; MSG SIZE  rcvd: 256


```

`+short`参数可以简化显示的结果
```
$ dig +short ns com

e.gtld-servers.net.
a.gtld-servers.net.
j.gtld-servers.net.
l.gtld-servers.net.
k.gtld-servers.net.
i.gtld-servers.net.
f.gtld-servers.net.
m.gtld-servers.net.
g.gtld-servers.net.
d.gtld-servers.net.
h.gtld-servers.net.
b.gtld-servers.net.
c.gtld-servers.net.
```

### DNS的记录类型

> 域名与IP之间的对应关系，成为`记录`(record)，根据使用场景，`记录`可以分为不同的类型(type)，前面已经看到了有`A`记录和`NS`记录。

常见的DNS记录类型如下:
```
- 1. `A`: 地址记录(Address),返回域名指向的IP地址。
- 2. `NS`: 域名服务器记录(Name Server),返回保存下一级域名信息的服务器地址。该记录只能设置为域名，不能设置为IP地址。
- 3. `MX`: 邮件记录(Mail eXchange), 返回接收电子邮件的服务器地址。
- 4. `CNAME`: 规范名称记录(Canonical Name), 返回另一个域名，即当前查询的域名是另一个域名的跳转，详见下文。
- 5. `PTR`: 逆向查询记录(Pointer Record),只用于从IP地址查询域名，详见下文。
```

一般来说，为了服务的安全可靠，至少应该有两条`NS`记录，而`A`记录和`MX`记录可以有多条，这样就提供了服务的冗余性，防止出现单点失败。

`CNAME`记录主要用于域名的内部跳转，为服务器配置提供灵活性，用户感知不到。举例来说，`www.baidu.com`这个域名就是一个`CNAME`记录。
```
; <<>> DiG 9.11.3-1ubuntu1.8-Ubuntu <<>> www.baidu.com
;; global options: +cmd

...

;; ANSWER SECTION:
www.baidu.com.		1150	IN	CNAME	www.a.shifen.com.
www.a.shifen.com.	249	IN	A	183.232.231.174
www.a.shifen.com.	249	IN	A	183.232.231.172

...
```

上面结果显示，`www.baidu.com`的CNAME记录指向`www.a.shifen.com`。也就是说，用户查询`www.baidu.com`的时候，实际上返回的是`www.a.shifen.com`的IP地址。
这样的好处是，变更服务器IP地址的时候，只要修改`www.a.shifen.com`这个域名就可以了。用户的`www.baidu.com`域名不用修改。

由于`CNAME`记录就是一个替换，所以域名一旦设置`CNAME`记录以后，就不能再设置其他记录了(比如`A`记录和`MX`记录)，这是为了防止冲突。由于顶级域名通常要设置`MX`记录，
所以一般不允许用户对顶级域名设置`CNAME`记录。

`PTR`记录一般用于从IP地址反查域名，`dig`命令的`-x`参数用于查询`PTR`记录。
逆向查询的一个应用，是可以防止垃圾邮件，即验证发送邮件的IP地址，是否真的有它所声称的域名。
`dig`命令可以查看指定的记录类型。
```
$ dig a github.com
$ dig ns github.com
$ dig mx github.com
```

### 其他DNS工具

> 除了dig外，还有一些其他小工具也可以使用

- 1. host命令
> `host`命令可以看做`dig`命令的简化版本，返回当前请求域名的各种记录。
```
$ host github.com

github.com has address 13.229.188.59
github.com mail is handled by 1 ASPMX.L.GOOGLE.com.
github.com mail is handled by 5 ALT2.ASPMX.L.GOOGLE.com.
github.com mail is handled by 10 ALT3.ASPMX.L.GOOGLE.com.
github.com mail is handled by 10 ALT4.ASPMX.L.GOOGLE.com.
github.com mail is handled by 5 ALT1.ASPMX.L.GOOGLE.com.

$ host facebook.github.com

facebook.github.com is an alias for github.github.io.
github.github.io has address 185.199.111.153
github.github.io has address 185.199.109.153
github.github.io has address 185.199.110.153
github.github.io has address 185.199.108.153
```

`host`命令也可以用于逆向查询，即从IP地址查询域名，等同于`dig -x <ip>`

```
$ host 13.229.188.59

59.188.229.13.in-addr.arpa domain name pointer ec2-13-229-188-59.ap-southeast-1.compute.amazonaws.com.

```

- 2. nslookup命令
> `nslookup`命令用于互动式地查询域名记录。

```
$ nslookup

> facebook.github.io
Server:		127.0.0.53
Address:	127.0.0.53#53

Non-authoritative answer:
Name:	facebook.github.io
Address: 185.199.109.153
Name:	facebook.github.io
Address: 185.199.110.153
>
```

- 3. whois命令
> `whois`命令用来查看域名的注册情况

```
$ whois github.com
```