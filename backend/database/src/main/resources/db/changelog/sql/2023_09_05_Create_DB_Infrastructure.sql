create schema if not exists "car_registrations";

create table if not exists "car_registrations"."department"
(
    "dep_code" varchar(40)  not null primary key,
    "dep_name" varchar(100) not null
    );
alter table "car_registrations"."department"
    add constraint "unique_department_code" unique ("dep_code");
alter table "car_registrations"."department"
    add constraint "unique_department_name" unique ("dep_name");
create index "if" not exists "index_department_code" on "car_registrations"."department" ("dep_code");
create index "if" not exists "index_department_name" on "car_registrations"."department" ("dep_name");



create table if not exists "car_registrations"."operation"
(
    "op_code" varchar(40)  not null primary key,
    "op_name" varchar(100) not null
    );
alter table "car_registrations"."operation"
    add constraint "unique_operation_code" unique ("op_code");
alter table "car_registrations"."operation"
    add constraint "unique_operation_name" unique ("op_name");
create index "if" not exists "index_operation_code" on "car_registrations"."operation" ("op_code");
create index "if" not exists "index_operation_name" on "car_registrations"."operation" ("op_name");



create table if not exists "car_registrations"."brand"
(
    "brand_id"   serial primary key,
    "brand_name" varchar(50) not null
    );
alter table "car_registrations"."brand"
    add constraint "unique_brand_name" unique ("brand_name");
create index "if" not exists "index_brand" on "car_registrations"."brand" ("brand_name");



create table if not exists "car_registrations"."model"
(
    "model_id"   serial primary key,
    "model_name" varchar(50
) not null
    );
alter table "car_registrations"."model"
    add constraint "unique_model_name" unique ("model_name");
create index "if" not exists "index_model" on "car_registrations"."model" ("model_name");



create table if not exists "car_registrations"."body_type"
(
    "body_type_id"   serial      not null primary key,
    "body_type_name" varchar(
    50
) not null
    );
alter table "car_registrations"."body_type"
    add constraint "unique_body_type_name" unique ("body_type_name");
create index "if" not exists "index_body_type" on "car_registrations"."body_type" ("body_type_name");



create table if not exists "car_registrations"."vehicle"
(
    "vehicle_id"      serial  not null
    primary
    key,
    "kind"
    varchar(50) not null,
    "brand_id"        integer not null,
    "model_id"        integer not null,
    "body_type_id"    integer not null,
    "fuel_type" varchar
(
    50
) not null,
    "color" varchar(50) not null,
    "make_year"       integer not null,
    "engine_capacity" integer,
    "own_weight"      integer,
    "total_weight" integer
    );
alter table "car_registrations"."vehicle"
    add constraint "unique_vehicle" unique ("kind", "brand_id", "model_id", "body_type_id", "fuel_type", "color",
                                            "make_year", "engine_capacity", "own_weight", "total_weight");
alter table "car_registrations"."vehicle"
    add constraint "fk_brand_id" foreign key ("brand_id") references "car_registrations"."brand" ("brand_id");
alter table "car_registrations"."vehicle"
    add constraint "fk_model_id" foreign key ("model_id") references "car_registrations"."model" ("model_id");
alter table "car_registrations"."vehicle"
    add constraint "fk_body_type_id" foreign key ("body_type_id") references "car_registrations"."body_type" ("body_type_id");
create index "if" not exists "index_vehicle_kind" on "car_registrations"."vehicle" ("kind");
create index "if" not exists "index_vehicle_fuel_type" on "car_registrations"."vehicle" ("fuel_type");
create index "if" not exists "index_vehicle_color" on "car_registrations"."vehicle" ("color");
create index "if" not exists "index_vehicle" on "car_registrations"."vehicle" ("kind","brand_id","model_id","body_type_id","fuel_type","color","make_year","engine_capacity","own_weight","total_weight");



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
    "purpose" varchar
(
    50
) not null
    );
alter table "car_registrations"."registration"
    add constraint "unique_registration" unique ("person_type", "registration_city_code", "department_id",
                                                 "operation_id", "date_registration", "number_registration", "vin",
                                                 "vehicle_id", "purpose");
alter table "car_registrations"."registration"
    add constraint "fk_department_id" foreign key ("department_id") references "car_registrations"."department" ("dep_code");
alter table "car_registrations"."registration"
    add constraint "fk_operation_id" foreign key ("operation_id") references "car_registrations"."operation" ("op_code");
alter table "car_registrations"."registration"
    add constraint "fk_vehicle_id" foreign key ("vehicle_id") references "car_registrations"."vehicle" ("vehicle_id");
create index "if" not exists "index_registration_vin" on "car_registrations"."registration" ("vin");
create index "if" not exists "index_registration_number_registration" on "car_registrations"."registration" ("number_registration");
create index "if" not exists "index_registration_date_registration" on "car_registrations"."registration" ("date_registration");
create index "if" not exists "index_registration_all_fields" on "car_registrations"."registration" ("person_type","registration_city_code","department_id","operation_id","date_registration","number_registration","vin","vehicle_id","purpose");



create view "car_registrations"."all_data" as
select "date_registration",
       "number_registration",
       "brand_name",
       "model_name",
       "body_type_name",
       "fuel_type",
       "engine_capacity",
       "make_year",
       "own_weight",
       "total_weight",
       "color",
       "kind",
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
       "car_registrations"."vehicle" as "v",
       "car_registrations"."registration" as "r"
 where "r"."department_id" = "dep"."dep_code"
   and "r"."operation_id" = "op"."op_code"
   and "r"."vehicle_id" = "v"."vehicle_id";