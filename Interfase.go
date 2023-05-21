package PHPer4GoLang

//    "io/fs"
//    "io/ioutil"
//    "log"
//    "os"
//    "fmt"

import (
    "reflect"
    "strconv"
)

func Strval(i interface{}) (string){
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

func Intval(i interface{}) (int){
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

func Boolval(i interface{}) (bool){

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














