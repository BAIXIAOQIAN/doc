## P2P穿透

### NAT简述

NAT(Network Address Translation), 网络地址转换，是一种广泛应用的解决IP短缺的有效方法，NAT将内网地址和端口号转换成合法的公网地址和端口号，建立一个会话，与公网主机进行通信。

### NAT的实现方式

NAT的实现方式有三种，静态转换、动态转换、端口多路复用

- 1.静态转换: 静态地址转换是将内部私网地址与合法公网地址进行一对一的转换，且每个内部地址的转换都是确定的。

- 2.动态转换: 是指将内部网络的私有IP转换为公网IP时，IP地址是不确定的，动态地址的转换是从合法地址池中动态选择一个未使用的地址来对内部私网地址进行转换。

- 3.端口多路复用: 是指改变外出数据包的源端口并进行端口转换，即端口地址转换，内部网络的所有主机均可共享一个合法外部IP地址实现对Internet的访问。

### NAT的类型

考虑到UDP的无状态特性，目前对NAT的实现大致可分为Full Cone、Restricted Cone 、Port Restricted Cone和Symmetric四种。值得指出的是，对于TCP协议而言，目前NAT中针对TCP的实现基本上是一致的，
这是因为TCP协议本身便是面向连接的。

**名词定义**
1.内部Tuple: 是指内部主机的私有地址和端口号所构成的二元组，即内部主机所发送报文的源地址、端口所构成的二元组。

2.外部Tuple: 是指内部Tuple经过NAT的源地址/端口转换之后，所获得的外部地址、端口所构成的二元组，即外部主机收到经过NAT转换之后的报文时，它所看到的该报文的源地址和源端口。

3.目标Tuple: 是指外部主机的地址、端口所构成的二元组，即内部主机所发送报文的目标地址、端口所构成的二元组。

**详细释义**
```
- 1. Full Cone NAT
    所有来自同一个内部Tuple X的请求均被NAT转换至同一个外部Tuple Y, 而不管这些请求是不是
```