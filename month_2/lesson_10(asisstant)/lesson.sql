-- 1 ===================================================================
CREATE OR REPLACE FUNCTION check_age() 
    RETURNS TRIGGER 
    LANGUAGE PLPGSQL AS 
$$ 
BEGIN 
    IF NEW.age > 18 THEN 
    RETURN NEW;
    END IF;
        RAISE INFO 'You are not allowed to visit this site. Please come back later.';
    RETURN NULL
END;  
$$;

CREATE TRIGGER check_user_age
BEFORE INSERT OR UPDATE ON users
FOR EACH ROW
EXECUTE PROCEDURE check_age();

insert into users (name, age)
values ('Omadbek', 19)

-- 2 ===================================================================
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE investor (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR NOT NULL CHECK (length(name) < 30),
    birth_date DATE CHECK (birth_date > '1950-01-01'),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE car (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    state_number VARCHAR UNIQUE,
    price NUMERIC NOT NULL,
    investor_id UUID NOT NULL REFERENCES investor(id),
    investor_percentage INT DEFAULT 70,
    status VARCHAR NOT NULL DEFAULT 'in_stock',
    mileage NUMERIC NOT NULL,
    oil_change_status BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE customer (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE orders (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    customer_id UUID NOT NULL REFERENCES customer(id),
    car_id UUID NOT NULL REFERENCES car(id),
    day_count INTEGER NOT NULL,
    total_price NUMERIC NOT NULL,
    paid_price NUMERIC,
    status VARCHAR,
    from_date TIMESTAMP NOT NULL CHECK (from_date::DATE >= CURRENT_DATE),
    to_date TIMESTAMP NOT NULL,
    give_km NUMERIC,
    receive_km NUMERIC,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


INSERT INTO investor (name, birth_date) VALUES ('Omadbek', '2002-02-25');

insert into car (state_number, price, investor_id, mileage) VALUES
('60 W 552 QA', 750000, '992d1254-7faa-4c2b-bf73-414b92f342e8', 15000),
('60 C 777 BA', 600000, '992d1254-7faa-4c2b-bf73-414b92f342e8', 15000),
('20 B 888 BB', 250000, '992d1254-7faa-4c2b-bf73-414b92f342e8', 15000);

INSERT INTO customer (name) VALUES
('Sarvarbek'),
('Davlatbek');

INSERT INTO orders (
    customer_id,
    car_id,
    day_count,
    total_price,
    status,
    from_date,
    to_date
) VALUES (
    '33870ba9-5aab-410d-83a7-6a36133dc96d',
    'dfcd9f95-682a-4c0c-9567-41dc97b7352f',
    2,
    (
        SELECT
            price * 2
        FROM car
        WHERE id = 'dfcd9f95-682a-4c0c-9567-41dc97b7352f'
    ),
    'new',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP + INTERVAL '2 day'
);



-- in_stock
-- booked
-- in_use

-- new
-- client_took
-- client_returned
CREATE OR REPLACE FUNCTION order_car()
RETURNS TRIGGER LANGUAGE PLPGSQL AS
$$
BEGIN 
    IF NEW.status NOT IN ('new', 'client_took', 'client_returned') THEN 
        RAISE EXCEPTION 'Status % is not allowed.', NEW.status;
        RETURN NULL;
    END IF;
    
    IF NEW.status = 'new' AND OLD.status IS NOT NULL THEN 
        UPDATE car SET status = 'booked' WHERE id = NEW.car_id;
        raise info 'Car status updated to booked.';
    ELSIF NEW.status = 'client_took' AND OLD.status = 'new' THEN 
        UPDATE car SET status = 'in_use' WHERE id = NEW.car_id;
        raise info 'Car status updated to in use.';
    ELSIF NEW.status = 'client_returned' AND OLD.status = 'client_took' THEN
        UPDATE car SET status = 'in_stock' WHERE id = NEW.car_id;
        raise info 'Car status updated to in stock.';
    END IF;
    
    RETURN NEW;
END;
$$;

CREATE TRIGGER order_car_tg
AFTER INSERT OR UPDATE ON orders
FOR EACH ROW
EXECUTE PROCEDURE order_car();