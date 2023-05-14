package PHPer4GoLang

import (
    "fmt"
    "runtime"
    "strings"
)

func Debugf(v ...interface{}){

    pc, file, line, ok := runtime.Caller(1) ;
    _ = pc ; _ = ok ;
    tmp := strings.Split(file,"/") ;
    basename := tmp[len(tmp)-1] ;

    // log.Print(v) ;

    out := "" ;
    if(true){
        strFmt := v[0].(string) ;
        switch(len(v)){
        case 1: out = strFmt ;
        case 2: out = fmt.Sprintf(strFmt,v[1]) ;
        case 3: out = fmt.Sprintf(strFmt,v[1],v[2]) ;
        case 4: out = fmt.Sprintf(strFmt,v[1],v[2],v[3]) ;
        case 5: out = fmt.Sprintf(strFmt,v[1],v[2],v[3],v[4]) ;
        case 6: out = fmt.Sprintf(strFmt,v[1],v[2],v[3],v[4],v[5]) ;
        case 7: out = fmt.Sprintf(strFmt,v[1],v[2],v[3],v[4],v[5],v[6]) ;
        case 8: out = fmt.Sprintf(strFmt,v[1],v[2],v[3],v[4],v[5],v[6],v[7]) ;
        case 9: out = fmt.Sprintf(strFmt,v[1],v[2],v[3],v[4],v[5],v[6],v[7],v[8]) ;
        }
    }

    fmt.Printf("%04d:%s:%s\n",line,basename,out) ;
}
