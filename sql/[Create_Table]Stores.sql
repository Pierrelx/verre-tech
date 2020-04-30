DROP TABLE IF EXISTS stores;
CREATE TABLE public.stores
(
	id SERIAL PRIMARY KEY,
    name character varying(150) NOT NULL,
    address character varying(300) NOT NULL,
    postal_code text NOT NULL,
    county text NOT NULL,
    city text NOT NULL,
    type text,
    latitude float,
    longitude float,
    created_on_utc integer NOT NULL,
    updated_on_utc integer NOT NULL
)
WITH (
	OIDS=FALSE
);
ALTER TABLE public.stores
	OWNER TO vtuser;
