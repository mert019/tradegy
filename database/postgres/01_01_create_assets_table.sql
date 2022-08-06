\connect Tradegy

-- Table: public.assets

-- DROP TABLE IF EXISTS public.assets;

CREATE SEQUENCE IF NOT EXISTS assets_id_seq;

CREATE TABLE IF NOT EXISTS public.assets
(
    id bigint NOT NULL DEFAULT nextval('assets_id_seq'::regclass),
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name character varying(255) COLLATE pg_catalog."default",
    code text COLLATE pg_catalog."default",
    api_id text COLLATE pg_catalog."default",
    type_id bigint,
    CONSTRAINT assets_pkey PRIMARY KEY (id),
    CONSTRAINT fk_assets_type FOREIGN KEY (type_id)
        REFERENCES public.enums (code) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.assets
    OWNER to postgres;
-- Index: idx_assets_deleted_at

-- DROP INDEX IF EXISTS public.idx_assets_deleted_at;

CREATE INDEX IF NOT EXISTS idx_assets_deleted_at
    ON public.assets USING btree
    (deleted_at ASC NULLS LAST)
    TABLESPACE pg_default;