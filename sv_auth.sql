drop schema if exists sv_auth;
create schema sv_auth;
use sv_auth;

create table roles
(
    id             bigint unsigned  not null,
    pid            bigint unsigned  not null,
    application_id bigint unsigned  not null,
    code           varchar(64)      not null,
    name           varchar(255)     not null,
    is_group       tinyint unsigned not null default 0,
    status         tinyint unsigned not null default 1,
    created_at     datetime(6)      not null default current_timestamp(6),
    created_by     bigint unsigned  not null,
    updated_at     datetime(6)      not null default current_timestamp(6) on update current_timestamp(6),
    updated_by     bigint unsigned  not null,
    primary key (id),
    constraint roles_chk_is_group check ( is_group = 0 or is_group = 1),
    constraint roles_chk_status check ( status = 0 or status = 1),
    constraint roles_fk_application_id foreign key (application_id) references `sv_sso`.`applications` (id) on update cascade on delete restrict,
    constraint roles_fk_created_by foreign key (created_by) references `sv_sso`.`users` (id) on update cascade on delete restrict,
    constraint roles_fk_updated_by foreign key (updated_by) references `sv_sso`.`users` (id) on update cascade on delete restrict,
    constraint unique roles_unique_index (application_id, pid, code)
) engine = InnoDB;

create table actions
(
    id             bigint unsigned  not null auto_increment,
    application_id bigint unsigned  not null,
    code           varchar(64)      not null,
    name           varchar(255)     not null,
    status         tinyint unsigned not null default 1,
    created_at     datetime(6)      not null default current_timestamp(6),
    created_by     bigint unsigned  not null,
    updated_at     datetime(6)      not null default current_timestamp(6) on update current_timestamp(6),
    updated_by     bigint unsigned  not null,
    primary key (id),
    constraint actions_chk_status check ( status = 0 or status = 1),
    constraint actions_fk_application_id foreign key (application_id) references `sv_sso`.`applications` (id) on update cascade on delete restrict,
    constraint actions_fk_created_by foreign key (created_by) references `sv_sso`.`users` (id) on update cascade on delete restrict,
    constraint actions_fk_updated_by foreign key (updated_by) references `sv_sso`.`users` (id) on update cascade on delete restrict,
    constraint unique actions_unique_index (application_id, code)
) engine = InnoDB;

create table menus
(
    id             bigint unsigned   not null auto_increment,
    pid            bigint unsigned   not null,
    application_id bigint unsigned   not null,
    code           varchar(64)       not null,
    name           varchar(255)      not null,
    sort_no        smallint unsigned not null,
    path           varchar(255)      not null,
    redirect       varchar(255),
    component      varchar(255)      not null,
    icon           varchar(255),
    hide           tinyint unsigned  not null default 0,
    status         tinyint unsigned  not null default 1,
    created_at     datetime(6)       not null default current_timestamp(6),
    created_by     bigint unsigned   not null,
    updated_at     datetime(6)       not null default current_timestamp(6) on update current_timestamp(6),
    updated_by     bigint unsigned   not null,
    primary key (id),
    constraint menus_chk_hide check ( hide = 0 or hide = 1),
    constraint menus_chk_status check ( status = 0 or status = 1),
    constraint menus_fk_application_id foreign key (application_id) references `sv_sso`.`applications` (id) on update cascade on delete restrict,
    constraint menus_fk_created_by foreign key (created_by) references `sv_sso`.`users` (id) on update cascade on delete restrict,
    constraint menus_fk_updated_by foreign key (updated_by) references `sv_sso`.`users` (id) on update cascade on delete restrict,
    constraint unique menus_unique_index (application_id, pid, code)
) engine = InnoDB;

create table action_menu
(
    id             bigint unsigned  not null,
    application_id bigint unsigned  not null,
    action_id      bigint unsigned  not null,
    menu_id        bigint unsigned  not null,
    status         tinyint unsigned not null default 1,
    created_at     datetime(6)      not null default current_timestamp(6),
    created_by     bigint unsigned  not null,
    updated_at     datetime(6)      not null default current_timestamp(6) on update current_timestamp(6),
    updated_by     bigint unsigned  not null,
    primary key (id),
    constraint action_menu_chk_status check ( status = 0 or status = 1),
    constraint action_menu_fk_action foreign key (action_id) references `actions` (id) on update cascade on delete cascade,
    constraint action_menu_fk_menu foreign key (menu_id) references `menus` (id) on update cascade on delete cascade,
    constraint action_menu_fk_application_id foreign key (application_id) references `sv_sso`.`applications` (id) on update cascade on delete restrict,
    constraint action_menu_fk_created_by foreign key (created_by) references `sv_sso`.`users` (id) on update cascade on delete restrict,
    constraint action_menu_fk_updated_by foreign key (updated_by) references `sv_sso`.`users` (id) on update cascade on delete restrict,
    index action_menu_action_id_index (action_id),
    index action_menu_menu_id_index (menu_id),
    constraint unique action_menu_unique_index (application_id, action_id, menu_id)
) engine = InnoDB;

create table role_menu
(
    id             bigint unsigned  not null,
    application_id bigint unsigned  not null,
    role_id        bigint unsigned  not null,
    menu_id        bigint unsigned  not null,
    status         tinyint unsigned not null default 1,
    created_at     datetime(6)      not null default current_timestamp(6),
    created_by     bigint unsigned  not null,
    updated_at     datetime(6)      not null default current_timestamp(6) on update current_timestamp(6),
    updated_by     bigint unsigned  not null,
    primary key (id),
    constraint role_menu_chk_status check ( status = 0 or status = 1),
    constraint role_menu_fk_role foreign key (role_id) references `roles` (id) on update cascade on delete cascade,
    constraint role_menu_fk_menu foreign key (menu_id) references `menus` (id) on update cascade on delete cascade,
    constraint role_menu_fk_application_id foreign key (application_id) references `sv_sso`.`applications` (id) on update cascade on delete restrict,
    constraint role_menu_fk_created_by foreign key (created_by) references `sv_sso`.`users` (id) on update cascade on delete restrict,
    constraint role_menu_fk_updated_by foreign key (updated_by) references `sv_sso`.`users` (id) on update cascade on delete restrict,
    index role_menu_role_id_index (role_id),
    index role_menu_menu_id_index (menu_id),
    constraint unique role_menu_unique_index (application_id, role_id, menu_id)
) engine = InnoDB;

create table role_user
(
    id             bigint unsigned  not null,
    application_id bigint unsigned  not null,
    role_id        bigint unsigned  not null,
    user_id        bigint unsigned  not null,
    status         tinyint unsigned not null default 1,
    created_at     datetime(6)      not null default current_timestamp(6),
    created_by     bigint unsigned  not null,
    updated_at     datetime(6)      not null default current_timestamp(6) on update current_timestamp(6),
    updated_by     bigint unsigned  not null,
    primary key (id),
    constraint role_user_chk_status check ( status = 0 or status = 1),
    constraint role_user_fk_role foreign key (role_id) references `roles` (id) on update cascade on delete cascade,
    constraint role_user_fk_user foreign key (user_id) references `sv_sso`.`users` (id) on update cascade on delete cascade,
    constraint role_user_fk_application_id foreign key (application_id) references `sv_sso`.`applications` (id) on update cascade on delete restrict,
    constraint role_user_fk_created_by foreign key (created_by) references `sv_sso`.`users` (id) on update cascade on delete restrict,
    constraint role_user_fk_updated_by foreign key (updated_by) references `sv_sso`.`users` (id) on update cascade on delete restrict,
    constraint unique role_user_unique_index (application_id, role_id, user_id),
    index role_user_role_id_index (role_id),
    index role_user_user_id_index (user_id)
) engine = InnoDB;

create table role_user__action
(
    id             bigint unsigned  not null,
    application_id bigint unsigned  not null,
    role_user_id   bigint unsigned  not null,
    action_id      bigint unsigned  not null,
    status         tinyint unsigned not null default 1,
    created_at     datetime(6)      not null default current_timestamp(6),
    created_by     bigint unsigned  not null,
    updated_at     datetime(6)      not null default current_timestamp(6) on update current_timestamp(6),
    updated_by     bigint unsigned  not null,
    primary key (id),
    constraint role_user__action_chk_status check ( status = 0 or status = 1),
    constraint role_user__action_fk_role_user foreign key (role_user_id) references `role_user` (id) on update cascade on delete cascade,
    constraint role_user__action_fk_action foreign key (action_id) references `actions` (id) on update cascade on delete cascade,
    constraint role_user__action_fk_application_id foreign key (application_id) references `sv_sso`.`applications` (id) on update cascade on delete restrict,
    constraint role_user__action_fk_created_by foreign key (created_by) references `sv_sso`.`users` (id) on update cascade on delete restrict,
    constraint role_user__action_fk_updated_by foreign key (updated_by) references `sv_sso`.`users` (id) on update cascade on delete restrict,
    index role_user__action_role_user_id_index (role_user_id),
    index role_user__action_action_id_index (action_id),
    constraint unique role_user__action_unique_index (application_id, role_user_id, action_id)
) engine = InnoDB;
