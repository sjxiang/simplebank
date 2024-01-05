package mail

import (
	"testing"

	"github.com/stretchr/testify/require"
	
	"github.com/sjxiang/simplebank/util"
)

func TestSendEmailWithGmail(t *testing.T) {
	// go test 设置 -short tag，会跳过该项测试 
	if testing.Short() {
		t.Skip()
	}

	config, err := util.LoadConfig("..")
	require.NoError(t, err)

	sender := NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)

	subject := "A test email"
	content := `
	<h1>Hello world</h1>
	<p>This is a test message from <a href="https://github.com/sjxiang">Sjxiang Github</a></p>
	`
	to := []string{"1535484943@qq.com"}
	attachFiles := []string{"../README.md"}

	err = sender.SendEmail(subject, content, to, nil, nil, attachFiles)
	require.NoError(t, err)
}

func TestSendEmailWithQQ(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	config, err := util.LoadConfig("..")
	require.NoError(t, err)

	qq := NewQQMailSender(config.QQEmailSenderName, config.QQEmailSenderAddress, config.QQEmailSenderPassword)

	subject := "A test email"
	content := `
	<h1>Hello world</h1>
	<p>This is a test message from <a href="https://github.com/sjxiang">Sjxiang Github</a></p>
	`
	to := []string{"sjxiang2023@gmail.com"}
	attachFiles := []string{"../README.md"}

	err = qq.SendEmail(subject, content, to, nil, nil, attachFiles)
	require.NoError(t, err)
}