CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users (
    user_id   UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    msisdn    VARCHAR(20) NOT NULL,
    name      VARCHAR(100) NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE
);

-- Test uchun 10 ta user qo'shamiz
INSERT INTO users (msisdn, name, is_active) VALUES
('+998901112233', 'Ali Valiyev', TRUE),
('+998901112234', 'Vali Karimov', TRUE),
('+998901112235', 'Sardor Tursunov', FALSE),
('+998901112236', 'Dilshod Raxmonov', TRUE),
('+998901112237', 'Azizbek Salimov', TRUE),
('+998901112238', 'Madina Qodirova', TRUE),
('+998901112239', 'Shahnoza Karimova', FALSE),
('+998901112240', 'Javohir Islomov', TRUE),
('+998901112241', 'Nigora Abduqodirova', TRUE),
('+998901112242', 'Sherzod Mamatov', TRUE);
