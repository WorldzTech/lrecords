package record

import "github.com/WorldzTech/lrecords/schema"

type Record struct {
	Schema *schema.RecordSchema `json:"schema"`
	Values []interface{}        `json:"values"`
}

func (r *Record) GetValueByName(name string) (interface{}, bool) {
	for i, field := range r.Schema.Fields {
		if field.Name == name {
			return r.Values[i], true
		}
	}
	return nil, false
}

func (r *Record) GetValueByIndex(index int) (interface{}, bool) {
	if index < 0 || index >= len(r.Values) {
		return nil, false
	}
	return r.Values[index], true
}

func (r *Record) SetValueByName(name string, value interface{}) bool {
	for i, field := range r.Schema.Fields {
		if field.Name == name {
			r.Values[i] = value
			return true
		}
	}
	return false
}

func (r *Record) SetValueByIndex(index int, value interface{}) bool {
	if index < 0 || index >= len(r.Values) {
		return false
	}
	r.Values[index] = value
	return true
}

func (r *Record) GetSchema() *schema.RecordSchema {
	return r.Schema
}

func (r *Record) GetValues() []interface{} {
	return r.Values
}

func (r *Record) Validate() bool {
	if len(r.Values) != len(r.Schema.Fields) {
		return false
	}
	for i, field := range r.Schema.Fields {
		value := r.Values[i]
		switch field.Type {
		case "string":
			if _, ok := value.(string); !ok {
				return false
			}
		case "int":
			if _, ok := value.(int); !ok {
				return false
			}
		case "float":
			if _, ok := value.(float64); !ok {
				return false
			}
		case "bool":
			if _, ok := value.(bool); !ok {
				return false
			}
		default:
			return false
		}
	}
	return true
}
