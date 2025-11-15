-- --------------------------------
-- 1. Пользователи
-- --------------------------------
INSERT INTO users (name, password_hash, email)
VALUES
    ('alice', '$2a$10$testhashalice', 'alice@example.com'),
    ('bob',   '$2a$10$testhashbob',   'bob@example.com'),
    ('charlie', '$2a$10$testhashcharlie', 'charlie@example.com'),
    ('diana', '$2a$10$testhashdiana', 'diana@example.com'),
    ('eve', '$2a$10$testhasheve', 'eve@example.com'),
    ('frank', '$2a$10$testhashfrank', 'frank@example.com')
ON CONFLICT (name) DO NOTHING;
-- --------------------------------
-- 2. Категории
-- --------------------------------
INSERT INTO categories (title)
VALUES
    ('Games'),
    ('Social'),
    ('Education'),
    ('Tools'),
    ('Music'),
    ('Video')
ON CONFLICT (title) DO NOTHING;

-- --------------------------------
-- 3. Разработчики
-- --------------------------------
INSERT INTO developers (name, password_hash, email)
VALUES
    ('SuperDev', '$2a$10$hashsuperdev', 'superdev@example.com'),
    ('GameStudio', '$2a$10$hashgamestudio', 'games@example.com'),
    ('EduSoft', '$2a$10$hashedusoft', 'edu@example.com'),
    ('TechInnovate', '$2a$10$hashtechinnovate', 'tech@example.com')
ON CONFLICT (name) DO NOTHING;

-- --------------------------------
-- 4. Приложения
-- --------------------------------
INSERT INTO apps (title, description, size_mb, age_rating, downloads,
                  version, link_cloud, developer_id, category_id, created_at)
VALUES
    ('Space Shooter', 'Dynamic space shooter game', 150.25, '12+', 15000,
     '1.2.3', 'https://example.com/app/spaceshooter.apk', 2, 1, NOW()),

    ('ChatWave', 'Modern messenger for everyone', 85.50, '3+', 800000,
     '3.0.1', 'https://example.com/app/chatwave.apk', 1, 2, NOW()),

    ('MathMaster', 'Learn math with interactive lessons', 200.00, '0+', 42000,
     '2.1.0', 'https://example.com/app/mathmaster.apk', 3, 3, NOW()),

    ('CleanPro', 'Phone cleaning tool', 40.10, '3+', 300000,
     '1.0.2', 'https://example.com/app/cleanpro.apk', 1, 4, NOW()),

    ('PhotoEdit Pro', 'Professional photo editing suite', 95.75, '4+', 250000,
     '4.2.1', 'https://example.com/app/photoeditpro.apk', 2, 5, NOW()),

    ('Fitness Tracker', 'Track your workouts and progress', 65.30, '12+', 180000,
     '2.5.0', 'https://example.com/app/fitnesstracker.apk', 3, 6, NOW()),

    ('MusicFlow', 'Stream millions of songs', 120.40, '12+', 950000,
     '5.1.3', 'https://example.com/app/musicflow.apk', 1, 7, NOW()),

    ('Budget Planner', 'Manage your finances easily', 35.20, '3+', 75000,
     '1.8.2', 'https://example.com/app/budgetplanner.apk', 4, 8, NOW()),

    ('Language Learn', 'Learn new languages fast', 180.90, '0+', 320000,
     '3.2.4', 'https://example.com/app/languagelearn.apk', 2, 3, NOW()),

    ('Weather Globe', 'Accurate weather forecasts', 28.60, '3+', 500000,
     '2.0.1', 'https://example.com/app/weatherglobe.apk', 3, 9, NOW()),

    ('Recipe Book', 'Thousands of cooking recipes', 75.80, '4+', 150000,
     '1.5.7', 'https://example.com/app/recipebook.apk', 4, 10, NOW()),

    ('Puzzle Quest', 'Challenging puzzle adventure', 210.15, '6+', 89000,
     '1.3.0', 'https://example.com/app/puzzlequest.apk', 1, 1, NOW()),

    ('Sleep Well', 'Sleep tracking and meditation', 42.25, '3+', 280000,
     '2.7.1', 'https://example.com/app/sleepwell.apk', 2, 6, NOW()),

    ('News Daily', 'Personalized news aggregator', 55.70, '12+', 420000,
     '3.1.8', 'https://example.com/app/newsdaily.apk', 3, 11, NOW()),

    ('Shopping Cart', 'Online shopping assistant', 88.45, '3+', 670000,
     '4.0.2', 'https://example.com/app/shoppingcart.apk', 4, 12, NOW())
ON CONFLICT (title) DO NOTHING;

-- --------------------------------
-- 5. Отзывы
-- --------------------------------
INSERT INTO reviews (user_id, app_id, score, comment)
VALUES
    (1, 1, 5, 'Awesome gameplay!'),
    (2, 1, 4, 'Good but difficult'),
    (1, 2, 5, 'Use it every day'),
    (3, 3, 3, 'Helpful for studying'),
    (2, 4, 4, 'Good cleaning tool!'),
    
    (4, 5, 5, 'Best photo editor ever!'),
    (5, 5, 4, 'Great filters and tools'),
    (6, 5, 3, 'Could use more features'),
    
    (1, 6, 5, 'Perfect for tracking workouts'),
    (3, 6, 4, 'Helped me lose 10kg!'),
    
    (2, 7, 5, 'Huge music library'),
    (4, 7, 2, 'Too many ads'),
    
    (5, 8, 4, 'Very useful for budgeting'),
    (6, 8, 5, 'Saved me so much money'),
    
    (1, 9, 5, 'Learning Spanish was easy'),
    (3, 9, 4, 'Good but pronunciation could be better'),
    
    (2, 10, 5, 'Accurate weather forecasts'),
    (4, 10, 3, 'Sometimes inaccurate'),
    
    (5, 11, 4, 'Great recipes for beginners'),
    (6, 11, 5, 'Cook delicious meals every day'),
    
    (1, 12, 5, 'Addictive puzzle game'),
    (3, 12, 4, 'Challenging but fun'),
    
    (2, 13, 5, 'Improved my sleep quality'),
    (4, 13, 4, 'Nice meditation exercises'),
    
    (5, 14, 3, 'Good news selection'),
    (6, 14, 4, 'Personalized feed works well'),
    
    (1, 15, 5, 'Shopping made easy'),
    (3, 15, 4, 'Great price comparisons')
ON CONFLICT DO NOTHING;

-- --------------------------------
-- 6. Пользовательские приложения
-- --------------------------------
INSERT INTO user_apps (user_id, app_id)
VALUES
    (1, 1),
    (1, 2),
    (2, 2),
    (3, 3),
    (2, 4)
ON CONFLICT DO NOTHING;
