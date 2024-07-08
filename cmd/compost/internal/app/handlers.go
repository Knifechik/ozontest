package app

import (
	"context"
	"fmt"
	"log"
)

func (a *App) CreatePost(ctx context.Context, title, content string, authorID int, commentAllowed bool) (*Post, error) {
	newPost := Post{
		Title:           title,
		Content:         content,
		AuthorID:        authorID,
		CommentsAllowed: commentAllowed,
	}
	post, err := a.repo.SavePost(ctx, newPost)
	if err != nil {
		return nil, fmt.Errorf("repo.SavePost: %w", err)
	}

	return post, nil
}

func (a *App) GetPost(ctx context.Context, id int) (*Post, error) {
	post, err := a.repo.PostByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("repo.PostByID: %w", err)
	}

	return post, nil
}

func (a *App) GetPosts(ctx context.Context) ([]Post, error) {
	posts, err := a.repo.GetAllPosts(ctx)
	if err != nil {
		return nil, fmt.Errorf("repo.GetAllPosts: %w", err)
	}

	return posts, nil
}

func (a *App) Comments(ctx context.Context, ID, limit int, offset *int) ([]Comment, error) {
	comments, err := a.repo.CommentsByID(ctx, ID, nil, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("repo.CommentsByID: %w", err)
	}

	return comments, nil
}

func (a *App) ChildComments(ctx context.Context, postID int, commentID *int, limit int, offset *int) ([]Comment, error) {
	comments, err := a.repo.CommentsByID(ctx, postID, commentID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("repo.CommentsByID: %w", err)
	}

	return comments, nil
}

func (a *App) CreateComment(ctx context.Context, postID int, content string, authorID int, parentCommentID *int) (*Comment, error) {
	err := commentLengthCheck(content)
	if err != nil {
		return nil, fmt.Errorf("commentLengthCheck: %w", err)
	}
	
	post, err := a.repo.PostByID(ctx, postID)
	if err != nil {
		return nil, fmt.Errorf("repo.PostByID: %w", err)
	}
	if !post.CommentsAllowed {
		return nil, ErrNotAllowed
	}

	newComment := Comment{
		PostID:          postID,
		Content:         content,
		AuthorID:        authorID,
		ParentCommentID: parentCommentID,
	}

	comment, err := a.repo.SaveComment(ctx, newComment)
	if err != nil {
		return nil, fmt.Errorf("repo.SaveComment: %w", err)
	}

	go func() {
		a.mu.Lock()
		defer a.mu.Unlock()

		val, ok := a.CommentsObserver[postID]
		if ok {
			log.Printf("CommentObserver: %v", val)

			val.mu.Lock()
			defer val.mu.Unlock()

			for _, observer := range val.subs {

				select {
				case <-ctx.Done():
					return
				case observer <- *comment:

				default:

				}
			}
		}

	}()

	return comment, nil
}

func (a *App) Subscriptions(ctx context.Context, postID, userID int) (<-chan Comment, error) {

	a.mu.Lock()
	defer a.mu.Unlock()
	if _, ok := a.CommentsObserver[postID]; !ok {
		newListen := &listener{
			postID: postID,
			subs:   make(map[int]chan Comment),
		}

		a.CommentsObserver[postID] = newListen
	}

	ch := make(chan Comment)

	if listen, ok := a.CommentsObserver[postID]; ok {
		listen.mu.Lock()
		listen.subs[userID] = ch
		listen.mu.Unlock()
	}

	go func() {
		<-ctx.Done()
		a.mu.Lock()
		defer a.mu.Unlock()

		listen := a.CommentsObserver[postID]
		listen.mu.Lock()
		defer listen.mu.Unlock()

		ch = listen.subs[userID]
		close(ch)

		delete(listen.subs, userID)

		if len(listen.subs) == 0 {
			delete(a.CommentsObserver, postID)
		}

		log.Println("deleted from map subs", a.CommentsObserver)
	}()

	//достать канал, закрыть канал, после удалить канал

	log.Println(a.CommentsObserver)

	return ch, nil
}
