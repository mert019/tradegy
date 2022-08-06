\connect Tradegy

-- Table: public.orders

-- DROP TABLE IF EXISTS public.orders;

CREATE SEQUENCE IF NOT EXISTS orders_id_seq;

CREATE TABLE IF NOT EXISTS public.orders
(
    id bigint NOT NULL DEFAULT nextval('orders_id_seq'::regclass),
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    buy_amount numeric,
    sell_amount numeric,
    "limit" numeric,
    user_id bigint,
    order_type_id bigint,
    buy_asset_id bigint,
    sell_asset_id bigint,
    order_status_id bigint,
    execution_date_time timestamp with time zone,
    CONSTRAINT orders_pkey PRIMARY KEY (id),
    CONSTRAINT fk_orders_buy_asset FOREIGN KEY (buy_asset_id)
        REFERENCES public.assets (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT fk_orders_order_status FOREIGN KEY (order_status_id)
        REFERENCES public.enums (code) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT fk_orders_order_type FOREIGN KEY (order_type_id)
        REFERENCES public.enums (code) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT fk_orders_sell_asset FOREIGN KEY (sell_asset_id)
        REFERENCES public.assets (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT fk_orders_user FOREIGN KEY (user_id)
        REFERENCES public.users (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.orders
    OWNER to postgres;
-- Index: idx_orders_deleted_at

-- DROP INDEX IF EXISTS public.idx_orders_deleted_at;

CREATE INDEX IF NOT EXISTS idx_orders_deleted_at
    ON public.orders USING btree
    (deleted_at ASC NULLS LAST)
    TABLESPACE pg_default;
