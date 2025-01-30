CREATE TABLE habit_logs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    habit_id UUID REFERENCES habits(id) ON DELETE CASCADE,
    log_date DATE NOT NULL, -- Tanggal pencatatan
    progress INT NOT NULL DEFAULT 0, -- Jumlah yang telah dicapai oleh user
    status VARCHAR(100) NOT NULL, -- Status pencatatan
    created_at TIMESTAMP DEFAULT now()
);
