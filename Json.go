package PHPer4GoLang

import (
    "fmt"
_   "log"
    "errors"
    "reflect"
    "encoding/json"
)

type JsonClass struct {
    JsonClass string "JsonClass" ;
    MapJson map[string]interface{} "MapJson" ;
} ;

func (class *JsonClass) Getter(key string) (interface{},string){
    v := class.MapJson[key] ;

    t := Gettype(v) ;
    x := reflect.ValueOf(v) ;

    switch(x.Kind()){
    default:{
            t = fmt.Sprintf("%v",x.Kind()) ;
        }
    }

    _ = t ;

    return v.(interface{}),t ;
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

func (class *JsonClass) RAW() (map[string]interface{}){
    return class.MapJson ;
}

func Json_decode(i ... interface{}) (JsonClass,error){

    var ret JsonClass ;
    ret.JsonClass = "JsonClass" ;

    err := errors.New("ASSERT") ;
    err = nil ;

    if(len(i) == 1){
        str := Strval(i[0]) ;
        json.Unmarshal([]byte(str),&(ret.MapJson)) ;
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









