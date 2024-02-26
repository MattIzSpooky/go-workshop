## 10. Kafka

Kafka is kind of like a message broker, but it is much more. Kafka has some features that resemble those of a message broker and adds on top of it.
It is mainly used as a message bus that allows for replay of streamed messages.
- [Kafka documentation](https://kafka.apache.org/documentation/#gettingStarted)
- [Differences between Kafka and a message broker](https://aws.amazon.com/compare/the-difference-between-rabbitmq-and-kafka/)

This example contains a [docker compose file](./docker-compose.yml) that sets up Kafka & [Zookeeper](https://zookeeper.apache.org/) for you.

Once Kafka is up and running, you can run [main.go](main.go). Note: it may take a little while for the messages to start coming in after writing them to Kafka.