# openstack
## Openstack 상세
### + 정의

+ 오픈스택(OpenStack)은 **Open Infrastructure ** 형태의 클라우드 컴퓨팅 오픈소스 프로젝트 <br>
+ **As A Service** & **Software Defined** <br>
+ 인프라의 추상화
+ 2012년 창설된 비영리 단체인 **OpenStack Foundation**에서 유지, 보수하고 있으며 아파치 라이센스하에 배포<br>
+ AMD, Intel, 캐노니컬, 수세 리눅스, Red Hat, Cisco , Dell, HP, IBM, NEC, VMWare, Yahoo등의 150개 이상의 회사가 이 프로젝트에 참가하고 있으며, 주로 리눅스 기반으로 운용과 개발이 이
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

## OpenStack Architecture
### + Architecture : Component
#### 주요 Component 
Service | Project Name | 비고
:---:|:---:|:---|
DashBoard | Horizone | web-based self-service portal
Compute | Nova | Manages the lifecycle of compute instances
Networking | Neutron | Network-Connectivity-as-a-Service
Identity | Keystone | authentication and authorization service
Image | Glance | Stores and retrieves virtual machine disk images
Telemerty | Ceilometer | Monitors and meters
Object Storage | Swift | Stores and retrieves arbitrary unstructured data objects
Block Storage | Cinder | Provides persistent block storage
Orchestration | Heat | Orchestrates multiple composite cloud applications
> *출처 : https://docs.openstack.org/ <br>
> ![OpenStack Logical Architecture](https://raw.githubusercontent.com/engineer-pjin/sre_component_foundation/master/image/openStack_simple_architecture.png)

<br><br>

### + Architecture : Design
+ 다양한 Component의 조합으로 구성되는 **Modular Architecture** <br>
+ Component를 통해 Backend Tech.를 제어 <br>

> *출처 : https://docs.openstack.org/ <br>
> ![OpenStack Logical Architecture](https://raw.githubusercontent.com/engineer-pjin/sre_component_foundation/master/image/openStack_logical_architecture.png)

<br><br> 

### + Architecture : Networking
#### Network zones
+ Underlay : defined as the physical network switching infrastructure
+ Overlay : defined as any L3 connectivity between the cloud components 
+ Edge : network traffic transitions from the cloud overlay or SDN networks into the traditional network environments

> *출처 : https://docs.openstack.org/ <br>
> ![OpenStack Logical Architecture](https://raw.githubusercontent.com/engineer-pjin/sre_component_foundation/master/image/openstack_l2_network.png)
> ![OpenStack Logical Architecture](https://raw.githubusercontent.com/engineer-pjin/sre_component_foundation/master/image/openstack_l3_network.png)

#### Traffic flow
+ East/West : traffic flow between workload within the cloud as well as the traffic flow between the compute nodes and storage nodes falls 
+ North/South : traffic between the workload and all external networks

> *출처 : https://docs.openstack.org/ <br>
> ![OpenStack Logical Architecture](https://raw.githubusercontent.com/engineer-pjin/sre_component_foundation/master/image/openstack_East_West_network.png)
> ![OpenStack Logical Architecture](https://raw.githubusercontent.com/engineer-pjin/sre_component_foundation/master/image/openstack_North_South_network2.png)
> ![OpenStack Logical Architecture](https://raw.githubusercontent.com/engineer-pjin/sre_component_foundation/master/image/openstack_North_South_network.png)

<br><br>

## HA 
### + HA : SPOF
+ OpenStack 에서의 SPOF는? : 모든 Component 
+ 각 Component 별 HA 구성을 진행
+ 플랫폼의 규모에 따라 Performance sizing(Component 분리)과는 별개로 Component 별 복수 구성

<br><br>

### + HA : Management 
+ 전통적 방식의 Component A/S Cluster 구성 : DRBD, Pacemaker, Corosync
+ 3 + 2n Node 구성 : HAProxy, Pacemaker, Galera Cluster

> *출처 : https://docs.openstack.org/ <br>
> ![OpenStack Logical Architecture](https://raw.githubusercontent.com/engineer-pjin/sre_component_foundation/master/image/openstack_ha_cont_network.jpg)
> ![OpenStack Logical Architecture](https://raw.githubusercontent.com/engineer-pjin/sre_component_foundation/master/image/openstack_ha_com_network.jpg)

<br><br>

### + HA : Network 
+ DVR(*Distributed Virtual Routing*) <br>
+- resides completely on the compute nodes for instances with a fixed or floating IP address using project networks on the same distributed virtual router<br>
+- 네트워크 노드에 집중되었던 Virtual Router 를 VM이 생성되는 각 Compute Node에 분산 배치<br>

> *출처 : https://docs.openstack.org/ <br>
> ![OpenStack Logical Architecture](https://raw.githubusercontent.com/engineer-pjin/sre_component_foundation/master/image/openstack_dvr_network.png)

<br><br>

### + HA : Storage
#### SDS (Software Defined Storage)<br>
+ Problem : Big & Many Data, Big & Many Data, Scale up Limit, Non Elastic
+ SDS’s Advantage : Scale Out, Elastic, Low Cost, Non lock-in, No limit on the machine
+ Examples : GlusterFS, Ceph, VMware Virtual SAN, EMC ScaleIO

#### Ceph<br>
+ Block & Object Storage & filesystem Support
+ Monitor – Cluster Map & OSD 관리, 3+2N 형태의 HA 구성
+ OSD – data read & write, **Object 단위** data 저장, CRUSH Algorithm
+ Monitor는 Client에 Cluster Map을 전달하며 이를 통해 어떤 OSD Nodes 에서 read & write 할지 결정


> *출처 : https://ceph.com/ <br>
> ![OpenStack Logical Architecture](https://raw.githubusercontent.com/engineer-pjin/sre_component_foundation/master/image/Ceph-diagram.png)
> ![OpenStack Logical Architecture](https://raw.githubusercontent.com/engineer-pjin/sre_component_foundation/master/image/ceph-stack.png)

<br><br>

## Performance Tuning 
### + Performance Tuning : pci passthrough
+ 컴퓨트 노드의 물리적 PCI 장치를 인스턴스에서 액세스하고 직접 제어<br>
+- GPU, NIC, Disk, HBAA등의 pci 장치 연결 가능


<br><br>

### + Performance Tuning : SR-IOV<br>
+ Single Root I/O Virtualization
+ NIC를 hardware 수준에서 multiple separate physical PCIe devices로 가상화하여 제공<br>
+ 하나의 Physical Function (PF)에서 Virtual Function(VF)이라고 하는 가상 PCI 카드를 생성하여 Hypervisor의 기능을 사용하지 않고도 VM에게 network interface 제공 가능 <br>
+ near-native performance로서 para-virtualized drivers, emulated access 보다 성능적 우위<br>

> *출처 : https://intel.com, https://www.redhat.com <br>
> ![OpenStack Logical Architecture](https://raw.githubusercontent.com/engineer-pjin/sre_component_foundation/master/image/perf_sriov.png)
> ![OpenStack Logical Architecture](https://raw.githubusercontent.com/engineer-pjin/sre_component_foundation/master/image/perf_sriov2.png)


### + Performance Tuning : DPDK
#### Latency <br>
+ Openstack을 구성하는 많은 logical devices에 의해 Latency 가 발생하며 이는 Network에 기반<br>
+ Hypervisor 는 network에 최적화 되도록 설계되지 않음<br>
+ Device는 interrupt 를 유발하며 kernel space 영역과 user space 영역간의 mem copy 발생<br>

> *출처 : https://intel.com <br>
> ![OpenStack Logical Architecture](https://raw.githubusercontent.com/engineer-pjin/sre_component_foundation/master/image/dpdk00.png)

#### Data Plan Developer Kit <br>
+ Data Plan Developer Kit <br>
+ 텔에서 개발한 고성능 패킷 처리 시스템 소프트웨어로 고속 패킷 처리를 위한 라이브러리와 드라이버의 집합<br>
리눅스나 윈도우의 커널 대신에 네트워크 패킷을 처리하는 응용 프로그램 - PMD (Poll Mode Driver)을 제공하고 전용 CPU 코어를 할당하여 커널을 거치지 않고 네트워크 카드에 도착한 패킷을 직접 받아서 빠르게 처리<br>
+ 인터럽트의 오버헤드를 줄이는 기술<br>
> ![OpenStack Logical Architecture](https://raw.githubusercontent.com/engineer-pjin/sre_component_foundation/master/image/dpdk01.png)

<br><br>

## OpenStack의 방향성
> 가상머신을 제공하는 플렛폼이 아닌 모든 인프라를 추상화 하고 서비스로 제공 

기술명 | 세부 
:---:|:---:
 Magnum | 컨테이너 오케스트레이션 엔진을 프로비저닝, 스케일링 및 관리를 하기 위한 서비스 집합
 Kuryr | 컨테이너 네트워킹 모델 ↔ OpenStack 네트워킹 모델, 컨테이너에서 Neutron에서 정의된 네트워크를 직접 사용
 Fuxi | 컨테이너 스토리지 모델 ↔ OpenStack 스토리지 모델
 Manila | 컨테이너/K8s에서 Cinder 및 Manila 볼륨을 직접 사용 가능
 Zun | OpenStack 인프라 환경 위에 컨테이너를 직접 관리, 컨테이너를 관리 가능한 모든 API를 제공함
 Kolla | OpenStack 핵심 구성 요소들을 컨테이너화하여 관리
 OpenStack-helm | Helm은 Kubernetes에 어플리케이션을 chart 단위로 배포하는 도구
 Airship | 오픈 인프라의 라이프 싸이클을 시작부터 끝까지 모두 관리
 Kata-Container | 표준 경량 VM 구현체 구축을 위한 오픈소스 프로젝트, 컨테이너 기술처럼 사용가능하게 구현하며 VM의 워크로드 격리
 StarlingX | Wind River® Titanium Cloud R5 제품을 기반으로 한 완벽한 기능의 고성능 Edge Cloud 소프트웨어 스택




