create table comments
(
    video_id     int           not null,
    user_id      int           not null comment '评论者ID',
    comment_text varchar(1024) null,
    comment_id   bigint auto_increment
        primary key,
    created_at   timestamp     null
);

create index favorites_video_id_index
    on comments (video_id);

create table favorites
(
    video_id int not null,
    user_id  int not null,
    primary key (video_id, user_id)
);

create table relations
(
    user_id     int not null,
    follower_id int not null,
    primary key (user_id, follower_id)
);

create table users
(
    user_id   int auto_increment
        primary key,
    name      varchar(64)                                      not null,
    signature varchar(100) default '这个人很懒，什么都没有留下' null,
    cover     varchar(512)                                     not null comment '用户个人页顶部大图',
    avatar    varchar(512)                                     not null,
    pass      varchar(60)                                      not null
);

create index users_user_id_name_index
    on users (user_id, name);

create table videos
(
    video_id   int auto_increment
        primary key,
    publish_at timestamp    not null,
    author_id  int          not null,
    play       varchar(512) not null,
    cover      varchar(512) not null,
    title      varchar(128) not null
);

