# docker 
> 참조 : https://www.redhat.com, https://www.docker.com/

## 개념
### + container

- 애플리케이션을 실제 구동 환경으로부터 추상화할 수 있는 논리 패키징 메커니즘을 제공<br>
- 코드와 라이브러리 및 의존성을 가진 모든 항목을 패키지화<br>
- Immutable infrastructure : 변경 불가능한 인프라 구조
- kernel layer와의 분리를 통한 ***경량화***
- 개발 단계에서 테스트 및 프로덕션 까지의 ***이식성과 일관성*** 유지<br>
- 개발, 배포, 관리의 lifecycle ***단순화*** 를 제공하며 CI/CD 최적화 <br><br>


### + Docker
-  Go 언어로 개발된 컨테이너 기반의 추상화를 제공하는 플랫폼 <br>
- LXC 기반으로 사용하다 0.9버전에서는 libcontainer Driver를 통해서 직접 kernel API에 접근하는 libcontainer 사용, 이후 [OCI](https://www.opencontainers.org/)를 준수하는 runC로 대체 <br>
- 2008년에 dotCloud를 통해 동명의 컨테이너 기술과 함께 Docker가 등장했으며, 2013년 [Pycon](https://pycon.org/)에서 공개<br>
<br>


- Docker Version

Capabilities|	Docker Engine - Community|	Docker Engine - Enterprise|	Docker Enterprise
:---:|:---:|:---:|:---:
컨테이너 엔진 및 오케스트레이션 네트워킹, 보안 내장| O|O|O
공인 된 인프라, 플러그인 및 ISV 컨테이너 || O| O
이미지 관리	|||O
컨테이너 앱 관리 |||O
이미지 보안 스캐닝 |||O
> [Docker Engine 릴리즈 노트](https://docs.docker.com/engine/release-notes/)<br>

<br><br>

### + Docker architecture
- 클라이언트 - 서버 아키텍처<br>
- Docker 데몬은 API 요청을 수신하고 이미지, 컨테이너, 네트워크 및 볼륨과 같은 Docker 객체를 관리<br>
- Docker 클라이언트와 데몬은 UNIX 소켓 또는 네트워크 인터페이스를 통해 REST API를 사용하여 통신<br>
- Docker 레지스트리 는 Docker 이미지를 저장<br>

![docker component](https://docs.docker.com/engine/images/engine-components-flow.png)<br><br>


### + 구성요소

Tech | 세부 
:---:|:---
Docker engine | daemon process, api를 통해 요청을 확인하고 Docker 구성 요소와 통신하여 서비스를 수행
Containerd | Docker와 Kubernetes는 물론 syscalls을 추상화하려는 다른 컨테이너 플랫폼이나 Linux, Windows, Solaris 또는 다른 OS에서 container를 관리해주는 daemon
cgroup | 프로세스 또는 프로세스 그룹의 리소스 사용을 제어하고 제한
namespaces | PID - process isolation<br> NET - managing network interfaces<br> IPC - managing access to IPC resources<br> MNT - managing filesystem mount points<br> UTS - isolating kernel and version identifiers
user-defined networks | bridge, overlay, macvlan 
union file systems | Image Layering and sharing - AUFS, overlayFS, DeviceMapper
<br>

![Docker architecture](https://i2.wp.com/blog.docker.com/wp-content/uploads/974cd631-b57e-470e-a944-78530aaa1a23-1.jpg?w=906&ssl=1)

<br><br>

## 테스트 환경 구성
### + 환경 
Docker Community Ver. && Ubuntu16.04

### + 설치
```
# apt update
# apt install software-properties-common apt-transport-https ca-certificates curl
# curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
# apt-key fingerprint 0EBFCD88
# add-apt-repository \
   "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
   $(lsb_release -cs) \
   stable"
# apt update 
# apt install docker-ce docker-ce-cli containerd.io
```

<br>

### + 설치 확인
```
# docker version
Client: Docker Engine - Community
 Version:           19.03.1
 API version:       1.40
 Go version:        go1.12.5
 Git commit:        74b1e89e8a
 Built:             Thu Jul 25 21:21:35 2019
 OS/Arch:           linux/amd64
 Experimental:      false

Server: Docker Engine - Community
 Engine:
  Version:          19.03.1
  API version:      1.40 (minimum version 1.12)
  Go version:       go1.12.5
  Git commit:       74b1e89e8a
  Built:            Thu Jul 25 21:20:09 2019
  OS/Arch:          linux/amd64
  Experimental:     false
 containerd:
  Version:          1.2.6
  GitCommit:        894b81a4b802e4eb2a91d1ce216b8817763c29fb
 runc:
  Version:          1.0.0-rc8
  GitCommit:        425e105d5a03fabd737a126ad93d62a9eeede87f
 docker-init:
  Version:          0.18.0
  GitCommit:        fec3683

# ps -ef |egrep -i "containerd|dockerd" |grep -v grep
root     18923     1  0 13:34 ?        00:00:00 /usr/bin/containerd
root     19028     1  0 13:34 ?        00:00:00 /usr/bin/dockerd -H fd:// --containerd=/run/containerd/containerd.sock

```
<br>

### + version 선택
```
# apt-cache madison docker-ce
 docker-ce | 5:19.03.1~3-0~ubuntu-xenial | https://download.docker.com/linux/ubuntu xenial/stable amd64 Packages
 docker-ce | 5:19.03.0~3-0~ubuntu-xenial | https://download.docker.com/linux/ubuntu xenial/stable amd64 Packages
 docker-ce | 5:18.09.8~3-0~ubuntu-xenial | https://download.docker.com/linux/ubuntu xenial/stable amd64 Packages
 docker-ce | 5:18.09.7~3-0~ubuntu-xenial | https://download.docker.com/linux/ubuntu xenial/stable amd64 Packages
 docker-ce | 5:18.09.6~3-0~ubuntu-xenial | https://download.docker.com/linux/ubuntu xenial/stable amd64 Packages
~

# apt install docker-ce=[version] docker-ce-cli=[version] containerd.io
```
<br><br>

## 테스트
### + Public Repo 연동
> https://hub.docker.com 계정 생성 필요 
```
# docker info
~

# docker login
Login with your Docker ID to push and pull images from Docker Hub. If you don't have a Docker ID, head over to https://hub.docker.com to create one.
Username: ****
Password:
WARNING! Your password will be stored unencrypted in /root/.docker/config.json.
Configure a credential helper to remove this warning. See
https://docs.docker.com/engine/reference/commandline/login/#credentials-store

Login Succeeded

# docker search ubuntu
NAME                                                      DESCRIPTION                                     STARS               OFFICIAL            AUTOMATED
ubuntu                                                    Ubuntu is a Debian-based Linux operating sys…   9779                [OK]            
dorowu/ubuntu-desktop-lxde-vnc                            Docker image to provide HTML5 VNC interface …   326                                     [OK]
rastasheep/ubuntu-sshd                                    Dockerized SSH service, built on top of offi…   225                                     [OK]
consol/ubuntu-xfce-vnc                                    Ubuntu container with "headless" VNC session…   184                                     [OK]
~
> stars는 이미지가 즐겨찾기된 횟수, official은 해당 이미지가 검증된 이미지임을 표시

# docker pull ubuntu
Using default tag: latest
latest: Pulling from library/ubuntu
7413c47ba209: Pull complete
0fe7e7cbb2e8: Pull complete
1d425c982345: Pull complete
344da5c95cec: Pull complete
Digest: sha256:c303f19cfe9ee92badbbbd7567bc1ca47789f79303ddcef56f77687d4744cd7a
Status: Downloaded newer image for ubuntu:latest
docker.io/library/ubuntu:latest

# docker images
REPOSITORY          TAG                 IMAGE ID            CREATED             SIZE
ubuntu              latest              3556258649b2        5 days ago          64.2MB
```
<br>

### + 기본 동작 
```
# iptables -nL 
# iptables -nL -t nat
> docker chain 확인

# docker run -i -t ubuntu /bin/echo 'hi pjin'
hi pjin
> 우분투 컨테이너가 주어진 잡을 실행하고 종료하는 명령, 일회성으로 수행하는 배치 잡에 적합
```
<br>

### + 도커 실행
```
# docker ps
CONTAINER ID        IMAGE               COMMAND             CREATED             STATUS              PORTS               NAMES
# docker ps -a
CONTAINER ID        IMAGE               COMMAND                 CREATED              STATUS                          PORTS               NAMES
2c2dd8c18d6c        ubuntu              "/bin/echo 'hi pjin'"   About a minute ago   Exited (0) About a minute ago                       bold_keller
> List containers

# docker inspect ubuntu
~
> low-level information on Docker objects

# docker run -i -t --name ubuntu-test01 ubuntu
root@fabdb18b69d4:/#
```
 > 로컬이미지가 없으면 docker pull로 땡겨옴<br>
 > docker create && docker start 자동 진행<br>
 > -i -t 옵션은 docker attach 로 컨테이너로 접속(-i : 표준입력 with /bin/bash, -t : tty 쉘 사용)<br>
  >> --add-host=[] : 컨테이너의 etc/hosts에 추가<br>
  >> -c : cpu-share=1024 & cgroup(control groups)<br>
  >> -d : deamon 실행

> 컨테이너안에 프로세스는 백그라운드로 실행중이어야 하며 실행중인 프로세스가 없다면 중지됨, 중지하지 않고 나오려면 ctl + P + q<br>
  
```
# docker ps
CONTAINER ID        IMAGE               COMMAND             CREATED             STATUS              PORTS               NAMES
fabdb18b69d4        ubuntu              "/bin/bash"         3 minutes ago       Up 3 minutes                            ubuntu-test01

# ip a
> veth 확인

# docker network inspect bridge
> 추가 정보 확인

# docker attach ubuntu-test01
root@fabdb18b69d4:/# hostname
fabdb18b69d4
root@fabdb18b69d4:/# exit
exit

# docker ps -a
CONTAINER ID        IMAGE               COMMAND                 CREATED             STATUS                      PORTS               NAMES
fabdb18b69d4        ubuntu              "/bin/bash"             6 minutes ago       Exited (0) 29 seconds ago                       ubuntu-test01
2c2dd8c18d6c        ubuntu              "/bin/echo 'hi pjin'"   11 minutes ago      Exited (0) 11 minutes ago                       bold_keller
> exit로 나오면 프로세스 중지

# docker rm fabdb18b69d4
fabdb18b69d4
# docker ps -a
CONTAINER ID        IMAGE               COMMAND                 CREATED             STATUS                      PORTS               NAMES
2c2dd8c18d6c        ubuntu              "/bin/echo 'hi pjin'"   12 minutes ago      Exited (0) 12 minutes ago                       bold_keller
```
<br>

### + 기본 네트워크
```
# docker network ls
NETWORK ID          NAME                DRIVER              SCOPE
a1d0bae0cdfa        bridge              bridge              local
0f65aefa12ab        host                host                local
e24369788d09        none                null                local

# docker network inspect [name]
~
```
> --net=[name]을 통해 구성 <br>
> bridge : docker0 bridge를 통해 network namespace 할당 <br>
> host : host와 네트워크 공유 <br>
> none : 네트워크 디바이스 미제공 <br>

<br>

### + 포트 바인딩
```
# docker run -itd -p 8888:80 --name web-test01 ubuntu
c884a5e3bcb67ec8a36fed4bc92b19aafaa505e94a062459e94c3de14ccef5d6

# docker ps
CONTAINER ID        IMAGE               COMMAND             CREATED             STATUS              PORTS                  NAMES
c884a5e3bcb6        ubuntu              "/bin/bash"         10 seconds ago      Up 9 seconds        0.0.0.0:8888->80/tcp   web-test01

# iptables -nL
# iptables -nL -t nat

# docker network ls
NETWORK ID          NAME                DRIVER              SCOPE
a1d0bae0cdfa        bridge              bridge              local
0f65aefa12ab        host                host                local
e24369788d09        none                null                local
# docker network inspect a1d0bae0cdfa

# netstat -antp |grep -i 8888
tcp6       0      0 :::8888                 :::*                    LISTEN      20557/docker-proxy

# docker attach web-test01
root@c884a5e3bcb6:/# apt update
root@c884a5e3bcb6:/# apt install nginx
root@c884a5e3bcb6:/# /etc/init.d/nginx start
root@c884a5e3bcb6:/# /etc/init.d/nginx status


# docker exec web-test01 ip a
# curl localhost:8888
```
<br>

### + 이미지 생성 및 push
```
# docker ps
CONTAINER ID        IMAGE               COMMAND             CREATED             STATUS              PORTS                  NAMES
c884a5e3bcb6        ubuntu              "/bin/bash"         9 minutes ago       Up 9 minutes        0.0.0.0:8888->80/tcp   web-test01

# docker commit -a "pjin" -m "nginx test" web-test01 pjin/web-test01:0.1
sha256:e2929a0e531b7799db17d887bedea5316a8fc2bcbcedfb5e39f15177177c7399

# docker images
REPOSITORY          TAG                 IMAGE ID            CREATED              SIZE
pjin/web-test01     0.1                 68fa1072ebc3        3 seconds ago        151MB
ubuntu              latest              3556258649b2        5 days ago           64.2MB

# docker push pjin/web-test01:0.1
The push refers to repository [docker.io/pjin/web-test01]
f30540dd671f: Pushed
b079b3fa8d1b: Pushed
a31dbd3063d7: Pushed
c56e09e1bd18: Pushed
543791078bdb: Pushed
0.1: digest: sha256:399cce3ab8cd39e03fdadc8348447be2b57ad30bf545b0060152eb54d28ee872 size: 1364
