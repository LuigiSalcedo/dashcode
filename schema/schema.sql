CREATE DATABASE dashcode;
USE dashcode;

CREATE TABLE USERS (
    ID INTEGER PRIMARY KEY NOT NULL, NAME VARCHAR(90) NOT NULL
);

CREATE TABLE LOGIN (
    ID INTEGER NOT NULL, EMAIL VARCHAR(90) PRIMARY KEY NOT NULL, PASSWORD VARCHAR(100) NOT NULL, 
    FOREIGN KEY(ID) REFERENCES USERS(ID)
);

CREATE TABLE GROUPS(
    ID INTEGER AUTO_INCREMENT, ID_CREATOR INTEGER, NAME VARCHAR(50) NOT NULL, DESCRIPTION VARCHAR(100),
    FOREIGN KEY(ID_CREATOR) REFERENCES USERS(ID)
);

CREATE TABLE GROUP_MEMBERS (
	ID_GROUP INTEGER,
	ID_USER INTEGER,
	FOREIGN KEY(ID_GROUP) REFERENCES groups(ID),
	FOREIGN KEY(ID_USER) REFERENCES users(ID)
);