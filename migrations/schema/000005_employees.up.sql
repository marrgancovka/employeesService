CREATE TABLE IF NOT EXISTS employees (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    surname TEXT NOT NULL,
    phone TEXT UNIQUE NOT NULL,
    company_id INTEGER REFERENCES companies(id) ON DELETE CASCADE NOT NULL,
    department_id INTEGER REFERENCES departments(id) ON DELETE SET NULL,
    passport_type TEXT NOT NULL,
    passport_number TEXT UNIQUE NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
)