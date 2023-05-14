package PHPer4GoLang

import ("strings")

func Explode(v ... interface{}) []string {
    separator := v[0].(string) ;
    str := v[1].(string) ;

    return strings.Split(str,separator) ;
}
