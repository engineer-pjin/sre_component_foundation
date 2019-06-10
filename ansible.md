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
log_path = ansible_log
host_key_checking = False

forks = 512

stdout_callback = skippy
timeout=129

[ssh_connection]
pipelining = True 
 > sudo 를 통한 테스트시에는 False 
```

## 실습
> **ansible directory 구성 요소**<br>
>> ansible.cfg : 지역 설정<br>
>> group_vars/all.yml : 그룹 변수 <br>
>> handlers/all.yml : handlers 설정 <br>
>> hosts/all : 호스트 리스트 <br>
>> sh/ : 스크립트 디렉토리 <br>
>> tasks/ : tasks 디렉토리 <br>
>> templates. : templates 디렉토리 <br>


**ansible**
```
- ssh key 등록 후
# vi hosts
test01a ansible_host=127.0.0.1 ansible_port=22 ansible_user=root
test02a ansible_host=<ip1> ansible_port=22 ansible_user=root
test03a ansible_host=<ip2> ansible_port=22 ansible_user=root


# ansible all -i hosts -m ping
vagrant-machine | SUCCESS => {
    "changed": false,
    "ping": "pong"
}
>> 노드들에 python2가 없으면 아래와 같은 에러 발생
    "changed": false,
    "module_stderr": "Shared connection to *** closed.\r\n",
    "module_stdout": "/bin/sh: 1: /usr/bin/python: not found\r\n",
    "msg": "MODULE FAILURE\nSee stdout/stderr for the exact error",
    "rc": 127

# ansible all -i hosts -m ping --list-hosts
> 실행 전 적용되는 노드들 확인

# ansible all -i hosts -m copy -a "src=test.txt dest=/root/"
# ansible all -i hosts -m shell -a "ls /root/"

# ansible all -i hosts -m apt -a "name=iptraf state=present"

# ansible-playbook -i hosts site.yml
# ansible-playbook -i hosts site.yml -l ops_cont --list-hosts
# ANSIBLE_DEBUG=y ansible-playbook -i hosts site.yml

# ansible all -i hosts -m shell -a uptime
# ansible all -i hosts -m shell -a "systemctl status ntp|grep -i active"


> ansible에서는 -C (or --check) 옵션을 포함하여 간단하게 수행이 가능하다 
> 모듈을 쓰지 않고 직접 쉘명령을 사용하는경우 주의해야 한다.
> 모듈은 테스트가 가능하지만 쉘명령은 실행 되면 그냥 반영되버릴 수 있다.)

# ansible testserver -C -m copy -a "src=/tmp/file.txt dest=/tmp"   <== 되는지 테스트
# ansible testserver -m copy -a "src=/tmp/file.txt dest=/tmp" <==실제 복사
```

<br>

**anssible playbook**
```
- fact 수집
> 파이썬이 없는 호스트 같이 하단에 셋업에 대한 부분은 점검하는 로직이 있음
> 이부분은 false 처리하면 속도는 빨라짐.. 

# vi site.yml
---
- name: 플레이북 튜토리얼
  hosts: all
  gather_facts: false  ## 셋업 환경 점검 여부
  become: true  
  tasks:

> 셋업값 확인
# ansible all -i hosts -m setup


- become
---
- name: 플레이북 튜토리얼
  hosts: all
  gather_facts: false  
  become: true  ## 관리자 권한 획득 여부 
  tasks:

> 권장은 플레이북에는 become:true, 필요에 따라 호스트/그룹 쪽에서 ansible_become_method 사용


- task : apt
  tasks:
    - name: iptraf-ng 설치
      apt:		  ## 모듈 네임 : ansible-doc -l
        name: iptraf-ng   ## 모듈 arg : value
        state: present

> state : present(default & 설치된 상태 installed), absent(미설치 상태 removed), latest(최신버전이 설치된 상태)
> task_directive : 테스크를 실행 단위별로 지정할 수 있는 지시자
 >> 테스크 리턴값을 변수에 설정, 루프 설정, 성공/실패 룰의 변경, 실행 조건의 설정


- service 모듈
    - name: 엔진엑스 설치
      apt:
        name: nginx
        state: present

    - name: 엔진엑스 서비스 시작과 자동 시작 설정
      service:
        name: nginx
        state: started
        enabled: true

> 서비스 모듈의 state : started, stoped, restard, reloaded
 >> systemd와 정상적으로 연동 됨을 확인

# ansible-doc [모듈 이름]
> 모듈 상세

- PLAY RECAP 
> ok : 상태가 수정되지 않았거나 부작용이 없는지 확인
> changed : 상태를 변경하는 처리가 실행
> failed : 처리에 실패
- unreachable : 작업 대상에 접속을 실패


- file 모듈
    - name: /tmp/dir1 create
      file:
        path: /tmp/dir1
        state: directory
> state 
 >> directory : 존재하지 않는다면 새로 생성
 >> file : 파일 속성 변경, 파일이 없으면 에러
 >> touch : 파일이 없으면 생성, 있으면 타임스태프 변경
 >> link & hard : 심볼릭 & 하드 링크 생성
 >> absent : 패스가 존자하면 파일과 디렉토리 삭제

    - name: 권한
      file:
        path: /tmp/dir1
        state: directory
        owner: ujin
        group: gjin
        mode: "u=rwx,g=rwx,o=r"
        recurse: true		## R옵션

https://docs.ansible.com/ansible/latest/modules/copy_module.html
- copy 모듈
    - name: 원격으로 파일 복사
      copy:
        src: /tmp/dir1/a
        dest: /tmp/dir1/1
        force: false	  ## 덮어 쓰기 안함
	backup: false	  ## 파일 변경 시 백업 파일 생성 안함(default)
> 디렉토리 자체를 복사하는 옵션이 있으나 파일 하나하나씩 검증하기 때문에 시간이 걸림
> src 는 ansible server, dest는 guest

- name: node expoter
  hosts: all
  gather_facts: false
  become: true
  tasks:
    - name: exporter copy
      copy:
        src: /root/wmp/ansible/site.retry
        dest: /root/
        owner: root
        group: root
        mode: 0755
        backup: yes

- user 
https://docs.ansible.com/ansible/latest/modules/user_module.html
    - name: group add
      group:
        name: test01
        gid: 1101

    - name: user add
      user:
        name: test01
        uid: 1101
        group: test01
        append: yes
        shell: /bin/bash
        home: /home/ttt01
        generate_ssh_key: yes
        ssh_key_bits: 2048
        ssh_key_file: .ssh/id_rsa


- lineinfile - 파일을 행단위로 수정, 일치하는게 없으면 마지막에 추가됨
    - name: 파일 라인 변경
      lineinfile:
        dest: /etc/ssh/sshd_config
        regexp: '^PermitRootLogin no\s+'
        line: PermitRootLogin yes
        validate: sshd -t -f %s		## 테스트, 반환코드가 0이 될때 검증에 성공한 것으로 판단하고 파일의 내용 변경, %s는 임시파일
      notify:				## 테스크의 실행결과가 change인 경우 지정된 이름의 핸들러를 테스크 동작 맨 마지막에 실행
        - sshd restart

  handlers:
    - name: sshd restart
      service:
        name: sshd
        state: restarted

>> 패턴와 일치하면 여러행을 치환하는 replace, 블록단위 치환을 하는 blockinfile 도 테스트

- blockinfile : 일치하는 내용이 없으면 추가


- command : 임의의 명령어 실행
> 명령을 실행한 후 종료상태가 0일때 성공 그외에는 실패
> 성공 시 실행 결과는 항상 changed

    - name : command 테스트
      command: "/usr/bin/ssh-keygen -b 2048 -t rsa -N '' -f /tmp/new-id_rsa"
      args:
        creates: /tmp/new-id_rsa
> args
 >> create 명령이 실행된 후에 생성될 파일의 경로를 지정, 파일이 존재하지 않을 경우에만 생성
 >> remove : 명령이 실행된 후에 삭제될 파일의 경로를 지정, 파일이 존재하는 경우만 삭제
 >> chdir : 명령을 실행할 때 기점이 되는 디렉토리를 지정
 >> executable : 명령을 실행할 때 사용될 셸의 경로, 지정하지 않으면 로그인한 사용자의 쉘 사용

> 제한사항 : 파이프와 리다이렉트 사용불가, $ 기호를 사용한 환경변수는 참조할 수 없음
> changed_when, failed_when : 실행된 결과를 재기록하기 위한 지시자
> 쉘의 기능을 사용하고 싶을때는 shell 모듈을 사용, 변수 사용시는 quote 필터 이용
    - name : 환경 변수 출력
      command: "echo {{ ansible_env.HOME | quote }}"
 >> gather_facts: true 인상태에서 진행해야함.. 즉 setup task 에서 ansible에 설정된 환경 값임
 >> 거의 비슷한 동작을 하는 shell 모듈이 있으나 멱등성이 보장이 안됨

>> env : 실시간 설정 변수는 ansible_env에서 안됨.. 아래 내용은 에러남
    - name:
      command: "echo {{ ansible_env.TEST | quote }}"
      environment:
        TEST: "imsy var"



- Command Line로 변수 전달
--extra-vars 옵션을 이용하여 실행시간에 CLI로 변수 값 전달 가능함.

---

- hosts: '{{ hosts }}'
  remote_user: '{{ user }}'

  tasks:
     - ...

ansible-playbook release.yml --extra-vars "hosts=vipers user=starbuck"

# 또는 JSON 형식으로
ansible-playbook release.yml --extra-vars \
  '{"hosts": "vipers", "user=starbuck", "ghosts":["inky","pinky","clyde","sue"], release: 1}'
단, 변수의 type을 살리고 싶으면 반드시 JSON 형식으로 전달할 것. 기본 key=value 형식은 모든 value가 string으로 인식됨.

@를 사용하여 JSON 파일로부터 바로 읽을 수도 있음(1.3 이후)
# ansible-playbook release.yml --extra-vars "@some_file.json"
> yml에서 직접 읽을 수도 있음
# ansible-playbook -i hosts -e '@extra-vars.yml' site.yml

# ansible-playbook -i hosts -e 'ntp_dst_ip=10.107.1.10 ntp_src_ip=non' site.yml
 > 인벤토리에 정의된 내용보다 우선함, 모든 값은 문자열이 됨, 수치나 부울값을 하려면 json 사용
# ansible-playbook -i hosts -e '{"ansible_port": 22}' site.yml
 > 인벤토리에 정의된 내용보다 우선함

- group_vars : 우선 순위가가장 낮아 기본변수를 지정하는데 유용, all < group << host 변수
http://theeye.pe.kr/archives/tag/group_vars
# cat group_vars/all.yml ## all
# cat group_vars/test02.yml ## group
--
ntp_dst_ip: <ip>
ansi		
# cat host_vars/test08a ## host

# cat hosts
~
[test02] ## host
test06a ansible_host=<ip> ntp_dst_ip=<ip>

[test02:vars] ## group
ntp_dst_ip=<ip>
>> 우선순위는 호스트 변수가 높기 때문에 기본은 그룹 변수로 두고 호스트 변수로 덮어쓰기

|-- group_vars
|   |-- all.yml    ## 3순위
|   `-- test02.yml ## 2순위
|-- host_vars
|   `-- test07a.yml ## 1순위

>> 인벤토리 파일의 변수 < 인벤터리 변수를 정의한 yarm < 플레이북의 변수를 정의한 yarm 파일
>> 위의 우선순위 안에서 호스트/그룹 변수의 순위가 나눠짐 : 테스트 필요
>> 변수는 알파벳 대소문자, 숫자, 언더바만 가능



- 동적 인벤토리 : 오픈 스택 및 베어그란트 별도 테스트, '엔서블 철저 입문:62'
> iaas(aws, openstack), monitoring, vagrant 등등
https://docs.openstack.org/openstack-ansible/newton/developer-docs/inventory.html
>> 앤서블 명령을 실행할 때 실시간으로 스크립트가 실행
>> 스크립트에서 외부 시스템에 접근할 때 동적으로 호스트 정보를 가져옴



```



