package sat

import (
	"testing"
)

var (
	dicter = DefaultDict()
)

func TestConvertPlain(t *testing.T) {
	cases := map[string]string{
		"繁体中文": "繁體中文",
		"还带来了音乐等方面的重大更新": "還帶來了音樂等方面的重大更新",
		"如果你是 iOS 14.5.1 用户，或者 Apple Music 的忠实粉丝，那么根本找不到理由去拒绝 iOS 14.6。": "如果你是 iOS 14.5.1 用戶，或者 Apple Music 的忠實粉絲，那麼根本找不到理由去拒絕 iOS 14.6。",
		"同花顺":         "同花順",
		"搜索":          "搜索",
		"返回":          "返回",
		"阿里巴巴":        "阿里巴巴",
		"长桥资讯":        "長橋資訊",
		"零一二三四五六七八九十": "零一二三四五六七八九十",
		"如果":          "如果",
		"我爱你":         "我愛你",
		"温暖的太阳":       "溫暖的太陽",
		"请输入遇到的问题":    "請輸入遇到的問題",
		"你或许会发现，如果你没有设定编辑器，当系统预设编辑器被执行时你很有可能会不知所措。 而在 Windows 系统下，使用系统预设编辑器编辑讯息的过程中，有可能因不熟悉而误用，进而导致某些 Git 操作过早结束": "你或許會發現，如果你沒有設定編輯器，當系統預設編輯器被執行時你很有可能會不知所措。 而在 Windows 系統下，使用系統預設編輯器編輯訊息的過程中，有可能因不熟悉而誤用，進而導致某些 Git 操作過早結束",
		"优趣汇或成为继宝尊电商之后第二家品牌电商港股。尿不湿品牌尤妮佳和化妆品品牌资生堂贡献超七成营收，营收连续增长，近三年经调整净利润均超 1 亿。":                                 "優趣匯或成為繼寶尊電商之后第二家品牌電商港股。尿不濕品牌尤妮佳和化妝品品牌資生堂貢獻超七成營收，營收連續增長，近三年經調整淨利潤均超 1 億。",
	}
	for sc, expected := range cases {
		actually := dicter.ReadReverse(sc)
		if expected != actually {
			t.Errorf("\nexpected: %s\nactually: %s", expected, actually)
		}
	}

}

func TestConvertHTML(t *testing.T) {
	source := `<p id="KN1928">如果你最近看到有人打<strong>王者荣耀</strong>时情绪波动较大，不一定是因为 <em>ta</em> 的队友太坑或对手太强，而或许只因 ta 升级了 iOS 14.5.1。</p>
	<p><img src="https://img.example.com/not/exist/image.jpg" alt="无效的图片"></p>
	<p>发热、掉帧、卡顿，iOS 14.5.1 的性能 bug，成为<a href="https://example.com" target="_blank" rel="nofollow">近期</a> iPhone 用户吐槽最多的话题。无论你是哪款 iPhone，升级 iOS 14.5.1 之后，都有可能遇到性能缩水的问题。</p>`

	expect := `<p id="KN1928">如果你最近看到有人打<strong>王者榮耀</strong>時情緒波動較大，不一定是因為 <em>ta</em> 的隊友太坑或對手太強，而或許只因 ta 升級了 iOS 14.5.1。</p>
	<p><img src="https://img.example.com/not/exist/image.jpg" alt="無效的圖片"></p>
	<p>發熱、掉幀、卡頓，iOS 14.5.1 的性能 bug，成為<a href="https://example.com" target="_blank" rel="nofollow">近期</a> iPhone 用戶吐槽最多的話題。無論你是哪款 iPhone，升級 iOS 14.5.1 之后，都有可能遇到性能縮水的問題。</p>`

	out := dicter.ReadReverse(source)
	if expect != out {
		t.Error(out)
	}
}
