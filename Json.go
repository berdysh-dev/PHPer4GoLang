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
    MapJson map[string]interface{} "MapJson" ;
} ;


func (class *JsonClass) Getter(i ... interface{}) (interface{},string){


    if(len(i) == 1){
        v := class.MapJson[i[0].(string)] ;
        x := reflect.ValueOf(v) ;
        k := fmt.Sprintf("%v",x.Kind()) ;
        return v.(interface{}),k ;
    }else{
        var v interface{} ;
        return v,"error" ;
    }
}

func (class *JsonClass) GetterMap(key string) (map[string]interface{},string){
    v := class.MapJson[key] ;

    t := Gettype(v) ;

    return v.(map[string]interface{}),t ;
}

func (class *JsonClass) GetterArray(key string) ([]interface{},string){
    v := class.MapJson[key] ;
    t := Gettype(v) ;
    return v.([]interface{}),t ;
}

func upperCamel(uc string) (string){
    return strings.ToUpper(uc[:1]) + uc[1:] ;
}

func (class *JsonClass) RAW() (map[string]interface{}){
    return class.MapJson ;
}

func convertUpper_r (i interface{}) (interface{}){
    k,t := Gettype2(i) ;
    switch(k){
    case "map":{
            if(t == "map[string]interface {}"){
                hash := make(map[string]interface{}) ;
                for kk , vv := range i.(map[string]interface {}){
                    hash[upperCamel(kk)] = convertUpper_r(vv) ;
                }
                return hash ;
            }
        }
    case "slice":{
            if(t == "[]interface {}"){
                var ar []interface{} ;
                for _ , vv := range i.([]interface {}){
                    ar = append(ar,convertUpper_r(vv)) ;
                }
                return ar ;
            }
        }
    default:{
            if(k == "float64"){
                return uint64(i.(float64)) ;
            }else{
                return i ;
            }
            // Debugf("[%v][%v]",k,t);
        }
    }
    return nil ;
}

func ConvertUpper (src map[string]interface{}) (map[string]interface{}){

    ret := make(map[string]interface{}) ;

    for lc , v := range src{
        ret[upperCamel(lc)] = convertUpper_r(v) ;
    }

    return ret ;
}

func Json_decode(i ... interface{}) (JsonClass,error){

    var ret JsonClass ;
    ret.JsonClass = "JsonClass" ;

    err := errors.New("ASSERT") ;
    err = nil ;

    if(len(i) > 0){
        str := Strval(i[0]) ;
        json.Unmarshal([]byte(str),&(ret.MapJson)) ;
        if(len(i) > 1){
            ret.MapJson = ConvertUpper(ret.MapJson) ;
        }
    }else{
        ret.MapJson = make(map[string]interface{}) ;
    }

    return ret , err ;
}

func Json_encode(i ... interface{}) (string,error){
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

        var dec map[string]interface{} ;

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









