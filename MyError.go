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

func E(v ... any) (error) {
    var err MyErrorStruct ;

    argc := len(v) ;
    off := 0 ;
    flag := 1 ;

    if(argc > 0){
        if(reflect.ValueOf(v[0]).Kind() == reflect.Int){
            flag = v[0].(int) ;
            argc -= 1 ;
            off += 1 ;
        }
    }

    if(flag == 0){ return nil ; }

    switch(argc){
    case 0: return nil ;
    case 1: err.mess = v[off].(string) ;
    case 2: err.mess = fmt.Sprintf(v[off].(string),v[off+1]) ;
    case 3: err.mess = fmt.Sprintf(v[off].(string),v[off+1],v[off+2]) ;
    }

    pc, file, line, ok := runtime.Caller(1) ;
    _ = pc ; _ = ok ;
    tmp := strings.Split(file,"/") ;
    err.file = tmp[len(tmp)-1] ;
    err.line = line ;

    return &err ;
}

