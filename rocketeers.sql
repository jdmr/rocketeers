use rocketeers;

drop table if exists user_images;
drop table if exists images;
drop table if exists user_roles;
drop table if exists roles;
drop table if exists users;

create table roles (
	id varchar(20) primary key
);

create table users (
	id varchar(50) primary key
    , first_name varchar(50) not null
    , last_name varchar(50) not null
    , birthdate date
    , gender varchar(10)
    , email varchar(255) unique
    , phone varchar(20)
    , carrier varchar(50)
    , image_url varchar(255)
);

create table user_roles (
	user_id varchar(50) not null
    , role_id varchar(20) not null
    , primary key (user_id, role_id)
    , index user_roles_user_idx (user_id)
    , index user_roles_role_idx (role_id)
    , foreign key (user_id) 
		references users(id)
        on delete cascade
	, foreign key (role_id)
		references roles(id)
        on delete cascade
);

create table images (
	id varchar(50) primary key
    , image blob not null
    , content_type varchar(20) not null
);

create table user_images (
	user_id varchar(50) not null
    , image_id varchar(50) not null
    , primary key (user_id, image_id)
    , index user_images_user_idx (user_id)
    , index user_images_image_idx (image_id)
    , foreign key (user_id)
		references users(id)
        on delete cascade
	, foreign key (image_id)
		references images(id)
        on delete cascade
);

insert into roles values('ADMIN');
insert into roles values('PATHFINDER');
insert into roles values('COUNSELOR');
insert into roles values('PARENT');