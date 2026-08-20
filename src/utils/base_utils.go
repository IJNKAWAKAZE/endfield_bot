package utils

import (
	"context"
	bot "endfield_bot/config"
	"fmt"
	"github.com/go-redis/redis/v8"
	tgbotapi "github.com/ijnkawakaze/telegram-bot-api"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/mxschmitt/playwright-go"
	"gorm.io/gorm"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

var ctx = context.Background()

var WebC = make(chan error, 10)

var browser playwright.Browser

type GroupInvite struct {
	Id           string    `json:"id" gorm:"primaryKey"`
	GroupName    string    `json:"groupName"`
	GroupNumber  int64     `json:"groupNumber"`
	UserName     string    `json:"userName"`
	UserNumber   int64     `json:"userNumber"`
	MemberName   string    `json:"memberName"`
	MemberNumber int64     `json:"memberNumber"`
	CreateTime   time.Time `json:"createTime" gorm:"autoCreateTime"`
	UpdateTime   time.Time `json:"updateTime" gorm:"autoUpdateTime"`
	Remark       string    `json:"remark"`
}

type GroupJoined struct {
	Id          string    `json:"id" gorm:"primaryKey"`
	GroupName   string    `json:"groupName"`
	GroupNumber int64     `json:"groupNumber"`
	News        int64     `json:"news"`
	Reg         int       `json:"reg"`
	Welcome     string    `json:"welcome"`
	RequestMode int64     `json:"requestMode"`
	CreateTime  time.Time `json:"createTime" gorm:"autoCreateTime"`
	UpdateTime  time.Time `json:"updateTime" gorm:"autoUpdateTime"`
	Remark      string    `json:"remark"`
}

// NewId 生成UUID
func NewId() string {
	id, _ := gonanoid.New(32)
	return id
}

// SaveInvite 保存邀请记录
func SaveInvite(message *tgbotapi.Message, member *tgbotapi.User) {
	id := NewId()
	groupInvite := GroupInvite{
		Id:           id,
		GroupName:    message.Chat.Title,
		GroupNumber:  message.Chat.ID,
		UserName:     message.From.FullName(),
		UserNumber:   message.From.ID,
		MemberName:   member.FullName(),
		MemberNumber: member.ID,
	}

	bot.DBEngine.Table("group_invite").Create(&groupInvite)
}

// SaveJoined 保存入群记录
func SaveJoined(message *tgbotapi.Message) {
	id := NewId()
	groupJoined := GroupJoined{
		Id:          id,
		GroupName:   message.Chat.Title,
		GroupNumber: message.Chat.ID,
		News:        0,
		Reg:         -1,
		Welcome:     "",
		RequestMode: 0,
	}

	bot.DBEngine.Table("group_joined").Create(&groupJoined)
}

// GetJoinedByChatId 查询入群记录
func GetJoinedByChatId(chatId int64) *gorm.DB {
	return bot.DBEngine.Raw("select * from group_joined where group_number = ? limit 1", chatId)
}

// DownloadFile 下载tg文件
func DownloadFile(fileId string) (io.ReadCloser, string) {
	fileUrl, _ := bot.Endfield.GetFileDirectURL(fileId)
	fileType := fileUrl[strings.LastIndex(fileUrl, ".")+1:]
	response, _ := http.Get(fileUrl)
	body := response.Body
	return body, fileType
}

// GetAccountByUserId 查询账号信息
func GetAccountByUserId(userId int64) *gorm.DB {
	return bot.DBEngine.Raw("select * from user_account where user_number = ? limit 1", userId)
}

// GetAccountByUserIdAndSklandId 查询账号信息
func GetAccountByUserIdAndSklandId(userId int64, sklandId string) *gorm.DB {
	return bot.DBEngine.Raw("select * from user_account where user_number = ? and skland_id = ?", userId, sklandId)
}

// GetAccountByUid 查询账号信息
func GetAccountByUid(userId int64, uid string) *gorm.DB {
	return bot.DBEngine.Raw("select t.* from user_account t, user_player t1 where t.id = t1.account_id and t.user_number = ? and t1.uid = ? limit 1", userId, uid)
}

// GetPlayersByUserId 查询绑定角色列表
func GetPlayersByUserId(userId int64) *gorm.DB {
	return bot.DBEngine.Raw("select * from user_player where user_number = ?", userId)
}

// GetPlayerByUserId 查询绑定角色
func GetPlayerByUserId(userId int64, uid string) *gorm.DB {
	return bot.DBEngine.Raw("select * from user_player where user_number = ? and uid = ?", userId, uid)
}

// GetAutoSign 查询自动签到用户
func GetAutoSign() *gorm.DB {
	return bot.DBEngine.Raw("select * from user_sign")
}

// GetAutoSignByUserId 查询自动签到用户
func GetAutoSignByUserId(userId int64) *gorm.DB {
	return bot.DBEngine.Raw("select * from user_sign where user_number = ?", userId)
}

// GetSanityReminders 获取开启理智提醒的用户
func GetSanityReminders() *gorm.DB {
	return bot.DBEngine.Raw("select * from user_sanity_reminder")
}

// GetSanityReminderByUserId 查询用户是否开启理智提醒
func GetSanityReminderByUserId(userId int64) *gorm.DB {
	return bot.DBEngine.Raw("select * from user_sanity_reminder where user_number = ? limit 1", userId)
}

// GetNewsGroups 获取开启消息推送的群组
func GetNewsGroups() []int64 {
	var groups []int64
	bot.DBEngine.Raw("select group_number from group_joined where news = 1 group by group_number").Scan(&groups)
	return groups
}

// RedisSet redis存值
func RedisSet(key string, val interface{}, expiration time.Duration) {
	err := bot.GoRedis.Set(ctx, key, val, expiration).Err()
	if err != nil {
		log.Println(err)
	}
}

// RedisGet redis取值
func RedisGet(key string) string {
	val, err := bot.GoRedis.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return ""
		}
		log.Println(err)
	}
	return val
}

// RedisIsExists 判断redis值是否存在
func RedisIsExists(key string) bool {
	val := RedisGet(key)
	if val == "" {
		return false
	}
	return true
}

// RedisDel redis根据key删除
func RedisDel(key string) {
	err := bot.GoRedis.Del(ctx, key).Err()
	if err != nil {
		log.Println(err)
	}
}

// RedisScanKeys 扫描匹配keys
func RedisScanKeys(match string) (*redis.ScanIterator, context.Context) {
	return bot.GoRedis.Scan(ctx, 0, match, 0).Iterator(), ctx
}

// RedisSetList redis添加链表元素
func RedisSetList(key string, val interface{}) {
	err := bot.GoRedis.RPush(ctx, key, val).Err()
	if err != nil {
		log.Println(err)
	}
}

// RedisGetList redis获取所有链表元素
func RedisGetList(key string) []string {
	val, err := bot.GoRedis.LRange(ctx, key, 0, -1).Result()
	if err != nil {
		if err == redis.Nil {
			return nil
		}
		log.Println(err)
	}
	return val
}

// RedisDelListItem redis移除链表元素
func RedisDelListItem(key string, val string) {
	err := bot.GoRedis.LRem(ctx, key, 0, val).Err()
	if err != nil {
		log.Println(err)
	}
}

// RedisAddSet redis集合添加元素
func RedisAddSet(key string, val string) {
	err := bot.GoRedis.SAdd(ctx, key, val).Err()
	if err != nil {
		log.Println(err)
	}
}

// RedisSetIsExists redis集合是否包含元素
func RedisSetIsExists(key string, val string) bool {
	exists, err := bot.GoRedis.SIsMember(ctx, key, val).Result()
	if err != nil {
		log.Println(err)
	}
	return exists
}

// RedisDelSetItem redis移除集合元素
func RedisDelSetItem(key string, val string) {
	err := bot.GoRedis.SRem(ctx, key, val).Err()
	if err != nil {
		log.Println(err)
	}
}

// Screenshot 屏幕截图
func Screenshot(url string, waitTime float64, scale float64) ([]byte, error) {
	if browser != nil && !browser.IsConnected() {
		browser = nil
	}
	if browser == nil {
		pw, err := playwright.Run()
		if err != nil {
			log.Println("未检测到playwright，开始自动安装...")
			if installErr := playwright.Install(&playwright.RunOptions{Browsers: []string{"chromium"}}); installErr != nil {
				return nil, fmt.Errorf("playwright安装失败: %w", installErr)
			}
			pw, err = playwright.Run()
			if err != nil {
				return nil, fmt.Errorf("playwright启动失败: %w", err)
			}
		}
		browser, err = pw.Chromium.Launch()
		if err != nil {
			log.Println(err)
			return nil, fmt.Errorf("playwright启动失败: %w", err)
		}
	}
	page, err := browser.NewPage(playwright.BrowserNewPageOptions{DeviceScaleFactor: &scale})
	if err != nil {
		return nil, fmt.Errorf("创建页面失败: %w", err)
	}
	defer func() {
		log.Println("关闭playwright")
		page.Close()
	}()
	log.Println("开始进行截图...")
	resp, err := page.Goto(url, playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateDomcontentloaded,
	})
	if err != nil {
		return nil, fmt.Errorf("页面加载失败: %w", err)
	}
	if resp != nil && resp.Status() >= 400 {
		return nil, fmt.Errorf("页面加载失败，状态码：%d", resp.Status())
	}
	if len(WebC) > 0 {
		e := <-WebC
		return nil, e
	}
	// 只等待截图需要的字体和图片，避免无关的持续请求阻塞截图。
	if _, err := page.WaitForFunction(`async () => {
    const fallback = new URL("/assets/state/default_char.png", window.location.href).href;
    const delay = (ms) => new Promise(resolve => setTimeout(resolve, ms));
    const waitForImage = (image) => {
        if (image.complete) return Promise.resolve();
        return new Promise(resolve => {
            const done = () => {
                image.removeEventListener("load", done);
                image.removeEventListener("error", done);
                resolve();
            };
            image.addEventListener("load", done, { once: true });
            image.addEventListener("error", done, { once: true });
            if (image.complete) done();
        });
    };
    const images = Array.from(document.images);
    const resources = (async () => {
        if (document.fonts && document.fonts.ready) await document.fonts.ready;
        await Promise.all(images.map(waitForImage));
        return false;
    })();
    const timedOut = await Promise.race([resources, delay(8000).then(() => true)]);
    const failedImages = images.filter(image => !image.complete || image.naturalWidth === 0);

    if (timedOut || failedImages.length > 0) {
        failedImages.forEach(image => {
            if (image.src !== fallback) {
                image.onerror = null;
                image.src = fallback;
            }
        });
        await Promise.all(failedImages.map(waitForImage));
    }

    await Promise.all(images.map(async image => {
        if (image.complete && image.naturalWidth > 0 && image.decode) {
            try {
                await image.decode();
            } catch (_) {
                // The image has settled; the fallback above handles failed loads.
            }
        }
    }));
    return true;
}`, nil, playwright.PageWaitForFunctionOptions{Timeout: playwright.Float(12000)}); err != nil {
		log.Println("等待图片和字体加载超时，继续截图:", err)
	}
	page.WaitForTimeout(waitTime)
	locator := page.Locator("#main")
	if v, err := locator.IsVisible(); err != nil || !v {
		log.Println("元素未加载取消截图操作")
		return nil, fmt.Errorf("元素未加载")
	}
	screenshot, err := locator.Screenshot(playwright.LocatorScreenshotOptions{Type: playwright.ScreenshotTypeJpeg})
	if err != nil {
		return nil, fmt.Errorf("截图失败: %w", err)
	}
	log.Println("截图完成...")
	return screenshot, nil
}

func GetImg(url string) []byte {
	var resp *http.Response
	var pic []byte
	times := 0
	for times < 3 {
		resp1, err := http.Get(url)
		resp = resp1
		if err != nil {
			log.Println("获取图片失败", err)
			times++
			continue
		}
		pic, _ = io.ReadAll(resp.Body)
		break
	}
	if resp != nil {
		defer resp.Body.Close()
	}
	return pic
}
