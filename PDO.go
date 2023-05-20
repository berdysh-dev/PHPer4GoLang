package PHPer4GoLang

import (
//    "log"
//    "strconv"
        "errors"
//        "context"
//    "time"
//    "net/http"
//    "io/ioutil"
//    "crypto/tls"

    "github.com/redis/go-redis/v9"
)

const (
    PARAM_INT   = 1 ;
    FETCH_ASSOC = 2 ;
) ;

type ClassPDO struct {
    host string ;
    port int ;

    client *redis.Client ;
} ;

type ClassPrepare struct {
} ;


func (class *ClassPrepare) BindValue (v ...interface{}) (error){
    err := errors.New("ASSERT") ;
    return err ;
}

func (class *ClassPrepare) Execute (v ...interface{}) (error){
    err := errors.New("ASSERT") ;
    return err ;
}

func (class *ClassPrepare) FetchAll (v ...interface{}) (error){
    err := errors.New("ASSERT") ;
    return err ;
}

func (class *ClassPDO) Prepare(v ...interface{}) (ClassPrepare,error){
    err := errors.New("ASSERT") ;

    var prepare ClassPrepare ;

    return prepare,err ;
}

func (class *ClassPDO) Close() (error){
    err := errors.New("ASSERT") ;
    err = nil ;
    return err ;
}

func PDO(v ...interface{}) (ClassPDO){
    var class ClassPDO ;

    return class ;
}


















