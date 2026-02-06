package bot

import (
	"testing"

	"github.com/stalkerxxl/telegnom/types"
)

//goland:noinspection GoMaybeNil,GoMaybeNil,GoMaybeNil,GoMaybeNil,GoMaybeNil,GoMaybeNil
func TestContext_EffectiveEntities(t *testing.T) {
	// 1. Тест для обычного сообщения
	chatID := int64(123)
	userID := int64(456)
	user := &types.User{ID: userID, FirstName: "Test"}
	chat := &types.Chat{ID: chatID, Type: types.ChatTypePrivate}

	updMsg := &types.Update{
		Message: &types.Message{
			ID:   1,
			Chat: chat,
			From: user,
			Text: "hello",
		},
	}

	ctxMsg := newContext(&Bot{}, updMsg)

	if ctxMsg.EffectiveChat().ID != chatID {
		t.Errorf("expected chat id %d, got %d", chatID, ctxMsg.EffectiveChat().ID)
	}
	if ctxMsg.EffectiveUser().ID != userID {
		t.Errorf("expected user id %d, got %d", userID, ctxMsg.EffectiveUser().ID)
	}

	// 2. Тест для CallbackQuery
	callbackUser := &types.User{ID: 789, FirstName: "CallbackUser"}
	updCallback := &types.Update{
		CallbackQuery: &types.CallbackQuery{
			ID:      "cb_id",
			From:    callbackUser,
			Message: &types.MaybeInaccessibleMessageData{}, // В реальности тут был бы маршалинг, но для теста проверим From
		},
	}
	// Добавим сообщение в CallbackQuery вручную через подмену структуры (имитация)
	// В нашей реализации CallbackQuery.EffectiveChat проверяет Message

	ctxCb := newContext(&Bot{}, updCallback)
	if ctxCb.EffectiveUser().ID != 789 {
		t.Errorf("expected callback user id 789, got %d", ctxCb.EffectiveUser().ID)
	}
}

func TestContext_Actions_Safety(t *testing.T) {
	// Проверяем, что методы не паникуют, если данных нет
	ctxEmpty := newContext(&Bot{}, &types.Update{})

	if msg, err := ctxEmpty.Send("hi"); msg != nil || err != nil {
		t.Error("Send should return nil, nil for empty update")
	}

	if msg, err := ctxEmpty.Reply("hi"); msg != nil || err != nil {
		t.Error("Reply should return nil, nil for empty update")
	}

	if err := ctxEmpty.Delete(); err != nil {
		t.Error("Delete should not return error for empty update")
	}

	if err := ctxEmpty.Answer("ok"); err != nil {
		t.Error("Answer should not return error for empty update")
	}
}
