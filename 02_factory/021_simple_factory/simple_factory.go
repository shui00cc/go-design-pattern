package main

import (
	"fmt"
)

type IRuleConfigParser interface {
	Parse(data []byte)
}

// 多个产品：以下是多个产品实现接口
// jsonRuleConfigParser 实现 Parse 接口
type jsonRuleConfigParser struct {
}

func (J jsonRuleConfigParser) Parse(data []byte) {
	panic("not implement")
}

// yamlRuleConfigParser 实现 Parse 接口
type yamlRuleConfigParser struct {
}

func (Y yamlRuleConfigParser) Parse(data []byte) {
	panic("not implement")
}

// 一个工厂：根据参数对应关系 返回一个对应的 IRuleConfigParser 接口
func NewIRuleConfigParser(t string) IRuleConfigParser {
	switch t {
	case "json":
		return jsonRuleConfigParser{}
	case "yaml":
		return yamlRuleConfigParser{}
	}
	return nil
}

func main() {
	// 创建json解析器
	jsonParser := NewIRuleConfigParser("json")
	fmt.Println(jsonParser)

	// 创建yaml解析器
	yamlParser := NewIRuleConfigParser("yaml")
	fmt.Println(yamlParser)

	// 创建未知类型
	unknownParser := NewIRuleConfigParser("unknown")
	fmt.Println(unknownParser)

	// 测试接口方法(panic占位,未实现)
	//jsonData := []byte(`{"name": "John", "age": 30}`)
	//jsonParser.Parse(jsonData)
}
