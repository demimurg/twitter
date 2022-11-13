package entity

type Tweet struct {
	ID         int
	Text       string
	Likes      int
	CommentIDs []string
}

type Comment struct {
	ID     int
	UserID string
	Text   string
	Likes  int
}
