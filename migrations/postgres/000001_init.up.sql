CREATE TABLE users (
    -- Уникальный идентификатор пользователя (автоинкремент)
                       id SERIAL PRIMARY KEY,
    -- Имя пользователя (обязательное текстовое поле)
                       name TEXT NOT NULL,
    -- Уникальный логин (например, username или email)
                       login TEXT NOT NULL UNIQUE,
    -- Хеш пароля (никогда не храним голый пароль!)
                       pass TEXT NOT NULL,
    -- Дата и время создания записи, по умолчанию = текущему времени
                       created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    -- Статус пользователя (например: зарегистрирован или разлогинен)
    -- Ограничиваем значения только на 'register' или 'logout'
                       status TEXT NOT NULL CHECK (status IN ('register', 'logout'))
);
