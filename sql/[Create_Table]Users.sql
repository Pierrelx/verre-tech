DROP TABLE IF EXISTS users;
CREATE TABLE public.users
(
	id SERIAL PRIMARY KEY,
    first_name character varying(150) NOT NULL,
    last_name character varying(150) NOT NULL,
    address character varying(300) NOT NULL,
    postal_code text NOT NULL,
    city text NOT NULL,
    email text NOT NULL,
    created_on_utc integer NOT NULL,
    updated_on_utc integer NOT NULL
)
WITH (
	OIDS=FALSE
);
ALTER TABLE public.users
	OWNER TO vtuser;
