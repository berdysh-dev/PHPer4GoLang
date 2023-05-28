package PHPer4GoLang

import (
_   "log"
    "fmt"
    "strings"
    "strconv"
    "errors"
    "reflect"
    "encoding/json"
)

type JsonClass struct {
    JsonClass string "JsonClass" ;
    MapJson map[string]any "MapJson" ;
} ;

func (class *JsonClass) Getter(anys ... any) (any,error){
    var k any ;
    var kind string ;
    var x any ;

    x = class.MapJson ;

    var args []any ;

    if(len(anys) == 1){
        t := fmt.Sprintf(`%T`,anys[0]) ;
        switch(t){
        case "[]string":{
                for _ , v := range anys[0].([]string){
                    args = append(args,v) ;
                }
            }
        case "[]interface {}":{
                for _ , v := range anys[0].([]interface {}){
                    args = append(args,v) ;
                }
            }
        case "[]any":{
                for _ , v := range anys[0].([]any){
                    args = append(args,v) ;
                }
            }
        default:{
                for _ , v := range anys[0].([4]string){
                    Debugf("!!![%v]",v) ;
                }
                return nil,errors.New(t) ;
            }
        }
    }else{
        args = anys ;
    }

    for iii:=0;iii<len(args);iii++ {

        switch(reflect.ValueOf(args[iii]).Kind()){
        case reflect.String:{
                k = args[iii].(string) ;
            }
        case reflect.Int:{
                k = args[iii].(int) ;
            }
        default:{
                Debugf("Other") ;
            }
        }

        kind = fmt.Sprintf("%v",reflect.ValueOf(x).Kind()) ;

        _ = kind ;
        _ = k ;

        // Debugf("kind[%v]",kind);

        switch(kind){
        case "map":{
                tmp := x.(map[string]any) ;
                v,ok := tmp[k.(string)] ;
                if(ok){
                    x = v ;
                    kind = fmt.Sprintf("%v",reflect.ValueOf(x).Kind()) ;
                }else{
                    return nil,errors.New("NotExists") ;
                }
            }
        case "slice":{
                tmp := x.([]any) ;

                kindK := fmt.Sprintf("%v",reflect.ValueOf(k).Kind()) ;

                if(kindK == "string"){
                    n,err := strconv.Atoi(k.(string)) ;
                    if(err == nil){
                        k = n ;
                    }else{
                        return nil,err ;
                    }
                }

                kindK = fmt.Sprintf("%v",reflect.ValueOf(k).Kind()) ;

                // Debugf("[%v]",kindK) ;

                if(kindK != "int"){
                    return nil,errors.New("NotInt") ;
                }
                
                if(k.(int) < 0){
                    k = (len(tmp) + k.(int)) ;

                    Debugf("!!!!!![%v]",k);
                }

                if(k.(int) < len(tmp)){
                    v := tmp[k.(int)] ;
                    x = v ;
                    kind = fmt.Sprintf("%v",reflect.ValueOf(x).Kind()) ;
                }else{
                    return nil,errors.New("TooBig") ;
                }
            }
        }
    }

    return x,nil ;
}

func (class *JsonClass) GetterMap(key string) (map[string]any,string){
    v := class.MapJson[key] ;

    t := Gettype(v) ;

    return v.(map[string]any),t ;
}

func (class *JsonClass) GetterArray(key string) ([]any,string){
    v := class.MapJson[key] ;
    t := Gettype(v) ;
    return v.([]any),t ;
}

func convertUpper(uc string) (string){
    return strings.ToUpper(uc[:1]) + uc[1:] ;
}

func convertLower(uc string) (string){
    return strings.ToLower(uc[:1]) + uc[1:] ;
}

func (class *JsonClass) Raw() (map[string]any){
    return class.MapJson ;
}

func checker(name string,v any) (error){
    err := errors.New("ASSERT") ;

    k := reflect.ValueOf(v).Kind() ;

    _ =  k ;

    //Debugf("Check[%v][%v][%v]",name,k,v) ;

    if(true){
        err = nil ;
    }

    return err ;
}

func convertUpper_r (src any) (any){
    switch(reflect.ValueOf(src).Kind()){
    case reflect.Map:{
            tmp := make(map[string]any) ;
            for k , v := range src.(map[string]any){

                kk := convertUpper(k) ;
                vv := convertUpper_r(v) ;

                err := checker(kk,vv) ;

                if(err == nil){
                    tmp[kk] = vv ;
                }
            }
            return tmp ;
        }
    case reflect.Slice:{
            var tmp []any ;
            for _ , v := range src.([]any){
                tmp = append(tmp,convertUpper_r(v)) ;
            }
            return tmp ;
        }
    case reflect.Float64:
        return uint64(src.(float64)) ;
    }
    return src ;
}

func Json_decode(i ... any) (JsonClass,error){

    var ret JsonClass ;
    ret.JsonClass = "JsonClass" ;

    err := errors.New("ASSERT") ;
    err = nil ;

    if(len(i) > 0){
        str := Strval(i[0]) ;
        json.Unmarshal([]byte(str),&(ret.MapJson)) ;
        if(len(i) > 1){
            ret.MapJson = convertUpper_r(ret.MapJson).(map[string]any) ;
        }
    }else{
        ret.MapJson = make(map[string]any) ;
    }

    return ret , err ;
}

func Json_encode(i ... any) (string,error){
    ret := "" ;

    err := errors.New("ASSERT") ;

    v := reflect.ValueOf(i[0]) ;

    if(v.Kind() == reflect.Struct){
        t := v.Type() ;
        for i := 0; i < t.NumField(); i++ {
            sf := t.Field(i)
            fmt.Printf("sf.Index=%v\n", sf.Index)
            fmt.Printf("sf.Name=%v\n", sf.Name)
            fmt.Printf("sf.Type=%v\n", sf.Type)
            fmt.Printf("sf.Tag=%v\n", sf.Tag)
            fmt.Printf("------------------\n")
        }
    }

    if(false){

        var dec map[string]any ;

        bytes, err := json.Marshal(dec) ;
        if(err == nil){
            ret = string(bytes) ;
        }else{
            Debugf("ERR[%v]",err);
        }
    }
    return ret,err ;
}

/*
    https://tech.yappli.io/entry/go_unmarshal_interface
    https://golangstart.com/structure_interface/
*/









