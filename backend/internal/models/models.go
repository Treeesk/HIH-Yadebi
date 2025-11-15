package models

import "time"

// ========================
// USERS
// ========================
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"login"`
	Password  string    `json:"password"` // хранится хеш!
	Email     string    `json:"email"`
	// CreatedAt time.Time `json:"created_at"`
}

// ========================
// CATEGORIES
// ========================
type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// ========================
// APPS
// ========================
type App struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	SizeMB      float64   `json:"size_mb"`
	AgeRating   string    `json:"age_rating"`
	Downloads   int64     `json:"downloads"`
	// Rating      float64   `json:"rating"`
	Version     string    `json:"version"`
	LinkApk     string    `json:"link_apk"`
	IconSmall   string    `json:"icon_small"`   // ссылка на маленькую иконку
	IconLarge   string    `json:"icon_large"`   // ссылка на большую иконку
	Screenshots string    `json:"screenshots"`  // JSON-массив ссылок на скриншоты

	DeveloperID int       `json:"developer_id"`
	CategoryID  int       `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"` // дата и время добавления приложения
}

// ========================
// REVIEWS
// ========================
type Review struct {
	ID      int    `json:"id"`
	UserID  int    `json:"user_id"`
	AppID   int    `json:"app_id"`
	Score   int    `json:"score"`
	Comment string `json:"comment"`
	// Helpful int    `json:"helpful"`
}

// ========================
// USER + APP (installed)
// ========================
type UserApp struct {
	ID     int `json:"id"`
	UserID int `json:"user_id"`
	AppID  int `json:"app_id"`
}

type Developer struct {
	ID        int       `json:"id"`
	Name      string    `json:"login"`
	Password  string    `json:"password"` // хранится хеш!
	Email     string    `json:"email"`
	// CreatedAt time.Time `json:"created_at"`
}