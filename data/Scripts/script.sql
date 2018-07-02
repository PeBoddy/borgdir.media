CREATE TABLE `Items` (
	`ItemID`	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	`Bezeichnung`	varchar ( 50 ) NOT NULL,
	`Kategorie`	varchar ( 20 ) NOT NULL,
	`InventarNummer`	INTEGER NOT NULL,
	`Lagerort`	varchar ( 30 ) NOT NULL,
	`Anzahl`	INTEGER,
	`Hinweis`	varchar ( 100 ),
	`BildURL`	varchar ( 100 ),
	`Status`	varchar ( 100 ),
	`Inhalt`	varchar ( 100 ),
	`AusgeliehenAm`	varchar ( 20 ),
	`RueckgabeAm`	varchar ( 20 ),
	`Entliehen`	INTEGER
);

CREATE TABLE `User` (
	`UserID`	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	`Name`	varchar ( 20 ) NOT NULL,
	`BildUrl`	varchar ( 30 ),
	`Typ`	varchar ( 20 ),
	`Status`	varchar ( 20 ),
	`Passwort`	varchar ( 100 ) NOT NULL,
	`Email`	varchar ( 30 ) NOT NULL,
	`PasswortKlar`	varchar ( 100 )
);

CREATE TABLE `Lend` (
	`LendID`	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	`UserID`	INTEGER NOT NULL,
	`ItemID`	INTEGER NOT NULL
);
