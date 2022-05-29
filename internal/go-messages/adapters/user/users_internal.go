package user

import "github.com/kaspermroz/go-message-backend/internal/go-messages/domain"

var UserOneID = domain.MustNewUUID("auth0|62920cb54abeda006ad5f206")
var UserTwoID = domain.MustNewUUID("auth0|62920cea79f99400690108b8")

var UserOne = domain.MustNewUser(UserOneID, domain.MustNewName("user1"))
var UserTwo = domain.MustNewUser(UserTwoID, domain.MustNewName("user2"))
