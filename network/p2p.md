## P2P穿透

### NAT简述
```
NAT(Network Address Translation), 网络地址转换，是一种广泛应用的解决IP短缺的有效方法，NAT将内网地址和端口号转换成合法的公网地址和端口号，建立一个会话，与公网主机进行通信。
```

### NAT的实现方式

```
NAT的实现方式有三种，静态转换、动态转换、端口多路复用

- 1.静态转换: 静态地址转换是将内部私网地址与合法公网地址进行一对一的转换，且每个内部地址的转换都是确定的。

- 2.动态转换: 是指将内部网络的私有IP转换为公网IP时，IP地址是不确定的，动态地址的转换是从合法地址池中动态选择一个未使用的地址来对内部私网地址进行转换。

- 3.端口多路复用: 是指改变外出数据包的源端口并进行端口转换，即端口地址转换，内部网络的所有主机均可共享一个合法外部IP地址实现对Internet的访问。
```

### NAT的类型

```
考虑到UDP的无状态特性，目前对NAT的实现大致可分为Full Cone、Restricted Cone 、Port Restricted Cone和Symmetric四种。值得指出的是，对于TCP协议而言，目前NAT中针对TCP的实现基本上是一致的，
这是因为TCP协议本身便是面向连接的。
```

**名词定义**
```
1.内部Tuple: 是指内部主机的私有地址和端口号所构成的二元组，即内部主机所发送报文的源地址、端口所构成的二元组。

2.外部Tuple: 是指内部Tuple经过NAT的源地址/端口转换之后，所获得的外部地址、端口所构成的二元组，即外部主机收到经过NAT转换之后的报文时，它所看到的该报文的源地址和源端口。

3.目标Tuple: 是指外部主机的地址、端口所构成的二元组，即内部主机所发送报文的目标地址、端口所构成的二元组。
```

**详细释义**
```
- 1. Full Cone NAT
    所有来自同一个内部Tuple X的请求均被NAT转换至同一个外部Tuple Y, 而不管这些请求是不是属于同一个应用或者是多个应用。除此之外，当X-Y的转换关系建立之后，任意外部主机均可随时将Y中的地址和端口作为目标地址和端口，
    向内部主机发送UDP报文，由于对外部请求的来源无任何限制，因此这种方式虽然足够简单，但却不那么安全。

- 2. Restricted Cone NAT
    它是Full Cone的受限版本，所有来自同一个内部Tuple X的请求均被NAT转换至同一个外部Tuple Y,这与Full Cone相同，但不同的是，只有当内部主机曾经发送过报文给外部主机后，外部主机才能以Y中的信息作为目标地址和端口，
    向内部主机发送UDP请求报文，这意味着，NAT设备只向内转发那些来自于当前已知的外部主机的UDP报文，从而保障了外部请求来源的安全性。

- 3. Port Restricted Cone NAT
    它是 Restricted Cone NAT 的进一步受限版，只有当内部主机曾经发送过报文给外部主机(假设其IP地址为Z且端口为P)之后，外部主机才能以Y中的信息作为目标地址和端口，向内部主机发送UDP报文，同时，其请求报文的源端口必须为P,
    这一要求进一步强化了对外部报文请求来源的限制，从而较Restrictd Cone更具安全性。

- 4. Symmetric NAT
    这是一种比所有Cone NAT都要更为灵活的转换方式，在Cone NAT中，内部主机的内部Tuple与外部Tuple的转换映射关系是独立于内部主机所发出的UDP报文中的目标地址及端口的，即与目标Tuple无关；在Symmetric NAT中，目标Tuple则成为了
    NAT设备建立转换关系的一个重要考量，只有来自于同一个内部Tuple、且针对同一目标Tuple的请求才会被NAT转换至同一个外部Tuple，否则的话，NAT将为之分配一个新的外部Tuple；举个例子：当内部主机以相同的内部Tuple对两个不同的目标
    Tuple发送UDP报文时，此时NAT将会为内部主机分配两个不同的外部Tuple,并且建立起两个不同的内、外部Tuple转换关系。与此同时，只有接收到了内部主机所发送的数据包的外部主机才能向内部主机返回UDP报文，这里对外部返回报文来源的限制是
    与Port Restricted Cone NAT一致的。
```

### P2P的NAT研究

####　第一部分：不同NAT实现方法的介绍

```
- 1.完全圆锥型NAT(Full Cone NAT): 内网主机建立一个UDP socket(LocallIP:LocallPort)第一次使用这个socket给外部主机发送数据时NAT会给其分配一个公网(PublicIP:PublicPort),以后用这个socket向外面任何主机发送数据都将使用
    这对(PublicIP:PublicPort)。此外，任何外部主机只要知道这个(PublicIP:PublicPort)就可以发送数据给(PublicIP:PublicPort),内网的主机就能收到这个数据包。

- 2.地址限制圆锥型NAT(Address Restricted Cone NAT): 内网主机建立一个UDP socket(LocallIP:LocallPort)第一次使用这个socket给外部主机发送数据时NAT会给其分配一个公网(PublicIP:PublicPort),以后用这个socket向外面任何主机发送数据都将使用
    这对(PublicIP:PublicPort)。此外，任何外部主机想要发送数据给这个内网主机，只要知道这个(PublicIP:PublicPort)并且内网主机之前用这个socket曾向这个外部主机IP发送过数据。只要满足这两个条件，这个外部主机就可以用自己的(IP,任何端口)发送数据给
    (PublicIP:PublicPort),内网的主机就能收到这个数据包。

- 3.端口限制圆锥型NAT(Port Restricted Cone NAT): 建立socket同上,此外，任何外部主机想要发送数据给这个内网主机，只要知道这个(PublicIP:PublicPort)并且内网主机之前用这个socket曾向这个外部主机(IP,Port)发送过数据。只要满足这两个条件，这个外部主
　　　机就可以用自己的(IP,Port)发送数据给(PublicIP:PublicPort),内网的主机就能收到这个数据包。

- 4.对称型NAT(Symmetric NAT):内网主机建立一个UDP socket(LocallIP:LocallPort)，当用这个socket第一次发送数据给外部主机1时，NAT为其映射一个(PublicIP-1,Port-1),以后内网主机发送给外部主机1的所有数据都是用这个(PublicIP-1,Port-1),
    如果内网主机同时用这个socket给外部主机2发送数据，第一次发送时，NAT会为其分配一个(PublicIp-2,Port-2),以后内网主机发送给外部主机2的所有数据都是用这个(PublicIP-2,Port-2),如果NAT有多于一个公网IP,则PublicIP-1和PubilcIP-2可能不同，如果
    NAT只有一个公网IP,则Port-1和Port-2肯定不同，也就是说一定不能是(PublicIP-1 == PublicIP-2) && (Port-1 == Port-2)。此外，如果任何外部主机想要发送数据给这个内网主机，那么它首先应该收到内网主机发送给他的数据，然后才能往回发送，否则即使它知道
    内网主机的一个(PublicIP,Port)也不能发送数据给内网主机，这种NAT无法实现UDP-P2P通信。
```

#### 第二部分：NAT类型检测

> 前提条件：有一个公网的Server并且绑定了两个公网IP(IP-1,IP-2)。这个Server做UDP监听(IP-1,Port-1),(IP-2,Port-2)并根据客户端的要求进行应答。

##### 第一步： 检测客户端是否有能力进行UDP通信以及客户端是否位于NAT后
```
客户端建立UDP socket 然后用这个socket 向服务器的(IP-1,Port-1)发送数据包要求服务器返回客户端的IP和Port,客户端发送请求后立即开始接收数据包，要设定socket Timeout(300ms),防止无限堵塞，重复这个过程若干次。如果每次都超时，
无法接收到服务器的回应，则说明客户端无法进行UDP通信，可能是防火墙或者NAT阻止UDP通信，这样的客户端也就不能P2P了(检测停止)。

当客户端能够接收到服务器的回应时，需要把服务器返回的客户端(IP,Port)和这个客户端socket的(LocallIP,LocalPort)比较。如果完全相同则客户端不在NAT后，这样的客户端具有公网IP可以直接监听UDP端口接收数据进行通信(检测停止)。否则客户端在
NAT后要做进一步的**NAT类型**检测(继续)。
```

##### 第二步： 检测客户端NAT是否是Full Cone NAT
```
客户端建立UDP socket然后用这个socket向服务器的(IP-1,Port-1)发送数据包要求服务器用另一对(IP-2,Port-2)响应客户端的请求，往回发送一个数据包，客户端发送请求后立即开始接收数据包，要设定socket Timeout(300ms), 防止无限堵塞，
重复这个过程若干次。如果每次都超时，无法接收到服务器的回应，则说明客户端的NAT不是一个Full Cone NAT,具体类型有待下一步检测(继续)。如果能够接收到服务器从(IP-2,Port-2)返回的应答UDP包，则说明客户端是一个Full Cone NAT，这样的
客户端能够进行UDP-P2P通信(检测停止)。
```

##### 第三步： 检测客户端NAT是否是Symmetric NAT
```
客户端建立UDP socket 然后用这个socket向服务器的(IP-1,Port-1)发送数据包要求服务器返回客户端的IP和Port,客户端发送请求后立即开始接收数据包，要设定socket Timeout(300ms),防止无限堵塞，重复这个过程直到收到回应(一定能够收到，因为第一步保证了
这个客户端可以进行UDP通信)。

用同样的方法用一个socket向服务器的(IP-2,Port-2)发送数据包要求服务器返回客户端的IP和Port。

比较上面两个过程从服务器返回的客户端(IP,Port),如果两个过程返回的(IP,Port)有一对不同则说明客户端为Symmetric NAT,这样的客户端无法进行UDP-P2P通信(检测停止)。否则继续检测。
```

##### 第四步： 检测客户端NAT是Restricted Cone NAT还是Port Restricted Cone NAT
```
客户端建立UDP socket然后用这个socket向服务器的(IP-1,Port-1)发送数据包要求服务器用IP-1和一个不同于Port-1的端口发送一个UDP数据包响应客户端，客户端发送请求后立即开始接收数据包，要设定socket Timeout(300ms),防止无限堵塞，重复这个过程若干次。
如果每次都超时，无法接收到服务器的回应，则说明客户端是一个Port Restricted Cone NAT,如果能够收到服务器的响应则说明客户端是一个Restricted Cone NAT。以上两种NAT都可以进行UDP-P2P通信。
```

***注意：*** 以上检测过程中只说明了是否可以进行UDP-P2P的打洞通信，具体怎么通信一般要借助于Rendezvous Server。另外对于Symmetric NAT不是说完全不能进行UDP-P2P打洞通信，可以进行端口预测打洞，不过不能保证成功。
