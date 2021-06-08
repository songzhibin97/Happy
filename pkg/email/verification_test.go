package email

import "testing"

func TestEmailAuth_Send(t *testing.T) {
	ea := EmailAuth{
		Host:     QQHost,
		Server:   QQServerAddr,
		Auth:     "718428482@qq.com",
		Password: "",
	}
	ea.Send(ea.CreateTemp("718428482@qq.com", "Seer", "帅彬"))

}
