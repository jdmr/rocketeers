drop database if exists pbe;
create database pbe;
use pbe;

create table questions (
	id varchar(50) primary key,
    book varchar(50) not null,
    chapter varchar(50) not null,
    verses varchar(50) not null,
    question varchar(500) not null
);

create table answers(
	id varchar(50) primary key,
    answer varchar(500) not null,
    status bit default 0,
    question_id varchar(50) not null,
    index answers_questions_idx (question_id),
    foreign key (question_id)
    references questions(id)
    on delete cascade
);

create table games(
	id varchar(50) primary key,
    name varchar(50) not null,
    seconds int not null,
    created datetime not null,
    questions int not null
);

create table game_chapters(
	id varchar(50) primary key,
    book varchar(50) not null,
    chapter varchar(50),
    game_id varchar(50),
    index game_chapters_games_idx(game_id),
    foreign key (game_id)
    references games(id)
    on delete cascade
);

create table game_questions(
	id varchar(50) primary key,
    game_id varchar(50),
    question_id varchar(50) not null,
    position smallint,
    is_current bit,
    index game_chapters_games_idx(game_id),
    index game_questions_questionos_idx(question_id),
    foreign key (game_id)
    references games(id)
    on delete cascade,
    foreign key (question_id)
    references questions(id)
    on delete cascade
);

create table teams(
	id varchar(50) primary key,
    name varchar(50) not null,
    game_id varchar(50) not null,
    index teams_games_idx(game_id),
    foreign key (game_id)
    references games(id)
    on delete cascade
);

create table team_answers(
	id varchar(50) primary key,
    game_id varchar(50) not null,
    team_id varchar(50) not null,
    answer_id varchar(50) not null,
    created datetime not null,
    index team_answers_games_idx(game_id),
    index team_answers_teams_idx(team_id),
    index team_answers_answers_idx(answer_id),
    foreign key (game_id)
    references games(id)
    on delete cascade,
    foreign key (team_id)
    references teams(id)
    on delete cascade,
    foreign key (answer_id)
    references answers(id)
    on delete cascade
);
