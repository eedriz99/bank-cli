CREATE TABLE
    NOT EXISTS accounts (
        -- id INT AUTO_INCREMENT,
        name TEXT NOT NULL,
        phone TEXT NOT NULL,
        acct_no TEXT NOT NULL,
        acct_type TEXT NOT NULL,
        bal REAL,
        lim REAL,
        PRIMARY KEY (acct_no)
    );

CREATE TABLE
    NOT EXISTS transactions (
        -- id INT AUTO_INCREMENT,
        acct TEXT NOT NULL,
        src TEXT NOT NULL,
        dest TEXT NOT NULL,
        amount REAL NOT NULL,
        is_success INTEGER,
        created_at TEXT DEFAULT (STRFTIME ('%Y-%m-%d %H:%M:%S', 'now')),
        FOREIGN KEY (acct) REFERENCES accounts(acct_no) ON DELETE CASCADE ON UPDATE CASCADE
    );