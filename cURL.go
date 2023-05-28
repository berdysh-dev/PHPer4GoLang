package PHPer4GoLang

import (
//    "log"
//    "strconv"
    "errors"
    "time"
    "net/http"
    "strings"
    "io"
    "io/ioutil"
    "crypto/tls"
)

type cURLHandler struct {
    url string ;
    method string ;
    returntransfer bool ;

    inited int ;

    postfields string ;

    is_post bool ;
    is_header bool ;

    opt_header []string ;

    http_Transport  http.Transport ;
    http_Client     http.Client ;
    http_Response   *http.Response ;
    TLS             *tls.ConnectionState ;

    base_uri        string ;
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

type GuzzleResponse struct {
    ch          *cURLHandler ;
    method      string ;
    url         string ;
    payload     string ;
    stat        int ;
} ;

type GuzzleBody struct {
    res *GuzzleResponse ;
} ;

func GuzzleHttpClient(anys ... any) (cURLHandler,error){
    err := errors.New("ASSERT") ;
    var ch cURLHandler ;
    ch.method = "GET" ;

    if(len(anys) > 0){
        for kk,vv := range anys[0].(map[string]any){
            Debugf("[%v][%v]",kk,vv);

            switch(kk){
            case "base_uri":{
                    ch.base_uri = vv.(string) ;
                }
            }
        }
    }

    err = nil ;
    return ch,err ;
}

func (ch *cURLHandler) Request(anys ... any) (GuzzleResponse,error) {
    err := MyErrorNew("HOHO[%v]",123) ;

    var res GuzzleResponse ;
    res.ch = ch ;

    if(len(anys) > 0){
        res.method = anys[0].(string) ;
    }

    if(len(anys) > 1){
        res.url = anys[1].(string) ;
    }

    if(len(anys) > 2){
        for kk,vv := range anys[2].(map[string]any){
            _ = vv ;
            switch(kk){
            case "http_errors":{
                }
            case "allow_redirects":{
                }
            case "headers":{
                }
            case "json":{
                }
            case "form_params":{
                }
            }
        }
    }

    res.payload = "RESrES" ;
    res.stat = 201 ;

    return res,err ;
}

func (res *GuzzleResponse) GetStatusCode (anys ... any) int {
    return res.stat ;
}

func (res *GuzzleResponse) GetBody(anys ... any) GuzzleBody {
    var body GuzzleBody ;
    body.res = res ;
    return body ;
}

func (body *GuzzleBody) GetContents(anys ... any) string {
    return body.res.payload ;
}

func Curl_init() (cURLHandler,error){
    err := errors.New("ASSERT") ;
    var ch cURLHandler ;

    ch.method = "GET" ;

    err = nil ;
    return ch,err ;
}

func Curl_setopt(v ...any){

    var ch *cURLHandler = v[0].(*cURLHandler) ;

    switch(v[1]){
    case CURLOPT_URL:{
            ch.url = Strval(v[2]) ;
            Debugf("CURLOPT_URL[%s]",ch.url) ;
        }
    case CURLOPT_HTTPHEADER:{
            ch.opt_header = v[2].([]string) ;
        }
    case CURLOPT_RETURNTRANSFER:{
            ch.returntransfer = Boolval(v[2]) ;
        }
    case CURLOPT_HEADER:{
            ch.is_header = Boolval(v[2]) ;
        }
    case CURLOPT_POST:{
            ch.is_post = Boolval(v[2]) ;
        }
    case CURLOPT_POSTFIELDS:{
            ch.postfields = Strval(v[2]) ;
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

    method := ch.method ;

    var io_Reader_Post io.Reader = nil ;

    if(ch.is_post == true){
        method = "POST" ;
        io_Reader_Post = strings.NewReader(ch.postfields) ;
    }

    http_NewRequest , err := http.NewRequest(method,ch.url,io_Reader_Post) ;

    if(err != nil){
        Debugf("ERR[%v]",err) ;
    }else{
        if(false){
            http_NewRequest.Header.Set("Content-Type", "application/x-www-form-urlencoded")
            http_NewRequest.Header.Set("Content-Type", "application/json")
        }

        if(len(ch.opt_header) > 0){
            for _,line := range ch.opt_header{
                match,err := Preg_match(`^([^\:]+)\s*:\s*(.*)$`,line) ;
                if(err == nil){
                    http_NewRequest.Header.Set(match[1],match[2]) ;
                }
            }
        }

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
                        _ = idx ;
                        _ = k ;
                        _ = v ;
                        // Debugf("[%v][%v][%v]",idx,k,v) ;
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

    if(ch.TLS != nil){
        if(ch.TLS.HandshakeComplete){
            info.Ssl_verify_result = 1 ;
        }else{
            info.Ssl_verify_result = 0 ;
        }
    }

    return info ;
}

















