package PHPer4GoLang

import (
    "os"
)

func Getenv(v interface{}) string {
    return os.Getenv(v.(string)) ;
}
