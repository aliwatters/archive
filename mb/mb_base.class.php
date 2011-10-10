<?php


class MONGOBASE {

    /* VERY VERY GENERIC CLASS THAT DOESNT DO ANYTHING -- JUST A FRAME WORK FOR SUBCLASSES */

    public $options = array();

    function __construct(){
        $this -> options();
    }
    
    function get_option($key) {
        return $this -> options[$key];
    }

    function set_option($key,$val) {
        $this->options[$key] = $val;
        return $val;
    }


    public function __($key) {
        if (function_exists('__')) return __($key) . "\n";
        
        return $key. "\n";
    }
    
    
	public function register_configuration_setting($key, $definition = false, $constant = false, $default = null){
    /* requires that values are defined somewhere - probably in the configuration config module */
		if ($default===null) $default = $constant;

		if( $definition !== false && $constant !== false){

			if (defined($definition)) $val = $constant;
            else $val = $default; 

		}else{

			if (defined($definition)) { $val = $default; }

		}
        $this -> set_option($key,$val);    
	}


    public function options() {

        if ($this->options !== null) return $this->options;

        $this -> options = array();

        if(!defined('TEST')) define('TEST', 'value');
		$this->register_configuration_setting('key','TEST', TEST);

        $this->options = $options;
        return $options;
    }

}


?>
