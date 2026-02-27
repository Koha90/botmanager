CREATE TABLE IF NOT EXISTS users (
  id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,

  tg_id BIGINT UNIQUE,
  tg_name TEXT,

  email TEXT UNIQUE,
  password_hash TEXT,

  role TEXT NOT NULL CHECK (role IN ('customer', 'admin')),
  balance BIGINT NOT NULL DEFAULT 0,

  is_enable BOOLEAN NOT NULL DEFAULT TRUE,

  admin_access_expires_at TIMESTAMP NULL,

  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_users_role ON users(role);
CREATE INDEX idx_users_is_enable ON users(is_enable);
