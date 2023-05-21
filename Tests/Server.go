package main

import ( "log" )

import  PHP "github.com/berdysh-dev/PHPer4GoLang"

func SelfTestDir(){
    // PHP.Debugf("テスト[%d][%d][%d]",1,2,3);

    if(false){
        PHP.Debugf("GOPATH[%s]",PHP.Getenv("GOPATH")) ;
        tmp := PHP.Explode("/",PHP.Getenv("GOPATH")) ;
        for idx, value := range tmp{
            PHP.Debugf("[%v][%v]",idx,value) ;
        }
    }

    dh,err := PHP.Opendir("/usr/local/src/Go/PHPer_local") ;

    if(err != nil){
        log.Fatal(err) ;
    }else{
        idx := 0 ;
        for{
            name,err := PHP.Readdir(&dh) ;
            if(err != nil){
                PHP.Debugf("%d:エラー[%v]",idx,err) ;
                break ;
            }else{
                st,_ := PHP.Stat(name) ;
                _ = st ;
                // log.Print(st) ;

                var mode int = 0 ;
                if(st.Mode.IsDir()){ mode = 1 ; }
                if(st.Mode.IsRegular()){ mode = 2 ; }

                PHP.Debugf("%d[%v]/mode[%v]/mtime[%v]/size[%v]",idx,name,mode,st.Mtime,st.Size) ;
            }
            idx += 1 ;
        }
        PHP.Closedir(&dh) ;
    }
}

func SelfTestPreg(){
    matches,err := PHP.Preg_match(`^(.)(.*)$`,`abc`) ;
    if(err != nil){
        PHP.Debugf("エラー[%v]",err) ;
    }else{
        for _, aaa := range matches{
            PHP.Debugf("[%s]",aaa) ;
         }
    }
    _ = matches ;
}

func SelfTestFopen(){

    str,err := PHP.File_get_contents("test") ;

    if(err != nil){
        PHP.Debugf("エラー[%v]",err) ;
    }else{
        PHP.Debugf("OK[%v]",str) ;
    }

    handle,err := PHP.Fopen("test", "r") ;
    if(err == nil){
        loop := 0 ;
        for {
            buf,err := PHP.Fread(&handle,3) ;
            if(err != nil){
                PHP.Debugf("%v:エラー[%v]",loop,err) ;
                break ;
            }else{
                PHP.Debugf("%v:OK[%v]",loop,buf) ;
            }
            loop += 1 ;
            if(loop > 3){ break ; }
        }
        PHP.Fclose(&handle) ;
    }
}

func SelfTestTime(){
    a := PHP.Gmtime(PHP.Time()) ;
    b := PHP.Localtime(PHP.Time()) ;

    year    := 2001 ;
    month   := 1 ;
    day     := 2 ;
    hour    := 3 ;
    minute  := 4 ;
    second  := 5 ;

    c := PHP.Localtime(PHP.Mktime(hour,minute,second,month,day,year)) ;

    PHP.Debugf("A[%04d/%02d/%02d %02d:%02d:%02d]",a.Tm_year+1900,a.Tm_mon+1,a.Tm_mday,a.Tm_hour,a.Tm_min,a.Tm_sec) ;
    PHP.Debugf("B[%04d/%02d/%02d %02d:%02d:%02d]",b.Tm_year+1900,b.Tm_mon+1,b.Tm_mday,b.Tm_hour,b.Tm_min,b.Tm_sec) ;
    PHP.Debugf("C[%04d/%02d/%02d %02d:%02d:%02d]",c.Tm_year+1900,c.Tm_mon+1,c.Tm_mday,c.Tm_hour,c.Tm_min,c.Tm_sec) ;
}

func SelfTestcURL(){

    ch,err := PHP.Curl_init() ;
    if(err != nil){

        count := 0 ;
        for{
            PHP.Curl_setopt(&ch,PHP.CURLOPT_URL,"http://berdysh.net/test.php") ;

            headers := []string{
                "Content-Type: application/json" ,
            } ;

//              "Content-Type: application/x-www-form-urlencoded"

//            PHP.Curl_setopt(&ch,PHP.CURLOPT_HTTPHEADER,headers) ;

            _ = headers ;

            PHP.Curl_setopt(&ch,PHP.CURLOPT_RETURNTRANSFER,1) ;
            PHP.Curl_setopt(&ch,PHP.CURLOPT_PROXYPORT,3128) ;
            PHP.Curl_setopt(&ch,PHP.CURLOPT_PROXY,"127.0.0.1") ;

            PHP.Curl_setopt(&ch,PHP.CURLOPT_POST,1);
            PHP.Curl_setopt(&ch,PHP.CURLOPT_POSTFIELDS,"{AAA: 123,BBB: 345}") ;

            payload,err := PHP.Curl_exec(&ch) ;

            if(err != nil){
                PHP.Debugf("ERR[&v]",err) ;
            }
            _ = payload ;

            info := PHP.Curl_getinfo(&ch) ;
            PHP.Debugf("STAT[%v]",info.Http_code) ;

            dec , err := PHP.Json_decode(payload) ;

            if(err == nil){
                if(true){
                    // log.Print(dec) ;
                    PHP.Json_encode(dec) ;
                }else{
                }
            }


            PHP.Curl_close(&ch) ;
            count += 1 ;
            if(count >= 1){ break ; }
        }
    }
}

func SelfTestRedis(){
    redis := PHP.Redis() ;

    err := redis.Connect("127.0.0.1",6379) ;
    if(err != nil){ PHP.Debugf("ERR[%v]",err) ; }

    err = redis.Set("KKK","kkk") ;
    if(err != nil){ PHP.Debugf("ERR[%v]",err) ; }

    v,err := redis.Get("KKK") ;
    if(err != nil){ PHP.Debugf("ERR[%v]",err) ; }

    PHP.Debugf("V[%v]",v) ;

    redis.Close() ;
}

func SelfTestPDO(){
    dsn         := "mysql:dbname=shop;host=127.0.0.1" ;
    user        := "root" ;
    password    := "root" ;

    dbh := PHP.PDO(dsn,user,password) ;
    prepare,err := dbh.Prepare("SQL") ;

    prepare.BindValue(":id", 3, PHP.PARAM_INT) ;
    prepare.Execute() ;
    result := prepare.FetchAll(PHP.FETCH_ASSOC) ;

    _ = prepare ;
    _ = err ;
    _ = result ;
}

func main(){

    PHP.Date_default_timezone_set("Asia/Tokyo") ;

    if(false){ SelfTestPreg() ; }
    if(false){ SelfTestDir() ; }
    if(false){ SelfTestFopen() ; }
    if(false){ SelfTestTime() ; }
    if(true){ SelfTestcURL() ; }
    if(false){ SelfTestRedis() ; }
    if(false){ SelfTestPDO() ; }
}





























