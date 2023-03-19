-- +goose Up
-- +goose StatementBegin
insert into printers (api_key, name, point_id, printer_type)
VALUES ('111', 'first', 1, 1),
       ('222', 'second', 1, 2),
       ('333', 'third', 2, 1);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
delete
from printers
where true;
-- +goose StatementEnd
