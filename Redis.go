package PHPer4GoLang

import (
//    "log"
//    "strconv"
        "errors"
        "context"
//    "time"
//    "net/http"
//    "io/ioutil"
//    "crypto/tls"

    "github.com/redis/go-redis/v9"
)

type ClassRedis struct {
    host string ;
    port int ;

    client *redis.Client ;
} ;

func (class *ClassRedis) Connect(v ...any) (error){
    err := errors.New("ASSERT") ;
    switch(len(v)){
    case 1:{
        }
    case 2:{
            class.host = v[0].(string) ;
            class.port = v[1].(int) ;
        }
    }

    class.client = redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB:       0,
    }) ;

    err = nil ;

    return err ;
}

func (class *ClassRedis) Set(key string,val string) (error){
    err := errors.New("ASSERT") ;

    ctx := context.Background() ;

    err = class.client.Set(ctx, key, val, 0).Err()

    return err ;
}

func (class *ClassRedis) Get(key string) (string,error){
    err := errors.New("ASSERT") ;

    ctx := context.Background() ;

    val := "" ;

    val, err = class.client.Get(ctx,key).Result() ;

    return val,err ;
}

func (class *ClassRedis) Close() (error){
    err := errors.New("ASSERT") ;
    err = nil ;
    // Debugf("[%s][%d]",class.host,class.port) ;
    return err ;
}

func Redis(v ... any) (ClassRedis){
    var class ClassRedis ;

    class.host = "localhost" ;
    class.port = 6379 ;

    return class ;
}


















