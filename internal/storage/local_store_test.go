package storage

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/Axpz/store/internal/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupTestStore(t *testing.T) (*LocalStore, string, func()) {
	// 创建临时目录
	tempDir, err := os.MkdirTemp("/tmp", "store-test-*")
	require.NoError(t, err)

	fmt.Printf("temp dir %s", tempDir)

	// 创建配置
	cfg := &config.Config{
		Storage: config.StorageConfig{
			Path: tempDir,
		},
	}

	// 创建存储实例
	store, err := NewLocalStore(cfg)
	require.NoError(t, err)

	// 返回清理函数
	cleanup := func() {
		os.RemoveAll(tempDir)
	}

	return store.(*LocalStore), tempDir, cleanup
}

func TestLocalStore_UserCRUD(t *testing.T) {
	store, _, _ := setupTestStore(t)
	// defer cleanup()

	// 测试创建用户
	user := User{
		ID:       "test-user-1",
		Username: "Test User",
	}
	err := store.Create(user)
	assert.NoError(t, err)

	// 测试获取用户
	got, err := store.Get(user.ID)
	assert.NoError(t, err)
	assert.Equal(t, user, got)

	// // 测试更新用户
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

func TestLocalStore_CommentCRUD(t *testing.T) {
	store, _, cleanup := setupTestStore(t)
	defer cleanup()

	// 测试创建评论
	comment := Comment{
		ID:      "test-comment-1",
		Content: "Test Comment",
		UserID:  "test-user-1",
	}
	err := store.CreateComment(comment)
	assert.NoError(t, err)

	// 测试获取评论
	got, err := store.GetComment(comment.ID)
	assert.NoError(t, err)
	assert.Equal(t, comment, got)

	// 测试更新评论
	comment.Content = "Updated Comment"
	err = store.UpdateComment(comment)
	assert.NoError(t, err)

	got, err = store.GetComment(comment.ID)
	assert.NoError(t, err)
	assert.Equal(t, comment, got)

	// 测试删除评论
	err = store.DeleteComment(comment.ID)
	assert.NoError(t, err)

	_, err = store.GetComment(comment.ID)
	assert.Error(t, err)
}

func TestLocalStore_ConcurrentAccess(t *testing.T) {
	store, _, cleanup := setupTestStore(t)
	defer cleanup()

	// 创建测试用户
	user := User{
		ID:       "test-user-1",
		Username: "Test User",
	}
	err := store.Create(user)
	assert.NoError(t, err)

	// 并发读取和更新
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			// 读取用户
			_, err := store.Get(user.ID)
			assert.NoError(t, err)

			// 更新用户
			user.Username = "Updated User"
			err = store.Update(user)
			assert.NoError(t, err)

			done <- true
		}()
	}

	// 等待所有 goroutine 完成
	for i := 0; i < 10; i++ {
		<-done
	}
}

func TestLocalStore_FileOperations(t *testing.T) {
	store, tempDir, cleanup := setupTestStore(t)
	defer cleanup()

	// 测试文件创建
	user := User{
		ID:       "test-user-1",
		Username: "Test User",
	}
	err := store.Create(user)
	assert.NoError(t, err)

	// 验证文件是否存在
	userFile := filepath.Join(tempDir, "users")
	_, err = os.Stat(userFile)
	assert.NoError(t, err)

	// 测试文件内容
	content, err := os.ReadFile(userFile)
	assert.NoError(t, err)
	assert.Contains(t, string(content), user.ID)
	assert.Contains(t, string(content), user.Username)

	// 测试文件更新
	user.Username = "Updated User"
	err = store.Update(user)
	assert.NoError(t, err)

	content, err = os.ReadFile(userFile)
	assert.NoError(t, err)
	assert.Contains(t, string(content), "Updated User")
}

func TestLocalStore_ErrorCases(t *testing.T) {
	store, _, cleanup := setupTestStore(t)
	defer cleanup()

	// 测试获取不存在的用户
	_, err := store.Get("non-existent")
	assert.Error(t, err)

	// 测试更新不存在的用户
	err = store.Update(User{ID: "non-existent"})
	assert.Error(t, err)

	// 测试删除不存在的用户
	err = store.Delete("non-existent")
	assert.Error(t, err)

	// 测试创建重复用户
	user := User{
		ID:       "test-user-1",
		Username: "Test User",
	}
	err = store.Create(user)
	assert.NoError(t, err)

	err = store.Create(user)
	assert.Error(t, err)

	// 测试评论相关错误
	_, err = store.GetComment("non-existent")
	assert.Error(t, err)

	err = store.UpdateComment(Comment{ID: "non-existent"})
	assert.Error(t, err)

	err = store.DeleteComment("non-existent")
	assert.Error(t, err)

	comment := Comment{
		ID:      "test-comment-1",
		Content: "Test Comment",
		UserID:  "test-user-1",
	}
	err = store.CreateComment(comment)
	assert.NoError(t, err)

	err = store.CreateComment(comment)
	assert.Error(t, err)
}
