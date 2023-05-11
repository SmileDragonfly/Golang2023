DROP TABLE IF EXISTS Wallet;
CREATE TABLE Wallet (
                        ID BIGSERIAL PRIMARY KEY,
                        CreateTime timestamp default CURRENT_TIMESTAMP,
                        UserName varchar(256) NOT NULL ,
                        Password varchar(128) NOT NULL ,
                        Mobile varchar(32) NOT NULL
);
