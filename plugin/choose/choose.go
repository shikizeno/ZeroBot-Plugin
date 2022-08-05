// Package choose 选择困难症帮手
package choose

import (
	"math/rand"
	"strconv"
	"strings"

	ctrl "github.com/FloatTech/zbpctrl"
	"github.com/FloatTech/zbputils/control"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

func init() {
	engine := control.Register("choose", &ctrl.Options[*zero.Ctx]{
		DisableOnDefault: false,
		Help: "choose\n" +
			"- 可口还是百事\n" +
			"- KFC还是华莱士\n" +
			"- 星星还是陈陈\n",
	})
	engine.OnPrefix("选择").SetBlock(true).Handle(handle)
}
func handle(ctx *zero.Ctx) {
	rawOptions := strings.Split(ctx.State["args"].(string), "还是")
	var options = make([]string, 0)
	for count, option := range rawOptions {
		options = append(options, strconv.Itoa(count+1)+", "+option)
	}
	result := rawOptions[rand.Intn(len(rawOptions))]
	name := ctx.Event.Sender.NickName
	ctx.SendChain(message.Text("> ", name, "\n",
		"你的选项有:", "\n",
		strings.Join(options, "\n"), "\n",
		"你最终会选: ", result,
	))
}
