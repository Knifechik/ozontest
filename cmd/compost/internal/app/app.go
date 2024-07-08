package app

import "sync"

type App struct {
	repo             Repo
	mu               sync.Mutex
	CommentsObserver map[int]*listener
}

type listener struct {
	postID int
	mu     sync.Mutex
	subs   map[int]chan Comment
}

func New(r Repo) *App {
	return &App{
		repo:             r,
		CommentsObserver: make(map[int]*listener),
	}
}
