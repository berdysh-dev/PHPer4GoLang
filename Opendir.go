package PHPer4GoLang

import (
    "errors"
)

func Opendir(v interface{}) (string,error) {
    err := errors.New("AAA") ;
    _ = err ;
    return "DH" , nil ;
}

func Readdir(v interface{}) (string,error) {
    err := errors.New("BBB") ;
    return "NAME" , err ;
}

func Closedir(v interface{}){
    return ;
}

