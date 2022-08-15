DROP DATABASE IF EXISTS ecommerce;

CREATE DATABASE ecommerce;
USE ecommerce;

DROP TABLE IF EXISTS Products;
DROP TABLE IF EXISTS Users;
DROP TABLE IF EXISTS Carts;
DROP TABLE IF EXISTS ProductInCart;

CREATE TABLE Products (
	Id INT NOT NULL AUTO_INCREMENT,
	Title VARCHAR(255),
	Description LONGTEXT,
	ImagePath VARCHAR(255),
    Category VARCHAR(255),
	Price FLOAT,
	Amount INT,
	PRIMARY KEY (Id)
);

CREATE TABLE Users (
	Id INT NOT NULL AUTO_INCREMENT,
	Username VARCHAR(255) UNIQUE,
	Password VARCHAR(255),
	FullName VARCHAR(255),	
	Email VARCHAR(255),
	IsAdmin BOOLEAN NOT NULL DEFAULT 0,
	PRIMARY KEY (Id)
);

CREATE TABLE Carts (
	Id INT NOT NULL AUTO_INCREMENT,
	UserId INT NOT NULL,
	CreatedAt DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY(Id),
	
	INDEX (UserId),
    FOREIGN KEY (UserId) REFERENCES Users(Id) ON UPDATE CASCADE
);

CREATE TABLE ProductInCart (
	Id INT NOT NULL AUTO_INCREMENT,
	CartId INT,
	ProductId INT,
	Quantity INT,
	PRIMARY KEY(Id),
	
	INDEX (CartId),
	INDEX (ProductId),
    
    FOREIGN KEY (CartId) REFERENCES Carts(Id) ON UPDATE CASCADE,
    FOREIGN KEY (ProductId) REFERENCES Products(Id) ON UPDATE CASCADE
);

INSERT INTO Products(Title, Description, ImagePath, Category, Price, Amount) VALUES ("NEW TITLE 1", "NEW DESCRIPTION 1", "NEW IMAGEPATH 1", "NEW CATEGORY 1", 0, 0);
INSERT INTO Products(Title, Description, ImagePath, Category, Price, Amount) VALUES ("NEW TITLE 2", "NEW DESCRIPTION 2", "NEW IMAGEPATH 2", "NEW CATEGORY 2",  0, 0);

INSERT INTO Users(Username, Password, FullName, Email) VALUES ("NEWUSER1", "NEWUSERPASSWORD1", "USER ONE", "newuser1@gmail.com");
INSERT INTO Users(Username, Password, FullName, Email) VALUES ("NEWUSER2", "NEWUSERPASSWORD2", "USER TWO", "newuser2@gmail.com");

INSERT INTO Carts(UserId) VALUES (1);
INSERT INTO Carts(UserId) VALUES (2);

INSERT INTO ProductInCart(CartId, ProductId, Quantity) VALUES (1, 1, 3);
INSERT INTO ProductInCart(CartId, ProductId, Quantity) VALUES (1, 2, 6);