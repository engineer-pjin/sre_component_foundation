# Golang 기초
## 소개
### 참조 사이트
[예제로 배우는 go 프로그래밍](http://golang.site/)<br>
[효과적인 Go 프로그래밍](https://gosudaweb.gitbooks.io/effective-go-in-korean/)<br>
[Tucker 의 GoLang 프로그래밍](https://www.youtube.com/watch?v=Tq3W8UyltFs&list=PLy-g2fnSzUTAaDcLW7hpq0e8Jlt7Zfgd6)<br>
[go tour](https://go-tour-kr.appspot.com/#1)<br><br>

### 특징
- class 가 없다. 단일 상속도 없고, 다중 상속도 없다.<br>
. 인터페이스(interfaces)로 다형성을 구현할 수 있다. 다른 OOP에서는 필드 없이, virtual 메서드로만 구성된 클래스 형태로 구현된다.<br>
. embedding으로 상속을 대신한다. 객체지향의 composition 모델과 비슷하다.<br>
https://golangkorea.github.io/post/go-start/object-oriented/

- 모듈화 의존성(빠른 빌드), 동적 타입 언어의 속성, 가비지컬렉션(실행 파일에서 지원), 병렬처리(추가 라이브러리가 아닌 언어 자체 지원)<br>
- 시스템 프로그래밍 언어 : 하드웨어나 운영체제에게 제공하는 서비스<br><br>

### Why go?
> [왜 golang인가?](https://andromedarabbit.net/%ec%99%9c-golang%ec%9d%b8%ea%b0%80/)
- Golang으로 작성한 세계 최대의 GitHub 저장소는 어디인가? >> [Kubernetes](https://kubernetes.io/)<br>
- 그외에도 docker, terraform, prometheus, chaos monkey, Fedora Core OS, influxDB, Istio<br>
- rest api를 통한 cloud native 구성에서 동시성 확보 및 심플<br>
- 배포 용이 : 운영체제와 프로그램 바이너리만(Kubernetes 처럼 클라우드 인프라를 떠받치는 시스템에선 매우 중요한 요소)<br>
- 동적 시스템에서 변경되는 리소스에 대해 별도의 정의 없이 활용 가능<br>

- Go 기반의 웹 프레임워크는 복잡한 비즈니스 로직을 단순화하는데 그다지 유용하지 않음<br>
. 데이터를 유연하게 다루는 능력은 Go의 장점이 아님

- goroutine : 쓰레드가 아닌 고루틴<br>
. 매번 커널 쓰레드를 생성하여 수행하지 않고 고루틴이 멀티플랙스를 이용해 쓰레드에 할당

- OS 쓰레드보다 훨씬 가볍게 비동기 Concurrent 처리를 구현하기 위하여 만든 것으로, 기본적으로 Go 런타임이 자체 관리 >> [Go 루틴 (goroutine)](http://golang.site/go/article/21-Go-%EB%A3%A8%ED%8B%B4-goroutine)

<br><br>

## 기초 문법
### Hello World
> With Visual Studio Code 

<br>

#### hello.go
```
package main // 패키지 명

import "fmt" // 사용할 패키지

func main() {
	fmt.Println("Hellow World")
}

```
<br>

### 변수 & 연산자
#### 변수 
+ 변수의 속성 : 이름, 타입, 값, 메모리 주소, 사이즈, 끝점(주소+사이즈)
+ int : 4/8 byte, int32 : 4byte(32bit), int64 : 8byte(64bit), int8 : 1byte(-128 ~ 127), uint8 : 1byte(0 ~ 255)<br>
. float32 : 4byte(32bit 숫자는 7개까지 표현), float64 : 8byte(64bit 숫자는 15개까지 표현)<br>
. string : rune(1~3byte)의 모임(배열)<br>
. bool : True / False<br>

+ 각 변수는 {} 기준으로 선언된 지역에서만 인식이 됨
+ 변수 선언 시 어떤 값으로 초기화 되는 지 명시적으로 스펙에 지정됨

<br>

#### 연산자
- 산술 연산자 : + - * / % ++ 11
- 관계 연산자 : == != <=
- 비트 연산자 : & | ^ << >>
- 논리 연산자 : && || !


```
package main

import "fmt"

func main() {
	var a int  // 변수 선언 방법 1 
	var b int = 4  // 변수 선언 방법 2
	a = 3  // 초기값 대입
	fmt.Println(a + b)

	c := 4  // 변수 선언 방법 3 && 초기화
	var d = 2  // 변수 선언 방법 4 && 초기값 대입
	fmt.Printf("c&d = %v\n", c&d)
	fmt.Printf("c|d = %v\n", c|d)
	fmt.Println("c^d = ", c^d)
	fmt.Println("c%d = ", c%d)

	fmt.Println("c << 1 = ", c<<1)    // *2와 동일, 쉬프트 연산이 곱셉이나 나눗셈보다 빠름
	fmt.Printf("c >> 1 = %v\n", c>>1) // /2와 동일

	var e bool = c > d // < > == !=
	var f bool = c < d
	fmt.Printf("c > d = %v\n", e)
	fmt.Printf("e && f = %v\n", e && f) // and
	fmt.Printf("e || f = %v\n", e || f) // or
}
```
> 참고 : 출력함수
> http://pyrasis.com/book/GoForTheReallyImpatient/Unit41<Br>
> func Print(a ...interface{}) (n int, err error): 값을 그 자리에 출력(새 줄로 넘어가지 않음)<Br>
> func Println(a ...interface{}) (n int, err error): 값을 출력한 뒤 새 줄로 넘어감(개행)<Br>
> func Printf(format string, a ...interface{}) (n int, err error): 형식을 지정하여 값을 출력<Br>

<br><br>

### 조건문
+ if
+ switch / case : 값을 입력하지 않으면 기본으로 True

#### if
```
package main 

import "fmt" 

func main() {
	g := 3
	h := 4
	if g > h {
		fmt.Println("g > h")
	} else if g < h {
		fmt.Println("g < h")
	} else if g == h {
		fmt.Println("g = h")
	} else {
		fmt.Println("??")
	}
}
```

<br>

#### switch / case
```
package main

import (
	"fmt"
)

func main() {
	n1 := 16
	n2 := 3
	line := "+" // 변경해서 확인

	switch line { 
	case "+":
		fmt.Printf("%d + %d = %d", n1, n2, n1+n2)
	case "-":
		fmt.Printf("%d - %d = %d", n1, n2, n1-n2)
	case "*":
		fmt.Printf("%d * %d = %d", n1, n2, n1*n2)
	case "/":
		fmt.Printf("%d / %d = %d", n1, n2, n1/n2)
	default:
		fmt.Println("잘못 입력하셨습니다.")

	}
}
```

<br><br>

### 반복문 : while은 없음
+ for
 - for의 조건 구간에 true가 들어가면 무한 루프<br>
  . break : 만나면 그냥 빠져나감<br>
  . continue : 조건에 맞으면 for문의 끝으로가서 다시 시작<br>


#### 기본 형태
```
package main

import "fmt"

func main() {

	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
		i += 2
	}
}
```

<br>

#### for 문에서 전처리, 후처리문 쓰기
```
package main

import "fmt"

func main() {
	var i int
	for i := 1; i < 10; i++ { 
		fmt.Println(i)
	}
	fmt.Println("최종 i의 값은", i)
}
```

<br>

#### continue, break
```
package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {
		if i == 5 {
			continue
		} else if i == 8 {
			break
		}
		fmt.Println(i)
	}
}
```

<br><br>

### **assignment01** : if 문과 for문을 활용하여 11~100 사이 선언한 두개의 숫자 중 처음 수에서 두번째 수를 뺀 값을 출력하고, 해당 결과가 11의 배수이거나 0이라면 종료 후 0을 리턴, 아니라면 첫번째 숫자의 1을 계속 더하여 11의 배수가 되는 횟수를 리턴해라.<br>
### 위 코드 작성 후 https://github.com/engineer-pjin/sre_component_foundation 레포 assignment 디렉토리에 추가 하시고 PR 요청해주시면 됩니다. (ex : assignment01_park.go)
> 참조 : https://wayhome25.github.io/git/2017/07/08/git-first-pull-request-story/

<br><br>

### 함수
#### 함수 : 반복되는 작업을 분리
 - 모듈화, 격리(분리 - decoupling)
 - 응집성(cohesive)는 높이고 종속성(dependency)은 낮춤 

#### 기본 형태
```
package main

import "fmt"

// add(x int, y int) == add(x, y int)
// 두개이상의 값을 반환 가능
func add(x int, y int) int { // 입력 int 2ea, 출력 int 1ea

	return x + y
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d + %d = %d\n", i, i+2, add(i, i+2))
	}

}
```

<br>

#### 재귀함수
 - 모든 재귀호출은 반목문으로 표현 가능
 - 반복문보다는 메모리 사용이 비효율적 : go는 함수형 언어가 아님, 느림
 - return을 통한 탈출 코드 필수
 - 재귀 호출이 필요한 경우
  .수학, 알고리즘 : 피보나치 수열 등

```
package main

import "fmt"

func recu(x int) {
	if x == 0 {
		return
	}
	fmt.Println(x)
	recu(x - 1)
}

func recusum(x int, s int) int {
	if x == 0 {
		return s
	}
	s += x
	return recusum(x-1, s)
}

func main() {
	recu(10)
	fmt.Println(recusum(10, 0))

}
```

<br>

####  pointer 
+ 포인터 연산이나 형변환 없이 명시적으로 사용
+ 변수는 "이름" "타입" "값" "메모리 주소" 를 가짐<br>
 . 포인터는 변수의 "메모리 주소" 
+ 함수로 포인터를 복사할때는 밸류 복사보다 양이 적다(메모리 주소만 복사되기 때문)

<br>

#### Pass By Value & Reference with pointer
```
package main

import "fmt"

func vadd(x int) {
	x++
}

func radd(y *int) {
	*y++
}

func main() {
	var a int
	var b *int

	b = &a
	a = 3
	fmt.Println(a, b, *b)

	*b = 5
	fmt.Println(a, b, *b)

	vadd(a)
	fmt.Println(a)

	radd(b)
	fmt.Println(a)

}


```

<br><br>

### Collection
자료구조를 표현한 형식

#### array
연속적인 메모리 공간에 동일한 타입의 데이타를 순서적으로 저장하는 자료구조
+ 배열 변수는 메모리상에서의 시작지점을 나타낸다.
+ utf8 의 글자는 1~3byte 임 (영어 1byte, 한글 3byte)
+ 문자열은 배열임, byte가 아닌 rune type은 한글등의 2byte 이상의 크기도 한글자로 인식해줌(utf8)

<br>

#### slice
동적 배열 : golang에서 주로 사용하는 타입
 - 배열의 크기가 변하면 새로운 메모리 공간을 만들고 복사(동적 배열은 실제 고정 배열을 가르키고 있음)
 - append로 추가시 할당된 공간을 넘게되면 복사가 되며, 그러면 원래 슬라이스와 다른 메모리 영역을 가짐 
  . 다른 메모리 영역을 가지려면 for문을 이용한 복사를 해라!

```
package main

import "fmt"

func RmBack(a []int, backnum int) ([]int, int) {
	return a[:len(a)-backnum], backnum
}

func main() {
	// array 선언
	var arr1 = [3]int{1, 2, 3}
    var arr2 = [...]int{1, 2, 3} //배열크기 자동으로
	fmt.Println(arr1, arr2)

    // slice 선언
	var a []int
	b := []int{1, 2, 3, 4}
	c := make([]int, 3, 8) // (type, len, cap)

	fmt.Println(len(a), cap(a))
	fmt.Println(len(b), cap(b), len(c), cap(c))
	fmt.Printf("%d %p %d %p\n", b, &b, c, &c)

	b = append(b, 1)
	c = append(c, 1)

	fmt.Println(len(b), cap(b), len(c), cap(c))
	fmt.Printf("%d %p %d %p\n", b, &b, c, &c) // 캐패시티의 변화에도 주소가 같음, 즉 메모리 복사가 발생하지 않음

	d := []int{1}
	e := append(d, 3)
	fmt.Printf("%p %p \n", d, e) // 변수가 바뀌고 캐패시티 변경 없음에도 주소가 달라짐

	st := make([]int, 7, 8)
	// 연산자 참조 http://pyrasis.com/book/GoForTheReallyImpatient/Unit13
	for i := 1; i <= len(st); i++ {
		if i == len(st) {
			break
		}
		st[i-1] = i
	}
	fmt.Println(st, st[3:6], st[:4]) // [1 2 3 4 5 6 0] [4 5 6] [1 2 3 4]

	st1 := st[:2]
	st1[0] = 10
	fmt.Println(st) // [10 2 3 4 5 6 0] st1은 st의 슬라이스된 부분의 주소값을 가르키기 때문에 st1을 변경하면 st도 변경됨
	var backnum int
	stb, backnum := RmBack(st, 3) // 두번째 인자만큼 슬라이스를 줄이는 함수
	fmt.Println(stb, backnum)
}

```

<br>

#### map
key & value 형식 = dict, hash table
```
package main

import "fmt"

func main() {
	//var mmap map[string]string     // 선언
	//mmap = make(map[string]string) // 초기화
	mmap := make(map[string]string) // 선언 초기화 동시에
	mmap["aa"] = "11"
	fmt.Println(mmap["aa"])
	fmt.Println(mmap["bb"]) //값이 없으면 기본값 출력(string 는 빈칸, int 0, bool은 false),

	mmap["cc"] = ""
	mmapNil1, ok1 := mmap["ss"] // 기본값이 없어서 기본값인지 원래 기본값인지 보기
	mmapNil2, ok2 := mmap["cc"] // 원래가 기본값이면 true 반환

	fmt.Println(mmapNil1, ok1, mmapNil2, ok2)

	delete(mmap, "cc")
	mmapNil2, ok2 = mmap["cc"] // 키에 해당하는 벨류 지우기
	fmt.Println(mmapNil1, ok1, mmapNil2, ok2)

	// 기본 map 선언 초기화 밸류넣기
	tickers := map[string]string{
		"GOOG": "Google Inc",
		"MSFT": "Microsoft",
		"FB":   "FaceBook",
		"AMZN": "Amazon",
	}
	fmt.Println(tickers["FB"])

	for key, value := range tickers { // 키와 밸류 모두 찾아내기, 무작위
		fmt.Println(key, value)
	}

}
```

<br><br>

### 구조체 : structure
새로운 타입 정의, 타언어의 class와 동일<br>
프로그래밍은 응집성(cohesive)을 높이고 종속성을 낮추는 방향으로 발전<br>
Custom Data Type을 표현하는데 사용되는데, 필드들의 집합체이며 필드들의 컨테이너<br> 
golang은 스트럭처에 속성 + 기능(method = Function)을 가짐<br>
관계에 따라 객체를 하나로 묶는, 관계는 기능을 정의할 수 있음<br>
 - 구조체 = 객체 (학생 구조체&객체 = 이름, 나이, 성별)<br>
 - 객체를 정의하고 객체간의 릴레이션쉽을  정의하는 것이 프로그래밍<br>
<br>

#### ex) 성적 처리 프로그램<br>
 - 객체 entity : 학생, 성적, 선생님
 - 관계 relationship : [선생님]이 [학생]의 [성적]을 "입력"한다, [학생]이 자신의 [성적]을 "조회"한다.
 - 기능 : 입력, 조회

<br>

#### Method
객체에 연결된 함수, structure가 필드만을 가지며 메서드는 별도로 분리되어 정의

```
package main

import (
	"fmt"
)

type Nums struct {
	Pnum int
	Snum int
}

func (n Nums) Add(Tnum int) int {	// Value Receiver
	n.Pnum += 1
	return n.Pnum + n.Snum + Tnum
}

func (n *Nums) Padd(Tnum int) int {	// Point Receiver
	n.Pnum += 1
	return n.Pnum + n.Snum + Tnum
}

func main() {
	var n Nums
	n.Pnum, n.Snum = 22, 11
	Tnum := 10

	fmt.Println(n.Pnum, n.Snum, Tnum, n.Pnum+n.Snum+Tnum)

	Addnum := n.Add(Tnum)
	fmt.Println(n.Pnum, n.Snum, Tnum, Addnum) // 함수의 변수 변경 내용이 메인에 영향을 주지 않음

	Paddnum := n.Padd(Tnum)
	fmt.Println(n.Pnum, n.Snum, Tnum, Paddnum) // 함수의 변수 변경 내용이 메인에 영향을 줌
}
```

<br><br>



### **assignment02** : 하단의 structure를 활용하여 학생의 성적을 입력, 수정, 삭제하는 메소드를 구현 및 실행<br>
```
type Student struct {
	name  string
	class int
	grade GradeResult
}

type GradeResult struct {
	name  string
	grade string
}
```
> 위 코드 작성 후 https://github.com/engineer-pjin/sre_component_foundation 레포 assignment 디렉토리에 추가 하시고 PR 요청해주시면 됩니다. (ex : assignment01_park.go)



<br><br>

### package
코드의 모듈화, 코드의 재사용 기능을 제공
> **GOROOT**는 보통은 mac, linux에서 /usr/local/go에 위치하는데 Go를 설치하면 Go관련된 실행파일, SDK 등이 위치<br>
> **GOPATH**는 커맨드라인에서 go get 명령어를 통해 받은 패키지나 라이브러리, 소스가 위치  

**gopath/main.go**
```
import (
	"fmt"
	"pjin"
)

func main() {
	fmt.Println("## golang go ##")
	pjin.OverlapEX()
}
```
<br>

**gopath/src/pjin/OverlapEX.go**
```
package pjin

import (
	"fmt"
)

func OverlapEX() {
	fmt.Println("## 중복 제거 ##")
	url := []string{"uu", "vv", "cc", "dd", "ff", "vv", "cc", "dd", "ee", "aa", "ww", "gg", "hh", "vv", "cc", "ss", "ee", "cc", "uu"}
	urlap := make(map[string]int)
	for i := 0; i < len(url); i++ {
		urlap[url[i]] += 1
	}
	fmt.Println(urlap)

	fmt.Println("")
	fmt.Println("## 정렬 ##")
	keytemp := make([]string, len(urlap))
	valuetemp := make([]int, len(urlap))
	tnum := 0

	for key, value := range urlap {
		keytemp[tnum] = key
		valuetemp[tnum] = value
		tnum += 1
	}
	fmt.Println(keytemp, valuetemp)

	for i := 1; i < len(urlap); i++ {
		for j := 0; j < i; j++ {
			if valuetemp[i] > valuetemp[j] { // 큰게 앞으로
				//if valuetemp[i] < valuetemp[j] {	// 큰게 뒤로
				valuetemp[i], valuetemp[j] = valuetemp[j], valuetemp[i]
				keytemp[i], keytemp[j] = keytemp[j], keytemp[i]
			}
		}
	}
	fmt.Println(keytemp, valuetemp)
}
```

<br><br>

### goroutine
+ kernel thread - wrapping -> GO(nm) thread
+ GO thread는 cpu의 코어에 최대한 가까우도록 kernel thread를 만들고 go thread를 그 안에 넣음<br>
 . 이로인해 context switching이 최소화 되도록 함<br>
 . Go 런타임이 자체 관리<br>
 . 매번 커널 쓰레드를 생성하여 수행하지 않고 고루틴이 멀티플랙스를 이용해 쓰레드에 할당<br>
 . OS 쓰레드보다 훨씬 가볍게 비동기 Concurrent 처리를 구현하기 위하여 만든 것으로, 기본적으로 Go 런타임이 자체 관리<br>
  
+ Go는 디폴트로 1개의 CPU를 사용한다. 즉, 여러 개의 Go 루틴을 만들더라도, 1개의 CPU에서 작업을 시분할하여 처리한다 (Concurrent 처리).<br>
 . 만약 머신이 복수개의 CPU를 가진 경우, Go 프로그램을 다중 CPU에서 병렬처리 (Parallel 처리)하게 할 수 있는데, 병렬처리를 위해서는 runtime.GOMAXPROCS(CPU수) 함수를 호출<br> 

+ deadlock : lock끼리 충돌, 쓰레드간 상대방이 lock을 풀어줘야 내가 lock을 잡는데 서로 상대방 락을 계속 대기상태<br>
  . 철학자들의 식사시간, 간헐적 발생으로 원인파악이 힘듬<br>

+ channel : 컨테이너 방식으로 쓰레드 간에 queue를 이용해 생산자-소비자 패턴 사용<br>
 > 참조 : https://brownbears.tistory.com/315

<br>

**goroutine**
```
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func say(s string, w *sync.WaitGroup) {
	defer w.Done()
	for i := 0; i < 10; i++ {
		fmt.Println(s, "***", i)
	}
}

func main() {
	wait := new(sync.WaitGroup)
	wait.Add(3)

	runtime.GOMAXPROCS(3)

	go say("Async1", wait)
	go say("Async2", wait)
	go say("Async3", wait)

	wait.Wait()
}
```

<br><br>

## web framework Gin
github : https://github.com/gin-gonic/gin<br>
특징 : 가벼움 - 마이크로 웹 프레임워크

### install
> go version 1.11+ is required
```
# go get -u github.com/gin-gonic/gin
```

<br>

#### example : ping/pong 
**main.go**
```
# vi main.go
package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() 
}

# go run main.go
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /ping                     --> main.main.func1 (3 handlers)
[GIN-debug] Environment variable PORT is undefined. Using port :8080 by default
[GIN-debug] Listening and serving HTTP on :8080

```
> visit 0.0.0.0:8080/ping (for windows "localhost:8080/ping") on browser

<br>


#### example os command
**main.go**
```
package main

import (
	"fmt"
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "health ok",
	})
}

func hostname(c *gin.Context) {
	cmd := exec.Command("hostname")
	stdoutStderr, _ := cmd.CombinedOutput()
	fmt.Println(string(stdoutStderr))
	c.JSON(200, gin.H{
		"message": string(stdoutStderr),
	})
}

func v2Any(c *gin.Context) {
	buf := make([]byte, 1024)
	v2Method := c.Request.Method
	n, _ := c.Request.Body.Read(buf)
	c.JSON(http.StatusOK, gin.H{"status": "good", "method": v2Method, "Body": string(buf[0:n])})
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	v1 := r.Group("/v1")
	{
		v1.GET("/health", health)
		v1.GET("/hostname", hostname)
	}

	v2 := r.Group("/v2")
	{
		v2.GET("/v2Get", v2Any)
		v2.POST("/v2Post", v2Any)
		v2.PUT("/v2Put", v2Any)
		v2.DELETE("/v2Delete", v2Any)
	}
	r.Run()
}

```

### **assignment03** : Rest API 정의에 따라 GET/POST/PUT/Delete 를 통해 학생의 성적을 생성/수정/삭제/표시 하는 api를 구현(assignment02 로직 활용) <br>
