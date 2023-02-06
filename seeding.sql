drop table if exists users cascade;
drop table if exists employees cascade;

CREATE TABLE public.users
(
    id bigserial NOT NULL,
    email character varying NOT NULL,
    password character varying NOT NULL,
    name character varying NOT NULL,
    role character varying NOT NULL,
    photo text,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp,
    PRIMARY KEY (id)
);

CREATE TABLE public.employees
(
    id bigserial NOT NULL,
    name character varying NOT NULL,
	job_description character varying NOT NULL,
	entry_date bigint NOT NULL,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp,
    PRIMARY KEY (id)
);

INSERT INTO users (
	created_at,
	updated_at,
	deleted_at,
	email,
	password,
	name,
	role,
	photo
) VALUES (
	'2023-01-30 15:34:49.549393',
	'2023-01-30 15:34:49.549393',
    NULL,
	'andreas@gmail.com',
	'$2a$10$UcKWRLEBlmd1O4q4kyKXXeX9USqyjexeZFlsViSZQojaS5sNajFGq',
	'Andreas Timothy',
	'admin',
	'https://media.licdn.com/dms/image/C4E03AQFWxbpWWPRg_A/profile-displayphoto-shrink_200_200/0/1643728652968?e=1680134400&v=beta&t=SHm2xgXOYFzsKwAFlbOpO_2Oi-8FRVqxmrOn0bQTt9M'
);

INSERT INTO users (
	created_at,
	updated_at,
	deleted_at,
	email,
	password,
	name,
	role,
	photo
) VALUES (
	'2023-01-30 15:44:49.549393',
	'2023-01-30 15:44:49.549393',
    NULL,
	'timothy@gmail.com',
	'$2a$10$UcKWRLEBlmd1O4q4kyKXXeX9USqyjexeZFlsViSZQojaS5sNajFGq',
	'Timothy Timothy',
	'user',
	''
);

INSERT INTO users (
	created_at,
	updated_at,
	deleted_at,
	email,
	password,
	name,
	role,
	photo
) VALUES (
	'2023-01-30 15:44:49.549393',
	'2023-01-30 15:44:49.549393',
    NULL,
	'andreastimothy@gmail.com',
	'$2a$10$UcKWRLEBlmd1O4q4kyKXXeX9USqyjexeZFlsViSZQojaS5sNajFGq',
	'Andreas Timothy',
	'guest',
	''
);

INSERT INTO users (
	created_at,
	updated_at,
	deleted_at,
	email,
	password,
	name,
	role,
	photo
) VALUES (
	'2023-01-30 15:44:49.549393',
	'2023-01-30 15:44:49.549393',
    NULL,
	'andreastimothyy@gmail.com',
	'$2a$10$UcKWRLEBlmd1O4q4kyKXXeX9USqyjexeZFlsViSZQojaS5sNajFGq',
	'Andreas Timothy',
	'user',
	'https://media.licdn.com/dms/image/C4E03AQFWxbpWWPRg_A/profile-displayphoto-shrink_200_200/0/1643728652968?e=1680134400&v=beta&t=SHm2xgXOYFzsKwAFlbOpO_2Oi-8FRVqxmrOn0bQTt9M'
);

INSERT INTO users (
	created_at,
	updated_at,
	deleted_at,
	email,
	password,
	name,
	role,
	photo
) VALUES (
	'2023-01-30 15:44:49.549393',
	'2023-01-30 15:44:49.549393',
    NULL,
	'andreas.timothy@shopee.com',
	'$2a$10$UcKWRLEBlmd1O4q4kyKXXeX9USqyjexeZFlsViSZQojaS5sNajFGq',
	'Andreas Timothy',
	'admin',
	''
);
