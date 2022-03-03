package api

import (
	"encoding/json"
	"fmt"
	"github.com/tkchry/nck-trampoline-bot/internal/api/request"
	"github.com/tkchry/nck-trampoline-bot/internal/db"
	"github.com/tkchry/nck-trampoline-bot/internal/domain/model"
	"github.com/tkchry/nck-trampoline-bot/internal/domain/model/member"
	"net/http"
)

func MessageHandler(w http.ResponseWriter, req *http.Request) {
	if method := req.Method; method != "POST" {
		return
	}

	body := make([]byte, req.ContentLength)
	req.Body.Read(body)
	var messageRequest request.MessageRequest
	json.Unmarshal(body, &messageRequest)

	bot := model.GetBot()
	bot.PushMessageForNotifyGroupId(messageRequest.Message)

	w.WriteHeader(200)
}

func SearchMemberHandler(w http.ResponseWriter, req *http.Request) {
	if method := req.Method; method != "POST" {
		return
	}

	body := make([]byte, req.ContentLength)
	req.Body.Read(body)
	var searchMemberRequest request.SearchMemberRequest
	json.Unmarshal(body, &searchMemberRequest)

	searchMemberId := member.NewId(searchMemberRequest.MemberId)
	mem := db.GetMember(searchMemberId)

	var content string
	if mem.IsEmpty() {
		content = fmt.Sprintf("MemberId: %d のメンバーは見つかりませんでした。", searchMemberId.Value)
	} else {
		content = fmt.Sprintf("MemberId: %d は %s さんです", mem.Id.Value, mem.Nickname.Value)
	}

	bot := model.GetBot()
	bot.PushMessageForNotifyGroupId(content)

	w.WriteHeader(200)
}
