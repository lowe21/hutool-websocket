BUILD_PATH = ./build
CONFIG_PATH = ./config
STATIC_PATH = ./static
TEMPLATE_PATH = ./template
APP_NAME = websocket

.PHONY: build
build:
	@mkdir -p ${BUILD_PATH}
	@cp -rf ${CONFIG_PATH} ${STATIC_PATH} ${TEMPLATE_PATH} ${BUILD_PATH}
	@go mod tidy
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${BUILD_PATH}/${APP_NAME} .
	@printf "#!/bin/sh\nps -ef|grep %s|grep -v grep|awk '{print \$$1}'|xargs kill -15" "${APP_NAME}" > ${BUILD_PATH}/stop.sh
	@chmod +x ${BUILD_PATH}/stop.sh
