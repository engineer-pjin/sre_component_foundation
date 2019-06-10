# prometheus
## 개념
> https://prometheus.io/docs/introduction/overview/

- SoundCloud 에서 제작 된 오픈 소스 시스템 모니터링, golang 
- 2016년에 Kubernetes에 이어 두 번째 프로젝트로 CNCF(Cloud Native Computing Foundation)에 가입
- HTTP full 방식으로 각 client 노드에서 export로부터 데이터를 수집하여 시계열 저장
- 각 서비스는 의존적이지 않으며 분산 구성에 적합(prometheus 서버끼리 연결을 통한 확장)
- exporter에서 서버가 데이터를 수집해서 저장하고 그라파나를 통해 보여주며, alertmanager를 통해 알람을 
![prometheus Architecture](https://raw.githubusercontent.com/engineer-pjin/sre_component_foundation/master/image/prometheus_architecture.png)

> 대표적인 export<br>
> node_exporter(기본 os 메트릭), mysqld_exporter, rabbitmq_exporter, nginx-vts-exporter...

<br>

## 구성
### server
> 패키지로 설치하면 테스트 당시 버전이 server version 0.16.2+ds, 소스버전은 2.4.3, 소스 설치 진행<br>

```
# mkdir /etc/prometheus /var/lib/prometheus
# cd /var/lib/prometheus
# curl -LO https://github.com/prometheus/prometheus/releases/download/v2.4.3/prometheus-2.4.3.linux-amd64.tar.gz
# sha256sum prometheus-2.4.3.linux-amd64.tar.gz
3aa063498ab3b4d1bee103d80098ba33d02b3fed63cb46e47e1d16290356db8a  prometheus-2.4.3.linux-amd64.tar.gz
# tar xvf prometheus-2.4.3.linux-amd64.tar.gz
# chown -R root.root ./*

# cp /var/lib/prometheus/prometheus-2.4.3.linux-amd64/prometheus /usr/bin/
# cp /var/lib/prometheus/prometheus-2.4.3.linux-amd64/promtool /usr/bin/
# cp -r prometheus-2.4.3.linux-amd64/consoles /etc/prometheus/
# cp -r prometheus-2.4.3.linux-amd64/console_libraries/ /etc/prometheus/
# vi /etc/prometheus/prometheus.yml
global:
  scrape_interval: 15s

scrape_configs:
  - job_name: 'prometheus'
    scrape_interval: 5s
    static_configs:
      - targets: ['localhost:80']

# prometheus \
    --config.file /etc/prometheus/prometheus.yml \
    --storage.tsdb.path /var/lib/prometheus/ \
    --web.console.templates=/etc/prometheus/consoles \
    --web.console.libraries=/etc/prometheus/console_libraries \
    --web.listen-address=:80
>> 정상적으로 동작하는지 확인 후 문제가 없다면 서비스 등록

# vi /etc/systemd/system/prometheus.service
[Unit]
Description=Prometheus
Wants=network-online.target
After=network-online.target

[Service]
User=root
Group=root
Type=simple
ExecStart=/usr/bin/prometheus \
    --config.file /etc/prometheus/prometheus.yml \
    --storage.tsdb.path /var/lib/prometheus/ \
    --web.console.templates=/etc/prometheus/consoles \
    --web.console.libraries=/etc/prometheus/console_libraries \
    --web.listen-address=:80 \
    --web.enable-admin-api

[Install]
WantedBy=multi-user.target

# systemctl daemon-reload
# systemctl enable prometheus.service
# systemctl start prometheus.service
# systemctl status prometheus.service
# netstat -antp |grep -i prome

```


### node exporter
```
# vi /etc/systemd/system/prometheus.service
[Unit]
Description=Prometheus
Wants=network-online.target
After=network-online.target

[Service]
User=root
Group=root
Type=simple
ExecStart=/usr/bin/prometheus \
    --config.file /etc/prometheus/prometheus.yml \
    --storage.tsdb.path /var/lib/prometheus/ \
    --web.console.templates=/etc/prometheus/consoles \
    --web.console.libraries=/etc/prometheus/console_libraries \
    --web.listen-address=:80 \
    --web.enable-admin-api

[Install]
WantedBy=multi-user.target

# systemctl daemon-reload
# systemctl enable prometheus.service
# systemctl start prometheus.service
# systemctl status prometheus.service
# netstat -antp |grep -i prome

> 테스트는 웹에서 http://<ip>:9100으로 확인 후 정상적으로 데이터가 보이는 지 확인 후 서버에 노드 추가
## 서버에서
# vi /etc/prometheus/prometheus.yml
...
scrape_configs:
...
- job_name: 'node_exporter'
    scrape_interval: 5s
    static_configs:
      - targets: ['<ip>:9100']
      
```
> https://prometheus.io/docs/prometheus/latest/configuration/configuration/#%3Cconsul_sd_config%3E<br>
> CONFIGURATION 문서를 보면 모니터링 대상을 Consul, DNS 기반의 서비스 디스커버리를 이용해서 동적으로 설정 가능<br>


### alertmanager
> 서버와 동일한 노드에서 진행한다고 가정<br>
> 기본 node-exporter의 동작여부 확인 후 <br>

```
# cd /var/lib/prometheus
# curl -LO https://github.com/prometheus/alertmanager/releases/download/v0.15.2/alertmanager-0.15.2.linux-amd64.tar.gz
# tar xvf alertmanager-0.15.2.linux-amd64.tar.gz
# chown -R root.root ./alertmanager*
# cd alertmanager-0.15.2.linux-amd64/

# cp alertmanager /usr/bin/
# cp amtool /usr/bin/
# cp alertmanager.yml /etc/prometheus/
# vi /etc/systemd/system/alertmanager.service
[Unit]
Description=Prometheus_alertmanager
Wants=network-online.target
After=network-online.target

[Service]
User=root
Group=root
Type=simple
ExecStart=/usr/bin/alertmanager \
    --config.file /etc/prometheus/alertmanager.yml

[Install]
WantedBy=multi-user.target

# vi /etc/prometheus/prometheus.yml
~
rule_files:
  - "alerts.yml"

alerting:
  alertmanagers:
  - static_configs:
    - targets:
      - localhost:9093
~

# vi /etc/prometheus/alerts.yml
groups:
  - name: default
    rules:
    - alert: InstanceDown
      expr: up == 0
      for: 1m
      labels:
        severity: critical
      annotations:
        summary: "Instance {{ $labels.instance }} down"
        description: "{{ $labels.instance }} of job {{ $labels.job }} has been down for more than 1 minutes."

# vi /etc/prometheus/alertmanager.yml
global:
  resolve_timeout: 5m
  smtp_smarthost : '***'
  smtp_from : 'sysadmin@***.com'

route:
  group_by: ['alertname']
  group_wait: 10s
  group_interval: 10s
  repeat_interval: 1h
  receiver: 'mail'

receivers:
- name: 'mail'
  email_configs:
  - to: 'pjin@***.com'

inhibit_rules:
  - source_match:
      severity: 'critical'
    target_match:
      severity: 'warning'
    equal: ['alertname', 'dev', 'instance']

# systemctl daemon-reload
# systemctl enable alertmanager.service
# systemctl start alertmanager.service
```

### grafana
```
# vi /etc/apt/sources.list
deb https://packagecloud.io/grafana/stable/debian/ jessie main

curl https://packagecloud.io/gpg.key | sudo apt-key add -
apt-get update
apt-get install grafana

or 
./grafana_repo.sh

or
apt-get install -y adduser libfontconfig
wget https://s3-us-west-2.amazonaws.com/grafana-releases/release/grafana_5.2.3_amd64.deb 
dpkg -i grafana_5.2.3_amd64.deb 
dpkg -i grafana_5.2.3_amd64.deb


vi /etc/grafana/grafana.ini

systemctl daemon-reload
systemctl enable grafana-server
systemctl start grafana-server

> 테스트
http://<ip>:3000/
admin / admin

# vi /etc/grafana/grafana.ini
[server]
http_port = 8080
# systemctl restart grafana-server

```

### webhook >> 작성 
> prometheus에서 지원하지 않는 line 알람을 위해 구성<br>
```
## webhook
http://guswnsxodlf.github.io/github-webhook-using-flask
http://yonggari.com/deploying-flask-app/

# apt install python3-pip 
# pip3 install virtualenv
# virtualenv -p python3 venv
# source /root/venv/bin/activate
(venv)# pip install flask uwsgi requests
(venv)# mkdir venv/project 
(venv)# cd venv/project
(venv)# vi app.py
from flask import Flask, request
app = Flask(__name__)

@app.route("/", methods=['post'])
def default():
        data = request.get_json()
        return 'success'


if __name__ == '__main__':
    app.run(host='0.0.0.0')
    
(venv)# flask run --host=0.0.0.0
>> flask 동작 여부 test

(venv)# vi wsgi.py
(venv)# uwsgi --socket 0.0.0.0:5000 --protocol=http -w wsgi:app
>> test : http://<ip>:5000/
(venv)# deactivate


# vi /root/venv/project/app.py
from flask import Flask, request
import requests

line_url = 'https://notify-api.line.me/api/notify'
line_token = '**'

app = Flask(__name__)

@app.route("/", methods=['post'])
def default():
  data = request.get_json()
  print(data)

  noti_line_message = {'message' : data['status'] + '\n' + 'description: '+ data['alerts'][0]['annotations']['description']}
  noti_line_header  = {'Authorization': 'Bearer ' + line_token}
  response = requests.post(url='https://notify-api.line.me/api/notify', headers=noti_line_header, data=noti_line_message)
  print(noti_line_message['message'])
  return 'success'


if __name__ == '__main__':
  app.run(host='0.0.0.0')


# vi ~/venv/project/webhook.ini
[uwsgi]
module = wsgi:app

master=true
processes=5

socket = webhook.sock
chmod-socket = 777
vacuum= true

die-on-term=true

# vi /etc/systemd/system/webhook.service
[Unit]
Description=flask webhook service
Wants=network-online.target
After=network-online.target

[Service]
User=root
Group=root
WorkingDirectory=/root/venv/project
Environment="PATH=/root/venv/bin"
ExecStart=/root/venv/bin/uwsgi --ini /root/venv/project/webhook.ini

[Install]
WantedBy=multi-user.target


# systemctl enable webhook.service
# systemctl start webhook.service


# apt install nginx
# vi /etc/nginx/sites-available/webhook
server {
    listen 8888;
    server_name server_domain_or_IP;

    location / {
        include uwsgi_params;
        uwsgi_pass unix:/root/venv/project/webhook.sock;
    }
}

# ln -s /etc/nginx/sites-available/webhook /etc/nginx/sites-enabled
# vi /etc/nginx/nginx.conf
> user root
# nginx -t
# rm /etc/nginx/sites-enabled/default
# rm /etc/nginx/sites-available/default


# systemctl status webhook prometheus alertmanager grafana-server nginx |grep -i active
```
