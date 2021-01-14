package test

import "gitlab.com/pinvest/internships/hydra/onboarding-dean/app/model/orm"

var UsersData = []orm.User{
	{
		Name:     "User 1",
		Email:    "user1@user.com",
		Password: "password",
	},
	{
		Name:     "User 2",
		Email:    "user2@user.com",
		Password: "password",
	},
}

var ThreadsData = []orm.Thread{
	{
		Name:   "Thread 1",
		Description: "Description 1",
		UserID: 1,
	},
	{
		Name:   "Thread 2",
		Description: "Description 1",
		UserID: 1,
	},
	{
		Name:   "Thread 3",
		Description: "Description 1",
		UserID: 2,
	},
	{
		Name:   "Thread 4",
		Description: "Description 1",
		UserID: 2,
	},
	{
		Name:   "Thread 5",
		Description: "Description 1",
		UserID: 2,
	},
}

var PostsData = []orm.Post{
	{
		Title:    "Post 1",
		Content:  "Post content 1",
		UserID:   1,
		ThreadID: 1,
	},
	{
		Title:    "Post 2",
		Content:  "Post content 2",
		UserID:   2,
		ThreadID: 1,
	},
	{
		Title:    "Post 3",
		Content:  "Post content 3",
		UserID:   2,
		ThreadID: 1,
	},
	{
		Title:    "Post 4",
		Content:  "Post content 4",
		UserID:   2,
		ThreadID: 2,
	},
	{
		Title:    "Post 5",
		Content:  "Post content 5",
		UserID:   1,
		ThreadID: 2,
	},
}

var TokensData = []orm.Token{
	{
		UserID:    uint(1),
		TokenUUID: "66f0ac33-f031-4ae5-8cae-cb5eef3536e6",
	},
	{
		UserID:    uint(2),
		TokenUUID: "e0cff24d-e353-448f-8328-9f8a5d9a4d00",
	},
}

var StringToken = []string{
	"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJPbmJvYXJkaW5nIERlYW4iLCJ1c2VyX2lkIjoxLCJ0b2tlbl91dWlkIjoiNjZmMGFjMzMtZjAzMS00YWU1LThjYWUtY2I1ZWVmMzUzNmU2In0.XH597NUrRydchWDpiw_ax104ymtldISH9riwNiQL7Rc",
	"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJPbmJvYXJkaW5nIERlYW4iLCJ1c2VyX2lkIjoyLCJ0b2tlbl91dWlkIjoiZTBjZmYyNGQtZTM1My00NDhmLTgzMjgtOWY4YTVkOWE0ZDAwIn0.VDAj69wpMXvmhwSpyRbmGo1eZKoDxZlwPevksioVwYg",
}
