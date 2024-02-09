CREATE TABLE books
(
    id        INT GENERATED ALWAYS AS IDENTITY,
    author_id INT NOT NULL,
    name      VARCHAR(50) UNIQUE NOT NULL,
    summary   text,
    PRIMARY KEY (id),
    CONSTRAINT fk_author
        FOREIGN KEY(author_id)
            REFERENCES authors(id)
            ON DELETE CASCADE
);