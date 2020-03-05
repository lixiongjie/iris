

# Iris 
 
 
 
 
 
 
 
#### 3-5

> 简单队列 & 工作模式

工作模式有负载的作用，一个消息只能被一个消费者消费

> 订阅模式

多个消费者同时消费，队列绑定交换机，一个消息多个消费者

> 路由模式

生产端指定消费端进行消费

> topic 模式

[反过来选](https://www.cnblogs.com/LUA123/p/8477387.html)
 
> rabbitmq 

docker pull rabbitmq:management

```yaml

docker run --name rabbitmq -d -p 15672:15672 -p 5672:5672 rabbitmq:management
docker stop rabbitmq
docker start rabbitmq
```

>控制台信息 

启动容器后，可以浏览器中访问http://localhost:15672来查看控制台信息。
 
RabbitMQ默认的用户名：guest，密码：guest

>  go get github.com/streadway/amqp



