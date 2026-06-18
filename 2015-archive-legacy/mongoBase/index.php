<?php

/* THIS IS AN EXAMPLE APP */

// This file is required and includes other files


//require_once(__DIR__.'/mb/mb_plugins.php');
//require_once(__DIR__.'/mb/mb_lang.php');
require_once(__DIR__.'/mb/mb_base.class.php');
require_once(__DIR__.'/mb/mb_db.class.php');
// require_once(__DIR__.'/mb/');

$mb = new MONGOBASE_DB;

var_dump($mb);

$add_object = array(
	'obj'	=> array(
		'key'	=> 'value'
	)
);

var_dump( $mb->mbsert($add_object) );

$edit_object = array(
	'obj'	=> array(
		'key'	=> time()
	),
	'id'	=> 'XXX' // Replace with valid ID to accomplish an edit
);
var_dump($mb->mbsert($edit_object));

var_dump($mb->find());

/* EXAMPLE OF DELETING AN OBJECT */
$delete_object = array(
	'id'	=> 'XXX' // Replace with valid ID to accomplish an edit
);
var_dump($mb->delete($delete_object));
