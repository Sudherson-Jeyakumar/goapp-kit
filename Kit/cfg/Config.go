package cfg

type Config interface {
	Set(value interface{}, path ...string)
	Value(...string) interface{}
}

type simpleConfig map[interface{}]interface{}

func Simple() simpleConfig {
	return map[interface{}]interface{}{}
}

func setInternal(s map[interface{}]interface{}, value interface{}, path ...string) {

	if len(path) == 1 {
		s[path[0]] = value
		return
	} else {
		if _, ok := s[path[0]]; !ok {
			s[path[0]] = map[interface{}]interface{}{}
		}
	}

	setInternal(s[path[0]].(map[interface{}]interface{}), value, path[1:]...)
}

func (s simpleConfig) Set(value interface{}, path ...string) {
	if len(path) == 0 {
		return
	}

	setInternal(s, value, path...)
}

func valueInternal(s map[interface{}]interface{}, path ...string) interface{} {
	v, ok := s[path[0]]
	if !ok {
		return nil
	}

	if len(path) == 1 {
		return s[path[0]]
	}

	if v1, ok := v.(map[interface{}]interface{}); !ok {
		return nil
	} else {
		return valueInternal(v1, path[1:]...)
	}
}

func (s simpleConfig) Value(path ...string) interface{} {

	if len(path) == 0 {
		return nil
	}

	return valueInternal(s, path...)
}
