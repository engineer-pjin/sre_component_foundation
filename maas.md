# MAAS
## MAAS(metal as a service) 
> 참조 : [MAAS 홈페이지](https://maas.io)
### + 정의
+ providing infrastructure coordination<br>
+ '지역 컨트롤러 (regiond)'를 뒷받침하는 중앙 계층화 된 아키텍처 <br>
+ postgres db 사용, ha 구성은 db를 기반으로 설계<br>
+ IPMI (IDRAC, ILO, IMM..) 지원<br>
+ GUI, CLI, API 제공

![maas_arch01][https://raw.githubusercontent.com/engineer-pjin/sre_component_foundation/master/image/maas_arch01.png]

<br>

### + Life Cycle

단계 | 세부 
:---:|:---:
New | 새로운 서버 등록, pxe boot를 위한 mac을 등록하며 BMC parameters를 통해 자동 등록 가능
Commissioning| 첫 부팅 시 pxe를 통해 임시 ubuntu os를 올리고 정보를 수집 
Ready | Commissioning이 정상적으로 완료되어 제어 가능 상태 
Allocated | network(boinding, addressing..) / disk(lvm, raid..) 설정 가능
Deploying | os 배포 단계
Releasing | 서버 사용 완료 후 공유 풀로 반환, 완료 후에는 ready 상태로 변환 됨

> 노드를 등록하고 PXE 부팅으로 들어가면 dhcp 범위내에서 ip를 받아와서 최소 커널로 메모리를 올려 부팅시키고 정해진 레포에서 패키지 가져와서
Commissioning 를 위한 정보를 가져오고 다시 꺼짐 <br>
> 해당 노드의 정보를 다가지고 오고 노드 리스트에서 상태가 레디가 되면 상세에 들어가서 take action의 deploy로 변경 - 이미지 선택 <br>
> 상태가 Deploying Ubuntu 16.04 LTS가 되고 여기서 pxe 부팅을 하면 os 설치됨 <br>

> 즉 기본 구성은 서버 구매 후 Commissioning을 통해 등록, 사용할 때는 ipmi 혹은 수동으로 pxe 부팅을 통해 OS를 설치하는 단계로 진행

<br>

### + 설치
```
# apt-add-repository ppa:maas/stable -y
# apt update
# apt install maas -y
# maas createadmin
Username: jin
Password:
Again:
Email: jin@jin.jin
Import SSH keys [] (lp:user-id or gh:user-id): root
SSH import protocol was not entered.  Using Launchpad protocol (default).
Unable to import SSH keys. There are no SSH keys for launchpad user root.
```
> gui 확인 : http://[ip]/MAAS