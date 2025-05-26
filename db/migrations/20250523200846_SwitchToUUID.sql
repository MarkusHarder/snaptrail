-- migrate:up
create extension if not exists "uuid-ossp";

-- add uuids to sessions and thumbnails
alter table sessions add column session_uuid uuid unique default uuid_generate_v4();
alter table thumbnails add column session_uuid uuid default uuid_generate_v4();
alter table thumbnails add column thumbnail_uuid uuid unique default uuid_generate_v4();

-- populate thumbnail fkey references with new session uuids
update thumbnails
set session_uuid = sessions.session_uuid
from sessions
where thumbnails.session_id = sessions.id;

-- change foreign keys
alter table thumbnails drop constraint thumbnails_session_id_fkey;
alter table thumbnails drop column session_id;

alter table thumbnails
add constraint thumbnails_session_uuid_fkey
foreign key (session_uuid) references sessions(session_uuid) on delete cascade;


-- change primary key, drop serial on sessions
alter table sessions drop constraint sessions_pkey;
alter table sessions add primary key (session_uuid);

alter table sessions drop column id;
alter table sessions rename column session_uuid to id;

-- change primary key, drop serial on thumbnails
alter table sessions drop constraint sessions_pkey;
alter table sessions add primary key (id);

alter table thumbnails drop column id;
alter table thumbnails rename column thumbnail_uuid to id;
alter table thumbnails rename column session_uuid to session_id;
alter table thumbnails rename constraint thumbnails_session_uuid_fkey to thumbnails_session_id_fkey;



-- migrate:down

-- add uuids to sessions and thumbnails
alter table sessions add column session_serial serial unique; 
alter table thumbnails add column session_serial serial;
alter table thumbnails add column thumbnail_serial serial unique;


UPDATE sessions
SET session_serial = nextval('sessions_session_serial_seq');

UPDATE thumbnails
SET thumbnail_serial = nextval('thumbnails_thumbnail_serial_seq');
-- populate thumbnail fkey references with new session uuids
update thumbnails
set session_serial = sessions.session_serial
from sessions
where thumbnails.session_id = sessions.id;

-- change foreign keys
alter table thumbnails
add constraint thumbnails_session_serial_fkey
foreign key (session_serial) references sessions(session_serial) on delete cascade;

alter table thumbnails drop constraint thumbnails_session_id_fkey;
alter table thumbnails drop column session_id;

-- change primary key, drop serial on sessions
alter table sessions drop constraint sessions_pkey;
alter table sessions add primary key (session_serial);

alter table sessions drop column id;
alter table sessions rename column session_serial to id;

-- change primary key, drop serial on thumbnails
alter table sessions drop constraint sessions_pkey;
alter table sessions add primary key (id);

alter table thumbnails drop column id;
alter table thumbnails rename column thumbnail_serial to id;
alter table thumbnails rename column session_serial to session_id;
alter table thumbnails rename constraint thumbnails_session_serial_fkey to thumbnails_session_id_fkey;

