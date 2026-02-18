-- Create some initial accounts
INSERT INTO accounts (owner, balance, currency) VALUES ('serge', 1000, 'USD');
INSERT INTO accounts (owner, balance, currency) VALUES ('apple', 500, 'USD');
INSERT INTO accounts (owner, balance, currency) VALUES ('google', 200, 'USD');

-- Create some history
INSERT INTO entries (account_id, amount) VALUES (1, 100);
INSERT INTO transfers (from_account_id, to_account_id, amount) VALUES (1, 2, 50);