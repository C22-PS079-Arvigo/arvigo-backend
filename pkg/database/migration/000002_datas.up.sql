-- auto-generated definition
DROP TABLE IF EXISTS addresses;
CREATE TABLE addresses
(
    id int unsigned auto_increment primary key,
    street text not null,
    province_id int not null,
    city_id int not null,
    district_id int not null,
    subdistrict_id int not null,
    postal_code_id int not null,
    created_at timestamp default CURRENT_TIMESTAMP null,
    updated_at timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP
);
CREATE INDEX idx_addresses_1 ON addresses (province_id, district_id, subdistrict_id, postal_code_id);


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
    personality_id int null,
    face_shape_id int null,
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
CREATE INDEX idx_users_3 ON users (personality_id);
CREATE INDEX idx_users_4 ON users (face_shape_id);


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
DROP TABLE IF EXISTS brands;
CREATE TABLE brands
(
    id int unsigned auto_increment primary key,
    name varchar(50) not null,
    category_id int not null,
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
    created_at timestamp default CURRENT_TIMESTAMP null,
    updated_at timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP
);
CREATE INDEX idx_merchants_1 ON merchants (is_recomendation);



-- auto-generated definition
DROP TABLE IF EXISTS products;
CREATE TABLE products
(
    id int unsigned auto_increment primary key,
    name varchar(200) not null,
    description text null,
    images text null,
    link_external varchar(100) not null,
    category_id int not null,
    brand_id int not null,
    merchant_id int default 0 not null, -- set 0 for product admin
    is_verified tinyint(1) default 0 not null,
    is_subscription_active tinyint(1) default 0 not null,
    created_at timestamp default CURRENT_TIMESTAMP null,
    updated_at timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP
);
CREATE INDEX idx_products_1 ON products (merchant_id);
CREATE INDEX idx_products_2 ON products (merchant_id, brand_id, category_id);

-- auto-generated definition
DROP TABLE IF EXISTS detail_linked_products;
CREATE TABLE detail_linked_products
(
    id int unsigned auto_increment primary key,
    initial_product_id int not null,
    merchant_product_id int not null,
    merchant_id int default 0 not null,
    created_at timestamp default CURRENT_TIMESTAMP null,
    updated_at timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP
);
CREATE INDEX idx_detail_linked_products_1 ON detail_linked_products (initial_product_id, merchant_product_id, merchant_id);

-- -- auto-generated definition
-- DROP TABLE IF EXISTS detail_product_categories;
-- CREATE TABLE detail_product_categories
-- (
--     id int unsigned auto_increment primary key,
--     product_id int not null,
--     category_id int not null,
--     created_at timestamp default CURRENT_TIMESTAMP null,
--     updated_at timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP
-- );
-- CREATE INDEX idx_detail_product_categories_1 ON detail_product_categories (product_id, category_id);


-- -- auto-generated definition
-- DROP TABLE IF EXISTS detail_product_brands;
-- CREATE TABLE detail_product_brands
-- (
--     id int unsigned auto_increment primary key,
--     product_id int not null,
--     brand_id int not null,
--     created_at timestamp default CURRENT_TIMESTAMP null,
--     updated_at timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP
-- );
-- CREATE INDEX idx_detail_product_brands_1 ON detail_product_brands (product_id, brand_id);

-- auto-generated definition
-- DROP TABLE IF EXISTS detail_product_reviews;
-- CREATE TABLE detail_product_reviews
-- (
--     id int unsigned auto_increment primary key,
--     product_id int not null,
--     user_id int not null,
--     comment varchar(200) not null,
--     rating double default 0 not null,
--     images text null,
--     created_at timestamp default CURRENT_TIMESTAMP null,
--     updated_at timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP
-- );
-- CREATE INDEX idx_detail_product_reviews_1 ON detail_product_reviews (product_id);
-- CREATE INDEX idx_detail_product_reviews_2 ON detail_product_reviews (user_id);


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
    category_id int not null,
    created_at timestamp default CURRENT_TIMESTAMP null,
    updated_at timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP
);
CREATE INDEX idx_tags_1 ON tags (category_id);


-- auto-generated definition
DROP TABLE IF EXISTS face_shapes;
CREATE TABLE face_shapes
(
    id int unsigned auto_increment primary key,
    name varchar(50) not null,
    created_at timestamp default CURRENT_TIMESTAMP null,
    updated_at timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP
);

-- auto-generated definition
DROP TABLE IF EXISTS detail_face_shape_tags;
CREATE TABLE detail_face_shape_tags
(
    id int unsigned auto_increment primary key,
    face_shape_id int not null,
    tag_id int not null,
    created_at timestamp default CURRENT_TIMESTAMP null,
    updated_at timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP
);

-- auto-generated definition
DROP TABLE IF EXISTS detail_product_tags;
CREATE TABLE detail_product_tags
(
    id int unsigned auto_increment primary key,
    tag_id int not null,
    product_id int not null,
    created_at timestamp default CURRENT_TIMESTAMP null,
    updated_at timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP
);
CREATE INDEX idx_detail_product_tags_1 ON detail_product_tags (product_id, tag_id);


-- auto-generated definition
DROP TABLE IF EXISTS detail_product_marketplaces;
CREATE TABLE detail_product_marketplaces
(
    id int unsigned auto_increment primary key,
    marketplace_id  int default 0 not null,
    product_id int not null,
    addresses_id int null,
    link varchar(100) null,
    clicked int default 0 not null,
    created_at timestamp default CURRENT_TIMESTAMP null,
    updated_at timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP
);
CREATE INDEX idx_detail_product_marketplaces_1 ON detail_product_marketplaces (product_id, marketplace_id);
CREATE INDEX idx_detail_product_marketplaces_2 ON detail_product_marketplaces (product_id, addresses_id);


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

CREATE INDEX idx_wishlists_1 ON wishlists (user_id, product_id);

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


-- auto-generated definition
DROP TABLE IF EXISTS personality_questionnaires;
CREATE TABLE personality_questionnaires (
    id INT PRIMARY KEY AUTO_INCREMENT,
    type VARCHAR(50) NOT NULL,
    question TEXT NOT NULL,
    created_at timestamp default CURRENT_TIMESTAMP null,
    updated_at timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP
);


-- auto-generated definition
DROP TABLE IF EXISTS user_personalities;
CREATE TABLE user_personalities (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    is_active INT DEFAULT 0 NOT NULL,
    ext_result int null,
    est_result int null,
    agr_result int null,
    csn_result int null,
    _result int null,
    EXT1 INT NOT NULL,EXT2 INT NOT NULL,EXT3 INT NOT NULL,EXT4 INT NOT NULL,EXT5 INT NOT NULL,EXT6 INT NOT NULL,EXT7 INT NOT NULL,EXT8 INT NOT NULL,EXT9 INT NOT NULL,EXT10 INT NOT NULL,EST1 INT NOT NULL,EST2 INT NOT NULL,EST3 INT NOT NULL,EST4 INT NOT NULL,EST5 INT NOT NULL,EST6 INT NOT NULL,EST7 INT NOT NULL,EST8 INT NOT NULL,EST9 INT NOT NULL,EST10 INT NOT NULL,AGR1 INT NOT NULL,AGR2 INT NOT NULL,AGR3 INT NOT NULL,AGR4 INT NOT NULL,AGR5 INT NOT NULL,AGR6 INT NOT NULL,AGR7 INT NOT NULL,AGR8 INT NOT NULL,AGR9 INT NOT NULL,AGR10 INT NOT NULL,CSN1 INT NOT NULL,CSN2 INT NOT NULL,CSN3 INT NOT NULL,CSN4 INT NOT NULL,CSN5 INT NOT NULL,CSN6 INT NOT NULL,CSN7 INT NOT NULL,CSN8 INT NOT NULL,CSN9 INT NOT NULL,CSN10 INT NOT NULL,OPN1 INT NOT NULL,OPN2 INT NOT NULL,OPN3 INT NOT NULL,OPN4 INT NOT NULL,OPN5 INT NOT NULL,OPN6 INT NOT NULL,OPN7 INT NOT NULL,OPN8 INT NOT NULL,OPN9 INT NOT NULL,OPN10 INT NOT NULL,
    created_at timestamp default CURRENT_TIMESTAMP null,
    updated_at timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP
);
CREATE INDEX idx_user_personalities_1 ON user_personalities (user_id, is_active);


-- auto-generated definition
DROP TABLE IF EXISTS user_personalities;
CREATE TABLE user_personalities (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    is_active INT DEFAULT 0 NOT NULL,
    ext_result int null,
    est_result int null,
    agr_result int null,
    csn_result int null,
    opn_result int null,
    EXT1 INT NOT NULL,EXT2 INT NOT NULL,EXT3 INT NOT NULL,EXT4 INT NOT NULL,EXT5 INT NOT NULL,EXT6 INT NOT NULL,EXT7 INT NOT NULL,EXT8 INT NOT NULL,EXT9 INT NOT NULL,EXT10 INT NOT NULL,EST1 INT NOT NULL,EST2 INT NOT NULL,EST3 INT NOT NULL,EST4 INT NOT NULL,EST5 INT NOT NULL,EST6 INT NOT NULL,EST7 INT NOT NULL,EST8 INT NOT NULL,EST9 INT NOT NULL,EST10 INT NOT NULL,AGR1 INT NOT NULL,AGR2 INT NOT NULL,AGR3 INT NOT NULL,AGR4 INT NOT NULL,AGR5 INT NOT NULL,AGR6 INT NOT NULL,AGR7 INT NOT NULL,AGR8 INT NOT NULL,AGR9 INT NOT NULL,AGR10 INT NOT NULL,CSN1 INT NOT NULL,CSN2 INT NOT NULL,CSN3 INT NOT NULL,CSN4 INT NOT NULL,CSN5 INT NOT NULL,CSN6 INT NOT NULL,CSN7 INT NOT NULL,CSN8 INT NOT NULL,CSN9 INT NOT NULL,CSN10 INT NOT NULL,OPN1 INT NOT NULL,OPN2 INT NOT NULL,OPN3 INT NOT NULL,OPN4 INT NOT NULL,OPN5 INT NOT NULL,OPN6 INT NOT NULL,OPN7 INT NOT NULL,OPN8 INT NOT NULL,OPN9 INT NOT NULL,OPN10 INT NOT NULL,
    created_at timestamp default CURRENT_TIMESTAMP null,
    updated_at timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP
);
CREATE INDEX idx_user_personalities_1 ON user_personalities (user_id, is_active);
