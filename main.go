//go:generate goversioninfo -icon=icon.ico -manifest=resource/goversioninfo.exe.manifest

package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)

var logStr = ""
var prefix = "喜欢甜文的宝子看过来，很高很高评分的甜文来啦，"
var suffix = "都看到这里了，给个关注呗，宝贝"

const (
	_ua = "Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36"
)

func main() {
	diff, err := getDateDiff()
	if err != nil {
		writeLog(err.Error())
		writeStr(logStr)
		pause()
		return
	}
	wd, err := os.Getwd()
	if err != nil {
		return
	}
	if diff > 24 &&
		!strings.Contains(wd, "benfei") &&
		!strings.Contains(wd, "MR") {
		// if diff > 24 {
		fmt.Println("使用余额不足，如需继续使用，请联系作者：buffer5")
		writeLog("使用余额不足，如需继续使用，请联系作者：buffer5")
		writeStr(logStr)
		pause()
		return
	}

	writeLog("欢迎使用梨涡文档处理工具V1.0，如有疑问请联系作者：buffer5\n")
	fmt.Println("欢迎使用梨涡文档处理工具V1.0，如有疑问请联系作者：buffer5")
	fmt.Println("\n提取知乎文案：输入1\n已有文案处理：输入2\n提起主页视频：输入3\n输入完成后回车键确认：")
	b := make([]byte, 999)
	os.Stdin.Read(b)
	s := strings.Split(string(b), "\n")[0]
	s = strings.Split(s, "\r")[0]
	if s == "1" { // 知乎链接不可用，暂时无法提取11
		err := getTextByUrl()
		if err != nil {
			writeLog(err.Error())
			writeStr(logStr)
			pause()
			return
		}
		s = "2"
	}
	if s == "2" {
		err := textProcess()
		if err != nil {
			writeLog(err.Error())
			writeStr(logStr)
			pause()
			return
		}
	}

	if s == "3" {

		// b := make([]byte, 999)
		// fmt.Println("请输入抖音主页链接：\n->")
		// os.Stdin.Read(b)
		// url := strings.Split(string(b), "\n")[0]
		// url = "https://www.douyin.com/aweme/v1/web/aweme/post/?device_platform=webapp&aid=6383&channel=channel_pc_web&sec_user_id=MS4wLjABAAAA_tn0xx5dNiKWmJtzPNGNmSbWI6c2-qqJfddiZ1yCJ-F_k7zgCnBYnnylPjjz8Lfo&max_cursor=1695961566000&locate_query=false&show_live_replay_strategy=1&need_time_list=0&time_list_query=0&whale_cut_token=&cut_version=1&count=18&publish_video_strategy_type=2&pc_client_type=1&version_code=170400&version_name=17.4.0&cookie_enabled=true&screen_width=1792&screen_height=1120&browser_language=zh-CN&browser_platform=MacIntel&browser_name=Chrome&browser_version=118.0.0.0&browser_online=true&engine_name=Blink&engine_version=118.0.0.0&os_name=Mac+OS&os_version=10.15.7&cpu_core_num=12&device_memory=8&platform=PC&downlink=10&effective_type=4g&round_trip_time=100&webid=7292596081962878476&msToken=UWeH8gLpgtOODgCe5iZU-0_QT9SDP4PWyZKb0U4Je42aFrYoCN31tbqIywG7Qz9vVKp0kIgsUiWuJirhdIdcZgl2x8TCnmuRNgnRnMZqhaW7oinrAoVotmWzqA==&X-Bogus=DFSzswVuqJUANa9otYkXdBaWqIyi"
		// rsp, err := http.Get(url)
		// if err != nil {
		// 	return
		// }
		// content, err := ioutil.ReadAll(rsp.Body)
		// if err != nil {
		// 	return
		// }
		file, err := os.ReadFile("./videourl.txt")
		if err != nil {
			return
		}
		split := strings.Split(string(file), "###")
		fileNames := ""
		oldStr := "《再临烟火》如果许沁没有孟家的帮衬，会过上什么样的生活？#人间烟火  #许沁孟宴臣番外篇 #许沁孟宴臣同人文 #超爆小故事.mp4\n《噢是皇帝》我意外能听到皇上的心声 #皇帝裴与宁芝芝加长版 #沙雕文 #皇帝芒果过敏后续 #女生必看 #满级皇帝后续.mp4\n《孟婆别笑》我到了地府，凭一己沙雕之力谋得职位#小说推荐 #女生必看 #沙雕文 #爆笑沙雕文.mp4\n《战场宠妻》夫君打了胜仗，八百里加急送回一个姑娘 #王妃芳心八音盒后续  #热热加急八音盒加长版 #陆宴沈惜文加长版后续.mp4\n《月光回程》白月光回国了，所有人都嘲笑，陆亦寒会扔掉我这个替身 #陆亦寒白月光回国加长版 #相似鉴定后续    #白月光拉着我做亲子鉴定后续 #陆亦寒楚言喻加长版.mp4\n《梨涡优选》某天，京城叶家大少爷身边出现了一个清纯可爱的女孩，京城的闲人议论纷纷，都在讨论叶少是选择天降的小美女，还是选择青梅竹马的我，而我，冷冷一笑，然后恨不得一脚踢开这个黏在我腿上求抱抱的当事人 #叶执姜眠知乎小说后续 #高甜来袭  #甜文小说 #叶执姜眠 #拿捏秋日的轻盈感.mp4\n《梨涡充电》打错电话捡了个男朋友，给我爸打电话时拨错了号码，接通后，我稀里糊涂喊了一声爸，对面沉默两秒，然后笑了，我没那么老，可以叫哥哥，然后，没过多久，他就成了我男朋友 #甜文小说 #蛋仔派对 #小说推荐 #宝藏小说.mp4\n《梨涡兔兔》影帝发博：一会儿去表白，成功了改名叫小兔子乖乖，失败了叫小兔子不乖。十分钟后，他改名：乖你奶个杀掉兔子，炖了兔头兔腿炫炫炫。全网沉默了 #影帝发博一会去表白完结后续加长版 #不乖难言青春镜像私权谋吻小说后续加长版 #高甜来袭 #沙雕小甜文小说.mp4\n《梨涡冤总》我在酒吧搭讪高冷帅哥，帅哥看我一眼，报出一串电话号码，按完最后一个数字，上面浮现五个字，冤种提款机 #小说推荐 #甜文 #一口气看完系列.mp4\n《梨涡冷暖》结婚半年，沈砚舟对我一直不冷不淡，某天我在他办公室撞见了女同事对他的暧昧纠缠，我看了几秒后轻声开口，既然你有了喜欢的人，我们离婚吧 #沈砚舟黎梨 #沈砚舟黎梨后续 #蛋仔派对 #炒鸡好看小说.mp4\n《梨涡回甜》重回高中时代，我去找我那会哭唧唧撒娇的奶狗老公，没想到昔日奶狗变狼狗，曾经烟酒不沾的小奶狗熟练地叼着烟，漂亮的眉眼像钩子似的，难道你喜欢我 #重回高中时代 #裴烨 #裴烨沈清 #甜文 #裴烨沈清知乎小说后续.mp4\n《梨涡失眠》网恋到了黏人小奶狗，声音很性感，整日姐姐这姐姐那 #网恋到了黏人小奶狗 #周妄 #周妄江年江秋完结 #抖音搜索流量来了 #甜文.mp4\n《梨涡慢热》#我和一个警察闪婚了完整版 #陈小乔陈正泽完整版 #我和一个警察闪婚了后续加长版 #高甜来袭 #甜文.mp4\n《梨涡换味》我给床搭子发消息，要出差三个月，结束关系吧，他已读不回，半夜突然敲响我公寓的门，我隔着防盗门，让他快走，不然报警，他在外面冷笑，许娇娇，你报警吧，?警察来了，我就举报你票昌 #何劲许娇娇完结  #许娇娇何劲  #许娇娇何劲后续  #许娇娇何劲完结.mp4\n《梨涡期待》我将吃了一口的小蛋糕递给顾倾北，他刚接过，蛋糕却被一个女生打翻了，顾倾北，你将来是万人之上的京圈太子爷，不需要在吃这个女人的东西，没想到.... #顾倾北慕晚桥 #顾倾北慕晚桥结局 #顾倾北慕晚桥小说加长版 #蛋仔派对 #抖音搜索流量来了.mp4\n《梨涡爆裂》顶流竹马说他家水管爆裂，跑来我家洗澡，结果我误接了当红小花从直播综艺打来的视频通话，哥哥在干吗呀，他在洗澡，要不你一会儿再打回来，就在此时，浴室的竹马忽然扬声喊我，帮我拿条毛巾，一句话，全网都沸腾了 #顶流竹马说他家水管爆裂跑来我家洗澡  #沈靳白芝芝  #沈靳白芝芝完结  #抖音搜索流量来了  #拿捏秋日的轻盈感.mp4\n《梨涡相亲》为了让对方知难而退，被迫去相亲，我胡诌，我不孕，对面帅哥神色惊讶，呦，巧了，我不育，我干脆脱下外套，露出里面的旺仔紧身衣，他挑眉，伸出脚上的黄金切尔西，我，遇到对手了 #被迫去相亲为了让对方知难而退 #陈淮之宋时微 #陈淮之宋时微后续 #蛋仔派对.mp4\n《梨涡知足》给球场上的校霸送情书，我塞到他手里拔腿就跑，回寝室一摸兜，情书还在，姨妈巾没了，第二天，校霸把我堵在宿舍楼下，他眼圈青黑，语气悲愤 #给球场上的校霸送情书 #周止琰 #周止琰江钰  #周止琰江钰知乎小说后续 #甜奶情书后续.mp4\n《梨涡祈福》在寺庙祈福烧香，愣是把前面帅哥的羽绒服烫出了三个洞，那个……你好，他有些不耐烦地回头，说吧，要什么，微信还是手机号，我硬着头皮开口，支付宝吧，微信里没钱 #在寺庙祈福烧香 #齐旻陈恬 #齐旻陈恬祈福 #恋动祈福小说全文 #甜文小说.mp4\n《梨涡粉底》我弟早恋了，我还在开会，班主任连着给我打了五个电话，当我推开办公室门的时候，再白的粉底都遮不住我比锅底还黑的脸，我弟站在墙角，吊儿郎当的抬着下巴，要多狂有多狂，在班主任的惊愕目光中，我冲过去在齐思宇脑袋上狠狠敲了一下，给我站好，齐思宇立马乖乖站的笔直 #甜文小说 #小说推荐 #超爆小故事?.mp4\n《梨涡绑票》绑架富二代后，我跟他爸打电话要钱，那头怒吼，鸽吻，吻，我为难地看了一眼宽肩窄腰的少爷，呃，好，少爷，我的裤子后面有两个兜，一个是空的，另一个，也是空的 #绑架富二代后我跟他爸打电话要钱后续 #林颂年温怡 #林颂年温怡后续 #林颂年温怡小说 #抖音搜索流量来了 #蛋仔派对.mp4\n《梨涡续甜》十七岁时，我网恋了一个男大学生，骗他说我已经二十岁了，可有天打视频电话的时候，他看见了我身上穿的高中校服，于是毅然决然地和我提了分手，现在我读大一了，而他竟然成为了我的专业课老师 #十七岁时我网恋了一个大学生 #江栩郑清若  #江栩郑清若网恋   #蛋仔派对  #超级好看的小说.mp4\n《梨涡聚意》我的同桌是出了名的高冷帅哥，但他打架超凶且从不正眼瞧人。我小心翼翼和他相处，生怕惹到他。可有天放学我折返回校，本该空无一人的教室里坐满了平时难管教的人物。我的高冷同桌坐在最中间，一脸得意：有什么可炫耀的？都没我家槿一漂亮。#我的同桌是出了名的高冷帅哥  #槿一谢沉 #谢沉槿一知乎小说 #高甜来袭 #你在抖音搜什么.mp4\n《梨涡背面》末世来临，我艰难苟活，谁知丧尸王吃了个恋爱脑，无可救药地爱上了我，每天打开门，他都会讨好般地送上新鲜的食物，直到他看了本名叫《奸商准则》的书 #末世来临我艰难苟活 #末世来临我艰难苟活后续 #末世来临我艰难苟活完结 #蛋仔派对 #女生必看.mp4\n《梨涡血糖》我低血糖晕倒在校草腿边，他却以为我在模仿鹅妈妈假摔，哟，模仿得挺到位啊，我要是有意识，高低得骂他几句，后来校草发现我是低血糖晕倒，悔得半夜都想起来扇自己巴掌 #小说推荐 #甜文 #宝藏小说推荐.mp4\n《梨涡误解》我把老板当成是我爸，微信连续要了一礼拜钱，第七天，他说他要出国没信号了，我问他出去玩为何不带我，半小时后，我坐到了他的私人飞机上，他问我，你就这么爱，一天都离不开我，连续加班七天后，整个团队的人都很崩溃 #我把老板当成是我爸 #林阿娇林泾川 #林阿娇 #林阿娇后续 #蛋仔派对 #女生必看.mp4\n《梨涡误诊》我癌症晚期，无所畏惧地亲了死对头，第二天，医院说我拿错了病例，我没病，是误诊 #我癌症晚期 #池野姜筱姜小小小说后续  #高评分甜文  #抖音搜索流量来了.mp4\n《梨涡跟班》跟室友一起嗑校草跟校霸的 cp，把游戏名改成了偷穆城裤衩，套李岩头上，当晚五排，对面打野跟射手追着我俩杀 #穆城李岩洛小小 #穆城李岩洛小小知乎后续 #高甜来袭 #抖音搜索流量来了 #拿捏秋日的轻盈感.mp4\n《梨涡身姿》骑小电驴回家，路遇几个帅到爆炸的特警小哥哥，我看得太入迷，撞上电线杆翻车了，然后喜提了公主抱和社死热搜，女子因沉迷帅气特警颜值骑车撞向了电线杆，所幸的是，我把最帅最痞的那个特警，变成了我男朋友 #甜文小说  #小说推荐  #超爆小故事.mp4\n《梨涡轻甜》小时候，我对竹马说，我喜欢你好兄弟，可别跟他说，直到他好兄弟婚礼，我掐着竹马的脖子，你嘴是真严啊，他冷嗤，再敢说一句喜欢他，小心老子捶你 #小时候我对竹马说我喜欢你好兄弟 #小时候我对竹马说 #江厘程茶 #江厘程茶知乎小说后续 #江厘程茶后续.mp4\n《梨涡闪现》半夜十二点，我突然闪现到了警校男神的床上，  和男神正好面面相觑，  就在男神准备逮捕我时，我抢先一步抱住了他的腰，蜷缩在他怀里，无辜道，哥哥，贴贴，怕怕 #甜文 #魏泽秋姜盈 #魏泽秋姜盈知乎小说 #魏泽秋姜盈加长版 #拿捏秋日的轻盈感.mp4\n《梨涡预测》我在灯红酒绿的酒吧发了条信息给季淮，宝贝，晚安，结果他回，我在你右后方的卡座，过来跟老子碰一杯，吓得我立刻拿起我的包，火速逃跑 #小说推荐 #甜文 #宝藏小说推荐.mp4\n《梨涡高数》我作为被认回豪门的真千金，回家的第一天，我就被那个霸占我身份二十年的假千金给了一个狠狠的下马威，一套高数题，然后我一道题没做对 #我作为被认回豪门的真千金 #苏思挽苏乔乔 #苏思挽苏乔乔小说 #苏思挽苏乔乔结局 #蛋仔派对.mp4\n《独特觉醒》这是你想要的结局吗 #孟许时分  #许沁好像觉醒的纸片人  #孟宴臣  #许沁.mp4\n《疯疯成名》沙雕女主穿进虐文，会擦出什么火花 #小说推荐  #女生必看  #拯救书荒  #沙雕文.mp4\n《穿穿三十》全班三十人穿到虐文女主身上成为各个器官，应该怎么协调身体 #小说推荐  #一口气看完系列  #拯救文荒  #已完结  #沙雕文.mp4\n《穿穿手机》穿越到霸总的手机了，看霸总都干了些什么 #小说推荐 #文荒推荐 #女生必看 #沙雕文.mp4\n《穿穿读心》沙雕女主穿进虐文，会擦出什么火花 #小说推荐  #女生必看  #拯救书荒  #沙雕文.mp4\n《穿穿躺平》沙雕女主穿进虐文，会擦出什么火花 #小说推荐  #女生必看  #拯救书荒  #沙雕文.mp4\n一口气看完大结局的免费甜文#高甜来袭 #一口气看完系列 #女生必看.mp4\n一口气看完结局的甜文来了#高甜来袭 #抖音搜索流量来了 #一口气看完系列 #你在抖音搜什么.mp4\n二十二岁芳龄被迫相亲。对方身高腿长、海归硕士，随便逗一下就脸红，优质到让我一见钟情。为了拿下他，我整天装得端庄贤淑、笑不露齿 #渣男海后巅峰对决完整版 #22岁芳龄被迫相亲加长版 #两相喜爱甜甜整装后续完结 #二十二岁芳龄被迫相亲段延乔千亿甜文后续.mp4\n可以一口气看完大结局的免费甜文#高甜来袭 #一口气看完系列 #女生必看 #青梅竹马文.mp4\n周念同学，借我520车费参加歌唱比赛，等我红了还你520万。六年后他红了，竟然还管我借钱#周念梁声超甜故事后续 #十八爱言后续结局 #周念梁声借钱宣传反诈骗 #紫府声念后续完结 #声声希月天文后续.mp4\n在公交车上被大帅哥踩了一脚，他没道歉。趁众人不注意，我偷偷捏了下他身旁的大妈。没人发现，我又捏了一下。猛地抬头，发现帅哥咬牙切齿地瞪我，你她么往哪里捏呢？ #环绕爱意后续加长版 #公交车上捏大帅哥妈妈的辟谷后续完结 #沙雕小甜文 #姜与乐谈秋树小甜文后续.mp4\n大学毕业后我把男友踹了，谁知他竟成了顶流影帝，被问到有没有心动过，他大胆承认有，还爆出自己被甩了 #大学毕业后我把男朋友踹了 #谢奕白霍渴完结版 #谢奕白霍渴知乎小说后续结局 #高甜来袭 #抖音搜索流量来了.mp4\n如果你是《我的人间烟火》中的许沁，你会选择孟宴臣还是宋焰  #我的人间烟火宋焰许沁  #我的人间烟火孟宴臣  #孟宴臣许沁伪骨科 #许沁说宋焰是她的命 #孟宴臣说又不是养不起.mp4\n婆婆刚给小姑子买完车就被小姑子试驾撞断了腿。 等我赶到医院的时候，一家人整整齐齐等着我付钱。 小姑子哭诉:“嫂子，我真不是故意的，你帮帮咱妈吧。”” 我沉吟片刻:“那定是车故意的。#婆媳 #女生必看 #徐天秦淼徐娇结局 #完结侠 #徐天秦淼徐娇完结.mp4\n学霸男友总爱管着我，我受不了提出分手。他竟然主动抛弃我，换了座位。 #学霸男友总爱管着我#宋迟张晚小说 #宋迟张晚谢阳 #高甜来袭 #抖音搜索流量来了.mp4\n带儿子参加综艺，没想到儿子竟然让观众给他找爸爸，我竖起手指发誓，儿子：可是你竖起的是中指呀 轻置臀于评论区，完结踢#桑言陈隶思完结侠 #带儿子参加综艺后续 #高甜来袭 #女生必看.mp4\n当红小生是我老公。公司不允许他爆出恋情，他就暗戳戳地搞事情。#当红小生是我老公后学 #靳文彦薇薇小说后续.mp4\n我一直暗恋的高冷校草塌房了，视频上，他穿着背心戴着安全帽，蹲在工地门口吃盒饭，被采访时，生气骂道，神经病，五百万能干嘛，一辆能看的跑车都买不到，全校的人嘲笑他贪慕虚荣，是个心比天高、一身A货的假富二代，我却偷偷砸钱把他追到手了 #我一直暗恋的高冷校草塌房了 #萧越曲筱筱小说 #萧越曲筱筱 #甜文小说 #蛋仔派对.mp4\n我假期兼职带娃，接了一个号称三岁很乖的娃，没想到说的一个，实际上是两个 #女生必看 #沙雕文 #兼职带娃 #未完结.mp4\n我和嫂子同时怀孕，双喜临门，我妈大手一挥给我俩定了高档月子中心抖音搜索「黑岩故事会」搜索1074309看原文#婆媳  #秦淼张家豪后续结局 #后续  #女生必看.mp4\n我和嫂子同时怀孕，我妈定了同一家高级月子中心，没想到我一直没见到嫂子#秦淼孙一乔月子中心后续#月子风波秦淼完结.mp4\n我哥喝多了，把兄弟的腹肌照发给了我，看一眼就忘不掉了，后来跟我哥和他兄弟聚餐，我喝多了，对这腹肌就是一顿操作 #江聿风江稚鱼陆闻笙小说加长版  #我哥喝醉了把兄弟的腹肌照发给了我  #甜文.mp4\n我带儿子上综艺，儿子当场让观众给他找爸爸，后续来了#带儿子参加综艺的小说 #妈妈与儿子的综艺 #带儿子参加节目 #童染江湛江芝陈聿思桑颜完结后续.mp4\n我意外绑定了缺德系统，需要对男主缺德才能完成任务回到现实 #宋书言肖桐桐沈柔江叙后续 #宋书言肖桐桐沈柔江叙后续结局 #显眼包 #女生必看 #划走你就草率了系列.mp4\n我是原告律师，法庭中场休息，我走进卫生间却突然被人按在墙上，还没来得及看清?炙热的唇贴了过来，刚才在庭上冰冷理性的辩方律师，此刻搂着我的腰身 #我是原告律师法庭中场休息小说后续  #林苏 #我是原告律师 #萧许林苏完结 #萧许林苏番外.mp4\n我是富商的女儿，用作冲喜嫁入侯府。结果成亲第一天，老侯爷喝药呛死了。第二天，小姑子落水溺死了。第三天，小叔子骑马摔死了。一个月后，整个侯府死得只剩我和小侯爷了 #沙雕文 #爆笑小故事 #相当炸裂.mp4\n我爸是超雄综合征患者，不仅家暴我妈，还重男轻女，于是我裹着身体撑到青春期，终于忍不住反击他 #家暴零容忍 #超雄综合征 #面具下的生活后续 #女生必看 #抖音搜索流量来了.mp4\n我的死对头最近十分不对劲，我只不过像往常一样怼了他两句，这厮没回嘴就算了，竟然还哭了。#我的死对头不太对劲后续加长版 #贺明州盛绵绵结局 #贺明周绵绵后续 #高甜来袭 #女生必看.mp4\n我穿成了一只大鹅，兴奋的尖叫#高鹅预警后续 #穿书成为一只大鹅后续完结.mp4\n我穿成了小说里的贫穷路人甲。当我看到男主扔了恶毒女配送的银行卡，呵斥她: 我不需要你的施舍。女主还附和: 别瞧不起我们，莫欺少年穷。我立马捡起了那张银行卡:我需要!大小姐，来施舍我吧!来看不起我吧!#我穿成了小说里的贫穷路人甲#姜思思孟静温心周礼后续  #姜思思孟静温心完结  #蛋仔派对 #姜思思孟静温心周礼.mp4\n我绑定了真话系统，男朋友背着我和初恋上了恋综，我是观察员#合约宝贝小说宋清语后续完整版 #陈梦瑶黎翘小说完结 #陈慕白霜宋清语真话系统完结加长版.mp4\n我转学了，得知我是校霸妹妹后，不少人跟我告白，想要当校霸的妹夫，当晚，校霸把一个表白的男生一脚踹开，随即把我搂在怀里 #转学后得知我是校霸妹妹小说知乎 #转学后我成了校霸的妹妹小说后续 #秦楠李季肖完整版 #秦楠李季肖小说校霸3 #高甜来袭.mp4\n找工作面试被问，三年之内是否准备要小孩，我脑袋一机灵，就说：已婚，老公他中看不中用。#面试甜言后续完结 #找工作面试被问三年之内是否准备要小孩 #唉的二八定律 #楚宴宋颜颜甜文大结局 #高甜来袭.mp4\n校霸打架在主席台上念检讨书，我一不小心听到他的心声，真没想到他是这样的校霸 #校霸上台念检讨书 #校霸打架在升旗台念检讨书完结 #月夕糖糖小说后续  #满级求偶周西西完结后续结局.mp4\n梨  涡  律  师  怀孕测出两道杠，我给律师前任打电话，产检费，孕期营养费，麻烦你出一下，?对方气笑了，陈夏，有必要让我提醒你一下，我们分手三年了 #怀孕测出两道杠我给律师前任打电话 #裴律陈夏 #甜文小说 #蛋仔派对.mp4\n梨 涡 交 警：瘸腿过马路被一米九的交警哥哥单手抱起，他捏捏我的肉肉脸，谁家小孩啊，你家大人呢，我，侮辱，这对154的女大学生来说是赤裸裸的侮辱 #小说推荐  #甜文小说 #高甜来袭.mp4\n梨 涡 保 姆。穿成总裁文里的保姆是什么体验，谢邀，上一秒刚配合管家说完台词，少爷好久都没这样笑过了，就在我混吃混喝混薪水的时候，男二出现了，这题我会，男二是男主弟弟，偏执阴险，要抢女主叶蝉的，可谁告诉我，为什么当男主皱着眉说，除了叶蝉，我什么都可以给你 #穿成总裁文里的秘书 #穿成总裁文里的保姆 #沈瞬 #沈瞬沈存叶蝉 #甜文小说.mp4\n梨 涡 保 镖。舔了霸总三年，我得到了一个BE结局，为了补偿我，分手那天，陆川让我选一样东西带走，什么都行，我毫不犹豫地指向站在他身后的男人 #舔了霸总三年 #陆川霍潮霍启成小说后续 #甜文小说 #蛋仔派对 #抖音搜索流量来了.mp4\n梨 涡 免 三：一口气看完系列来了哦，宝子们，点个关注点个赞吧，我弟弟带着好几个大帅哥回来了，他们一群人看到我，向我打招呼，阿姨好，这，就离谱，不过，看着他们的脸，我强忍着内心的狂喜，温柔的说，坐呀，随便坐，最好坐我腿上，我一个嘴瓢，差点说出来 #一口气看完系列 #甜文小说 #炒鸡好看的小说.mp4\n梨 涡 免 二：一口气看完系列来了哦，宝子们，点个关注点个赞吧，逃婚后，我爹停了我所有银行卡，我发朋友圈发牢骚，其实我难过的时候，只需要一个拥抱和一千万而已，五秒后，死对头黎应向我转账一千万，五分钟之后，他一通电话打过来，下楼，来送拥抱了 #蛋仔派对 #甜文小说 #一口气看完系列 .mp4\n梨 涡 免 四：#一口气看完系列 #甜文小说 #小说推荐 ?.mp4\n梨 涡 再 会。19 岁这年，我雇了个保镖，人狠，话不多，可以面不改色地接受我的骚扰，我各种勾人的招数，在他这回回碰壁 #徐青野林汀晚 #徐青野林汀晚知乎小说后续 #甜文小说 #蛋仔派对 #原神夏日回响音乐会.mp4\n梨 涡 出 神：江城三中的孟辞以脾气臭、拳头硬闻名，然而，最近有同学撞见他苦着脸向新来的女转学生讨饶，我的小祖宗，你就背个英语单词吧 #小说推荐  #甜文小说 #高甜来袭?.mp4\n梨 涡 刺 猬。我的老板出差了，他让我把他家里的宠物刺猬领回家好好供起来，小刺猬的pp肉嘟嘟的，像水蜜桃一样可爱，我忍不住一戳再戳，我揉得正上头的时候，小刺猬猛地转过身来怒瞪着我，戳得很上瘾是吧，老子要开除你一百次 #高甜来袭 #抖音搜索流量来了 #甜文小说 #蛋仔派对.mp4\n梨 涡 单 车：男朋友每个月都给我二十万，我每天，恪守女德，绝不看他手机，从不追问行踪，偶尔撞到他和其他女生逛街，我比他还紧张，埋头跑得飞快，生怕坏他好事 #一口气看完系列 #甜文 #小说推荐.mp4\n梨 涡 哭 包：我的竹马是个性格高冷的爱哭鬼，人生最大梦想就是搞钱娶媳妇，进大学，我调侃他，让他好好地把握机会追校花，他捂着被子就哭了，说我欺负人，后来我才知道他一直想娶的姑娘原来是我呀 #甜文小说 #小说推荐 .mp4\n梨 涡 坦 白。我绑定了真话系统，男朋友背着我和初恋上了恋综，我是观察员#合约宝贝小说宋清语后续完整版 #陈梦瑶黎翘小说完结 #陈慕白霜宋清语真话系统完结加长版.mp4\n梨 涡 夏 至。竹马在篮球场上打球，我在场边加油打气，却被竹马嫌弃太吵，后来竹马的女友在边大喊加油时，竹马一脸心疼，让她注意嗓子，篮球赛上，所有人都以为我是来给竹马加油的，我却在对手进球的时候，欢呼喝彩  #竹马在篮球场上打球 #顾珩许妍林芷莹 #甜文小说 #抖音搜索流量来了.mp4\n梨 涡 好 感。穿书后，系统要我攻略病娇太子，为了回家，我做了，太子没人疼，我疼，太子没爱，我爱，太子没人护，我护，好不容易刷满好感度，完成任务，前脚我刚跑路回到家，后脚系统又让我穿了回去 #穿书后系统要我攻略病娇太子 #裴珺南音宁 #裴珺南音宁小说后续 #甜文小说 #抖音搜索流量来了.mp4\n梨 涡 小 鹿：我听过不少说什么青梅竹马不抵天降，我不相信直到我意外发现一本书，那里面的主角是我的竹马，而我只是个炮灰，我竟是一本书中的人物 #青梅竹马不抵天降  #林深时鹿 #甜文小说 #蛋仔派对.mp4\n梨 涡 尾 巴。DNA鉴定结果出来了，我和一个男生被抱错了20年，我现在的父母是房地产大亨，得知他们的孩子其实是个儿子，哭兮兮抱着我  #dna鉴定结果出来了 #丁柠谢行 #丁柠和行哥 #丁柠 #丁柠和谢行小说全文结局.mp4\n梨 涡 心 声。裴清然是个哑巴帅哥，先天哑巴，直到有一天，我竟然觉醒了读心术，还只能听见裴清然的内心的声音，于是我就发现了一个秘密    #裴清然苏小小吸血鬼 #裴清然苏小小吸血鬼配享太庙  #裴清然苏小小吸血鬼结局 #甜文小说.mp4\n梨 涡 拱 月。从乡下中学转到贵族学校后，我一不小心得了年级第一，校长女儿质疑我作弊，还说我这种乡下来的土包子不配和她在一个学校，更不配跟校草搭话 #从乡下中学转到贵族学校后我一不小心得了年级第一 #林安好 #林安好时锐小说 #林安好后续 #林安好时锐白优优.mp4\n梨 涡 换 运。闺蜜主动帮我给校霸送告白信，还说校霸答应了，让我当众亲他一口，他好公开我，实际上她早就跟校霸在一起了，我笑笑，好，亲，扭头亲上了上一世一直守护我的竹马 #闺蜜主动帮我给校霸送告白信 #阮湘 #阮湘唐怡小说 #阮湘唐怡小说后续 #甜文小说.mp4\n梨 涡 旗 袍。我弟高考是我送去的，当时我穿了一件高开叉旗袍，他买的，而且他用攒了一年的零花钱贿赂我，让我穿着送他去考试，美其名曰，祝他旗开得胜 #我弟高考是我送去的 #林知之 #林知之原净小说后续 #甜文小说 #抖音搜索流量来了.mp4\n梨 涡 无 言。陆琛说我是他追求者里最舔的，舔狗里最美的，我每次都笑笑不说话，直到他的白月光回国，陆琛叼着烟和我提分手，被我一巴掌把烟都抽飞了 #陆琛说我是他追求者里最舔的 #陆琛 #陆琛秦舒颜 #陆琛秦舒颜裴延 #甜文小说.mp4\n梨 涡 气 质：我喜欢竹马十年，他厌烦我也十年，终于有一天，我将他吃干抹净，还留下两百块钱，竹马倍感羞辱，四处，追杀，我，可他哪儿都找不到我，躲了他三年，直到父亲去世，我成了孤儿 #我喜欢竹马十年 #贺放闻栀 #贺放闻栀完结 #贺放闻栀小说 #贺放闻栀后续.mp4\n梨 涡 烟 雨：和竹马结婚一年，他经常绯闻缠身，直到他白月光回国的那天，我没回去，他遍地找我，差点掀了整座城，接通电话时，他声音发抖，烟烟，为什么不回家，我下班回到家的时候，二楼传来一阵嬉闹声，刚换好鞋，一个裹着浴袍的女人跑了下来  #和竹马结婚一年他经常绯闻缠身 #陆尧许南烟 #陆尧许南烟后续 #蛋仔派对 #甜文小说.mp4\n梨 涡 独 钟：我把喜欢的电竞选手照片发给了游戏搭子，这哥们长得够帅，游戏搭子，嗯，你要不看看他同队的那个江时延，比他帅上几千几万倍，操作也 6，我故意抬杠，太拽了，不喜欢 #小说推荐 #甜文 #炒鸡好看的故事 #超爆小故事?.mp4\n梨 涡 畅 聊：我跟电信骗子畅聊三天，成功被警察传唤，警局里，我一抬头，霍，巧了，前男友你咋还活着呢 #甜文小说 #小说推荐 .mp4\n梨 涡 痴 迷：我和施一青梅竹马，我们俩是幼儿园同学，三岁的时候，我妈妈给我穿了一条红色小裙子参加幼儿园舞蹈表演，一下子迷住了好多同学，其中施一最主动，他拿个小板凳坐在我旁边，对我说，欢欢，你可真好看，你当我女朋友好不好 小说推荐 #一口气看完系列 #甜文?.mp4\n梨 涡 绿 友：加班的男友突然给我发消息，已有女友，互删谢谢，我缓缓的打了一个问号过去，于是看到了传说中的红色感叹号，我的男朋友有女朋友了，那我是谁 #甜文小说  #小说推荐 #超爆小故事?.mp4\n梨 涡 胎 记：那天，孤儿院门口，一对中年夫妻找了过来。他们红着眼找我打听消息：请问，这里有个叫雯雯的女孩吗？叔叔阿姨，你们是谁？?我歪着头，露出脖子后的胎记，跟他们女儿一模一样的胎记 #甜文小说  #小说推荐  #超爆小故事.mp4\n梨 涡 芝 芝：我觉醒的时候，故事已经接近尾声，团宠女主得到了包括我爸、我哥和我未婚夫在内的所有人的爱，而我则奄奄一息地躺在医院的病床上，就在我即将陷入书中的昏迷结局前，病床前忽然多出一个人 #我觉醒的时候故事已经接近尾声  #李苏叶程在河小说 #李苏叶程在河完结 #蛋仔派对 #甜文小说 .mp4\n梨 涡 记 录。综艺上游戏失败，被问到和好友最近的聊天内容，彼时刚吵完架的酷拽竹马正不停给我发消息，吵完架就冷暴力我是吧，可恶，就你有脾气是吧，那你怎么不把我拉黑了，服了，别理我呗，下一秒，支付宝到账声响起，支付宝到账52000元 #谭绪路稚遥 #蛋仔派对 #谭绪路椎谣江妄后续 #甜文小说 #抖音搜索流量来了 .mp4\n梨 涡 蹙 眉；做乳腺检查却遇到了前男友，顾谨之冰凉的手指触到我的那一刻，脸颊连着耳根，滚烫到了极点，顾谨之眉头微蹙，有几分忧心，陆宛虞，你就是这样照顾自己的，正在熬夜肝产品资料时左胸一阵针刺样的疼痛袭来，回想这样的情况好像已经连续出现了一个月，急忙打开手机挂了个湘安医院乳腺专家号 #小说推荐 #甜文 #炒鸡好看的故事 #超爆小故事.mp4\n梨 涡 转 换。攻略者攻略男主失败后，将攻略目标换成了男主的兄弟，也就是我的竹马江拓，想让男主追妻火葬场，她在我把吃不掉的早餐扔到江拓怀里时，大声斥责我说江拓不是我的垃圾桶，还在江拓帮我带饭打水时，认真地告诉他，他不是我的奴隶 #攻略者攻略男主失败后 #江拓姜栀黎染结局 #江拓姜#江拓姜黎染完结 #江拓差栀黎染周佚后续  .mp4\n梨 涡 辣 妹：和室友打赌输了，他们让我穿辣妹装给校霸送水1于是，我翻出了我压箱底的一条黑色百褶裙，撸了一个全妆，室友还帮我卷了一个大波浪，我迅速的去超市买了一瓶气泡水，和室友便出发了到了篮球场 #甜文小说 #小说推荐 #超爆小故事 .mp4\n梨 涡 雇 主。我舔了江岸三年，给他做饭洗衣，还帮他写选修作业，后来有人问他，江岸，林晩月舔了你这么久了，你就没动心，昏暗的灯光下，我听见他轻笑了一声，嗯，她只是个保姆而已啊 #我舔了江岸三年 #林晩月江岸#林晩月江岸小说  #蛋仔派对 #抖音搜索流量来了.mp4\n甜甜甜甜，真的太甜了，我愿意打10分，一口气看完系列，?全文已更新大结局，放心观看 #高甜来袭 #女生必看 #甜文 #抖音搜索流量来了 #艾特你想艾特的人.mp4\n穿越后我绑定了缺德系统，只要缺德就获得奖励 未完结，已留臀，等踢 #穿越后我绑定了缺德系统后续 #解压小游戏 #超爆小故事 #缺德系统.mp4\n被男朋友耍了之后。我一怒之下冲进了男寝 303，将一份变态辣的大盘鸡盖在了他的头上。旁边校霸无辜躺枪，被甩了一身的红点子。之后校霸围堵我。我盯着他的红内裤跑了神。看哪呢。我脱口：你今年本命年吗。 #被男朋友耍了以后我一怒之下冲进男寝303 #舒郎林梦白小说完整版 #高甜来袭 #抖音搜索流量来了 #舒郎林梦白知乎小说加长版.mp4\n闺蜜的手机来电，你朋友好像喝醉了，现在正在大街上，抱着我朋友的大腿不撒手，非说我朋友是渣男 #周末闺蜜一个电话把我吵醒 #夏橙子许诚 #许诚夏橙子 #许诚夏橙子结局 #抖音搜索流量来了.mp4\n闺蜜被拐，我也跟上车：挤一挤，挤一挤，我也上去，我朋友在上面。谁料人贩子嫌弃地看了我一眼：你太蠢了，不要！我被羞辱得爆哭出声，冲上去乱啃他屁股，然后掀飞了他的车子。#闺蜜被拐我也跟上车后续完结 #女大学生跟闺蜜被拐 #她会发疯2完结后续加长版 #顾瑶周柏苏眠徐烟沙雕小说后续完结 #孤注一掷反诈骗发疯文学.mp4\n靠近老板五米内，我就能听到他的心声 #听到老板的心声小说后续完结 #林安安谢屿听见老板心声甜文后续 #林安安听见老板心声完整版 #甜文.mp4\n"
		oldStrArr := strings.Split(oldStr, "\n")
		for _, s := range split {
			stru := make([]*AwemeStruct, 0)
			err := json.Unmarshal([]byte(s), &stru)
			if err != nil {
				return
			}
			flag := true
			for _, struc := range stru {

				for _, str := range oldStrArr {
					if strings.Contains(str, struc.Desc) {
						flag = false
						break
					}
				}
				if !flag {
					continue
				}
				videoByte, err := http.Get(struc.Video.PlayAddr.UrlList[2])
				if err != nil {
					return
				}
				bytesData, err := ioutil.ReadAll(videoByte.Body)
				if err != nil {
					return
				}
				err = os.WriteFile(struc.Desc+".mp4", bytesData, os.ModePerm)
				if err != nil {
					fmt.Println(err.Error())
					return
				}
				fileNames += struc.Desc + "\n"
			}
		}
		os.WriteFile("FileName"+time.Now().Format("20060102150405")+".txt", []byte(fileNames), os.ModePerm)
	}
	writeStr(logStr)
	pause()
}

func textProcess() error {
	_, err := os.Open("./log")
	if err != nil {
		if os.IsNotExist(err) {
			err := os.Mkdir("./log", os.ModePerm)
			if err != nil {
				fmt.Println("异常退出")
				return err
			}
		}
	}

	files := make([]string, 0)
	_, err = os.Open("./source")
	if err != nil {
		writeLog(err.Error())
		if os.IsNotExist(err) {
			fmt.Println("source文件夹不存在，已自动创建，请将源文档放入source文件夹内")
			writeLog("source文件夹不存在，已自动创建，请将源文档放入source文件夹内后，重新运行软件")
			err := os.Mkdir("./source", os.ModePerm)
			if err != nil {
				writeLog(err.Error())
				return err
			}
		}
	}
	_, err = os.Open("./target")
	if err != nil {
		writeLog(err.Error())
		if os.IsNotExist(err) {
			fmt.Println("target文件夹不存在，已自动创建")
			writeLog("target文件夹不存在，已自动创建")
			err := os.Mkdir("./target", os.ModePerm)
			if err != nil {
				writeLog(err.Error())
				return err
			}
		}
	}
	err = filepath.Walk("./source", func(path string, info fs.FileInfo, err error) error {
		if info == nil {
			return nil
		}

		if info.IsDir() {
			return nil
		}
		if strings.HasSuffix(path, "txt") {
			files = append(files, path)
			writeLog("找到文件：" + path)
		}
		return nil
	})
	if err != nil {
		writeLog(err.Error())
		return err
	}
	fmt.Println("待处理文件：\n" + strings.Join(files, ",\n"))
	writeLog("待处理文件：\n" + strings.Join(files, ",\n"))
	for _, file := range files {
		err := process(file)
		if err != nil {
			writeLog(err.Error())
			return err
		}
	}
	writeLog("恭喜你，处理完成！")
	return nil
}

func getDateDiff() (float64, error) {
	nowNet, err := http.Get("http://time.tianqi.com/")
	if err != nil {
		fmt.Println("网络连接失败！")
		writeLog("网络连接失败")
		return -1, err
	}
	strDate := nowNet.Header.Get("Date")
	split := strings.Split(strDate, " ")
	fmt.Println(split)
	y, _ := strconv.Atoi(split[3])
	m, _ := MonthMap[split[2]]
	d, _ := strconv.Atoi(split[1])
	timeNet := strings.Split(split[4], ":")
	h, _ := strconv.Atoi(timeNet[0])
	M, _ := strconv.Atoi(timeNet[1])
	s, _ := strconv.Atoi(timeNet[2])
	lastAuthTime := time.Date(2023, 10, 28, 22, 00, 00, 0, time.Local)
	now := time.Date(y, m, d, h+8, M, s, 0, time.Local)
	diff := now.Sub(lastAuthTime).Hours()
	return diff, nil
}

// 通过链接获取文案，并写成txt文档，如果是小程序，则需要返回文字 TODO
func getTextByUrl() error {
	fmt.Println("请输入知乎链接(https://www.zhihu.com/market/paid_column/xxx/section/xxxx类型的专栏链接无法获取60%的字数)：")
	b := make([]byte, 999)
	os.Stdin.Read(b)
	url := string(b)
	if strings.HasPrefix(url, "https://soia.zhihu.com/tab/") {
		codes := strings.Split(url, "mst=")
		if len(codes) < 2 {
			fmt.Println("链接错误")
			writeLog("链接错误：" + url)
		}
		sec := strings.Split(strings.ReplaceAll(codes[1], "\r", "\n"), "\n")[0]
		url = "https://story.zhihu.com/blogger/next-manuscript/paid_column/" + sec
	}
	content, err := getHtml(url, 180)
	if err != nil {
		return err
	}
	// 处理标签，提取文案
	contentAfterTrim := trimHtml(content)

	randStr := time.Now().Format("20060102150405")
	cFile, err := os.Create("./source/" + randStr + ".txt")
	if err != nil {
		return err
	}
	_, err = cFile.Write([]byte(contentAfterTrim))
	if err != nil {
		return err
	}
	defer cFile.Close()

	return nil
}

func getHtml(url string, ottime int) (html string, err error) {
	// 参数配置
	options := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", true), // 是否打开浏览器调试
		chromedp.UserAgent(_ua),         // 设置User-Agent
		chromedp.ExecPath("C:\\Program Files\\Google\\Chrome\\Application\\chrome.exe"),
	}
	options = append(chromedp.DefaultExecAllocatorOptions[:], options...)
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), options...)
	defer cancel()

	// 创建chrome实例
	ctx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	defer cancel()
	// 设置超时时间
	ctx, cancel = context.WithTimeout(ctx, time.Duration(ottime)*time.Second)
	defer cancel()

	if err := chromedp.Run(ctx,
		chromedp.Tasks{
			// 打开导航
			chromedp.Navigate(url),
			// 等待元素加载完成
			chromedp.WaitVisible("body", chromedp.ByQuery),
			// 获取html
			chromedp.OuterHTML("html", &html, chromedp.ByQuery),
		},
	); err != nil {
		return "", err
	}
	return html, err
}

// https://www.zhihu.com/market/paid_column/1688557124559572992/section/1699829542443814912
// https://www.zhihu.com/question/512785324/answer/3218241171?utm_psn=1694272271680839680
// https://zhuanlan.zhihu.com/p/654036002

func trimHtml(s string) string {
	destStr := ""
	s = strings.Split(s, "<body")[1]
	reg := regexp.MustCompile("<path.*?>.*?</path>")
	s = string(reg.ReplaceAll([]byte(s), []byte("")))
	reg = regexp.MustCompile("<p.*?>(.*?)</p>")
	allString := reg.FindAllString(s, -1)
	for _, item := range allString {
		reg = regexp.MustCompile("<p.*?>")
		newStr := string(reg.ReplaceAll([]byte(item), []byte("\n")))
		newStr = strings.ReplaceAll(newStr, "</p>", "，")
		reg = regexp.MustCompile("<span.*>")
		newStr = string(reg.ReplaceAll([]byte(newStr), []byte("，")))
		reg = regexp.MustCompile("</span>")
		newStr = string(reg.ReplaceAll([]byte(newStr), []byte("，")))
		newStr = strings.ReplaceAll(newStr, "扫码下载知乎 App", "")
		newStr = strings.ReplaceAll(newStr, "<br>", "\n")
		destStr += newStr
	}
	return destStr
}

func writeLog(s string) {
	logStr += "\n" + s
}

func writeStr(s string) {
	os.WriteFile("./log/text_process.log", []byte(s), os.ModePerm)
}

func pause() {
	fmt.Println("文件已经处理完成，请到target文件夹查看！")
	fmt.Println("如有疑问请联系作者：buffer5")
	b := make([]byte, 1)
	os.Stdin.Read(b)
}

func process(fileName string) error {
	writeLog("处理文件：" + fileName)

	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	r := bufio.NewReader(file)
	str := ""
	for {
		line, err := r.ReadString('\n')
		str += line
		if err == io.EOF {
			break
		} else if err != nil {
			writeLog("error reading file " + err.Error())
			return err
		}

	}
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	if strings.ContainsAny(wd, "benfei") ||
		strings.ContainsAny(wd, "MR") {
		str = prefix + str
	}

	// 换行空格 \n　　=>\n
	reg := regexp.MustCompile("\n[　　]*")
	str = string(reg.ReplaceAll([]byte(str), []byte("\n")))

	// 换行\r
	reg = regexp.MustCompile("\r+")
	str = string(reg.ReplaceAll([]byte(str), []byte("")))

	// 章节序号
	reg = regexp.MustCompile("\n\\d+[.|，|、|,]*\n*")
	str = string(reg.ReplaceAll([]byte(str), []byte("\n")))

	reg = regexp.MustCompile("[\u4e00-\u9fa5]\\.+")
	strs := reg.FindAllString(str, -1)
	for _, strT := range strs {
		strT1 := strings.ReplaceAll(strT, ".", "")
		str = strings.ReplaceAll(str, strT, strT1)
	}

	reg = regexp.MustCompile("[\u4e00-\u9fa5][:|：]")
	strs = reg.FindAllString(str, -1)
	for _, strT := range strs {
		strT1 := strings.ReplaceAll(strT, ":", "，")
		strT1 = strings.ReplaceAll(strT1, "：", "，")
		str = strings.ReplaceAll(str, strT, strT1)
	}

	// 所有符号替换成中文逗号
	reg = regexp.MustCompile("[—|、|《|》|『|』|“|”|\"|【|】|\\[|\\]|「|」|{|}|\\(|\\)|（|）|？|\\?|！|\\!|,|。|：|~|~|～|(&nbsp;)]")
	str = string(reg.ReplaceAll([]byte(str), []byte("，")))

	// 所有符号替换成中文逗号
	reg = regexp.MustCompile("\\.{2,}")
	str = string(reg.ReplaceAll([]byte(str), []byte("，")))

	// 中国 人=>中国人
	reg = regexp.MustCompile("[\u4e00-\u9fa5] +")
	strs = reg.FindAllString(str, -1)
	for _, strT := range strs {
		strT1 := strings.ReplaceAll(strT, " ", "")
		str = strings.ReplaceAll(str, strT, strT1)
	}

	// 15 岁=>15岁
	reg = regexp.MustCompile("\\d +[\u4e00-\u9fa5]")
	strs = reg.FindAllString(str, -1)
	for _, strT := range strs {
		strT1 := strings.ReplaceAll(strT, " ", "")
		str = strings.ReplaceAll(str, strT, strT1)
	}

	// 替换逗号后，出现换行tab
	reg = regexp.MustCompile("　+")
	str = string(reg.ReplaceAll([]byte(str), []byte("")))

	reg = regexp.MustCompile("\n，")
	str = string(reg.ReplaceAll([]byte(str), []byte("")))

	// 省略号
	reg = regexp.MustCompile("…+")
	str = string(reg.ReplaceAll([]byte(str), []byte("\n")))

	// 所有换行符删除
	reg = regexp.MustCompile("\n+")
	str = string(reg.ReplaceAll([]byte(str), []byte("")))

	// 替换成逗号后，可能出现多个逗号相连，将多个逗号替换为一个逗号
	reg = regexp.MustCompile("，+")
	str = string(reg.ReplaceAll([]byte(str), []byte("，")))

	reg = regexp.MustCompile("\\d+[\u4e00-\u9fa5]")
	strs = reg.FindAllString(str, -1)
	if len(strs) >= 5 {
		fmt.Println("检测到可能是章节序号的数字，请选择是否需要处理：\n1-删除\n2-保留\n3-全部忽略")
		fmt.Println(strings.Join(strs, "\n"))
		for _, strT := range strs {
			fmt.Println(strT)
			b := make([]byte, 999)
			os.Stdin.Read(b)
			s := strings.Split(string(b), "\n")[0]
			s = strings.Split(s, "\r")[0]
			if s == "1" {
				reg = regexp.MustCompile("\\d+")
				strT1 := string(reg.ReplaceAll([]byte(strT), []byte("")))
				str = strings.ReplaceAll(str, strT, strT1)
			}
			if s == "2" {
				continue
			}
			if s == "3" {
				break
			}
		}
	}

	// 替换敏感词
	for {
		if strings.HasPrefix(str, "，") {
			str = strings.Replace(str, "，", "", 1)
		} else {
			break
		}
	}
	for k, v := range SensitiveWordsMap {
		str = strings.ReplaceAll(str, k, v)
	}

	if strings.ContainsAny(wd, "benfei") ||
		strings.ContainsAny(wd, "MR") {
		if !strings.HasSuffix(str, "，") {
			str += "，"
		}
		str += suffix
	}

	targetName := strings.Split(fileName, "\\")
	if len(targetName) == 1 {
		targetName = strings.Split(fileName, "/")
	}
	err = os.WriteFile("./target/副本"+targetName[1], []byte(str), os.ModePerm)
	if err != nil {
		writeLog(err.Error())
		return err
	}
	writeLog("文件：" + fileName + "处理完成")
	return nil
}

var SensitiveWordsMap = map[string]string{
	"婊子": "小可爱",
	"他妈": "特么",
	"妈的": "玛德",
	"贱人": "下等人",
	"自杀": "寻短见",
	"变态": "变小态",
}

var MonthMap = map[string]time.Month{
	"Jan":  time.January,
	"Feb":  time.February,
	"Mar":  time.March,
	"Apr":  time.April,
	"May":  time.May,
	"Jun":  time.June,
	"Jul":  time.July,
	"Aug":  time.August,
	"Sept": time.September,
	"Oct":  time.October,
	"Nov":  time.November,
	"Dec":  time.December,
}

type AwemeStruct struct {
	AwemeId string `json:"aweme_id"`
	Desc    string `json:"desc"`
	Video   struct {
		PlayAddr struct {
			UrlList []string `json:"url_list"`
		} `json:"play_addr"`
	} `json:"video"`
}
