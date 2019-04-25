drop table if exists users;

create table users (
  id char(36) not null,
  created datetime not null default now(),
  updated datetime not null default now(),
  username varchar(20) not null unique,
  password_hash varchar(60) not null,
  email varchar(50) not null,
  first_name varchar(50) not null default '',
  last_name varchar(50) not null default '',
  gravatar_hash varchar(32) not null default '',
  is_locked tinyint not null default 0,
  
  primary key (id)
);