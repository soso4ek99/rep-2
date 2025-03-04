CREATE TABLE tasks (
                       id SERIAL PRIMARY KEY,
                       task VARCHAR(255) NOT NULL,
                       is_done BOOLEAN DEFAULT FALSE,
                        user_id INTEGER,
                       /*user_id INTEGER REFERENCES users(id),*/
                       created_at TIMESTAMP NOT NULL DEFAULT NOW(),
                       updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
                       deleted_at TIMESTAMP DEFAULT NULL
);