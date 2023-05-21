package PHPer4GoLang

import (
    "fmt"
    "log"
    "errors"
    "reflect"
    "encoding/json"
)

type JsonClass struct {
    JsonClass string "JsonClass" ;
    Data map[string]interface{} "Data" ;
} ;

/*
func (class *JsonClass) Data() (map[string]interface{}){
    return class.Data ;
}
*/

func Json_decode(i ... interface{}) (JsonClass,error){

    var ret JsonClass ;

    err := errors.New("ASSERT") ;
    err = nil ;
    str := Strval(i[0]) ;

    ret.JsonClass = "JsonClass" ;

    json.Unmarshal([]byte(str),&(ret.Data)) ;

    if(false){
        log.Print(ret) ;
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









