DROP TABLE users;
CREATE TABLE users(
  id SERIAL PRIMARY KEY,
  email VARCHAR(100) NOT NULL UNIQUE,
  name VARCHAR(100) NOT NULL,
  dob DATE NOT NULL
);


INSERT INTO users (email, name, dob) VALUES ('mozarikfahrozi@gmail.com', 'zein', '11-04-2000');

SELECT * FROM users;
