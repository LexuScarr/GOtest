-- Создание таблицы news
CREATE TABLE IF NOT EXISTS news
(
    id      BIGSERIAL PRIMARY KEY,
    title   VARCHAR(128) NOT NULL,
    content TEXT         NOT NULL
    );

-- Создание таблицы news_categories
CREATE TABLE IF NOT EXISTS news_categories
(
    id          BIGSERIAL PRIMARY KEY,
    news_id     BIGINT                NOT NULL,
    category_id BIGINT                NOT NULL,
    CONSTRAINT news_categories_pk UNIQUE (news_id, category_id),
    FOREIGN KEY (news_id) REFERENCES news (id) ON DELETE CASCADE
    );

-- Создание таблицы categories
CREATE TABLE IF NOT EXISTS categories
(
    id   BIGSERIAL PRIMARY KEY,
    name VARCHAR(128) NOT NULL
    );

-- Вставка тестовых данных в таблицу news
INSERT INTO news (title, content) VALUES ('news 1','description 1');
INSERT INTO news (title, content) VALUES ('news 2','2. description 2');
INSERT INTO news (title, content) VALUES ('news 3','3. description 3');
INSERT INTO news (title, content) VALUES ('news 4','4. description 4');
INSERT INTO news (title, content) VALUES ('news 5','5. description 5');

-- Вставка тестовых данных в таблицу categories
INSERT INTO categories (name) VALUES ('category 1');
INSERT INTO categories (name) VALUES ('category 2');
INSERT INTO categories (name) VALUES ('category 3');
INSERT INTO categories (name) VALUES ('category 4');

-- Вставка данных в таблицу news_categories для связывания новостей с категориями
INSERT INTO news_categories (news_id, category_id) VALUES (1, 1);
INSERT INTO news_categories (news_id, category_id) VALUES (1, 2);
INSERT INTO news_categories (news_id, category_id) VALUES (1, 3);
INSERT INTO news_categories (news_id, category_id) VALUES (2, 1);
INSERT INTO news_categories (news_id, category_id) VALUES (2, 2);
INSERT INTO news_categories (news_id, category_id) VALUES (3, 1);
INSERT INTO news_categories (news_id, category_id) VALUES (4, 1);
INSERT INTO news_categories (news_id, category_id) VALUES (5, 1);
INSERT INTO news_categories (news_id, category_id) VALUES (5, 2);
INSERT INTO news_categories (news_id, category_id) VALUES (5, 3);
INSERT INTO news_categories (news_id, category_id) VALUES (5, 4);
