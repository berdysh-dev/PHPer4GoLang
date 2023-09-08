package PHPer4GoLang

import (
    "fmt"
    "runtime"
    "strings"
    "reflect"
)

type MyErrorStruct struct {
    line    int ;
    file    string ;
    mess    string ;
}

func (err *MyErrorStruct) Error() (string) {
    if(err.mess != ""){
        return fmt.Sprintf("%04d:%s:%v",err.line,err.file,err.mess) ;
    }else{
        return fmt.Sprintf("%04d:%s",err.line,err.file) ;
    }
}

func (err *MyErrorStruct) Line() (string) {
    return "HOHO" ;
}

func Err(v ... any) (error) {
    var err MyErrorStruct ;

    argc := len(v) ;
    n := 0 ;
    flag := 1 ;

    if(argc > 0){
        if(reflect.ValueOf(v[0]).Kind() == reflect.Int){
            flag = v[0].(int) ;
            argc -= 1 ;
            n += 1 ;
        }
    }

    if(flag == 0){ return nil ; }

    switch(argc){
    case 0: return nil ;
    case 1: err.mess = v[n].(string) ;
    case 2: err.mess = fmt.Sprintf(v[n].(string),v[n+1]) ;
    case 3: err.mess = fmt.Sprintf(v[n].(string),v[n+1],v[n+2]) ;
    case 4: err.mess = fmt.Sprintf(v[n].(string),v[n+1],v[n+2],v[n+3]) ;
    case 5: err.mess = fmt.Sprintf(v[n].(string),v[n+1],v[n+2],v[n+3],v[n+4]) ;
    case 6: err.mess = fmt.Sprintf(v[n].(string),v[n+1],v[n+2],v[n+3],v[n+4],v[n+5]) ;
    case 7: err.mess = fmt.Sprintf(v[n].(string),v[n+1],v[n+2],v[n+3],v[n+4],v[n+5],v[n+6]) ;
    }

    pc, file, line, ok := runtime.Caller(1) ;
    _ = pc ; _ = ok ;
    tmp := strings.Split(file,"/") ;
    err.file = tmp[len(tmp)-1] ;
    err.line = line ;

    return &err ;
}

