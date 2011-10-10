<?php

class MONGOBASE {

	private function mongo(){
		/*
		For authentication and security, mongodb need to execute/run with --auth parameter
		Reference: http://www.mongodb.org/display/DOCS/Security+and+Authentication
		*/
		try{
			$mb = mb_load();
			$options = $mb->options();
			if(isset($GLOBALS['_mb_cache']['mongo'])) return $GLOBALS['_mb_cache']['mongo'];
			if ($options['db_replicas'] === true && $options['db_user'] !==''){
				//replica and database need authentication
				$m = new Mongo("mongodb://{$options['db_user']}:{$options['db_pass']}@{$options['db_host']}:{$options['db_port']}/{$options['db_name']}", array('replicaSet' => true));
			} elseif ($options['db_replicas'] === true) {
				//replica set and no auth.
				$m = new Mongo("mongodb://{$options['db_host']}:{$options['db_port']}/{$options['db_name']}", array('replicaSet' => true));

			} elseif ($options['db_user']!=false) { //database need authentication
				$m = new Mongo("mongodb://{$options['db_user']}:{$options['db_pass']}@{$options['db_host']}:{$options['db_port']}/{$options['db_name']}");

			} else {
				//default without auth and replica
				$m = new Mongo("mongodb://{$options['db_host']}:{$options['db_port']}/{$options['db_name']}");

			}
			$GLOBALS['_mb_cache']['mongo'] = $m;
			return $m;
		} catch (MongoConnectionException $e) {
            return(__('Error connecting to MongoDB server'));
        } catch (MongoException $e) {
            return(__('Error: ').$e->getMessage());
        } catch (MongoCursorException $e) {
            return(__('Error: probably username password in config').$e->getMessage());
        }
	}

	private function register_configuration_setting($definition = false, $constant = false, $default = null){
		if($default===null) $default = $constant;
		if(($definition)&&($constant)){
			if(defined($definition)){ $this_setting = $constant; }else{ $this_setting = $default; }
			return $this_setting;
		}else{
			if($definition){
				if(defined($definition)){ $this_setting = $default; }
				return $this_setting;
			}else{
				return $default;
			}
		}
	}

	private function arrayed($these_objs){
        if(is_object($these_objs)){
            $objects = array();
            foreach($these_objs as $this_obj) {
                $this_object = array();
                foreach($this_obj as $key => $value){
                    $this_object[$key] = $value;
                } $objects[] = $this_object;
            }
            if(is_array($objects)){
                if(!empty($objects)){
                    return $objects;
                }
            }
        }
    }
	
	public function options(){

        /* IF GOT OPTIONS IN CACHE, RETURN THEM */
		if(isset($GLOBALS['_mb_cache']['mb_options'])){
            return $GLOBALS['_mb_cache']['mb_options'];
        }else{

            /* ELSE - GATHER CONFIG SETTINGS */
			if(!defined('MONGODB_NAME')) define('MONGODB_NAME', 'mongobase');
			$mb_options['db_name'] = $this->register_configuration_setting('MONGODB_NAME', MONGODB_NAME);

			if(!defined('MONGODB_HOST')) define('MONGODB_HOST', 'localhost');
			$mb_options['db_host'] = $this->register_configuration_setting('MONGODB_HOST', MONGODB_HOST);

			if(!defined('MONGODB_USERNAME')) define('MONGODB_USERNAME', false);
			$mb_options['db_user'] = $this->register_configuration_setting('MONGODB_USERNAME', MONGODB_USERNAME);

			if(!defined('MONGODB_PASSWORD')) define('MONGODB_PASSWORD', false);
			$mb_options['db_pass'] = $this->register_configuration_setting('MONGODB_PASSWORD', MONGODB_PASSWORD);

			if(!defined('MONGODB_PORT')) define('MONGODB_PORT', '27017');
			$mb_options['db_port'] = $this->register_configuration_setting('MONGODB_PORT', MONGODB_PORT);

			if(!defined('MONGODB_REPLICAS')) define('MONGODB_REPLICAS', false);
			$mb_options['db_replicas'] = $this->register_configuration_setting('MONGODB_REPLICAS', MONGODB_REPLICAS);
			
			/* STORE AND RETURN OPTIONS */
            $GLOBALS['_mb_cache']['mb_options'] = $mb_options;
            return $mb_options;

        }
    }

	/* mbsert() allow for intelligent inserting and (or) updating */
	public function mbsert($options = false){
		var_dump($options);
		$default_options = array(
            'col'	=> 'mbsert',
            'obj'   => false,
			'id'	=> false
        );
        if(is_array($options)){
            $settings = array_merge($default_options,$options);
        }else{
            $settings = $default_options;
        }

		try{

			/* GOT OBJECT */
			$m = $this->mongo();
			$options = $this->options();
			$db = $m->$options['db_name'];
			$collection = $db->$settings['col'];
			$mongo_id = new MongoID($settings['id']);
			$key = array("_id"=>$mongo_id);
			$data = $settings['obj'];
			$results = $db->command( array(
				'findAndModify' => $settings['col'],
				'query' => $key,
				'update' => $data,
				'new' => true,
				'upsert' => true,
				'fields' => array( '_id' => 1 )
			) );
			return $results['value']['_id'];

		} catch (MongoConnectionException $e) {
            return(__('Error connecting to MongoDB server'));
        } catch (MongoException $e) {
            return(__('Error: ').$e->getMessage());
        } catch (MongoCursorException $e) {
            return(__('Error: probably username password in config').$e->getMessage());
        }

	}

	public function find($options = false){
		$default_options = array(
            'col'		=> 'mbsert',
            'where'		=> array(),
			'limit'		=> false,
			'offset'	=> false,
			'order_by'	=> false,
			'order'		=> false,
			'id'		=> false
        );
        if(is_array($options)){
            $settings = array_merge($default_options,$options);
        }else{
            $settings = $default_options;
        } if($settings['order_by']){
			if($settings['order']!='desc'){ $order_value=1; }else{ $order_value=-1; }
			$sort_clause = array($order_by=>$order_value);
		}else{ $sort_clause = array(); }

		try{

			$m = $this->mongo();
			$options = $this->options();
			$db = $m->$options['db_name'];
			$collection = $db->$settings['col'];
			$results = $this->arrayed($collection->find($settings['where'])->sort($sort_clause)->skip($settings['offset'])->limit($settings['limit']));
			return $results;

		} catch (MongoConnectionException $e) {
            return(__('Error connecting to MongoDB server'));
        } catch (MongoException $e) {
            return(__('Error: ').$e->getMessage());
        } catch (MongoCursorException $e) {
            return(__('Error: probably username password in config').$e->getMessage());
        }
	}

	public function delete($options = false){
		$default_options = array(
            'col'		=> 'mbsert',
			'id'		=> false
        );
        if(is_array($options)){
            $settings = array_merge($default_options,$options);
        }else{
            $settings = $default_options;
        }

		try{

			$m = $this->mongo();
			$options = $this->options();
			$db = $m->$options['db_name'];
			$collection = $db->$settings['col'];
			$criteria = array(
				'_id' => new MongoId($settings['id']),
			);
			$progress = $collection->remove($criteria,array('safe'=>true));
			return $progress['n'];

		} catch (MongoConnectionException $e) {
            return(__('Error connecting to MongoDB server'));
        } catch (MongoException $e) {
            return(__('Error: ').$e->getMessage());
        } catch (MongoCursorException $e) {
            return(__('Error: probably username password in config').$e->getMessage());
        }
	}

}