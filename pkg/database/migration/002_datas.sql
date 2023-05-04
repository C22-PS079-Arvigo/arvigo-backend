-- auto-generated definition
DROP TABLE IF EXISTS addresses;
CREATE TABLE addresses
(
    id int unsigned auto_increment primary key,
    user_id int not null,
    street text not null,
    province_id int not null,
    district_id int not null,
    subdistrict_id int not null,
    postal_code_id int not null,
    created_at timestamp default CURRENT_TIMESTAMP null,
    updated_at timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP
);
CREATE INDEX idx_addresses_1 ON addresses (user_id);
CREATE INDEX idx_addresses_2 ON addresses (province_id, district_id, subdistrict_id, postal_code_id);


-- auto-generated definition
DROP TABLE IF EXISTS users;
CREATE TABLE users
(
    id int unsigned auto_increment primary key,
    email varchar(255) not null,
    password varchar(255) not null,
    role_id int not null,
    full_name varchar(255) not null,
    gender varchar(10) null,
    date_of_birth date null,
    place_of_birth varchar(50) null,
    is_complete_personality_test int default 0 not null,
    is_complete_face_test int default 0 not null,
    personality_tag_id int null,
    face_shape_tag_id int null,
    is_verified int default 0 not null,
    avatar varchar(200) null,
    addresses_id int null,
    merchant_id int null,
    created_at timestamp default CURRENT_TIMESTAMP null,
    updated_at timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP,
    constraint email unique (email)
);
CREATE INDEX idx_users_1 ON users (addresses_id);
CREATE INDEX idx_users_2 ON users (merchant_id);
CREATE INDEX idx_users_3 ON users (personality_tag_id);
CREATE INDEX idx_users_4 ON users (face_shape_tag_id);


-- auto-generated definition
DROP TABLE IF EXISTS roles;
CREATE TABLE roles
(
    id int unsigned auto_increment primary key,
    name varchar(50) not null,
    created_at timestamp default CURRENT_TIMESTAMP null,
    updated_at timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP
);

-- auto-generated definition
DROP TABLE IF EXISTS categories;
CREATE TABLE categories
(
    id int unsigned auto_increment primary key,
    name varchar(50) not null,
    created_at timestamp default CURRENT_TIMESTAMP null,
    updated_at timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP
);

-- auto-generated definition
DROP TABLE IF EXISTS brand;
CREATE TABLE brand
(
    id int unsigned auto_increment primary key,
    name varchar(50) not null,
    created_at timestamp default CURRENT_TIMESTAMP null,
    updated_at timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP
);

-- auto-generated definition
DROP TABLE IF EXISTS marketplaces;
CREATE TABLE marketplaces
(
    id int unsigned auto_increment primary key,
    name varchar(50) not null,
    image varchar(200) not null,
    created_at timestamp default CURRENT_TIMESTAMP null,
    updated_at timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP
);

-- auto-generated definition
DROP TABLE IF EXISTS merchants;
CREATE TABLE merchants
(
    id int unsigned auto_increment primary key,
    name varchar(50) not null,
    is_recomendation int default 0 not null,
    addresses_id int default 0 not null,
    created_at timestamp default CURRENT_TIMESTAMP null,
    updated_at timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP
);
CREATE INDEX idx_merchants_1 ON merchants (addresses_id);



-- auto-generated definition
DROP TABLE IF EXISTS products;
CREATE TABLE products
(
    id int unsigned auto_increment primary key,
    name varchar(200) not null,
    description text null,
    images text null,
    link_external varchar(100) not null,
    merchant_id int not null,
    created_at timestamp default CURRENT_TIMESTAMP null,
    updated_at timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP
);
CREATE INDEX idx_products_1 ON products (merchant_id);


-- auto-generated definition
DROP TABLE IF EXISTS detail_product_categories;
CREATE TABLE detail_product_categories
(
    id int unsigned auto_increment primary key,
    product_id int not null,
    category_id int not null,
    created_at timestamp default CURRENT_TIMESTAMP null,
    updated_at timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP
);
CREATE INDEX idx_detail_product_categories_1 ON detail_product_categories (product_id);
CREATE INDEX idx_detail_product_categories_2 ON detail_product_categories (category_id);


-- auto-generated definition
DROP TABLE IF EXISTS detail_product_reviews;
CREATE TABLE detail_product_reviews
(
    id int unsigned auto_increment primary key,
    product_id int not null,
    user_id int not null,
    comment varchar(200) not null,
    rating double default 0 not null,
    images text null,
    created_at timestamp default CURRENT_TIMESTAMP null,
    updated_at timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP
);
CREATE INDEX idx_detail_product_reviews_1 ON detail_product_reviews (product_id);
CREATE INDEX idx_detail_product_reviews_2 ON detail_product_reviews (user_id);


-- auto-generated definition
DROP TABLE IF EXISTS detail_product_variants;
CREATE TABLE detail_product_variants
(
    id int unsigned auto_increment primary key,
    name varchar(200) not null,
    link_ar varchar(100) not null,
    is_primary_variant int default 0 not null,
    product_id int not null,
    created_at timestamp default CURRENT_TIMESTAMP null,
    updated_at timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP
);
CREATE INDEX idx_detail_product_variants_1 ON detail_product_variants (product_id);


-- auto-generated definition
DROP TABLE IF EXISTS tags;
CREATE TABLE tags
(
    id int unsigned auto_increment primary key,
    name varchar(200) not null,
    type varchar(50) not null,
    created_at timestamp default CURRENT_TIMESTAMP null,
    updated_at timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP
);
CREATE INDEX idx_tags_1 ON tags (type);


-- auto-generated definition
DROP TABLE IF EXISTS detail_product_tags;
CREATE TABLE detail_product_tags
(
    id int unsigned auto_increment primary key,
    tag_id varchar(200) not null,
    product_id int not null,
    created_at timestamp default CURRENT_TIMESTAMP null,
    updated_at timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP
);
CREATE INDEX idx_detail_product_tags_1 ON detail_product_tags (tag_id);
CREATE INDEX idx_detail_product_tags_2 ON detail_product_tags (product_id);


-- auto-generated definition
DROP TABLE IF EXISTS detail_product_marketplaces;
CREATE TABLE detail_product_marketplaces
(
    id int unsigned auto_increment primary key,
    marketplace_id int not null,
    product_id int not null,
    link varchar(100) not null,
    clicked int default 0 not null,
    created_at timestamp default CURRENT_TIMESTAMP null,
    updated_at timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP
);
CREATE INDEX idx_detail_product_marketplaces_1 ON detail_product_marketplaces (product_id);
CREATE INDEX idx_detail_product_marketplaces_2 ON detail_product_marketplaces (marketplace_id);


-- auto-generated definition
DROP TABLE IF EXISTS wishlists;
CREATE TABLE wishlists
(
    id int unsigned auto_increment primary key,
    user_id int not null,
    product_id int not null,
    created_at timestamp default CURRENT_TIMESTAMP null,
    updated_at timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP
);

CREATE INDEX idx_wishlists_1 ON wishlists (product_id);
CREATE INDEX idx_wishlists_2 ON wishlists (user_id);

-- auto-generated definition
DROP TABLE IF EXISTS detail_user_subscriptions;
CREATE TABLE detail_user_subscriptions
(
    id int unsigned auto_increment primary key,
    user_id int not null,
    price int not null,
    unique_code int not null,
    subscription_start timestamp null,
    subscription_end timestamp null,
    is_verified int default 0 not null,
    message varchar(100) null,
    paid_at timestamp null,
    created_at timestamp default CURRENT_TIMESTAMP null,
    updated_at timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP
);

CREATE INDEX idx_detail_user_subscriptions_1 ON detail_user_subscriptions (user_id);
