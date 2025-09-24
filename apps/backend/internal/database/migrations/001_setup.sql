-- safe camel-case converter
CREATE OR REPLACE FUNCTION camel(input_row anyelement)
RETURNS jsonb
LANGUAGE plpgsql
STRICT
IMMUTABLE
AS $$
DECLARE
    result jsonb := '{}'::jsonb;
    rec record;
    key_text text;
    cleaned text;
    camel_key text;
BEGIN
    FOR rec IN
        SELECT key, value
        FROM jsonb_each(to_jsonb(input_row))
    LOOP
        key_text := rec.key;
        -- 1) replace underscores with spaces -> "user_name" -> "user name"
        -- 2) initcap -> "User Name"
        -- 3) remove spaces -> "UserName"
        -- 4) lowercase first char -> "userName"
        cleaned := regexp_replace(initcap(regexp_replace(key_text, '_', ' ', 'g')), E'\\s', '', 'g');
        camel_key := lower(substring(cleaned, 1, 1)) || substring(cleaned, 2);
        result := result || jsonb_build_object(camel_key, rec.value);
    END LOOP;

    RETURN result;
END;
$$;

-- trigger to set updated_at timestamp on row update
CREATE OR REPLACE FUNCTION trigger_set_updated_at()
RETURNS trigger
LANGUAGE plpgsql
AS $$
BEGIN
    NEW.updated_at := CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$;
