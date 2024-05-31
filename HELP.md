# sqlc did not support trigger on mysql 
if the schema contains below trigger syntax, it will reports 
```shell
syntax error near "DELIMITER ;;"
```

```mysql
DELIMITER ;;
CREATE TRIGGER generate_email_for_user
    BEFORE INSERT
    ON `goadmin_users`
    FOR EACH ROW
BEGIN
    IF NEW.email IS NULL THEN
        SET NEW.email = CONCAT(
                CHAR(ROUND(RAND() * 25) + 97),
                CHAR(ROUND(RAND() * 25) + 97),
                CHAR(ROUND(RAND() * 25) + 97),
                CHAR(ROUND(RAND() * 25) + 97),
                CHAR(ROUND(RAND() * 25) + 97),
                CHAR(ROUND(RAND() * 25) + 97),
                CHAR(ROUND(RAND() * 25) + 97),
                CHAR(ROUND(RAND() * 25) + 97),
                CHAR(ROUND(RAND() * 25) + 97),
                '@',
                CHAR(ROUND(RAND() * 25) + 97),
                CHAR(ROUND(RAND() * 25) + 97),
                CHAR(ROUND(RAND() * 25) + 97),
                CHAR(ROUND(RAND() * 25) + 97),
                CHAR(ROUND(RAND() * 25) + 97),
                CHAR(ROUND(RAND() * 25) + 97),
                '.com'
                        );
    END IF;
END ;;
DELIMITER ;;
```