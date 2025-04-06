package store

import (
	"sync"
	"time"
)

type User struct {
	ID       string
	Username string
}

type Post struct {
	ID        string
	Content   string
	AuthorID  string
	Timestamp time.Time
}

type Store struct {
	mu        sync.RWMutex
	users     map[string]*User
	posts     map[string][]*Post
	followers map[string][]string
	userPosts map[string][]*Post
}

func NewStore() *Store {
	s := &Store{
		users:     make(map[string]*User),
		posts:     make(map[string][]*Post),
		followers: make(map[string][]string),
		userPosts: make(map[string][]*Post),
	}
	s.initializeMockData()
	return s
}

func (s *Store) initializeMockData() {
	users := []*User{
		{ID: "1", Username: "user1"},
		{ID: "2", Username: "user2"},
		{ID: "3", Username: "user3"},
		{ID: "4", Username: "user4"},
		{ID: "5", Username: "user5"},
	}

	for _, user := range users {
		s.users[user.ID] = user
	}

	s.followers["1"] = []string{"2", "3", "4"}
	s.followers["2"] = []string{"1", "3"}
	s.followers["3"] = []string{"1", "4"}
	s.followers["4"] = []string{"1", "2"}
	s.followers["5"] = []string{"1"}

	now := time.Now()
	for _, user := range users {
		posts := make([]*Post, 0)
		for i := 0; i < 5; i++ {
			post := &Post{
				ID:        user.ID + "-post-" + string(rune('a'+i)),
				Content:   "Post " + string(rune('a'+i)) + " from " + user.Username,
				AuthorID:  user.ID,
				Timestamp: now.Add(-time.Duration(i) * time.Hour),
			}
			posts = append(posts, post)
		}
		s.userPosts[user.ID] = posts
	}
}

func (s *Store) GetUserPosts(userID string) []*Post {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.userPosts[userID]
}

func (s *Store) GetFollowers(userID string) []string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.followers[userID]
}

func (s *Store) GetUser(userID string) *User {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.users[userID]
}
