ALTER TABLE IF EXISTS refresh_tokens DROP CONSTRAINT IF EXISTS refresh_tokens_customer_id_fkey;
DROP TABLE IF EXISTS refresh_tokens;
DROP TABLE IF EXISTS customers;

ALTER TABLE IF EXISTS admin_users DROP CONSTRAINT IF EXISTS admin_users_role_id_fkey;
DROP TABLE IF EXISTS admin_users;

ALTER TABLE IF EXISTS authorization_rules DROP CONSTRAINT IF EXISTS authorization_rules_role_id_fkey;
DROP TABLE IF EXISTS authorization_rules;

DROP TABLE IF EXISTS authorization_roles;

DROP TYPE gender;
