## Docker 运行

### 安装docker

- 官方文档  https://docs.docker.com/engine/install/
- 阿里镜像  https://developer.aliyun.com/mirror/docker-ce
- 清华镜像  https://mirrors.tuna.tsinghua.edu.cn/help/docker-ce/

如果非root用户，把需要操作的用户添加到docker group，然后重启shell

```shell
sudo groupadd docker
sudo usermod -aG docker ${USER}
sudo service docker restart
```

### 运行docker

需要复制文件`deploy/nginx.conf`,`deploy/Dockerfile`,`deploy/config-temp.yaml`,`deploy/run-docker.sh`到工作目录

修改config-temp.yaml里的配置文件(必选，按照实际情况)

修改nginx.conf里的配置(可选，需要同步修改dockerfile里的文件位置)

然后运行下面的命令

```
docker build -t buaashow:1.0 .
# -d 后台运行
# -p 映射端口:内部端口
# -name 容器名字
# buaashow:1.0为docker build时的-t的参数
docker run -d -p 8080:80 --name buaashow-v1 buaashow:1.0

# 进入容器操作
docker exec -it ${CONTAINER ID} /bin/sh
```

### 文件路径映射

运行容器时指定挂载目录

```shell
docker run -v /hostTest:/conainterTest 
```

同时修改配置文件的`static`选项
