-- ============================
-- 1. Пользователи
-- ============================
INSERT INTO users (name, password_hash, email) VALUES
                                                   ('alice', '$2a$10$testhashalice', 'alice@example.com'),
                                                   ('bob', '$2a$10$testhashbob', 'bob@example.com'),
                                                   ('charlie', '$2a$10$testhashcharlie', 'charlie@example.com'),
                                                   ('diana', '$2a$10$testhashdiana', 'diana@example.com'),
                                                   ('eve', '$2a$10$testhasheve', 'eve@example.com'),
                                                   ('frank', '$2a$10$testhashfrank', 'frank@example.com')
    ON CONFLICT DO NOTHING;

-- ============================
-- 2. Категории
-- ============================
INSERT INTO categories (title) VALUES
                                   ('Games'), ('Social'), ('Education'), ('Tools'),
                                   ('Music'), ('Video')
    ON CONFLICT DO NOTHING;

-- ============================
-- 3. Разработчики
-- ============================
INSERT INTO developers (name, password_hash, email) VALUES
                                                        ('SuperDev', '$2a$10$hashsuperdev', 'superdev@example.com'),
                                                        ('GameStudio', '$2a$10$hashgamestudio', 'games@example.com'),
                                                        ('EduSoft', '$2a$10$hashedusoft', 'edu@example.com'),
                                                        ('TechInnovate', '$2a$10$hashtechinnovate', 'tech@example.com')
    ON CONFLICT DO NOTHING;

-- ============================
-- 4. Приложения
-- ============================
INSERT INTO apps (title, description, size_mb, age_rating, downloads,
                  version, link_cloud, developer_id, category_id, created_at)
VALUES
    ('Space Shooter', 'Dynamic space shooter game', 150.25, '12+', 15000,
     '1.2.3', 'https://example.com/app/spaceshooter.apk',
     (SELECT id FROM developers WHERE name='GameStudio'),
     (SELECT id FROM categories WHERE title='Games'),
     NOW()),

    ('ChatWave', 'Modern messenger for everyone', 85.50, '3+', 800000,
     '3.0.1', 'https://example.com/app/chatwave.apk',
     (SELECT id FROM developers WHERE name='SuperDev'),
     (SELECT id FROM categories WHERE title='Social'),
     NOW()),

    ('MathMaster', 'Learn math with interactive lessons', 200.00, '0+', 42000,
     '2.1.0', 'https://example.com/app/mathmaster.apk',
     (SELECT id FROM developers WHERE name='EduSoft'),
     (SELECT id FROM categories WHERE title='Education'),
     NOW()),

    ('CleanPro', 'Phone cleaning tool', 40.10, '3+', 300000,
     '1.0.2', 'https://example.com/app/cleanpro.apk',
     (SELECT id FROM developers WHERE name='SuperDev'),
     (SELECT id FROM categories WHERE title='Tools'),
     NOW()),

    ('PhotoEdit Pro', 'Professional photo editing suite', 95.75, '4+', 250000,
     '4.2.1', 'https://example.com/app/photoeditpro.apk',
     (SELECT id FROM developers WHERE name='GameStudio'),
     (SELECT id FROM categories WHERE title='Music'),
     NOW()),

    ('Fitness Tracker', 'Track your workouts and progress', 65.30, '12+', 180000,
     '2.5.0', 'https://example.com/app/fitnesstracker.apk',
     (SELECT id FROM developers WHERE name='EduSoft'),
     (SELECT id FROM categories WHERE title='Video'),
     NOW()),

    ('MusicFlow', 'Stream millions of songs', 120.40, '12+', 950000,
     '5.1.3', 'https://example.com/app/musicflow.apk',
     (SELECT id FROM developers WHERE name='SuperDev'),
     (SELECT id FROM categories WHERE title='Education'),
     NOW()),

    ('Budget Planner', 'Manage your finances easily', 35.20, '3+', 75000,
     '1.8.2', 'https://example.com/app/budgetplanner.apk',
     (SELECT id FROM developers WHERE name='TechInnovate'),
     (SELECT id FROM categories WHERE title='Tools'),
     NOW()),

    ('Language Learn', 'Learn new languages fast', 180.90, '0+', 320000,
     '3.2.4', 'https://example.com/app/languagelearn.apk',
     (SELECT id FROM developers WHERE name='GameStudio'),
     (SELECT id FROM categories WHERE title='Education'),
     NOW()),

    ('Weather Globe', 'Accurate weather forecasts', 28.60, '3+', 500000,
     '2.0.1', 'https://example.com/app/weatherglobe.apk',
     (SELECT id FROM developers WHERE name='EduSoft'),
     (SELECT id FROM categories WHERE title='Video'),
     NOW()),

    ('Recipe Book', 'Thousands of cooking recipes', 75.80, '4+', 150000,
     '1.5.7', 'https://example.com/app/recipebook.apk',
     (SELECT id FROM developers WHERE name='TechInnovate'),
     (SELECT id FROM categories WHERE title='Music'),
     NOW()),

    ('Puzzle Quest', 'Challenging puzzle adventure', 210.15, '6+', 89000,
     '1.3.0', 'https://example.com/app/puzzlequest.apk',
     (SELECT id FROM developers WHERE name='SuperDev'),
     (SELECT id FROM categories WHERE title='Games'),
     NOW()),

    ('Sleep Well', 'Sleep tracking and meditation', 42.25, '3+', 280000,
     '2.7.1', 'https://example.com/app/sleepwell.apk',
     (SELECT id FROM developers WHERE name='GameStudio'),
     (SELECT id FROM categories WHERE title='Video'),
     NOW()),

    ('News Daily', 'Personalized news aggregator', 55.70, '12+', 420000,
     '3.1.8', 'https://example.com/app/newsdaily.apk',
     (SELECT id FROM developers WHERE name='EduSoft'),
     (SELECT id FROM categories WHERE title='Social'),
     NOW()),

    ('Shopping Cart', 'Online shopping assistant', 88.45, '3+', 670000,
     '4.0.2', 'https://example.com/app/shoppingcart.apk',
     (SELECT id FROM developers WHERE name='TechInnovate'),
     (SELECT id FROM categories WHERE title='Tools'),
     NOW())
        ,
    ('MindRelax', 'Meditation and breathing exercises', 52.30, '3+', 210000,
     '1.9.0', 'https://example.com/app/mindrelax.apk',
     (SELECT id FROM developers WHERE name='SuperDev'),
     (SELECT id FROM categories WHERE title='Health'),
     NOW() - INTERVAL '1 day'),

    ('CryptoWatch', 'Track your cryptocurrency portfolio', 33.80, '12+', 540000,
     '4.3.7', 'https://example.com/app/cryptowatch.apk',
     (SELECT id FROM developers WHERE name='TechInnovate'),
     (SELECT id FROM categories WHERE title='Tools'),
     NOW() - INTERVAL '2 days'),

    ('FlashCards Pro', 'Study with smart flashcards', 120.20, '0+', 88000,
     '2.0.4', 'https://example.com/app/flashcardspro.apk',
     (SELECT id FROM developers WHERE name='EduSoft'),
     (SELECT id FROM categories WHERE title='Education'),
     NOW() - INTERVAL '5 days'),

    ('Galaxy Racer', 'High-speed space racing game', 250.75, '6+', 190000,
     '1.1.0', 'https://example.com/app/galaxyracer.apk',
     (SELECT id FROM developers WHERE name='GameStudio'),
     (SELECT id FROM categories WHERE title='Games'),
     NOW() - INTERVAL '7 days'),

    ('Task Organizer', 'Daily planner and to-do list', 28.50, '3+', 72000,
     '1.6.2', 'https://example.com/app/taskorganizer.apk',
     (SELECT id FROM developers WHERE name='TechInnovate'),
     (SELECT id FROM categories WHERE title='Tools'),
     NOW() - INTERVAL '8 days'),

    ('VideoMaster', 'Professional video editing tool', 310.10, '12+', 305000,
     '3.4.1', 'https://example.com/app/videomaster.apk',
     (SELECT id FROM developers WHERE name='GameStudio'),
     (SELECT id FROM categories WHERE title='Video'),
     NOW() - INTERVAL '10 days'),

    ('EcoTracker', 'Reduce your carbon footprint', 46.25, '3+', 53000,
     '1.2.1', 'https://example.com/app/ecotracker.apk',
     (SELECT id FROM developers WHERE name='SuperDev'),
     (SELECT id FROM categories WHERE title='Tools'),
     NOW() - INTERVAL '12 days'),

    ('StudyMate', 'AI-based study assistant', 140.35, '0+', 91000,
     '2.2.0', 'https://example.com/app/studymate.apk',
     (SELECT id FROM developers WHERE name='EduSoft'),
     (SELECT id FROM categories WHERE title='Education'),
     NOW() - INTERVAL '15 days'),

    ('RPG Legends', 'Epic RPG adventure game', 420.50, '12+', 760000,
     '1.7.3', 'https://example.com/app/rpglegends.apk',
     (SELECT id FROM developers WHERE name='GameStudio'),
     (SELECT id FROM categories WHERE title='Games'),
     NOW() - INTERVAL '17 days'),

    ('PhotoLite', 'Lightweight photo editor', 33.10, '3+', 128000,
     '1.0.5', 'https://example.com/app/photolite.apk',
     (SELECT id FROM developers WHERE name='TechInnovate'),
     (SELECT id FROM categories WHERE title='Music'),
     NOW() - INTERVAL '20 days'),

    ('Language Pro', 'Advanced language learning', 185.80, '0+', 280000,
     '3.5.1', 'https://example.com/app/languagepro.apk',
     (SELECT id FROM developers WHERE name='EduSoft'),
     (SELECT id FROM categories WHERE title='Education'),
     NOW() - INTERVAL '22 days'),

    ('CalorieWatch', 'Track calories and diet', 55.30, '3+', 67000,
     '2.0.3', 'https://example.com/app/caloriewatch.apk',
     (SELECT id FROM developers WHERE name='SuperDev'),
     (SELECT id FROM categories WHERE title='Tools'),
     NOW() - INTERVAL '30 days'),

    ('NewsWorld', 'International news aggregator', 89.40, '12+', 430000,
     '2.9.0', 'https://example.com/app/newsworld.apk',
     (SELECT id FROM developers WHERE name='TechInnovate'),
     (SELECT id FROM categories WHERE title='Social'),
     NOW() - INTERVAL '31 days'),

    ('MusicBeat', 'Find trending DJ tracks', 130.25, '6+', 520000,
     '1.8.7', 'https://example.com/app/musicbeat.apk',
     (SELECT id FROM developers WHERE name='GameStudio'),
     (SELECT id FROM categories WHERE title='Music'),
     NOW() - INTERVAL '33 days'),

    ('ForestRun', 'Nature-themed endless runner', 95.65, '3+', 150000,
     '1.3.1', 'https://example.com/app/forestrun.apk',
     (SELECT id FROM developers WHERE name='GameStudio'),
     (SELECT id FROM categories WHERE title='Games'),
     NOW() - INTERVAL '40 days'),

    ('BudgetEasy', 'Simple finance tracking app', 39.10, '3+', 42000,
     '1.4.4', 'https://example.com/app/budgeteasy.apk',
     (SELECT id FROM developers WHERE name='TechInnovate'),
     (SELECT id FROM categories WHERE title='Tools'),
     NOW() - INTERVAL '45 days'),

    ('SleepTracker+', 'Advanced sleep monitoring', 110.50, '3+', 120000,
     '2.8.0', 'https://example.com/app/sleeptracker.apk',
     (SELECT id FROM developers WHERE name='SuperDev'),
     (SELECT id FROM categories WHERE title='Video'),
     NOW() - INTERVAL '48 days'),

    ('PuzzleWorld', 'Colorful puzzle challenges', 205.30, '6+', 60000,
     '1.2.0', 'https://example.com/app/puzzleworld.apk',
     (SELECT id FROM developers WHERE name='GameStudio'),
     (SELECT id FROM categories WHERE title='Games'),
     NOW() - INTERVAL '50 days'),

    ('TravelGuide', 'Explore world attractions', 140.40, '3+', 310000,
     '2.1.7', 'https://example.com/app/travelguide.apk',
     (SELECT id FROM developers WHERE name='EduSoft'),
     (SELECT id FROM categories WHERE title='Education'),
     NOW() - INTERVAL '60 days'),

    ('VoiceType', 'AI-powered speech-to-text', 85.75, '3+', 210000,
     '4.0.1', 'https://example.com/app/voicetype.apk',
     (SELECT id FROM developers WHERE name='TechInnovate'),
     (SELECT id FROM categories WHERE title='Tools'),
     NOW() - INTERVAL '65 days')

    ON CONFLICT DO NOTHING;

-- ============================
-- 5. Отзывы
-- ============================
INSERT INTO reviews (user_id, app_id, score, comment)
VALUES
    ((SELECT id FROM users WHERE name='alice'),
     (SELECT id FROM apps WHERE title='Space Shooter'),
     5, 'Awesome gameplay!'),

    ((SELECT id FROM users WHERE name='bob'),
     (SELECT id FROM apps WHERE title='Space Shooter'),
     4, 'Good but difficult'),

    ((SELECT id FROM users WHERE name='alice'),
     (SELECT id FROM apps WHERE title='ChatWave'),
     5, 'Use it every day'),

    ((SELECT id FROM users WHERE name='charlie'),
     (SELECT id FROM apps WHERE title='MathMaster'),
     3, 'Helpful for studying'),

    ((SELECT id FROM users WHERE name='bob'),
     (SELECT id FROM apps WHERE title='CleanPro'),
     4, 'Good cleaning tool!'),

    ((SELECT id FROM users WHERE name='diana'),
     (SELECT id FROM apps WHERE title='PhotoEdit Pro'),
     5, 'Best photo editor ever!'),

    ((SELECT id FROM users WHERE name='eve'),
     (SELECT id FROM apps WHERE title='PhotoEdit Pro'),
     4, 'Great filters and tools'),

    ((SELECT id FROM users WHERE name='frank'),
     (SELECT id FROM apps WHERE title='PhotoEdit Pro'),
     3, 'Could use more features')
    ON CONFLICT DO NOTHING;
