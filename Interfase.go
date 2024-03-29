package PHPer4GoLang

//    "io/fs"
//    "io/ioutil"
//    "log"
//    "os"

import (
    "reflect"
    "strconv"
    "strings"
    "fmt"
)

func Strval(i any) (string){
    ret := "" ;

    v := reflect.ValueOf(i) ;

    switch(v.Kind()){
    case reflect.String:{
            ret = i.(string) ;
        }
    case reflect.Int:{
            ret = strconv.Itoa(i.(int)) ;
        }
    case reflect.Bool:{
            if(i.(bool)){
                ret = "1" ;
            }else{
                ret = "0" ;
            }
        }
    default:{
            Debugf("ERR:Kind[%v]",v.Kind()) ;
        }
    }

    return ret ;
}

func Intval(i any) (int){
    ret := 0 ;

    v := reflect.ValueOf(i) ;

    switch(v.Kind()){
    case reflect.String:{
            // Debugf("strinと判定[%s]",i.(string)) ;
            num , err := strconv.Atoi(i.(string)) ;
            if(err == nil){
                // Debugf("Atoi 成功[%d]",num) ;
                ret = num ;
            }else{
                // Debugf("Atoi 失敗[%v]",err) ;
            }
        }
    case reflect.Int:{
            ret = i.(int) ;
        }
    case reflect.Bool:{
            if(i.(bool)){
                ret = 1 ;
            }else{
                ret = 0 ;
            }
        }
    default:{
            // Debugf("ERR:Kind[%v]",v.Kind()) ;
        }
    }

    return ret ;
}

func Boolval(i any) (bool){

    ret := false ;

    v := reflect.ValueOf(i)

    switch(v.Kind()){
    case reflect.Bool:{
            ret = i.(bool) ;
        }
    case reflect.Int:{
            if(i.(int) != 0){
                ret = true ;
            }else{
                ret = false ;
            }
        }
    case reflect.String:{
            // Debugf("strinと判定[%s]",i.(string)) ;
            num , err := strconv.Atoi(i.(string)) ;
            if(err == nil){
                // Debugf("Atoi 成功[%d]",num) ;
                if(num != 0){
                    ret = true ;
                }else{
                    ret = false ;
                }
            }else{
                Debugf("Atoi 失敗[%v]",err) ;
            }
        }
    case reflect.Invalid:{
            ret = false ;
        }
    default:{
            Debugf("ERR:Kind[%v]",v.Kind()) ;
            ret = false ;
        }
    }

    return ret ;
}

func Gettype(i any) (string){
    return fmt.Sprintf(`%T`,i) ;
}

func Gettype2(i any) (string,string){
    v := reflect.ValueOf(i)

    k := fmt.Sprintf(`%v`,v.Kind()) ;
    t := fmt.Sprintf(`%T`,i) ;

    t = strings.Replace(t, "interface{}", "any", -1) ;
    t = strings.Replace(t, "interface {}", "any", -1) ;

    // Debugf("![%v]",t) ;

    return k,t ;
}












