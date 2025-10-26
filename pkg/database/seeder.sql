USE habit_tracker;

-- ============================================================
-- USERS
-- ============================================================
INSERT INTO users (name, email, password, created_at, updated_at)
VALUES
('Andi Pratama', 'andi@example.com', '$2a$10$UGll22w/xT/VFTDrYJFW2u0ZywN6kMvQJlV5NBoaMl5SKVficLw0O', NOW(), NOW()),
('Budi Santoso', 'budi@example.com', '$2a$10$UGll22w/xT/VFTDrYJFW2u0ZywN6kMvQJlV5NBoaMl5SKVficLw0O', NOW(), NOW()),
('Citra Lestari', 'citra@example.com', '$2a$10$UGll22w/xT/VFTDrYJFW2u0ZywN6kMvQJlV5NBoaMl5SKVficLw0O', NOW(), NOW());

-- ============================================================
-- HABITS
-- ============================================================
INSERT INTO habits (user_id, title, description, created_at, updated_at)
VALUES
(1, 'Olahraga pagi', 'Lari pagi minimal 20 menit', NOW(), NOW()),
(1, 'Membaca buku', 'Baca minimal 10 halaman buku', NOW(), NOW()),
(2, 'Meditasi', 'Meditasi 10 menit setelah bangun tidur', NOW(), NOW()),
(3, 'Belajar coding', 'Latihan coding minimal 1 jam', NOW(), NOW());

-- ============================================================
-- HABIT LOGS
-- ============================================================
INSERT INTO habit_logs (habit_id, user_id, log_date, status, note, created_at, updated_at)
VALUES
(1, 1, '2025-10-23', 'done', 'Lari di taman', NOW(), NOW()),
(1, 1, '2025-10-24', 'done', 'Cuaca cerah, lari 25 menit', NOW(), NOW()),
(2, 1, '2025-10-24', 'missed', 'Kelelahan', NOW(), NOW()),
(3, 2, '2025-10-24', 'done', 'Meditasi 15 menit', NOW(), NOW()),
(4, 3, '2025-10-24', 'done', 'Belajar GoLang di YouTube', NOW(), NOW());

-- ============================================================
-- SLEEP LOGS
-- ============================================================
INSERT INTO sleep_logs (user_id, sleep_start, sleep_end, created_at, updated_at)
VALUES
(1, '2025-10-23 22:30:00', '2025-10-24 06:30:00', NOW(), NOW()),
(2, '2025-10-23 23:00:00', '2025-10-24 07:00:00', NOW(), NOW()),
(3, '2025-10-23 21:45:00', '2025-10-24 05:30:00', NOW(), NOW());

-- ============================================================
-- ACTIVITIES
-- ============================================================
INSERT INTO activities (user_id, activity_date, title, duration_minutes, notes, created_at, updated_at)
VALUES
(1, '2025-10-24', 'Bekerja di kantor', 480, 'Meeting dan laporan', NOW(), NOW()),
(1, '2025-10-24', 'Nonton film', 120, 'Film dokumenter', NOW(), NOW()),
(2, '2025-10-24', 'Mengajar kelas online', 180, 'Materi Python dasar', NOW(), NOW()),
(3, '2025-10-24', 'Mengerjakan proyek freelance', 240, 'Fixing bug UI', NOW(), NOW());

-- ============================================================
-- DAILY STORIES
-- ============================================================
INSERT INTO daily_stories (user_id, story_date, story_text, mood, created_at, updated_at)
VALUES
(1, '2025-10-24', 'Hari ini produktif dan senang sekali!', 'happy', NOW(), NOW()),
(2, '2025-10-24', 'Capek tapi puas dengan hasil kerja.', 'neutral', NOW(), NOW()),
(3, '2025-10-24', 'Sedikit stres karena banyak deadline.', 'sad', NOW(), NOW());

-- ============================================================
-- REMINDERS
-- ============================================================
INSERT INTO reminders (user_id, habit_id, reminder_time, message, is_active, created_at, updated_at)
VALUES
(1, 1, '07:00:00', 'Saatnya olahraga pagi!', TRUE, NOW(), NOW()),
(1, 2, '20:30:00', 'Jangan lupa baca buku sebelum tidur', TRUE, NOW(), NOW()),
(2, 3, '06:30:00', 'Meditasi pagi sebentar yuk', TRUE, NOW(), NOW()),
(3, 4, '21:00:00', 'Ayo lanjut belajar coding!', TRUE, NOW(), NOW());
