SRV=rouyi-go-service
VERSION=latest

.PHONY: build
build:
	rm -rf ./${SRV} # 清理二进制文件
	go env -w CGO_ENABLED=0
	go build -a -ldflags '-extldflags "-static"' -o ${SRV} main.go
	go env -w CGO_ENABLED=1

.PHONY: docker
docker:
	# 登录镜像仓库
	docker login
	docker build . -t ${SRV}:${VERSION}

	docker tag ${SRV}:${VERSION} ${SRV}:${VERSION}
	docker push ${SRV}:${VERSION}
	rm -rf ./${SRV} # 清理二进制文件

.PHONY: docs
docs:
	# 生成 swagger 文档
	swag init -o ./app/swagger

start:
	go run main.go