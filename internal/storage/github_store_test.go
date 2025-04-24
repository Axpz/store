package storage

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/Axpz/store/internal/config"
	"github.com/stretchr/testify/assert"
)

// setupGitHubTestStore 设置测试环境
func setupGitHubTestStore() (StoreInterface, error) {
	// 创建配置
	cfg := &config.Config{
		GitHub: config.GitHubConfig{
			Token: os.Getenv("GITHUB_API_TOKEN"),
			Repo: config.RepoConfig{
				Owner:  "Axpz",
				Name:   "store",
				Branch: "main",
				Tables: config.TablesConfig{
					Path:     "tables",
					Users:    "users.json",
					Comments: "comments.json",
				},
			},
		},
		Storage: config.StorageConfig{
			Type: "github",
		},
	}

	// 创建存储实例
	store, err := NewGitHubStore(cfg)
	if err != nil {
		return nil, fmt.Errorf("创建 GitHub 存储失败: %v", err)
	}

	// 初始化用户表
	initialUsers := map[string]User{}
	store.(*GitHubStore).users = initialUsers

	// 保存初始用户表
	err = store.(*GitHubStore).saveUsers()
	if err != nil {
		return nil, fmt.Errorf("保存初始用户表失败: %v", err)
	}

	return store, nil
}

func TestGitHubStore_UserCRUD(t *testing.T) {
	store, err := setupGitHubTestStore()
	if err != nil {
		t.Fatalf("设置测试环境失败: %v", err)
	}

	// 测试创建用户
	user := User{
		ID:       "test-user-1",
		Username: "Test User",
		Email:    "test@example.com",
		Plan:     "free",
		Created:  time.Now().Unix(),
		Updated:  time.Now().Unix(),
	}
	err = store.Create(user)
	assert.NoError(t, err)

	// 测试获取用户
	got, err := store.Get(user.ID)
	assert.NoError(t, err)
	assert.Equal(t, user, got)

	// 测试更新用户
	user.Username = "Updated User"
	err = store.Update(user)
	assert.NoError(t, err)

	got, err = store.Get(user.ID)
	assert.NoError(t, err)
	assert.Equal(t, user, got)

	// 测试删除用户
	err = store.Delete(user.ID)
	assert.NoError(t, err)

	_, err = store.Get(user.ID)
	assert.Error(t, err)
}

// func TestGitHubStore_CommentCRUD(t *testing.T) {
// 	store, mockClient, cleanup := setupGitHubTestStore(t)
// 	defer cleanup()

// 	// 设置模拟行为
// 	mockClient.On("GetContents", mock.Anything, "owner", "repo", "tables/comments.json", mock.Anything).Return(nil, nil, nil)
// 	mockClient.On("CreateFile", mock.Anything, "owner", "repo", "tables/comments.json", mock.Anything).Return(nil, nil, nil)

// 	// 测试创建评论
// 	comment := Comment{
// 		ID:      "test-comment-1",
// 		Content: "Test Comment",
// 		UserID:  "test-user-1",
// 		Created: time.Now().Unix(),
// 		Updated: time.Now().Unix(),
// 	}
// 	err := store.CreateComment(comment)
// 	assert.NoError(t, err)

// 	// 设置模拟行为 - 获取评论
// 	commentData := map[string]Comment{
// 		comment.ID: comment,
// 	}
// 	jsonData, _ := json.Marshal(commentData)
// 	mockClient.store["tables/comments.json"] = jsonData
// 	mockClient.On("GetContents", mock.Anything, "owner", "repo", "tables/comments.json", mock.Anything).Return(nil, nil, nil)

// 	// 测试获取评论
// 	got, err := store.GetComment(comment.ID)
// 	assert.NoError(t, err)
// 	assert.Equal(t, comment, got)

// 	// 设置模拟行为 - 更新评论
// 	mockClient.On("CreateFile", mock.Anything, "owner", "repo", "tables/comments.json", mock.Anything).Return(nil, nil, nil)

// 	// 测试更新评论
// 	comment.Content = "Updated Comment"
// 	err = store.UpdateComment(comment)
// 	assert.NoError(t, err)

// 	// 设置模拟行为 - 获取更新后的评论
// 	commentData[comment.ID] = comment
// 	jsonData, _ = json.Marshal(commentData)
// 	mockClient.store["tables/comments.json"] = jsonData
// 	mockClient.On("GetContents", mock.Anything, "owner", "repo", "tables/comments.json", mock.Anything).Return(nil, nil, nil)

// 	got, err = store.GetComment(comment.ID)
// 	assert.NoError(t, err)
// 	assert.Equal(t, comment, got)

// 	// 设置模拟行为 - 删除评论
// 	mockClient.On("CreateFile", mock.Anything, "owner", "repo", "tables/comments.json", mock.Anything).Return(nil, nil, nil)

// 	// 测试删除评论
// 	err = store.DeleteComment(comment.ID)
// 	assert.NoError(t, err)

// 	// 设置模拟行为 - 获取不存在的评论
// 	delete(commentData, comment.ID)
// 	jsonData, _ = json.Marshal(commentData)
// 	mockClient.store["tables/comments.json"] = jsonData
// 	mockClient.On("GetContents", mock.Anything, "owner", "repo", "tables/comments.json", mock.Anything).Return(nil, nil, nil)

// 	_, err = store.GetComment(comment.ID)
// 	assert.Error(t, err)
// }

// func TestGitHubStore_ConcurrentAccess(t *testing.T) {
// 	store, mockClient, cleanup := setupGitHubTestStore(t)
// 	defer cleanup()

// 	// 设置模拟行为
// 	mockClient.On("GetContents", mock.Anything, "owner", "repo", "tables/users.json", mock.Anything).Return(nil, nil, nil)
// 	mockClient.On("CreateFile", mock.Anything, "owner", "repo", "tables/users.json", mock.Anything).Return(nil, nil, nil)

// 	// 创建测试用户
// 	user := User{
// 		ID:       "test-user-1",
// 		Username: "Test User",
// 		Email:    "test@example.com",
// 		Plan:     "free",
// 		Created:  time.Now().Unix(),
// 		Updated:  time.Now().Unix(),
// 	}
// 	err := store.Create(user)
// 	assert.NoError(t, err)

// 	// 设置模拟行为 - 获取用户
// 	userData := map[string]User{
// 		user.ID: user,
// 	}
// 	jsonData, _ := json.Marshal(userData)
// 	mockClient.store["tables/users.json"] = jsonData
// 	mockClient.On("GetContents", mock.Anything, "owner", "repo", "tables/users.json", mock.Anything).Return(nil, nil, nil)

// 	// 并发读取和更新
// 	done := make(chan bool)
// 	for i := 0; i < 10; i++ {
// 		go func() {
// 			// 读取用户
// 			_, err := store.Get(user.ID)
// 			assert.NoError(t, err)

// 			// 更新用户
// 			user.Username = "Updated User"
// 			err = store.Update(user)
// 			assert.NoError(t, err)

// 			done <- true
// 		}()
// 	}

// 	// 等待所有 goroutine 完成
// 	for i := 0; i < 10; i++ {
// 		<-done
// 	}
// }

// func TestGitHubStore_ErrorCases(t *testing.T) {
// 	store, mockClient, cleanup := setupGitHubTestStore(t)
// 	defer cleanup()

// 	// 设置模拟行为 - 获取不存在的用户
// 	mockClient.On("GetContents", mock.Anything, "owner", "repo", "tables/users.json", mock.Anything).Return(nil, nil, nil)
// 	userData := make(map[string]User)
// 	jsonData, _ := json.Marshal(userData)
// 	mockClient.store["tables/users.json"] = jsonData

// 	// 测试获取不存在的用户
// 	_, err := store.Get("non-existent")
// 	assert.Error(t, err)

// 	// 测试更新不存在的用户
// 	err = store.Update(User{ID: "non-existent"})
// 	assert.Error(t, err)

// 	// 测试删除不存在的用户
// 	err = store.Delete("non-existent")
// 	assert.Error(t, err)

// 	// 设置模拟行为 - 创建用户
// 	mockClient.On("CreateFile", mock.Anything, "owner", "repo", "tables/users.json", mock.Anything).Return(nil, nil, nil)

// 	// 测试创建重复用户
// 	user := User{
// 		ID:       "test-user-1",
// 		Username: "Test User",
// 		Email:    "test@example.com",
// 		Plan:     "free",
// 		Created:  time.Now().Unix(),
// 		Updated:  time.Now().Unix(),
// 	}
// 	err = store.Create(user)
// 	assert.NoError(t, err)

// 	// 设置模拟行为 - 获取用户
// 	userData[user.ID] = user
// 	jsonData, _ = json.Marshal(userData)
// 	mockClient.store["tables/users.json"] = jsonData
// 	mockClient.On("GetContents", mock.Anything, "owner", "repo", "tables/users.json", mock.Anything).Return(nil, nil, nil)

// 	err = store.Create(user)
// 	assert.Error(t, err)

// 	// 设置模拟行为 - 获取不存在的评论
// 	mockClient.On("GetContents", mock.Anything, "owner", "repo", "tables/comments.json", mock.Anything).Return(nil, nil, nil)
// 	commentData := make(map[string]Comment)
// 	jsonData, _ = json.Marshal(commentData)
// 	mockClient.store["tables/comments.json"] = jsonData

// 	// 测试评论相关错误
// 	_, err = store.GetComment("non-existent")
// 	assert.Error(t, err)

// 	err = store.UpdateComment(Comment{ID: "non-existent"})
// 	assert.Error(t, err)

// 	err = store.DeleteComment("non-existent")
// 	assert.Error(t, err)

// 	// 设置模拟行为 - 创建评论
// 	mockClient.On("CreateFile", mock.Anything, "owner", "repo", "tables/comments.json", mock.Anything).Return(nil, nil, nil)

// 	comment := Comment{
// 		ID:      "test-comment-1",
// 		Content: "Test Comment",
// 		UserID:  "test-user-1",
// 		Created: time.Now().Unix(),
// 		Updated: time.Now().Unix(),
// 	}
// 	err = store.CreateComment(comment)
// 	assert.NoError(t, err)

// 	// 设置模拟行为 - 获取评论
// 	commentData[comment.ID] = comment
// 	jsonData, _ = json.Marshal(commentData)
// 	mockClient.store["tables/comments.json"] = jsonData
// 	mockClient.On("GetContents", mock.Anything, "owner", "repo", "tables/comments.json", mock.Anything).Return(nil, nil, nil)

// 	err = store.CreateComment(comment)
// 	assert.Error(t, err)
// }

// func TestGitHubStore_LoadTable(t *testing.T) {
// 	store, mockClient, cleanup := setupGitHubTestStore(t)
// 	defer cleanup()

// 	// 准备测试数据
// 	testData := map[string]User{
// 		"test-user-1": {
// 			ID:       "test-user-1",
// 			Username: "Test User",
// 			Email:    "test@example.com",
// 			Plan:     "free",
// 			Created:  time.Now().Unix(),
// 			Updated:  time.Now().Unix(),
// 		},
// 	}

// 	// 将测试数据序列化为 JSON
// 	jsonData, err := json.Marshal(testData)
// 	require.NoError(t, err)

// 	// 将数据存储到模拟客户端
// 	mockClient.store["tables/users.json"] = jsonData

// 	// 设置模拟行为
// 	mockClient.On("GetContents", mock.Anything, "owner", "repo", "tables/users.json", mock.Anything).Return(nil, nil, nil)

// 	// 测试加载表
// 	err = store.loadUsers()
// 	assert.NoError(t, err)

// 	// 验证数据是否正确加载
// 	user, err := store.Get("test-user-1")
// 	assert.NoError(t, err)
// 	assert.Equal(t, testData["test-user-1"], user)
// }

// func TestGitHubStore_SaveTable(t *testing.T) {
// 	store, mockClient, cleanup := setupGitHubTestStore(t)
// 	defer cleanup()

// 	// 设置模拟行为
// 	mockClient.On("GetContents", mock.Anything, "owner", "repo", "tables/users.json", mock.Anything).Return(nil, nil, nil)
// 	mockClient.On("CreateFile", mock.Anything, "owner", "repo", "tables/users.json", mock.Anything).Return(nil, nil, nil)

// 	// 创建测试用户
// 	user := User{
// 		ID:       "test-user-1",
// 		Username: "Test User",
// 		Email:    "test@example.com",
// 		Plan:     "free",
// 		Created:  time.Now().Unix(),
// 		Updated:  time.Now().Unix(),
// 	}

// 	// 保存用户
// 	err := store.Create(user)
// 	assert.NoError(t, err)

// 	// 验证数据是否正确保存
// 	savedData, ok := mockClient.store["tables/users.json"]
// 	assert.True(t, ok)

// 	var savedUsers map[string]User
// 	err = json.Unmarshal(savedData, &savedUsers)
// 	assert.NoError(t, err)

// 	savedUser, ok := savedUsers[user.ID]
// 	assert.True(t, ok)
// 	assert.Equal(t, user, savedUser)
// }
