package data_type
import "strings"
import "fmt"
type stringType struct {
	 val string //字符串底层是一个[]byte 数组切片
}

func (a *stringType) Add(str string)  {
	 a.val=a.val+str;
}
func (a *stringType) Index(str string) int {
	  return strings.Index(a.val,str)
}
func (a * stringType) Split(str string) [] string {
	return strings.Split(a.val,str)
}
func (a * stringType) Len() int  {
	val:=a.val;
	return len(val)
}
func (a *stringType) RealLen() int  {
	real:=[]rune(a.val)
	for index,item:=range real{
		fmt.Println(string(item),index)
	}
	return len(real)
}
func (a * stringType) Reverse() string {
   var strSlice []byte=[]byte(a.val);
    len:=len(strSlice);
    lastIndex:=len-1;
	for i:=0;i<len/2;i++ {
		strSlice[i],strSlice[lastIndex-i]=strSlice[lastIndex-i],strSlice[i]
	}
    return  string(strSlice)
}
func (a * stringType) ReverseChiese(str string) string {
	  var strSlice [] rune=[] rune(str);
	  len:=len(strSlice);
	  lastIndex:=len-1;
	for i:=0;i<len/2;i++ {
		strSlice[i],strSlice[lastIndex-i]=strSlice[lastIndex-i],strSlice[i]
	}
	return  string(strSlice)
}
func (a *stringType) IsChinese(str string) [] string {
     var strSlice [] rune=[] rune(str)
     var strChinese [] string;
     var strIndex int
	for _,item:=range strSlice{
		 nStr:=string([] rune {item});
		 nStrSlice:=[] byte(nStr);
		if len(nStrSlice)>1 {
			strChinese=append(strChinese,string(nStrSlice))
			strIndex++
		}
	}
	return  strChinese
}
func NewStr(val string) *stringType{
	return &stringType{val:val}
}
