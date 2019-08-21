# Kubernetes
> 참조 : [쿠버네티스 홈페이지](https://kubernetes.io), [쿠버네티스 시작하기](http://acornpub.co.kr/book/kubernetes-up-and-running), [쿠버네티스 마스터](http://acornpub.co.kr/book/mastering-kubernetes)
<br>

## 기본
### + Kubernetes?
- ***Desired State Management***
- 구글 내부 SRE에서 개발된 borg(Large-scale cluster management)를 기반으로 2014년에 오픈소스로 공개했고, 현재는 CNCF(Cloud Native Computing Foundation)에서 관리
- kubernetes 이름은 키잡이, 파일럿이라는 의미의 그리스어에서 유래
- container화 된 어플리케이션을 위한 배포 플랫폼
- 운영에 있어 best practices에 기반하여 디자인
- 어플리케이션의 lifecycle과 scaling을 관리

<br>

### + 상세
**특징**
- desired state : ***object*** 를 ***label*** 로 구분하여 yaml 파일에 ***선언***
- 속도 : 높은 가용성, 불변성(immutable infrastructure), 선언형, 자가 치유
- 확장성 : 분리된 아키텍처(decoupled architecture), 쉬운 확장 및 예측, msa를 통한 팀의 확장, 일관성과 확장성에 대한 고려사항 분리
- 인프라 추상화 및 효율성

**version**
+ xyz : x major, y minor, z patch
- 2019.08 기준 v1.15.0

**limit**
- node 5000EA 
- pod 150000EA 
- container 300000EA 
- pod 100EA per node


<br><br>

##  structure
![Kubernetes structure](https://raw.githubusercontent.com/engineer-pjin/sre_component_foundation/master/image/post-ccm-arch.png)

- 쿠버네티스 마스터 노드 프로세스<br>
. kube-apiserver<br>
. kube-controller-manager<br>
. kube-scheduler<br>

- 쿠버네티스 노드<br>
. kubelet : 쿠버네티스 마스터와 통신<br>
. kube-proxy : 각 노드의 쿠버네티스 네트워킹 서비스를 반영하는 네트워크 프록시<br>

<br>

### + layer
- 어플리케이션 : 컨테이너로 구동된 어플리케이션이 동작
- Data Plane : 컨테이너가 실행되고 네트워크에 연결될 수 있게 CPU, 메모리, 네트워크, 스토리지와 같은 능력을 제공
- Control Plane : 컨테이너의 라이프사이클을 정의, 배포, 관리하기 위한 API와 인터페이스들을 노출하는 컨테이너 오케스트레이션 레이어

![Kubernetes solutions](https://raw.githubusercontent.com/engineer-pjin/sre_component_foundation/master/image/kubernetessolutions.PNG)
> 관리주체에 따른 범위 구분

<br><br>

## 구성요소
### + basic object
Kubernetes API의 추상화 된 객체, desired state<br>
![Kubernetes objects](https://raw.githubusercontent.com/engineer-pjin/sre_component_foundation/master/image/sisdig_4.png)<br>

 - pod <br>
  . 고래(컨테이너)의 작은 그룹<br>
  . 네트워크 네임스페이스 및 볼륨 공유, ip는 서비스에서는 사용하지 않음<br>
  . Pod는 여러 컨테이너를 가질 수 있지만 대부분 1~2개로 구성<br>
  . 스케일링 또한 컨테이너가 아닌 Pod단위로 수행<br>
  . pod는 하나의 물리적 노드에서 실행<br>
    - pause 컨테이너가 생성되고 이를 통해 linux의 namespace 공유
  . mortal 오브젝트 : 상태에 대해 보장하지 않음
 - service <br>
  . pod의 endpoint를 관리하고 외부에서 Service를 통해 Pod에 접근<br>
  . pod에게 자신의 IP 주소와 포드 집합에 대한 단일 DNS 이름을 제공하고 프록시로 로드밸런싱을 수행<br>
  . 각 서비스가 고유 한 IP를 수신하도록하기 위해 내부 할당자가 자동으로 etcd 의 전역 할당 맵을 업데이트<br>
   - etcd : 클러스터 데이터를 담는 키-값 저장소
   - 서비스 환경 변수와 dns 두가지의 모드 제공<br>
  . service proxy : Pod-to-Service 및 External-to-Service 네트워킹 관리, 서비스의 가상 IP를 서비스가 제어하는 ​​백엔드 포드의 IP로 변환
   - User space proxy mode : 외부 접근을 iptables를 통해 kube-proxy로 전달, round-robin algorithm<br>
   - iptables proxy mode : 외부접근이 iptables에서 pod로 direct로 전달, 사용자 공간과 커널 공간 사이를 전환 할 필요없이 Linux netfilter가 트래픽을 처리, 10,000개 이상 서비스에서 느려지며 pod 미응답 시 실패 처리<br>
   - IPVS(IP Virtual Server) 프록시 모드 : netfilter 후크 기능을 기반으로하지만 해시 테이블을 기본 데이터 구조로 사용하고 커널 공간에서 l4 layer로 작동, 라운드로빈/최소연결/대상 해시 등 다양한 밸런싱 옵션 제공<br>
 - volume : PV(PersistentVolume - 리소스)와 PVC(PersistentVolumeClaim - 요청)를 통해 별도의 lifecycle 관리<br>
  . Control Plane의 인터페이스를 통해 연동, CSI(Container Storage Interface)<br>
  . shared block storage : 기본구성으로 사용됨<br>
   - nfs, iscsi, fc, ceph, glusterFS, vsphereVolume, aws EBS, azure Disk 등 다양한 백앤드 지원<br>
   - rook :  Kubernetes에서 Ceph를 yaml에 선언된 상태로 클러스터 배포부터 관리 까지 제공<br>
  . object storage 사용시 별도의 app 필요(ex MinIO - aws s3 연동, ibmcloud-object-storage-plugin)<br>
 - namespace : 가상 클러스터<br>
  . 물리적 클러스터를 통해 지원되는 여러 가상 클러스터를 지원, 여러 사용자간에 클러스터 리소스를 나누는 방법<br>
  . 레이블 을 사용 하여 동일한 네임 스페이스 내의 다른 리소스를 구별<br>
  . DNS entry form은 <service-name>.<namespace-name>.svc.cluster.local<br>

<br>

### + Controllers 
basic objects를 기반으로 정의된 형상을 관리하고 부가 기능 및 편의 기능을 제공하는 higher-level 추상화 객체<br>
 - ReplicaSet : 정의된 수의 포드 단위 복제본 유지<br>
 - Deployments <br>
  . 컨테이너를 어떻게 생성하고 업데이트해야 하는지를 지시<br>
  . 자동 복구(self-healing) 메커니즘을 제공<br>
 - StatefulSets<br>
  . stateful applications을 관리하는 workload API object<br>
  . Kubernetes 1.9 버전부터 정식 지원 <br>
  . PersistentVolume Provisioner 에 의해 제공된 볼륨 사용<br>
 - DaemonSet : 클러스터 내 모든 노드에 정의된 컨테이너 생성<br>
 - Jobs <br>
  . 배치 성격의 컨테이너를 실행하고 종료 까지 추적<br>
  . 병렬 배치 배치 가능 <br>
 - ingress<br>
  . 로드밸런싱을 위한 유연하고 독립적이며 이식 가능한 방법
  . 클러스터 외부에서 접근하는 요청들에 대한 응답을 정의하며 HTTP(S)기반의 L7 로드밸런싱, 경로 라우팅, ssl 인증서 등을 제공하고 백엔드 테크와 연동<br>
  . pod ip로 요청 전달<br>
  . Controller 유형 : ingress-nginx, Kong, haproxy-ingress, F5 Container Ingress, openstack octavia-ingress, ingress-gce, AWS ALB Ingress ..<br>
![Kubernetes nginx Ingress](https://raw.githubusercontent.com/engineer-pjin/sre_component_foundation/master/image/NGINX-Ingress-Controller-4-services.png)
<br><br>

### + HPA(Horizontal Pod Autoscaler)
**scale 유형에 따른 구분** <br> 

name | scale point | detail
---|:---|:---
CA(Cluster Autoscaler) | Kubernetes node | cloud platform과 연동
VPA(Vertical Pod Autoscaler) | pod scale up | scale up 시 pod 재시작
HPA(Horizontal Pod Autoscaler) | pod scale out | K
<br> 

**HPA 상세** <br>
- CPU 사용량 (또는 사용자 정의 메트릭, 아니면 다른 애플리케이션 지원 메트릭)을 관찰하여 레플리케이션 컨트롤러, 디플로이먼트 또는 레플리카 셋의 pod 개수를 자동으로 스케일<br>
 . metrics.k8s.io : 리소스 메트릭, 클러스터 애드온<br>
 . custom.metrics.k8s.io : 메트릭 솔루션 공급 업체에서 제공하는 “어댑터” API 서버에서 제공(ex: Prometheus)<br>
 . external.metrics.k8s.io : 클러스터 외부에서 오는 메트릭을 기반, HPA v2 API extension proposal에서 제안 및 도입<br>
  > Custom Metrics Adapter는 Custom Metrics API와 External Metrics API를 모두 제공 할 것으로 예상되지만 이는 필수 사항이 아니며 두 API를 별도로 구현하여 사용할 수 있음
- Kubernetes 1.6부터 멀티 메트릭을 기반으로 스케일링을 지원<br>
- 정해진 주기 동안 컨트롤러 관리자는 각 HorizontalPodAutoscaler 정의에 지정된 메트릭에 대해 리소스 사용률을 질의 후 레플리케이션 컨트롤러 또는 디플로이먼트에서 레플리카 개수를 주기적으로 조정<br>
 . --horizontal-pod-autoscaler-sync-period : 컨트롤 루프 주기<br>
 . --horizontal-pod-autoscaler-upscale-delay : scale out후 upscale을 위한 delay<br>
 . --horizontal-pod-autoscaler-downscale-delay : scale out 후 downscale을 위한 delay<br>
 ![Kubernetes HPA](https://raw.githubusercontent.com/engineer-pjin/sre_component_foundation/master/image/k8s_hpa.PNG)

<br><br>

## 클러스터 연동
### + multizone cluster
- version 1.2 부터 제공
- 별도 존(aws az와 같은 개념) 에서 구현되는 Cluster Federation feature의 경량 버전<br>
- 전체 클러스터의 단일 정적 엔드포인트를 제공하고 클러스터 포드를 지정된 zone에 구성<br>
- 현재 GCE 및 AWS에서 지원<br>
![google k8s engine, regional cluster](https://raw.githubusercontent.com/engineer-pjin/sre_component_foundation/master/image/gcp-google-kubernetes-engine-regional-clusterbcum.png)
> https://cloud.google.com/kubernetes-engine/docs/concepts/regional-clusters?hl=ko
<br>

<br>

### + kubefed(Kubernetes Cluster Federation)
> https://github.com/kubernetes-sigs/kubefed

 - Cluster Federation feature은 별도의 Federation Control Plane 필요<br>
 - 단일 API 엔드포인트에서 여러 Kubernetes 클러스터의 구성을 조정<br>
 - multi-geo applications 배포나 재해복구와 같은 다중 클러스터 사용 사례의 기초<br>
  . 베타 개발 중<br>
 - Multi-Cluster Ingress DNS <br>
 - KubeFed에서 제공하는 추상화 컨셉<br>
  . Templates : 클러스터 전반에서 공통된 리소스 표현을 정의<br>
  . Placement : 리소스가 표시 될 클러스터를 정의<br>
  . Overrides : 템플릿에 적용 할 클러스터 단위 필드 수준 변형을 정의<br>
 - 제공되는 building blocks : Status, Policy, Scheduling <br>

![concepts](https://raw.githubusercontent.com/engineer-pjin/sre_component_foundation/master/image/kubefed_concepts.png)

<br><br>

## 설치
### + 환경
- os : ubuntu 16.04
- node
- network
- 배포 tool : Kubespray

<br><br>
### + Kubespray
- 배포 툴 종류 : kubeadm, kops, rancher, Kubespray
