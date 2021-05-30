FROM golang:alpine

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io,direct

WORKDIR /home/
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk add --no-cache git && \
    git clone https://gitee.com/moglin/buaashow.git && \
    git clone -b gh-pages https://gitee.com/moglin/buaashow-admin.git &&\
    git clone -b gh-pages https://gitee.com/moglin/workstation.git 

RUN cd buaashow && go build -o buaashow

FROM nginx:alpine

ENV GIN_MODE=release
COPY nginx.conf /etc/nginx/conf.d/default.conf
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk add --no-cache  gettext tzdata   && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" >  /etc/timezone && \
    date && \
    apk del tzdata

WORKDIR /home/

COPY config-temp.yaml ./config.yaml
COPY run-docker.sh ./run.sh
RUN chmod +x run.sh
RUN mkdir resources && mkdir static 
COPY --from=0 /home/buaashow-admin static/admin 
COPY --from=0 /home/workstation static/front
COPY --from=0 /home/buaashow/buaashow buaashow

ENTRYPOINT ./run.sh