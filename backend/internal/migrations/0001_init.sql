-- =========================
-- 1. Таблица: пользователи
-- =========================
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE
);

-- =========================
-- 2. Таблица: категории
-- =========================
CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL UNIQUE
);

-- =========================
-- 3. Таблица: приложения
-- =========================
CREATE TABLE apps (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    size_mb NUMERIC(10,2),
    age_rating VARCHAR(50),
    downloads BIGINT DEFAULT 0,
    -- rating NUMERIC(3,2) DEFAULT 0,
    version VARCHAR(50) NOT NULL DEFAULT '1.0.0',
    link_apk TEXT,
    developer_id INTEGER NOT NULL REFERENCES developer(id) ON DELETE CASCADE,
    category_id INTEGER REFERENCES categories(id) ON DELETE SET NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()  -- дата и время добавления приложения
);

-- =========================
-- 4. Таблица: отзывы
-- =========================
CREATE TABLE reviews (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    app_id INTEGER NOT NULL REFERENCES apps(id) ON DELETE CASCADE,
    score INTEGER NOT NULL CHECK (score BETWEEN 1 AND 5),
    comment TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW() 
    -- helpful INTEGER DEFAULT 0,
    UNIQUE(user_id, app_id) -- один отзыв на приложение от пользователя
);

-- ========================================
-- 5. Таблица: пользователь + приложение
--    (например: установленные приложения)
-- ========================================
CREATE TABLE user_apps (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    app_id INTEGER NOT NULL REFERENCES apps(id) ON DELETE CASCADE,
    UNIQUE(user_id, app_id) -- пользователь не может иметь одно приложение 2 раза
);

CREATE TABLE developers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE
);