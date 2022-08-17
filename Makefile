VERSION=2.8.4

init:
	docker pull nats:${VERSION}

server:
	docker run -p 4222:4222 nats:${VERSION}
