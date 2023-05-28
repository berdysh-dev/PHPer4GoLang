package PHPer4GoLang

import (
_   "log"
    "fmt"
    "strings"
    "errors"
    "reflect"
    "encoding/json"
)

type JsonClass struct {
    JsonClass string "JsonClass" ;
    MapJson map[string]any "MapJson" ;
} ;


func (class *JsonClass) Getter(anys ... any) (any,string){
    if(len(anys) == 1){
        v := class.MapJson[anys[0].(string)] ;
        x := reflect.ValueOf(v) ;
        k := fmt.Sprintf("%v",x.Kind()) ;
        return v.(any),k ;
    }else{
        var v any ;
        return v,"error" ;
    }
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

func checker(name string,v any) (any){
    k := reflect.ValueOf(v).Kind() ;
    Debugf("Check[%v][%v][%v]",name,k,v) ;
    return v ;
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









