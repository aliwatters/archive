<?php

class MONGOBASE {

	/* VERY VERY GENERIC CLASS THAT DOESNT DO ANYTHING -- JUST A FRAME WORK FOR SUBCLASSES */

	public $options = array();

	function __construct(){
		$this->options();
		// Subclasses should call this with parent::__construct();
	}

	public function settings($options,$defaults = array()){
		if(! $options === false && is_array($options)){
			return array_merge($default,$options);
		}
		return $defaults;
	}

	public function get_option($key){
		return $this->options[$key];
	}

	public function set_option($key,$val){
		$this->options[$key] = $val;
		return $val;
	}


	public function __($key){
		if (function_exists('__')) return __($key) . "\n";
		return $key. "\n";
	}

	public function register_configuration_setting($key, $definition = false, $constant = false, $default = null){
		/* requires that values are defined somewhere - probably in the config module */
		if($default===null) $default = $constant;
		if($definition !== false && $constant !== false){
			if(defined($definition)) $val = $constant;
			else $val = $default;
		}else{
			$val = $default;
		}
		$this -> set_option($key,$val);
	}

	public function options(){
		if($this->options !== null) return $this->options;
		$this->options = array();
		if(!defined('DEMO_CONFIG')) define('DEMO_CONFIG', 'default_value');
		$this->register_configuration_setting('demo_key', 'DEMO_CONFIG', DEMO_CONFIG);
		$this->options = $options;
		return $options;
	}

}