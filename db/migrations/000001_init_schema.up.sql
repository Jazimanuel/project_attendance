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
  student_id INT REFERENCES students(id) ON DELETE CASCADE NOT NULL,
  course_id INT REFERENCES courses(id) ON DELETE CASCADE NOT NULL,
  check_in TIMESTAMP NOT NULL DEFAULT now(),
  check_out TIMESTAMP,
  status varchar(20) NOT NULL ,
  aborted BOOLEAN DEFAULT false,
  created_at TIMESTAMP DEFAULT now()
);
