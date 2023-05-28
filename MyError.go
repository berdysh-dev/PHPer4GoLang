package PHPer4GoLang

import (
    "fmt"
    "runtime"
    "strings"
)

type MyErrorStruct struct {
    line    int ;
    file    string ;
    mess    string ;
}

func (err *MyErrorStruct) Error() (string) {
    return fmt.Sprintf("%04d:%s:%v",err.line,err.file,err.mess) ;
}

func MyErrorNew(v ... any) (error) {
    var err MyErrorStruct ;

    switch(len(v)){
    case 1: err.mess = v[0].(string) ;
    case 2: err.mess = fmt.Sprintf(v[0].(string),v[1]) ;
    case 3: err.mess = fmt.Sprintf(v[0].(string),v[1],v[2]) ;
    case 4: err.mess = fmt.Sprintf(v[0].(string),v[1],v[2],v[3]) ;
    case 5: err.mess = fmt.Sprintf(v[0].(string),v[1],v[2],v[3],v[4]) ;
    }

    pc, file, line, ok := runtime.Caller(1) ;
    _ = pc ; _ = ok ;
    tmp := strings.Split(file,"/") ;
    err.file = tmp[len(tmp)-1] ;
    err.line = line ;

    return &err ;
}

