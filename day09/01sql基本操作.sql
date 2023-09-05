show databases;
show create database go10;
create database go10 character set utf8mb4;
alter database go10 character set utf8mb4

use go10;
show tables;
drop table emp;
create table emp
(
    id int primary key auto_increment,
    name varchar(32) not null ,
    age int,
    gender enum("male","female","other") default "male",
    tel char(11) unique ,
    salary decimal(8,2),
    birth date
) engine = InnoDB default character set utf8mb4;
desc emp;
show create table emp;
select * from emp;
alter table emp
    add dep varchar(32) after salary;
select * from emp;
show create table emp;
drop table emp;
desc emp;
alter table emp add dep varchar(5) after salary;
alter table emp add city varchar(5) after salary;
INSERT INTO emp (name,gender,age,dep,city,salary) VALUES
                                                      ("yuan","male",24,"教学部","河北省",8000),
                                                      ("eric","male",34,"销售部","山东省",8000),
                                                      ("rain","male",28,"销售部","山东省",10000),
                                                      ("alvin","female",22,"教学部","北京",9000),
                                                      ("George", "male",24,"教学部","河北省",6000),
                                                      ("danae", "male",32,"运营部","北京",12000),
                                                      ("Sera", "male",38,"运营部","河北省",7000),
                                                      ("Echo", "male",19,"运营部","河北省",9000),
                                                      ("Abel", "female",24,"销售部","北京",9000);

select * from emp;
alter table emp modify salary decimal(9,3);
select * from emp where age>=24;
select * from emp where dep="教学部" and gender="male"
select * from emp order by age desc ;
select * from emp order by salary,age;
select dep,max(age) from emp group by dep;

-- 查询每个部门年龄最大的员工姓名
SELECT name,dep,age FROM emp WHERE age in (SELECT max(age) FROM emp GROUP BY dep) ;

select dep,avg(salary) from emp group by dep;
-- 查询教学部的员工最高工资
select dep ,max(salary) from emp group by dep having dep="教学部";
-- 查询平均薪水超过8000的部门
select dep,avg(salary) from emp group by dep having avg(salary)>8000;
select * from emp;
select dep,group_concat(name) from emp group by dep;
select count(name) from emp;
use go10
create table book(
                     id int primary key auto_increment,
                     title varchar(32),
                     price double(5,2),
    pub_id int not null
) engine=INNODB charset=utf8mb4;
create table publisher(
                          id int primary key auto_increment,
                          name varchar(32),
                          email varchar(32),
                          addr varchar(32)
)engine=INNODB charset=utf8mb4;
INSERT INTO book(title,price,pub_id) VALUES
                                         ('西游记',15,1),
                                         ('三国演义',45,2),
                                         ('红楼梦',66,3),
                                         ('水浒传',21,2),
                                         ('红与黑',67,3),
                                         ('乱世佳人',44,6),
                                         ('飘',56,1),
                                         ('放风筝的人',78,3);
INSERT INTO publisher(id,name,email,addr) VALUES
                                              (1,'清华出版社',"123","bj"),
                                              (2,'北大出版社',"234","bj"),
                                              (3,'机械工业出版社',"345","nj"),
                                              (4,'邮电出版社',"456","nj"),
                                              (5,'电子工业出版社',"567","bj"),
                                              (6,'人民大学出版社',"678","bj");
CREATE TABLE author(
                       id INT PRIMARY KEY AUTO_INCREMENT,
                       NAME VARCHAR(32) NOT NULL
)ENGINE=INNODB CHARSET=utf8;
CREATE TABLE book2author(
                            id INT NOT NULL UNIQUE AUTO_INCREMENT,
                            author_id INT NOT NULL,
                            book_id INT NOT NULL
)ENGINE=INNODB CHARSET=utf8;
INSERT INTO author(NAME) VALUES
                             ('yuan'),
                             ('rain'),
                             ('alvin'),
                             ('eric');
INSERT INTO book2author(author_id,book_id) VALUES
                                               (1,1),
                                               (1,2),
                                               (2,1),
                                               (3,3),
                                               (3,4),
                                               (1,3);
CREATE TABLE authorDetail(
                             id INT PRIMARY KEY AUTO_INCREMENT,
                             tel VARCHAR(32),
                             addr VARCHAR(32),
                             author_id INT NOT NULL unique -- 也可以给author添加一个关联字段：   alter table author add authorDetail_id INT NOT NULL
)ENGINE=INNODB CHARSET=utf8;
INSERT INTO authorDetail(tel,addr,author_id) VALUES
                                                 ("110","北京",1),
                                                 ("911","成都",2),
                                                 ("119","上海",3),
                                                 ("111","广州",4);
-- 多表查询 TODO
-- 外键约束 TODO