# Exchange

The core idea in the messaging model in rabbitMQ is that _the producer never sends amy messages directly to a queue_.
**Quite often the producer doesn't even know if a message will be delivered to any queue at all**.

The producer can only send messages to an **exchange**. Exchange **reveives** messages from producers and **pushes** them to the queues.

**NOTE** : The exchange must exactly know what to do with the message it receives.

Should it be appended to a particular queue? Should it be appended to many queues? Or should it get discarded. The rules for that are defined by the **exchange type**.

![](2021-09-15-01-13-57.png)

Exchange types:

- direct
- topic
- headers
- fanout

In case we use **nameless** exchange, it will route messages to the queue _with the name specified by ROUTING_KEY paramater_, if it exists.

```Go
err = ch.Publish(
  "",     // exchange
  q.Name, // routing key
  false,  // mandatory
  false,  // immediate
  amqp.Publishing{
    ContentType: "text/plain",
    Body:        []byte(body),
})
```
