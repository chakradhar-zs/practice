USE zopstore;

CREATE TABLE brands(
    id INT,
    name VARCHAR(50) UNIQUE NOT NULL ,
    PRIMARY KEY (id)
);

INSERT INTO brands VALUES (1,'Amul');

INSERT INTO brands VALUES (2,'Skybag');

CREATE TABLE products(
    id INT,
    name VARCHAR(50) NOT NULL ,
    description VARCHAR(500) NOT NULL ,
    price INT NOT NULL ,
    quantity INT NOT NULL ,
    category varchar(30) NOT NULL ,
    brand_id INT NOT NULL ,
    status ENUM('Available','Out of Stock','Discontinued'),
    PRIMARY KEY (id),
    FOREIGN KEY (brand_id) REFERENCES brands(id)
);

INSERT INTO products VALUES (1,'Amul Ghee','pure',200,2,'ghee',1,'Available');

INSERT INTO products VALUES (2,'Bag','comfort',3000,1,'travel bags',2,'Available');

