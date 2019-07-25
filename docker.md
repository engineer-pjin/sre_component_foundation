# docker 
> 참조 : https://www.redhat.com, https://www.docker.com/

## 개념
### + container

- 애플리케이션을 실제 구동 환경으로부터 추상화할 수 있는 논리 패키징 메커니즘을 제공<br>
- 코드와 라이브러리 및 의존성을 가진 모든 항목을 패키지화<br>
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


