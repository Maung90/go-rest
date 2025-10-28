-- ============================================================
-- DATABASE: habit_tracker
-- ============================================================
CREATE DATABASE IF NOT EXISTS habit_tracker;
USE habit_tracker;

-- ============================================================
-- TABLE: users
-- ============================================================
CREATE TABLE users (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(150) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- ============================================================
-- TABLE: refresh_tokens
-- ============================================================
CREATE TABLE refresh_tokens (
    id VARCHAR(36) PRIMARY KEY,
    user_id BIGINT UNSIGNED NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- ============================================================
-- TABLE: habits
-- ============================================================
CREATE TABLE habits (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT UNSIGNED NOT NULL,
    title VARCHAR(150) NOT NULL,
    description TEXT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- ============================================================
-- TABLE: habit_logs
-- (menandai habit dilakukan di tanggal tertentu)
-- ============================================================
CREATE TABLE habit_logs (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    habit_id BIGINT UNSIGNED NOT NULL,
    user_id BIGINT UNSIGNED NOT NULL,
    log_date DATE NOT NULL,
    status ENUM('done','missed') DEFAULT 'done',
    note VARCHAR(255) NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (habit_id) REFERENCES habits(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- ============================================================
-- TABLE: sleep_logs
-- ============================================================
CREATE TABLE sleep_logs (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT UNSIGNED NOT NULL,
    sleep_start DATETIME NOT NULL,
    sleep_end DATETIME NOT NULL,
    duration_hours DECIMAL(4,2) GENERATED ALWAYS AS (TIMESTAMPDIFF(MINUTE, sleep_start, sleep_end) / 60) STORED,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- ============================================================
-- TABLE: activities
-- ============================================================
CREATE TABLE activities (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT UNSIGNED NOT NULL,
    activity_date DATE NOT NULL,
    title VARCHAR(150) NOT NULL,
    duration_minutes INT DEFAULT 0,
    notes TEXT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- ============================================================
-- TABLE: daily_stories
-- ============================================================
CREATE TABLE daily_stories (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT UNSIGNED NOT NULL,
    story_date DATE NOT NULL,
    story_text VARCHAR(255) NOT NULL,
    mood ENUM('happy','neutral','sad','angry','excited') DEFAULT 'neutral',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- ============================================================
-- TABLE: reminders (opsional)
-- ============================================================
CREATE TABLE reminders (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT UNSIGNED NOT NULL,
    habit_id BIGINT UNSIGNED NULL,
    reminder_time TIME NOT NULL,
    message VARCHAR(255) DEFAULT 'Jangan lupa lakukan habit kamu!',
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (habit_id) REFERENCES habits(id) ON DELETE CASCADE
);
-- ============================================================
-- VIEW: daily_summary 
-- ============================================================
CREATE OR REPLACE VIEW daily_summary AS
SELECT 
    u.id AS user_id,
    u.name AS user_name,
    d.story_date AS date,
    d.story_text,
    d.mood,
    COUNT(DISTINCT hl.id) AS total_habits_done,
    COUNT(DISTINCT a.id) AS total_activities,
    ROUND(AVG(sl.duration_hours), 2) AS avg_sleep_hours
FROM users u
LEFT JOIN daily_stories d ON u.id = d.user_id
LEFT JOIN habit_logs hl ON hl.user_id = u.id AND hl.log_date = d.story_date
LEFT JOIN activities a ON a.user_id = u.id AND a.activity_date = d.story_date
LEFT JOIN sleep_logs sl ON sl.user_id = u.id AND DATE(sl.sleep_start) = d.story_date
GROUP BY u.id, d.story_date, d.story_text, d.mood;

