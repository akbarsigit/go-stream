.PHONY: services
services:
	buf generate --template buf.gen.svc.yaml --path api/stream/v1

.PHONY: gateway
gateway:
	buf generate --template buf.gen.gw.yaml --path api/gateway/v1

.PHONY: all
all: services gateway