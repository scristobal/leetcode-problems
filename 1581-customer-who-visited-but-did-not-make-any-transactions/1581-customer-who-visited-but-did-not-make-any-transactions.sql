# Write your MySQL query statement below
select customer_id, COUNT(*) as count_no_trans
from Visits left join Transactions on Visits.visit_id = Transactions.visit_id
group by customer_id, transaction_id
having Transactions.transaction_id is null