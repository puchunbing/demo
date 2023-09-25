package controllers

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"time"
)

type ContextTestControllers struct {
}

var (
	ContextTestControllersInstance *ContextTestControllers
)

func init() {
	ContextTestControllersInstance = &ContextTestControllers{}
}

func (this *ContextTestControllers) TimeOutTest(c *gin.Context) {
	contont, err := geturlContent("https://github.com/", 5*time.Second)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(contont)
}

func geturlContent(url string, timeout time.Duration) (string, error) {
	type Resp struct {
		content string
		err     error
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	ch := make(chan Resp, 1)
	go func() {
		res, err := http.Get(url)
		if err != nil {
			ch <- Resp{"", err}
			return
		}
		defer res.Body.Close()

		boby, err := io.ReadAll(res.Body)
		if err != nil {
			ch <- Resp{"", err}
			return
		}
		ch <- Resp{string(boby), nil}

	}()
	select {
	case resp := <-ch:
		return resp.content, resp.err
	case <-ctx.Done(): //超时
		return "", errors.New("超时了")

	}
}

func (this *ContextTestControllers) TimeAfterDemo(c *gin.Context) {
	ch := make(chan string)
	go func() {
		time.Sleep(5 * time.Second)
		ch <- "124"
	}()
	select {
	case resp := <-ch:
		fmt.Println(resp)
		return
	case <-time.After(3 * time.Second): //超时
		fmt.Println("超时了")
		return
	}
}
