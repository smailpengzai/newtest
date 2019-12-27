package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"strings"
	"time"
)

func main() {
	//beego.Notice(fmt.Sprintf("上周的周一是：%v",GetLastWeekFirstDate()))
	//beego.Notice(getFirstDateOfWeek())
	getWeekFastDay(49)
}

/**
获取本周周一 ~~~~ 星期天的日期
*/
func getFirstDateOfWeek() string {
	now := time.Now()

	offset := int(time.Monday - now.Weekday())

	//offsetEnd := int( now.Weekday()-time.Sunday )
	if offset > 0 {
		offset = -6
	}

	weekStartDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	weekEndDate := weekStartDate.AddDate(0, 0, 6)
	//time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offsetEnd)
	weekDay := weekStartDate.Format("2006-01-02") + " ~ " + strings.Split(weekEndDate.Format("2006-01-02"), "-")[2]
	_, weekWeek := now.ISOWeek()
	return fmt.Sprintf("第%v周(%s)", weekWeek, weekDay)
}

/**
获取本周周一 的日期
*/
func getFirstDayOfWeek() string {

}

/**
获取上周的周一日期
*/
func GetLastWeekFirstDate() (lastWeekMonday time.Time) {
	thisWeekMonday := getFirstDayOfWeek()
	TimeMonday, _ := time.Parse("2006-01-02", thisWeekMonday)
	lastWeekMonday = TimeMonday.AddDate(0, 0, -7)
	//weekMonday = lastWeekMonday.Format("2006-01-02")
	return
}

func getWeekFastDay(isweek int) {
	now := time.Now()
	_, nowWeek := now.ISOWeek()
	diffWeek := nowWeek - isweek
	beego.Notice(diffWeek)
	thisWeekMonday := getFirstDayOfWeek()
	TimeMonday, _ := time.Parse("2006-01-02", thisWeekMonday)
	startday := TimeMonday.AddDate(0, 0, -7*diffWeek)
	beego.Notice(startday.Format("2006-01-02 00:00:00"))
}
