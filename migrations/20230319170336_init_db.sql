-- +goose Up
-- +goose StatementBegin
CREATE TABLE printers
(
    id           serial,
    api_key      text unique not null,
    name         text default 'printer',
    point_id     int         not null,
    printer_type int  default 1,
    PRIMARY KEY (id)
);

CREATE TABLE checks
(
    id          serial,
    printer_id  text,
    check_order text,
    status      text,
    file_path   text,
    check_type  text,
    PRIMARY KEY (id),
    foreign key (printer_id) REFERENCES printers (api_key)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE printers;
DROP TABLE checks;
-- +goose StatementEnd
