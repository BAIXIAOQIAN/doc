## COAP

### 一、什么是COAP协议？

    是一种应用层协议，运行于UDP协议之上，COAP协议非常小巧，最小的数据包仅为4字节。
    基于REST架构，传输层为UDP，网络层为6LowPAN(面向低功耗无线局域网的Ipv6)

### 二、COAP协议怎么用？

    有四种消息类型，采用与HTTP相同的请求响应工作模式。

```
    - CON: 需要被确认的请求，如果CON请求被发送，那么对方必须做出回应

    - NON: 不需要被确认的请求

    - ACK: 应答消息，接收到CON消息的响应

    - RST: 复位消息
```

    消息的格式 = 固定长度的头部Header + 可选个数的Option + 负载Payload

### 三、简单的发布/订阅模型

    COAP协议通过扩展简单的实现了发布/订阅模型

```
    - Subject(主题): 代表COAP服务器上的资源，该资源状态随时可能发生变化

    - Observer(观察者): 代表对某个COAP资源感兴趣的客户端 COAP Client

    - Rejistation(登记): 观察者需要向服务器 COAP Server 登记感兴趣的 SubJect

    - Notification(通知): 当COAP服务器观察到某个Subject发生变化时，会主动向对该Subject感兴趣的观察者列表里的每个观察者发送其订阅的Subject最新状态数据
```

    观察协议在COAP协议的基础上增加了一个 Observer Option，通过该Option来实现订阅/发布模型管理。