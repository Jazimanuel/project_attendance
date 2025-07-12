CREATE TABLE students (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  email TEXT UNIQUE NOT NULL,
  created_at TIMESTAMP DEFAULT now()
);

CREATE TABLE courses (
  id SERIAL PRIMARY KEY,
  title TEXT NOT NULL,
  instructor TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT now()
);

CREATE TABLE admins (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  email TEXT UNIQUE NOT NULL,
  role TEXT DEFAULT 'admin',
  created_at TIMESTAMP DEFAULT now()
);

CREATE TABLE attendance (
  id SERIAL PRIMARY KEY,
  student_id INT REFERENCES students(id) ON DELETE CASCADE,
  course_id INT REFERENCES courses(id) ON DELETE CASCADE,
  check_in TIMESTAMP,
  check_out TIMESTAMP,
  aborted BOOLEAN DEFAULT false,
  created_at TIMESTAMP DEFAULT now()
);
