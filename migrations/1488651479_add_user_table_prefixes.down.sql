ALTER TABLE users RENAME user_first_name TO first_name;
ALTER TABLE users RENAME user_last_name TO last_name;
ALTER TABLE users RENAME user_password TO password;
ALTER TABLE users RENAME user_email TO email;
ALTER TABLE users RENAME user_id TO id;

ALTER TABLE users RENAME user_auth_method TO auth_method;
ALTER TABLE users RENAME user_user_type TO user_type;

ALTER TABLE user_types RENAME user_type_type TO user_type_desc;_
ALTER TABLE user_types RENAME user_type_id TO id;

ALTER TABLE authentication_methods RENAME authentication_method_type TO auth_method_desc; 
ALTER TABLE authentication_methods RENAME authentication_method_id TO id;

