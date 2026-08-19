-- +goose Up
CREATE TABLE records_2026 (
    id INT GENERATED ALWAYS AS IDENTITY (START WITH 1 INCREMENT BY 1) PRIMARY KEY,
    record_date DATE NOT NULL,
    record_title TEXT NOT NULL,
    record_type TEXT NOT NULL,
    updated_at TIMESTAMP NOT NULL
);
-- +goose Down
DROP TABLE records_2026;
--category options (book, movie, tv show, other)