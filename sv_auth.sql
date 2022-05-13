drop schema if exists sv_auth;
create schema sv_auth;
use sv_auth;

create table roles
(
    id             bigint       not null,
    pid            bigint       not null,
    application_id bigint       not null,
    code           varchar(64)  not null,
    name           varchar(255) not null,
    is_group       tinyint(1)   not null default 0,
    status         tinyint(1)   not null default 1,
    created_at     datetime(6)  not null default current_timestamp(6),
    created_by     bigint       not null,
    updated_at     datetime(6)  not null default current_timestamp(6) on update current_timestamp(6),
    updated_by     bigint       not null,
    primary key (id),
    constraint roles_chk_is_group check ( is_group = 0 or is_group = 1),
    constraint roles_chk_status check ( status = 0 or status = 1),
    constraint roles_fk_application_id foreign key (application_id) references `sv_sso`.`applications` (id) on update cascade on delete restrict,
    constraint roles_fk_created_by foreign key (created_by) references `sv_sso`.`users` (id) on update cascade on delete restrict,
    constraint roles_fk_updated_by foreign key (updated_by) references `sv_sso`.`users` (id) on update cascade on delete restrict,
    constraint unique roles_unique1_index (application_id, code, name),
    constraint unique roles_unique2_index (application_id, code),
    constraint unique roles_unique3_index (application_id, name)
) engine = InnoDB;

create table actions
(
    id             bigint       not null auto_increment,
    application_id bigint       not null,
    code           varchar(64)  not null,
    name           varchar(255) not null,
    status         tinyint(1)   not null default 1,
    created_at     datetime(6)  not null default current_timestamp(6),
    created_by     bigint       not null,
    updated_at     datetime(6)  not null default current_timestamp(6) on update current_timestamp(6),
    updated_by     bigint       not null,
    primary key (id),
    constraint actions_chk_status check ( status = 0 or status = 1),
    constraint actions_fk_application_id foreign key (application_id) references `sv_sso`.`applications` (id) on update cascade on delete restrict,
    constraint actions_fk_created_by foreign key (created_by) references `sv_sso`.`users` (id) on update cascade on delete restrict,
    constraint actions_fk_updated_by foreign key (updated_by) references `sv_sso`.`users` (id) on update cascade on delete restrict,
    constraint unique actions_unique1_index (application_id, code, name),
    constraint unique actions_unique2_index (application_id, code),
    constraint unique actions_unique3_index (application_id, name)
) engine = InnoDB;

create table menus
(
    id             bigint            not null auto_increment,
    pid            bigint            not null,
    application_id bigint            not null,
    code           varchar(64)       not null,
    name           varchar(255)      not null,
    sort_no        smallint unsigned not null,
    path           varchar(255)      not null,
    redirect       varchar(255),
    component      varchar(255)      not null,
    icon           varchar(255),
    hide           tinyint(1)        not null default 0,
    status         tinyint(1)        not null default 1,
    created_at     datetime(6)       not null default current_timestamp(6),
    created_by     bigint            not null,
    updated_at     datetime(6)       not null default current_timestamp(6) on update current_timestamp(6),
    updated_by     bigint            not null,
    primary key (id),
    constraint menus_chk_hide check ( hide = 0 or hide = 1),
    constraint menus_chk_status check ( status = 0 or status = 1),
    constraint menus_fk_application_id foreign key (application_id) references `sv_sso`.`applications` (id) on update cascade on delete restrict,
    constraint menus_fk_created_by foreign key (created_by) references `sv_sso`.`users` (id) on update cascade on delete restrict,
    constraint menus_fk_updated_by foreign key (updated_by) references `sv_sso`.`users` (id) on update cascade on delete restrict,
    constraint unique menus_unique1_index (application_id, code, name),
    constraint unique menus_unique2_index (application_id, code),
    constraint unique menus_unique3_index (application_id, name),
    constraint unique menus_unique4_index (application_id, pid, path),
    constraint unique menus_unique5_index (application_id, pid, code)
) engine = InnoDB;

create table action_menu
(
    action_id  bigint      not null,
    menu_id    bigint      not null,
    status     tinyint(1)  not null default 1,
    created_at datetime(6) not null default current_timestamp(6),
    created_by bigint      not null,
    updated_at datetime(6) not null default current_timestamp(6) on update current_timestamp(6),
    updated_by bigint      not null,
    primary key (action_id, menu_id),
    constraint action_menu_chk_status check ( status = 0 or status = 1),
    constraint action_menu_fk_action foreign key (action_id) references `actions` (id) on update cascade on delete cascade,
    constraint action_menu_fk_menu foreign key (menu_id) references `menus` (id) on update cascade on delete cascade,
    constraint action_menu_fk_created_by foreign key (created_by) references `sv_sso`.`users` (id) on update cascade on delete restrict,
    constraint action_menu_fk_updated_by foreign key (updated_by) references `sv_sso`.`users` (id) on update cascade on delete restrict,
    index action_menu_action_id_index (action_id),
    index action_menu_menu_id_index (menu_id)
) engine = InnoDB;

create table role_menu
(
    role_id    bigint      not null,
    menu_id    bigint      not null,
    status     tinyint(1)  not null default 1,
    created_at datetime(6) not null default current_timestamp(6),
    created_by bigint      not null,
    updated_at datetime(6) not null default current_timestamp(6) on update current_timestamp(6),
    updated_by bigint      not null,
    primary key (role_id, menu_id),
    constraint role_menu_chk_status check ( status = 0 or status = 1),
    constraint role_menu_fk_role foreign key (role_id) references `roles` (id) on update cascade on delete cascade,
    constraint role_menu_fk_menu foreign key (menu_id) references `menus` (id) on update cascade on delete cascade,
    constraint role_menu_fk_created_by foreign key (created_by) references `sv_sso`.`users` (id) on update cascade on delete restrict,
    constraint role_menu_fk_updated_by foreign key (updated_by) references `sv_sso`.`users` (id) on update cascade on delete restrict,
    index role_menu_role_id_index (role_id),
    index role_menu_menu_id_index (menu_id)
) engine = InnoDB;

create table role_user
(
    id         bigint      not null,
    role_id    bigint      not null,
    user_id    bigint      not null,
    status     tinyint(1)  not null default 1,
    created_at datetime(6) not null default current_timestamp(6),
    created_by bigint      not null,
    updated_at datetime(6) not null default current_timestamp(6) on update current_timestamp(6),
    updated_by bigint      not null,
    primary key (id),
    constraint role_user_chk_status check ( status = 0 or status = 1),
    constraint role_user_fk_role foreign key (role_id) references `roles` (id) on update cascade on delete cascade,
    constraint role_user_fk_user foreign key (user_id) references `sv_sso`.`users` (id) on update cascade on delete cascade,
    constraint role_user_fk_created_by foreign key (created_by) references `sv_sso`.`users` (id) on update cascade on delete restrict,
    constraint role_user_fk_updated_by foreign key (updated_by) references `sv_sso`.`users` (id) on update cascade on delete restrict,
    constraint unique role_user_unique_index (role_id, user_id),
    index role_user_role_id_index (role_id),
    index role_user_user_id_index (user_id)
) engine = InnoDB;

create table role_user__action
(
    role_user_id bigint      not null,
    action_id    bigint         not null,
    status       tinyint(1)  not null default 1,
    created_at   datetime(6) not null default current_timestamp(6),
    created_by   bigint      not null,
    updated_at   datetime(6) not null default current_timestamp(6) on update current_timestamp(6),
    updated_by   bigint      not null,
    primary key (role_user_id, action_id),
    constraint role_user__action_chk_status check ( status = 0 or status = 1),
    constraint role_user__action_fk_role_user foreign key (role_user_id) references `role_user` (id) on update cascade on delete cascade,
    constraint role_user__action_fk_action foreign key (action_id) references `actions` (id) on update cascade on delete cascade,
    constraint role_user__action_fk_created_by foreign key (created_by) references `sv_sso`.`users` (id) on update cascade on delete restrict,
    constraint role_user__action_fk_updated_by foreign key (updated_by) references `sv_sso`.`users` (id) on update cascade on delete restrict,
    index role_user__action_role_user_id_index (role_user_id),
    index role_user__action_action_id_index (action_id)
) engine = InnoDB;