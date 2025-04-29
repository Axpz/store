package storage

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/Axpz/store/internal/config"
	"github.com/google/go-github/v45/github"
	"golang.org/x/oauth2"
)

// GitHubStore 实现使用 GitHub 作为存储
type GitHubStore struct {
	client *github.Client

	Store
}

// NewGitHubStore 创建一个新的 GitHub 存储
func NewGitHubStore(cfg *config.Config) (StoreInterface, error) {
	// 创建 OAuth2 客户端
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: cfg.GitHub.Token},
	)
	tc := oauth2.NewClient(ctx, ts)

	// 创建 GitHub 客户端
	client := github.NewClient(tc)

	store := &GitHubStore{
		client: client,
		Store:  NewStore(cfg),
	}

	return store, nil
}

// Create 创建新用户
func (s *GitHubStore) Create(user User) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// 确保用户表已加载
	if err := s.loadUsers(); err != nil {
		return err
	}

	if _, exists := s.users[user.ID]; exists {
		return fmt.Errorf("用户已存在")
	}

	s.users[user.ID] = user
	return s.saveUsers()
}

// Get 获取用户
func (s *GitHubStore) Get(id string) (User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var result User

	// 确保用户表已加载
	if err := s.loadUsers(); err != nil {
		return result, err
	}

	result, exists := s.users[id]
	if !exists {
		return result, fmt.Errorf("用户不存在")
	}

	return result, nil
}

// Update 更新用户
func (s *GitHubStore) Update(user User) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// 确保用户表已加载
	if err := s.loadUsers(); err != nil {
		return err
	}

	if _, exists := s.users[user.ID]; !exists {
		return fmt.Errorf("用户不存在")
	}

	s.users[user.ID] = user
	return s.saveUsers()
}

// Delete 删除用户
func (s *GitHubStore) Delete(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// 确保用户表已加载
	if err := s.loadUsers(); err != nil {
		return err
	}

	if _, exists := s.users[id]; !exists {
		return fmt.Errorf("用户不存在")
	}

	delete(s.users, id)
	return s.saveUsers()
}

// CreateComment 创建新评论
func (s *GitHubStore) CreateComment(comment Comment) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// 确保评论表已加载
	if err := s.loadComments(); err != nil {
		return err
	}

	if _, exists := s.comments[comment.ID]; exists {
		return fmt.Errorf("评论已存在")
	}

	s.comments[comment.ID] = comment
	return s.saveComments()
}

// GetComment 获取评论
func (s *GitHubStore) GetComment(id string) (Comment, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var result Comment

	// 确保评论表已加载
	if err := s.loadComments(); err != nil {
		return result, err
	}

	comment, exists := s.comments[id]
	if !exists {
		return result, fmt.Errorf("评论不存在")
	}

	result = comment
	return result, nil
}

// UpdateComment 更新评论
func (s *GitHubStore) UpdateComment(comment Comment) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// 确保评论表已加载
	if err := s.loadComments(); err != nil {
		return err
	}

	if _, exists := s.comments[comment.ID]; !exists {
		return fmt.Errorf("评论不存在")
	}

	s.comments[comment.ID] = comment
	return s.saveComments()
}

// DeleteComment 删除评论
func (s *GitHubStore) DeleteComment(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// 确保评论表已加载
	if err := s.loadComments(); err != nil {
		return err
	}

	if _, exists := s.comments[id]; !exists {
		return fmt.Errorf("评论不存在")
	}

	delete(s.comments, id)
	return s.saveComments()
}

// loadTable 加载指定表的数据
func (s *GitHubStore) loadTable(tableName string, data any) error {
	content, _, _, err := s.client.Repositories.GetContents(
		s.ctx,
		s.config.GitHub.Repo.Owner,
		s.config.GitHub.Repo.Name,
		s.config.GetTablePath(tableName),
		&github.RepositoryContentGetOptions{
			Ref: s.config.GitHub.Repo.Branch,
		},
	)
	if err != nil {
		return fmt.Errorf("获取文件内容失败: %v", err)
	}

	// 解码 base64 内容
	decoded, err := base64.StdEncoding.DecodeString(*content.Content)
	if err != nil {
		return fmt.Errorf("解码文件内容失败: %v", err)
	}

	// 解析 JSON
	if err := json.Unmarshal(decoded, data); err != nil {
		return fmt.Errorf("解析 JSON 失败: %v", err)
	}

	return nil
}

// saveTable 保存指定表的数据
func (s *GitHubStore) saveTable(tableName string, data any) (err error) {
	doSave := func() error {
		// 序列化为 JSON
		jsonData, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			return fmt.Errorf("序列化 JSON 失败: %v", err)
		}

		// 获取当前文件的 SHA
		content, _, _, err := s.client.Repositories.GetContents(
			s.ctx,
			s.config.GitHub.Repo.Owner,
			s.config.GitHub.Repo.Name,
			s.config.GetTablePath(tableName),
			&github.RepositoryContentGetOptions{
				Ref: s.config.GitHub.Repo.Branch,
			},
		)
		if err != nil {
			return fmt.Errorf("获取文件 SHA 失败: %v", err)
		}

		// 创建提交
		opts := &github.RepositoryContentFileOptions{
			Message: github.String(fmt.Sprintf("Update %s table", tableName)),
			Content: jsonData,
			SHA:     content.SHA,
			Branch:  github.String(s.config.GitHub.Repo.Branch),
		}

		_, _, err = s.client.Repositories.CreateFile(
			s.ctx,
			s.config.GitHub.Repo.Owner,
			s.config.GitHub.Repo.Name,
			s.config.GetTablePath(tableName),
			opts,
		)
		if err != nil {
			return fmt.Errorf("创建文件失败: %v", err)
		}

		return nil
	}

	return s.throttlesaver.RequestMustSave(doSave)
}

func (s *GitHubStore) loadUsers() error {
	if s.loaded["users"] {
		return nil
	}

	if err := s.loadTable("users", &s.users); err != nil {
		return fmt.Errorf("加载用户表失败: %v", err)
	}
	s.loaded["users"] = true
	return nil
}

func (s *GitHubStore) saveUsers() error {
	return s.saveTable("users", s.users)
}

func (s *GitHubStore) loadComments() error {
	if s.loaded["comments"] {
		return nil
	}

	if err := s.loadTable("comments", &s.comments); err != nil {
		return fmt.Errorf("加载用户表失败: %v", err)
	}
	s.loaded["comments"] = true
	return nil
}

func (s *GitHubStore) saveComments() error {
	return s.saveTable("comments", s.comments)
}
