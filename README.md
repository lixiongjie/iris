

# Iris 
 
 
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



