package graph

import (
	"context"
	"sort"
	"time"

	"github.com/Pawan2061/timeline_grpc_go/server/graph/generated"
	"github.com/Pawan2061/timeline_grpc_go/server/graph/model"
)

func (r *Resolver) GetTimeline(ctx context.Context, userID string) ([]*model.Post, error) {
	followedUsers := r.store.GetFollowers(userID)
	if len(followedUsers) == 0 {
		return []*model.Post{}, nil
	}

	postsChan := make(chan []*model.Post, len(followedUsers))

	for _, followedUserID := range followedUsers {
		go func(userID string) {
			storePosts := r.store.GetUserPosts(userID)
			posts := make([]*model.Post, len(storePosts))
			for i, post := range storePosts {
				posts[i] = &model.Post{
					ID:        post.ID,
					Content:   post.Content,
					Author:    post.AuthorID,
					Timestamp: post.Timestamp.Format(time.RFC3339),
				}
			}
			postsChan <- posts
		}(followedUserID)
	}

	var allPosts []*model.Post
	for i := 0; i < len(followedUsers); i++ {
		posts := <-postsChan
		allPosts = append(allPosts, posts...)
	}

	sort.Slice(allPosts, func(i, j int) bool {
		timeI, _ := time.Parse(time.RFC3339, allPosts[i].Timestamp)
		timeJ, _ := time.Parse(time.RFC3339, allPosts[j].Timestamp)
		return timeI.After(timeJ)
	})

	if len(allPosts) > 20 {
		allPosts = allPosts[:20]
	}

	return allPosts, nil
}

func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
