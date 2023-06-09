package settings

type Settings_Report struct {
	// 截图开关：True：开启截图，False：关闭截图
	is_screen_shot bool
	// 打印测试日志开关：True：开启打印日志，False：关闭打印日志
	is_write_log bool
	// 打印系统日志开关：True：开启打印日志，False：关闭打印日志
	is_write_system_log bool
	// 日志中是否显示查找元素信息：True：显示，False：不显示
	is_write_find_element_log bool
	// 日志打印级别 1级：业务级别 2级：包含断言（默认） 3级：代码级别
	log_level int
	// 本地日志文件夹路径，也可以使用绝对路径
	log_dir string
	// 本地截图文件夹路径
	screen_shot_dir string
	// 服务器存储图片的路径，仅在向服务器写日志时生效
	server_img_dir string
	// 服务器默认图片路径，日志无截图时使用
	server_default_img_dir string
}
