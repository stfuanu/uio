```bash
# To set up a local MySQL server with a database named auth and a table named user on port 3306, you can follow these steps:

# Install MySQL server on your local machine. You can use the following command on Ubuntu or Debian-based Linux distributions:

sudo apt-get update
sudo apt-get install mysql-server
Once the MySQL server is installed, start the server using the following command:


sudo systemctl start mysql
# Log in to the MySQL server as the root user:

sudo mysql -u root
# Create a new database named auth:

CREATE DATABASE auth;
USE auth;
```

```mysql
CREATE TABLE user (
  id int unsigned NOT NULL AUTO_INCREMENT,
  first_name varchar(50) NOT NULL,
  last_name varchar(50) NOT NULL,
  email varchar(100) NOT NULL UNIQUE,
  password char(60) NOT NULL,
  status_id tinyint unsigned NOT NULL DEFAULT 1,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted tinyint unsigned NOT NULL DEFAULT 0,
  PRIMARY KEY (id),
  KEY status_id (status_id)
);

```

```mysql
CREATE USER 'adminuser'@'localhost' IDENTIFIED BY 'password123';
GRANT ALL PRIVILEGES ON auth.* TO 'adminuser'@'localhost';

```
