package main

import (
	"fmt"
	"log"
)

// 请求搜索引擎，需要传递若干参数：
// query, count, page, user_agent, timeout, src, ip, correct, refer
// 如果使用构造函数或函数调用时传入，大量参数不利于维护，且增加新参数，更改较大。
// 这种情况适用于 建造者模式

func main() {
	fmt.Println("hello")
	var eng Engine
	// 正常的建造者模式
	eb := EngineBuilder{}
	eng = eb.SetQuery("nba").SetCount(10).SetPage(1).Build()
	eng.Search()

	// 未设置query的建造
	eb2 := EngineBuilder{}
	eng = eb2.SetCount(10).SetPage(1).Build()
	eng.Search()

}

// Engine 真正的引擎类
type Engine struct {
	Query     string
	Count     int64
	Page      int64
	UserAgent string
	Timeout   int64
	Src       string
	IP        string
	Correct   bool
	Refer     string
}

// Search 发起引擎请求
func (eng Engine) Search() {
	fmt.Printf("Start Search ...\nQuery: %s\nCount: %d\nPage: %d\nSearch Success ...\n", eng.Query, eng.Count, eng.Page)
}

// EngineBuilder 引擎建造者
type EngineBuilder struct {
	Query     string
	Count     int64
	Page      int64
	UserAgent string
	Timeout   int64
	Src       string
	IP        string
	Correct   bool
	Refer     string
}

// Build 建造引擎的方法
func (eb EngineBuilder) Build() Engine {
	// 判断参数是否合法
	if eb.Query == "" {
		log.Fatal("Query empty.")
	}

	return Engine{
		Query: eb.Query,
		Count: eb.Count,
		Page:  eb.Page,
	}
}

// SetQuery 设置引擎Query
func (eb EngineBuilder) SetQuery(query string) EngineBuilder {
	eb.Query = query
	return eb
}

// SetCount 设置请求数量
func (eb EngineBuilder) SetCount(n int64) EngineBuilder {
	eb.Count = n
	return eb
}

// SetPage 设置请求页数
func (eb EngineBuilder) SetPage(n int64) EngineBuilder {
	eb.Page = n
	return eb
}

// ...省略其它方法
