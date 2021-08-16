package optimization

import (
	"errors"
	"log"
	"testing"
)

// Test__retry_Run 始终失败
func Test__retry_Run(t *testing.T) {
	count := 0
	if err := NewRetry().
		Subject(func() error {
			log.Println("Subject body : helloWorld")

			if count < 2 {
				return errors.New("擦")
			}
			count++
			return nil
		}).Sleep(2).
		Count(2).
		Run(); err != nil {
		log.Println(err.Error())
		return
	}

	log.Println("执行成功")
}

// Test__retry_OK 重试后成功
func Test__retry_OK(t *testing.T) {
	count := 0
	if err := NewRetry().
		Subject(func() error {
			log.Println("Subject body : helloWorld")
			count++
			if count < 2 {
				return errors.New("擦")
			}

			return nil
		}).Sleep(2).
		Count(2).
		Run(); err != nil {
		log.Println(err.Error())
		return
	}

	log.Println("执行成功")
}

// Test__retry_OK 重试后成功
func Test__retry_Normal(t *testing.T) {
	count := 0
	if err := NewRetry().
		Subject(func() error {
			log.Println("Subject body : helloWorld")
			count++
			println(count)
			return nil
		}).Sleep(2).
		Count(2).
		Run(); err != nil {
		log.Println(err.Error())
		return
	}

	log.Println("执行成功")
}
