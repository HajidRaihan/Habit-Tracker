CREATE TABLE habits (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(id) NOT NULL,
    name VARCHAR(100) NOT NULL,
    goal VARCHAR(255) NOT NULL,
    time JSONB NOT NULL,
    created_at TIMESTAMP DEFAULT now()
)