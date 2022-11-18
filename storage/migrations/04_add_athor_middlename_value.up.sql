BEGIN;

INSERT INTO author (id, firstname, middlename, lastname) VALUES ('3e1dfc06-dcf6-41fc-b3cc-7c0563fdfab4', 'John','Mr', 'Doe') ON CONFLICT DO NOTHING;
INSERT INTO author (id, firstname,middlename, lastname) VALUES ('24000e82-9c48-4297-a442-ecd1ad55791a','yorqin','Mr','Baqoyev') ON CONFLICT DO NOTHING;

INSERT INTO article (id, title, body, author_id) VALUES ('26e2aebc-9771-45ba-8577-ef1a2e7b4171', 'Lorem 4', 'Body 4', '3e1dfc06-dcf6-41fc-b3cc-7c0563fdfab4') ON CONFLICT DO NOTHING;
INSERT INTO article (id, title, body, author_id) VALUES ('9900756f-e3ed-4dd7-a3a8-4e3cef248cce', 'Lorem 5', 'Body 5', '24000e82-9c48-4297-a442-ecd1ad55791a') ON CONFLICT DO NOTHING;
INSERT INTO article (id, title, body, author_id) VALUES ('9900756f-e3ed-4dd7-a3a8-4e3cef248ccd', 'Lorem 6', 'Body 6', '24000e82-9c48-4297-a442-ecd1ad55791a') ON CONFLICT DO NOTHING;
INSERT INTO article (id, title, body, author_id) VALUES ('3e451dc4-42e8-4dbc-a70b-edee8f6452bb','Lorem 7','Body 7','3e1dfc06-dcf6-41fc-b3cc-7c0563fdfab4') ON CONFLICT DO NOTHING;

COMMIT;