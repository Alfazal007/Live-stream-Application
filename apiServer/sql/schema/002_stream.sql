-- +goose Up
create table stream (
  id varchar(11) primary key,
	admin_id uuid,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	started BOOLEAN default(false),
	ended BOOLEAN default(false),
  FOREIGN KEY (admin_id) REFERENCES users(id)
);

-- +goose Down
drop table stream;
