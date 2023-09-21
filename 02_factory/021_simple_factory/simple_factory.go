package factory

type IRuleConfigParser interface {
	Parse(data []byte)
}

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

// 根据参数对应关系 返回一个对应的 IRuleConfigParser 接口
func NewIRuleConfigParser(t string) IRuleConfigParser {
	switch t {
	case "json":
		return jsonRuleConfigParser{}
	case "yaml":
		return yamlRuleConfigParser{}
	}
	return nil
}
