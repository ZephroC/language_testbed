
drop table if exists titles cascade;
create table titles (
    title_id serial primary key,
    title   varchar(200) not null
);

drop table if exists books;
create table books (
    book_id serial primary key,
    title int references titles (title_id) on delete cascade
);

drop view if exists books_view;
create view books_view as
select b.book_id, t.title
from books as b
left join titles as t
on b.title=t.title_id;

insert into titles (title) values('Rivers of London'), ('Moon Over Soho'),('Whispers Underground');
insert into books (title) select title_id from titles;
