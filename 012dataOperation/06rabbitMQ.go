RabbitMQ介绍
1. RabbitMQ介绍
1.1.1. MQ简介
简单释义
消息总线(Message Queue)，是一种跨进程、异步的通信机制，用于上下游传递消息。由消息系统来确保消息的可靠传递。

背景描述
当前市面上mq的产品很多，比如RabbitMQ、Kafka、ActiveMQ、ZeroMQ和阿里巴巴捐献给Apache的RocketMQ。甚至连redis这种NoSQL都支持MQ的功能。

适用场景
上下游逻辑解耦&&物理解耦
保证数据最终一致性
广播
错峰流控等等
1.1.2. RabbitMQ的特点
RabbitMQ是由Erlang语言开发的AMQP的开源实现。

AMQP：Advanced Message Queue，高级消息队列协议。它是应用层协议的一个开放标准，为面向消息的中间件设计，基于此协议的客户端与消息中间件可传递消息，并不受产品、开发语言灯条件的限制。

可靠性(Reliablity)：使用了一些机制来保证可靠性，比如持久化、传输确认、发布确认。
灵活的路由(Flexible Routing)：在消息进入队列之前，通过Exchange来路由消息。对于典型的路由功能，Rabbit已经提供了一些内置的Exchange来实现。针对更复杂的路由功能，可以将多个Exchange绑定在一起，也通过插件机制实现自己的Exchange。
消息集群(Clustering)：多个RabbitMQ服务器可以组成一个集群，形成一个逻辑Broker。
高可用(Highly Avaliable Queues)：队列可以在集群中的机器上进行镜像，使得在部分节点出问题的情况下队列仍然可用。
多种协议(Multi-protocol)：支持多种消息队列协议，如STOMP、MQTT等。
多种语言客户端(Many Clients)：几乎支持所有常用语言，比如Java、.NET、Ruby等。
管理界面(Management UI)：提供了易用的用户界面，使得用户可以监控和管理消息Broker的许多方面。
跟踪机制(Tracing)：如果消息异常，RabbitMQ提供了消息的跟踪机制，使用者可以找出发生了什么。
插件机制(Plugin System)：提供了许多插件，来从多方面进行扩展，也可以编辑自己的插件。
1.1.3. rabbitmq简单使用


所有MQ产品从模型抽象来说，都是一样的过程：

消费者(consumer)订阅某个队列。
生产者(product)创建消息，然后发布到队列中(queue)，最终将消息发送到监听的消费者。
这只是最简单抽象的描述，具体到RabbitMQ则由更详细的概念需要解释。



Broker：标识消息队列服务器实体.
Virtual Host：虚拟主机。标识一批交换机、消息队列和相关对象。虚拟主机是共享相同的身份认证和加密环境的独立服务器域。每个vhost本质上就是一个mini版的RabbitMQ服务器，拥有自己的队列、交换器、绑定和权限机制。vhost是AMQP概念的基础，必须在链接时指定，RabbitMQ默认的vhost是 /。
Exchange：交换器，用来接收生产者发送的消息并将这些消息路由给服务器中的队列。
Queue：消息队列，用来保存消息直到发送给消费者。它是消息的容器，也是消息的终点。一个消息可投入一个或多个队列。消息一直在队列里面，等待消费者连接到这个队列将其取走。
Banding：绑定，用于消息队列和交换机之间的关联。一个绑定就是基于路由键将交换机和消息队列连接起来的路由规则，所以可以将交换器理解成一个由绑定构成的路由表。
Channel：信道，多路复用连接中的一条独立的双向数据流通道。新到是建立在真实的TCP连接内地虚拟链接，AMQP命令都是通过新到发出去的，不管是发布消息、订阅队列还是接收消息，这些动作都是通过信道完成。因为对于操作系统来说，建立和销毁TCP都是非常昂贵的开销，所以引入了信道的概念，以复用一条TCP连接。
Connection：网络连接，比如一个TCP连接。
Publisher：消息的生产者，也是一个向交换器发布消息的客户端应用程序。
Consumer：消息的消费者，表示一个从一个消息队列中取得消息的客户端应用程序。
Message：消息，消息是不具名的，它是由消息头和消息体组成。消息体是不透明的，而消息头则是由一系列的可选属性组成，这些属性包括routing-key(路由键)、priority(优先级)、delivery-mode(消息可能需要持久性存储[消息的路由模式])等。
1.1.4. RabbitMQ的六种工作模式
simple简单模式


消息产生着§将消息放入队列
消息的消费者(consumer) 监听(while) 消息队列,如果队列中有消息,就消费掉,消息被拿走后,自动从队列中删除(隐患 消息可能没有被消费者正确处理,已经从队列中消失了,造成消息的丢失)应用场景:聊天(中间有一个过度的服务器;p端,c端)
work工作模式(资源的竞争)


消息产生者将消息放入队列消费者可以有多个,消费者1,消费者2,同时监听同一个队列,消息被消费?C1 C2共同争抢当前的消息队列内容,谁先拿到谁负责消费消息(隐患,高并发情况下,默认会产生某一个消息被多个消费者共同使用,可以设置一个开关(syncronize,与同步锁的性能不一样) 保证一条消息只能被一个消费者使用)
应用场景:红包;大项目中的资源调度(任务分配系统不需知道哪一个任务执行系统在空闲,直接将任务扔到消息队列中,空闲的系统自动争抢)
publish/subscribe发布订阅(共享资源)


X代表交换机rabbitMQ内部组件,erlang 消息产生者是代码完成,代码的执行效率不高,消息产生者将消息放入交换机,交换机发布订阅把消息发送到所有消息队列中,对应消息队列的消费者拿到消息进行消费
相关场景:邮件群发,群聊天,广播(广告)
routing路由模式


消息生产者将消息发送给交换机按照路由判断,路由是字符串(info) 当前产生的消息携带路由字符(对象的方法),交换机根据路由的key,只能匹配上路由key对应的消息队列,对应的消费者才能消费消息;
根据业务功能定义路由字符串
从系统的代码逻辑中获取对应的功能字符串,将消息任务扔到对应的队列中业务场景:error 通知;EXCEPTION;错误通知的功能;传统意义的错误通知;客户通知;利用key路由,可以将程序中的错误封装成消息传入到消息队列中,开发者可以自定义消费者,实时接收错误;
topic 主题模式(路由模式的一种)


星号井号代表通配符
星号代表多个单词,井号代表一个单词
路由功能添加模糊匹配
消息产生者产生消息,把消息交给交换机
交换机根据key的规则模糊匹配到对应的队列,由队列的监听消费者接收消息消费
RPC (先不做解释后续补充)


RabbitMQ安装
1. RabbitMQ安装
1.1.1. win下安装
第一步：下载并安装erlang
原因：RabbitMQ服务端代码是使用并发式语言Erlang编写的，安装Rabbit MQ的前提是安装Erlang。
下载地址：http://www.erlang.org/downloads


根据本机位数选择erlang下载版本。

下载完是这么个东西：


双击，点next就可以。


选择一个自己想保存的地方，然后next、finish就可以。
 

安装完事儿后要记得配置一下系统的环境变量。 此电脑-->鼠标右键“属性”-->高级系统设置-->环境变量-->“新建”系统环境变量


变量名：ERLANG_HOME

变量值就是刚才erlang的安装地址，点击确定。

然后双击系统变量path
点击“新建”，将%ERLANG_HOME%\bin加入到path中。

最后windows键+R键，输入cmd，再输入erl，看到版本号就说明erlang安装成功了。


第二步：下载并安装RabbitMQ
下载地址：http://www.rabbitmq.com/download.html


双击下载后的.exe文件，安装过程与erlang的安装过程相同。
RabbitMQ安装好后接下来安装RabbitMQ-Plugins。打开命令行cd，输入RabbitMQ的sbin目录。
我的目录是：E:\RabbitMQ\rabbitmq_server-3.8.1\sbin

然后在后面输入rabbitmq-plugins enable rabbitmq_management命令进行安装



如果出现下面的提示表示运行成功



输入命令：rabbitmq-server.bat

如果出现下面的提示表示启动成功



rabbitmq启动成功，浏览器中http://localhost:15672，



输入guest,guest进入rabbitMQ管理控制台：



注意

上面的安装步骤如果出现下面错误

输入rabbitmq-plugins enable rabbitmq_management命令出现


这样的提示解决办法

解决方法：

将 C:\Users\Administrator\.erlang.cookie

同步至C:\Windows\System32\config\systemprofile\.erlang.cookie

同时删除：C:\Users\Administrator\AppData\Roaming\RabbitMQ目录

重新输入rabbitmq-plugins enable rabbitmq_management命令

运行命令：rabbitmq-server.bat
如果出现下面的提示：



这个是因为rabbit已经启动了，不能再次启动，通过tasklist指令，发现进程是存在的：

    tasklist | find /i "erl"


如果有结果，那么说明已经启动了，通过任务管理器kill掉进程再次启动即可。

ctrl+alt+delete进入任务管理器



右键->结束任务就好了

Simple模式
1. Simple模式


消息产生着§将消息放入队列
消息的消费者(consumer) 监听(while) 消息队列,如果队列中有消息,就消费掉,消息被拿走后,自动从队列中删除(隐患 消息可能没有被消费者正确处理,已经从队列中消失了,造成消息的丢失)应用场景:聊天(中间有一个过度的服务器;p端,c端)
做simple简单模式之前首先我们新建一个Virtual Host并且给他分配一个用户名，用来隔离数据，根据自己需要自行创建

第一步



第二步



第三步



第四步



第五步



1.1.1. 代码层面
目录结构



kuteng-RabbitMQ

-RabbitMQ

--rabitmq.go //这个是RabbitMQ的封装

-SimlpePublish

--mainSimlpePublish.go //Publish 先启动

-SimpleRecieve

--mainSimpleRecieve.go

rabitmq.go代码

package RabbitMQ

import (
    "fmt"
    "log"

    "github.com/streadway/amqp"
)

//连接信息amqp://kuteng:kuteng@127.0.0.1:5672/kuteng这个信息是固定不变的amqp://事固定参数后面两个是用户名密码ip地址端口号Virtual Host
const MQURL = "amqp://kuteng:kuteng@127.0.0.1:5672/kuteng"

//rabbitMQ结构体
type RabbitMQ struct {
    conn    *amqp.Connection
    channel *amqp.Channel
    //队列名称
    QueueName string
    //交换机名称
    Exchange string
    //bind Key 名称
    Key string
    //连接信息
    Mqurl string
}

//创建结构体实例
func NewRabbitMQ(queueName string, exchange string, key string) *RabbitMQ {
    return &RabbitMQ{QueueName: queueName, Exchange: exchange, Key: key, Mqurl: MQURL}
}

//断开channel 和 connection
func (r *RabbitMQ) Destory() {
    r.channel.Close()
    r.conn.Close()
}

//错误处理函数
func (r *RabbitMQ) failOnErr(err error, message string) {
    if err != nil {
        log.Fatalf("%s:%s", message, err)
        panic(fmt.Sprintf("%s:%s", message, err))
    }
}

//创建简单模式下RabbitMQ实例
func NewRabbitMQSimple(queueName string) *RabbitMQ {
    //创建RabbitMQ实例
    rabbitmq := NewRabbitMQ(queueName, "", "")
    var err error
    //获取connection
    rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
    rabbitmq.failOnErr(err, "failed to connect rabb"+
        "itmq!")
    //获取channel
    rabbitmq.channel, err = rabbitmq.conn.Channel()
    rabbitmq.failOnErr(err, "failed to open a channel")
    return rabbitmq
}

//直接模式队列生产
func (r *RabbitMQ) PublishSimple(message string) {
    //1.申请队列，如果队列不存在会自动创建，存在则跳过创建
    _, err := r.channel.QueueDeclare(
        r.QueueName,
        //是否持久化
        false,
        //是否自动删除
        false,
        //是否具有排他性
        false,
        //是否阻塞处理
        false,
        //额外的属性
        nil,
    )
    if err != nil {
        fmt.Println(err)
    }
    //调用channel 发送消息到队列中
    r.channel.Publish(
        r.Exchange,
        r.QueueName,
        //如果为true，根据自身exchange类型和routekey规则无法找到符合条件的队列会把消息返还给发送者
        false,
        //如果为true，当exchange发送消息到队列后发现队列上没有消费者，则会把消息返还给发送者
        false,
        amqp.Publishing{
            ContentType: "text/plain",
            Body:        []byte(message),
        })
}

//simple 模式下消费者
func (r *RabbitMQ) ConsumeSimple() {
    //1.申请队列，如果队列不存在会自动创建，存在则跳过创建
    q, err := r.channel.QueueDeclare(
        r.QueueName,
        //是否持久化
        false,
        //是否自动删除
        false,
        //是否具有排他性
        false,
        //是否阻塞处理
        false,
        //额外的属性
        nil,
    )
    if err != nil {
        fmt.Println(err)
    }

    //接收消息
    msgs, err := r.channel.Consume(
        q.Name, // queue
        //用来区分多个消费者
        "", // consumer
        //是否自动应答
        true, // auto-ack
        //是否独有
        false, // exclusive
        //设置为true，表示 不能将同一个Conenction中生产者发送的消息传递给这个Connection中 的消费者
        false, // no-local
        //列是否阻塞
        false, // no-wait
        nil,   // args
    )
    if err != nil {
        fmt.Println(err)
    }

    forever := make(chan bool)
    //启用协程处理消息
    go func() {
        for d := range msgs {
            //消息逻辑处理，可以自行设计逻辑
            log.Printf("Received a message: %s", d.Body)

        }
    }()

    log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
    <-forever

}
mainSimlpePublish.go代码

package main

import (
    "fmt"

    "github.com/student/kuteng-RabbitMQ/RabbitMQ"
)

func main() {
    rabbitmq := RabbitMQ.NewRabbitMQSimple("" +
        "kuteng")
    rabbitmq.PublishSimple("Hello kuteng222!")
    fmt.Println("发送成功！")
}
mainSimpleRecieve.go代码

package main

import "github.com/student/kuteng-RabbitMQ/RabbitMQ"

func main() {
    rabbitmq := RabbitMQ.NewRabbitMQSimple("" +
        "kuteng")
    rabbitmq.ConsumeSimple()
}



Work模式
1. Work模式（工作模式，一个消息只能被一个消费者获取）


消息产生者将消息放入队列消费者可以有多个,消费者1,消费者2,同时监听同一个队列,消息被消费?C1 C2共同争抢当前的消息队列内容,谁先拿到谁负责消费消息(隐患,高并发情况下,默认会产生某一个消息被多个消费者共同使用,可以设置一个开关(syncronize,与同步锁的性能不一样) 保证一条消息只能被一个消费者使用)
应用场景:红包;大项目中的资源调度(任务分配系统不需知道哪一个任务执行系统在空闲,直接将任务扔到消息队列中,空闲的系统自动争抢)
目录结构



kuteng-RabbitMQ

-RabbitMQ

--rabitmq.go //这个是RabbitMQ的封装和Simple模式代码一样

-SimlpePublish

--mainSimlpePublish.go //Publish 先启动

-SimpleRecieve1

--mainSimpleRecieve.go

-SimpleRecieve2

--mainSimpleRecieve.go

注意

Work模式和Simple模式相比代码并没有发生变化只是多了一个消费者

rabitmq.go代码

package RabbitMQ

import (
    "fmt"
    "log"

    "github.com/streadway/amqp"
)

//连接信息amqp://kuteng:kuteng@127.0.0.1:5672/kuteng这个信息是固定不变的amqp://事固定参数后面两个是用户名密码ip地址端口号Virtual Host
const MQURL = "amqp://kuteng:kuteng@127.0.0.1:5672/kuteng"

//rabbitMQ结构体
type RabbitMQ struct {
    conn    *amqp.Connection
    channel *amqp.Channel
    //队列名称
    QueueName string
    //交换机名称
    Exchange string
    //bind Key 名称
    Key string
    //连接信息
    Mqurl string
}

//创建结构体实例
func NewRabbitMQ(queueName string, exchange string, key string) *RabbitMQ {
    return &RabbitMQ{QueueName: queueName, Exchange: exchange, Key: key, Mqurl: MQURL}
}

//断开channel 和 connection
func (r *RabbitMQ) Destory() {
    r.channel.Close()
    r.conn.Close()
}

//错误处理函数
func (r *RabbitMQ) failOnErr(err error, message string) {
    if err != nil {
        log.Fatalf("%s:%s", message, err)
        panic(fmt.Sprintf("%s:%s", message, err))
    }
}

//创建简单模式下RabbitMQ实例
func NewRabbitMQSimple(queueName string) *RabbitMQ {
    //创建RabbitMQ实例
    rabbitmq := NewRabbitMQ(queueName, "", "")
    var err error
    //获取connection
    rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
    rabbitmq.failOnErr(err, "failed to connect rabb"+
        "itmq!")
    //获取channel
    rabbitmq.channel, err = rabbitmq.conn.Channel()
    rabbitmq.failOnErr(err, "failed to open a channel")
    return rabbitmq
}

//直接模式队列生产
func (r *RabbitMQ) PublishSimple(message string) {
    //1.申请队列，如果队列不存在会自动创建，存在则跳过创建
    _, err := r.channel.QueueDeclare(
        r.QueueName,
        //是否持久化
        false,
        //是否自动删除
        false,
        //是否具有排他性
        false,
        //是否阻塞处理
        false,
        //额外的属性
        nil,
    )
    if err != nil {
        fmt.Println(err)
    }
    //调用channel 发送消息到队列中
    r.channel.Publish(
        r.Exchange,
        r.QueueName,
        //如果为true，根据自身exchange类型和routekey规则无法找到符合条件的队列会把消息返还给发送者
        false,
        //如果为true，当exchange发送消息到队列后发现队列上没有消费者，则会把消息返还给发送者
        false,
        amqp.Publishing{
            ContentType: "text/plain",
            Body:        []byte(message),
        })
}

//simple 模式下消费者
func (r *RabbitMQ) ConsumeSimple() {
    //1.申请队列，如果队列不存在会自动创建，存在则跳过创建
    q, err := r.channel.QueueDeclare(
        r.QueueName,
        //是否持久化
        false,
        //是否自动删除
        false,
        //是否具有排他性
        false,
        //是否阻塞处理
        false,
        //额外的属性
        nil,
    )
    if err != nil {
        fmt.Println(err)
    }

    //接收消息
    msgs, err := r.channel.Consume(
        q.Name, // queue
        //用来区分多个消费者
        "", // consumer
        //是否自动应答
        true, // auto-ack
        //是否独有
        false, // exclusive
        //设置为true，表示 不能将同一个Conenction中生产者发送的消息传递给这个Connection中 的消费者
        false, // no-local
        //列是否阻塞
        false, // no-wait
        nil,   // args
    )
    if err != nil {
        fmt.Println(err)
    }

    forever := make(chan bool)
    //启用协程处理消息
    go func() {
        for d := range msgs {
            //消息逻辑处理，可以自行设计逻辑
            log.Printf("Received a message: %s", d.Body)

        }
    }()

    log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
    <-forever

}
mainSimlpePublish.go代码

package main

import (
    "fmt"
    "strconv"
    "time"

    "github.com/student/kuteng-RabbitMQ/RabbitMQ"
)

func main() {
    rabbitmq := RabbitMQ.NewRabbitMQSimple("" +
        "kuteng")

    for i := 0; i <= 100; i++ {
        rabbitmq.PublishSimple("Hello kuteng!" + strconv.Itoa(i))
        time.Sleep(1 * time.Second)
        fmt.Println(i)
    }
}
mainSimpleRecieve.go代码(两个消费端的代码是一样的)

package main

import "github.com/student/kuteng-RabbitMQ/RabbitMQ"

func main() {
    rabbitmq := RabbitMQ.NewRabbitMQSimple("" +
        "kuteng")
    rabbitmq.ConsumeSimple()
}


Publish模式
1. Publish模式（订阅模式，消息被路由投递给多个队列，一个消息被多个消费者获取）


X代表交换机rabbitMQ内部组件,erlang 消息产生者是代码完成,代码的执行效率不高,消息产生者将消息放入交换机,交换机发布订阅把消息发送到所有消息队列中,对应消息队列的消费者拿到消息进行消费
相关场景:邮件群发,群聊天,广播(广告)
目录结构



kuteng-RabbitMQ

-RabbitMQ

--rabitmq.go //这个是RabbitMQ的封装

-Pub

--mainPub.go //Publish 先启动

-Sub

--mainSub.go

-Sub2

--mainSub.go

rabitmq.go代码

package RabbitMQ

import (
    "fmt"
    "log"

    "github.com/streadway/amqp"
)

//连接信息
const MQURL = "amqp://kuteng:kuteng@127.0.0.1:5672/kuteng"

//rabbitMQ结构体
type RabbitMQ struct {
    conn    *amqp.Connection
    channel *amqp.Channel
    //队列名称
    QueueName string
    //交换机名称
    Exchange string
    //bind Key 名称
    Key string
    //连接信息
    Mqurl string
}

//创建结构体实例
func NewRabbitMQ(queueName string, exchange string, key string) *RabbitMQ {
    return &RabbitMQ{QueueName: queueName, Exchange: exchange, Key: key, Mqurl: MQURL}
}

//断开channel 和 connection
func (r *RabbitMQ) Destory() {
    r.channel.Close()
    r.conn.Close()
}

//错误处理函数
func (r *RabbitMQ) failOnErr(err error, message string) {
    if err != nil {
        log.Fatalf("%s:%s", message, err)
        panic(fmt.Sprintf("%s:%s", message, err))
    }
}

//订阅模式创建RabbitMQ实例
func NewRabbitMQPubSub(exchangeName string) *RabbitMQ {
    //创建RabbitMQ实例
    rabbitmq := NewRabbitMQ("", exchangeName, "")
    var err error
    //获取connection
    rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
    rabbitmq.failOnErr(err, "failed to connect rabbitmq!")
    //获取channel
    rabbitmq.channel, err = rabbitmq.conn.Channel()
    rabbitmq.failOnErr(err, "failed to open a channel")
    return rabbitmq
}

//订阅模式生产
func (r *RabbitMQ) PublishPub(message string) {
    //1.尝试创建交换机
    err := r.channel.ExchangeDeclare(
        r.Exchange,
        "fanout",
        true,
        false,
        //true表示这个exchange不可以被client用来推送消息，仅用来进行exchange和exchange之间的绑定
        false,
        false,
        nil,
    )

    r.failOnErr(err, "Failed to declare an excha"+
        "nge")

    //2.发送消息
    err = r.channel.Publish(
        r.Exchange,
        "",
        false,
        false,
        amqp.Publishing{
            ContentType: "text/plain",
            Body:        []byte(message),
        })
}

//订阅模式消费端代码
func (r *RabbitMQ) RecieveSub() {
    //1.试探性创建交换机
    err := r.channel.ExchangeDeclare(
        r.Exchange,
        //交换机类型
        "fanout",
        true,
        false,
        //YES表示这个exchange不可以被client用来推送消息，仅用来进行exchange和exchange之间的绑定
        false,
        false,
        nil,
    )
    r.failOnErr(err, "Failed to declare an exch"+
        "ange")
    //2.试探性创建队列，这里注意队列名称不要写
    q, err := r.channel.QueueDeclare(
        "", //随机生产队列名称
        false,
        false,
        true,
        false,
        nil,
    )
    r.failOnErr(err, "Failed to declare a queue")

    //绑定队列到 exchange 中
    err = r.channel.QueueBind(
        q.Name,
        //在pub/sub模式下，这里的key要为空
        "",
        r.Exchange,
        false,
        nil)

    //消费消息
    messges, err := r.channel.Consume(
        q.Name,
        "",
        true,
        false,
        false,
        false,
        nil,
    )

    forever := make(chan bool)

    go func() {
        for d := range messges {
            log.Printf("Received a message: %s", d.Body)
        }
    }()
    fmt.Println("退出请按 CTRL+C\n")
    <-forever
}
mainPub.go代码

package main

import (
    "fmt"
    "strconv"
    "time"

    "github.com/student/kuteng-RabbitMQ/RabbitMQ"
)

func main() {
    rabbitmq := RabbitMQ.NewRabbitMQPubSub("" +
        "newProduct")
    for i := 0; i < 100; i++ {
        rabbitmq.PublishPub("订阅模式生产第" +
            strconv.Itoa(i) + "条" + "数据")
        fmt.Println("订阅模式生产第" +
            strconv.Itoa(i) + "条" + "数据")
        time.Sleep(1 * time.Second)
    }

}
mainSub.go代码(两个消费者代码是一样的)

package main

import "github.com/student/kuteng-RabbitMQ/RabbitMQ"

func main() {
    rabbitmq := RabbitMQ.NewRabbitMQPubSub("" +
        "newProduct")
    rabbitmq.RecieveSub()
}
赏


Routing模式
1. Routing模式(路由模式，一个消息被多个消费者获取，并且消息的目标队列可被生产者指定)


消息生产者将消息发送给交换机按照路由判断,路由是字符串(info) 当前产生的消息携带路由字符(对象的方法),交换机根据路由的key,只能匹配上路由key对应的消息队列,对应的消费者才能消费消息;
根据业务功能定义路由字符串
从系统的代码逻辑中获取对应的功能字符串,将消息任务扔到对应的队列中业务场景:error 通知;EXCEPTION;错误通知的功能;传统意义的错误通知;客户通知;利用key路由,可以将程序中的错误封装成消息传入到消息队列中,开发者可以自定义消费者,实时接收错误;
目录结构



kuteng-RabbitMQ

-RabbitMQ

--rabitmq.go //这个是RabbitMQ的封装

-publish

--mainpublish.go //Publish 先启动

-recieve1

--mainrecieve.go

-recieve2

--mainrecieve.go

rabitmq.go代码

package RabbitMQ

import (
    "fmt"
    "log"

    "github.com/streadway/amqp"
)

//连接信息
const MQURL = "amqp://kuteng:kuteng@127.0.0.1:5672/kuteng"

//rabbitMQ结构体
type RabbitMQ struct {
    conn    *amqp.Connection
    channel *amqp.Channel
    //队列名称
    QueueName string
    //交换机名称
    Exchange string
    //bind Key 名称
    Key string
    //连接信息
    Mqurl string
}

//创建结构体实例
func NewRabbitMQ(queueName string, exchange string, key string) *RabbitMQ {
    return &RabbitMQ{QueueName: queueName, Exchange: exchange, Key: key, Mqurl: MQURL}
}

//断开channel 和 connection
func (r *RabbitMQ) Destory() {
    r.channel.Close()
    r.conn.Close()
}

//错误处理函数
func (r *RabbitMQ) failOnErr(err error, message string) {
    if err != nil {
        log.Fatalf("%s:%s", message, err)
        panic(fmt.Sprintf("%s:%s", message, err))
    }
}

//路由模式
//创建RabbitMQ实例
func NewRabbitMQRouting(exchangeName string, routingKey string) *RabbitMQ {
    //创建RabbitMQ实例
    rabbitmq := NewRabbitMQ("", exchangeName, routingKey)
    var err error
    //获取connection
    rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
    rabbitmq.failOnErr(err, "failed to connect rabbitmq!")
    //获取channel
    rabbitmq.channel, err = rabbitmq.conn.Channel()
    rabbitmq.failOnErr(err, "failed to open a channel")
    return rabbitmq
}

//路由模式发送消息
func (r *RabbitMQ) PublishRouting(message string) {
    //1.尝试创建交换机
    err := r.channel.ExchangeDeclare(
        r.Exchange,
        //要改成direct
        "direct",
        true,
        false,
        false,
        false,
        nil,
    )

    r.failOnErr(err, "Failed to declare an excha"+
        "nge")

    //2.发送消息
    err = r.channel.Publish(
        r.Exchange,
        //要设置
        r.Key,
        false,
        false,
        amqp.Publishing{
            ContentType: "text/plain",
            Body:        []byte(message),
        })
}

//路由模式接受消息
func (r *RabbitMQ) RecieveRouting() {
    //1.试探性创建交换机
    err := r.channel.ExchangeDeclare(
        r.Exchange,
        //交换机类型
        "direct",
        true,
        false,
        false,
        false,
        nil,
    )
    r.failOnErr(err, "Failed to declare an exch"+
        "ange")
    //2.试探性创建队列，这里注意队列名称不要写
    q, err := r.channel.QueueDeclare(
        "", //随机生产队列名称
        false,
        false,
        true,
        false,
        nil,
    )
    r.failOnErr(err, "Failed to declare a queue")

    //绑定队列到 exchange 中
    err = r.channel.QueueBind(
        q.Name,
        //需要绑定key
        r.Key,
        r.Exchange,
        false,
        nil)

    //消费消息
    messges, err := r.channel.Consume(
        q.Name,
        "",
        true,
        false,
        false,
        false,
        nil,
    )

    forever := make(chan bool)

    go func() {
        for d := range messges {
            log.Printf("Received a message: %s", d.Body)
        }
    }()

    fmt.Println("退出请按 CTRL+C\n")
    <-forever
}
mainpublish.go代码

package main

import (
    "fmt"
    "strconv"
    "time"

    "github.com/student/kuteng-RabbitMQ/RabbitMQ"
)

func main() {
    kutengone := RabbitMQ.NewRabbitMQRouting("kuteng", "kuteng_one")
    kutengtwo := RabbitMQ.NewRabbitMQRouting("kuteng", "kuteng_two")
    for i := 0; i <= 100; i++ {
        kutengone.PublishRouting("Hello kuteng one!" + strconv.Itoa(i))
        kutengtwo.PublishRouting("Hello kuteng Two!" + strconv.Itoa(i))
        time.Sleep(1 * time.Second)
        fmt.Println(i)
    }

}
recieve1/mainrecieve.go代码

package main

import "github.com/student/kuteng-RabbitMQ/RabbitMQ"

func main() {
    kutengone := RabbitMQ.NewRabbitMQRouting("kuteng", "kuteng_one")
    kutengone.RecieveRouting()
}
recieve2/mainrecieve.go代码

package main

import "github.com/student/kuteng-RabbitMQ/RabbitMQ"

func main() {
    kutengtwo := RabbitMQ.NewRabbitMQRouting("kuteng", "kuteng_two")
    kutengtwo.RecieveRouting()
}


Topic模式
1. Topic模式（话题模式，一个消息被多个消费者获取，消息的目标queue可用BindingKey以通配符，（#：一个或多个词，*：一个词）的方式指定


星号井号代表通配符
星号代表多个单词,井号代表一个单词
路由功能添加模糊匹配
消息产生者产生消息,把消息交给交换机
交换机根据key的规则模糊匹配到对应的队列,由队列的监听消费者接收消息消费
目录结构



kuteng-RabbitMQ

-RabbitMQ

--rabitmq.go //这个是RabbitMQ的封装

-publish

--mainpublish.go //Publish 先启动

-recieve1

--mainrecieve.go

-recieve2

--mainrecieve.go

rabitmq.go代码

package RabbitMQ

import (
    "fmt"
    "log"

    "github.com/streadway/amqp"
)

//连接信息
const MQURL = "amqp://kuteng:kuteng@127.0.0.1:5672/kuteng"

//rabbitMQ结构体
type RabbitMQ struct {
    conn    *amqp.Connection
    channel *amqp.Channel
    //队列名称
    QueueName string
    //交换机名称
    Exchange string
    //bind Key 名称
    Key string
    //连接信息
    Mqurl string
}

//创建结构体实例
func NewRabbitMQ(queueName string, exchange string, key string) *RabbitMQ {
    return &RabbitMQ{QueueName: queueName, Exchange: exchange, Key: key, Mqurl: MQURL}
}

//断开channel 和 connection
func (r *RabbitMQ) Destory() {
    r.channel.Close()
    r.conn.Close()
}

//错误处理函数
func (r *RabbitMQ) failOnErr(err error, message string) {
    if err != nil {
        log.Fatalf("%s:%s", message, err)
        panic(fmt.Sprintf("%s:%s", message, err))
    }
}

//话题模式
//创建RabbitMQ实例
func NewRabbitMQTopic(exchangeName string, routingKey string) *RabbitMQ {
    //创建RabbitMQ实例
    rabbitmq := NewRabbitMQ("", exchangeName, routingKey)
    var err error
    //获取connection
    rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
    rabbitmq.failOnErr(err, "failed to connect rabbitmq!")
    //获取channel
    rabbitmq.channel, err = rabbitmq.conn.Channel()
    rabbitmq.failOnErr(err, "failed to open a channel")
    return rabbitmq
}

//话题模式发送消息
func (r *RabbitMQ) PublishTopic(message string) {
    //1.尝试创建交换机
    err := r.channel.ExchangeDeclare(
        r.Exchange,
        //要改成topic
        "topic",
        true,
        false,
        false,
        false,
        nil,
    )

    r.failOnErr(err, "Failed to declare an excha"+
        "nge")

    //2.发送消息
    err = r.channel.Publish(
        r.Exchange,
        //要设置
        r.Key,
        false,
        false,
        amqp.Publishing{
            ContentType: "text/plain",
            Body:        []byte(message),
        })
}

//话题模式接受消息
//要注意key,规则
//其中“*”用于匹配一个单词，“#”用于匹配多个单词（可以是零个）
//匹配 kuteng.* 表示匹配 kuteng.hello, kuteng.hello.one需要用kuteng.#才能匹配到
func (r *RabbitMQ) RecieveTopic() {
    //1.试探性创建交换机
    err := r.channel.ExchangeDeclare(
        r.Exchange,
        //交换机类型
        "topic",
        true,
        false,
        false,
        false,
        nil,
    )
    r.failOnErr(err, "Failed to declare an exch"+
        "ange")
    //2.试探性创建队列，这里注意队列名称不要写
    q, err := r.channel.QueueDeclare(
        "", //随机生产队列名称
        false,
        false,
        true,
        false,
        nil,
    )
    r.failOnErr(err, "Failed to declare a queue")

    //绑定队列到 exchange 中
    err = r.channel.QueueBind(
        q.Name,
        //在pub/sub模式下，这里的key要为空
        r.Key,
        r.Exchange,
        false,
        nil)

    //消费消息
    messges, err := r.channel.Consume(
        q.Name,
        "",
        true,
        false,
        false,
        false,
        nil,
    )

    forever := make(chan bool)

    go func() {
        for d := range messges {
            log.Printf("Received a message: %s", d.Body)
        }
    }()

    fmt.Println("退出请按 CTRL+C\n")
    <-forever
}
mainpublish.go代码

package main

import (
    "fmt"
    "strconv"
    "time"

    "github.com/student/kuteng-RabbitMQ/RabbitMQ"
)

func main() {
    kutengOne := RabbitMQ.NewRabbitMQTopic("exKutengTopic", "kuteng.topic.one")
    kutengTwo := RabbitMQ.NewRabbitMQTopic("exKutengTopic", "kuteng.topic.two")
    for i := 0; i <= 100; i++ {
        kutengOne.PublishTopic("Hello kuteng topic one!" + strconv.Itoa(i))
        kutengTwo.PublishTopic("Hello kuteng topic Two!" + strconv.Itoa(i))
        time.Sleep(1 * time.Second)
        fmt.Println(i)
    }

}
recieve1/mainrecieve.go代码

package main

import "github.com/student/kuteng-RabbitMQ/RabbitMQ"

func main() {
    kutengOne := RabbitMQ.NewRabbitMQTopic("exKutengTopic", "#")
    kutengOne.RecieveTopic()
}
recieve2/mainrecieve.go代码

package main

import "github.com/student/kuteng-RabbitMQ/RabbitMQ"

func main() {
    kutengOne := RabbitMQ.NewRabbitMQTopic("exKutengTopic", "kuteng.*.two")
    kutengOne.RecieveTopic()
}