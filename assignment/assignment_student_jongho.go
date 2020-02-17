package main

import "fmt"

type GradeResult struct {
	name  string
	grade string
}

type Student struct {
	name  string
	class int
	grade GradeResult
}

func (g *Student) insert(new_grade string) {
	g.grade.grade = new_grade
}

func (g *Student) update(new_grade string) {
	g.grade.grade = new_grade
}

func (g *Student) delete(delte_name string) {
	g.grade.grade = "none"
}

func main() {
	s := Student{}

	s.name = "Chae Jongho"
	s.class = 3
	s.grade.name = "Math"
	s.grade.grade = "A+"

	fmt.Println("초기 입력 값 출력 :")
	fmt.Println(s)

	// Grade 입력 메소드 활용
	s.insert("A")
	fmt.Println("Grade A 입력 :")
	fmt.Println(s)

	// Grade 수정 메소드 활용
	s.update("B-")
	fmt.Println("Grade B-로 수정 :")
	fmt.Println(s)

	// Grade 삭제 메소드 활용
	s.delete("Chae Jongho")
	fmt.Println("Grade 삭제 :")
	fmt.Println(s)
}
