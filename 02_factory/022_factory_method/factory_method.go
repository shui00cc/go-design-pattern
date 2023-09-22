package main

import "fmt"

// 抽象产品
type IRuleConfigParser interface {
	Parse(data []byte)
}

// 具体产品 及其方法实现
type jsonRuleConfigParser struct {
}

func (J jsonRuleConfigParser) Parse(data []byte) {
	panic("implement me")
}

type yamlRuleConfigParser struct {
}

func (Y yamlRuleConfigParser) Parse(data []byte) {
	panic("implement me")
}

// 抽象工厂
type IRuleConfigParserFactory interface {
	CreateParser() IRuleConfigParser
}

// 具体工厂
type yamlRuleConfigParserFactory struct {
}

func (y yamlRuleConfigParserFactory) CreateParser() IRuleConfigParser {
	fmt.Println("yaml解析器的具体创建逻辑")
	return yamlRuleConfigParser{}
}

type jsonRuleConfigParserFactory struct {
}

func (j jsonRuleConfigParserFactory) CreateParser() IRuleConfigParser {
	fmt.Println("json解析器的具体创建逻辑")
	return jsonRuleConfigParser{}
}

// 用一个简单工厂封装工厂方法
func NewIRuleConfigParserFactory(t string) IRuleConfigParserFactory {
	switch t {
	case "json":
		return jsonRuleConfigParserFactory{}
	case "yaml":
		return yamlRuleConfigParserFactory{}
	}
	fmt.Println("未知类型，返回默认json解析器")
	return jsonRuleConfigParserFactory{}
}

func main() {
	// 创建json解析器工厂，进而创建实例
	jsonParserFactory := NewIRuleConfigParserFactory("json")
	jsonParser := jsonParserFactory.CreateParser()
	fmt.Println(jsonParser)

	// 创建yaml解析器工厂，进而创建实例
	yamlParserFactory := NewIRuleConfigParserFactory("yaml")
	yamlParser := yamlParserFactory.CreateParser()
	fmt.Println(yamlParser)

	// 创建未知类型解析器工厂，进而创建实例
	unknownParserFactory := NewIRuleConfigParserFactory("unknown")
	unknownParser := unknownParserFactory.CreateParser()
	fmt.Println(unknownParser)

	// 测试接口方法(panic占位,未实现)
	//jsonData := []byte(`{"name": "John", "age": 30}`)
	//jsonParser.Parse(jsonData)
}
