\connect Tradegy

-- Table: public.enums

-- DROP TABLE IF EXISTS public.enums;

CREATE SEQUENCE IF NOT EXISTS enums_id_seq;

CREATE TABLE IF NOT EXISTS public.enums
(
    id bigint NOT NULL DEFAULT nextval('enums_id_seq'::regclass),
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    area text COLLATE pg_catalog."default",
    name text COLLATE pg_catalog."default",
    code bigint,
    CONSTRAINT enums_pkey PRIMARY KEY (id),
    CONSTRAINT enums_code_key UNIQUE (code)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.enums
    OWNER to postgres;
-- Index: idx_enums_deleted_at

-- DROP INDEX IF EXISTS public.idx_enums_deleted_at;

CREATE INDEX IF NOT EXISTS idx_enums_deleted_at
    ON public.enums USING btree
    (deleted_at ASC NULLS LAST)
    TABLESPACE pg_default;
