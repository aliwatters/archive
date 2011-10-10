<?php

function mb_isset($array,$count=0,$field=false){
	$is_set = false;
	if($field){
		if(isset($array[$field])){
			$array[] = $array[$field];
		}else{
			$array = false;
		}
	}
	if(isset($array)){
		if(is_array($array)){
			if(count($array)>$count){
				$is_set = true;
			}
		}
	} return $is_set;
}

function mb_load(){
	if(isset($GLOBALS['_mb_cache']['mb'])) return $GLOBALS['_mb_cache']['mb'];
    $mb = new MONGOBASE();
    $GLOBALS['_mb_cache']['mb'] = $mb;
    return $mb;
}