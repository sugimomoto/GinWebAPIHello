drop table posts;

CREATE TABLE `sample`.`tasks` (
  `Id` INT NOT NULL AUTO_INCREMENT,
  `Subject` VARCHAR(200) NOT NULL,
  `Priority` VARCHAR(200) NULL,
  `CreatedAt` TIMESTAMP NOT NULL,
  PRIMARY KEY (`Id`));
);

