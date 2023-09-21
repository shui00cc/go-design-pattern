package factory

import (
	"testing"
)

func TestNewIRuleConfigParser(t *testing.T) {
	// 测试返回json解析器
	jsonParser := NewIRuleConfigParser("json")
	if jsonParser == nil {
		t.Error("Expected jsonParser to be non-nil")
	}

	// 测试返回yaml解析器
	yamlParser := NewIRuleConfigParser("yaml")
	if yamlParser == nil {
		t.Error("Expected yamlParser to be non-nil")
	}

	// 测试未知类型
	unknownParser := NewIRuleConfigParser("unknown")
	if unknownParser != nil {
		t.Error("Expected unknownParser to be nil")
	}

	//测试接口方法
	//jsonData := []byte(`{"name": "John", "age": 30}`)
	//jsonParser.Parse(jsonData)
}
