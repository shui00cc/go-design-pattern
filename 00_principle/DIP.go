package main

import "fmt"

// 抽象层
type MessageSender interface {
	SendMessage(message string, to string) error
}

// 实现层
type EmailSender struct{}

func (e EmailSender) SendMessage(message string, to string) error {
	fmt.Printf("发送邮件给 %s: %s\n", to, message)
	return nil
}

type SMSSender struct{}

func (s SMSSender) SendMessage(message string, to string) error {
	fmt.Printf("发送短信给 %s: %s\n", to, message)
	return nil
}

// 业务层
func main() {
	// 使用 EmailSender 发送邮件
	emailSender := EmailSender{}
	err := send(emailSender, "Hello from EmailSender", "example@example.com")
	if err != nil {
		fmt.Println(err)
	}

	// 使用 SMSSender 发送消息
	smsSender := SMSSender{}
	err = send(smsSender, "Hello from SMSSender", "123456789")
	if err != nil {
		fmt.Println(err)
	}
	// 新增消息发送模式时，只需要在实现层中新增，调用send方法即可
}

// send 充当了业务层和实现层之间的桥梁
func send(sender MessageSender, message string, to string) error {
	return sender.SendMessage(message, to)
}
