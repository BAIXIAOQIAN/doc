## MQTT协议

### MQTT协议是什么
> MQTT是一个客户端服务端结构的`发布/订阅`模式的消息传输协议。
- 设计思想: 轻巧、开放、简单、规范、易于实现
- 应用场景: 机器与机器的通信(M2M)、物联网环境(IOT)

### MQTT控制报文的结构

MQTT控制报文由三部分组成:

|字段名|描述|
|---|---|
|Fixed header | 固定报头，所有控制报文都包含 |
|Variable header | 可变报头，部分控制报文包含 |
|Payload | 有效载荷，部分控制报文包含 |

#### 固定报头(Fixed header)
> 每个MQTT控制报文都包含一个固定报头。

![](./image/mqtt_fixed_haeder.png)

##### MQTT控制报文的类型(MQTT Control Packet type)

- 位置: Byte1中bits 7-4 ,表示为4位无符号值

![](./image/mqtt_control_packet_type.png)

##### 标志(Flags)

- 位置: Byte1中bits 3-0

- 在不使用标识位的消息类型中，标识位被作为保留位。如果收到无效的标志时，接收端必须关闭网络连接。

![](./image/mqtt_flags.png)

```
- DUP: 发布消息的副本。用来在保证消息的可靠传输，如果设置为1，则在下面的变长中增加MessageId,并且需要回复确认，
以保证消息传输完成，但不能用于检测消息重复发送。

- Qos: 发布消息的服务质量，即: 保证消息传递的次数
     00: 至多一次，即: <= 1
     01: 至少一次，即: >= 1
     10: 只有一次，即: == 1
     11: 预留

- RETAIN: 发布保留标识，表示服务器要保留这次推送的信息，如果有新的订阅者出现，就把这消息推送给它，如果没有，
那么推送至当前订阅者后释放。
```
##### 剩余长度(Remaining Length)

- 位置: 从第二个字节(Byte2)开始

- 剩余长度表示当前报文剩余部分的字节数，包括可变报头和负载的数据。剩余长度不包括用于编码剩余长度字段本身的字节数。

- 剩余长度字段使用一个变长编码方案，对于小于128的值使用单字节编码；更大的值按下面的方式处理:低7位有效位用于编码数据，
最高有效位用于指示是否有更多的字节，且按照大端方式进行编码。因此每个字节可以编码128个数值和一个延续位(continuation bit),
剩余长度字段最大4个字节。

|字节数|最小值|最大值|
|---|---|---|
|1|0(0x00)|127(0x7F)|
|2|128(0x80,0x01)|16383(0xFF,0x7F)|
|3|16384(0x80,0x80,0x01)|2097151(0xFF,0xFF,0x7F)|
|4|2097152(0x80,0x80,0x80,0x01)|268435455(0xFF,0xFF,0xFF,0x7F)|

#### 可变报头
> 某些MQTT控制报文包含一个可变报头部分，它在固定报头和负载之间，可变报头的内容根据报文类型的不同而不同。可变报头的报文标识符(Packet Identifier)字段存在于多个类型的报文中。

##### 报文标识符(Packet Identifier)

|Bit|7-0|
|---|---|
|Byte1| 报文标识符MSB|
|Byte2| 报文标识符LSB|

![](./image/mqtt_packet.png)

```
很多类型数据包中都包括一个2字节的数据包标识字段，这些类型的包有: PUBLISH(Qos > 0)、PUBACK、PUBREC、PUBREL、PUBCOMP
、SUBSCRIBE、SUBACK、UNSUBSCRIBE、UNSUBACK

客户端和服务端彼此独立的分配报文标识符，因此，客户端服务端组合使用相同的报文标识符可以实现并发的消息交换。
```

#### 有效载荷(Payload)
> 某些MQTT控制报文在报文的最后部分包含一个有效载荷，对于PUBLISH来说有效载荷就是应用消息。

![](./image/mqtt_payload.png)

```
Payload消息体是MQTT数据包的第三部分，不好CONNECT、SUBSCRIBE、SUBACK、UNSUBSCRIBE四中类型的消息:

- CONNECT: 消息体内容主要是: 客户端的ClientID、订阅者的Topic、Message以及用户名和密码

- SUBSCRIBE: 消息体内容是一系列的要订阅的主题以及Qos

- SUBACK: 消息体内容是服务器对于SUBSCRIBE所申请的主题及Qos进行确认和回复

- UNSUBSCRIBE: 消息体内容是要订阅的主题
```

### MQTT控制报文

#### CONNECT-连接服务端
> 客户端到服务端的网络连接建立后，客户端发送给服务端的第一个报文必须是CONNECT报文，在一个网络连接上，客户端只能发送一次CONNECT报文。
> 服务端必须将客户端发送的第二个CONNECT报文当做协议违规处理并断开客户端的连接。
> 有效载荷包含一个或多个编码的字段。包括客户端的唯一标识符，Will主题，Will消息，用户名和密码。除了客户端标识之外，其他的字段都是可选的，
> 基于标志位来决定可变报头中是否需要包含这些字段

##### 固定报头(Fixed header)

- Byte1前四位(报文类型): 0001，后四位(Reserved保留位): 0000

- Byte2...剩余长度:等于可变报头的长度(10字节)加上有效载荷的长度。

##### 可变报头
```
CONNECT报文的可变报头按下列次序包含四个字段:协议名(Protocol Name),协议级别(Protocol Level),连接标志(Connect Flags)和保持连接(Keep Alive)
```

- 协议名(Protocol Name)
|协议名|说明|内容|
|---|---|---|
|byte1|长度MSB(0)|00000000|
|byte2|长度LSB(4)|00000100|
|byte3|'M'|01001101|
|byte4|'Q'|01010001|
|byte5|'T'|01010100|
|byte6|'T'|01010100|

```
协议名是表示`MQTT`的UTF-8编码的字符串。MQTT规范的后续版本不会改变这个字符串的偏移和长度。
如果协议名不正确，服务端可以断开客户端的连接，也可以按照某些其它规范继续处理CONNECT报文，
对于后一种情况，按照本规范，服务端不能继续处理CONNECT报文。
```
- 协议级别(Protocol Level)

|协议级别|说明|内容|
|---|---|---|
|byte7|Level(4)|00000100|

```
客户端用8位的无符号值表示协议的修订版本，对于3.1.1版本协议，协议级别字段的值是4(0x04).如果发现不支持的协议级别，服务端必须给客户端发送一个返回码0x01(不支持的协议级别)，
的CONNACK报文响应CONNECT报文，然后断开客户端的连接。
```

- 连接标志(Connect Flags)
> 连接标志字节包含一些用于指定MQTT连接行为的参数。它还指出有效载荷中的字段是否存在

![](./image/mqtt_connect_flag.png)

服务端必须验证CONNECT控制报文的保留标志位(第0位)是否为0，如果不为0必须断开客户端连接。

- 清理会话(Clean Session)

位置： 连接标志字节的第1位,这个二进制位指定了会话状态的处理方式。
```
客户端和服务端可以保存会话状态，以支持跨网络连接的可靠消息传输。这个标志位用于控制会话状态的生存时间。
如果清理会话（CleanSession）标志被设置为0，服务端必须基于当前会话（使用客户端标识符识别）的状态恢复与客户端的通信。
如果没有与这个客户端标识符关联的会话，服务端必须创建一个新的会话，当连接断开后，客户端和服务端必须保存会话信息。
当清理会话标志为0的会话连接断开之后，服务端必须将之后的QoS 1和QoS 2级别的消息保存为会话状态的一部分，如果这
些消息匹配断开连接时客户端的任何订阅。服务端也可以保存满足相同条件的QoS 0级别的消息。
```

