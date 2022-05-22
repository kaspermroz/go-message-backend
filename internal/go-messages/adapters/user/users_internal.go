package user

import "github.com/kaspermroz/go-message-backend/internal/go-messages/domain"

var UserOneID = domain.MustNewUUID("2137")
var UserTwoID = domain.MustNewUUID("1337")

var UserOne = domain.MustNewUser(UserOneID, domain.MustNewName("user1"))
var UserTwo = domain.MustNewUser(UserTwoID, domain.MustNewName("user2"))
