package settings

type Settings_Server struct {
	// 发送日志到服务器的开关：True：发送日志，False 不发送日志
	is_request bool
	// 日志发送URL
	is_request_url string
}
