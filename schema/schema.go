package schema

import (
	"encoding/json"
	"os"
)

type RecordField struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type RecordSchema struct {
	Fields []RecordField `json:"fields"`
}

func NewRecordSchemaFromFile(file *os.File) (*RecordSchema, error) {
	var s RecordSchema
	if err := json.NewDecoder(file).Decode(&s); err != nil {
		return nil, err
	}
	return &s, nil
}

func (s *RecordSchema) GetFieldNames() []string {
	names := make([]string, len(s.Fields))
	for i, field := range s.Fields {
		names[i] = field.Name
	}
	return names
}

func (s *RecordSchema) GetFieldTypes() []string {
	types := make([]string, len(s.Fields))
	for i, field := range s.Fields {
		types[i] = field.Type
	}
	return types
}

func (s *RecordSchema) GetFieldTypeByName(name string) (string, bool) {
	for _, field := range s.Fields {
		if field.Name == name {
			return field.Type, true
		}
	}
	return "", false
}
