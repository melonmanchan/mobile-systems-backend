ALTER TABLE users RENAME first_name TO user_first_name;
ALTER TABLE users RENAME last_name TO user_last_name;
ALTER TABLE users RENAME password TO user_password;
ALTER TABLE users RENAME email TO user_email;
ALTER TABLE users RENAME id TO user_id;

ALTER TABLE users RENAME auth_method TO user_auth_method;
ALTER TABLE users RENAME user_type TO user_user_type;

ALTER TABLE user_types RENAME user_type_desc TO user_type_type;
ALTER TABLE user_types RENAME id TO user_type_id;

ALTER TABLE authentication_methods RENAME auth_method_desc TO authentication_method_type;
ALTER TABLE authentication_methods RENAME id TO authentication_method_id;

