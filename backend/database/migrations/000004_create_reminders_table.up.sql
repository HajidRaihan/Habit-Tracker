CREATE TABLE reminders (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    habit_id UUID REFERENCES habits(id) ON DELETE CASCADE,
    reminder_time TIME NOT NULL, -- Waktu pengingat
    created_at TIMESTAMP DEFAULT now()
);
