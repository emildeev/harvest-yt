package port

import (
	tracker "github.com/emildeev/yandex-tracker-go"

	ytrackercore "github.com/emildeev/harvest-yt/internal/core/y_tracker"
)

func UserToCore(user *tracker.User) ytrackercore.User {
	return ytrackercore.User{
		UID:         user.UID,
		Login:       user.Login,
		TrackerUID:  user.TrackerUid,
		PassportUid: user.PassportUid,
		CloudUid:    user.CloudUid,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Display:     user.Display,
		Email:       user.Email,
	}
}
