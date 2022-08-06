\connect Tradegy

-- Insert Order Types
INSERT INTO public.enums(
	created_at, updated_at, area, name, code)
	VALUES (NOW(), NOW(), 'Order', 'InitializeUser', 10000)
ON CONFLICT DO NOTHING;

INSERT INTO public.enums(
	created_at, updated_at, area, name, code)
	VALUES (NOW(), NOW(), 'Order', 'MarketOrderBuy', 10001)
ON CONFLICT DO NOTHING;

INSERT INTO public.enums(
	created_at, updated_at, area, name, code)
	VALUES (NOW(), NOW(), 'Order', 'MarketOrderSell', 10002)
ON CONFLICT DO NOTHING;

INSERT INTO public.enums(
	created_at, updated_at, area, name, code)
	VALUES (NOW(), NOW(), 'Order', 'LimitOrderBuy', 10003)
ON CONFLICT DO NOTHING;

INSERT INTO public.enums(
	created_at, updated_at, area, name, code)
	VALUES (NOW(), NOW(), 'Order', 'LimitOrderSell', 10004)
ON CONFLICT DO NOTHING;

INSERT INTO public.enums(
	created_at, updated_at, area, name, code)
	VALUES (NOW(), NOW(), 'Order', 'StopOrderSell', 10005)
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
	VALUES (NOW(), NOW(), 'OrderStatus', 'CanceledByUser', 12001)
ON CONFLICT DO NOTHING;

INSERT INTO public.enums(
	created_at, updated_at, area, name, code)
	VALUES (NOW(), NOW(), 'OrderStatus', 'CanceledBySystem', 12002)
ON CONFLICT DO NOTHING;

INSERT INTO public.enums(
	created_at, updated_at, area, name, code)
	VALUES (NOW(), NOW(), 'OrderStatus', 'Executed', 12003)
ON CONFLICT DO NOTHING;
