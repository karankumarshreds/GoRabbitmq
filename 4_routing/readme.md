# Exchange type == "DIRECT"

We were using a `fanout` exchange, which doesn't give us much flexibility - it's only capable of mindless broadcasting.

We will use a `direct` exchange instead. The routing algorithm behind a direct exchange is simple - a message goes to the queues whose `queue-bind routingKey` (sub) exactly matches the `publish routingKey` (pub) of the message.
