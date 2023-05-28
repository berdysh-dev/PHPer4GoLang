package PHPer4GoLang

import (
    "os"
)

func Getenv(v any) string {
    return os.Getenv(v.(string)) ;
}
