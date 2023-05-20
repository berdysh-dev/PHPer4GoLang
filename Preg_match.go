package PHPer4GoLang

//    "io/fs"
//    "io/ioutil"
//    "log"
//    "os"
//    "fmt"

import (
    "errors"
    "regexp"
)

func Preg_match(pattern string,subject string) ([]string,error) {
    var err error = nil ;

    if(false){
        err = errors.New("ERR") ;
    }

    var exp = regexp.MustCompile(pattern) ;
    matches := exp.FindAllStringSubmatch(subject,-1) ;

    return matches[0] , err ;
}


