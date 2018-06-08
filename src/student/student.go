package student
var arg [] int
type student struct {
	name string
	age int
}
type all struct {
	Students *student
}

func (a all) Add(num int)  {
	   c:=append(arg,num);
       arg=c;
}
func (a all) Get() [] int {
	return arg;
}
func NewAll(name string,age int) *all {
	return &all{Students:&student{name:name,age:age}}
}
