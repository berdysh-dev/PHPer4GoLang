package PHPer4GoLang

//    "io/fs"
//    "io/ioutil"
//    "log"
//    "os"
//    "fmt"

import (
    "time"
    "log"
)

const location = "Asia/Tokyo" ;

type TypeLocalTime struct {
    Tm_sec      int ;
    Tm_min      int ;
    Tm_hour     int ;
    Tm_mday     int ;
    Tm_mon      int ;
    Tm_year     int ;
    Tm_wday     int ;
    Tm_yday     int ;
    Tm_isdst    int ;
}

func Mon2mon(mon time.Month) (int){
    ret := 0 ;
    switch(mon){
    case time.January   : ret = 1 ;
    case time.February  : ret = 2 ;
    case time.March     : ret = 3 ;
    case time.April     : ret = 4 ;
    case time.May       : ret = 5 ;
    case time.June      : ret = 6 ;
    case time.July      : ret = 7 ;
    case time.August    : ret = 8 ;
    case time.September : ret = 9 ;
    case time.October   : ret = 10 ;
    case time.November  : ret = 11 ;
    case time.December  : ret = 12 ;
    }
    return ret ;
}

func Date_default_timezone_set(location string){
    loc , err := time.LoadLocation(location) ;

    if(err != nil){
        log.Print(err) ;
    }else{
        time.Local = loc ;
    }
}

func Time() (int64) {
    unixtime := time.Now().Unix() ;
    return unixtime ;
}

func Gmtime(unixtime int64) (TypeLocalTime) {
    var tm TypeLocalTime ;

    loc , _ := time.LoadLocation("Etc/GMT") ;

    local := time.Unix(unixtime,0) ;
    gmt := local.In(loc) ;

    tm.Tm_year  = gmt.Year() - 1900 ;
    tm.Tm_mon   = Mon2mon(gmt.Month()) - 1 ;
    tm.Tm_mday  = gmt.Day() ;
    tm.Tm_hour  = gmt.Hour() ;
    tm.Tm_min   = gmt.Minute() ;
    tm.Tm_sec   = gmt.Second() ;

    return tm ;
}

func Localtime(unixtime int64) (TypeLocalTime) {
    var tm TypeLocalTime ;

    local := time.Unix(unixtime,0) ;

    tm.Tm_year  = local.Year() - 1900 ;
    tm.Tm_mon   = Mon2mon(local.Month()) - 1 ;
    tm.Tm_mday  = local.Day() ;
    tm.Tm_hour  = local.Hour() ;
    tm.Tm_min   = local.Minute() ;
    tm.Tm_sec   = local.Second() ;

    return tm ;
}

func Mktime(hour int,minute int,second int,month int,day int,year int) (int64){
    return 0 ;
}














