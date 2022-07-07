create table t_hs_info
(
    id         bigint unsigned auto_increment comment '自增id'
        primary key,
    created_at datetime     null comment '创建时间',
    updated_at datetime     null comment '更新时间',
    deleted_at datetime     null comment '删除时间',
    title      varchar(256) not null comment '标题',
    url        varchar(256) null comment '网站url',
    m3u8_url   varchar(256) null comment '下载url',
    class_id   bigint       null comment '分类id',
    platform   varchar(255) null comment '平台名称',
    page       bigint       null comment '页码',
    location   varchar(255) null comment '位置',
    constraint idx_title_unique
        unique (title)
);

create index idx_m3u8_url
    on t_hs_info (m3u8_url);

create index idx_t_hs_info_deleted_at
    on t_hs_info (deleted_at);

create index sel_title_normal
    on t_hs_info (title);

