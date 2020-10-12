package result

import "time"

type Posts []Post

func (p Posts) String() string {
	panic("implement me")
}

type Post struct {
	Id            int
	Title         string
	Body          string
	Draft         bool
	Archived      bool
	Url           string
	CreatedAt     *time.Time
	UpdatedAt     *time.Time
	Scope         string
	SharingUrl    string
	Tags          []string
	User          User
	StarCount     int
	GoodJobsCount int
	Comments      []Comment
	Groups        Groups
}

type Comment struct {
	Id        int
	Body      string
	CreatedAt *time.Time
	User      User
}
