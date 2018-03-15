<?php

//echo "\r\ntest from php\r\n";

date_default_timezone_set("Asia/Shanghai");
echo "\r\n".date("Y-m-d H:i:s a l");

if ($argc > 1) {
    parse_str ( $argv [1], $param );
    foreach ( $param as $k => $v ) {
        echo "\r\ndefine $k=$v";
        $param[$k] = $v;
    }
}
sleep(1)
?>