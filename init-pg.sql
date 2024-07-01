CREATE TABLE IF NOT EXISTS movies (
  id  SERIAL PRIMARY KEY not null,
  title TEXT,
  url TEXT,
  release_date DATE not null default CURRENT_DATE
);

INSERT INTO movies (title, url) VALUES ('Fast and Furious 21', 'https://fnf.com/21/'); 
INSERT INTO movies (title, url) VALUES ('Fast and Furious 17', 'https://fnf.com/17/'); 
INSERT INTO movies (title, url) VALUES ('Fall Guy', 'https://fallguy.com/'); 
INSERT INTO movies (title, url) VALUES ('Rise of GraphQL Warrior Pt1', 'https://riseofgraphqlwarriorpt1.com/'); 