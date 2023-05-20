package PHPer4GoLang

//    "io/fs"
//    "io/ioutil"
//    "log"

import (
    "os"
    "errors"
    "io/fs"
)

type DirHandler struct {
    path    string ;
    fs_FS   fs.FS ;
    paths   []string ;
    idx     int ;
}

func Opendir(path string) (DirHandler,error) {

    var dh DirHandler ;
    var err error ;

/*
    https://text.baldanders.info/golang/abstract-filesystem/
*/

    dh.idx = 0 ;
    dh.fs_FS = os.DirFS(path) ;

    dh.paths, err = fs.Glob(dh.fs_FS, "*") ;

    if err != nil {
        Debugf("Error[%v]",err) ;
    }else{
        // log.Print(dh.paths);
        // for idx, name := range dh.paths{ Debugf("Path[%v][%v]",idx,name) ; }
    }

    return dh,err ;

/*
    files, err := ioutil.ReadDir(path) ;
    _ = files ;
    for _, fs_FileInfo := range files{ Debugf("Name[%s]",fs_FileInfo.Name()); }
*/

}

func Readdir(dh *DirHandler) (string,error) {

    err := errors.New("EOF") ;

    // Debugf("[%d/%d]",dh.idx,len(dh.paths)) ;

    if(dh.idx >= len(dh.paths)){
        return "EOF" , err ;
    }else{
        err = nil ;
        name := dh.paths[dh.idx] ;
        dh.idx += 1 ;
        return name , err ;
    }
}

func Closedir(dh *DirHandler){
    return ;
}

