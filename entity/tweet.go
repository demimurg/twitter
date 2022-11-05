package entity

type Tweet struct {
	ID         string
	Text       string
	Likes      int
	CommentIDs []string
}

type Comment struct {
	ID              string
	UserID          string
	Text            string
	Likes           int
	ReplyCommentIDs []string
}

type NewsFeed struct {
	ID     string
	Tweets []Tweet
}
