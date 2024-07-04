FROM centos
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai  /etc/localtime
WORKDIR /app
COPY  admin  .
COPY  conf/conf.yaml .
ENTRYPOINT ./admin -c conf.yaml
EXPOSE 8888

# docker run  -d --name appstore -p 8886:8888   13.209.13.178:5000/server