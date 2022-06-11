package user

import "github.com/kaspermroz/go-message-backend/internal/go-messages/domain"

var UserOneID = domain.MustNewUUID("auth0|62920cb54abeda006ad5f206")
var UserTwoID = domain.MustNewUUID("auth0|62920cea79f99400690108b8")
var E2EUserID = domain.MustNewUUID("auth0|62a47030624b9df84675d21d")

var UserOne = domain.MustNewUser(UserOneID, domain.MustNewName("user1"))
var UserTwo = domain.MustNewUser(UserTwoID, domain.MustNewName("user2"))
var E2EUser = domain.MustNewUser(E2EUserID, domain.MustNewName("e2e"))
