package PHPer4GoLang

import (
//    "log"
//    "strconv"
    "errors"
    "time"
    "net/http"
    "io/ioutil"
    "crypto/tls"
)

type cURLHandler struct {
    url string ;
    method string ;
    returntransfer int ;

    inited int ;

    http_Transport  http.Transport ;
    http_Client     http.Client ;
    http_Response   *http.Response ;
    TLS             *tls.ConnectionState ;
} ;

type cURLInfo struct {
    Url                         string ;
    Content_type                string ;
    Http_code                   int ;
    Header_size                 int ;
    Request_size                int ;
    Filetime                    int ;
    Ssl_verify_result           int ;
    Redirect_count              int ;
    Total_time                  int ;
    Namelookup_time             int ;
    Connect_time                int ;
    Pretransfer_time            int ;
    Size_upload                 int ;
    Size_download               int ;
    Speed_download              int ;
    Speed_upload                int ;
    Download_content_length     int64 ;
    Upload_content_length       int ;
    Starttransfer_time          int ;
    Redirect_time               int ;
    Certinfo                    []string ;
    Redirect_url                string ;
} ;

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

        ch.http_Transport.MaxIdleConns = 10 ;
        ch.http_Transport.IdleConnTimeout = 30 * time.Second ;
        ch.http_Transport.DisableCompression = true ;

        ch.http_Client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
            return http.ErrUseLastResponse
        } ;

        ch.http_Client.Transport = &(ch.http_Transport) ;

        ch.inited = 1 ;
    }

    http_NewRequest , err := http.NewRequest(ch.method,ch.url, nil) ;

    if(err != nil){
        Debugf("ERR[%v]",err) ;
    }else{
        http_NewRequest.Header.Add("If-None-Match", `W/"wyzzy"`) ;

        ch.http_Response , err = ch.http_Client.Do(http_NewRequest) ;

        if(err != nil){
            Debugf("ERR[%v]",err) ;
        }else{
            defer ch.http_Response.Body.Close() ;
            body, err := ioutil.ReadAll(ch.http_Response.Body) ;
            if(err != nil){
                Debugf("ERR[%v]",err) ;
            }else{
                payload = string(body) ;

                ch.TLS = ch.http_Response.TLS ;

                for k,vs := range ch.http_Response.Header{
                    for idx,v := range vs{
                        Debugf("[%v][%v][%v]",idx,k,v) ;
                    }
                }
            }
        }
    }

    return payload , err ;
}

func Curl_close(ch *cURLHandler){
}

func Curl_getinfo(ch *cURLHandler) (cURLInfo){
    var info cURLInfo ;

    // Debugf("ServerName[%v]",ch.TLS.ServerName) ;

    info.Http_code = ch.http_Response.StatusCode ;

    info.Download_content_length = ch.http_Response.ContentLength ;

    if(ch.TLS.HandshakeComplete){
        info.Ssl_verify_result = 1 ;
    }else{
        info.Ssl_verify_result = 0 ;
    }

    return info ;
}

















