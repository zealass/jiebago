package analyse

import (
	"github.com/wangbin/jiebago"
	"testing"
)

var (
	test_contents = []string{
		"这是一个伸手不见五指的黑夜。我叫孙悟空，我爱北京，我爱Python和C++。",
		"我不喜欢日本和服。",
		"雷猴回归人间。",
		"工信处女干事每月经过下属科室都要亲口交代24口交换机等技术性器件的安装工作",
		"我需要廉租房",
		"永和服装饰品有限公司",
		"我爱北京天安门",
		"abc",
		"隐马尔可夫",
		"雷猴是个好网站",
		"“Microsoft”一词由“MICROcomputer（微型计算机）”和“SOFTware（软件）”两部分组成",
		"草泥马和欺实马是今年的流行词汇",
		"伊藤洋华堂总府店",
		"中国科学院计算技术研究所",
		"罗密欧与朱丽叶",
		"我购买了道具和服装",
		"PS: 我觉得开源有一个好处，就是能够敦促自己不断改进，避免敞帚自珍",
		"湖北省石首市",
		"湖北省十堰市",
		"总经理完成了这件事情",
		"电脑修好了",
		"做好了这件事情就一了百了了",
		"人们审美的观点是不同的",
		"我们买了一个美的空调",
		"线程初始化时我们要注意",
		"一个分子是由好多原子组织成的",
		"祝你马到功成",
		"他掉进了无底洞里",
		"中国的首都是北京",
		"孙君意",
		"外交部发言人马朝旭",
		"领导人会议和第四届东亚峰会",
		"在过去的这五年",
		"还需要很长的路要走",
		"60周年首都阅兵",
		"你好人们审美的观点是不同的",
		"买水果然后来世博园",
		"买水果然后去世博园",
		"但是后来我才知道你是对的",
		"存在即合理",
		"的的的的的在的的的的就以和和和",
		"I love你，不以为耻，反以为rong",
		"因",
		"",
		"hello你好人们审美的观点是不同的",
		"很好但主要是基于网页形式",
		"hello你好人们审美的观点是不同的",
		"为什么我不能拥有想要的生活",
		"后来我才",
		"此次来中国是为了",
		"使用了它就可以解决一些问题",
		",使用了它就可以解决一些问题",
		"其实使用了它就可以解决一些问题",
		"好人使用了它就可以解决一些问题",
		"是因为和国家",
		"老年搜索还支持",
		"干脆就把那部蒙人的闲法给废了拉倒！RT @laoshipukong : 27日，全国人大常委会第三次审议侵权责任法草案，删除了有关医疗损害责任“举证倒置”的规定。在医患纠纷中本已处于弱势地位的消费者由此将陷入万劫不复的境地。 ",
		"大",
		"",
		"他说的确实在理",
		"长春市长春节讲话",
		"结婚的和尚未结婚的",
		"结合成分子时",
		"旅游和服务是最好的",
		"这件事情的确是我的错",
		"供大家参考指正",
		"哈尔滨政府公布塌桥原因",
		"我在机场入口处",
		"邢永臣摄影报道",
		"BP神经网络如何训练才能在分类时增加区分度？",
		"南京市长江大桥",
		"应一些使用者的建议，也为了便于利用NiuTrans用于SMT研究",
		"长春市长春药店",
		"邓颖超生前最喜欢的衣服",
		"胡锦涛是热爱世界和平的政治局常委",
		"程序员祝海林和朱会震是在孙健的左面和右面, 范凯在最右面.再往左是李松洪",
		"一次性交多少钱",
		"两块五一套，三块八一斤，四块七一本，五块六一条",
		"小和尚留了一个像大和尚一样的和尚头",
		"我是中华人民共和国公民;我爸爸是共和党党员; 地铁和平门站",
		"张晓梅去人民医院做了个B超然后去买了件T恤",
		"AT&T是一件不错的公司，给你发offer了吗？",
		"C++和c#是什么关系？11+122=133，是吗？PI=3.14159",
		"你认识那个和主席握手的的哥吗？他开一辆黑色的士。",
		"枪杆子中出政权"}
	Tags = [][]string{
		[]string{"Python", "C++", "\u4f38\u624b\u4e0d\u89c1\u4e94\u6307", "\u5b59\u609f\u7a7a", "\u9ed1\u591c", "\u5317\u4eac", "\u8fd9\u662f", "\u4e00\u4e2a"},
		[]string{"\u548c\u670d", "\u559c\u6b22", "\u65e5\u672c"},
		[]string{"\u96f7\u7334", "\u4eba\u95f4", "\u56de\u5f52"},
		[]string{"\u5de5\u4fe1\u5904", "\u5973\u5e72\u4e8b", "24", "\u4ea4\u6362\u673a", "\u79d1\u5ba4", "\u4eb2\u53e3", "\u5668\u4ef6", "\u6280\u672f\u6027", "\u4e0b\u5c5e", "\u4ea4\u4ee3", "\u6bcf\u6708", "\u5b89\u88c5", "\u7ecf\u8fc7", "\u5de5\u4f5c"},
		[]string{"\u5ec9\u79df\u623f", "\u9700\u8981"},
		[]string{"\u9970\u54c1", "\u6c38\u548c", "\u670d\u88c5", "\u6709\u9650\u516c\u53f8"},
		[]string{"\u5929\u5b89\u95e8", "\u5317\u4eac"},
		[]string{"abc"},
		[]string{"\u9a6c\u5c14\u53ef\u592b"},
		[]string{"\u96f7\u7334", "\u7f51\u7ad9"},
		[]string{"SOFTware", "Microsoft", "MICROcomputer", "\u5fae\u578b", "\u4e00\u8bcd", "\u8f6f\u4ef6", "\u8ba1\u7b97\u673a", "\u7ec4\u6210", "\u90e8\u5206"},
		[]string{"\u8349\u6ce5\u9a6c", "\u6b3a\u5b9e", "\u8bcd\u6c47", "\u6d41\u884c", "\u4eca\u5e74"},
		[]string{"\u6d0b\u534e\u5802", "\u603b\u5e9c", "\u4f0a\u85e4"},
		[]string{"\u4e2d\u56fd\u79d1\u5b66\u9662\u8ba1\u7b97\u6280\u672f\u7814\u7a76\u6240"},
		[]string{"\u6731\u4e3d\u53f6", "\u7f57\u5bc6\u6b27"},
		[]string{"\u9053\u5177", "\u670d\u88c5", "\u8d2d\u4e70"},
		[]string{"\u81ea\u73cd", "\u655e\u5e1a", "PS", "\u5f00\u6e90", "\u4e0d\u65ad\u6539\u8fdb", "\u6566\u4fc3", "\u597d\u5904", "\u907f\u514d", "\u80fd\u591f", "\u89c9\u5f97", "\u5c31\u662f", "\u81ea\u5df1", "\u4e00\u4e2a"},
		[]string{"\u77f3\u9996\u5e02", "\u6e56\u5317\u7701"},
		[]string{"\u5341\u5830\u5e02", "\u6e56\u5317\u7701"},
		[]string{"\u603b\u7ecf\u7406", "\u8fd9\u4ef6", "\u5b8c\u6210", "\u4e8b\u60c5"},
		[]string{"\u4fee\u597d", "\u7535\u8111"},
		[]string{"\u4e00\u4e86\u767e\u4e86", "\u505a\u597d", "\u8fd9\u4ef6", "\u4e8b\u60c5"},
		[]string{"\u5ba1\u7f8e", "\u89c2\u70b9", "\u4eba\u4eec", "\u4e0d\u540c"},
		[]string{"\u7f8e\u7684", "\u7a7a\u8c03", "\u6211\u4eec", "\u4e00\u4e2a"},
		[]string{"\u7ebf\u7a0b", "\u521d\u59cb\u5316", "\u6ce8\u610f", "\u6211\u4eec"},
		[]string{"\u597d\u591a", "\u539f\u5b50", "\u5206\u5b50", "\u7ec4\u7ec7", "\u4e00\u4e2a"},
		[]string{"\u9a6c\u5230\u529f\u6210"},
		[]string{"\u65e0\u5e95\u6d1e"},
		[]string{"\u9996\u90fd", "\u5317\u4eac", "\u4e2d\u56fd"},
		[]string{"\u5b59\u541b\u610f"},
		[]string{"\u9a6c\u671d\u65ed", "\u5916\u4ea4\u90e8", "\u53d1\u8a00\u4eba"},
		[]string{"\u7b2c\u56db\u5c4a", "\u4e1c\u4e9a", "\u5cf0\u4f1a", "\u9886\u5bfc\u4eba", "\u4f1a\u8bae"},
		[]string{"\u4e94\u5e74", "\u8fc7\u53bb"},
		[]string{"\u5f88\u957f", "\u9700\u8981"},
		[]string{"60", "\u9605\u5175", "\u5468\u5e74", "\u9996\u90fd"},
		[]string{"\u5ba1\u7f8e", "\u4f60\u597d", "\u89c2\u70b9", "\u4eba\u4eec", "\u4e0d\u540c"},
		[]string{"\u4e16\u535a\u56ed", "\u6c34\u679c", "\u7136\u540e"},
		[]string{"\u4e16\u535a\u56ed", "\u6c34\u679c", "\u7136\u540e"},
		[]string{"\u540e\u6765", "\u4f46\u662f", "\u77e5\u9053"},
		[]string{"\u5408\u7406", "\u5b58\u5728"},
		[]string{},
		[]string{"rong", "love", "\u4e0d\u4ee5\u4e3a\u803b", "\u4ee5\u4e3a"},
		[]string{},
		[]string{},
		[]string{"hello", "\u5ba1\u7f8e", "\u4f60\u597d", "\u89c2\u70b9", "\u4eba\u4eec", "\u4e0d\u540c"},
		[]string{"\u7f51\u9875", "\u57fa\u4e8e", "\u5f62\u5f0f", "\u4e3b\u8981"},
		[]string{"hello", "\u5ba1\u7f8e", "\u4f60\u597d", "\u89c2\u70b9", "\u4eba\u4eec", "\u4e0d\u540c"},
		[]string{"\u60f3\u8981", "\u62e5\u6709", "\u4e3a\u4ec0\u4e48", "\u751f\u6d3b", "\u4e0d\u80fd"},
		[]string{"\u540e\u6765"},
		[]string{"\u6b64\u6b21", "\u4e3a\u4e86", "\u4e2d\u56fd"},
		[]string{"\u89e3\u51b3", "\u4f7f\u7528", "\u4e00\u4e9b", "\u95ee\u9898", "\u53ef\u4ee5"},
		[]string{"\u89e3\u51b3", "\u4f7f\u7528", "\u4e00\u4e9b", "\u95ee\u9898", "\u53ef\u4ee5"},
		[]string{"\u89e3\u51b3", "\u5176\u5b9e", "\u4f7f\u7528", "\u4e00\u4e9b", "\u95ee\u9898", "\u53ef\u4ee5"},
		[]string{"\u597d\u4eba", "\u89e3\u51b3", "\u4f7f\u7528", "\u4e00\u4e9b", "\u95ee\u9898", "\u53ef\u4ee5"},
		[]string{"\u662f\u56e0\u4e3a", "\u56fd\u5bb6"},
		[]string{"\u8001\u5e74", "\u641c\u7d22", "\u652f\u6301"},
		[]string{"\u95f2\u6cd5", "\u4e2d\u672c", "laoshipukong", "RT", "27", "\u8d23\u4efb\u6cd5", "\u8499\u4eba", "\u4e07\u52ab\u4e0d\u590d", "\u4e3e\u8bc1", "\u5012\u7f6e", "\u533b\u60a3", "\u90a3\u90e8", "\u62c9\u5012", "\u4fb5\u6743", "\u5168\u56fd\u4eba\u5927\u5e38\u59d4\u4f1a", "\u8349\u6848", "\u5883\u5730", "\u7ea0\u7eb7", "\u5220\u9664", "\u5f31\u52bf"},
		[]string{},
		[]string{},
		[]string{"\u5728\u7406", "\u786e\u5b9e"},
		[]string{"\u957f\u6625", "\u6625\u8282", "\u8bb2\u8bdd", "\u5e02\u957f"},
		[]string{"\u7ed3\u5a5a", "\u5c1a\u672a"},
		[]string{"\u5206\u5b50", "\u7ed3\u5408"},
		[]string{"\u65c5\u6e38", "\u6700\u597d", "\u670d\u52a1"},
		[]string{"\u7684\u786e", "\u8fd9\u4ef6", "\u4e8b\u60c5"},
		[]string{"\u6307\u6b63", "\u53c2\u8003", "\u5927\u5bb6"},
		[]string{"\u584c\u6865", "\u54c8\u5c14\u6ee8", "\u516c\u5e03", "\u539f\u56e0", "\u653f\u5e9c"},
		[]string{"\u5165\u53e3\u5904", "\u673a\u573a"},
		[]string{"\u90a2\u6c38\u81e3", "\u6444\u5f71", "\u62a5\u9053"},
		[]string{"\u533a\u5206\u5ea6", "BP", "\u795e\u7ecf\u7f51\u7edc", "\u8bad\u7ec3", "\u5206\u7c7b", "\u624d\u80fd", "\u5982\u4f55", "\u589e\u52a0"},
		[]string{"\u957f\u6c5f\u5927\u6865", "\u5357\u4eac\u5e02"},
		[]string{"SMT", "NiuTrans", "\u4f7f\u7528\u8005", "\u4fbf\u4e8e", "\u7528\u4e8e", "\u5efa\u8bae", "\u5229\u7528", "\u4e3a\u4e86", "\u7814\u7a76", "\u4e00\u4e9b"},
		[]string{"\u957f\u6625\u5e02", "\u836f\u5e97", "\u957f\u6625"},
		[]string{"\u9093\u9896\u8d85", "\u751f\u524d", "\u8863\u670d", "\u559c\u6b22"},
		[]string{"\u653f\u6cbb\u5c40", "\u70ed\u7231", "\u5e38\u59d4", "\u80e1\u9526\u6d9b", "\u548c\u5e73", "\u4e16\u754c"},
		[]string{"\u53f3\u9762", "\u5b59\u5065", "\u8303\u51ef", "\u674e\u677e\u6d2a", "\u6731\u4f1a\u9707", "\u6d77\u6797", "\u5de6\u9762", "\u7a0b\u5e8f\u5458", "\u518d\u5f80"},
		[]string{"\u4e00\u6b21\u6027", "\u591a\u5c11"},
		[]string{"\u56db\u5757", "\u4e94\u5757", "\u4e09\u5757", "\u4e00\u65a4", "\u4e24\u5757", "\u4e00\u672c", "\u4e00\u5957", "\u4e00\u6761"},
		[]string{"\u548c\u5c1a", "\u548c\u5c1a\u5934", "\u4e00\u6837", "\u4e00\u4e2a"},
		[]string{"\u548c\u5e73\u95e8", "\u5171\u548c\u515a", "\u5730\u94c1", "\u515a\u5458", "\u516c\u6c11", "\u7238\u7238", "\u4e2d\u534e\u4eba\u6c11\u5171\u548c\u56fd"},
		[]string{"\u5f20\u6653\u6885", "T\u6064", "B\u8d85", "\u533b\u9662", "\u4eba\u6c11", "\u7136\u540e"},
		[]string{"offer", "AT&T", "\u4e0d\u9519", "\u4e00\u4ef6", "\u516c\u53f8"},
		[]string{"c#", "PI", "C++", "3.14159", "133", "122", "11", "\u5173\u7cfb", "\u4ec0\u4e48"},
		[]string{"\u7684\u58eb", "\u7684\u54e5", "\u4ed6\u5f00", "\u63e1\u624b", "\u4e00\u8f86", "\u9ed1\u8272", "\u4e3b\u5e2d", "\u8ba4\u8bc6", "\u90a3\u4e2a"},
		[]string{"\u67aa\u6746\u5b50", "\u653f\u6743"},
	}
)

func TestExtractTags(t *testing.T) {
	jiebago.SetDictionary("../dict.txt")
	SetIdf("idf.txt")
	for index, sentence := range test_contents {
		result := ExtractTags(sentence, 20)
		if len(result) != len(Tags[index]) {
			t.Errorf("%s = %v", sentence, result)
		}
		for i, tag := range result {
			if tag != Tags[index][i] {
				t.Error(tag)
			}
		}
	}
}
