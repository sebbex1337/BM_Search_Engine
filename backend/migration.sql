-- Insert existing data from pages table into pages_fts table
INSERT INTO pages_fts (title, content, language)
SELECT title, content, language FROM pages;