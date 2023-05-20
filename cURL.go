package PHPer4GoLang

import (
    "log"
    "errors"
//    "time"
    "net/http"
    "io/ioutil"
)

type cURLHandler struct {
    url string ;
    method string ;
    returntransfer int ;

    inited int ;

    http_Transport  http.Transport ;
    http_Client     http.Client ;
}


func Curl_init() (cURLHandler,error){
    err := errors.New("ASSERT") ;
    var ch cURLHandler ;

    ch.method = "GET" ;

    return ch,err ;
}

func Curl_setopt(v ...interface{}){

    var ch *cURLHandler = v[0].(*cURLHandler) ;

    switch(v[1]){
    case CURLOPT_URL:{
            Debugf("CURLOPT_URL[%v]",v[2]);
            ch.url = v[2].(string) ;
        }
    case CURLOPT_RETURNTRANSFER:{
            ch.returntransfer = v[2].(int) ;
        }
    }
}

func Curl_exec(ch *cURLHandler) (string,error){

    payload := "" ;

    err := errors.New("ASSERT") ;

    if(ch.inited != 1){

/*
        ch.http_Transport = &http.Transport{
            MaxIdleConns:       10,
            IdleConnTimeout:    30 * time.Second,
            DisableCompression: true,
        } ;

        ch.http_Client = &http.Client{
            CheckRedirect: func(req *http.Request, via []*http.Request) error {
                return http.ErrUseLastResponse
            },
            Transport: ch.http_Transport,
        } ;
*/

        ch.inited = 1 ;
    }

    http_NewRequest , err := http.NewRequest(ch.method,ch.url, nil) ;

    if(err != nil){
        Debugf("ERR[%v]",err) ;
    }else{
        http_NewRequest.Header.Add("If-None-Match", `W/"wyzzy"`) ;

        http_Response , err := ch.http_Client.Do(http_NewRequest) ;

        if(err != nil){
            Debugf("ERR[%v]",err) ;
        }else{
            defer http_Response.Body.Close() ;
            body, err := ioutil.ReadAll(http_Response.Body) ;
            if(err != nil){
                Debugf("ERR[%v]",err) ;
            }else{
                payload = string(body) ;
                log.Print(http_Response.Status) ;
            }
        }
    }

    return payload , err ;
}

func Curl_close(ch *cURLHandler){
}


















