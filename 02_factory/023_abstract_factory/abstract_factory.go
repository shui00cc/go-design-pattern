package main

import "fmt"

// ----抽象产品----
//
//	抽象产品1：规则解析器
type IRuleConfigParser interface {
	Parse(data []byte)
}

// 抽象产品2：系统解析器
type ISystemConfigParser interface {
	ParseSystem(data []byte)
}

// ----具体产品----
//
//	具体产品1：json规则解析器
type jsonRuleConfigParser struct {
}

func (J jsonRuleConfigParser) Parse(data []byte) {
	panic("implement me")
}

// 具体产品2：json系统解析器
type jsonSystemConfigParser struct {
}

func (Y jsonSystemConfigParser) ParseSystem(data []byte) {
	panic("implement me")
}

// 具体产品3：yaml规则解析器
type yamlRuleConfigParser struct {
}

func (Y yamlRuleConfigParser) Parse(data []byte) {
	panic("implement me")
}

// 具体产品4：yaml系统解析器
type yamlSystemConfigParser struct {
}

func (Y yamlSystemConfigParser) ParseSystem(data []byte) {
	panic("implement me")
}

// ----抽象父类工厂----
type IConfigParserFactory interface {
	CreateRuleParser() IRuleConfigParser
	CreateSystemParser() ISystemConfigParser
}

// ----具体子类工厂----
//
//	具体子类工厂1：json解析器工厂
type jsonConfigParserFactory struct {
}

func (j jsonConfigParserFactory) CreateRuleParser() IRuleConfigParser {
	fmt.Println("json规则解析器的具体创建逻辑")
	return jsonRuleConfigParser{}
}
func (j jsonConfigParserFactory) CreateSystemParser() ISystemConfigParser {
	fmt.Println("json系统解析器的具体创建逻辑")
	return jsonSystemConfigParser{}
}

// 具体子类工厂2：yaml解析器工厂
type yamlConfigParserFactory struct {
}

func (y yamlConfigParserFactory) CreateRuleParser() IRuleConfigParser {
	fmt.Println("yaml规则解析器的具体创建逻辑")
	return yamlRuleConfigParser{}
}
func (y yamlConfigParserFactory) CreateSystemParser() ISystemConfigParser {
	fmt.Println("yaml系统解析器的具体创建逻辑")
	return yamlSystemConfigParser{}
}

func main() {
	// 创建json解析器工厂，进而创建实例
	jsonParserFactory := jsonConfigParserFactory{}
	jsonRuleParser := jsonParserFactory.CreateRuleParser()
	jsonSystemParser := jsonParserFactory.CreateSystemParser()
	fmt.Println(jsonRuleParser, jsonSystemParser)

	// 创建yaml解析器工厂，进而创建实例
	yamlParserFactory := yamlConfigParserFactory{}
	yamlRuleParser := yamlParserFactory.CreateRuleParser()
	yamlSystemParser := yamlParserFactory.CreateSystemParser()
	fmt.Println(yamlRuleParser, yamlSystemParser)

	// 测试接口方法(panic占位,未实现)
	//jsonData := []byte(`{"name": "John", "age": 30}`)
	//jsonParser.Parse(jsonData)
}
