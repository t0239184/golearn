create table users (
    ID bigint(20) not null auto_increment primary key,
    Account varchar(50) not null,
    Password varchar(50) not null,
    Status varchar(1) not null default '0',
    Creator varchar(50) not null,
    CreateAt datetime not null,
    Updater varchar(50),
    UpdateAt datetime,
    Deleter varchar(50),
    DeleteAt datetime,
    IsDelete tinyint(1) not null default 0
);