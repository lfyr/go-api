IMAGE_NAME=13.229.13.172:5000/admin:latest

# 编译
build:
	@cd app/web && go build

# 构建镜像
image: build
	@docker build -t $(IMAGE_NAME) .

# 推送镜像
push: image
	@docker push $(IMAGE_NAME)

# docker run  -d --name appstore -p 8886:8888   13.209.13.178:5000/server