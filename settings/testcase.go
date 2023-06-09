package settings

type Settings_TestCase struct {
	// 运行模式(Remote,Browser)
	run_type string
	// 测试方案配置文件夹路径(支持绝对路径)
	xml_dir string
	// 测试用例配置文件夹路径(必须相对于项目根目录)
	testcase_dir string
	// 待上传文件路径
	files_dir string
	// start表示一个案例开始前启用（通过getObj），其他值表示用例执行时通过getDriver创建
	create_type string
}
