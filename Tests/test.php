<?php

    date_default_timezone_set("Asia/Tokyo") ;

    class time{

        function test(){
            $b = localtime(time(),1) ;

            $year    = 2001 ;
            $month   = 1 ;
            $day     = 2 ;
            $hour    = 3 ;
            $minute  = 4 ;
            $second  = 5 ;

            $c = localtime(mktime($hour,$minute,$second,$month,$day,$year),1) ;

            $d = localtime(gmmktime($hour,$minute,$second,$month,$day,$year),1) ;

            print_r($b) ;
            print_r($c) ;
            print_r($d) ;

            printf("B[%04d/%02d/%02d %02d:%02d:%02d]\n",$b['tm_year']+1900,$b['tm_mon']+1,$b['tm_mday'],$b['tm_hour'],$b['tm_min'],$b['tm_sec']) ;
            printf("C[%04d/%02d/%02d %02d:%02d:%02d]\n",$c['tm_year']+1900,$c['tm_mon']+1,$c['tm_mday'],$c['tm_hour'],$c['tm_min'],$c['tm_sec']) ;
            printf("D[%04d/%02d/%02d %02d:%02d:%02d]\n",$d['tm_year']+1900,$d['tm_mon']+1,$d['tm_mday'],$d['tm_hour'],$d['tm_min'],$d['tm_sec']) ;
        }
    }

    class curl{
        function test(){
            $db = [] ;
            foreach(get_defined_constants() as $k => $v){
                if(substr($k,0,4) === 'CURL'){
                    $db[$k] = $v ;
                }
            }
            ksort($db) ;
            foreach($db as $k => $v){
                printf("    %s = %d ;\n",$k,$v) ;
            }
        }
    }

    if(0){ if($ctx = new time()){ $ctx->test() ; } }
    if(1){ if($ctx = new curl()){ $ctx->test() ; } }
