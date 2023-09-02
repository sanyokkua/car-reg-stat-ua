create schema if not exists "car_registrations";

create table if not exists "car_registrations"."department"
(
    "dep_code" varchar(40)  not null primary key,
    "dep_name" varchar(100) not null
);

create table if not exists "car_registrations"."operation"
(
    "op_code" varchar(40)  not null primary key,
    "op_name" varchar(100) not null
);

create table if not exists "car_registrations"."brand"
(
    "brand_id"   serial primary key,
    "brand_name" varchar(50) not null
);

create table if not exists "car_registrations"."model"
(
    "model_id"   serial primary key,
    "model_name" varchar(50) not null
);

create table if not exists "car_registrations"."body_type"
(
    "body_type_id"   serial      not null primary key,
    "body_type_name" varchar(50) not null
);

create table if not exists "car_registrations"."fuel_type"
(
    "fuel_type_id"   serial      not null primary key,
    "fuel_type_name" varchar(50) not null
);

create table if not exists "car_registrations"."color"
(
    "color_id"   serial      not null primary key,
    "color_name" varchar(50) not null
);

create table if not exists "car_registrations"."kind"
(
    "kind_id"   serial      not null primary key,
    "kind_name" varchar(50) not null
);

create table if not exists "car_registrations"."vehicle"
(
    "vehicle_id"      serial  not null primary key,
    "kind_id"         integer not null,
    "brand_id"        integer not null,
    "model_id"        integer not null,
    "body_type_id"    integer not null,
    "fuel_type_id"    integer not null,
    "color_id"        integer not null,
    "make_year"       integer not null,
    "engine_capacity" integer,
    "own_weight"      integer,
    "total_weight"    integer
);

create table if not exists "car_registrations"."registration"
(
    "registration_id"        serial      not null primary key,
    "person_type"            varchar(2)  not null,
    "registration_city_code" varchar(100),
    "department_id"          varchar(40) not null,
    "operation_id"           varchar(40) not null,
    "date_registration"      date        not null,
    "number_registration"    varchar(20),
    "vin"                    varchar(200),
    "vehicle_id"             integer     not null,
    "purpose"                varchar(50) not null
);


alter table "car_registrations"."vehicle"
    add constraint "fk_kind_id" foreign key ("kind_id") references "car_registrations"."kind" ("kind_id");
alter table "car_registrations"."vehicle"
    add constraint "fk_brand_id" foreign key ("brand_id") references "car_registrations"."brand" ("brand_id");
alter table "car_registrations"."vehicle"
    add constraint "fk_model_id" foreign key ("model_id") references "car_registrations"."model" ("model_id");
alter table "car_registrations"."vehicle"
    add constraint "fk_body_type_id" foreign key ("body_type_id") references "car_registrations"."body_type" ("body_type_id");
alter table "car_registrations"."vehicle"
    add constraint "fk_fuel_type_id" foreign key ("fuel_type_id") references "car_registrations"."fuel_type" ("fuel_type_id");
alter table "car_registrations"."vehicle"
    add constraint "fk_color_id" foreign key ("color_id") references "car_registrations"."color" ("color_id");

alter table "car_registrations"."registration"
    add constraint "fk_department_id" foreign key ("department_id") references "car_registrations"."department" ("dep_code");
alter table "car_registrations"."registration"
    add constraint "fk_operation_id" foreign key ("operation_id") references "car_registrations"."operation" ("op_code");
alter table "car_registrations"."registration"
    add constraint "fk_vehicle_id" foreign key ("vehicle_id") references "car_registrations"."vehicle" ("vehicle_id");

alter table "car_registrations"."kind"
    add constraint "unique_kind" unique ("kind_name");
alter table "car_registrations"."brand"
    add constraint "unique_brand" unique ("brand_name");
alter table "car_registrations"."model"
    add constraint "unique_model" unique ("model_name");
alter table "car_registrations"."body_type"
    add constraint "unique_body_type" unique ("body_type_name");
alter table "car_registrations"."fuel_type"
    add constraint "unique_fuel_type" unique ("fuel_type_name");
alter table "car_registrations"."color"
    add constraint "unique_color" unique ("color_name");
alter table "car_registrations"."department"
    add constraint "unique_department" unique ("dep_name");
alter table "car_registrations"."operation"
    add constraint "unique_operation" unique ("op_name");

create index if not exists "index_kind" on "car_registrations"."kind" ("kind_name");
create index if not exists "index_brand" on "car_registrations"."brand" ("brand_name");
create index if not exists "index_model" on "car_registrations"."model" ("model_name");
create index if not exists "index_body_type" on "car_registrations"."body_type" ("body_type_name");
create index if not exists "index_fuel_type" on "car_registrations"."fuel_type" ("fuel_type_name");
create index if not exists "index_color" on "car_registrations"."color" ("color_name");
create index if not exists "index_department_name" on "car_registrations"."department" ("dep_name");
create index if not exists "index_operation_name" on "car_registrations"."operation" ("op_name");
create index if not exists "index_vehicle" on "car_registrations"."vehicle" ("kind_id", "brand_id", "model_id",
                                                                             "body_type_id", "fuel_type_id", "color_id",
                                                                             "make_year", "engine_capacity",
                                                                             "own_weight", "total_weight");

create view "car_registrations"."all_data" as
select "date_registration",
       "number_registration",
       "brand_name",
       "model_name",
       "body_type_name",
       "fuel_type_name",
       "engine_capacity",
       "make_year",
       "own_weight",
       "total_weight",
       "color_name",
       "kind_name",
       "purpose",
       "vin",
       "dep_code",
       "dep_name",
       "op_code",
       "op_name",
       "person_type",
       "registration_city_code"
  from "car_registrations"."department" as "dep",
       "car_registrations"."operation" as "op",
       "car_registrations"."brand" as "brand",
       "car_registrations"."model" as "model",
       "car_registrations"."body_type" as "body",
       "car_registrations"."fuel_type" as "fuel",
       "car_registrations"."color" as "color",
       "car_registrations"."kind" as "kind",
       "car_registrations"."vehicle" as "v",
       "car_registrations"."registration" as "r"
 where "r"."department_id" = "dep"."dep_code"
   and "r"."operation_id" = "op"."op_code"
   and "r"."vehicle_id" = "v"."vehicle_id";