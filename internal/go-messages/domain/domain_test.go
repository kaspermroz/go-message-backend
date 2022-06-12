package domain_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/kaspermroz/go-message-backend/internal/go-messages/domain"
)

func Test_Message(t *testing.T) {
	failedTxt, err := domain.NewText("")
	require.Error(t, err)
	require.Zero(t, failedTxt)

	text, err := domain.NewText("test")
	require.NoError(t, err)
	require.NotZero(t, text)

	uuid := domain.MustNewUUID("1234")
	failedMsg, err := domain.NewMessage(uuid, domain.Text{})
	require.Error(t, err)
	require.Zero(t, failedMsg)

	msg, err := domain.NewMessage(uuid, text)
	require.NoError(t, err)
	require.NotZero(t, msg)
	require.Equal(t, msg.AuthorID(), uuid)
	require.Equal(t, msg.Text(), text)
}

func Test_User(t *testing.T) {
	uuid := domain.MustNewUUID("1234")
	failedName, err := domain.NewName("")
	require.Error(t, err)
	require.Zero(t, failedName)

	name, err := domain.NewName("Test user")
	require.NoError(t, err)
	require.NotZero(t, name)

	failedUser, err := domain.NewUser(uuid, domain.Name{})
	require.Error(t, err)
	require.Zero(t, failedUser)

	user, err := domain.NewUser(uuid, name)
	require.NoError(t, err)
	require.NotZero(t, user)
	require.Equal(t, user.UUID(), uuid)
	require.Equal(t, user.Name(), name)
}

func Test_Chat(t *testing.T) {
	uuid1 := domain.MustNewUUID("4321")
	uuid2 := domain.MustNewUUID("4322")
	title1 := domain.MustNewTitle("test chat")
	title2 := domain.MustNewTitle("test chat 2")
	user1 := domain.MustNewUser(domain.MustNewUUID("1234"), domain.MustNewName("user1"))
	user2 := domain.MustNewUser(domain.MustNewUUID("2234"), domain.MustNewName("user2"))

	failedChat1, err := domain.NewChat(uuid1, title1, []domain.User{})
	require.Error(t, err)
	require.Zero(t, failedChat1)
	failedChat2, err := domain.NewChat(uuid1, domain.Title{}, []domain.User{user1, user2})
	require.Error(t, err)
	require.Zero(t, failedChat2)

	chat1, err := domain.NewChat(uuid1, title1, []domain.User{user1})
	require.NoError(t, err)
	require.NotZero(t, chat1)
	chat2, err := domain.NewChat(uuid2, title2, []domain.User{user2})
	require.NoError(t, err)
	require.NotZero(t, chat2)

	require.True(t, chat1.HasUser(user1.UUID()))
	require.True(t, chat2.HasUser(user2.UUID()))

	err = chat1.AddMessage(domain.MustNewMessage(user2.UUID(), domain.MustNewText("test message 1")))
	require.Error(t, err, "user not present in this chat")
	require.Len(t, chat1.Messages(), 0)

	err = chat2.AddMessage(domain.MustNewMessage(user2.UUID(), domain.MustNewText("test message 2")))
	require.NoError(t, err)
	require.Len(t, chat2.Messages(), 1)
}
