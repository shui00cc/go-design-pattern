package main

import "fmt"

// 定义一个配置结构体
type Config struct {
	Timeout    int
	MaxRetries int
	Verbose    bool
}

// 定义一个函数类型，接受一个配置结构体作为参数
type Option func(*Config)

// 函数选项1：设置超时时间
func WithTimeout(timeout int) Option {
	return func(c *Config) {
		c.Timeout = timeout
	}
}

// 函数选项2：设置最大重试次数
func WithMaxRetries(maxRetries int) Option {
	return func(c *Config) {
		c.MaxRetries = maxRetries
	}
}

// 函数选项3：设置是否启用详细日志
func WithVerbose(verbose bool) Option {
	return func(c *Config) {
		c.Verbose = verbose
	}
}

// 主要函数，接受一个可变数量的函数选项
func Connect(url string, options ...Option) {
	// 默认配置
	config := Config{
		Timeout:    10,
		MaxRetries: 3,
		Verbose:    false,
	}

	// 应用所有传递进来的选项
	for _, option := range options {
		option(&config)
	}

	// 打印配置信息（演示用途）
	fmt.Printf("Connecting to %s with config: %+v\n", url, config)
}

func main() {
	// 使用函数选项模式调用Connect函数
	Connect("example.com",
		WithTimeout(5),
		WithMaxRetries(2),
		//WithVerbose(true),
	)
}
