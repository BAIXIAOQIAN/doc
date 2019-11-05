## tcpdump常用命令

[TOC]

### 默认启动
```
tcpdump
```

### 监视指定网络接口的数据包
```
tcpdump -i eth1
```

### 监视指定主机的数据包

- 打印所有进入或离开sundown的数据包
```
tcpdunp host sundown
```

- 也可以指定ip，例如截获所有210.27.48.1 的主机收到的和发出的所有的数据包
```
tcpdump host 210.27.48.1
```

- 打印helios 与 hot 或者与 ace 之间通信的数据包
```
tcpdump host helios and \( hot or ace \)
```

- 截获主机210.27.48.1 和主机210.27.48.2 或210.27.48.3的通信
```
tcpdump host 210.27.48.1 and \ (210.27.48.2 or 210.27.48.3 \) 
```

- 打印ace与任何其他主机之间通信的IP 数据包, 但不包括与helios之间的数据包.
```
tcpdump ip host ace and not helios
```

- 如果想要获取主机210.27.48.1除了和主机210.27.48.2之外所有主机通信的ip包，使用命令：
```
tcpdump ip host 210.27.48.1 and ! 210.27.48.2
```

- 截获主机hostname发送的所有数据
```
tcpdump -i eth0 src host hostname
```

- 监视所有送到主机hostname的数据包
```
tcpdump -i eth0 dst host hostname
```

### 监视指定主机和端口的数据包

- 如果想要获取主机210.27.48.1接收或发出的telnet包，使用如下命令
```
tcpdump tcp port 23 and host 210.27.48.1
```

- 对本机的udp 123 端口进行监视 123 为ntp的服务端口
```
tcpdump udp port 123 
```

### 监视指定网络的数据包

- 打印本地主机与Berkeley网络上的主机之间的所有通信数据包(nt: ucb-ether, 此处可理解为'Berkeley网络'的网络地址,此表达式最原始的含义可表达为: 打印网络地址为ucb-ether的所有数据包)
```
tcpdump net ucb-ether
```

- 打印所有通过网关snup的ftp数据包(注意, 表达式被单引号括起来了, 这可以防止shell对其中的括号进行错误解析)
```
tcpdump 'gateway snup and (port ftp or ftp-data)'
```

- 打印所有源地址或目标地址是本地主机的IP数据包
```
tcpdump ip and not net localnet
```

### 监视指定协议的数据包

- 打印TCP会话中的的开始和结束数据包, 并且数据包的源或目的不是本地网络上的主机.(nt: localnet, 实际使用时要真正替换成本地网络的名字))
```
tcpdump 'tcp[tcpflags] & (tcp-syn|tcp-fin) != 0 and not src and dst net localnet'
```

- 打印所有源或目的端口是80, 网络层协议为IPv4, 并且含有数据,而不是SYN,FIN以及ACK-only等不含数据的数据包.(ipv6的版本的表达式可做练习)
```
tcpdump 'tcp port 80 and (((ip[2:2] - ((ip[0]&0xf)<<2)) - ((tcp[12]&0xf0)>>2)) != 0)'
```

- 打印长度超过576字节, 并且网关地址是snup的IP数据包
```
tcpdump 'gateway snup and ip[2:2] > 576'
```

- 打印所有IP层广播或多播的数据包， 但不是物理以太网层的广播或多播数据报
```
tcpdump 'ether[0] & 1 = 0 and ip[16] >= 224'
```

- 打印除'echo request'或者'echo reply'类型以外的ICMP数据包( 比如,需要打印所有非ping 程序产生的数据包时可用到此表达式 .
(nt: 'echo reuqest' 与 'echo reply' 这两种类型的ICMP数据包通常由ping程序产生))
```
tcpdump 'icmp[icmptype] != icmp-echo and icmp[icmptype] != icmp-echoreply'
```

### tcpdump与wireshark

Wireshark(以前是ethereal)是Windows下非常简单易用的抓包工具。但在Linux下很难找到一个好用的图形化抓包工具。
还好有Tcpdump。我们可以用Tcpdump + Wireshark 的完美组合实现：在 Linux 里抓包，然后在Windows 里分析包。

```
tcpdump tcp -i eth1 -t -s 0 -c 100 and dst port ! 22 and src net 192.168.1.0/24 -w ./target.cap


(1)tcp: ip icmp arp rarp 和 tcp、udp、icmp这些选项等都要放到第一个参数的位置，用来过滤数据报的类型
(2)-i eth1 : 只抓经过接口eth1的包
(3)-t : 不显示时间戳
(4)-s 0 : 抓取数据包时默认抓取长度为68字节。加上-S 0 后可以抓到完整的数据包
(5)-c 100 : 只抓取100个数据包
(6)dst port ! 22 : 不抓取目标端口是22的数据包
(7)src net 192.168.1.0/24 : 数据包的源网络地址为192.168.1.0/24
(8)-w ./target.cap : 保存成cap文件，方便用ethereal(即wireshark)分析
```

### 使用tcpdump抓取HTTP包

```
tcpdump  -XvvennSs 0 -i eth0 tcp[20:2]=0x4745 or tcp[20:2]=0x4854
```

0x4745 为"GET"前两个字母"GE",0x4854 为"HTTP"前两个字母"HT"。


tcpdump 对截获的数据并没有进行彻底解码，数据包内的大部分内容是使用十六进制的形式直接打印输出的。显然这不利于分析网络故障，通常的解决办法是先使用带-w参数的tcpdump 截获数据并保存到文件中，然后再使用其他程序(如Wireshark)进行解码分析。当然也应该定义过滤规则，以避免捕获的数据包填满整个硬盘。