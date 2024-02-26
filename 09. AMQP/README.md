## 9. AMQP

Message brokers are important for integrating with other software, or other modules within the same software.
You put a message on a queue/topic and a subscriber of said queue or topic will read the message and do start doing this.
This makes it relatively easy to add asynchronous behavior to the architecture of your application.

- [What is a message broker?](https://www.ibm.com/topics/message-brokers)
- [What is a wire protocol?](https://en.wikipedia.org/wiki/Wire_protocol)

This example shows you how to read/write to an AMQP topic/queue.
The AMQP broker used in this example is ActiveMQ Artemis because it is modern and I felt like it.
It also supports AMQP 1.0 out of the box, unlike Rabbit MQ, which uses AMQP 0.9.1, though with the use of a plugin AMQP 1.0 is supported.

[What is AMQP?](https://en.wikipedia.org/wiki/Advanced_Message_Queuing_Protocol)

This example contains a [docker compose file](./docker-compose.yml) that sets up the message broker for you.
Once it is running you can go to http://0.0.0.0:8161/console to view the ActiveMQ Artemis console.

Once Artemis is running, you can run [main.go](main.go).