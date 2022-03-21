ORM:
ORM  对象关系映射（英语：Object Relational Mapping，简称ORM，或O/RM，或O/R mapping），是一种程序设计技术，用于实现面向对象编程语言里不同类型系统的数据之间的转换。从效果上说，它其实是创建了一个可在编程语言里使用的“虚拟对象数据库”。如今已有很多免费和付费的ORM产品，而有些程序员更倾向于创建自己的ORM工具。

crud:
crud是指在做计算处理时的增加(Create)、检索(Retrieve)、更新(Update)和删除(Delete)几个单词的首字母简写。crud主要被用在描述软件系统中数据库或者持久层的基本操作功能。

//对 object 操作的四个方法 Read / Insert / Update / Delete

RPC:
RPC是远程过程调用（Remote Procedure Call）的缩写形式。SAP系统RPC调用的原理其实很简单，有一些类似于三层构架的C/S系统，第三方的客户程序通过接口调用SAP内部的标准或自定义函数，获得函数返回的数据进行处理后显示或打印。

FAQ:
FAQ一般指常见问题解答。 常见问题解答（FAQ，frequently-asked questions）是使新用户熟悉规则的一种方法。

CORS:
CORS是一个W3C标准，全称是"跨域资源共享"（Cross-origin resource sharing）。

RMI:
重量级框架: RMI（即Remote Method Invoke 远程方法调用）。在Java中，只要一个类extends了java.rmi.Remote接口，即可成为存在于服务器端的远程对象，供客户端访问并提供一定的服务。JavaDoc描述：Remote 接口用于标识其方法可以从非本地虚拟机上调用的接口。任何远程对象都必须直接或间接实现此接口。只有在“远程接口”（扩展 java.rmi.Remote 的接口）中指定的这些方法才可远程使用。


ESB: 
比较重量级
企业服务总线，即ESB全称为Enterprise Service Bus，指的是传统中间件技术与XML、Web服务等技术结合的产物。ESB提供了网络中最基本的连接中枢，是构筑企业神经系统的必要元素。
面向服务的体系结构已经逐渐成为IT集成的主流技术。面向服务的体系结构(service-oriented architecture，SOA)是一种软件系统设计方法，通过已经发布的和可发现的接口为终端用户应用程序或其它服务提供服务。

SOA:
面向服务架构（SOA）是一个组件模型，它将应用程序的不同功能单元（称为服务）进行拆分，并通过这些服务之间定义良好的接口和协议联系起来。接口是采用中立的方式进行定义的，它应该独立于实现服务的硬件平台、操作系统和编程语言。这使得构建在各种各样的系统中的服务可以以一种统一和通用的方式进行交互。


gRPC (gRPC Remote Procedure Calls) 是 Google 发起的一个开源远程过程调用系统，该系统基于 HTTP/2 协议传输，本文介绍 gRPC 的基础概念，首先通过关系图直观展示这些基础概念之间关联，介绍异步 gRPC 的 Server 和 Client 的逻辑；然后介绍 RPC 的类型，阅读和抓包分析 gRPC 的通信过程协议，gRPC 上下文；最后分析 grpc.pb.h 文件的内容，包括 Stub 的能力、Service 的种类以及与核心库的关系。
