package app

import "sync"

type App struct {
	repo             Repo
	mu               sync.Mutex
	CommentsObserver map[int]*listener //postID
}

type listener struct {
	postID int
	mu     sync.Mutex
	subs   map[int]chan Comment //userID
}

func New(r Repo) *App {
	return &App{
		repo:             r,
		CommentsObserver: make(map[int]*listener),
	}
}
