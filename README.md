# Cortito

This is a simple URL shortener created using microservices.

There are two simple micro services using [Go Micro](https://github.com/micro/go-micro). All microservices use Protocol Buffers for communication.
The services are exposed using an API gateway also using [Micro API](https://micro.mu).

### Common package

All services use common functionality like database access, connection to other services, reading configuration, etc.

The [`common`](./common/) package contains these functionality to avoid code duplication.

### Account service

The [`account`](./account/) service handles users and authentication. It exposes functions to create, update and authenticate a user.
Password can be generated using BCrypt or Argon2. This can be changed in the configuration file.

### Shortener service

The [`shortener`](./shortener/) service is responsible for creating and resolving shortened links. These links are
created per-user, but everyone can follow the link.

### API

The [`api`](./api/) interfaces with the services through a single endpoint. The Micro API gateway is responsible for selecting the node each
request is going to be handled by. This is useful in a multi-node environment.

### Web endpoint

The [`web`](./web/) endpoint is responsible for resolving shortened URLs and redirect the user to the actual URL.
When somebody hits a shortened link the visit for that link is increased.

### Frontend

The [`frontend`](./frontend/) is a React application where users can sign up, create links and see the usage of a particular link.

#### TODO

- Add support for private links
- Generate shortened links a message queue.
- Add WebSockets to the frontend.
- Add information about the client following a link.
