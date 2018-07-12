package builder

type Student struct {
	Name string
	Age int
	Sex string
	No int
}

type StudentBuilder struct {
	student *Student
}

func NewStudentBuilder() *StudentBuilder {

	return &StudentBuilder{student:&Student{}}
}

func (s * StudentBuilder)Name(name string) *StudentBuilder  {
	s.student.Name = name
	return s
}
func (s * StudentBuilder)Sex(sex string) *StudentBuilder  {
	s.student.Sex = sex
	return s
}
func (s * StudentBuilder)No(no int) *StudentBuilder  {
	s.student.No = no
	return s
}
func (s * StudentBuilder)Build() *Student  {
	return s.student
}


