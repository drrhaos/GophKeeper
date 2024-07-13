CREATE TABLE IF NOT EXISTS public.users
(
    id serial NOT NULL,
    login character varying COLLATE pg_catalog."default" NOT NULL,
    password character varying(64) COLLATE pg_catalog."default" NOT NULL,
    registered_at timestamp with time zone NOT NULL,
    last_time timestamp with time zone NOT NULL,
    CONSTRAINT users_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;


CREATE TABLE IF NOT EXISTS public.store
(
    id serial NOT NULL,
    user_id bigint NOT NULL,
    uuid character varying COLLATE pg_catalog."default" NOT NULL,
    login character varying COLLATE pg_catalog."default",
    password character varying COLLATE pg_catalog."default",
    data text COLLATE pg_catalog."default",
    card_number character varying COLLATE pg_catalog."default",
    card_cvc character varying COLLATE pg_catalog."default",
    card_date character varying COLLATE pg_catalog."default",
    card_owner character varying COLLATE pg_catalog."default",
    file_name character varying COLLATE pg_catalog."default",
    update_at timestamp with time zone NOT NULL,
    CONSTRAINT store_pkey PRIMARY KEY (id),
    CONSTRAINT "user" FOREIGN KEY (user_id)
        REFERENCES public.users (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE CASCADE
        NOT VALID
)

TABLESPACE pg_default;

