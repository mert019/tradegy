\connect Tradegy

do $$ 
declare
    cryptocurrencyCode integer := (SELECT code FROM public.enums WHERE area = 'Asset' AND name = 'Cryptocurrency');
    forexCode integer := (SELECT code FROM public.enums WHERE area = 'Asset' AND name = 'Forex');
begin
    INSERT INTO public.assets(id, created_at, updated_at, name, code, api_id, type_id)
	VALUES
        (1, NOW(), NOW(), 'U.S. Dollar', 'USD', 'usd', forexCode),
        (2, NOW(), NOW(), 'Bitcoin', 'BTC', 'bitcoin', cryptocurrencyCode),
        (3, NOW(), NOW(), 'Dogecoin', 'DOGE', 'dogecoin', cryptocurrencyCode)
    ON CONFLICT DO NOTHING;
end $$;
