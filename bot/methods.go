package bot

import (
	"github.com/stalkerxxl/telegnom/types"
)

// GetMe - a simple method for testing your bot's authentication token. Requires
// no parameters. Returns basic information about the bot in form of a types.User
// object.
//
// See https://core.telegram.org/bots/api#getme
func (b *Bot) GetMe() (*types.User, error) {
	r := &types.User{}
	err := b.apiRequest("getMe", nil, r)
	return r, err
}

// LogOut - use this method to log out from the cloud Bot API server before
// launching the bot locally. You must log out the bot before running it locally,
// otherwise there is no guarantee that the bot will receive updates. After a
// successful call, you can immediately log in on a local server, but will not be
// able to log in back to the cloud Bot API server for 10 minutes. Returns True
// on success. Requires no parameters.
//
// See https://core.telegram.org/bots/api#logout
func (b *Bot) LogOut() (bool, error) {
	var r bool
	err := b.apiRequest("logOut", nil, &r)
	return r, err
}

// Close - use this method to close the bot instance before moving it from one
// local server to another. You need to delete the webhook before calling this
// method to ensure that the bot isn't launched again after server restart. The
// method will return error 429 in the first 10 minutes after the bot is
// launched. Returns True on success. Requires no parameters.
//
// See https://core.telegram.org/bots/api#close
func (b *Bot) Close() (bool, error) {
	var r bool
	err := b.apiRequest("close", nil, &r)
	return r, err
}

// SendMessage - use this method to send text messages. On success, the sent
// types.Message is returned.
//
// See https://core.telegram.org/bots/api#sendmessage
func (b *Bot) SendMessage(p *SendMessageParams) (*types.Message, error) {
	r := &types.Message{}
	err := b.apiRequest("sendMessage", p, r)
	return r, err
}

// ForwardMessage - use this method to forward messages of any kind. Service
// messages and messages with protected content can't be forwarded. On success,
// the sent types.Message is returned.
//
// See https://core.telegram.org/bots/api#forwardmessage
func (b *Bot) ForwardMessage(p *ForwardMessageParams) (*types.Message, error) {
	r := &types.Message{}
	err := b.apiRequest("forwardMessage", p, r)
	return r, err
}

// ForwardMessages - use this method to forward multiple messages of any kind. If
// some of the specified messages can't be found or forwarded, they are skipped.
// Service messages and messages with protected content can't be forwarded. Album
// grouping is kept for forwarded messages. On success, an array of types.MessageID of
// the sent messages is returned.
//
// See https://core.telegram.org/bots/api#forwardmessages
func (b *Bot) ForwardMessages(p *ForwardMessagesParams) ([]types.MessageID, error) {
	var r []types.MessageID
	err := b.apiRequest("forwardMessages", p, &r)
	return r, err
}

// CopyMessage - use this method to copy messages of any kind. Service messages,
// paid media messages, giveaway messages, giveaway winners messages, and invoice
// messages can't be copied. A quiz poll can be copied only if the value of the
// field correct_option_id is known to the bot. The method is analogous to the
// method ForwardMessage, but the copied message doesn't have a link to the
// original message. Returns the types.MessageID of the sent message on success.
//
// See https://core.telegram.org/bots/api#copymessage
func (b *Bot) CopyMessage(p *CopyMessageParams) (*types.MessageID, error) {
	r := &types.MessageID{}
	err := b.apiRequest("copyMessage", p, r)
	return r, err
}

// CopyMessages - use this method to copy messages of any kind. If some of the
// specified messages can't be found or copied, they are skipped. Service
// messages, paid media messages, giveaway messages, giveaway winners messages,
// and invoice messages can't be copied. A quiz poll can be copied only if the
// value of the field correct_option_id is known to the bot. The method is
// analogous to the method ForwardMessages, but the copied messages don't have a
// link to the original message. Album grouping is kept for copied messages. On
// success, an array of types.MessageID of the sent messages is returned.
//
// See https://core.telegram.org/bots/api#copymessages
func (b *Bot) CopyMessages(p *CopyMessagesParams) ([]types.MessageID, error) {
	var r []types.MessageID
	err := b.apiRequest("copyMessages", p, &r)
	return r, err
}

// SendPhoto - use this method to send photos. On success, the sent types.Message is
// returned
//
// Пример использования (фото с диска):
//
//	p := &bot.SendPhotoParams{
//	    ChatID: 12345,
//	    Photo: &params.InputFile{Path: "test.png"},
//	    Caption: "Фото через attach://",
//	}
//	bot.SendPhoto(p)
//
// See https://core.telegram.org/bots/api#sendphoto
func (b *Bot) SendPhoto(p *SendPhotoParams) (*types.Message, error) {
	res := &types.Message{}
	err := b.apiRequest("sendPhoto", p, res)
	return res, err
}

// SendAudio - use this method to send audio files, if you want Telegram clients
// to display them in the music player. Your audio must be in the .MP3 or .M4A
// format. On success, the sent types.Message is returned. Bots can currently send
// audio files of up to 50 MB in size, this limit may be changed in the future.
//
// For sending voice messages, use the SendVoice method instead.
//
// See https://core.telegram.org/bots/api#sendaudio
func (b *Bot) SendAudio(p *SendAudioParams) (*types.Message, error) {
	res := &types.Message{}
	err := b.apiRequest("sendAudio", p, res)
	return res, err
}

// SendDocument - use this method to send general files. On success, the sent
// types.Message is returned. Bots can currently send files of any type of up to 50 MB
// in size, this limit may be changed in the future.
//
// See https://core.telegram.org/bots/api#senddocument
func (b *Bot) SendDocument(p *SendDocumentParams) (*types.Message, error) {
	res := &types.Message{}
	err := b.apiRequest("sendDocument", p, res)
	return res, err
}

// SendVideo - use this method to send video files. Telegram clients support
// MPEG4 videos (other formats may be sent as types.Document). On success, the sent
// types.Message is returned. Bots can currently send video files of up to 50 MB in
// size, this limit may be changed in the future.
//
// See https://core.telegram.org/bots/api#sendvideo
func (b *Bot) SendVideo(p *SendVideoParams) (*types.Message, error) {
	res := &types.Message{}
	err := b.apiRequest("sendVideo", p, res)
	return res, err
}

// SendAnimation - use this method to send animation files (GIF or H.264/MPEG-4 AVC
// video without sound). On success, the sent types.Message is returned. Bots can
// currently send animation files of up to 50 MB in size, this limit may be
// changed in the future.
//
// See https://core.telegram.org/bots/api#sendanimation
func (b *Bot) SendAnimation(p *SendAnimationParams) (*types.Message, error) {
	res := &types.Message{}
	err := b.apiRequest("sendAnimation", p, res)
	return res, err
}

// SendVoice - use this method to send audio files, if you want Telegram clients
// to display the file as a playable voice message. For this to work, your audio
// must be in an .OGG file encoded with OPUS, or in .MP3 format, or in .M4A
// format (other formats may be sent as [types.Audio] or [types.Document]). On
// success, the sent types.Message is returned. Bots can currently send voice
// messages of up to 50 MB in size, this limit may be changed in the future.
//
// See https://core.telegram.org/bots/api#sendvoice
func (b *Bot) SendVoice(p *SendVoiceParams) (*types.Message, error) {
	res := &types.Message{}
	err := b.apiRequest("sendVoice", p, res)
	return res, err
}

// SendVideoNote - As of v.4.0, Telegram clients support rounded square MPEG4
// videos of up to 1 minute long, use this method to send video messages. On
// success, the sent [types.Message] is returned.
//
// See https://core.telegram.org/bots/api#sendvideonote
func (b *Bot) SendVideoNote(p *SendVideoNoteParams) (*types.Message, error) {
	res := &types.Message{}
	err := b.apiRequest("sendVideoNote", p, res)
	return res, err
}

// SendMediaGroup - use this method to send a group of photos, videos, documents or
// audios as an album. Documents and audio files can be only grouped in an album
// with messages of the same type. On success, an array of types.Message that were sent is
// returned.
//
// See https://core.telegram.org/bots/api#sendmediagroup
func (b *Bot) SendMediaGroup(p *SendMediaGroupParams) ([]types.Message, error) {
	var res []types.Message
	err := b.apiRequest("sendMediaGroup", p, &res)
	return res, err
}

// SendStory - use this method to send a story. On success, the sent
// types.Message is returned.
//
// See https://core.telegram.org/bots/api#sendstory
func (b *Bot) SendStory(p *SendStoryParams) (*types.Message, error) {
	res := &types.Message{}
	err := b.apiRequest("sendStory", p, res)
	return res, err
}

// SendLocation - use this method to send point on the map. On success, the sent
// types.Message is returned.
//
// See https://core.telegram.org/bots/api#sendlocation
func (b *Bot) SendLocation(p *SendLocationParams) (*types.Message, error) {
	res := &types.Message{}
	err := b.apiRequest("sendLocation", p, res)
	return res, err
}

// SendVenue - Use this method to send information about a venue. On success, the
// sent types.Message is returned
//
// See https://core.telegram.org/bots/api#sendvenue
func (b *Bot) SendVenue(p *SendVenueParams) (*types.Message, error) {
	res := &types.Message{}
	err := b.apiRequest("sendVenue", p, res)
	return res, err
}

// SendContact - use this method to send phone contacts. On success, the sent
// types.Message is returned.
//
// See https://core.telegram.org/bots/api#sendcontact
func (b *Bot) SendContact(p *SendContactParams) (*types.Message, error) {
	res := &types.Message{}
	err := b.apiRequest("sendContact", p, res)
	return res, err
}

// SendPoll - use this method to send a native poll. On success, the sent
// types.Message is returned.
//
// See https://core.telegram.org/bots/api#sendpoll
func (b *Bot) SendPoll(p *SendPollParams) (*types.Message, error) {
	res := &types.Message{}
	err := b.apiRequest("sendPoll", p, res)
	return res, err
}

// SendChecklist - use this method to send a checklist on behalf of a connected
// business account. On success, the sent types.Message is returned.
//
// See https://core.telegram.org/bots/api#sendchecklist
func (b *Bot) SendChecklist(p *SendChecklistParams) (*types.Message, error) {
	res := &types.Message{}
	err := b.apiRequest("sendChecklist", p, res)
	return res, err
}

// SendDice - use this method to send an animated emoji that will display a
// random value. On success, the sent types.Message is returned.
//
// See https://core.telegram.org/bots/api#senddice
func (b *Bot) SendDice(p *SendDiceParams) (*types.Message, error) {
	res := &types.Message{}
	err := b.apiRequest("sendDice", p, res)
	return res, err
}

// SendMessageDraft - use this method to stream a partial message to a user while
// the message is being generated; supported only for bots with forum topic mode
// enabled. Returns True on success.
//
// See https://core.telegram.org/bots/api#sendmessagedraft
func (b *Bot) SendMessageDraft(p *SendMessageDraftParams) (bool, error) {
	var r bool
	err := b.apiRequest("sendMessageDraft", p, &r)
	return r, err
}

// SendChatAction - use this method when you need to tell the user that something
// is happening on the bot's side. The status is set for 5 seconds or less (when
// a message arrives from your bot, Telegram clients clear its typing status).
// Returns True on success.
//
// We only recommend using this method when a response from the bot will take a
// noticeable amount of time to arrive.
//
// See https://core.telegram.org/bots/api#sendchataction
func (b *Bot) SendChatAction(p *SendChatActionParams) (bool, error) {
	var r bool
	err := b.apiRequest("sendChatAction", p, &r)
	return r, err
}

// SetMessageReaction - use this method to change the chosen reactions on a
// message. Service messages of some types can't be reacted to. Automatically
// forwarded messages from a channel to its discussion group have the same
// available reactions as messages in the channel. Bots can't use paid reactions.
// Returns True on success.
//
// See https://core.telegram.org/bots/api#setmessagereaction
func (b *Bot) SetMessageReaction(p *SetMessageReactionParams) (bool, error) {
	var r bool
	err := b.apiRequest("setMessageReaction", p, &r)
	return r, err
}

// GetUserProfilePhotos - use this method to get a list of profile pictures for a
// user. Returns a types.UserProfilePhotos object.
//
// See https://core.telegram.org/bots/api#getuserprofilephotos
func (b *Bot) GetUserProfilePhotos(p *GetUserProfilePhotosParams) (*types.UserProfilePhotos, error) {
	r := &types.UserProfilePhotos{}
	err := b.apiRequest("getUserProfilePhotos", p, r)
	return r, err
}

// SetUserEmojiStatus changes the emoji status for a given user that previously
// allowed the bot to manage their emoji status via the Mini App method
// RequestEmojiStatusAccess. Returns True on success.
//
// See https://core.telegram.org/bots/api#setuseremojistatus
func (b *Bot) SetUserEmojiStatus(p *SetUserEmojiStatusParams) (bool, error) {
	var r bool
	err := b.apiRequest("setUserEmojiStatus", p, &r)
	return r, err
}

// GetFile use this method to get basic information about a file and prepare it
// for downloading. For the moment, bots can download files of up to 20MB in
// size. On success, a types.File object is returned. The file can then be downloaded
// via the link
//
//	https://api.telegram.org/file/bot<token>/<file_path>
//
// where <file_path> is taken from the response. It is guaranteed that the link will be
// valid for at least 1 hour. When the link expires, a new one can be requested
// by calling getFile again.
//
// Note: This function may not preserve the original file name and MIME type. You
// should save the file's MIME type and name (if available) when the File object
// is received.
//
// See https://core.telegram.org/bots/api#getfile
func (b *Bot) GetFile(p *GetFileParams) (*types.File, error) {
	r := &types.File{}
	err := b.apiRequest("getFile", p, r)
	return r, err
}

// BanChatMember - use this method to ban a user in a group, a supergroup or a
// channel. In the case of supergroups and channels, the user will not be able to
// return to the chat on their own using invite links, etc., unless unbanned
// first. The bot must be an administrator in the chat for this to work and must
// have the appropriate administrator rights. Returns True on success.
//
// See https://core.telegram.org/bots/api#banchatmember
func (b *Bot) BanChatMember(p *BanChatMemberParams) (bool, error) {
	var r bool
	err := b.apiRequest("banChatMember", p, &r)
	return r, err
}

// UnbanChatMember - use this method to unban a previously banned user in a
// supergroup or channel. The user will not return to the group or channel
// automatically, but will be able to join via link, etc. The bot must be an
// administrator for this to work. By default, this method guarantees that after
// the call the user is not a member of the chat, but will be able to join it. So
// if the user is a member of the chat they will also be removed from the chat.
// If you don't want this, use the parameter only_if_banned. Returns True on
// success.
//
// See https://core.telegram.org/bots/api#unbanchatmember
func (b *Bot) UnbanChatMember(p *UnbanChatMemberParams) (bool, error) {
	var r bool
	err := b.apiRequest("unbanChatMember", p, &r)
	return r, err
}

// RestrictChatMember - use this method to restrict a user in a supergroup. The
// bot must be an administrator in the supergroup for this to work and must have
// the appropriate administrator rights. Pass True for all permissions to lift
// restrictions from a user. Returns True on success.
//
// See https://core.telegram.org/bots/api#restrictchatmember
func (b *Bot) RestrictChatMember(p *RestrictChatMemberParams) (bool, error) {
	var r bool
	err := b.apiRequest("restrictChatMember", p, &r)
	return r, err
}

// PromoteChatMember - use this method to promote or demote a user in a
// supergroup or a channel. The bot must be an administrator in the chat for this
// to work and must have the appropriate administrator rights. Pass False for all
// boolean parameters to demote a user. Returns True on success.
//
// See https://core.telegram.org/bots/api#promotechatmember
func (b *Bot) PromoteChatMember(p *PromoteChatMemberParams) (bool, error) {
	var r bool
	err := b.apiRequest("promoteChatMember", p, &r)
	return r, err
}

// SetChatAdministratorCustomTitle - use this method to set a custom title for an
// administrator in a supergroup promoted by the bot. Returns True on success.
//
// See https://core.telegram.org/bots/api#setchatadministratorcustomtitle
func (b *Bot) SetChatAdministratorCustomTitle(p *SetChatAdministratorCustomTitleParams) (bool, error) {
	var r bool
	err := b.apiRequest("setChatAdministratorCustomTitle", p, &r)
	return r, err
}

// BanChatSenderChat - use this method to ban a channel chat in a supergroup or a
// channel. Until the chat is unbanned, the owner of the banned chat won't be
// able to send messages on behalf of any of their channels. The bot must be an
// administrator in the supergroup or channel for this to work and must have the
// appropriate administrator rights. Returns True on success.
//
// See https://core.telegram.org/bots/api#banchatsenderchat
func (b *Bot) BanChatSenderChat(p *BanChatSenderChatParams) (bool, error) {
	var r bool
	err := b.apiRequest("banChatSenderChat", p, &r)
	return r, err
}

// UnbanChatSenderChat - use this method to unban a previously banned channel
// chat in a supergroup or channel. The bot must be an administrator for this to
// work and must have the appropriate administrator rights. Returns True on
// success.
//
// See https://core.telegram.org/bots/api#unbanchatsenderchat
func (b *Bot) UnbanChatSenderChat(p *UnbanChatSenderChatParams) (bool, error) {
	var r bool
	err := b.apiRequest("unbanChatSenderChat", p, &r)
	return r, err
}

// SetChatPermissions - use this method to set default chat permissions for all
// members. The bot must be an administrator in the group or a supergroup for
// this to work and must have the can_restrict_members administrator rights.
// Returns True on success.
//
// See https://core.telegram.org/bots/api#setchatpermissions
func (b *Bot) SetChatPermissions(p *SetChatPermissionsParams) (bool, error) {
	var r bool
	err := b.apiRequest("setChatPermissions", p, &r)
	return r, err
}

// ExportChatInviteLink - use this method to generate a new primary invite link
// for a chat; any previously generated primary link is revoked. The bot must be
// an administrator in the chat for this to work and must have the appropriate
// administrator rights. Returns the new invite link as String on success.
//
// Note: Each administrator in a chat generates their own invite links. Bots
// can't use invite links generated by other administrators. If you want your bot
// to work with invite links, it will need to generate its own link using
// ExportChatInviteLink or by calling the GetChat method. If your bot needs to
// generate a new primary invite link replacing its previous one, use
// ExportChatInviteLink again.
//
// See https://core.telegram.org/bots/api#exportchatinvitelink
func (b *Bot) ExportChatInviteLink(p *ExportChatInviteLinkParams) (string, error) {
	var r string
	err := b.apiRequest("exportChatInviteLink", p, &r)
	return r, err
}

// CreateChatInviteLink - use this method to create an additional invite link for
// a chat. The bot must be an administrator in the chat for this to work and must
// have the appropriate administrator rights. The link can be revoked using the
// method RevokeChatInviteLink. Returns the new invite link as types.ChatInviteLink
// object.
//
// See https://core.telegram.org/bots/api#createchatinvitelink
func (b *Bot) CreateChatInviteLink(p *CreateChatInviteLinkParams) (*types.ChatInviteLink, error) {
	r := &types.ChatInviteLink{}
	err := b.apiRequest("createChatInviteLink", p, r)
	return r, err
}

// EditChatInviteLink - use this method to edit a non-primary invite link created
// by the bot. The bot must be an administrator in the chat for this to work and
// must have the appropriate administrator rights. Returns the edited invite link
// as a types.ChatInviteLink object.
//
// See https://core.telegram.org/bots/api#editchatinvitelink
func (b *Bot) EditChatInviteLink(p *EditChatInviteLinkParams) (*types.ChatInviteLink, error) {
	r := &types.ChatInviteLink{}
	err := b.apiRequest("editChatInviteLink", p, r)
	return r, err
}

// CreateChatSubscriptionInviteLink - use this method to create a subscription
// invite link for a channel chat. The bot must have the can_invite_users
// administrator rights. The link can be edited using the method
// EditChatSubscriptionInviteLink or revoked using the method
// RevokeChatInviteLink. Returns the new invite link as a types.ChatInviteLink object.
//
// See https://core.telegram.org/bots/api#createchatsubscriptioninvitelink
func (b *Bot) CreateChatSubscriptionInviteLink(p *CreateChatSubscriptionInviteLinkParams) (*types.ChatInviteLink, error) {
	r := &types.ChatInviteLink{}
	err := b.apiRequest("createChatSubscriptionInviteLink", p, r)
	return r, err
}

// EditChatSubscriptionInviteLink - use this method to edit a subscription invite
// link created by the bot. The bot must have the can_invite_users administrator
// rights. Returns the edited invite link as a types.ChatInviteLink object.
//
// See https://core.telegram.org/bots/api#editchatsubscriptioninvitelink
func (b *Bot) EditChatSubscriptionInviteLink(p *EditChatSubscriptionInviteLinkParams) (*types.ChatInviteLink, error) {
	r := &types.ChatInviteLink{}
	err := b.apiRequest("editChatSubscriptionInviteLink", p, r)
	return r, err
}

// RevokeChatInviteLink - use this method to revoke an invite link created by the
// bot. If the primary link is revoked, a new link is automatically generated.
// The bot must be an administrator in the chat for this to work and must have
// the appropriate administrator rights. Returns the revoked invite link as
// types.ChatInviteLink object.
//
// See https://core.telegram.org/bots/api#revokechatinvitelink
func (b *Bot) RevokeChatInviteLink(p *RevokeChatInviteLinkParams) (*types.ChatInviteLink, error) {
	r := &types.ChatInviteLink{}
	err := b.apiRequest("revokeChatInviteLink", p, r)
	return r, err
}

// ApproveChatJoinRequest - use this method to approve a chat join request. The
// bot must be an administrator in the chat for this to work and must have the
// can_invite_users administrator right. Returns True on success.
//
// See https://core.telegram.org/bots/api#approvechatjoinrequest
func (b *Bot) ApproveChatJoinRequest(p *ApproveChatJoinRequestParams) (bool, error) {
	var r bool
	err := b.apiRequest("approveChatJoinRequest", p, &r)
	return r, err
}

// DeclineChatJoinRequest - use this method to decline a chat join request. The
// bot must be an administrator in the chat for this to work and must have the
// can_invite_users administrator right. Returns True on success.
//
// See https://core.telegram.org/bots/api#declinechatjoinrequest
func (b *Bot) DeclineChatJoinRequest(p *DeclineChatJoinRequestParams) (bool, error) {
	var r bool
	err := b.apiRequest("declineChatJoinRequest", p, &r)
	return r, err
}

// SetChatPhoto - use this method to set a new profile photo for the chat. Photos
// can't be changed for private chats. The bot must be an administrator in the
// chat for this to work and must have the appropriate administrator rights.
// Returns True on success.
//
// See https://core.telegram.org/bots/api#setchatphoto
func (b *Bot) SetChatPhoto(p *SetChatPhotoParams) (bool, error) {
	var r bool
	err := b.apiRequest("setChatPhoto", p, &r)
	return r, err
}

// DeleteChatPhoto - use this method to delete a chat photo. Photos can't be
// changed for private chats. The bot must be an administrator in the chat for
// this to work and must have the appropriate administrator rights. Returns True
// on success.
//
// See https://core.telegram.org/bots/api#deletechatphoto
func (b *Bot) DeleteChatPhoto(p *DeleteChatPhotoParams) (bool, error) {
	var r bool
	err := b.apiRequest("deleteChatPhoto", p, &r)
	return r, err
}

// SetChatTitle - use this method to change the title of a chat. Titles can't be
// changed for private chats. The bot must be an administrator in the chat for
// this to work and must have the appropriate administrator rights. Returns True
// on success.
//
// See https://core.telegram.org/bots/api#setchattitle
func (b *Bot) SetChatTitle(p *SetChatTitleParams) (bool, error) {
	var r bool
	err := b.apiRequest("setChatTitle", p, &r)
	return r, err
}

// SetChatDescription - use this method to change the description of a group, a
// supergroup or a channel. The bot must be an administrator in the chat for this
// to work and must have the appropriate administrator rights. Returns True on
// success.
//
// See https://core.telegram.org/bots/api#setchatdescription
func (b *Bot) SetChatDescription(p *SetChatDescriptionParams) (bool, error) {
	var r bool
	err := b.apiRequest("setChatDescription", p, &r)
	return r, err
}

// PinChatMessage - use this method to add a message to the list of pinned
// messages in a chat. In private chats and channel direct messages chats, all
// non-service messages can be pinned. Conversely, the bot must be an
// administrator with the 'can_pin_messages' right or the 'can_edit_messages'
// right to pin messages in groups and channels respectively. Returns True on
// success.
//
// See https://core.telegram.org/bots/api#pinchatmessage
func (b *Bot) PinChatMessage(p *PinChatMessageParams) (bool, error) {
	var r bool
	err := b.apiRequest("pinChatMessage", p, &r)
	return r, err
}

// UnpinChatMessage - use this method to remove a message from the list of pinned
// messages in a chat. In private chats and channel direct messages chats, all
// messages can be unpinned. Conversely, the bot must be an administrator with
// the 'can_pin_messages' right or the 'can_edit_messages' right to unpin
// messages in groups and channels respectively. Returns True on success.
//
// See https://core.telegram.org/bots/api#unpinchatmessage
func (b *Bot) UnpinChatMessage(p *UnpinChatMessageParams) (bool, error) {
	var r bool
	err := b.apiRequest("unpinChatMessage", p, &r)
	return r, err
}

// UnpinAllChatMessages - use this method to clear the list of pinned messages in
// a chat. In private chats and channel direct messages chats, no additional
// rights are required to unpin all pinned messages. Conversely, the bot must be
// an administrator with the 'can_pin_messages' right or the 'can_edit_messages'
// right to unpin all pinned messages in groups and channels respectively.
// Returns True on success.
//
// See https://core.telegram.org/bots/api#unpinallchatmessages
func (b *Bot) UnpinAllChatMessages(p *UnpinAllChatMessagesParams) (bool, error) {
	var r bool
	err := b.apiRequest("unpinAllChatMessages", p, &r)
	return r, err
}

// LeaveChat - use this method for your bot to leave a group, supergroup or
// channel. Returns True on success.
//
// See https://core.telegram.org/bots/api#leavechat
func (b *Bot) LeaveChat(p *LeaveChatParams) (bool, error) {
	var r bool
	err := b.apiRequest("leaveChat", p, &r)
	return r, err
}

// GetChat - use this method to get up-to-date information about the chat.
// Returns a types.ChatFullInfo object on success.
//
// See https://core.telegram.org/bots/api#getchat
func (b *Bot) GetChat(p *GetChatParams) (*types.ChatFullInfo, error) {
	r := &types.ChatFullInfo{}
	err := b.apiRequest("getChat", p, r)
	return r, err
}

// GetChatAdministrators - use this method to get a list of administrators in a
// chat, which aren't bots. Returns an Array of types.ChatMemberData objects.
//
// See https://core.telegram.org/bots/api#getchatadministrators
func (b *Bot) GetChatAdministrators(p *GetChatAdministratorsParams) ([]types.ChatMemberData, error) {
	var r []types.ChatMemberData
	err := b.apiRequest("getChatAdministrators", p, &r)
	return r, err
}

// GetChatMemberCount - use this method to get the number of members in a chat.
// Returns Int on success.
//
// See https://core.telegram.org/bots/api#getchatmembercount
func (b *Bot) GetChatMemberCount(p *GetChatMemberCountParams) (int, error) {
	var r int
	err := b.apiRequest("getChatMemberCount", p, &r)
	return r, err
}

// GetChatMember - use this method to get information about a member of a chat.
// The method is only guaranteed to work for other users if the bot is an
// administrator in the chat. Returns a types.ChatMemberData object on success
//
// See https://core.telegram.org/bots/api#getchatmember
func (b *Bot) GetChatMember(p *GetChatMemberParams) (*types.ChatMemberData, error) {
	r := &types.ChatMemberData{}
	err := b.apiRequest("getChatMember", p, r)
	return r, err
}

// SetChatStickerSet - use this method to set a new group sticker set for a
// supergroup. The bot must be an administrator in the chat for this to work and
// must have the appropriate administrator rights. use the field
// can_set_sticker_set optionally returned in GetChat requests to check if the
// bot can use this method. Returns True on success.
//
// See https://core.telegram.org/bots/api#setchatstickerset
func (b *Bot) SetChatStickerSet(p *SetChatStickerSetParams) (bool, error) {
	var r bool
	err := b.apiRequest("setChatStickerSet", p, &r)
	return r, err
}

// DeleteChatStickerSet - use this method to delete a group sticker set from a
// supergroup. The bot must be an administrator in the chat for this to work and
// must have the appropriate administrator rights. use the field
// can_set_sticker_set optionally returned in GetChat requests to check if the
// bot can use this method. Returns True on success.
//
// See https://core.telegram.org/bots/api#deletechatstickerset
func (b *Bot) DeleteChatStickerSet(p *DeleteChatStickerSetParams) (bool, error) {
	var r bool
	err := b.apiRequest("deleteChatStickerSet", p, &r)
	return r, err
}

// GetForumTopicIconStickers - use this method to get custom emoji stickers,
// which can be used as a forum topic icon by any user. Requires no parameters.
// Returns an Array of types.Sticker objects.
//
// See https://core.telegram.org/bots/api#getforumtopiciconstickers
func (b *Bot) GetForumTopicIconStickers() ([]types.Sticker, error) {
	var r []types.Sticker
	err := b.apiRequest("getForumTopicIconStickers", nil, &r)
	return r, err
}

// CreateForumTopic - use this method to create a topic in a forum supergroup
// chat. The bot must be an administrator in the chat for this to work and must
// have the can_manage_topics administrator rights. Returns information about the
// created topic as a types.ForumTopic object.
//
// See https://core.telegram.org/bots/api#createforumtopic
func (b *Bot) CreateForumTopic(p *CreateForumTopicParams) (*types.ForumTopic, error) {
	r := &types.ForumTopic{}
	err := b.apiRequest("createForumTopic", p, r)
	return r, err
}

// EditForumTopic - use this method to edit name and icon of a topic in a forum
// supergroup chat or a private chat with a user. In the case of a supergroup
// chat the bot must be an administrator in the chat for this to work and must
// have the can_manage_topics administrator rights, unless it is the creator of
// the topic. Returns True on success.
//
// See https://core.telegram.org/bots/api#editforumtopic
func (b *Bot) EditForumTopic(p *EditForumTopicParams) (bool, error) {
	var r bool
	err := b.apiRequest("editForumTopic", p, &r)
	return r, err
}

// CloseForumTopic - use this method to close an open topic in a forum supergroup
// chat. The bot must be an administrator in the chat for this to work and must
// have the can_manage_topics administrator rights, unless it is the creator of
// the topic. Returns True on success.
//
// See https://core.telegram.org/bots/api#closeforumtopic
func (b *Bot) CloseForumTopic(p *CloseForumTopicParams) (bool, error) {
	var r bool
	err := b.apiRequest("closeForumTopic", p, &r)
	return r, err
}

// ReopenForumTopic - use this method to reopen a closed topic in a forum
// supergroup chat. The bot must be an administrator in the chat for this to work
// and must have the can_manage_topics administrator rights, unless it is the
// creator of the topic. Returns True on success.
//
// See https://core.telegram.org/bots/api#reopenforumtopic
func (b *Bot) ReopenForumTopic(p *ReopenForumTopicParams) (bool, error) {
	var r bool
	err := b.apiRequest("reopenForumTopic", p, &r)
	return r, err
}

// DeleteForumTopic - use this method to delete a forum topic along with all its
// messages in a forum supergroup chat or a private chat with a user. In the case
// of a supergroup chat the bot must be an administrator in the chat for this to
// work and must have the can_delete_messages administrator rights. Returns True
// on success.
//
// See https://core.telegram.org/bots/api#deleteforumtopic
func (b *Bot) DeleteForumTopic(p *DeleteForumTopicParams) (bool, error) {
	var r bool
	err := b.apiRequest("deleteForumTopic", p, &r)
	return r, err
}

// UnpinAllForumTopicMessages - use this method to clear the list of pinned
// messages in a forum topic in a forum supergroup chat or a private chat with a
// user. In the case of a supergroup chat the bot must be an administrator in the
// chat for this to work and must have the can_pin_messages administrator right
// in the supergroup. Returns True on success.
//
// See https://core.telegram.org/bots/api#unpinallforumtopicmessages
func (b *Bot) UnpinAllForumTopicMessages(p *UnpinAllForumTopicMessagesParams) (bool, error) {
	var r bool
	err := b.apiRequest("unpinAllForumTopicMessages", p, &r)
	return r, err
}

// EditGeneralForumTopic - use this method to edit the name of the 'General'
// topic in a forum supergroup chat. The bot must be an administrator in the chat
// for this to work and must have the can_manage_topics administrator rights.
// Returns True on success.
//
// See https://core.telegram.org/bots/api#editgeneralforumtopic
func (b *Bot) EditGeneralForumTopic(p *EditGeneralForumTopicParams) (bool, error) {
	var r bool
	err := b.apiRequest("editGeneralForumTopic", p, &r)
	return r, err
}

// CloseGeneralForumTopic - use this method to close an open 'General' topic in a
// forum supergroup chat. The bot must be an administrator in the chat for this
// to work and must have the can_manage_topics administrator rights. Returns True
// on success.
//
// See https://core.telegram.org/bots/api#closegeneralforumtopic
func (b *Bot) CloseGeneralForumTopic(p *CloseGeneralForumTopicParams) (bool, error) {
	var r bool
	err := b.apiRequest("closeGeneralForumTopic", p, &r)
	return r, err
}

// ReopenGeneralForumTopic - use this method to reopen a closed 'General' topic
// in a forum supergroup chat. The bot must be an administrator in the chat for
// this to work and must have the can_manage_topics administrator rights. The
// topic will be automatically unhidden if it was hidden. Returns True on
// success.
//
// See https://core.telegram.org/bots/api#reopengeneralforumtopic
func (b *Bot) ReopenGeneralForumTopic(p *ReopenGeneralForumTopicParams) (bool, error) {
	var r bool
	err := b.apiRequest("reopenGeneralForumTopic", p, &r)
	return r, err
}

// HideGeneralForumTopic - use this method to hide the 'General' topic in a forum
// supergroup chat. The bot must be an administrator in the chat for this to work
// and must have the can_manage_topics administrator rights. The topic will be
// automatically closed if it was open. Returns True on success.
//
// See https://core.telegram.org/bots/api#hidegeneralforumtopic
func (b *Bot) HideGeneralForumTopic(p *HideGeneralForumTopicParams) (bool, error) {
	var r bool
	err := b.apiRequest("hideGeneralForumTopic", p, &r)
	return r, err
}

// UnhideGeneralForumTopic - use this method to unhide the 'General' topic in a
// forum supergroup chat. The bot must be an administrator in the chat for this
// to work and must have the can_manage_topics administrator rights. Returns True
// on success.
//
// See https://core.telegram.org/bots/api#unhidegeneralforumtopic
func (b *Bot) UnhideGeneralForumTopic(p *UnhideGeneralForumTopicParams) (bool, error) {
	var r bool
	err := b.apiRequest("unhideGeneralForumTopic", p, &r)
	return r, err
}

// UnpinAllGeneralForumTopicMessages - use this method to clear the list of
// pinned messages in a General forum topic. The bot must be an administrator in
// the chat for this to work and must have the can_pin_messages administrator
// right in the supergroup. Returns True on success.
//
// See https://core.telegram.org/bots/api#unpinallgeneralforumtopicmessages
func (b *Bot) UnpinAllGeneralForumTopicMessages(p *UnpinAllGeneralForumTopicMessagesParams) (bool, error) {
	var r bool
	err := b.apiRequest("unpinAllGeneralForumTopicMessages", p, &r)
	return r, err
}

// AnswerCallbackQuery - use this method to send answers to callback queries sent
// from inline keyboards. The answer will be displayed to the user as a
// notification at the top of the chat screen or as an alert. On success, True is
// returned.
//
// Alternatively, the user can be redirected to the specified Game URL. For this
// option to work, you must first create a game for your bot via @BotFather and
// accept the terms. Otherwise, you may use links like t.me/your_bot?start=XXXX
// that open your bot with a parameter.
//
// See https://core.telegram.org/bots/api#answercallbackquery
func (b *Bot) AnswerCallbackQuery(p *AnswerCallbackQueryParams) (bool, error) {
	var r bool
	err := b.apiRequest("answerCallbackQuery", p, &r)
	return r, err
}

// GetUserChatBoosts - use this method to get the list of boosts added to a chat
// by a user. Requires administrator rights in the chat. Returns a types.UserChatBoosts
// object.
//
// See https://core.telegram.org/bots/api#getuserchatboosts
func (b *Bot) GetUserChatBoosts(p *GetUserChatBoostsParams) (*types.UserChatBoosts, error) {
	r := &types.UserChatBoosts{}
	err := b.apiRequest("getUserChatBoosts", p, r)
	return r, err
}

// GetBusinessConnection - use this method to get information about the
// connection of the bot with a business account. Returns a types.BusinessConnection
// object on success.
//
// See https://core.telegram.org/bots/api#getbusinessconnection
func (b *Bot) GetBusinessConnection(p *GetBusinessConnectionParams) (*types.BusinessConnection, error) {
	r := &types.BusinessConnection{}
	err := b.apiRequest("getBusinessConnection", p, r)
	return r, err
}

// SetMyCommands - use this method to change the list of the bot's commands. See
// this manual for more details about bot commands. Returns True on success.
//
// See https://core.telegram.org/bots/api#setmycommands
func (b *Bot) SetMyCommands(p *SetMyCommandsParams) (bool, error) {
	var r bool
	err := b.apiRequest("setMyCommands", p, &r)
	return r, err
}

// DeleteMyCommands - use this method to delete the list of the bot's commands
// for the given scope and user language. After deletion, higher level commands
// will be shown to affected users. Returns True on success.
//
// See https://core.telegram.org/bots/api#deletemycommands
func (b *Bot) DeleteMyCommands(p *DeleteMyCommandsParams) (bool, error) {
	var r bool
	err := b.apiRequest("deleteMyCommands", p, &r)
	return r, err
}

// GetMyCommands - use this method to get the current list of the bot's commands
// for the given scope and user language. Returns an Array of types.BotCommand objects.
// If commands aren't set, an empty list is returned.
//
// See https://core.telegram.org/bots/api#getmycommands
func (b *Bot) GetMyCommands(p *GetMyCommandsParams) ([]types.BotCommand, error) {
	var r []types.BotCommand
	err := b.apiRequest("getMyCommands", p, &r)
	return r, err
}

// SetMyName - use this method to change the bot's name. Returns True on success.
//
// See https://core.telegram.org/bots/api#setmyname
func (b *Bot) SetMyName(p *SetMyNameParams) (bool, error) {
	var r bool
	err := b.apiRequest("setMyName", p, &r)
	return r, err
}

// GetMyName - use this method to get the current bot name for the given user
// language. Returns types.BotName on success.
func (b *Bot) GetMyName(p *GetMyNameParams) (*types.BotName, error) {
	r := &types.BotName{}
	err := b.apiRequest("getMyName", p, r)
	return r, err
}

// SetMyDescription - use this method to change the bot's description, which is
// shown in the chat with the bot if the chat is empty. Returns True on success.
//
// See https://core.telegram.org/bots/api#setmydescription
func (b *Bot) SetMyDescription(p *SetMyDescriptionParams) (bool, error) {
	var r bool
	err := b.apiRequest("setMyDescription", p, &r)
	return r, err
}

// GetMyDescription - use this method to get the current bot description for the
// given user language. Returns types.BotDescription on success.
func (b *Bot) GetMyDescription(p *GetMyDescriptionParams) (*types.BotDescription, error) {
	r := &types.BotDescription{}
	err := b.apiRequest("getMyDescription", p, r)
	return r, err
}

// SetMyShortDescription - use this method to change the bot's short description,
// which is shown on the bot's profile page and is sent together with the link
// when users share the bot. Returns True on success.
//
// See https://core.telegram.org/bots/api#getmyshortdescription
func (b *Bot) SetMyShortDescription(p *SetMyShortDescriptionParams) (bool, error) {
	var r bool
	err := b.apiRequest("setMyShortDescription", p, &r)
	return r, err
}

// GetMyShortDescription - use this method to get the current bot short
// description for the given user language. Returns types.BotShortDescription on
// success.
//
// See https://core.telegram.org/bots/api#getmyshortdescription
func (b *Bot) GetMyShortDescription(p *GetMyShortDescriptionParams) (*types.BotShortDescription, error) {
	r := &types.BotShortDescription{}
	err := b.apiRequest("getMyShortDescription", p, r)
	return r, err
}

// SetChatMenuButton - use this method to change the bot's menu button in a
// private chat, or the default menu button. Returns True on success.
//
// See https://core.telegram.org/bots/api#setchatmenubutton
func (b *Bot) SetChatMenuButton(p *SetChatMenuButtonParams) (bool, error) {
	var r bool
	err := b.apiRequest("setChatMenuButton", p, &r)
	return r, err
}

// GetChatMenuButton - use this method to get the current value of the bot's menu
// button in a private chat, or the default menu button. Returns types.MenuButtonData on
// success.
//
// See https://core.telegram.org/bots/api#getchatmenubutton
func (b *Bot) GetChatMenuButton(p *GetChatMenuButtonParams) (*types.MenuButtonData, error) {
	r := &types.MenuButtonData{}
	err := b.apiRequest("getChatMenuButton", p, r)
	return r, err
}

// SetMyDefaultAdministratorRights - use this method to change the default
// administrator rights requested by the bot when it's added as an administrator
// to groups or channels. These rights will be suggested to users, but they are
// free to modify the list before adding the bot. Returns True on success.
//
// See https://core.telegram.org/bots/api#setmydefaultadministratorrights
func (b *Bot) SetMyDefaultAdministratorRights(p *SetMyDefaultAdministratorRightsParams) (bool, error) {
	var r bool
	err := b.apiRequest("setMyDefaultAdministratorRights", p, &r)
	return r, err
}

// GetMyDefaultAdministratorRights - use this method to get the current default
// administrator rights of the bot. Returns types.ChatAdministratorRights on success.
//
// See https://core.telegram.org/bots/api#getmydefaultadministratorrights
func (b *Bot) GetMyDefaultAdministratorRights(p *GetMyDefaultAdministratorRightsParams) (*types.ChatAdministratorRights, error) {
	r := &types.ChatAdministratorRights{}
	err := b.apiRequest("getMyDefaultAdministratorRights", p, r)
	return r, err
}

// GetAvailableGifts returns the list of gifts that can be sent by the bot to
// users and channel chats. Requires no parameters. Returns a types.Gifts object.
//
// See https://core.telegram.org/bots/api#getavailablegifts
func (b *Bot) GetAvailableGifts() (*types.Gifts, error) {
	r := &types.Gifts{}
	err := b.apiRequest("getAvailableGifts", nil, r)
	return r, err
}

// SendGift - sends a gift to the given user or channel chat. The gift can't be
// converted to Telegram Stars by the receiver. Returns True on success.
//
// See https://core.telegram.org/bots/api#sendgift
func (b *Bot) SendGift(p *SendGiftParams) (bool, error) {
	var r bool
	err := b.apiRequest("sendGift", p, &r)
	return r, err
}

// GiftPremiumSubscription - Gifts a Telegram Premium subscription to the given
// user. Returns True on success.
//
// See https://core.telegram.org/bots/api#giftpremiumsubscription
func (b *Bot) GiftPremiumSubscription(p *GiftPremiumSubscriptionParams) (bool, error) {
	var r bool
	err := b.apiRequest("giftPremiumSubscription", p, &r)
	return r, err
}

// VerifyUser - verifies a user on behalf of the organization which is
// represented by the bot. Returns True on success.
//
// See https://core.telegram.org/bots/api#verifyuser
func (b *Bot) VerifyUser(p *VerifyUserParams) (bool, error) {
	var r bool
	err := b.apiRequest("verifyUser", p, &r)
	return r, err
}

// VerifyChat - verifies a chat on behalf of the organization which is
// represented by the bot. Returns True on success.
//
// See https://core.telegram.org/bots/api#verifychat
func (b *Bot) VerifyChat(p *VerifyChatParams) (bool, error) {
	var r bool
	err := b.apiRequest("verifyChat", p, &r)
	return r, err
}

// RemoveUserVerification - Removes verification from a user who is currently
// verified on behalf of the organization represented by the bot. Returns True on
// success.
//
// See https://core.telegram.org/bots/api#removeuserverification
func (b *Bot) RemoveUserVerification(p *RemoveUserVerificationParams) (bool, error) {
	var r bool
	err := b.apiRequest("removeUserVerification", p, &r)
	return r, err
}

// RemoveChatVerification - removes verification from a chat that is currently
// verified on behalf of the organization represented by the bot. Returns True on
// success.
//
// See https://core.telegram.org/bots/api#removechatverification
func (b *Bot) RemoveChatVerification(p *RemoveChatVerificationParams) (bool, error) {
	var r bool
	err := b.apiRequest("removeChatVerification", p, &r)
	return r, err
}

// ReadBusinessMessage - marks incoming message as read on behalf of a business
// account. Requires the can_read_messages business bot right. Returns True on
// success.
//
// See https://core.telegram.org/bots/api#readbusinessmessage
func (b *Bot) ReadBusinessMessage(p *ReadBusinessMessageParams) (bool, error) {
	var r bool
	err := b.apiRequest("readBusinessMessage", p, &r)
	return r, err
}

// DeleteBusinessMessages - delete messages on behalf of a business account.
// Requires the can_delete_sent_messages business bot right to delete messages
// sent by the bot itself, or the can_delete_all_messages business bot right to
// delete any message. Returns True on success.
//
// See https://core.telegram.org/bots/api#deletebusinessmessages
func (b *Bot) DeleteBusinessMessages(p *DeleteBusinessMessagesParams) (bool, error) {
	var r bool
	err := b.apiRequest("deleteBusinessMessages", p, &r)
	return r, err
}

// SetBusinessAccountName - changes the first and last name of a managed business
// account. Requires the can_change_name business bot right. Returns True on
// success.
//
// See https://core.telegram.org/bots/api#setbusinessaccountname
func (b *Bot) SetBusinessAccountName(p *SetBusinessAccountNameParams) (bool, error) {
	var r bool
	err := b.apiRequest("setBusinessAccountName", p, &r)
	return r, err
}

// SetBusinessAccountUsername - changes the username of a managed business
// account. Requires the can_change_username business bot right. Returns True on
// success.
//
// See https://core.telegram.org/bots/api#setbusinessaccountusername
func (b *Bot) SetBusinessAccountUsername(p *SetBusinessAccountUsernameParams) (bool, error) {
	var r bool
	err := b.apiRequest("setBusinessAccountUsername", p, &r)
	return r, err
}

// SetBusinessAccountBio - changes the bio of a managed business account.
// Requires the can_change_bio business bot right. Returns True on success.
//
// See https://core.telegram.org/bots/api#setbusinessaccountbio
func (b *Bot) SetBusinessAccountBio(p *SetBusinessAccountBioParams) (bool, error) {
	var r bool
	err := b.apiRequest("setBusinessAccountBio", p, &r)
	return r, err
}

// SetBusinessAccountProfilePhoto - Changes the profile photo of a managed
// business account. Requires the can_edit_profile_photo business bot right.
// Returns True on success.
//
// See https://core.telegram.org/bots/api#setbusinessaccountprofilephoto
func (b *Bot) SetBusinessAccountProfilePhoto(p *SetBusinessAccountProfilePhotoParams) (bool, error) {
	var r bool
	err := b.apiRequest("setBusinessAccountProfilePhoto", p, &r)
	return r, err
}

// RemoveBusinessAccountProfilePhoto - Removes the current profile photo of a
// managed business account. Requires the can_edit_profile_photo business bot
// right. Returns True on success.
//
// See https://core.telegram.org/bots/api#removebusinessaccountprofilephoto
func (b *Bot) RemoveBusinessAccountProfilePhoto(p *RemoveBusinessAccountProfilePhotoParams) (bool, error) {
	var r bool
	err := b.apiRequest("removeBusinessAccountProfilePhoto", p, &r)
	return r, err
}

// SetBusinessAccountGiftSettings - Changes the privacy settings pertaining to
// incoming gifts in a managed business account. Requires the
// can_change_gift_settings business bot right. Returns True on success.
//
// See https://core.telegram.org/bots/api#setbusinessaccountgiftsettings
func (b *Bot) SetBusinessAccountGiftSettings(p *SetBusinessAccountGiftSettingsParams) (bool, error) {
	var r bool
	err := b.apiRequest("setBusinessAccountGiftSettings", p, &r)
	return r, err
}

// GetBusinessAccountStarBalance - returns the amount of Telegram Stars owned by
// a managed business account. Requires the can_view_gifts_and_stars business bot
// right. Returns types.StarAmount on success.
//
// See https://core.telegram.org/bots/api#getbusinessaccountstarbalance
func (b *Bot) GetBusinessAccountStarBalance(p *GetBusinessAccountStarBalanceParams) (*types.StarAmount, error) {
	r := &types.StarAmount{}
	err := b.apiRequest("getBusinessAccountStarBalance", p, r)
	return r, err
}

// TransferBusinessAccountStars - transfers Telegram Stars from the business
// account balance to the bot's balance. Requires the can_transfer_stars business
// bot right. Returns True on success.
//
// See https://core.telegram.org/bots/api#transferbusinessaccountstars
func (b *Bot) TransferBusinessAccountStars(p *TransferBusinessAccountStarsParams) (bool, error) {
	var r bool
	err := b.apiRequest("transferBusinessAccountStars", p, &r)
	return r, err
}

// GetBusinessAccountGifts - returns the gifts received and owned by a managed
// business account. Requires the can_view_gifts_and_stars business bot right.
// Returns types.OwnedGifts on success.
//
// See https://core.telegram.org/bots/api#getbusinessaccountgifts
func (b *Bot) GetBusinessAccountGifts(p *GetBusinessAccountGiftsParams) (*types.OwnedGifts, error) {
	r := &types.OwnedGifts{}
	err := b.apiRequest("getBusinessAccountGifts", p, r)
	return r, err
}

// GetUserGifts - Returns the gifts owned and hosted by a user. Returns
// types.OwnedGifts on success.
//
// See https://core.telegram.org/bots/api#getusergifts
func (b *Bot) GetUserGifts(p *GetUserGiftsParams) (*types.OwnedGifts, error) {
	r := &types.OwnedGifts{}
	err := b.apiRequest("getUserGifts", p, r)
	return r, err
}

// GetChatGifts - Returns the gifts owned by a chat. Returns types.OwnedGifts on success.
//
// See https://core.telegram.org/bots/api#getchatgifts
func (b *Bot) GetChatGifts(p *GetChatGiftsParams) (*types.OwnedGifts, error) {
	r := &types.OwnedGifts{}
	err := b.apiRequest("getChatGifts", p, r)
	return r, err
}

// ConvertGiftToStars - Converts a given regular gift to Telegram Stars. Requires
// the can_convert_gifts_to_stars business bot right. Returns True on success.
//
// See https://core.telegram.org/bots/api#convertgifttostars
func (b *Bot) ConvertGiftToStars(p *ConvertGiftToStarsParams) (bool, error) {
	var r bool
	err := b.apiRequest("convertGiftToStars", p, &r)
	return r, err
}

// UpgradeGift - Upgrades a given regular gift to a unique gift. Requires the
// can_transfer_and_upgrade_gifts business bot right. Additionally, requires the
// can_transfer_stars business bot right if the upgrade is paid. Returns True on
// success.
//
// See https://core.telegram.org/bots/api#upgradegift
func (b *Bot) UpgradeGift(p *UpgradeGiftParams) (bool, error) {
	var r bool
	err := b.apiRequest("upgradeGift", p, &r)
	return r, err
}

// TransferGift - transfers an owned unique gift to another user. Requires the
// can_transfer_and_upgrade_gifts business bot right. Requires can_transfer_stars
// business bot right if the transfer is paid. Returns True on success.
//
// See https://core.telegram.org/bots/api#transfergift
func (b *Bot) TransferGift(p *TransferGiftParams) (bool, error) {
	var r bool
	err := b.apiRequest("transferGift", p, &r)
	return r, err
}

// PostStory - Posts a story on behalf of a managed business account. Requires
// the can_manage_stories business bot right. Returns types.Story on success.
//
// See https://core.telegram.org/bots/api#poststory
func (b *Bot) PostStory(p *PostStoryParams) (*types.Story, error) {
	r := &types.Story{}
	err := b.apiRequest("postStory", p, r)
	return r, err
}

// RepostStory - reposts a story on behalf of a business account from another
// business account. Both business accounts must be managed by the same bot, and
// the story on the source account must have been posted (or reposted) by the
// bot. Requires the can_manage_stories business bot right for both business
// accounts. Returns types.Story on success.
//
// See https://core.telegram.org/bots/api#repoststory
func (b *Bot) RepostStory(p *RepostStoryParams) (*types.Story, error) {
	r := &types.Story{}
	err := b.apiRequest("repostStory", p, r)
	return r, err
}

// EditStory - Edits a story previously posted by the bot on behalf of a managed
// business account. Requires the can_manage_stories business bot right. Returns
// types.Story on success.
//
// See https://core.telegram.org/bots/api#editstory
func (b *Bot) EditStory(p *EditStoryParams) (*types.Story, error) {
	r := &types.Story{}
	err := b.apiRequest("editStory", p, r)
	return r, err
}

// DeleteStory - Deletes a story previously posted by the bot on behalf of a
// managed business account. Requires the can_manage_stories business bot right.
// Returns True on success.
//
// See https://core.telegram.org/bots/api#deletestory
func (b *Bot) DeleteStory(p *DeleteStoryParams) (bool, error) {
	var r bool
	err := b.apiRequest("deleteStory", p, &r)
	return r, err
}

// SendPaidMedia - use this method to send paid media. On success, the sent
// types.Message is returned.
//
// See https://core.telegram.org/bots/api#sendpaidmedia
func (b *Bot) SendPaidMedia(p *SendPaidMediaParams) (*types.Message, error) {
	r := &types.Message{}
	err := b.apiRequest("sendPaidMedia", p, r)
	return r, err
}

// EditMessageText use this method to edit text and game messages. On success, if
// the edited message is not an inline message, the edited types.Message is
// returned, otherwise True is returned (in types.MessageOrBool structure). Note
// that business messages that were not sent by the bot and do not contain an
// inline keyboard can only be edited within 48 hours from the time they were
// sent.
//
// See https://core.telegram.org/bots/api#editmessagetext
func (b *Bot) EditMessageText(p *EditMessageTextParams) (*types.MessageOrBool, error) {
	r := &types.MessageOrBool{}
	err := b.apiRequest("editMessageText", p, r)
	return r, err
}

// EditMessageCaption use this method to edit captions of messages. On success,
// if the edited message is not an inline message, the edited types.Message is
// returned, otherwise True is returned (in types.MessageOrBool structure). Note
// that business messages that were not sent by the bot and do not contain an
// inline keyboard can only be edited within 48 hours from the time they were
// sent.
//
// See https://core.telegram.org/bots/api#editmessagecaption
func (b *Bot) EditMessageCaption(p *EditMessageCaptionParams) (*types.MessageOrBool, error) {
	r := &types.MessageOrBool{}
	err := b.apiRequest("editMessageCaption", p, r)
	return r, err
}

// EditMessageMedia use this method to edit animation, audio, document, photo, or
// video messages, or to add media to text messages. If a message is part of a
// message album, then it can be edited only to an audio for audio albums, only
// to a document for document albums and to a photo or a video otherwise. When an
// inline message is edited, a new file can't be uploaded; use a previously
// uploaded file via its file_id or specify a URL. On success, if the edited
// message is not an inline message, the edited types.Message is returned,
// otherwise True is returned (in types.MessageOrBool structure). Note that
// business messages that were not sent by the bot and do not contain an inline
// keyboard can only be edited within 48 hours from the time they were sent.
//
// See https://core.telegram.org/bots/api#editmessagemedia
func (b *Bot) EditMessageMedia(p *EditMessageMediaParams) (*types.MessageOrBool, error) {
	r := &types.MessageOrBool{}
	err := b.apiRequest("editMessageMedia", p, r)
	return r, err
}

// EditMessageLiveLocation use this method to edit live location messages. A
// location can be edited until its live_period expires or editing is explicitly
// disabled by a call to StopMessageLiveLocation. On success, if the edited
// message is not an inline message, the edited types.Message is returned, otherwise
// True is returned (in types.MessageOrBool structure).
//
// See https://core.telegram.org/bots/api#editmessagelivelocation
func (b *Bot) EditMessageLiveLocation(p *EditMessageLiveLocationParams) (*types.MessageOrBool, error) {
	r := &types.MessageOrBool{}
	err := b.apiRequest("editMessageLiveLocation", p, r)
	return r, err
}

// StopMessageLiveLocation use this method to stop updating a live location
// message before live_period expires. On success, if the message is not an
// inline message, the edited types.Message is returned, otherwise True is returned.
// (in types.MessageOrBool structure).
//
// See https://core.telegram.org/bots/api#stopmessagelivelocation
func (b *Bot) StopMessageLiveLocation(p *StopMessageLiveLocationParams) (*types.MessageOrBool, error) {
	r := &types.MessageOrBool{}
	err := b.apiRequest("stopMessageLiveLocation", p, r)
	return r, err
}

// EditMessageChecklist use this method to edit a checklist on behalf of a
// connected business account. On success, the edited types.Message is returned.
//
// See https://core.telegram.org/bots/api#editmessagechecklist
func (b *Bot) EditMessageChecklist(p *EditMessageChecklistParams) (*types.Message, error) {
	r := &types.Message{}
	err := b.apiRequest("editMessageChecklist", p, r)
	return r, err
}

// EditMessageReplyMarkup Use this method to edit only the reply markup of
// messages. On success, if the edited message is not an inline message, the
// edited types.Message is returned, otherwise True is returned (in
// types.MessageOrBool structure). Note that business messages that were not sent
// by the bot and do not contain an inline keyboard can only be edited within 48
// hours from the time they were sent.
//
// See https://core.telegram.org/bots/api#editmessagereplymarkup
func (b *Bot) EditMessageReplyMarkup(p *EditMessageReplyMarkupParams) (*types.MessageOrBool, error) {
	r := &types.MessageOrBool{}
	err := b.apiRequest("editMessageReplyMarkup", p, r)
	return r, err
}

// StopPoll - use this method to stop a poll which was sent by the bot. On
// success, the stopped types.Poll is returned.
//
// See https://core.telegram.org/bots/api#stoppoll
func (b *Bot) StopPoll(p *StopPollParams) (*types.Poll, error) {
	r := &types.Poll{}
	err := b.apiRequest("stopPoll", p, r)
	return r, err
}

// ApproveSuggestedPost use this method to approve a suggested post in a direct
// messages chat. The bot must have the 'can_post_messages' administrator right
// in the corresponding channel chat. Returns True on success.
//
// See https://core.telegram.org/bots/api#approvesuggestedpost
func (b *Bot) ApproveSuggestedPost(p *ApproveSuggestedPostParams) (bool, error) {
	var r bool
	err := b.apiRequest("approveSuggestedPost", p, &r)
	return r, err
}

// DeclineSuggestedPost use this method to decline a suggested post in a direct
// messages chat. The bot must have the 'can_manage_direct_messages'
// administrator right in the corresponding channel chat. Returns True on
// success.
//
// See https://core.telegram.org/bots/api#declinesuggestedpost
func (b *Bot) DeclineSuggestedPost(p *DeclineSuggestedPostParams) (bool, error) {
	var r bool
	err := b.apiRequest("declineSuggestedPost", p, &r)
	return r, err
}

// DeleteMessage use this method to delete a message, including service messages,
// with the following limitations:
//
// - A message can only be deleted if it was sent less than 48 hours ago.
//
// - Service messages about a supergroup, channel, or forum topic creation can't be deleted.
//
// - A dice message in a private chat can only be deleted if it was sent more than 24 hours ago.
//
// - Bots can delete outgoing messages in private chats, groups, and supergroups.
//
// - Bots can delete incoming messages in private chats.
//
// - Bots granted can_post_messages permissions can delete outgoing messages in channels.
//
// - If the bot is an administrator of a group, it can delete any message there.
//
// - If the bot has can_delete_messages administrator right in a supergroup or a channel, it can delete any message there.
//
// - If the bot has can_manage_direct_messages administrator right in a channel, it can delete any message in the corresponding direct messages chat.
//
// Returns True on success.
//
// See https://core.telegram.org/bots/api#deletemessage
func (b *Bot) DeleteMessage(p *DeleteMessageParams) (bool, error) {
	var r bool
	err := b.apiRequest("deleteMessage", p, &r)
	return r, err
}

// DeleteMessages use this method to delete multiple messages simultaneously. If
// some of the specified messages can't be found, they are skipped. Returns True
// on success.
//
// See https://core.telegram.org/bots/api#deletemessages
func (b *Bot) DeleteMessages(p *DeleteMessagesParams) (bool, error) {
	var r bool
	err := b.apiRequest("deleteMessages", p, &r)
	return r, err
}

// SendInvoice use this method to send invoices. On success, the sent
// types.Message is returned.
//
// See https://core.telegram.org/bots/api#sendinvoice
func (b *Bot) SendInvoice(p *SendInvoiceParams) (*types.Message, error) {
	r := &types.Message{}
	err := b.apiRequest("sendInvoice", p, r)
	return r, err
}

// CreateInvoiceLink use this method to create a link for an invoice. Returns the
// created invoice link as String on success.
//
// See https://core.telegram.org/bots/api#createinvoicelink
func (b *Bot) CreateInvoiceLink(p *CreateInvoiceLinkParams) (string, error) {
	var r string
	err := b.apiRequest("createInvoiceLink", p, &r)
	return r, err
}

// AnswerShippingQuery If you sent an invoice requesting a shipping address and
// the parameter is_flexible was specified, the TelegramBot API will send a [types.Update] with
// a shipping_query field to the bot. use this method to reply to shipping
// queries. On success, True is returned.
//
// See https://core.telegram.org/bots/api#answershippingquery
func (b *Bot) AnswerShippingQuery(p *AnswerShippingQueryParams) (bool, error) {
	var r bool
	err := b.apiRequest("answerShippingQuery", p, &r)
	return r, err
}

// AnswerPreCheckoutQuery Once the user has confirmed their payment and shipping
// details, the BotAPI sends the final confirmation in the form of a [types.Update]
// with the field pre_checkout_query. use this method to respond to such
// pre-checkout queries. On success, True is returned. Note: The BotAPI must
// receive an answer within 10 seconds after the pre-checkout query was sent.
//
// See https://core.telegram.org/bots/api#answerprecheckoutquery
func (b *Bot) AnswerPreCheckoutQuery(p *AnswerPreCheckoutQueryParams) (bool, error) {
	var r bool
	err := b.apiRequest("answerPreCheckoutQuery", p, &r)
	return r, err
}

// GetMyStarBalance - a method to get the current Telegram Stars balance of the
// bot. Requires no parameters. Returns [types.StarAmount] on success.
//
// See https://core.telegram.org/bots/api#getmystarbalance
func (b *Bot) GetMyStarBalance() (*types.StarAmount, error) {
	r := &types.StarAmount{}
	err := b.apiRequest("getMyStarBalance", nil, r)
	return r, err
}

// GetStarTransactions returns the bot's Telegram Star transactions in
// chronological order. On success, returns a [types.StarTransactions] object.
//
// See https://core.telegram.org/bots/api#getstartransactions
func (b *Bot) GetStarTransactions(p *GetStarTransactionsParams) (*types.StarTransactions, error) {
	r := &types.StarTransactions{}
	err := b.apiRequest("getStarTransactions", p, r)
	return r, err
}

// RefundStarPayment refunds a successful payment in Telegram Stars. Returns True
// on success.
//
// See https://core.telegram.org/bots/api#refundstarpayment
func (b *Bot) RefundStarPayment(p *RefundStarPaymentParams) (bool, error) {
	var r bool
	err := b.apiRequest("refundStarPayment", p, &r)
	return r, err
}

// EditUserStarSubscription allows the bot to cancel or re-enable extension of a
// subscription paid in Telegram Stars. Returns True on success.
//
// See https://core.telegram.org/bots/api#edituserstarsubscription
func (b *Bot) EditUserStarSubscription(p *EditUserStarSubscriptionParams) (bool, error) {
	var r bool
	err := b.apiRequest("editUserStarSubscription", p, &r)
	return r, err
}

// SetPassportDataErrors informs a user that some of the Telegram Passport
// elements they provided contains errors. The user will not be able to re-submit
// their Passport to you until the errors are fixed (the contents of the field
// for which you returned the error must change). Returns True on success.
//
// use this if the data submitted by the user doesn't satisfy the standards your
// service requires for any reason. For example, if a birthday date seems
// invalid, a submitted document is blurry, a scan shows evidence of tampering,
// etc. Supply some details in the error message to make sure the user knows how
// to correct the issues.
//
// See https://core.telegram.org/bots/api#setpassportdataerrors
func (b *Bot) SetPassportDataErrors(p *SetPassportDataErrorsParams) (bool, error) {
	var r bool
	err := b.apiRequest("setPassportDataErrors", p, &r)
	return r, err
}

// SendSticker - use this method to send static .WEBP, animated .TGS, or video
// .WEBM stickers. On success, the sent types.Message is returned.
//
// See https://core.telegram.org/bots/api#sendsticker
func (b *Bot) SendSticker(p *SendStickerParams) (*types.Message, error) {
	r := &types.Message{}
	err := b.apiRequest("sendSticker", p, r)
	return r, err
}

// GetStickerSet - use this method to get a sticker set. On success, a
// types.StickerSet object is returned.
//
// See https://core.telegram.org/bots/api#getstickerset
func (b *Bot) GetStickerSet(p *GetStickerSetParams) (*types.StickerSet, error) {
	r := &types.StickerSet{}
	err := b.apiRequest("getStickerSet", p, r)
	return r, err
}

// GetCustomEmojiStickers - use this method to get a sticker set. On success, a
// types.StickerSet object is returned
//
// See https://core.telegram.org/bots/api#getcustomemojistickers
func (b *Bot) GetCustomEmojiStickers(p *GetCustomEmojiStickersParams) ([]types.Sticker, error) {
	var r []types.Sticker
	err := b.apiRequest("getCustomEmojiStickers", p, &r)
	return r, err
}

// UploadStickerFile - use this method to upload a file with a sticker for later
// use in the CreateNewStickerSet, AddStickerToSet, or ReplaceStickerInSet
// methods (the file can be used multiple times). Returns the uploaded types.File on
// success.
//
// See https://core.telegram.org/bots/api#uploadstickerfile
func (b *Bot) UploadStickerFile(p *UploadStickerFileParams) (*types.File, error) {
	r := &types.File{}
	err := b.apiRequest("uploadStickerFile", p, r)
	return r, err
}

// CreateNewStickerSet - use this method to create a new sticker set owned by
// the bot. The bot will be able to edit the sticker set thus created. Returns
// True on success.
//
// See https://core.telegram.org/bots/api#createnewstickerset
func (b *Bot) CreateNewStickerSet(p *CreateNewStickerSetParams) (bool, error) {
	var r bool
	err := b.apiRequest("createNewStickerSet", p, &r)
	return r, err
}

// AddStickerToSet - use this method to add a new sticker to a set created by the
// bot. Emoji sticker sets can have up to 200 stickers. Other sticker sets can
// have up to 120 stickers. Returns True on success
//
// See https://core.telegram.org/bots/api#addstickertoset
func (b *Bot) AddStickerToSet(p *AddStickerToSetParams) (bool, error) {
	var r bool
	err := b.apiRequest("addStickerToSet", p, &r)
	return r, err
}

// SetStickerPositionInSet - use this method to move a sticker in a set created by
// the bot to a specific position. Returns True on success.
//
// See https://core.telegram.org/bots/api#setstickerpositioninset
func (b *Bot) SetStickerPositionInSet(p *SetStickerPositionInSetParams) (bool, error) {
	var r bool
	err := b.apiRequest("setStickerPositionInSet", p, &r)
	return r, err
}

// DeleteStickerFromSet - use this method to delete a sticker from a set created
// by the bot. Returns True on success.
//
// See https://core.telegram.org/bots/api#deletestickerfromset
func (b *Bot) DeleteStickerFromSet(p *DeleteStickerFromSetParams) (bool, error) {
	var r bool
	err := b.apiRequest("deleteStickerFromSet", p, &r)
	return r, err
}

// ReplaceStickerInSet - use this method to replace an existing sticker in a
// sticker set with a new one. The method is equivalent to calling
// DeleteStickerFromSet, then AddStickerToSet, then SetStickerPositionInSet.
// Returns True on success.
//
// See https://core.telegram.org/bots/api#replacestickerinset
func (b *Bot) ReplaceStickerInSet(p *ReplaceStickerInSetParams) (bool, error) {
	var r bool
	err := b.apiRequest("replaceStickerInSet", p, &r)
	return r, err
}

// SetStickerEmojiList - use this method to change the list of emoji assigned to a
// regular or custom emoji sticker. The sticker must belong to a sticker set
// created by the bot. Returns True on success.
//
// See https://core.telegram.org/bots/api#setstickeremojilist
func (b *Bot) SetStickerEmojiList(p *SetStickerEmojiListParams) (bool, error) {
	var r bool
	err := b.apiRequest("setStickerEmojiList", p, &r)
	return r, err
}

// SetStickerKeywords - use this method to change search keywords assigned to a
// regular or custom emoji sticker. The sticker must belong to a sticker set
// created by the bot. Returns True on success.
//
// See https://core.telegram.org/bots/api#setstickerkeywords
func (b *Bot) SetStickerKeywords(p *SetStickerKeywordsParams) (bool, error) {
	var r bool
	err := b.apiRequest("setStickerKeywords", p, &r)
	return r, err
}

// SetStickerMaskPosition - use this method to change the types.MaskPosition of a
// mask sticker. The sticker must belong to a sticker set created by the bot.
// Returns True on success.
//
// See https://core.telegram.org/bots/api#setstickermaskposition
func (b *Bot) SetStickerMaskPosition(p *SetStickerMaskPositionParams) (bool, error) {
	var r bool
	err := b.apiRequest("setStickerMaskPosition", p, &r)
	return r, err
}

// SetStickerSetTitle - use this method to set the title of a created sticker
// set. Returns True on success.
//
// See https://core.telegram.org/bots/api#setstickersettitle
func (b *Bot) SetStickerSetTitle(p *SetStickerSetTitleParams) (bool, error) {
	var r bool
	err := b.apiRequest("setStickerSetTitle", p, &r)
	return r, err
}

// SetStickerSetThumbnail - use this method to set the thumbnail of a sticker
// set. The format of the thumbnail file must match the format of the stickers in
// the set. Returns True on success.
//
// See https://core.telegram.org/bots/api#setstickersetthumbnail
func (b *Bot) SetStickerSetThumbnail(p *SetStickerSetThumbnailParams) (bool, error) {
	var r bool
	err := b.apiRequest("setStickerSetThumbnail", p, &r)
	return r, err
}

// SetCustomEmojiStickerSetThumbnail - use this method to set the thumbnail of a
// custom emoji sticker set. Returns True on success.
//
// See https://core.telegram.org/bots/api#setcustomemojistickersetthumbnail
func (b *Bot) SetCustomEmojiStickerSetThumbnail(p *SetCustomEmojiStickerSetThumbnailParams) (bool, error) {
	var r bool
	err := b.apiRequest("setCustomEmojiStickerSetThumbnail", p, &r)
	return r, err
}

// DeleteStickerSet - use this method to delete a sticker set created by the
// bot. Returns True on success.
//
// See https://core.telegram.org/bots/api#deletestickerset
func (b *Bot) DeleteStickerSet(p *DeleteStickerSetParams) (bool, error) {
	var r bool
	err := b.apiRequest("deleteStickerSet", p, &r)
	return r, err
}

// SetWebhook - use this method to specify a URL and receive incoming updates
// via an outgoing webhook. Whenever there is an update for the bot, we will send
// an HTTPS POST request to the specified URL, containing a JSON-serialized
// Update. In case of an unsuccessful delivery, we will repeat the request and
// give up after a reasonable amount of attempts. Returns True on success.
//
// See https://core.telegram.org/bots/api#setwebhook
func (b *Bot) SetWebhook(p *SetWebhookParams) (bool, error) {
	var r bool
	err := b.apiRequest("setWebhook", p, &r)
	return r, err
}

// DeleteWebhook - use this method to remove webhook integration if you decide
// to switch back to getUpdates. Returns True on success.
//
// See https://core.telegram.org/bots/api#deletewebhook
func (b *Bot) DeleteWebhook(p *DeleteWebhookParams) (bool, error) {
	var r bool
	err := b.apiRequest("deleteWebhook", p, &r)
	return r, err
}

// GetWebhookInfo - use this method to get current webhook status. Requires no
// parameters. On success, returns a WebhookInfo object. If the bot is using
// getUpdates, will return an object with the url field empty.
//
// See https://core.telegram.org/bots/api#getwebhookinfo
func (b *Bot) GetWebhookInfo() (*types.WebhookInfo, error) {
	r := &types.WebhookInfo{}
	err := b.apiRequest("getWebhookInfo", nil, r)
	return r, err
}
