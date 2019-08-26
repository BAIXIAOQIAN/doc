[TOC]

- 参考文章: https://www.jianshu.com/p/14835303b07e

## redis的路由查询

- 将请求发送到任意节点，接收到请求的节点会将查询请求发送到正确的节点上执行。
- 开源方案
    - Redis-cluster
- 基本原理

![](./image/redis_cluster.png)