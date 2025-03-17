package email

type mockEmail struct{}

// Send 发送邮件
func (m *mockEmail) Send() error {
	return nil
}

// SetTo 设置收件人
func (m *mockEmail) SetTo(to ...string) Email {
	return m
}

// SetSubject 设置邮件主题
func (m *mockEmail) SetSubject(subject string) Email {
	return m
}

// SetBody 设置邮件正文
func (m *mockEmail) SetBody(body string, contentType ...string) Email {
	return m
}

// SetAttach 设置附件
func (m *mockEmail) SetAttach(attach ...string) Email {
	return m
}

// SetCc 设置抄送人
func (m *mockEmail) SetCc(cc ...string) Email {
	return m
}

// NewMockEmail 创建邮件模拟
func NewMockEmail() Email {
	return &mockEmail{}
}
