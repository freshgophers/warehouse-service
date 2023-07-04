DO $$
  DECLARE
    country_id UUID;
    store_id UUID;
    city_id UUID;

  BEGIN
    -- EXTENSIONS --
    CREATE EXTENSION IF NOT EXISTS pgcrypto;

    -- TABLES --
    CREATE TABLE IF NOT EXISTS inventories (
    created_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    id             UUID PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
    store_id       VARCHAR NOT NULL,
    catalog_id     VARCHAR,
    product_id     VARCHAR,
    quantity       INTEGER NOT NULL DEFAULT 0,
    quantity_min   INTEGER NOT NULL DEFAULT 0,
    quantity_max   INTEGER NOT NULL DEFAULT 0,
    price          NUMERIC NOT NULL DEFAULT 0,
    price_special  NUMERIC NOT NULL DEFAULT 0,
    price_previous NUMERIC NOT NULL DEFAULT 0,
    is_available   BOOLEAN NOT NULL DEFAULT FALSE
    );


    CREATE TABLE IF NOT EXISTS countries (
        created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        id          UUID PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
        name        VARCHAR NOT NULL UNIQUE
    );

    CREATE TABLE IF NOT EXISTS cities (
        created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        id          UUID PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
        country_id  UUID NOT NULL,
        name        VARCHAR NOT NULL UNIQUE,
        geocenter   VARCHAR NOT NULL,
        FOREIGN KEY (country_id) REFERENCES countries (id)
    );

    CREATE TABLE IF NOT EXISTS currencies (
        created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        id          UUID PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
        country_id  UUID NOT NULL,
        sign        VARCHAR NOT NULL UNIQUE,
        decimals    INTEGER NOT NULL,
        prefix      BOOLEAN NOT NULL DEFAULT FALSE,
        FOREIGN KEY (country_id) REFERENCES countries (id)
    );


    CREATE TABLE IF NOT EXISTS stores (
        created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        id          UUID PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
        merchant_id VARCHAR NOT NULL,
        city_id     UUID NOT NULL,
        name        VARCHAR NOT NULL,
        address     VARCHAR NOT NULL,
        location    VARCHAR NOT NULL,
        rating      NUMERIC NOT NULL DEFAULT 0,
        is_active   BOOLEAN NOT NULL DEFAULT FALSE,
        FOREIGN KEY (city_id) REFERENCES cities (id)

    );

    CREATE TABLE IF NOT EXISTS schedules (
        created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        id          UUID PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
        store_id    UUID NOT NULL UNIQUE,
        periods     JSONB NOT NULL,
        is_active   BOOLEAN NOT NULL DEFAULT FALSE,
        FOREIGN KEY (store_id) REFERENCES stores (id) ON DELETE CASCADE
    );

    CREATE TABLE IF NOT EXISTS deliveries (
        created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        id          UUID PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
        store_id    UUID NOT NULL UNIQUE,
        periods     JSONB NOT NULL,
        areas       JSONB NOT NULL,
        is_active   BOOLEAN NOT NULL DEFAULT FALSE,
        FOREIGN KEY (store_id) REFERENCES stores (id) ON DELETE CASCADE
    );

    -- DATA --
    INSERT INTO countries (name)
    VALUES ('Kazakhstan')
    RETURNING id INTO country_id;

    INSERT INTO cities (country_id, name, geocenter)
    VALUES (country_id, 'Almaty', '43.238949, 76.889709')
    RETURNING id INTO city_id;

    INSERT INTO currencies (country_id, sign, decimals, prefix)
    VALUES (country_id, '₸', 0, FALSE);


  INSERT INTO stores (merchant_id, city_id, name, address, location, rating, is_active)
  VALUES ('merchant_id',city_id, 'YourStoreName', 'YourAddress', 'YourLocation', 0, true)
  RETURNING id INTO store_id;

  -- Insert data into the schedules table
  INSERT INTO schedules (store_id, periods, is_active)
  VALUES (
    store_id,
    '[{"day": "Monday", "from": "09:00", "to":"17:00"}, {"day": "Thursday", "from": "09:00", "to":"17:00"}]',
    true
  );

  -- Insert data into the deliveries table
  INSERT INTO deliveries (store_id, periods, areas, is_active)
  VALUES (
    store_id,
    '[{"day": "Monday", "from": "09:00", "to":"17:00"}, {"day": "Thursday", "from": "09:00", "to":"17:00"}]',
    '[{"latitude": "31:3214", "longitude": "21:3123"}, {"latitude": "21:312314", "longitude": "43:312313"}]',
    true
  );

  COMMIT;
END $$;