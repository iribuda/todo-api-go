CREATE TABLE category (
  categoryId int PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL
);

CREATE TABLE task (
  taskId int PRIMARY KEY NOT NULL AUTO_INCREMENT,
  title varchar(100) NOT NULL,
  text varchar(255) DEFAULT NULL,
  deadline date DEFAULT NULL,
  categoryId int NOT NULL,
  done tinyint(1) NOT NULL DEFAULT '0',

  FOREIGN KEY (categoryId) REFERENCES category(categoryId) ON DELETE CASCADE
);

CREATE TABLE `user` (
  `userId` int PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `username` varchar(50) NOT NULL,
  `email` varchar(100) NOT NULL,
  `password` varchar(255) NOT NULL
) ;

CREATE TABLE `user_task` (
  `userId` int NOT NULL,
  `taskId` int NOT NULL,

  FOREIGN KEY (userId) REFERENCES user(userId) ON DELETE CASCADE,
  FOREIGN KEY (taskId) REFERENCES task(taskId) ON DELETE CASCADE
) ;

INSERT INTO `category` (`name`) VALUES
('Projekt'),
('Studium'),
('Persönlich');

INSERT INTO `task` (`title`, `text`, `deadline`, `categoryId`, `done`) VALUES
('Datenbank erstellen', 'Datenbank soll erstellt und ausgefüllt werden', '2024-07-03', 1, 1),
('API entwickeln', 'Alle sollen implementiert werden', '2024-07-03', 1, 0),
('Prüfungsvorbereitung', 'Vorbereitung für taktisches Informationsmanagement', '2024-07-03', 2, 0);

INSERT INTO `user` (`username`, `email`, `password`) VALUES
('aygerim', 'aybu@gmail.com', '$2a$10$bdCyh8X3UEdymoQLVfUweuIizGkiLrb7UZu/F90bcy6EdIrw0Fl0i'),
('ann', 'ann@gmail.com', '$2a$10$HU5mjB.yHtB08YoRyR8qA.TFKEkXXeL5H9MjcFBONzlm9gWHI7...');

INSERT INTO `user_task` (`userId`, `taskId`) VALUES
(1, 1),
(1, 2),
(1, 3);