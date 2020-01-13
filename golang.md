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
> assignment : if 문과 for문을 활용하여 11~100 사이 선언한 두개의 숫자의 값을 출력하고, 해당 결과가 11의 배수이거나 0이라면 종료 후 0을 리턴, 아니라면 첫번째 숫자의 1을 계속 더하여 11의 배수가 되는 횟수를 리턴해라. 