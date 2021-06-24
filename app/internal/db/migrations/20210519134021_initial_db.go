package migrations

import "github.com/wshaman/migrate"

func init() {
	migrate.RegisterSQL(20210519134021,
		"Acme Developer <my-email@acme.corp>",
		"initial_db",
		`
-- create users table
create table users
(
	id serial not null
		constraint users_pk
			primary key,
	name text,
	email text
);

create unique index users_email_uindex
	on users (email);
`,
		`
drop index users_email_uindex;
drop table users;
`,
	)
}
