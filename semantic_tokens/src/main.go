package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq" // postgres driver
)

// 常量定义
const (
	// 服务器配置
	ServerPort = ":8080"
	Timeout    = 30 * time.Second

	// 数据库配置
	MaxDBConnections = 10
	DBTimeout        = 5 * time.Second
)

// 用户角色枚举
type Role int

const (
	Admin Role = iota
	Editor
	Viewer
)

// String 将角色转换为字符串
func (r Role) String() string {
	return [...]string{"Admin", "Editor", "Viewer"}[r]
}

// User 表示系统中的用户
type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role      Role      `json:"role"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
}

// UserService 处理用户相关操作
type UserService struct {
	db  *sql.DB
	mu  sync.RWMutex
	log *log.Logger
}

// Error 类型定义
type ServiceError struct {
	Code    int
	Message string
	Err     error
}

// Error 实现 error 接口
func (e *ServiceError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

// NewUserService 创建新的用户服务
func NewUserService(db *sql.DB) *UserService {
	return &UserService{
		db:  db,
		mu:  sync.RWMutex{},
		log: log.New(os.Stdout, "[USER-SERVICE] ", log.LstdFlags),
	}
}

// GetUser 通过ID获取用户
func (s *UserService) GetUser(ctx context.Context, id string) (*User, error) {
	// 创建带超时的上下文
	ctx, cancel := context.WithTimeout(ctx, DBTimeout)
	defer cancel()

	// 准备SQL查询
	query := `SELECT id, name, email, role, active, created_at FROM users WHERE id = $1`

	var user User
	err := s.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Role,
		&user.Active,
		&user.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, &ServiceError{
				Code:    http.StatusNotFound,
				Message: "用户不存在",
				Err:     err,
			}
		}
		return nil, &ServiceError{
			Code:    http.StatusInternalServerError,
			Message: "获取用户失败",
			Err:     err,
		}
	}

	return &user, nil
}

// CreateUser 创建新用户
func (s *UserService) CreateUser(ctx context.Context, user *User) error {
	// 生成UUID
	if user.ID == "" {
		user.ID = uuid.New().String()
	}

	// 设置创建时间
	user.CreatedAt = time.Now().UTC()

	// 使用事务
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return &ServiceError{
			Code:    http.StatusInternalServerError,
			Message: "开始事务失败",
			Err:     err,
		}
	}
	defer tx.Rollback() // 如果提交成功，这将是无操作

	// 准备SQL语句
	query := `
		INSERT INTO users (id, name, email, role, active, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	// 执行查询
	_, err = tx.ExecContext(
		ctx,
		query,
		user.ID,
		user.Name,
		user.Email,
		user.Role,
		user.Active,
		user.CreatedAt,
	)

	if err != nil {
		return &ServiceError{
			Code:    http.StatusInternalServerError,
			Message: "创建用户失败",
			Err:     err,
		}
	}

	// 提交事务
	if err = tx.Commit(); err != nil {
		return &ServiceError{
			Code:    http.StatusInternalServerError,
			Message: "提交事务失败",
			Err:     err,
		}
	}

	s.log.Printf("用户创建成功: %s", user.ID)
	return nil
}

// UserHandler HTTP处理器
type UserHandler struct {
	service *UserService
}

// NewUserHandler 创建用户处理器
func NewUserHandler(service *UserService) *UserHandler {
	return &UserHandler{service: service}
}

// GetUser 处理获取用户请求
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	// 从URL获取用户ID
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "必须提供用户ID", http.StatusBadRequest)
		return
	}

	// 获取用户
	user, err := h.service.GetUser(r.Context(), id)
	if err != nil {
		if svcErr, ok := err.(*ServiceError); ok {
			http.Error(w, svcErr.Message, svcErr.Code)
		} else {
			http.Error(w, "服务器内部错误", http.StatusInternalServerError)
		}
		return
	}

	// 返回JSON响应
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "编码响应失败", http.StatusInternalServerError)
		return
	}
}

// 并发处理示例
func processConcurrently(users []User) {
	var wg sync.WaitGroup
	results := make(chan string, len(users))

	for _, user := range users {
		wg.Add(1)
		// 使用goroutine并发处理
		go func(u User) {
			defer wg.Done()

			// 模拟处理时间
			time.Sleep(100 * time.Millisecond)

			// 处理结果
			result := fmt.Sprintf("处理用户: %s 完成", u.Name)
			results <- result
		}(user)
	}

	// 等待所有goroutine完成
	go func() {
		wg.Wait()
		close(results)
	}()

	// 收集结果
	for result := range results {
		fmt.Println(result)
	}
}

// 泛型函数示例 (Go 1.18+)
func Filter[T any](slice []T, f func(T) bool) []T {
	var result []T
	for _, v := range slice {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	// 设置日志
	logger := log.New(os.Stdout, "[MAIN] ", log.LstdFlags)
	logger.Println("启动服务...")

	// 连接数据库
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		logger.Fatalf("连接数据库失败: %v", err)
	}
	defer db.Close()

	db.SetMaxOpenConns(MaxDBConnections)
	db.SetMaxIdleConns(MaxDBConnections / 2)
	db.SetConnMaxLifetime(time.Hour)

	// 测试数据库连接
	if err := db.Ping(); err != nil {
		logger.Fatalf("数据库Ping失败: %v", err)
	}
	logger.Println("数据库连接成功")

	// 创建服务和处理器
	userService := NewUserService(db)
	userHandler := NewUserHandler(userService)

	// 设置路由
	http.HandleFunc("/user", userHandler.GetUser)

	// 使用defer语句进行资源清理示例
	file, err := os.Create("server.log")
	if err != nil {
		logger.Fatalf("创建日志文件失败: %v", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			logger.Printf("关闭日志文件失败: %v", err)
		}
	}()

	// 使用goroutine启动服务器
	go func() {
		logger.Printf("HTTP服务器监听在%s\n", ServerPort)
		err := http.ListenAndServe(ServerPort, nil)
		if err != nil {
			logger.Fatalf("HTTP服务器启动失败: %v", err)
		}
	}()

	// 无限循环，保持主程序运行
	select {}
}
