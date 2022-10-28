\connect Tradegy

-- Insert Order Types
INSERT INTO public.enums(
	created_at, updated_at, area, name, code)
	VALUES (NOW(), NOW(), 'Order', 'Initialize User', 10000)
ON CONFLICT DO NOTHING;

INSERT INTO public.enums(
	created_at, updated_at, area, name, code)
	VALUES (NOW(), NOW(), 'Order', 'Market Order', 10001)
ON CONFLICT DO NOTHING;

INSERT INTO public.enums(
	created_at, updated_at, area, name, code)
	VALUES (NOW(), NOW(), 'Order', 'Limit Order Buy', 10002)
ON CONFLICT DO NOTHING;

INSERT INTO public.enums(
	created_at, updated_at, area, name, code)
	VALUES (NOW(), NOW(), 'Order', 'Limit Order Sell', 10003)
ON CONFLICT DO NOTHING;

INSERT INTO public.enums(
	created_at, updated_at, area, name, code)
	VALUES (NOW(), NOW(), 'Order', 'Stop Order Sell', 10004)
ON CONFLICT DO NOTHING;

-- Insert Asset Types
INSERT INTO public.enums(
	created_at, updated_at, area, name, code)
	VALUES (NOW(), NOW(), 'Asset', 'Cryptocurrency', 11000)
ON CONFLICT DO NOTHING;

INSERT INTO public.enums(
	created_at, updated_at, area, name, code)
	VALUES (NOW(), NOW(), 'Asset', 'Forex', 11001)
ON CONFLICT DO NOTHING;

-- Insert Order Statuses
INSERT INTO public.enums(
	created_at, updated_at, area, name, code)
	VALUES (NOW(), NOW(), 'OrderStatus', 'Open', 12000)
ON CONFLICT DO NOTHING;

INSERT INTO public.enums(
	created_at, updated_at, area, name, code)
	VALUES (NOW(), NOW(), 'OrderStatus', 'Cancelled', 12001)
ON CONFLICT DO NOTHING;

INSERT INTO public.enums(
	created_at, updated_at, area, name, code)
	VALUES (NOW(), NOW(), 'OrderStatus', 'Executed', 12002)
ON CONFLICT DO NOTHING;
