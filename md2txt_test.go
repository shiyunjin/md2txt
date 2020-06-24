package md2txt_test

import (
	"testing"

	"github.com/shiyunjin/md2txt"
)

func TestToTxt(t *testing.T) {
	txt := md2txt.ToTXT("v2.1.2版本的百度，应用包，360渠道包\nV 2.1.1 更新日志：\n##### 功能优化\n- 手机通知栏点击可进入加速页\n- 本地添加游戏区服列表样式调整\n- 可通过延迟检测提示直接开启双通道\n- 改善了加速器在5.0-6.0系统的界面适配及运行稳定性\n##### Bug修复\n- 内嵌网页或跳转失败\n 安卓5.0.2系统无法登录\n(点击查看)[https://www.taptap.com/app/180786/review?order=update#review-list]\n(点击查看)[https://www.taptap.com/app/180786/review?order=update#review-list](点击查看)[https://www.taptap.com/app/180786/review?order=update#review-list][点击查看|https://www.taptap.com/app/180786/review?order=update#review-list](点击查看)[https://www.taptap.com/app/180786/review?order=update#review-list][点击查看|https://www.taptap.com/app/180786/review?order=update#review-list]")
	t.Log(txt)
}
