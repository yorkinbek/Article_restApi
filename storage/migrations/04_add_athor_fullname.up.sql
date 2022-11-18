AlTER TABLE author ADD COLUMN fullname VARCHAR(100);
UPDATE author SET fullname= firstname || ' ' || lastname ||  
(SELECT CASE WHEN middlename IS NULL THEN '' ELSE ' ' || middlename END AS middlename);