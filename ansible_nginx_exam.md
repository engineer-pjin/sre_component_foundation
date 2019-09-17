## ansible example
```
# vi site.yml
---
- name: pkg install
  hosts: all
  gather_facts: false
  become: true
  vars:
  vars_files:
  tasks:
    - include: tasks/nginx.yml

  handlers:
    - include: handlers/all.yml


# vi tasks/nginx.yml
- name: nginx 설치
  apt:
    name: nginx
    state: present

- name: nginx source 복사
  template:
    src: index.html.j2
    dest: /var/www/html/index.html
  notify:
    - nginx restart


# vi handlers/all.yml
---
- name: nginx restart
  service:
    name: nginx
    state: restarted


# vi templates/index.html.j2
## {{ ansible_managed }}
<!DOCTYPE html>
<html>
<head>
<title>Welcome to ansible test {{ ansible_host }}</title>
<p><em>Thank you for using ansible</em></p>
</body>
</html>
```