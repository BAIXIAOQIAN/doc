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

