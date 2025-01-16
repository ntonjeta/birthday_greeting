DROP TABLE IF EXISTS test.FRIENDS;
CREATE TABLE test.FRIENDS (
    id MEDIUMINT NOT NULL AUTO_INCREMENT,
    name VARCHAR(30) NOT NULL,
    surname VARCHAR(30) NOT NULL,
    birthday TIMESTAMP NOT NULL,
    PRIMARY KEY (id)
);

INSERT INTO test.FRIENDS (id, name, surname, birthday) values (1, 'Mary', 'Ann', '1974-11-11 00:00:00');
INSERT INTO test.FRIENDS (id, name, surname, birthday) values (2, 'John', 'Doe', '1974-10-1 00:00:00');