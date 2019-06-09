# Ansible 기초
## 1 개념
### 1-1 기본
- Automation Tool
- Push Type : Agentles
- Configuraion as Code & Infrastructure as Code
- devops & CI/CD
- **멱등성 (Idempotency)** 
```
대상의 상태는 어때야 한다를 정의
어떤 작업을 여러번 실행해도 결과가 항상 같음
ex) '파일 a를 어디에 복사한다'가 아닌 '파일 a가 어디에 있다'
```

<br>

### 1-2 상세
**1) 기본 동작** : playbook에 정의된 대로 inventory 대상으로 module들을 실행 

**2) 구성 요소**
```
- Inventory : 조작대상이 되는 서버 접속 정보를 표시하는 정의 with dynamic inventory
- module :  실행되는 동작 (작업을 실행하기전 상태를 확인하고 변경이 있을때만 동작 처리)
- playbook : yaml 형식의 스크립트(코드), yaml은 프로그래밍 언어가 아닌 데이터의 표현 형식
```

**3) 재사용 : 변수 & 롤(role - 플레이북을 각 시스템에서 공통으로 사용하는 단위로 분리한 것)** 
```
playbook 내에서 변수를 참조할 때, Jinja2 사용
```
<br>

**4) install with ubuntu 16.04** <br>
ansible 의 릴리즈 주기는 보통 4개월이어서 개발 및 테스트를 위해서는 pip를, 안정적인 운영을 위해서는 패키지 설치를 권장 <br>
```
## 패키지 방식으로 설치
# apt-get update
# apt-get install software-properties-common
# apt-add-repository ppa:ansible/ansible
 >> https가 오픈되어야 함.. 아니면 "ERROR: '~ansible' user or team does not exist"
# apt-get update
# apt-get install ansible
# ansible --version
ansible 2.7.0
  config file = /etc/ansible/ansible.cfg
  configured module search path = [u'/root/.ansible/plugins/modules', u'/usr/share/ansible/plugins/modules']
  ansible python module location = /usr/lib/python2.7/dist-packages/ansible
  executable location = /usr/bin/ansible
  python version = 2.7.12 (default, Dec  4 2017, 14:50:18) [GCC 5.4.0 20160609]

# ansible-doc -l
```
**5) 설정**<br>
- ANSIBLE_CONFIG 환경변수에서 지정한 파일
- 현재 디렉토리에 있는 ansible.cfg   
- 사용자 홈 디렉토리 아래의 .ansible.cfg, 사용자 레벨의 기본설정
- /etc/ansible/ansible.cfg  (글로벌 기본설정)
- 테스트 간 추천 설정
```
[defaults]
host_key_checking = False

forks = 128

stdout_callback = debug

timeout=30

[ssh_connection]
pipelining = True 
 > sudo 를 통한 테스트시에는 False 
```

## 실습
