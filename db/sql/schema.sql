CREATE TABLE users (
id SERIAL PRIMARY KEY,
google_id TEXT UNIQUE NOT NULL,
email TEXT UNIQUE NOT NULL,
name TEXT NOT NULL,
avatar_url TEXT,
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
role INT not null default 1
);

CREATE TABLE products (
id SERIAL PRIMARY KEY,
name TEXT NOT NULL,
description TEXT not null ,
price int NOT NULL,
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
special INT references special(id),
option_1_name TEXT NOT NULL,
option_2_name TEXT NOT NULL
);

CREATE TABLE product_variants (
id SERIAL PRIMARY KEY,
product_id INT REFERENCES products(id) ON DELETE CASCADE not null,
sku TEXT UNIQUE NOT NULL,
price int NOT NULL,
stock_quantity INT NOT NULL DEFAULT 0,
option_1_value TEXT NOT NULL,
option_2_value TEXT NOT NULL,
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE product_images (
id SERIAL PRIMARY KEY,
product_id INT REFERENCES products(id) ON DELETE CASCADE NOT NULL,
image_url TEXT NOT NULL,
is_primary BOOLEAN DEFAULT FALSE NOT NULL,
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE categories (
id SERIAL PRIMARY KEY,
name TEXT NOT NULL,
description TEXT
);

CREATE TABLE product_categories (
product_id INT REFERENCES products(id) ON DELETE CASCADE,
category_id INT REFERENCES categories(id) ON DELETE CASCADE,
PRIMARY KEY (product_id, category_id)
);


CREATE TABLE basket_items (
id SERIAL PRIMARY KEY,
user_id INT REFERENCES users(id) ON DELETE CASCADE NOT NULL ,
product_variant_id INT REFERENCES product_variants(id) ON DELETE CASCADE NOT NULL,
quantity INT NOT NULL DEFAULT 1,
created_at TIMESTAMP DEFAULT NOW(),
updated_at TIMESTAMP DEFAULT NOW()
);



CREATE TABLE favourite_items (
id SERIAL PRIMARY KEY,
user_id INT REFERENCES users(id) not null ,
product_id INT REFERENCES products(id) not null,
created_at TIMESTAMP DEFAULT NOW(),
updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE orders (
id SERIAL PRIMARY KEY,
user_id INT REFERENCES users(id) not null ,
status VARCHAR(50) DEFAULT 'pending' not null,
total_price int NOT NULL,
shipping_id INT REFERENCES shippings(id) not null ,
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE order_items (
id SERIAL PRIMARY KEY,
order_id INT REFERENCES orders(id) ON DELETE CASCADE not null,
product_variant_id INT REFERENCES product_variants(id) not null,
quantity INT NOT NULL DEFAULT 1,
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE conversations (
id SERIAL PRIMARY KEY,
name VARCHAR(255) DEFAULT 'Unnamed conversation' not null ,
order_id INT REFERENCES orders(id) not null,
created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE conversation_messages (
id SERIAL PRIMARY KEY,
conversation_id INT REFERENCES conversations(id) NOT NULL,
sender_id INT REFERENCES users(id) NOT NULL,
content TEXT NOT NULL,
created_at TIMESTAMP DEFAULT NOW() NOT NULL,
read_at TIMESTAMP,
status VARCHAR(50) DEFAULT 'unread'
);

CREATE TABLE reviews (
id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
sender_id INT REFERENCES users(id) NOT NULL,
product_id INT REFERENCES products(id) NOT NULL,
rating INT CHECK (rating >= 1 AND rating <= 5) not null,
content TEXT NOT NULL,
created_at TIMESTAMP DEFAULT NOW() NOT NULL
);

CREATE TABLE conversations_users (
id SERIAL PRIMARY KEY,
conversation_id INT REFERENCES conversations(id) NOT NULL,
user_id INT REFERENCES users(id) NOT NULL
);

CREATE TABLE shippings (
id           SERIAL PRIMARY KEY,
user_id      INT REFERENCES users (id) NOT NULL,
phone_number VARCHAR(255)              not null,
address      VARCHAR(255)              not null,
house_number VARCHAR(255)              not null,
postal_code  VARCHAR(255)              not null,
city         VARCHAR(255)              not null,
notes        VARCHAR(255) DEFAULT ''   not null
);

CREATE TABLE special (
id SERIAL PRIMARY KEY,
name VARCHAR(100) not null
);