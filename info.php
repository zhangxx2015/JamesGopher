<?php

if ($argc > 1) {
    parse_str ( $argv [1], $param );
    foreach ( $param as $k => $v ) {
        echo "\r\n\tdefine: $k=$v";
        $param[$k] = $v;
		echo "\r\n\t>>>$k:$param[$k]";
    }
	$key="form";
	echo "\r\n\r\n\t$key:$param[$key]";
	
}

?>