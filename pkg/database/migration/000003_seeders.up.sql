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
    
    ('Foundation', 2, DEFAULT, DEFAULT),
    ('Concealer', 2, DEFAULT, DEFAULT),
    ('Powder', 2, DEFAULT, DEFAULT),
    ('Blush', 2, DEFAULT, DEFAULT),
    ('Highlighter', 2, DEFAULT, DEFAULT),
    ('Bronzer', 2, DEFAULT, DEFAULT),
    ('Eyeshadow', 2, DEFAULT, DEFAULT),
    ('Eyeliner', 2, DEFAULT, DEFAULT),
    ('Mascara', 2, DEFAULT, DEFAULT),
    ('Lipstick', 2, DEFAULT, DEFAULT),
    ('Lip Gloss', 2, DEFAULT, DEFAULT),
    ('Lip Liner', 2, DEFAULT, DEFAULT),
    ('Eyebrow', 2, DEFAULT, DEFAULT),
    ('Setting Spray', 2, DEFAULT, DEFAULT);

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
    ('EXT1', 'I am the life of the party', DEFAULT, DEFAULT),
    ('EXT2', 'I dont talk a lot', DEFAULT, DEFAULT),
    ('EXT3', 'I feel comfortable around people', DEFAULT, DEFAULT),
    ('EXT4', 'I keep in the background', DEFAULT, DEFAULT),
    ('EXT5', 'I start conversations', DEFAULT, DEFAULT),
    ('EXT6', 'I have little to say', DEFAULT, DEFAULT),
    ('EXT7', 'I talk to a lot of different people at parties', DEFAULT, DEFAULT),
    ('EXT8', 'I dont like to draw attention to myself', DEFAULT, DEFAULT),
    ('EXT9', 'I dont mind being the center of attention', DEFAULT, DEFAULT),
    ('EXT10', 'I am quiet around strangers', DEFAULT, DEFAULT),
    ('EST1', 'I get stressed out easily', DEFAULT, DEFAULT),
    ('EST2', 'I am relaxed most of the time', DEFAULT, DEFAULT),
    ('EST3', 'I worry about things', DEFAULT, DEFAULT),
    ('EST4', 'I seldom feel blue', DEFAULT, DEFAULT),
    ('EST5', 'I am easily disturbed', DEFAULT, DEFAULT),
    ('EST6', 'I get upset easily', DEFAULT, DEFAULT),
    ('EST7', 'I change my mood a lot', DEFAULT, DEFAULT),
    ('EST8', 'I have frequent mood swings', DEFAULT, DEFAULT),
    ('EST9', 'I get irritated easily', DEFAULT, DEFAULT),
    ('EST10', 'I often feel blue', DEFAULT, DEFAULT),
    ('AGR1', 'I feel little concern for others', DEFAULT, DEFAULT),
    ('AGR2', 'I am interested in people', DEFAULT, DEFAULT),
    ('AGR3', 'I insult people', DEFAULT, DEFAULT),
    ('AGR4', 'I sympathize with others feelings', DEFAULT, DEFAULT),
    ('AGR5', 'I am not interested in other peoples problems', DEFAULT, DEFAULT),
    ('AGR6', 'I have a soft heart', DEFAULT, DEFAULT),
    ('AGR7', 'I am not really interested in others', DEFAULT, DEFAULT),
    ('AGR8', 'I take time out for others', DEFAULT, DEFAULT),
    ('AGR9', 'I feel others emotions', DEFAULT, DEFAULT),
    ('AGR10', 'I make people feel at ease', DEFAULT, DEFAULT),
    ('CSN1', 'I am always prepared', DEFAULT, DEFAULT),
    ('CSN2', 'I leave my belongings around', DEFAULT, DEFAULT),
    ('CSN3', 'I pay attention to details', DEFAULT, DEFAULT),
    ('CSN4', 'I make a mess of things', DEFAULT, DEFAULT),
    ('CSN5', 'I get chores done right away', DEFAULT, DEFAULT),
    ('CSN6', 'I often forget to put things back in their proper place', DEFAULT, DEFAULT),
    ('CSN7', 'I like order', DEFAULT, DEFAULT),
    ('CSN8', 'I shirk my duties', DEFAULT, DEFAULT),
    ('CSN9', 'I follow a schedule', DEFAULT, DEFAULT),
    ('CSN10', 'I am exacting in my work', DEFAULT, DEFAULT),
    ('OPN1', 'I have a rich vocabulary', DEFAULT, DEFAULT),
    ('OPN2', 'I have difficulty understanding abstract ideas', DEFAULT, DEFAULT),
    ('OPN3', 'I have a vivid imagination', DEFAULT, DEFAULT),
    ('OPN4', 'I am not interested in abstract ideas', DEFAULT, DEFAULT),
    ('OPN5', 'I have excellent ideas', DEFAULT, DEFAULT),
    ('OPN6', 'I do not have a good imagination', DEFAULT, DEFAULT),
    ('OPN7', 'I am quick to understand things', DEFAULT, DEFAULT),
    ('OPN8', 'I use difficult words', DEFAULT, DEFAULT),
    ('OPN9', 'I spend time reflecting on things', DEFAULT, DEFAULT),
    ('OPN10', 'I am full of ideas', DEFAULT, DEFAULT);

INSERT INTO marketplaces (name, image, created_at, updated_at)
VALUES
    ('Website', "", DEFAULT, DEFAULT),
    ('Tokopedia', "", DEFAULT, DEFAULT),
    ('Shopee', "", DEFAULT, DEFAULT);