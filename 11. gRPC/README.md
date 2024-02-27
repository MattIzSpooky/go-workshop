# 10. gRPC

[gRPC](https://grpc.io/) is a way of communicating over HTTP, like [REST](https://restfulapi.net/), [SOAP](https://www.indeed.com/career-advice/career-development/what-is-soap-api) and [GraphQL](https://graphql.org/). 
It allows for [rpc](https://en.wikipedia.org/wiki/Remote_procedure_call) calls over HTTP/2 using protobuf (default) or json.

It offers many benefits over the other methods of communication, some of which are speed and consistency. 
It being powered by [protobuf](https://protobuf.dev/overview/) by default makes it pretty fast as it is a binary format.
This has some benefits when it comes to message sizes, un- & marshalling performance compared to REST/GraphQL (JSON) and SOAP (XML, the horror oh the horror 😰).
My favorite benefit is consistency. A protobuf file, like seen [here](./chat.proto) declares how should be communicated with a gRPC server.
Code is generated from one of those files, be it C#, Java, C++ or Go. It does not matter, they can communicate with each other through gRPC.
A protobuf file is kind of like a [WSDL](https://www.soapui.org/docs/soap-and-wsdl/working-with-wsdls/) but without the pain associated with it and SOAP in general.

Another neat feature of gRPC is the built-in support for streaming. gRPC supports:
- Client-side streaming
- Server-side streaming
- Bidirectional streaming (client & server)

This example contains a chat application. Both the [server](./server-app) and the [client](./client-app).
Messages are being streamed to the server and the server will broadcast a received message to other users in the same chat room.
Is this a perfect example? Not really, things can probably be cleaned up or written in another way, but I had fun writing it and learned a lot about gRPC and its quirks. 
It also shows one of gRPCs strongest benefits, streaming. But it does not neglect simple gRPC calls.

To run the applications you first need to follow the [prerequisites section in the quickstart guide guide on the grpc website](https://grpc.io/docs/languages/go/quickstart/#prerequisites).

To (re)generate the code, run the following command:

```bash
protoc --go_out=generated --go_opt=paths=import \
    --go-grpc_out=generated --go-grpc_opt=paths=import \
    proto/chat.proto
```

Both the client and the server can run locally. They just need to be compiled first. 
cd to each folder and run `go build .`