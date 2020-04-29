package main

import (
	"fmt"
	"log"
	"strings"
)

/*
*	根据配置文件的后缀（json、xml、yaml、properties），选择不同的解析器（JsonRuleConfigParser、XmlRuleConfigParser……），将存储在文件中的配置解析成内存对象 RuleConfig。
 */

// RuleConfig 解析出的规则
type RuleConfig struct{}

// RuleConfigParser 解析器接口
type RuleConfigParser interface {
	Parse(string) RuleConfig
}

// JSONRuleConfigParser json 解析器
type JSONRuleConfigParser struct{}

// Parse 解析文件
func (rcp JSONRuleConfigParser) Parse(file string) RuleConfig {
	fmt.Println("parse json file.")
	return RuleConfig{}
}

// XMLRuleConfigParser xml 解析器
type XMLRuleConfigParser struct{}

// Parse 解析文件
func (rcp XMLRuleConfigParser) Parse(file string) RuleConfig {
	fmt.Println("xml json file.")
	return RuleConfig{}
}

// YAMLRuleConfigParser yaml 解析器
type YAMLRuleConfigParser struct{}

// Parse 解析文件
func (rcp YAMLRuleConfigParser) Parse(file string) RuleConfig {
	fmt.Println("yaml json file.")
	return RuleConfig{}
}

func main() {
	fmt.Println("hello")

	// 一般做法; 有大量的 if..else

	var filename string = "test.json"
	var rcp RuleConfigParser

	log.Println("一般做法, if..else..")
	if FileExtension(filename, "json") {
		rcp = JSONRuleConfigParser{}
	} else if FileExtension(filename, "xml") {
		rcp = XMLRuleConfigParser{}
	} else if FileExtension(filename, "yaml") {
		rcp = YAMLRuleConfigParser{}
	}
	_ = rcp.Parse(filename)

	// 简单工厂模式
	// 对于简单工厂，如果我们要添加新的 parser，势必要改动到 RuleConfigParserFactory 的代码，
	// 那这是不是违反开闭原则呢？实际上，如果不是需要频繁地添加新的 parser，只是偶尔修改一下
	// RuleConfigParserFactory 代码，稍微不符合开闭原则，也是完全可以接受的。
	log.Println("简单工厂模式：")
	rcpf := RuleConfigParserFactory{}
	rcp2 := rcpf.CreateRCP(filename)
	_ = rcp2.Parse(filename)

	// 工厂方法
	// 如果完全要避免上述简单工厂的弊端，需要使用工厂方法；
	// 实际就是对上述简单工厂再使用一个工厂，成为工厂的工厂，这样就会减少对 RuleConfigParserFactory 的改动.
	log.Println("工厂方法模式：")
	rcp3 := NewRuleConfigParserFactory(filename)
	_ = rcp3.Parse(filename)

	// 抽象工厂
	// 涉及到多个产品组合的情况，抽象出工厂，工厂实现各类组合形式的构造。
}

// FileExtension 判断文件是否某个类型
func FileExtension(filename, extension string) bool {
	return strings.HasSuffix(filename, "."+extension)
}

// ====简单工厂====

// RuleConfigParserFactory RuleConfigParser 简单工厂模式
type RuleConfigParserFactory struct {
}

// CreateRCP 创建 RuleConfigParser
func (rcpf RuleConfigParserFactory) CreateRCP(filename string) RuleConfigParser {
	var rcp RuleConfigParser
	if FileExtension(filename, "json") {
		rcp = JSONRuleConfigParser{}
	} else if FileExtension(filename, "xml") {
		rcp = XMLRuleConfigParser{}
	} else if FileExtension(filename, "yaml") {
		rcp = YAMLRuleConfigParser{}
	}
	return rcp
}

// ====工厂方法(工厂的工厂)====

// IRuleConfigParserFactory 工厂接口
type IRuleConfigParserFactory interface {
	CreateFactory() RuleConfigParser
}

// JSONRuleConfigParserFactory JSON工厂实现
type JSONRuleConfigParserFactory struct{}

// CreateFactory JSON工厂实现
func (jrcpf JSONRuleConfigParserFactory) CreateFactory() RuleConfigParser {
	return JSONRuleConfigParser{}
}

// XMLRuleConfigParserFactory XML工厂实现
type XMLRuleConfigParserFactory struct{}

// CreateFactory XML工厂实现
func (xrcpf XMLRuleConfigParserFactory) CreateFactory() RuleConfigParser {
	return XMLRuleConfigParser{}
}

// YAMLRuleConfigParserFactory YAML工厂实现
type YAMLRuleConfigParserFactory struct{}

// CreateFactory YAML工厂实现
func (yrcpf YAMLRuleConfigParserFactory) CreateFactory() RuleConfigParser {
	return YAMLRuleConfigParser{}
}

// NewRuleConfigParserFactory 工厂的工厂，创建对应类型的工厂，如果新增类型，增加工厂实现，这里直接调用对应的工厂来创建
// 这里虽然也涉及 if..else ，在有新增类型时也涉及更改，不过工厂的逻辑都在单独的工厂类里实现了，这里的工作很简单，满足开闭原则。
// 当对象的创建逻辑比较复杂，不只是简单的 new 一下就可以，而是要组合其他类对象，做各种初始化操作的时候，我们推荐使用工厂方法模式，
// 将复杂的创建逻辑拆分到多个工厂类中，让每个工厂类都不至于过于复杂。
// 而使用简单工厂模式，将所有的创建逻辑都放到一个工厂类中，会导致这个工厂类变得很复杂。
func NewRuleConfigParserFactory(filename string) RuleConfigParser {
	var rcpf IRuleConfigParserFactory
	if FileExtension(filename, "json") {
		rcpf = JSONRuleConfigParserFactory{}
	} else if FileExtension(filename, "xml") {
		rcpf = XMLRuleConfigParserFactory{}
	} else if FileExtension(filename, "yaml") {
		rcpf = YAMLRuleConfigParserFactory{}
	}
	return rcpf.CreateFactory()
}
