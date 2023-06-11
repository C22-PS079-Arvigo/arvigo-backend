-- auto-generated definition
INSERT INTO face_shapes (name, created_at, updated_at)
VALUES
    ('circle', DEFAULT, DEFAULT),
    ('heart', DEFAULT, DEFAULT),
    ('oblong', DEFAULT, DEFAULT),
    ('oval', DEFAULT, DEFAULT),
    ('square', DEFAULT, DEFAULT),
    ('triangle', DEFAULT, DEFAULT);

-- auto-generated definition
INSERT INTO categories (name, created_at, updated_at)
VALUES
    ('Glasses', DEFAULT, DEFAULT),
    ('Makeup', DEFAULT, DEFAULT);

-- auto-generated definition
INSERT INTO brands (name, category_id, created_at, updated_at)
VALUES
    ('CHANEL', 1, DEFAULT, DEFAULT),
    ('Oakley', 1, DEFAULT, DEFAULT),
    ('Ray-Ban', 1, DEFAULT, DEFAULT),
    ('Baleno', 1, DEFAULT, DEFAULT),
    ('Calvin Klein', 1, DEFAULT, DEFAULT),
    ('Emporio Armani', 1, DEFAULT, DEFAULT),
    ('Police', 1, DEFAULT, DEFAULT),
    ('Quiksilver', 1, DEFAULT, DEFAULT),
    ('Gucci', 1, DEFAULT, DEFAULT),
    ('Bottega Veneta', 1, DEFAULT, DEFAULT),
    
    ('Skintific', 2, DEFAULT, DEFAULT),
    ('Wardah', 2, DEFAULT, DEFAULT),
    ('Purbasari', 2, DEFAULT, DEFAULT),
    ('Somethinc', 2, DEFAULT, DEFAULT),
    ('Y.O.U', 2, DEFAULT, DEFAULT),
    ('Luxcrime', 2, DEFAULT, DEFAULT),
    ('Inez', 2, DEFAULT, DEFAULT),
    ('Viva Cosmetics', 2, DEFAULT, DEFAULT),
    ('Sariayu', 2, DEFAULT, DEFAULT),
    ('Emina', 2, DEFAULT, DEFAULT);


-- auto-generated definition
INSERT INTO tags (name, category_id, created_at, updated_at)
VALUES
    ('Aviator', 1, DEFAULT, DEFAULT),
    ('Cat Eye', 1, DEFAULT, DEFAULT),
    ('Square', 1, DEFAULT, DEFAULT),
    ('Oversized', 1, DEFAULT, DEFAULT),
    ('Rimless', 1, DEFAULT, DEFAULT),
    ('Round', 1, DEFAULT, DEFAULT),
    
    ('Red Shades', 2, DEFAULT, DEFAULT),
    ('Cluster', 2, DEFAULT, DEFAULT),
    ('Green', 2, DEFAULT, DEFAULT),

    ('Nude Shades', 2, DEFAULT, DEFAULT),
    ('Natural', 2, DEFAULT, DEFAULT),
    ('Blue', 2, DEFAULT, DEFAULT),

    ('Pink Shades', 2, DEFAULT, DEFAULT),
    ('Cat', 2, DEFAULT, DEFAULT),
    ('Brown', 2, DEFAULT, DEFAULT),

    ('Orange Shades', 2, DEFAULT, DEFAULT),
    ('Open', 2, DEFAULT, DEFAULT),
    ('Hazel', 2, DEFAULT, DEFAULT),

    ('Berry Shades', 2, DEFAULT, DEFAULT),
    ('Violet', 2, DEFAULT, DEFAULT);

-- auto-generated definition
INSERT INTO detail_face_shape_tags (face_shape_id, tag_id, created_at, updated_at)
VALUES
    (1, 3, DEFAULT, DEFAULT),
    (2, 1, DEFAULT, DEFAULT),
    (2, 5, DEFAULT, DEFAULT),
    (3, 4, DEFAULT, DEFAULT),
    (4, 1, DEFAULT, DEFAULT),
    (4, 2, DEFAULT, DEFAULT),
    (4, 3, DEFAULT, DEFAULT),
    (4, 4, DEFAULT, DEFAULT),
    (4, 5, DEFAULT, DEFAULT),
    (4, 6, DEFAULT, DEFAULT),
    (5, 2, DEFAULT, DEFAULT),
    (5, 6, DEFAULT, DEFAULT),
    (6, 1, DEFAULT, DEFAULT),
    (6, 5, DEFAULT, DEFAULT);

-- auto-generated definition
INSERT INTO roles (name, created_at, updated_at)
VALUES
    ('dashboard-app', DEFAULT, DEFAULT),
    ('user', DEFAULT, DEFAULT),
    ('merchant', DEFAULT, DEFAULT);

-- auto-generated definition
INSERT INTO personality_questionnaires (type, question, created_at, updated_at)
VALUES
    ('EXT1', 'Saya adalah orang yang paling ramai di pesta', DEFAULT, DEFAULT),
    ('EXT2', 'Saya tidak banyak bicara', DEFAULT, DEFAULT),
    ('EXT3', 'Saya merasa nyaman di sekitar orang lain', DEFAULT, DEFAULT),
    ('EXT4', 'Saya cenderung berada di latar belakang', DEFAULT, DEFAULT),
    ('EXT5', 'Saya memulai percakapan', DEFAULT, DEFAULT),
    ('EXT6', 'Saya jarang memiliki banyak kata-kata', DEFAULT, DEFAULT),
    ('EXT7', 'Saya berbicara dengan banyak orang yang berbeda di pesta', DEFAULT, DEFAULT),
    ('EXT8', 'Saya tidak suka menarik perhatian pada diri sendiri', DEFAULT, DEFAULT),
    ('EXT9', 'Saya tidak keberatan menjadi pusat perhatian', DEFAULT, DEFAULT),
    ('EXT10', 'Saya cenderung diam di sekitar orang asing', DEFAULT, DEFAULT),
    ('EST1', 'Saya mudah stres', DEFAULT, DEFAULT),
    ('EST2', 'Saya santai sebagian besar waktu', DEFAULT, DEFAULT),
    ('EST3', 'Saya khawatir tentang hal-hal', DEFAULT, DEFAULT),
    ('EST4', 'Saya jarang merasa sedih', DEFAULT, DEFAULT),
    ('EST5', 'Saya mudah terganggu', DEFAULT, DEFAULT),
    ('EST6', 'Saya mudah marah', DEFAULT, DEFAULT),
    ('EST7', 'Saya sering mengubah mood', DEFAULT, DEFAULT),
    ('EST8', 'Saya sering berganti mood', DEFAULT, DEFAULT),
    ('EST9', 'Saya mudah tersinggung', DEFAULT, DEFAULT),
    ('EST10', 'Saya sering merasa sedih', DEFAULT, DEFAULT),
    ('AGR1', 'Saya kurang peduli terhadap orang lain', DEFAULT, DEFAULT),
    ('AGR2', 'Saya tertarik pada orang lain', DEFAULT, DEFAULT),
    ('AGR3', 'Saya menghina orang lain', DEFAULT, DEFAULT),
    ('AGR4', 'Saya merasa simpati terhadap perasaan orang lain', DEFAULT, DEFAULT),
    ('AGR5', 'Saya tidak tertarik dengan masalah orang lain', DEFAULT, DEFAULT),
    ('AGR6', 'Saya memiliki hati yang lembut', DEFAULT, DEFAULT),
    ('AGR7', 'Saya tidak begitu tertarik dengan orang lain', DEFAULT, DEFAULT),
    ('AGR8', 'Saya meluangkan waktu untuk orang lain', DEFAULT, DEFAULT),
    ('AGR9', 'Saya merasakan emosi orang lain', DEFAULT, DEFAULT),
    ('AGR10', 'Saya membuat orang merasa nyaman', DEFAULT, DEFAULT),
    ('CSN1', 'Saya selalu siap', DEFAULT, DEFAULT),
    ('CSN2', 'Saya meninggalkan barang-barang saya di mana saja', DEFAULT, DEFAULT),
    ('CSN3', 'Saya memperhatikan detail', DEFAULT, DEFAULT),
    ('CSN4', 'Saya membuat kekacauan', DEFAULT, DEFAULT),
    ('CSN5', 'Saya menyelesaikan pekerjaan segera', DEFAULT, DEFAULT),
    ('CSN6', 'Saya sering lupa untuk meletakkan barang sesuai tempatnya', DEFAULT, DEFAULT),
    ('CSN7', 'Saya suka keteraturan', DEFAULT, DEFAULT),
    ('CSN8', 'Saya menghindari tanggung jawab saya', DEFAULT, DEFAULT),
    ('CSN9', 'Saya mengikuti jadwal', DEFAULT, DEFAULT),
    ('CSN10', 'Saya teliti dalam pekerjaan saya', DEFAULT, DEFAULT),
    ('OPN1', 'Saya memiliki kosa kata yang kaya', DEFAULT, DEFAULT),
    ('OPN2', 'Saya kesulitan memahami konsep abstrak', DEFAULT, DEFAULT),
    ('OPN3', 'Saya memiliki imajinasi yang kuat', DEFAULT, DEFAULT),
    ('OPN4', 'Saya tidak tertarik dengan konsep abstrak', DEFAULT, DEFAULT),
    ('OPN5', 'Saya memiliki ide-ide yang luar biasa', DEFAULT, DEFAULT),
    ('OPN6', 'Saya tidak memiliki imajinasi yang baik', DEFAULT, DEFAULT),
    ('OPN7', 'Saya cepat memahami hal-hal', DEFAULT, DEFAULT),
    ('OPN8', 'Saya menggunakan kata-kata yang sulit', DEFAULT, DEFAULT),
    ('OPN9', 'Saya menghabiskan waktu untuk merenungkan hal-hal', DEFAULT, DEFAULT),
    ('OPN10', 'Saya penuh dengan ide-ide', DEFAULT, DEFAULT);

INSERT INTO marketplaces (name, image, created_at, updated_at)
VALUES
    ('Website', "", DEFAULT, DEFAULT),
    ('Tokopedia', "", DEFAULT, DEFAULT),
    ('Shopee', "", DEFAULT, DEFAULT);

INSERT INTO users (email, password, role_id, full_name, gender, date_of_birth, place_of_birth,
                    is_complete_personality_test, is_complete_face_test, personality_id, face_shape_id,
                    is_verified, avatar, addresses_id, merchant_id, created_at, updated_at)
VALUES ('admin@gmail.com', '$2a$04$LuV7xCkK2l/La9qtQRo0nu3KAwefFppTYwiJPpS/4iypJrpYpQGmu', 1, 'Admin', '', null, '', 0,
        0, 0, 0, 0, '', 0, 0, '2023-06-03 07:14:18', '2023-06-03 07:14:56');