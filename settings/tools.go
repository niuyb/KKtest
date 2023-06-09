package settings

import "strings"

func GetScriptDir() string {
	return "1"
}


func InitSrcPath(){
	
}

func GetSrcDir(srcPath, srcDefaultPath string) (SRC_dir string) {

	defer func() {
		if err := recover(); err != nil {
			if len(srcPath) != 0 {
				// 源码模块路径可改，需根据SRC实际位置修改
				SRC_dir = strings.Replace(srcPath, "\\", "/", -1)
			} else {
				SRC_dir = GetScriptDir()
			}
		}
	}()
	SRC_dir = srcDefaultPath
	


	return SRC_dir

}
