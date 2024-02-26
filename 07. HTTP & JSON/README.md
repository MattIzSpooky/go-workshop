## 7. HTTP & JSON

Go provides a powerful low-level HTTP API.
This application shows how to use the standard http module.

I made the choice not to use any off-the-shelf HTTP frameworks for this example to 
show how easy it is to create your own HTTP server application using only the standard library.

This demo uses just **4.5mb** of RAM on my M1 Mac.

This example also comes with a [docker file](./Dockerfile) to show how easy it is to dockerize a go application.
Also, a [docker compose](./docker-compose.yml) to simplify starting it.

However, for production use. It might be worth using a framework instead.
I would recommend one of the following:
- https://github.com/gin-gonic/gin 
- https://github.com/gofiber/fiber
- https://github.com/beego/beego

