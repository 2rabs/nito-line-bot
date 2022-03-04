package controller

import (
	"encoding/json"
	"fmt"
	"github.com/tkchry/nck-trampoline-bot/internal/api/request"
	"github.com/tkchry/nck-trampoline-bot/internal/database"
	"github.com/tkchry/nck-trampoline-bot/internal/domain/model/member"
	"github.com/tkchry/nck-trampoline-bot/internal/line"
	"gorm.io/gorm"
	"net/http"
)

type MemberController struct {
	db  *gorm.DB
	bot *line.Bot
}

func NewMemberController(
	db *gorm.DB,
	bot *line.Bot,
) *MemberController {
	return &MemberController{
		db:  db,
		bot: bot,
	}
}

func (c MemberController) MessageHandler(w http.ResponseWriter, req *http.Request) {
	if method := req.Method; method != "POST" {
		return
	}

	body := make([]byte, req.ContentLength)
	req.Body.Read(body)
	var messageRequest request.MessageRequest
	json.Unmarshal(body, &messageRequest)

	c.bot.PushMessageForNotifyGroupId(messageRequest.Message)

	w.WriteHeader(200)
}

func (c MemberController) SearchMemberHandler(w http.ResponseWriter, req *http.Request) {
	if method := req.Method; method != "POST" {
		return
	}

	body := make([]byte, req.ContentLength)
	req.Body.Read(body)
	var searchMemberRequest request.SearchMemberRequest
	json.Unmarshal(body, &searchMemberRequest)

	searchMemberId := member.NewId(searchMemberRequest.MemberId)
	mem := database.GetMember(c.db, searchMemberId)

	var content string
	if mem.IsEmpty() {
		content = fmt.Sprintf("MemberId: %d のメンバーは見つかりませんでした。", searchMemberId.Value)
	} else {
		content = fmt.Sprintf("MemberId: %d は %s さんです", mem.Id.Value, mem.Nickname.Value)
	}

	c.bot.PushMessageForNotifyGroupId(content)

	w.WriteHeader(200)
}
