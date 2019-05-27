# openstack
## Openstack 상세
### + 정의
+ 오픈스택(OpenStack)은 **Open Infrastructure Platform** 형태의 클라우드 컴퓨팅 오픈소스 프로젝트 <br>
+ **As A Service** & **Software Defined** <br>
+ 2012년 창설된 비영리 단체인 **OpenStack Foundation**에서 유지, 보수하고 있으며 아파치 라이센스하에 배포<br>
+ AMD, Intel, 캐노니컬, 수세 리눅스, Red Hat, Cisco , Dell, HP, IBM, NEC, VMWare, Yahoo!등의 150개 이상의 회사가 이 프로젝트에 참가하고 있으며, 주로 리눅스 기반으로 운용과 개발이 이
루어 지고 있음<br>
+ [openstack 홈페이지](https://www.openstack.org/)
<br><br>

### + 상세
##### OpenSource Echosystem<br>
+ 대부분의 Cloud에서 연동을 제공하며 OpenStack 자체가 그 연동의 중심이 되는 플랫폼<br>
+ 각 Component의 독립적인 동작과 분산 구성을 통해 기본적인 HA 구성 지원<br>
+ 다수의 벤더에서 드라이버 제공<br>

#### Elastic
+ 사업 자체에 탄력성을 부여
+ 모든 resource의 가상화 및 통합 관리를 통해 유연한 service 제공 가능
+ 사업의 체질 개선을 위한 cloud service 를 유연하게 제공

#### Tech. internalization
+ 가용되는 인프라-플랫폼-소프트웨어의 통합을 통한 기술 내재화
+ 향 후 기술 중심의 인력 양성을 통해  rehat, hp등 벤더에 의존되지 않는 운영을 가능하게 해줌
+ 최소의 인력으로 인프라 운영을 가능하게 함

#### general
+ 기존의 사용하고 있는 대부분의 장비 사용 가능
+ 다수의 벤더에서 드라이버 제공
+ 클라우드 간 용이한 마이그레이션
+ Software Defined
<br><br>


### + Architecture Design
+ 다양한 Component의 조합으로 구성되는 **Modular Architecture** <br>
> *출처 : https://docs.openstack.org/arch-design/* <br>
> ![OpenStack Logical Architecture](https://raw.githubusercontent.com/engineer-pjin/sre_component_foundation/master/image/openStack_simple_architecture.png)
> ![OpenStack Logical Architecture](https://raw.githubusercontent.com/engineer-pjin/sre_component_foundation/master/image/openStack_logical_architecture.png)

