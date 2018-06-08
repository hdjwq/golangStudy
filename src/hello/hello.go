package main

import (
       "time_demo"
       "fmt"
)


func main()  {
     a:=time_demo.NewTime();
     str:=a.GetTime("yy:mm:dd hh:ff:ss")
     fmt.Printf("时间格式为%s\n",str)
     timeStamp:=a.GetTimeStamp();
     day,err:=a.Difference("2018-06-09 17:50:00")
       if err!=nil {
              fmt.Println(err)
       }
       fmt.Println(day)
     fmt.Println(timeStamp)
}


