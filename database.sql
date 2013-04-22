--The database-scheme of gofire
--the database is mainly used for storing user, passwords, sessions
CREATE TABLE gf_user(
	id BIGSERIAL PRIMARY KEY,
	login text,
	pw text,
	session text,
	mod INTEGER
);

INSERT INTO gf_user(
	login,
	pw,
	session,
	mod
	)VALUES(
		'admin',
		'61646d696ecf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e',
	'3161646d696ecf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e',
	0
);
