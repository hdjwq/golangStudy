package time_demo

import (
	"time"
	"strings"
	"strconv"
)
type time_obj struct {
     now map[string] int
}

func (a time_obj) GetTime(str string) string {
	 var strTime string;
	 timeMap:=a.now;
	 strTime=strings.ToLower(str);
	 strTime=strings.Replace(strTime, "yy", strconv.Itoa(timeMap["yy"]), -1);
	 strTime=strings.Replace(strTime, "mm", strconv.Itoa(timeMap["mm"]), -1);
	 strTime=strings.Replace(strTime, "dd", strconv.Itoa(timeMap["dd"]), -1);
	 strTime=strings.Replace(strTime, "hh", strconv.Itoa(timeMap["hh"]), -1);
	 strTime=strings.Replace(strTime, "ff", strconv.Itoa(timeMap["ff"]), -1);
	 strTime=strings.Replace(strTime, "ss", strconv.Itoa(timeMap["ss"]), -1);
	 return strTime
}
func (a time_obj) GetTimeStamp() int64{
	 timeMap:=a.now
	 str:=time.Date(timeMap["yy"],time.Month(timeMap["mm"]),timeMap["dd"],timeMap["hh"],timeMap["ff"],timeMap["ss"],0,time.Local);
	 return  str.Unix();
}
func (a time_obj) Difference(str string) (float32,error) {
	 date,err:=time.Parse("2006-01-02 15:04:05",str);
	 if err!=nil {
		return 0,err
	 }
	 nowTime:=a.GetTimeStamp();
	 cj:=float32(date.Unix()-nowTime);
	 day:=cj/60/60/24
     return day,nil
}
func NewTime()  *time_obj{
	 now:=time.Now();
	 timeMap:=make(map[string] int);
	 timeMap["yy"]=now.Year();
	 timeMap["mm"]=int(now.Month());
	 timeMap["dd"]=now.Day();
	 timeMap["hh"]=now.Hour();
	 timeMap["ff"]=now.Minute();
	 timeMap["ss"]=now.Second();
	 return  &time_obj{now:timeMap}
}
