package in_memory

import (
	"context"
	"fmt"
	"ozon_test_compost/cmd/compost/internal/app"
	"sync"
	"time"
)

var _ app.Repo = &Repo{}

type Post struct {
	ID              int       `db:"id" json:"id"`
	Title           string    `db:"title" json:"title"`
	Content         string    `db:"content" json:"content"`
	AuthorID        int       `db:"author_id" json:"authorId"`
	CommentsAllowed bool      `db:"comments_allowed" json:"commentsAllowed"`
	CreatedAt       time.Time `db:"created_at" json:"createdAt"`
}

type Comment struct {
	ID              int       `db:"id" json:"id"`
	PostID          int       `db:"post_id" json:"postId"`
	Content         string    `db:"content" json:"content"`
	AuthorID        int       `db:"author_id" json:"authorId"`
	ParentCommentID *int      `db:"parent_comment_id" json:"parentCommentId,omitempty"`
	CreatedAt       time.Time `db:"created_at" json:"createdAt"`
}

type postInfo struct {
	posts *Post

	mu       sync.Mutex
	comments map[int]*Comment
}

type Repo struct {
	mu             sync.Mutex
	posts          map[int]*postInfo
	postCounter    int
	commentCounter int
}

func New() *Repo {
	return &Repo{
		posts: make(map[int]*postInfo),
	}
}

func (r *Repo) SavePost(ctx context.Context, post app.Post) (*app.Post, error) {
	p := postConvert(post)

	r.postCounter++
	p.ID = r.postCounter
	p.CreatedAt = time.Now()

	pInfo := &postInfo{
		posts:    p,
		mu:       sync.Mutex{},
		comments: make(map[int]*Comment),
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	r.posts[p.ID] = pInfo

	return p.Convert(), nil
}

func (r *Repo) PostByID(ctx context.Context, id int) (*app.Post, error) {

	r.mu.Lock()
	defer r.mu.Unlock()

	p, ok := r.posts[id]
	if !ok {
		return nil, fmt.Errorf("post not found")
	}

	return p.posts.Convert(), nil
}

func (r *Repo) GetAllPosts(ctx context.Context) ([]app.Post, error) {

	r.mu.Lock()
	defer r.mu.Unlock()

	posts := make([]app.Post, 0, len(r.posts))
	for _, p := range r.posts {
		posts = append(posts, *p.posts.Convert())
	}

	return posts, nil
}

func (r *Repo) SaveComment(ctx context.Context, comment app.Comment) (*app.Comment, error) {
	c := commentConvert(comment)

	r.mu.Lock()
	defer r.mu.Unlock()

	pInfo, ok := r.posts[c.PostID]
	if !ok {
		return nil, fmt.Errorf("post not found")
	}

	r.commentCounter++

	c.CreatedAt = time.Now()
	c.ID = r.commentCounter

	pInfo.mu.Lock()
	defer pInfo.mu.Unlock()

	pInfo.comments[c.ID] = c

	return c.Convert(), nil

}

func (r *Repo) CommentsByID(ctx context.Context, id int, parentID *int, limit int, offset *int) ([]app.Comment, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	pInfo, ok := r.posts[id]
	if !ok {
		return nil, fmt.Errorf("post not found")
	}

	pInfo.mu.Lock()
	defer pInfo.mu.Unlock()

	comments := make([]app.Comment, 0, len(pInfo.comments))
	
	for _, val := range pInfo.comments {

		if parentID != nil && val.ParentCommentID != nil {
			if *val.ParentCommentID == *parentID {
				comments = append(comments, *val.Convert())
			}
			continue
		}

		if val.ParentCommentID == parentID {
			comments = append(comments, *val.Convert())
		}
	}

	return comments, nil
}
