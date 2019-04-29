# git foundation
> **참조 사이트**<br>
> [git 공식 메뉴얼](https://git-scm.com/book/ko/v2) <br>
> [참조 사이트 : 생활 코딩 - 지옥에서 온 Git](https://opentutorials.org/course/2708)

## 상세
+ 리누즈 토발즈가 처음에 만들기 시작<br>
+ Git은 데이터를 파일 시스템 스냅샷의 연속으로 취급하고 크기가 아주 작다. Git은 커밋하거나 프로젝트의 상태를 저장할 때마다 파일이 존재하는 그 순간을 중요하게 여긴다. 파일이 달라지지 않았으면 Git은 성능을 위해서 파일을 새로 저장하지 않는다. 단지 이전 상태의 파일에 대한 링크만 저장한다. Git은 데이터를 스냅샷의 스트림처럼 취급한다. <br>
+ add (index에 기록) > commit : (오브젝트로 기록)<br>
&nbsp;&nbsp;- working directory & index, staging area, cache & repository<br>
![git 프로젝트의 세가지 단계](https://git-scm.com/book/en/v2/images/areas.png)


+ 파일의 내용이 같다면 git 내부에서 object(sha1 해쉬알고리즘)는 동일하게 취급됨, 중복 제거<br>
+ commit 자체도 오브젝트로 취급되어 기록된다.(커밋당 하나의 오브젝트) <br>
+ 오브젝트는 blob(file) tree(directory) commit 으로 구분됨<br>
파일의 내용이 변하고 커밋될때마다 오브젝트도 변함,<br>

<br><br>
## 기본 CLI
> ubuntu 16.04 기준 <br>

**설치**

```
# apt install git
```

<br>

**초기화**

```
# git init
```
> init 되지 않는 폴더라면 아래와 같은 에러 발생<br>
**fatal: Not a git repository (or any of the parent directories): .git**

<br>

**사용자 정보 입력 및 확인** 

```
# git config --global user.name pjin
# git config --global user.email pjin@wemakeprice.com

# git config --list
```

<br>

**상태 확인** 

```
# git status
```

<br>

**index에 기록**

```
# git add [file name]
```

<br>

**object로 기록**

```
# git commit
```
> vi editor<br>
> commit message를 넣음<br>

<br>

**add 와 commit 동시**

```
# git commit -am 'commit message'
```
> a:add 없이, m:메세지 직접<br>
> tracking 하고 있지 않는 파일은 해당 되지 않음

<br>

**commit 로그 확인**

```
# git commit -am 'commit message'
```

<br>

**commit 제외 리스트 작성**

```
# vi .gitignore
imsy_log.txt
tmp/
```
> 버전 관리에서 제외할 파일 목록을 지정

<br>

**버전 간의 차이점을 비교할 때**

```
# git diff '버전 id'..'버전 id2'
```

<br>

**git add하기 전과 add한 후의 파일 내용을 비교할 때**

```
# git diff
```

<br>

**reset**

```
# git reset '버전 id'
```
> reset을 실행하면 돌아 가려는 커밋으로 리파지토리는 재설정되고, 해당 커밋 이후의 이력은 사라짐<br>
> 지정되는 version이 내가 원하는 마지막 버전<br>
> .git/logs/HEAD 에 가면 reset 했다는 로그도 보임<br>
>
> hard, soft, mixed(default)<br>
>> 
working directroy | index | repository
---|:---:|---:
working tree| staging area | history
working copy| cache | tree 


 >> soft : repository <br>
 >> mixed : repository + index<br>
 >> hard : repository + index + working directroy<br>
 >> git reset --hard 까지만 하면 가장 최신 커밋으로 돌아감.. 즉 수정하다가 수정한거 초기화 하고 싶으면 사용

<br>

**revert**

```
# git revert '버전 id'
```
> revert를 실행하면 되돌리겠다는 기록을 남기고 기존 버전으로 돌리기<br>
 >> revert를 사용하여 해당 커밋의 내용만 되돌릴 수 있음, 즉 revert로 지정된 버전의 commit은 무효화<br>
 >> 원격 리파지토리에 push 를 한 상태에서 reset을 사용하면 reset 하기 이전으로 되돌리기 전까지는 push 할 수 없음 
     그래서 이미 push 한 코드라면 미련을 버리고 revert를 실행하여야 함
     
<br>

**원격 저장소 github**
> 분산 버전 관리 툴인 깃(Git)을 사용하는 프로젝트를 지원하는 웹호스팅 서비스<br>
> gui 제공, MS 인수<br>
> 무료 계정도 비공개 저장소를 사용 가능하나 협업이 3명까지만 가능<br>

등록 및 push
```
# git remote add origin  https://github.com/engineer-pjin/sre_component_foundation
# git remote -v
origin  https://github.com/engineer-pjin/test.git (fetch)
origin  https://github.com/engineer-pjin/test.git (push)

# git push -u origin master 

# git push
```

<br>

git의 소스코드를 지역저장소로 가져오기
```
# git clone https://github.com/git/git.git [directory]
```

<br>

pull
```
# git pull
```

<br>

신규 서버에서 git pull을 하려면
```
# git init
# git remote add origin  https://github.com/engineer-pjin/sre_component_foundation
# git remote -v
# git pull origin master
Username for 'https://github.com': engineer-pjin
Password for 'https://engineer-pjin@github.com':
From https://github.com/engineer-pjin/sre_component_foundation
 * branch            master     -> FETCH_HEAD
```


<br><br>

## 실습 1
> 기본 CLI 실습

```
# git init
Initialized empty Git repository in /root/git/.git/

# ls -al
total 12
drwxr-xr-x  3 root root 4096 Apr 29 15:30 .
drwx------ 12 root root 4096 Apr 29 15:30 ..
drwxr-xr-x  7 root root 4096 Apr 29 15:30 .git

# git config --global user.name engineer.pjin

# git config --global user.email engineer.pjin@gmail.com

# git config --list
user.name=engineer.pjin
user.email=engineer.pjin@gmail.com
core.repositoryformatversion=0
core.filemode=true
core.bare=false
core.logallrefupdates=true

# vi test01.txt
test01-1
test01-2
test01-3
test01-4

# vi test02.txt
test02-1
test02-2
test02-3
test02-4

# git status
On branch master

Initial commit

Untracked files:
  (use "git add <file>..." to include in what will be committed)

        test01.txt
        test02.txt

nothing added to commit but untracked files present (use "git add" to track)

# git add test01.txt test02.txt

# git status
On branch master

Initial commit

Changes to be committed:
  (use "git rm --cached <file>..." to unstage)

        new file:   test01.txt
        new file:   test02.txt

# git commit -m 'first commit'
[master (root-commit) 9dd6472] first commit
 2 files changed, 8 insertions(+)
 create mode 100644 test01.txt
 create mode 100644 test02.txt

# git log
commit 9dd647264c29d9865f2596c94c6dc09480dee758
Author: engineer.pjin <engineer.pjin@gmail.com>
Date:   Mon Apr 29 15:39:12 2019 +0900

    first commit

# cat .git/refs/heads/master
9dd647264c29d9865f2596c94c6dc09480dee758

# vi test03.txt
test03-1
test03-2
test03-3
test03-4

# git add test03.txt
root@test01a.dev.jin.sysadmin:~/git

# git status
On branch master
Changes to be committed:
  (use "git reset HEAD <file>..." to unstage)

        new file:   test03.txt

# rm test03.txt

# git status
On branch master
Changes to be committed:
  (use "git reset HEAD <file>..." to unstage)

        new file:   test03.txt

Changes not staged for commit:
  (use "git add/rm <file>..." to update what will be committed)
  (use "git checkout -- <file>..." to discard changes in working directory)

        deleted:    test03.txt

# cat test03.txt
test03-1
test03-2
test03-3
test03-4

# git status
On branch master
Changes to be committed:
  (use "git reset HEAD <file>..." to unstage)

        new file:   test03.txt

# git commit -am 'test03.txt add'
[master 05ae8f4] test03.txt add
 1 file changed, 4 insertions(+)
 create mode 100644 test03.txt

# git status
On branch master
nothing to commit, working directory clean

# git log
commit 05ae8f45c5fbffd45cf7a2b54a6df095de294bd9
Author: engineer.pjin <engineer.pjin@gmail.com>
Date:   Mon Apr 29 15:44:30 2019 +0900

    test03.txt add

commit 9dd647264c29d9865f2596c94c6dc09480dee758
Author: engineer.pjin <engineer.pjin@gmail.com>
Date:   Mon Apr 29 15:39:12 2019 +0900

    first commit

# cat .git/refs/heads/master
05ae8f45c5fbffd45cf7a2b54a6df095de294bd9

# git diff 9dd647264c29d9865f2596c94c6dc09480dee758..05ae8f45c5fbffd45cf7a2b54a6df095de294bd9
diff --git a/test03.txt b/test03.txt
new file mode 100644
index 0000000..39191bf
--- /dev/null
+++ b/test03.txt
@@ -0,0 +1,4 @@
+test03-1
+test03-2
+test03-3
+test03-4

# rm test03.txt

# git diff
diff --git a/test03.txt b/test03.txt
deleted file mode 100644
index 39191bf..0000000
--- a/test03.txt
+++ /dev/null
@@ -1,4 +0,0 @@
-test03-1
-test03-2
-test03-3
-test03-4

# git status
On branch master
Changes not staged for commit:
  (use "git add/rm <file>..." to update what will be committed)
  (use "git checkout -- <file>..." to discard changes in working directory)

        deleted:    test03.txt

no changes added to commit (use "git add" and/or "git commit -a")

# vi test02.txt
test02-1
test02-2
test02-3
test02-5

# git status
On branch master
Changes not staged for commit:
  (use "git add/rm <file>..." to update what will be committed)
  (use "git checkout -- <file>..." to discard changes in working directory)

        modified:   test02.txt
        deleted:    test03.txt

no changes added to commit (use "git add" and/or "git commit -a")

# git diff
diff --git a/test02.txt b/test02.txt
index 4bd88a4..6dcfd9a 100644
--- a/test02.txt
+++ b/test02.txt
@@ -1,4 +1,4 @@
 test02-1
 test02-2
 test02-3
-test02-4
+test02-5
diff --git a/test03.txt b/test03.txt
deleted file mode 100644
index 39191bf..0000000
--- a/test03.txt
+++ /dev/null
@@ -1,4 +0,0 @@
-test03-1
-test03-2
-test03-3
-test03-4

# git commit -am "test03.txt del, test02.txt edit"
[master 9d43228] test03.txt del, test02.txt edit
 2 files changed, 1 insertion(+), 5 deletions(-)
 delete mode 100644 test03.txt

# git log
commit 9d432281cf5ac5693ef0cf5bebdfc62fd819ffb4
Author: engineer.pjin <engineer.pjin@gmail.com>
Date:   Mon Apr 29 15:48:36 2019 +0900

    test03.txt del, test02.txt edit

commit 05ae8f45c5fbffd45cf7a2b54a6df095de294bd9
Author: engineer.pjin <engineer.pjin@gmail.com>
Date:   Mon Apr 29 15:44:30 2019 +0900

    test03.txt add

commit 9dd647264c29d9865f2596c94c6dc09480dee758
Author: engineer.pjin <engineer.pjin@gmail.com>
Date:   Mon Apr 29 15:39:12 2019 +0900

    first commit

# touch test.log

# vi .gitignore
test.log

# git status
On branch master
Untracked files:
  (use "git add <file>..." to include in what will be committed)

        .gitignore

nothing added to commit but untracked files present (use "git add" to track)

# git add .gitignore

# git commit -am "gitignore add"
[master 046905c] gitignore add
 1 file changed, 1 insertion(+)
 create mode 100644 .gitignore

# git log
commit 046905c1be4adcaf0e243015ae85648d8e451b05
Author: engineer.pjin <engineer.pjin@gmail.com>
Date:   Mon Apr 29 15:52:39 2019 +0900

    gitignore add

commit 9d432281cf5ac5693ef0cf5bebdfc62fd819ffb4
Author: engineer.pjin <engineer.pjin@gmail.com>
Date:   Mon Apr 29 15:48:36 2019 +0900

    test03.txt del, test02.txt edit

commit 05ae8f45c5fbffd45cf7a2b54a6df095de294bd9
Author: engineer.pjin <engineer.pjin@gmail.com>
Date:   Mon Apr 29 15:44:30 2019 +0900

    test03.txt add

commit 9dd647264c29d9865f2596c94c6dc09480dee758
Author: engineer.pjin <engineer.pjin@gmail.com>
Date:   Mon Apr 29 15:39:12 2019 +0900

    first commit

# git reset --hard 9d432281cf5ac5693ef0cf5bebdfc62fd819ffb4
HEAD is now at 9d43228 test03.txt del, test02.txt edit

# ls -al
total 20
drwxr-xr-x  3 root root 4096 Apr 29 16:07 .
drwx------ 12 root root 4096 Apr 29 15:51 ..
drwxr-xr-x  8 root root 4096 Apr 29 16:07 .git
-rw-r--r--  1 root root    0 Apr 29 15:51 test.log
-rw-r--r--  1 root root   36 Apr 29 15:33 test01.txt
-rw-r--r--  1 root root   36 Apr 29 15:47 test02.txt

# cat test02.txt
test02-1
test02-2
test02-3
test02-5

# cat .git/logs/HEAD
0000000000000000000000000000000000000000 9dd647264c29d9865f2596c94c6dc09480dee758 engineer.pjin <engineer.pjin@gmail.com> 1556519952 +0900      commit (initial): first commit
9dd647264c29d9865f2596c94c6dc09480dee758 05ae8f45c5fbffd45cf7a2b54a6df095de294bd9 engineer.pjin <engineer.pjin@gmail.com> 1556520270 +0900      commit: test03.txt add
05ae8f45c5fbffd45cf7a2b54a6df095de294bd9 9d432281cf5ac5693ef0cf5bebdfc62fd819ffb4 engineer.pjin <engineer.pjin@gmail.com> 1556520516 +0900      commit: test03.txt del, test02.txt edit
9d432281cf5ac5693ef0cf5bebdfc62fd819ffb4 046905c1be4adcaf0e243015ae85648d8e451b05 engineer.pjin <engineer.pjin@gmail.com> 1556520759 +0900      commit: gitignore add
046905c1be4adcaf0e243015ae85648d8e451b05 9d432281cf5ac5693ef0cf5bebdfc62fd819ffb4 engineer.pjin <engineer.pjin@gmail.com> 1556521658 +0900      reset: moving to 9d432281cf5ac5693ef0cf5bebdfc62fd819ffb4

# git log
commit 9d432281cf5ac5693ef0cf5bebdfc62fd819ffb4
Author: engineer.pjin <engineer.pjin@gmail.com>
Date:   Mon Apr 29 15:48:36 2019 +0900

    test03.txt del, test02.txt edit

commit 05ae8f45c5fbffd45cf7a2b54a6df095de294bd9
Author: engineer.pjin <engineer.pjin@gmail.com>
Date:   Mon Apr 29 15:44:30 2019 +0900

    test03.txt add

commit 9dd647264c29d9865f2596c94c6dc09480dee758
Author: engineer.pjin <engineer.pjin@gmail.com>
Date:   Mon Apr 29 15:39:12 2019 +0900

    first commit
    
# ls
test.log  test01.txt  test02.txt

# git status
On branch master
Untracked files:
  (use "git add <file>..." to include in what will be committed)

        test.log

nothing added to commit but untracked files present (use "git add" to track)

# git add test.log

# git commit -am 'test.log add'
[master ae0f59a] test.log add
 1 file changed, 0 insertions(+), 0 deletions(-)
 create mode 100644 test.log

# git log
commit ae0f59a3f524380a248697af38e9766500084136
Author: engineer.pjin <engineer.pjin@gmail.com>
Date:   Mon Apr 29 16:18:42 2019 +0900

    test.log add

commit 9d432281cf5ac5693ef0cf5bebdfc62fd819ffb4
Author: engineer.pjin <engineer.pjin@gmail.com>
Date:   Mon Apr 29 15:48:36 2019 +0900

    test03.txt del, test02.txt edit

commit 05ae8f45c5fbffd45cf7a2b54a6df095de294bd9
Author: engineer.pjin <engineer.pjin@gmail.com>
Date:   Mon Apr 29 15:44:30 2019 +0900

    test03.txt add

commit 9dd647264c29d9865f2596c94c6dc09480dee758
Author: engineer.pjin <engineer.pjin@gmail.com>
Date:   Mon Apr 29 15:39:12 2019 +0900

    first commit

# git revert ae0f59a3f524380a248697af38e9766500084136
[master f51e653] Revert "test.log add"
 1 file changed, 0 insertions(+), 0 deletions(-)
 delete mode 100644 test.log

# git log
commit f51e6536d8e893674b01695bce0fe345ea4df6f0
Author: engineer.pjin <engineer.pjin@gmail.com>
Date:   Mon Apr 29 16:20:03 2019 +0900

    Revert "test.log add"

    This reverts commit ae0f59a3f524380a248697af38e9766500084136.

commit ae0f59a3f524380a248697af38e9766500084136
Author: engineer.pjin <engineer.pjin@gmail.com>
Date:   Mon Apr 29 16:18:42 2019 +0900

    test.log add

commit 9d432281cf5ac5693ef0cf5bebdfc62fd819ffb4
Author: engineer.pjin <engineer.pjin@gmail.com>
Date:   Mon Apr 29 15:48:36 2019 +0900

    test03.txt del, test02.txt edit

commit 05ae8f45c5fbffd45cf7a2b54a6df095de294bd9
Author: engineer.pjin <engineer.pjin@gmail.com>
Date:   Mon Apr 29 15:44:30 2019 +0900

    test03.txt add

commit 9dd647264c29d9865f2596c94c6dc09480dee758
Author: engineer.pjin <engineer.pjin@gmail.com>
Date:   Mon Apr 29 15:39:12 2019 +0900

    first commit

# ls
test01.txt  test02.txt


```

<br><br>

## 더 알아보기 - branch
> 개념을 익힌 후 실습은 필요 시 진행

+ **branch** : 모든 버전 관리 시스템은 브랜치를 지원한다. 개발을 하다 보면 코드를 여러 개로 복사해야 하는 일이 자주 생긴다. 코드를 통째로 복사하고 나서 원래 코드와는 상관없이 독립적으로 개발을 진행할 수 있는데, 이렇게 독립적으로 개발하는 것이 브랜치다.
+ 각 브랜치를 하나의 “실험실” 로 생각
![](https://git-scm.com/book/en/v2/images/lr-branches-2.png)
![](https://git-scm.com/book/en/v2/images/topic-branches-1.png)

### CLI
**생성 및 전환**

```
# git branch
# git branch "새로운 브랜치 이름"
# git branch -d
# git branch -D
> 병합하지 않은 브랜치를 강제 삭제할 때 
# git checkout "전환하려는 브랜치 이름"
> 브랜치를 전환(체크아웃)할 때, 그냥 log에 보이는 id를 넣는다면 그 버젼으로 돌아가기도 함
# git checkout -b "생성하고 전환할 브랜치 이름"
> 브랜치를 생성하고 전환까지 할 때 
> 브랜치 변경에 따라 파일의 내용 및 신규/삭제 파일은 각 브랜치를 따라간다

# git checkout [commit id]
> 해당 커밋 상태로 돌아감

# git log master..jin -p
> 마스터에는 없고 jin에는 있는거
# git diff master..jin

> A 브랜치로 B 브랜치를 병합할 때 (master ← jin)
# git checkout master
# git merge jin
```


**branch CLI 그래프**

```
# git log --branches --decorate --graph
* commit 6f26986f2c7235da8feaec06725461b7594f6234 (HEAD -> master)
| Author: pjin <pjin.test>
| Date:   Sun Jan 27 23:33:54 2019 +0900
|
|     v4
|
| * commit 17446713468c934dd28f3b3baae10ded71436ffa (jin)
| | Author: pjin <pjin.test>
| | Date:   Sun Jan 27 23:32:52 2019 +0900
| |
| |     v5
| |
| * commit d7b967ae9a54fa52299664edcaab9faba26823f1
| | Author: pjin <pjin.test>
| | Date:   Sun Jan 27 23:26:04 2019 +0900
| |
| |     v4
| |
| * commit 63701734c989279254121be9e9cc22bf55aee6b1
|/  Author: pjin <pjin.test>
|   Date:   Sun Jan 27 23:22:30 2019 +0900
|
|       v3
|
* commit 7f9444eb2dfd50b561b20a35fb52b706c5d87733
  Author: pjin <pjin.test>
  Date:   Sun Jan 27 23:10:26 2019 +0900
```

**merge conflict**

> git merge 시 서로 다른 부분을 추가했으면 추가된 부분이 모두 합쳐짐<br>
>> base, local, remote : 세 오브젝트에서 머지 내용 조절<br>
> 같은 부분을 수정했나면 conflict 이 발생하고 해당 파일에 표시가 됨.. 그 표시되는 애를 수정하면 됨<br>


>충돌이 생기면 아래와 같은 메시지 발생
>>fix conflict and run "gir commit"
