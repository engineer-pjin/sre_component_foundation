# Ansible 기초
## 개념
### + 기본
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

### + 상세
#### 기본 동작 : playbook에 정의된 대로 inventory 대상으로 module들을 실행 <br>

#### 구성 요소 
```
- Inventory : 조작대상이 되는 서버 접속 정보를 표시하는 정의 with dynamic inventory
- module :  실행되는 동작 (작업을 실행하기전 상태를 확인하고 변경이 있을때만 동작 처리)
- playbook : yaml 형식의 스크립트(코드), yaml은 프로그래밍 언어가 아닌 데이터의 표현 형식
```
<br>

#### 재사용 : 변수 & 롤(role - 플레이북을 각 시스템에서 공통으로 사용하는 단위로 분리한 것) <br>

```
playbook 내에서 변수를 참조할 때, Jinja2 사용
```
<br>

#### install with ubuntu 16.04<br>
> ansible 의 릴리즈 주기는 보통 4개월이어서 개발 및 테스트를 위해서는 pip를, 안정적인 운영을 위해서는 패키지 설치를 권장 <br>
```
## pip 방식
# wget https://bootstrap.pypa.io/get-pip.py 
# python2 get-pip.py
~
Installing collected packages: pip, setuptools, wheel
Successfully installed pip-18.1 setuptools-40.4.3 wheel-0.32.1

# pip2 install ansible==2.7.0
~
Successfully built ansible pycparser
paramiko 2.4.2 has requirement cryptography>=1.5, but you'll have cryptography 1.2.3 which is incompatible.
Installing collected packages: pycparser, cffi, bcrypt, pynacl, paramiko, ansible
Successfully installed ansible-2.7.0 bcrypt-3.1.4 cffi-1.11.5 paramiko-2.4.2 pycparser-2.19 pynacl-1.3.0

# pip list

# ansible-doc -l
```

<br>


#### 설정<br>
+ ANSIBLE_CONFIG 환경변수에서 지정한 파일
+ 현재 디렉토리에 있는 ansible.cfg   
+ 사용자 홈 디렉토리 아래의 .ansible.cfg, 사용자 레벨의 기본설정
+ /etc/ansible/ansible.cfg  (글로벌 기본설정)
+ 테스트 간 추천 설정
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
>> tasks/ : tasks 디렉토리 <br>
>> templates. : templates 디렉토리 <br>

<br>

### + ansible
```
+ ssh key 등록 후
# vi hosts/all
test01a ansible_host=127.0.0.1 ansible_port=22 ansible_user=root
test02a ansible_host=<ip1> ansible_port=22 ansible_user=root
test03a ansible_host=<ip2> ansible_port=22 ansible_user=root

+ ping module 이용하여 확인
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

+ 실행 전 적용되는 노드들 확인
# ansible all -i hosts -m ping --list-hosts

+ module test 
# ansible all -i hosts -m copy -a "src=test.txt dest=/root/"
# ansible all -i hosts -m shell -a "ls /root/"

# ansible all -i hosts -m apt -a "name=iptraf-ng state=present"

# ansible all -i hosts -m shell -a uptime
# ansible all -i hosts -m shell -a "systemctl status ntp|grep -i active"


> ansible에서는 -C (or --check) 옵션을 포함하여 간단하게 수행이 가능하다 
> 모듈을 쓰지 않고 직접 쉘명령을 사용하는경우 주의해야 한다.
> 모듈은 테스트가 가능하지만 쉘명령은 실행 되면 그냥 반영되버릴 수 있다.)

# ansible testserver -C -m copy -a "src=/tmp/file.txt dest=/tmp"   <== 되는지 테스트
# ansible testserver -m copy -a "src=/tmp/file.txt dest=/tmp" <==실제 복사
```

<br>

### + ansible playbook
> fact 수집
>> 파이썬이 없는 호스트 같이 하단에 셋업에 대한 부분은 점검하는 로직이 있음
>> 이부분은 false 처리하면 속도는 빨라짐.. 

```
# ansible-playbook -i hosts site.yml
# ansible-playbook -i hosts site.yml -l ops_cont --list-hosts
# ANSIBLE_DEBUG=y ansible-playbook -i hosts site.yml
```
<br>

#### site.yml
```
---
- name: 플레이북 튜토리얼
  hosts: all
  gather_facts: false  ## 셋업 환경 점검 여부
  become: true  ## 관리자 권한 획득 여부
  tasks:

> 셋업값 확인
# ansible all -i hosts -m setup
```
> 권장은 플레이북에는 become:true, 필요에 따라 호스트/그룹 쪽에서 ansible_become_method 사용

<br>


#### PLAY RECAP 
> ok : 상태가 수정되지 않았거나 부작용이 없는지 확인<br>
> changed : 상태를 변경하는 처리가 실행<br>
> failed : 처리에 실패<br>
> unreachable : 작업 대상에 접속을 실패<br>

<br>

#### ansible-doc [모듈 이름]
> 모듈 상세 <br>
> 참조 사이트 : https://docs.ansible.com/ansible/latest/index.html

<br>

#### task : apt
> https://docs.ansible.com/ansible/latest/modules/apt_module.html?highlight=apt

```
  tasks:
    - name: iptraf-ng 설치
      apt:		  ## 모듈 네임 : ansible-doc -l
        name: iptraf-ng   ## 모듈 arg : value
        state: present
```
> state 
>> present(default & 설치된 상태 installed)<br>
>> absent(미설치 상태 removed)<br>
>>latest(최신버전이 설치된 상태)<br>

> task_directive : 테스크를 실행 단위별로 지정할 수 있는 지시자
 >> 테스크 리턴값을 변수에 설정, 루프 설정, 성공/실패 룰의 변경, 실행 조건의 설정

<br>

#### task : service
> https://docs.ansible.com/ansible/latest/modules/service_module.html?highlight=service
```
    - name: 엔진엑스 설치
      apt:
        name: nginx
        state: present

    - name: 엔진엑스 서비스 시작과 자동 시작 설정
      service:
        name: nginx
        state: started
        enabled: true
```
> 서비스 모듈의 state : started, stoped, restard, reloaded

<br>

#### task : file
> https://docs.ansible.com/ansible/latest/modules/file_module.html?highlight=file
````
    - name: /tmp/dir1 create
      file:
        path: /tmp/dir1
        state: directory
````
> state 
 >> directory : 존재하지 않는다면 새로 생성<br>
 >> file : 파일 속성 변경, 파일이 없으면 에러<br>
 >> touch : 파일이 없으면 생성, 있으면 타임스태프 변경<br>
 >> link & hard : 심볼릭 & 하드 링크 생성<br>
 >> absent : 패스가 존자하면 파일과 디렉토리 삭제<br>


```
    - name: 권한
      file:
        path: /tmp/dir1
        state: directory
        owner: ujin
        group: gjin
        mode: "u=rwx,g=rwx,o=r"
        recurse: true		## R옵션
```

<br>

#### task : copy
> https://docs.ansible.com/ansible/latest/modules/copy_module.html?highlight=copy
````
    - name: 원격으로 파일 복사
      copy:
        src: /tmp/dir1/a
        dest: /tmp/dir1/1
        force: false	  ## 덮어 쓰기 안함
	      backup: false	  ## 파일 변경 시 백업 파일 생성 안함(default)
````
> 디렉토리 자체를 복사하는 옵션이 있으나 파일 하나하나씩 검증하기 때문에 시간이 걸림<br>
> src 는 ansible server, dest는 guest<br>

<br>

#### task : user 
> https://docs.ansible.com/ansible/latest/modules/user_module.html

```
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

```

<br>

#### task : command 
> 임의의 명령어 실행<br>
> 명령을 실행한 후 종료상태가 0일때 성공 그외에는 실패 <br>
> 성공 시 실행 결과는 항상 changed<br>
```
    - name : command 테스트
      command: "/usr/bin/ssh-keygen -b 2048 -t rsa -N '' -f /tmp/new-id_rsa"
      args:
        creates: /tmp/new-id_rsa
```
> args
 >> create 명령이 실행된 후에 생성될 파일의 경로를 지정, 파일이 존재하지 않을 경우에만 생성<br>
 >> remove : 명령이 실행된 후에 삭제될 파일의 경로를 지정, 파일이 존재하는 경우만 삭제<br>
 >> chdir : 명령을 실행할 때 기점이 되는 디렉토리를 지정<br>
 >> executable : 명령을 실행할 때 사용될 셸의 경로, 지정하지 않으면 로그인한 사용자의 쉘 사용<br>
 
 <br>

> 제한사항 : 파이프와 리다이렉트 사용불가, $ 기호를 사용한 환경변수는 참조할 수 없음<br>
> changed_when, failed_when : 실행된 결과를 재기록하기 위한 지시자<br>
> 쉘의 기능을 사용하고 싶을때는 shell 모듈을 사용, 변수 사용시는 quote 필터 이용<br>
> 거의 비슷한 동작을 하는 shell 모듈이 있으나 멱등성이 보장이 안됨<br>

<br>

#### handlers - task 동작 중 notify 시 실행
```
- name: ntp temp
  template:
    src: ntp.conf.j2
    dest: /etc/ntp.conf
  notify:
    - ntp restart
```
```
  handlers:
  - name: ntp restart
    service:
      name: ntp
      state: restarted

```

<br>


#### 변수
- Command Line로 변수 전달
> --extra-vars 옵션을 이용하여 실행시간에 CLI로 변수 값 전달 가능함.
```
- hosts: '{{ hosts }}'
  remote_user: '{{ user }}'

  tasks:
     - ...
```

```
 ansible-playbook release.yml --extra-vars "hosts=vipers user=starbuck"

# 또는 JSON 형식으로
ansible-playbook release.yml --extra-vars \
  '{"hosts": "vipers", "user=starbuck", "ghosts":["inky","pinky","clyde","sue"], release: 1}'
```
> 변수의 type을 유지하려면 JSON 형식으로 해야 하며 key/value 형식은 string으로 인식<br>
> @를 사용하여 JSON 파일 사용 가능<br>
```
# ansible-playbook release.yml --extra-vars "@some_file.json"
# ansible-playbook -i hosts -e '@extra-vars.yml' site.yml

# ansible-playbook -i hosts -e 'ntp_dst_ip=[ip] ntp_src_ip=non' site.yml
 > 인벤토리에 정의된 내용보다 우선함, 모든 값은 문자열이 됨, 수치나 부울값을 하려면 json 사용
# ansible-playbook -i hosts -e '{"ansible_port": 22}' site.yml
 > 인벤토리에 정의된 내용보다 우선함
```

<br>

- group_vars : 우선 순위가가장 낮아 기본변수를 지정하는데 유용, all < group << host 변수
```
|-- group_vars
|   |-- all.yml    ## 3순위
|   `-- test02.yml ## 2순위
|-- host_vars
|   `-- test07a.yml ## 1순위
```
> 인벤토리 파일의 변수 < 인벤터리 변수를 정의한 yarm < 플레이북의 변수를 정의한 yarm 파일
> 위의 우선순위 안에서 호스트/그룹 변수의 순위가 나눠짐 : 테스트 필요
> 변수는 알파벳 대소문자, 숫자, 언더바만 가능

<br>

- 동적 인벤토리 
> iaas(aws, openstack), monitoring, vagrant 등등 <br>
https://docs.openstack.org/openstack-ansible/newton/developer-docs/inventory.html
>> 앤서블 명령을 실행할 때 실시간으로 스크립트가 실행
>> 스크립트에서 외부 시스템에 접근할 때 동적으로 호스트 정보를 가져옴





