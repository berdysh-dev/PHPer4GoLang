ckage PHPer4GoLang

type MyErrorStruct struct {
    line    int ;
    file    string ;
    mess    string ;
}

func (err * MyErrorStruct) Error() string {
    return fmt.Sprintf("HOGE") ;
}

func MyErrorNew(mess string) MyErrorStruct {
    var err MyErrorStruct ;
    err.mess = mess ;
    return err ;
}
