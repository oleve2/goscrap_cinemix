123

--scipts for SQLITE database

--// какие есть авторы
select substr( author_title, 0, instr(author_title, '-')-1) as author, count(1) as cnt
from cine_rows 
where 1=1
--and author_title like '%yared%'
group by substr( author_title, 0, instr(author_title, '-')-1)
order by cnt desc

-- какие треки композитора проигрывались
select * from cine_rows
where 1=1 
and author_title like '%Alexandre Desplat%'
and album like '%harry%'

-- автор и кино (группировки)
select substr( author_title, 0, instr(author_title, '-')-1) as author, album, count(1) as cnt
from cine_rows
where 1=1
and author_title like '%yared%'
group by substr( author_title, 0, instr(author_title, '-')-1), album
order by cnt desc



select album, count(1) as cnt from cine_rows
group by album
order by cnt desc







