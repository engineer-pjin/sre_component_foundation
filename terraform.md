## terraform
## 기본
### + 특성
```
- Infrastructure as Code
- server provisioning tool, written in Golang
- HCL로 리소스를 선언
- 인프라스트럭처를 선언적인 코드로 작성
- Support Multi Cloud Provider(AWS, GCP, OpenStack, Vmware...)
```
> 참조 : [MAAS 홈페이지](https://www.terraform.io/)

### + 명세: https://www.terraform.io/docs/providers/openstack/
### + 배포 단계
```
스텝 1 - HCL(Hashicorp Configuration Language) 언어로 필요한 리소스를 선언
   >> 테라폼은 기본적으로 특정 디렉터리에 있는 모든  .tf 확장자를 가진 파일을 읽어들인 후, 리소스 생성, 수정, 삭제 작업을 진행
스텝 2 - 선언된 리소스들이 생성가능한지 계획(Plan)을 확인
스텝 3 - 선언된 리소스들을 프로바이더를 통해 적용(Apply)
```
### + 파일 상세
```
openstack_provider.tf : provider description
vars.tf : 변수 description
terraform.tfstate & .backup : terraform 의 리소스 description (Do not edit)
**.tf : 선언되는 리소스
```
> 다수 인원 혹은 terraform을 backend로 둔 platform을 운영한다면 terraform.tfstate은 별도 저장소에서 공유되어야 함

<br><br>

## 인프라 구성 정보
> ver 20190617

### + host list : com01, com02
### + Hardware spec : 4core 8G
### + az : nova
### + ip range
```
  192.168.18.21 ~ 100 /24
```
### + image
```
  Ubuntu16
  cirros
```
### + flavor
```
  1v_1G_5G
  2v_2G_5G
```
### + network : test-net

<br><br>

## 실습
### + 설치
```
# wget https://releases.hashicorp.com/terraform/0.11.11/terraform_0.11.11_linux_amd64.zip
# unzip terraform_0.11.11_linux_amd64.zip
# ls
terraform  terraform_0.11.11_linux_amd64.zip
 > 바이너리 파일 하나만 있음
# cp terraform /bin/

```

### + 파일 상세
+ openstack_provider.tf
```
provider "openstack" {
  user_name   = "**"
  tenant_name = "**"
  password    = "**"
  auth_url    = "http://**:5000/v2.0"
  region      = "RegionOne"
}
```
> **# terraform init** 으로 초기화 

+ vars.tf
```
variable "network01" {
  type = "string"
  default = "test-net"
}

variable "az01" {
  type = "string"
  default = "nova"
}


variable "compute_lists" {
  type = "map"
  default = {
    compute_01 =  "com01"
    compute_02 =  "com02"
  }
}
```

+ test_instance.tf
```
variable "test_instance_ips" {
  description   = "test instance ip list"
  type          = "list"
  default       = [
  "192.168.18.22",
  "192.168.18.23"
  ]
}

resource "openstack_compute_instance_v2" "test01" {
  name              = "test1"
  image_name        = "ubuntu16"
  flavor_name       = "2v_2G_5G"
  availability_zone = "${var.az01}:${lookup(var.compute_lists, "compute_01")}"

  network {
    name        = "${var.network01}"
    fixed_ip_v4 = "192.168.18.21"
  }
}

output "test01 names" {
  value = "${openstack_compute_instance_v2.test01.*.name}"
}

resource "openstack_compute_instance_v2" "test_instance" {
  count             = "${length(var.test_instance_ips)}"
  name              = "test${count.index+2}"
  image_name        = "ubuntu16"
  flavor_name       = "2v_2G_5G"
  availability_zone = "${var.az01}:${lookup(var.compute_lists, "compute_01")}"

  network {
    name        = "${var.network01}"
    fixed_ip_v4 = "${element(var.test_instance_ips, count.index)}"
  }
}

output "test_instance names" {
  value = "${openstack_compute_instance_v2.test_instance.*.name}"

```

### + PLAN
```
# terraform plan
~
Plan: 3 to add, 0 to change, 0 to destroy.
~
```

### + APPLY
```
# terraform apply
~
Plan: 3 to add, 0 to change, 0 to destroy.               
                                                         
Do you want to perform these actions?                    
  Terraform will perform the actions described above.    
  Only 'yes' will be accepted to approve.                
                                                         
  Enter a value: <여기에 yes 입력 시 인프라에 적용 됨>
~
Apply complete! Resources: 3 added, 0 changed, 0 destroyed.       
                                                                  
Outputs:                                                          
                                                                  
test01 names = [                                                  
    test01                                                        
]                                                                 
test_instance names = [                                           
    test2,                                                        
    test3                                                         
]                                                                                                     
```