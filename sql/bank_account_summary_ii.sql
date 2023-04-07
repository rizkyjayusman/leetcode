SELECT u.name as NAME, SUM(t.amount) as BALANCE FROM Transactions t LEFT JOIN Users u ON t.account = u.account GROUP BY t.account HAVING BALANCE > 10000;