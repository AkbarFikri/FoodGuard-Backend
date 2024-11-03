CREATE TABLE nutritions (
    id VARCHAR(26) PRIMARY KEY,
    user_id VARCHAR(26) NOT NULL REFERENCES users(id),
    name VARCHAR(255) NOT NULL,
    type VARCHAR(100) NOT NULL ,
    score REAL NOT NULL ,
    carbohydrate DOUBLE PRECISION NOT NULL ,
    sugar DOUBLE PRECISION NOT NULL ,
    calorie DOUBLE PRECISION NOT NULL ,
    fat DOUBLE PRECISION NOT NULL ,
    protein DOUBLE PRECISION NOT NULL ,
    recommendation TEXT NOT NULL ,
    created_at TIMESTAMP NOT NULL
);